package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
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
	entries, err := os.ReadDir("views")
	if err != nil {
		fmt.Println("Error reading directory views")
		os.Exit(1)
	}
	for _, entry := range entries {
		entry := entry.Name()
		entry = strings.TrimSuffix(entry, ".html")
		http.HandleFunc("/"+entry, func(w http.ResponseWriter, r *http.Request) {
			content, err := os.ReadFile("views/" + entry + ".html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, string(content))
		})
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
