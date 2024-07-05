package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type mainPage struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

var Data map[string]mainPage

func parseStory(jsonFile []byte) {
	err := json.Unmarshal(jsonFile, &Data)
	if err != nil {
		fmt.Println("error in unmarshall jsonfile")
		panic(err)
	}
}

func storyHandler(w http.ResponseWriter, r *http.Request) {
	var arc string
	if r.URL.Path == "/" {
		arc = "intro"
	} else {
		arc = strings.TrimLeft(r.URL.Path, "/")
	}

	t, err := template.ParseFiles("temp.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(Data)
	fmt.Println(arc, Data[arc])
	err = t.Execute(w, Data[arc])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	jsonFile, err := os.ReadFile("gopher.json")
	if err != nil {
		fmt.Println("error in opening file")
		panic(err)
	}
	parseStory(jsonFile)

	mux := http.DefaultServeMux

	mux.HandleFunc("/", storyHandler)

	http.ListenAndServe(":8080", mux)
	// fmt.Println(story)
}
