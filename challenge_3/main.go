package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"
)

type mainPage struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]mainPage

func parseStory(jsonFile []byte) Story {
	var temp Story
	err := json.Unmarshal(jsonFile, &temp)
	if err != nil {
		fmt.Println("error in unmarshall jsonfile")
		panic(err)
	}

	return temp
}

func main() {
	jsonFile, err := os.ReadFile("gopher.json")
	if err != nil {
		fmt.Println("error in opening file")
		panic(err)
	}
	story := parseStory(jsonFile)

	mux := http.DefaultServeMux

	mux.Handle("/", story)

	http.ListenAndServe(":8080", mux)
	fmt.Println(story)
}
