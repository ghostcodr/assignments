package structs

import "strings"

type answer struct {
	text string
}

func ToAnswerArray(answersStr string) []answer {
	var answers []answer
	for _, ansStr := range strings.Split(answersStr, ";") {
		answers = append(answers, answer{
			text: strings.TrimSpace(ansStr),
		})
	}
	return answers
}
