package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type PagesHandler struct {
	templates *template.Template
}

func NewPagesHandler(dir string) (*PagesHandler, error) {
	tmpl, err := template.ParseGlob(filepath.Join(dir, "*.html"))
	if err != nil {
		return nil, err
	}
	return &PagesHandler{templates: tmpl}, nil
}

func (p *PagesHandler) Login(w http.ResponseWriter, r *http.Request) {
	p.templates.ExecuteTemplate(w, "login.html", nil)
}

func (p *PagesHandler) Register(w http.ResponseWriter, r *http.Request) {
	p.templates.ExecuteTemplate(w, "register.html", nil)
}

func (p *PagesHandler) Chat(w http.ResponseWriter, r *http.Request) {
	p.templates.ExecuteTemplate(w, "chat.html", nil)
}
