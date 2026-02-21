package server

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed templates/*.html.tmpl
var templatesFS embed.FS

//go:embed assets/*
var assetsFS embed.FS

type Server struct {
	templates templateSet
	mux       *http.ServeMux
}

func New() (*Server, error) {
	templates, err := newTemplateSet()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	server := &Server{templates: templates, mux: mux}

	assetFiles, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		return nil, err
	}

	mux.HandleFunc("/", server.handleIndex)
	mux.HandleFunc("/emails/", server.handleEmails)
	mux.HandleFunc("/blocks/", server.handleBlocks)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assetFiles))))

	return server, nil
}

func (server *Server) Handler() http.Handler {
	return server.mux
}

func (server *Server) ListenAndServe(address string) error {
	if strings.TrimSpace(address) == "" {
		address = ":8380"
	}

	return http.ListenAndServe(address, server.Handler())
}
