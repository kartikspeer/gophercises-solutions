package main

import (
	"challenge_2/urlshort"
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()
	mapHandler := urlshort.YamlHandler("urls.yaml", mux)
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world! this is default mux")
}
