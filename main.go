package main

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/lditzel94/resume-portfolio/handlers"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	root := filepath.Dir(file)

	dataPath := filepath.Join(root, "data", "resume.yaml")
	tmplDir := filepath.Join(root, "templates")
	staticDir := filepath.Join(root, "static")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.ResumeHandler(dataPath, tmplDir))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	log.Println("Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
