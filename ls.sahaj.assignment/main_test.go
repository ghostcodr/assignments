package main

import (
	"github.com/stretchr/testify/assert"
	"ls.sahaj.assignment/test/structs"
	"testing"
)

func TestFindSolution(t *testing.T) {
	lines := parseFile("test_file.txt")
	problem := structs.NewProblem(lines)
	solution, err := problem.Solve()
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"Grevy's zebra and the mountain zebra",
		"aims to breed zebras that are phenotypically similar to the quagga",
		"horses and donkeys",
		"the plains zebra, the Grevy's zebra and the mountain zebra",
		"subgenus Hippotigris"}, solution)
}
