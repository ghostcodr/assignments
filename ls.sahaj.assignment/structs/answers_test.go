package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAnswerArray(t *testing.T) {
	ansText := "subgenus Hippotigris; the plains zebra, the Grévy's zebra and the mountain zebra;horses and donkeys;aims to breed zebras that are phenotypically similar to the quagga; Grévy's zebra and the mountain zebra"
	answerArray := ToAnswerArray(ansText)
	assert.ElementsMatch(t, []answer{{
		text: "subgenus Hippotigris",
	}, {
		text: "the plains zebra, the Grévy's zebra and the mountain zebra",
	}, {
		text: "horses and donkeys",
	}, {
		text: "aims to breed zebras that are phenotypically similar to the quagga",
	}, {
		text: "Grévy's zebra and the mountain zebra",
	},
	}, answerArray)
}
