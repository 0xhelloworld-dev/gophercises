package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

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
	// Using custom template + custom path function
	tpl := template.Must(template.New("").Parse(storyTempl))
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl), cyoa.WithPathFunc(customPathFn))

	//default template
	//h := cyoa.NewHandler(story)

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func customPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/cyoa" || path == "/cyoa/" {
		path = "/cyoa/intro"
	}
	//returns "intro" string
	return path[len("/cyoa/"):]
}

var storyTempl = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <title>title</title>
</head>
<body>
    <h1> {{.Title}} </h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li><a href ="/cyoa/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`
