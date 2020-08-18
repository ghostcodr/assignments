package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToQuestionArray(t *testing.T) {
	strArr := []string{"abc", "dev", "xyz"}
	questionArray := ToQuestionArray(strArr)
	assert.ElementsMatch(t, []question{{
		ques:   "abc",
		quesNo: 0,
		words:  []string{"abc"},
	}, {
		ques:   "dev",
		quesNo: 1,
		words:  []string{"dev"},
	}, {
		ques:   "xyz",
		quesNo: 2,
		words:  []string{"xyz"},
	}}, questionArray)
}

func TestToQuestionArrayMultiwords(t *testing.T) {
	strArr := []string{"What is your name?", "How are you?"}
	questionArray := ToQuestionArray(strArr)
	assert.ElementsMatch(t, []question{{
		ques:   "What is your name?",
		quesNo: 0,
		words:  []string{"What", "is", "your", "name"},
	}, {
		ques:   "How are you?",
		quesNo: 1,
		words:  []string{"How", "are", "you"},
	}}, questionArray)
}
