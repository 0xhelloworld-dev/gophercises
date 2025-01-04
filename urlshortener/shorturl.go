package main

import (
	"fmt"
	"net/http"

	//"github.com/0xhelloworld-dev/urlshortener/urlshort"
	"github.com/0xhelloworld-dev/urlshortener/urlshort"
)

func main() {
	mux := defaultMux()

	// ~~~ default setting ~~~
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// ~~~~ yaml setting ~~~
	// TODO: need to set urlshort.MapHandler
	//get yaml file based on user input
	//be cautious of spacing here, make sure you don't have any indents at the beginning of each line
	//properties of the same record (such as "url") need to be indented with two spaces
	yamlString := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
 `

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlString), mapHandler)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	fmt.Printf("Webserver started\n")
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
