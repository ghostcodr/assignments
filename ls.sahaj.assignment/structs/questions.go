package structs

type question struct {
	ques   string
	quesNo int
	words  []string
}

func ToQuestionArray(questionsStr []string) []question {
	var questions []question
	for idx, quesStr := range questionsStr {
		questions = append(questions, question{
			ques:   quesStr,
			quesNo: idx,
			words:  GetWords(quesStr),
		})
	}
	return questions
}
