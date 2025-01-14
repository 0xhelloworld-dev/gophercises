package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/0xhelloworld-dev/cyoa"
)

func main() {
	fmt.Println("Hello!!!")
	port := flag.Int("port", 3000, "listening port for app")
	filename := flag.String("file", "/Users/jonathanchua/Desktop/goProjects/gophercises/cyoa/gopher.json", "JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)
	f, err := os.Open(*filename) //flag.String returns the address of a string variable that stores the value https://pkg.go.dev/flag#String
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonToStory(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
