package urlshort

import (
	"fmt"
	"html"
	"net/http"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type urlStruct struct {
	sURL string `yaml:path`
	url  string `yaml:url`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	fmt.Println("Got Heer")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Response :: %v", html.EscapeString(r.URL.Path))
		req := html.EscapeString(r.URL.Path)
		// find the matching URL
		k, pres := pathsToUrls[r.URL.Path]
		if pres {
			// fmt.Fprintf(w, "### %v will be redirected to %v \n", req, k)
			http.Redirect(w, r, k, http.StatusPermanentRedirect)

		} else {
			sr := strings.NewReplacer("/", "")
			req := string(sr.Replace(req))
			// fmt.Fprintf(w, "### Go to %v\n", req)
			// http.Redirect(w, r, req, http.StatusPermanentRedirect)
			fmt.Fprintf(w, "Requested path has no shorthand :: %v", html.EscapeString(req))
		}

	}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	return h1
}

func generalYamlMapper(yml []byte) (map[string]string, error) {

	ymlData := []map[string]string{}
	err := yaml.Unmarshal([]byte(yml), &ymlData)
	if err != nil {
		panic(err)
	}
	urlMap := make(map[string]string)
	for _, k := range ymlData {
		// fmt.Printf("## YAML : %v , %T \n", ymlData[i], ymlData[i])
		urlMap[k["path"]] = k["url"]
	}
	return urlMap, err
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathsToUrls, err := generalYamlMapper(yml)

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Response :: %v", html.EscapeString(r.URL.Path))
		req := html.EscapeString(r.URL.Path)
		// find the matching URL
		k, pres := pathsToUrls[r.URL.Path]
		if pres {
			// fmt.Fprintf(w, "### %v will be redirected to %v \n", req, k)
			http.Redirect(w, r, k, http.StatusPermanentRedirect)

		} else {
			sr := strings.NewReplacer("/", "")
			req := string(sr.Replace(req))
			// fmt.Fprintf(w, "### Go to %v\n", req)
			// http.Redirect(w, r, req, http.StatusPermanentRedirect)
			fmt.Fprintf(w, "Requested path has no shorthand :: %v", html.EscapeString(req))
		}

	}

	return h1, err
}
