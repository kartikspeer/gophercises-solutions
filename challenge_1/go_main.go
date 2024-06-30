package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

func askQuestions(problems []quiz) int {
	score := 0
	for _, q := range problems {
		fmt.Println(q.question)
		var ans string
		fmt.Scanln(&ans)
		if ans == q.answer {
			score++
		}
	}

	return score
}
func main() {
	// f := flag.String("f", "problems.csv", "input file in csv format")
	// flag.Parse()
	answerkey, err := parsefile("problems.csv")
	fmt.Println(answerkey)
	fmt.Println(err)

	correctScore := askQuestions(*answerkey)

	fmt.Println("your correct score is: ", correctScore)
}
