package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tanerijun/urlshort/handler"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: urlshort <path_to_file>")
		os.Exit(1)
	}

	fpath := os.Args[1]

	mux := defaultMux() // used as a fallback for non-existent path

	h, err := getHandler(fpath, mux)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is live on port 8080.")
	http.ListenAndServe(":8080", h)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)
	return mux
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Url shortener is running.")
}

func getHandler(fpath string, fallback http.Handler) (http.HandlerFunc, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	var h http.HandlerFunc
	switch ext := filepath.Ext(fpath); ext {
	case ".yaml":
		h, err = handler.YAMLHandler(data, fallback)
		if err != nil {
			return nil, err
		}
	case ".json":
		h, err = handler.JSONHandler(data, fallback)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	return h, nil
}
