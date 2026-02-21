package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

func (server *Server) handleIndex(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	themeName, selectedTheme := selectedThemeFromRequest(request.URL.Query().Get("theme"))
	emailItems, err := buildEmailItems(themeName, selectedTheme)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	blockGroups, err := buildBlockItems(themeName, selectedTheme)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	setHTMLHeader(writer)
	if err := server.templates.index.ExecuteTemplate(writer, "index.html.tmpl", indexViewData{
		Title:        "Myrtle example directory",
		ThemeOptions: themeOptions(themeName),
		Theme:        themeName,
		EmailItems:   emailItems,
		BlockGroups:  blockGroups,
	}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) handleEmails(writer http.ResponseWriter, request *http.Request) {
	server.handlePreview(
		writer,
		request,
		"/emails/",
		"/emails/",
		"email.html.tmpl",
		server.templates.email,
		"Email: ",
		buildExampleEmail,
	)
}

func (server *Server) handleBlocks(writer http.ResponseWriter, request *http.Request) {
	server.handlePreview(
		writer,
		request,
		"/blocks/",
		"/blocks/",
		"block.html.tmpl",
		server.templates.block,
		"Block: ",
		buildBlockEmail,
	)
}

func (server *Server) handlePreview(
	writer http.ResponseWriter,
	request *http.Request,
	pathPrefix string,
	previewPrefix string,
	templateName string,
	pageTemplate *template.Template,
	titlePrefix string,
	builder func(name string, selectedTheme theme.Theme) (*myrtle.Email, error),
) {
	path := strings.TrimPrefix(request.URL.Path, pathPrefix)
	if path == "" || path == "/" {
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}

	themeName, selectedTheme := selectedThemeFromRequest(request.URL.Query().Get("theme"))
	isHTML := strings.HasSuffix(path, "/html")
	name := strings.TrimSuffix(strings.TrimSuffix(path, "/html"), "/")

	email, err := builder(name, selectedTheme)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	htmlOutput, err := email.HTML()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTML {
		setHTMLHeader(writer)
		_, _ = writer.Write([]byte(htmlOutput))
		return
	}

	markdownOutput, err := email.Text()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	setHTMLHeader(writer)
	if err := pageTemplate.ExecuteTemplate(writer, templateName, previewViewData{
		Title:    titlePrefix + name,
		Theme:    themeName,
		Name:     name,
		Preview:  previewPrefix + name + "/html?theme=" + themeName,
		Markdown: markdownOutput,
	}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
