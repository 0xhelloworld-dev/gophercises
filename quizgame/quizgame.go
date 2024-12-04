package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/maja42/goval"
)

type QuizResult struct {
	totalScore int
	maxScore   int
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.String("time", "30", "specify the time in seconds for this to run")
	flag.Parse()

	timeSeconds, err := strconv.Atoi(*timeLimit)

	//wait for user input to begin quiz
	var userinput string
	fmt.Println("Press enter to start quiz...")
	fmt.Scanln(&userinput)
	fmt.Println("Beginning quiz.... ")

	//open csv file and parse records
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
		os.Exit(1)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	//need to initialize channels - channel for timer and results
	timerChan := make(chan bool, 1)
	resultChan := make(chan QuizResult, 1)

	//start the timer in go routine
	go func() {
		time.Sleep(time.Duration(timeSeconds) * time.Second)
		timerChan <- true
	}()

	go func() {
		result := gradeQuiz(data) //when finished send result to the channel
		resultChan <- result
	}()

	select {
	case result := <-resultChan:
		fmt.Printf("[+] Final Score [Finished]: %d / %d\n", result.totalScore, result.maxScore)
	case <-timerChan:
		fmt.Printf("Time's up!")
		exit(fmt.Sprintf("[+] Final Score [Time Limit]: incomplete"))
	}

}

func checkAnswer(challenge string, studentanswer int) bool {
	eval := goval.NewEvaluator()

	//fmt.Printf("\t [+] Challenge: %s\n", challenge)

	evaluateResult, err := eval.Evaluate(challenge, nil, nil)
	if err != nil {
		panic(err)
	}
	correctAnswer := evaluateResult.(int)
	return correctAnswer == studentanswer

}

func gradeQuiz(quizdata [][]string) QuizResult {
	result := QuizResult{
		maxScore:   len(quizdata),
		totalScore: 0,
	}

	for rowIndex, rowValues := range quizdata {
		//fmt.Printf("Row Value: %d:%v\n", rowIndex, rowValues)
		var challenge string
		var studentanswer int

		for columnIndex, columnValue := range rowValues {
			if columnIndex == 0 {
				challenge = columnValue
				fmt.Printf("%d. %s=", rowIndex+1, challenge)
			}
			if columnIndex == 1 {
				fmt.Scanln(&studentanswer)
				//studentanswer, err = strconv.Atoi(columnValue) //convert columnValue into an int
			}
		}

		checkAnswer := checkAnswer(challenge, studentanswer)
		if checkAnswer {
			fmt.Printf("Answer is correct\n")
			result.totalScore++
		} else {
			fmt.Printf("Answer is incorrect\n")
		}
		//fmt.Printf("\t \t [+] Current score: %d\n", totalScore)
		//fmt.Printf("[+] Challenge: %s, Answer:%s, checkAnswer: %v, finalScore: %s\n", challenge, strconv.Itoa(answer), checkAnswer, strconv.Itoa(finalScore)) // convert answer back to a string
	}

	return result
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
