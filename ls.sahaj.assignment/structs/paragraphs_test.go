package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var p = paragraph{
	text: "The Moon is a barren, rocky world without air and water. It has dark lava plain on its surface. The Moon is filled wit craters. It has no light of its own. It gets its light from the Sun. The Moo keeps changing its shape as it moves round the Earth. It spins on its axis in 27.3 days stars were named after the Edwin Aldrin were the first ones to set their foot on the Moon on 21 July 1969 They reached the Moon in their space craft named Apollo II.",
}

func TestParagraphGetSentences(t *testing.T) {
	sentences := p.getSentences()
	assert.Equal(t, 7, len(sentences))
	assert.Equal(t, "The Moon is a barren, rocky world without air and water", sentences[0].text)
	assert.Equal(t, "It has dark lava plain on its surface", sentences[1].text)
	assert.Equal(t, "The Moon is filled wit craters", sentences[2].text)
	assert.Equal(t, "It has no light of its own", sentences[3].text)
	assert.Equal(t, "It gets its light from the Sun", sentences[4].text)
	assert.Equal(t, "The Moo keeps changing its shape as it moves round the Earth", sentences[5].text)
	assert.Equal(t, "It spins on its axis in 27.3 days stars were named after the Edwin Aldrin were the first ones to set their foot on the Moon on 21 July 1969 They reached the Moon in their space craft named Apollo II.", sentences[6].text)
}

func TestParagraph_GetMatchingSentences(t *testing.T) {
	sentences, sentenceNotFound := p.GetMatchingSentences(answer{
		text: "rocky world without air",
	})
	assert.Nil(t, sentenceNotFound)
	assert.Equal(t, 1, len(sentences))
	assert.Equal(t, "The Moon is a barren, rocky world without air and water", sentences[0].text)
}

func TestSentence_IsContains(t *testing.T) {
	sentence := Sentence{
		text: "It has dark lava plain on its surface",
	}
	assert.True(t, sentence.IsContains("lava plain"))
	assert.False(t, sentence.IsContains("lava plainn"))
}

func TestSentence_SplitToWords(t *testing.T) {
	sentence := Sentence{
		text: "It has dark lava plain on its surface",
	}
	words := sentence.SplitToWords("lava plain on its")
	assert.Equal(t, []string{"It", "has", "dark", "surface"}, words)
}
