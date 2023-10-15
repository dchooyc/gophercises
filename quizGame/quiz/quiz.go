package quiz

import (
	"fmt"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func ConvertToProblems(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func StartQuiz(problems []Problem, timeLimit *int) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.Question)
		answerCh := make(chan string)
		go getAnswer(answerCh)

		select {
		case <-timer.C:
			fmt.Println("\nTimes up!")
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.Answer {
				correct++
				fmt.Println("Correct!")
			}
		}
	}
}

func getAnswer(answerCh chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	answerCh <- answer
}
