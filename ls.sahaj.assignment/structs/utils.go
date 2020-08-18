package structs

import "strings"

var suffixToRemove = []string{",", "?", "."}
var keyWords = []string{"which", "what", "why", "where", "their", "is", "i", "am", "are", "a", "the", "of", "to", "do", "and"}

func GetWords(str string) []string {
	var words []string
	for _, word := range strings.Split(str, " ") {
		words = append(words, TrimSuffix(word))
	}
	return words
}
func TrimSuffix(str string) string {
	for _, suffix := range suffixToRemove {
		str = strings.TrimSuffix(strings.TrimSpace(str), suffix)
	}
	return str
}

func Find(slice []string, val string) (int, bool) {
	for idx, item := range slice {
		if item == val {
			return idx, true
		}
	}
	return -1, false
}