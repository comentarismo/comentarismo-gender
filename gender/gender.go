package gender

import (
	"comentarismo-gender/lang"
	"encoding/csv"
	"encoding/gob"
	redis "gopkg.in/redis.v3"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	English_ignore_words_map    = make(map[string]int)
	Portuguese_ignore_words_map = make(map[string]int)
	Spanish_ignore_words_map    = make(map[string]int)
	Italian_ignore_words_map    = make(map[string]int)
	French_ignore_words_map     = make(map[string]int)

	RedisClient  *redis.Client
	Redis_prefix = "genderbayes:"
	correction   = 0.1
)

type GenderReport struct {
	Code   int    `json:"code"`
	Error  string `json:"error"`
	Gender string `json:"gender"`
}

type GenderStruct struct {
	Count    int64  `json:"count"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Language string `json:"gender"`
}

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
var GENDER_DEBUG = os.Getenv("GENDER_DEBUG")

var LEARNGENDER = os.Getenv("LEARNGENDER")

var GenderList []GenderStruct

// replace \_.,<>:;~+|\[\]?`"!@#$%^&*()\s chars with whitespace
// re.sub(r'[\_.,<>:;~+|\[\]?`"!@#$%^&*()\s]', ' '
func Tidy(s string) (safe string) {
	reg, err := regexp.Compile("[\\_.,:;~+|\\[\\]?`\"!@#$%^&*()\\s]+")
	if err != nil {
		Debug("Error: Tidy, ", err)
		return
	}

	text_in_lower := strings.ToLower(s)
	safe = reg.ReplaceAllLiteralString(text_in_lower, " ")
	return
}

// Serialize this classifier to a file.
func WriteToFile(name string) (err error) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error when WriteToFile(), ", err)
		return
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(GenderList)
	return
}

