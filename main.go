package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func startHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("views/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	})

	filepath.Walk("views", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error reading directory views")
			os.Exit(1)
		}
		if len(path) > 6 {
			path = strings.Replace(path, "\\", "/", -1)
			if path[len(path)-5:] == ".html" {
				http.HandleFunc("/"+path[6:len(path)-5], func(w http.ResponseWriter, r *http.Request) {
					content, err := os.ReadFile(path)
					if err != nil {
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
						return
					}
					fmt.Fprint(w, string(content))
				})
			}
		}

		return nil
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
