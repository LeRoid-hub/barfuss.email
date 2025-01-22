package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

type SlugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct {
}

func (f FileReader) Read(slug string) (string, error) {
	file, err := os.Open(slug + ".md")
	if err != nil {
		return "", err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func SiteHandler(s SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data Data
		slug := r.PathValue("slug")
		if slug == "" {
			slug = "index"
		}
		content, err := s.Read(slug)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		rest, err := frontmatter.Parse(strings.NewReader(content), &data)
		if err != nil {
			http.Error(w, "Error parsing frontmatter"+err.Error(), http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer

		md := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("github"),
				),
			),
		)

		err = md.Convert(rest, &buf)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
			return
		}

		data.Content = template.HTML(buf.String())

		err = Template.Execute(w, data)

		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	}
}

func english() http.HandlerFunc {
	html := `
		<!DOCTYPE html>
		<html>
		<body>

		<audio controls>
		<source src="/assets/english.wav" type="audio/wav">
		Your browser does not support the audio element.
		</audio>

		<p> Downlaod: <a href="/assets/english.wav">english.wav</a> </p>

		</body>
		</html>

	`

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}
}

type Data struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
}

var (
	Template template.Template
)

func main() {
	tpl, err := template.ParseFiles("template.html")
	Template = *tpl

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/english", english())
	mux.HandleFunc("/", SiteHandler(FileReader{}))
	mux.HandleFunc("/{slug}", SiteHandler(FileReader{}))

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