// De-Serialize this classifier from a file.
func ReadFromFile(name string) (GenderList []GenderStruct, err error) {
	file, err := os.Open(name)
	if err != nil {
		log.Println("Error when ReadFromFile(), ", err)
		return
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(&GenderList)
	return
}

func LearnFromFile(name string) (GenderList []GenderStruct, err error) {
	GenderList, err = ReadFromFile(name)
	if err != nil {
		log.Println("Error when LearnFromFile() ReadFromFile(), ", err)
		return
	}
	for _, v := range GenderList {
		Train(v.Gender, v.Name, v.Language)
	}
	return
}

// tidy the input text, ignore those text composed with less than 2 chars
func Tokenizer(s string, ignore_words_map map[string]int) (res []string) {
	words := strings.Fields(Tidy(s))
	// this slice's length should be initialized to 0
	// otherwise, the first element will be the whitespace(empty string)
	res = make([]string, 0)

	for _, word := range words {
		strings.TrimSpace(word)
		word = strings.ToLower(word)

		_, omit := ignore_words_map[word]
		if omit || len(word) <= 2 {
			continue
		}
		res = append(res, word)
	}

	return
}

// compute word occurances
func Occurances(words []string) (counts map[string]uint) {
	counts = make(map[string]uint)
	for _, word := range words {
		if _, ok := counts[word]; ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	//Debug("compute word occurances, ", counts)
	return
}

func Flush() {
	reply := RedisClient.SMembers(Redis_prefix + "categories")

	for _, key := range reply.Val() {
		RedisClient.Del(Redis_prefix + string(key))
	}

	RedisClient.Del(Redis_prefix + "categories")
}

func Train(categories, text, l string) {
	RedisClient.SAdd(Redis_prefix+"categories", categories)

	detectedLang := ""
	var err error
	if l == "" {
		Debug("Train, Lang could not detect language, will try to Guess ", detectedLang)
		detectedLang, err = lang.Guess(text)
		if err != nil {
			detectedLang = "en"
		}
	} else {
		detectedLang = l
	}

	//Debug("Train, Lang detected: ", detectedLang)

	token_occur := GetOccurances(detectedLang, text)

	for word, count := range token_occur {
		//Debug("Train, ", Redis_prefix + categories, word, count)
		RedisClient.HIncrBy(Redis_prefix+categories, word, int64(count))
	}
}

func Untrain(categories, text, l string) {

	detectedLang := ""
	var err error
	if l == "" {
		detectedLang, err = lang.Guess(text)
		if err != nil {
			detectedLang = "en"
		}
	} else {
		detectedLang = l
	}

	Debug("Untrain, Lang detected: ", detectedLang)

	token_occur := GetOccurances(detectedLang, text)

	for word, count := range token_occur {
		reply := RedisClient.HGet(Redis_prefix+categories, word)

		cur, _ := strconv.ParseUint(string(reply.Val()), 10, 0)
		if cur != 0 {
			inew := cur - uint64(count)
			if inew > 0 {
				RedisClient.HSet(Redis_prefix+categories, word, strconv.Itoa(int(inew)))
			} else {
				RedisClient.HDel(Redis_prefix+categories, word)
			}
		}
	}

	if Tally(categories) == 0 {
		RedisClient.Del(Redis_prefix + categories)
		RedisClient.SRem(Redis_prefix+"categories", categories)
	}
}

func Classify(text, lang string) (key string) {
	scores := Score(text, lang)
	Debug("Classify, Scores: ", scores)
	max := 0.0
	if scores != nil {
		for k, v := range scores {
			if v <= max {
				max = v
				key = k
			}
		}

		Debug("Classify, key: ", key, max)
		//if key == "bad" && max == 0 {
		//	Debug("Will reclassify false gender to not gender as score is too low for being gender ")
		//	key = "good"
		//}

		return
	}
	key = "I dont know"
	Debug("Error: Could not Classify, text: ", text, key)
	return
}

func Score(text, l string) (res map[string]float64) {

	detectedLang := ""
	var err error
	if l == "" {
		detectedLang, err = lang.Guess(text)
		if err != nil {
			detectedLang = "en"
		}
	} else {
		detectedLang = l
	}

	Debug("Score, Lang detected: ", detectedLang)

	token_occur := GetOccurances(detectedLang, text)

	Debug("Score, token_occur, ", token_occur)
	res = make(map[string]float64)

	reply := RedisClient.SMembers(Redis_prefix + "categories")

	Debug("Score, reply, ", reply)
	for v1, category := range reply.Val() {
		Debug("Score, range reply.Val() ", v1, category)
		tally := Tally(category)
		Debug("Score, tally, ", tally)
		if tally == 0 {
			continue
		}

		res[category] = 0.0
		for word, v := range token_occur {
			Debug("Score, range token_occur,", word, ", count:", v)

			Debug("Score, will run RedisClient.HGet,", Redis_prefix+category, word)
			score := RedisClient.HGet(Redis_prefix+category, word)
			Debug("Score, result of RedisClient.HGet,", score.Val())

			if score == nil {
				continue
			}

			targetVal := score.Val()
			if targetVal == "" {
				continue
			}

			iVal, err := strconv.ParseFloat(targetVal, 64)
			if err != nil {
				Debug("Error: Score, ", err)
				return nil
			}

			Debug("Score, ival ", iVal)

			if iVal == 0.0 {
				iVal = correction
			}

			res[category] += math.Log(iVal / float64(tally))
			Debug("Score, res[category], ", category, res[category])
		}
	}

	return res
}

var supportedLang []string = []string{
	"pt",
	"fr",
	"it",
	"es",
	"en",
}

func GetOccurances(lang, text string) (counts map[string]uint) {
	//check if lang is supported
	supported := false
	for _, v := range supportedLang {
		if v == lang {
			supported = true
		}
	}
	if !supported {
		Debug("WARN: Will use default lang EN ", lang, supportedLang)
		counts = Occurances(Tokenizer(text, English_ignore_words_map))
	} else if strings.ContainsAny(lang, "en") {
		counts = Occurances(Tokenizer(text, English_ignore_words_map))
	} else if strings.ContainsAny(lang, "pt") {
		counts = Occurances(Tokenizer(text, Portuguese_ignore_words_map))
	} else if lang == "es" {
		counts = Occurances(Tokenizer(text, Spanish_ignore_words_map))
	} else if lang == "it" {
		counts = Occurances(Tokenizer(text, Italian_ignore_words_map))
	} else if lang == "fr" {
		counts = Occurances(Tokenizer(text, French_ignore_words_map))
	} else {
		Debug("ERROR: GetOccurances, Could not identify Language, ", lang)
	}
	return
}

//female = good
//male =  bad
func Gender(gender string) string {
	if gender == "F" {
		return "good"
	} else {
		return "bad"
	}
}

func Tally(category string) (sum uint64) {
	vals := RedisClient.HVals(Redis_prefix + category)

	for _, val := range vals.Val() {
		iVal, err := strconv.ParseUint(string(val), 10, 0)
		if err != nil {
			Debug("Error: Tally, ", err)
			return
		}

		sum += iVal
	}
	return sum
}

// init function, load the configs
// fill english_ignore_words_map
func init() {
	if REDIS_HOST == "" {
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_PORT = "6379"
	}
	if REDIS_PASSWORD == "" {
	}

	// get redis connection info
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASSWORD, // no password set
		DB:       0,              // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		Debug("Error: init, Can't connect to Redis Server", err)
		panic("Can't connect to Redis Server")
	}
	Debug(pong)
}

func StartLanguageGender(filename string, done chan bool) {
	Debug("StartLanguageGender init, ", filename)
	filename = GetPWD(filename)
	file, _ := os.Open(filename)
	defer file.Close()
	reader := csv.NewReader(file)

	result, _ := reader.ReadAll()
	for _, record := range result {

		count, _ := strconv.ParseInt(record[2], 10, 8)
		name := strings.ToLower(record[0])
		g := Gender(record[1])
		//idx := 0
		//for idx <= int(count) {
		//Debug("Train(g, name, lang)", g, name, "en")
		Train(g, name, "en")
		//idx++
		//}
		//
		GenderList = append(GenderList, GenderStruct{
			Count:    count,
			Name:     name,
			Gender:   g,
			Language: "en",
		})
	}
	Debug("finished parsing", filename)
	Debug("StartLanguageGender end")
	done <- true
}

func GetPWD(targetFile string) (pdw string) {
	path, _ := os.Getwd()
	pdw = path + targetFile
	Debug("GetPWD pdw -> ", pdw, " path, ", path, "targetfile: ", targetFile)
	if _, err := os.Stat(pdw); os.IsNotExist(err) {
		pdw = path + "/.." + targetFile
		Debug("", pdw)
	}
	return
}

func Debug(v ...interface{}) {
	if GENDER_DEBUG == "true" {
		log.Println(v)
	}
}
