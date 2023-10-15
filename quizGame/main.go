package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"quizgame/quiz"
)

func main() {
	csvFilename := flag.String("csv", "quiz.csv", "Target csv file for quiz")
	timeLimit := flag.Int("limit", 5, "Time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening file: %s\n", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Error reading csv file")
	}

	problems := quiz.ConvertToProblems(lines)
	quiz.StartQuiz(problems, timeLimit)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
