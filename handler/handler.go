package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}

	if err := yaml.Unmarshal(yml, &pathToUrls); err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathToUrl := range pathToUrls {
			if pathToUrl.Path == r.URL.Path {
				http.Redirect(w, r, pathToUrl.URL, http.StatusFound)
				return
			}
		}

		fallback.ServeHTTP(w, r)
	}, nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//	{
//	  "path": "/some-path",
//	  "url": "https://www.some-url.com/demo"
//	}
//
// The only errors that can be returned all related to having
// invalid JSON data.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathToUrls []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}

	if err := json.Unmarshal(jsn, &pathToUrls); err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathToUrl := range pathToUrls {
			if pathToUrl.Path == r.URL.Path {
				http.Redirect(w, r, pathToUrl.URL, http.StatusFound)
			}
		}

		fallback.ServeHTTP(w, r)
	}, nil
}
