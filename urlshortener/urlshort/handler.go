package urlshort

import (
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathKey := r.URL.Path
		if destination, ok := pathsToUrls[pathKey]; ok {
			fmt.Printf("Redirecting user to %s", destination)
			http.Redirect(w, r, destination, http.StatusSeeOther)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type yamlRecord struct {
	Path string `yaml:"path"` //these are called yaml tags
	Url  string `yaml:"url"`  //these should match the keys in the yaml file. helps yaml.Unmarshal match yaml keys to struct fields
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var records []yamlRecord
	err := yaml.Unmarshal(yml, &records)
	if err != nil {
		return nil, err
	}
	pathToUrls := make(map[string]string)
	for _, record := range records {
		pathToUrls[record.Path] = record.Url
	}

	return MapHandler(pathToUrls, fallback), nil
}
