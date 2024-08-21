package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Printf("Request path: %s\n", path)
		dest, found := pathsToUrls[path]
		if found {
			fmt.Printf("Redirecting to: %s\n", dest)
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fmt.Println("Path not found, using fallback")
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []PathURL
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		fmt.Printf("YAML unmarshal error: %v\n", err)
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		fmt.Printf("Path: %s, URL: %s\n", pu.Path, pu.URL)
		pathsToUrls[pu.Path] = pu.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}
