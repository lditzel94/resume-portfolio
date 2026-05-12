package handlers

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lditzel94/resume-portfolio/models"
	"gopkg.in/yaml.v3"
)

var buildID = fmt.Sprintf("%d", time.Now().Unix())

type PortfolioPage struct {
	Resume       *models.Resume
	T            *models.I18n
	Lang         string
	Section      string // current section: home, summary, experience, skills, links
	Initials     string // computed from Resume.Meta.Name (e.g. "LD")
	LinkedInUser string // trailing username only (e.g. "luciano-ditzel")
	EmailB64     string // base64-encoded email for client-side decode (avoids raw DOM exposure)
	BuildID      string // cache-bust suffix for static assets
	Status       string // availability: "open" (default), "busy" or "away" — set via STATUS env var
}

var validStatuses = map[string]bool{"open": true, "busy": true, "away": true}

func getStatus() string {
	s := strings.ToLower(strings.TrimSpace(os.Getenv("STATUS")))
	if validStatuses[s] {
		return s
	}
	return "open"
}

func linkedInUser(url string) string {
	trimmed := strings.Trim(url, "/")
	if i := strings.LastIndex(trimmed, "/"); i >= 0 {
		return trimmed[i+1:]
	}
	return trimmed
}

func computeInitials(name string) string {
	parts := strings.Fields(name)
	var b strings.Builder
	for i, p := range parts {
		if i >= 2 {
			break
		}
		b.WriteRune([]rune(p)[0])
	}
	return strings.ToUpper(b.String())
}

var validSections = map[string]bool{
	"home":       true,
	"summary":    true,
	"experience": true,
	"skills":     true,
	"links":      true,
}

func getLang(r *http.Request) string {
	if lang := r.URL.Query().Get("lang"); lang == "en" || lang == "es" {
		return lang
	}
	if c, err := r.Cookie("lang"); err == nil {
		if c.Value == "en" || c.Value == "es" {
			return c.Value
		}
	}
	return "en"
}

func loadI18n(dir, lang string) (*models.I18n, error) {
	data, err := os.ReadFile(filepath.Join(dir, lang+".yaml"))
	if err != nil {
		return nil, err
	}
	var i models.I18n
	if err := yaml.Unmarshal(data, &i); err != nil {
		return nil, err
	}
	return &i, nil
}

func parsePortfolioTemplates(tmplDir string) (*template.Template, error) {
	return template.ParseFiles(
		tmplDir+"/portfolio/layout.html",
		tmplDir+"/portfolio/sections/home.html",
		tmplDir+"/portfolio/sections/summary.html",
		tmplDir+"/portfolio/sections/experience.html",
		tmplDir+"/portfolio/sections/skills.html",
		tmplDir+"/portfolio/sections/links.html",
	)
}

func getSection(r *http.Request) string {
	if s := r.URL.Query().Get("section"); validSections[s] {
		return s
	}
	return "home"
}

func loadPage(dataPath, i18nDir string, r *http.Request) (*PortfolioPage, error) {
	resume, err := loadResume(dataPath)
	if err != nil {
		return nil, err
	}
	lang := getLang(r)
	t18n, err := loadI18n(i18nDir, lang)
	if err != nil {
		return nil, err
	}
	return &PortfolioPage{
		Resume:       resume,
		T:            t18n,
		Lang:         lang,
		Section:      getSection(r),
		Initials:     computeInitials(resume.Meta.Name),
		LinkedInUser: linkedInUser(resume.Meta.LinkedIn),
		EmailB64:     base64.StdEncoding.EncodeToString([]byte(resume.Meta.Email)),
		BuildID:      buildID,
		Status:       getStatus(),
	}, nil
}

func PortfolioHandler(dataPath, i18nDir, tmplDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Lang query param: set cookie + redirect, preserving section if present
		if q := r.URL.Query().Get("lang"); q == "en" || q == "es" {
			http.SetCookie(w, &http.Cookie{
				Name:     "lang",
				Value:    q,
				Path:     "/",
				MaxAge:   86400 * 365,
				SameSite: http.SameSiteLaxMode,
			})
			redirect := "/"
			if s := r.URL.Query().Get("section"); validSections[s] && s != "home" {
				redirect = "/?section=" + s
			}
			http.Redirect(w, r, redirect, http.StatusSeeOther)
			return
		}

		page, err := loadPage(dataPath, i18nDir, r)
		if err != nil {
			http.Error(w, "data load failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		t, err := parsePortfolioTemplates(tmplDir)
		if err != nil {
			http.Error(w, "template parse failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := t.ExecuteTemplate(w, "portfolio-layout", page); err != nil {
			http.Error(w, "render failed: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// SectionHandler serves a single section partial for HTMX swapping.
func SectionHandler(dataPath, i18nDir, tmplDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/section/")
		if !validSections[name] {
			http.NotFound(w, r)
			return
		}

		page, err := loadPage(dataPath, i18nDir, r)
		if err != nil {
			http.Error(w, "data load failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		page.Section = name

		t, err := parsePortfolioTemplates(tmplDir)
		if err != nil {
			http.Error(w, "template parse failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := t.ExecuteTemplate(w, "section-"+name, page); err != nil {
			http.Error(w, "render failed: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
