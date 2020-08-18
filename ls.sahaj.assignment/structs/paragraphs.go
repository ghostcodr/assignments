package structs

import (
	"regexp"
	"strings"
)

var sentenceRegEx = regexp.MustCompile("\\. ")

type paragraph struct {
	text string
}

func (p paragraph) getSentences() []Sentence {
	var sentences []Sentence
	for _, sentenceStr := range sentenceRegEx.Split(p.text, -1) {
		sentences = append(sentences,
			Sentence{
				text:  sentenceStr,
			},
		)
	}
	return sentences
}

func (p paragraph) GetMatchingSentences(ans answer) ([]Sentence, *SentenceNotFound) {
	var sentences []Sentence
	for _, sentence := range p.getSentences() {
		if sentence.IsContains(ans.text) {
			sentences = append(sentences, sentence)
		}
	}
	if len(sentences) > 0 {
		return sentences, nil
	}
	return nil, &SentenceNotFound{Text: ans.text}
}

type Sentence struct {
	text  string
}

func (s Sentence) IsContains(subStr string) bool {
	return strings.Contains(s.text, subStr)
}

func (s Sentence) SplitToWords(subStr string) []string {
	var words []string
	for _, subSentence := range strings.Split(s.text, subStr) {
		words = append(words, GetWords(strings.TrimSpace(subSentence))...)
	}
	return words
}
