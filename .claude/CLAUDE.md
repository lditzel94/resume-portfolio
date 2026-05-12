# resume-portfolio

Self-hosted Go + HTMX resume manager. Single YAML file as data source, no database.

## Architecture

- `main.go` — HTTP server, routing
- `handlers/resume.go` — public resume page (reads YAML, renders templates)
- `handlers/admin.go` — admin CRUD (not yet implemented)
- `handlers/export.go` — PDF export via chromedp (not yet implemented)
- `models/resume.go` — data structs matching resume.yaml schema
- `templates/` — Go html/template files; layout.html wraps everything
- `templates/partials/` — HTMX-swappable fragments (experience, education, skills)
- `static/css/resume.css` — all styles + `@media print` for PDF output
- `data/resume.yaml` — single source of truth; edit here or via /admin

## Running

```bash
# Dev (live reload) — auto-detects Windows / Linux / macOS
make dev

# Plain run (no reload)
make run

# Build binary into ./tmp/
make build
# → http://localhost:8080
```

Air configs live in `.air/` (`linux.toml` and `windows.toml`); the Makefile picks
the right one based on `$(OS)`. On Windows you need `make` (Git Bash, scoop, or
chocolatey provide it). Without `make`, fall back to:
- Linux / macOS: `air -c .air/linux.toml`
- Windows:       `air -c .air\windows.toml`

## Updating resume data

Edit `data/resume.yaml` directly, or use the admin UI at `/admin` (once implemented).
Air will detect YAML changes and rebuild automatically.

## ATS optimization approach

- `<h1>` for name, `<h2>` for section headings, `<h3>` for job titles
- No layout tables or CSS columns that break ATS parsers
- Skills rendered as plain comma-separated text inside `<dl>`
- Dates in "Month YYYY" format throughout
- PDF exported via chromedp printing the same HTML — fully selectable text, no rasterization
- Export filename: `FirstName_LastName_Resume.pdf`

## Common tasks

- Add experience entry → append to `experience:` list in resume.yaml
- Add skill category → append to `skills:` list
- Change theme colors → edit CSS variables in `static/css/resume.css`
- Export PDF → GET /export/pdf (once implemented)

## Not yet implemented

- `/admin` — inline editing UI
- `/export/pdf` — chromedp PDF generation
- Authentication (out of scope; intended for localhost use only)
