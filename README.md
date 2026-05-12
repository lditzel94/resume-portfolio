# resume-portfolio

Self-hosted personal portfolio and ATS-optimized resume in Go + HTMX.
A single YAML file is the source of truth — no database, no build step.

- **Interactive portfolio** at `/` — animated orb, glass-pill bottom nav,
  EN/ES language switching, contact modal.
- **Printable ATS resume** at `/resume` — semantic HTML with `@media print`
  rules for clean PDF output.

## Stack

- Go `net/http` + `html/template` (no framework)
- HTMX for section-swap navigation
- YAML for data (`data/resume.yaml`, `data/i18n/{en,es}.yaml`)
- Vanilla CSS with view transitions

## Quick start

```bash
make dev    # live reload via air (auto-detects Windows / Linux / macOS)
make run    # one-shot go run
make build  # build a binary into ./tmp/
```

Then open <http://localhost:8080>.

If you don't have `make` on Windows, fall back to:

```powershell
air -c .air\windows.toml
```

## Project layout

```
main.go                       HTTP server, routing
handlers/
  portfolio.go                interactive portfolio + section partials
  resume.go                   printable ATS resume
models/
  resume.go                   YAML schema for resume data
  i18n.go                     YAML schema for EN/ES translations
templates/
  resume.html, layout.html    ATS resume page
  portfolio/
    layout.html               topbar + stage + bottom pill nav
    sections/                 home, summary, experience, skills, links
data/
  resume.yaml                 experience, skills, education, projects
  i18n/{en,es}.yaml           translatable copy
static/
  css/{portfolio,resume}.css
  img/profile.png
  js/htmx.min.js
.air/                         air live-reload configs (linux + windows)
Makefile                      cross-platform dev/run/build targets
```

## Editing content

- **Experience, skills, certifications** → `data/resume.yaml`
- **Translated copy** (taglines, labels, experience bullet translations)
  → `data/i18n/en.yaml` and `data/i18n/es.yaml`
- **Theme** → CSS variables at the top of `static/css/portfolio.css`

YAML files are reloaded on every request — no restart needed after editing
copy. Go code and template structure changes do require a restart (or run
`make dev` for auto-reload).

## License

MIT
