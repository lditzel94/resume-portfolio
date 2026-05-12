package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lditzel94/resume-portfolio/handlers"
)

func main() {
	// BASE_DIR overrides the asset root (defaults to the current working
	// directory). In containerized deploys the Dockerfile sets WORKDIR /app
	// and copies data/templates/static there, so the default works.
	root := "."
	if v := os.Getenv("BASE_DIR"); v != "" {
		root = v
	}

	dataPath := filepath.Join(root, "data", "resume.yaml")
	i18nDir := filepath.Join(root, "data", "i18n")
	tmplDir := filepath.Join(root, "templates")
	staticDir := filepath.Join(root, "static")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.PortfolioHandler(dataPath, i18nDir, tmplDir))
	mux.HandleFunc("/section/", handlers.SectionHandler(dataPath, i18nDir, tmplDir))
	mux.HandleFunc("/resume", handlers.ResumeHandler(dataPath, tmplDir))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Listening on http://0.0.0.0%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
