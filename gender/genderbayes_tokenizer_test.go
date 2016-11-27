package gender_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"comentarismo-gender/gender"
)

func TestPortuguese_tokenizer(t *testing.T) {
	test_string := "escola exemplo estudante $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"escola","exemplo","estudante", "fajs", "ldkfj"}

	words := gender.Tokenizer(test_string,gender.Portuguese_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}

func TestEnglish_tokenizer(t *testing.T) {
	test_string := "love again fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"love","again","fjalsdfj", "fajs", "ldkfj"}

	words := gender.Tokenizer(test_string,gender.English_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}


func TestSpanish_tokenizer(t *testing.T) {
	test_string := "amor parecer fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"amor","parecer","fjalsdfj", "fajs", "ldkfj"}

	words := gender.Tokenizer(test_string,gender.Spanish_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}

func TestItalian_tokenizer(t *testing.T) {
	test_string := "pizza fianco fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"pizza","fianco","fjalsdfj", "fajs", "ldkfj"}

	words := gender.Tokenizer(test_string,gender.Italian_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}

func TestFrench_tokenizer(t *testing.T) {
	test_string := "creme plusieurs fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"creme","plusieurs","fjalsdfj", "fajs", "ldkfj"}

	words := gender.Tokenizer(test_string,gender.French_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}