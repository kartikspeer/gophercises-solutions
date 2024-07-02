package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type conf struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func MapHandler(paths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := paths[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YamlHandler(file string, fallback http.Handler) http.HandlerFunc {
	var c []conf
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error in reading")
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {

		fmt.Println("error in marshal")
		panic(err)
	}
	pathToUrl := buildMap(c)
	fmt.Println(pathToUrl)
	mapHandler := MapHandler(pathToUrl, fallback)

	return mapHandler
}

func JsonHandler(file string, fallback http.Handler) http.HandlerFunc {
	var c []conf
	jsonFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error in reading")
		panic(err)
	}
	err = json.Unmarshal(jsonFile, &c)
	if err != nil {

		fmt.Println("error in marshal")
		panic(err)
	}
	pathToUrl := buildMap(c)
	fmt.Println(pathToUrl)
	mapHandler := MapHandler(pathToUrl, fallback)

	return mapHandler
}

func buildMap(c []conf) map[string]string {
	pathToUrl := make(map[string]string)

	for _, temp := range c {
		pathToUrl[temp.Path] = temp.Url
	}

	return pathToUrl
}
