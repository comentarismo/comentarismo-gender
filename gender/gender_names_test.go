package gender_test

import (
	"testing"
	"comentarismo-gender/gender"
	"log"
)

func TestClassifyGenderEn(t *testing.T) {
	log.Println("Will start server on learning mode, default to English. ")

	lang := "en"
	gender.Train("good", "rachael",lang)
	gender.Train("good", "hannah",lang)
	gender.Train("good", "norah",lang)
	gender.Train("good", "claire",lang)
	gender.Train("good", "lauren",lang)
	gender.Train("good", "marsha",lang)

	gender.Train("bad", "matt",lang)
	gender.Train("bad", "matthew",lang)
	gender.Train("bad", "hank",lang)
	gender.Train("bad", "mark",lang)
	gender.Train("bad", "edward",lang)
	gender.Train("bad", "henry",lang)
	gender.Train("bad", "charlie",lang)
	gender.Train("bad", "ben",lang)

	//FEMALE

	targetWord := "matt"
	class := gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "rachael"
	class = gender.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "hannah"
	class = gender.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "norah"
	class = gender.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "claire"
	class = gender.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "lauren"
	class = gender.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "marsha"
	class = gender.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	//MALE

	targetWord = "matthew"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}
	targetWord = "hank"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}
	targetWord = "mark"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}
	targetWord = "edward"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}
	targetWord = "henry"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}
	targetWord = "charlie"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}
	targetWord = "ben"
	class = gender.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

}

func init() {
	gender.Flush()
}
