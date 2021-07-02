package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	h "./LIB"
)

func readYAMLFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	//fmt.Println(string(data))
	return data, err
}

func main() {
	yamlFilename := flag.String("yaml", "pathslist.yml", "List of short urls'.")

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := h.MapHandler(pathsToUrls, mux)
	// fmt.Println("Starting the server on :8080")
	// log.Fatal(http.ListenAndServe(":8080", mapHandler))

	// Build the YAMLHandler using the mapHandler as the
	// 	// fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: 'https://github.com/gophercises/urlshort'
	// - path: /urlshort-final
	//   url: 'https://github.com/gophercises/urlshort/tree/solution'
	// `

	yaml, err := readYAMLFile(*yamlFilename)

	if err != nil {
		panic(err)
	}

	yamlHandler, err := h.YAMLHandler([]byte(yaml), mapHandler)
	// fmt.Println("## yamlHandler :: ", yamlHandler)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
