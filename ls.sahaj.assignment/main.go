package main

import (
	"bufio"
	"fmt"
	"ls.sahaj.assignment/test/structs"
	"os"
)

func main() {
	findSolution("test_file.txt")
}

func findSolution(inputFile string) {
	lines := parseFile(inputFile)
	problem := structs.NewProblem(lines)
	solution, err := problem.Solve()
	if err != nil {
		fmt.Print(err.Error())
	}
	for _, ans := range solution {
		fmt.Println(ans)
	}
}

func parseFile(inputFile string) []string {
	inFile, err := os.Open(inputFile)
	defer inFile.Close()
	if err != nil {
		fmt.Printf("failed to read file %s", inputFile)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
