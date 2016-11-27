package gender_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"comentarismo-gender/gender"
	"fmt"
	"log"
	"runtime"
)

func TestFemaleGenderHandler(t *testing.T) {

	Convey("Should Learn female names in english and report gender female for a female name", t, func() {

		//var start = 1950
		//var end = 2012
		runtime.GOMAXPROCS(runtime.NumCPU())

		var start = 2012
		var end = 2012
		log.Println("Will start server on learning mode")

		done := make(chan bool, end - start)
		for i := start; i <= end; i++ {
			targetFile := fmt.Sprintf("/en/yob%d.txt", i)
			log.Println("Will learn ", targetFile)
			go gender.StartLanguageGender(targetFile, done)
		}
		for j := start; j <= end; j++ {
			<-done
		}

		//now ask question

		lang := "en"

		targetWord := "matt"
		class := gender.Classify(targetWord, lang)
		if class != "bad" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}

		targetWord = "rachael"
		class = gender.Classify(targetWord, lang)
		if class != "good" {
			t.Errorf("Classify failed, word (%s) should be good, result: %s", targetWord, class)
		}

		targetWord = "hannah"
		class = gender.Classify(targetWord, lang)
		if class != "good" {
			t.Errorf("Classify failed, word (%s) should be good, result: %s", targetWord, class)
		}

		targetWord = "norah"
		class = gender.Classify(targetWord, lang)
		if class != "good" {
			t.Errorf("Classify failed, word (%s) should be good, result: %s", targetWord, class)
		}

		targetWord = "claire"
		class = gender.Classify(targetWord, lang)
		if class != "good" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}

		targetWord = "lauren"
		class = gender.Classify(targetWord, lang)
		if class != "good" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}

		targetWord = "marsha"
		class = gender.Classify(targetWord, lang)
		if class != "good" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}

		//MALE

		targetWord = "hank"
		class = gender.Classify(targetWord, lang)
		if class != "bad" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}
		targetWord = "mark"
		class = gender.Classify(targetWord, lang)
		if class != "bad" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}
		targetWord = "edward"
		class = gender.Classify(targetWord, lang)
		if class != "bad" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}
		targetWord = "ben"
		class = gender.Classify(targetWord, lang)
		if class != "bad" {
			t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
		}
	})

}

func init() {
	gender.Flush()
}