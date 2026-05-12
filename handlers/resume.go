package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/lditzel94/resume-portfolio/models"
	"gopkg.in/yaml.v3"
)


func loadResume(path string) (*models.Resume, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var r models.Resume
	if err := yaml.Unmarshal(f, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func ResumeHandler(dataPath, tmplDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resume, err := loadResume(dataPath)
		if err != nil {
			http.Error(w, "failed to load resume data", http.StatusInternalServerError)
			return
		}

		t, err := template.ParseFiles(
			tmplDir+"/layout.html",
			tmplDir+"/resume.html",
			tmplDir+"/partials/experience.html",
			tmplDir+"/partials/education.html",
			tmplDir+"/partials/skills.html",
		)
		if err != nil {
			http.Error(w, "failed to parse templates: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := t.ExecuteTemplate(w, "layout", resume); err != nil {
			http.Error(w, "template execution failed: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
