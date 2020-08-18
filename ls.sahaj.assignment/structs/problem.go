package structs

import "fmt"

type Problem struct {
	paragraph paragraph
	questions []question
	answers   []answer
}

func NewProblem(lines []string) Problem {
	return Problem{paragraph: paragraph{
		text: lines[0],
	},
		questions: ToQuestionArray(lines[1:6]),
		answers:   ToAnswerArray(lines[6]),
	}
}

func (problem Problem) Solve() ([]string, error) {
	solution := make([]string, 5)
	for _, answer := range problem.answers {
		sentences, sentenceNotFound := problem.paragraph.GetMatchingSentences(answer)
		if sentenceNotFound != nil {
			return nil, sentenceNotFound
		}
		questions, err := findMatchingQuestions(problem.questions, sentences, answer)
		if err != nil {
			fmt.Printf("Matching question not found for Ans %s", answer.text)
			fmt.Println()
		} else {
			if len(questions) > 1 {
				for _, question := range questions {
					exists := solution[question.quesNo]
					if len(exists) > 1 {
						continue
					}
					solution[question.quesNo] = answer.text
				}
			} else {
				solution[questions[0].quesNo] = answer.text
			}
		}
	}
	return solution, nil
}

func findMatchingQuestions(questions []question, sentences []Sentence, answer answer) ([]question, error) {
	var matchingQuestions []question
	if len(sentences) > 1 {
		for _, sentence := range sentences {
			question, err := findBestMatchingQuestion(questions, sentence, answer)
			if err != nil {
				continue
			}
			matchingQuestions = append(matchingQuestions, question)
		}
	} else {
		question, err := findBestMatchingQuestion(questions, sentences[0], answer)
		if err != nil {
			return nil, err
		}
		matchingQuestions = append(matchingQuestions, question)
	}
	return matchingQuestions, nil
}

func findBestMatchingQuestion(questions []question, sentences Sentence, answer answer) (question, error) {
	var nearestQuestion question
	matchingWordCount := 0
	for _, question := range questions {
		count := 0
		var matchingWords []string
		for _, questionWord := range question.words {
			_, ok := Find(keyWords, questionWord)
			if ok {
				continue
			}
			_, ok = Find(matchingWords, questionWord)
			if ok {
				continue
			}

			ansWords := sentences.SplitToWords(answer.text)
			_, found := Find(ansWords, questionWord)
			if found {
				matchingWords = append(matchingWords, questionWord)
				count++
			}
		}
		if count == 0 {
			continue
		}

		if matchingWordCount == 0 || matchingWordCount < count {
			nearestQuestion = question
			matchingWordCount = count
		}
	}
	return nearestQuestion, nil
}
