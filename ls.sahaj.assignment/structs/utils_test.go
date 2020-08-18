package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	assert.Equal(t, TrimSuffix("example,"), "example")
	assert.Equal(t, TrimSuffix("example"), "example")
	assert.Equal(t, TrimSuffix("example."), "example")
	assert.Equal(t, TrimSuffix("example!"), "example!")
}

func TestFind(t *testing.T) {
	arr := []string{"1", "2", "3"}
	i, b := Find(arr, "1")
	assert.True(t, b)
	assert.Equal(t, 0, i)
	i, b = Find(arr, "4")
	assert.False(t, b)
	assert.Equal(t, -1, i)
}

func TestGetWords(t *testing.T) {
	sentence := "Test get words"
	words := GetWords(sentence)
	assert.ElementsMatch(t, []string{"Test", "get", "words"}, words)
	sentence = "Test"
	words = GetWords(sentence)
	assert.ElementsMatch(t, []string{"Test"}, words)
}
