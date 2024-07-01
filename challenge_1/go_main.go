package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"flag"
	"time"
)

type quiz struct {
	question, answer string
}

func parsefile(filepath string) (*[]quiz, error) {
	fmt.Println(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err, "-failed to open file")
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err, "-failed reading csv file")
		return nil, err
	}
	result := make([]quiz, len(records))
	i := 0
	for _, temp := range records {
		result[i].question = temp[0]
		result[i].answer = temp[1]
		i++
	}

	return &result, nil
}

func askQuestions(problems []quiz, timer *time.Timer) int {
	timerend := false
	go func() {
		<-timer.C
		timerend = true
	}()
	score := 0
	for _, q := range problems {
		fmt.Println(q.question)
		var ans string
		fmt.Scanln(&ans)
		if ans == q.answer {
			score++
		}
		if timerend {
			fmt.Println("timeout")
			return score
		}
	}

	return score
}

func main() {
	file := flag.String("file", "csvfile", "input file in csv format")
	timelimit := flag.Int("timelimit", 30, "The time limit of the quiz in seconds")
	flag.Parse()
	answerkey, err := parsefile(*file)
	if err != nil {
		fmt.Println("error while parsing csv: ", err)
		return
	}

	timer := time.NewTimer(time.Second * time.Duration(*timelimit))
	correctScore := askQuestions(*answerkey, timer)

	fmt.Println("your correct score is: ", correctScore)
}
