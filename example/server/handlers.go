package server

import (
	"html/template"
	"net/http"
	"net/url"
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
		Title:        "Example Emails",
		ThemeOptions: themeOptions(themeName),
		Theme:        themeName,
		EmailItems:   emailItems,
		BlockGroups:  blockGroups,
		SMTPEnabled:  server.smtp != nil,
		SMTPDefault:  server.defaultTo,
		SendStatus:   parseSendStatus(request),
	}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) handleSendEmail(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if server.smtp == nil {
		http.Error(writer, "smtp is not configured", http.StatusNotFound)
		return
	}

	if err := request.ParseForm(); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(request.FormValue("name"))
	to := strings.TrimSpace(request.FormValue("to"))
	themeName, selectedTheme := selectedThemeFromRequest(request.FormValue("theme"))

	redirectWithStatus := func(success bool, message string) {
		values := url.Values{}
		values.Set("theme", themeName)
		values.Set("send_name", name)
		if success {
			values.Set("send_ok", "1")
		} else {
			values.Set("send_error", message)
		}

		http.Redirect(writer, request, "/?"+values.Encode(), http.StatusSeeOther)
	}

	if name == "" {
		redirectWithStatus(false, "missing email key")
		return
	}

	if to == "" {
		redirectWithStatus(false, "recipient is required")
		return
	}

	email, err := buildExampleEmail(name, selectedTheme)
	if err != nil {
		redirectWithStatus(false, "unknown example email")
		return
	}

	htmlBody, err := email.HTML()
	if err != nil {
		redirectWithStatus(false, "failed to render html")
		return
	}

	textBody, err := email.Text()
	if err != nil {
		redirectWithStatus(false, "failed to render text fallback")
		return
	}

	subject := "Myrtle example: " + goEmailName(name)
	if err := server.smtp.Send(to, subject, htmlBody, textBody); err != nil {
		redirectWithStatus(false, err.Error())
		return
	}

	redirectWithStatus(true, "sent")
}

func parseSendStatus(request *http.Request) *sendStatus {
	if request == nil {
		return nil
	}

	name := strings.TrimSpace(request.URL.Query().Get("send_name"))
	if name == "" {
		return nil
	}

	if request.URL.Query().Get("send_ok") == "1" {
		return &sendStatus{
			Name:    name,
			Success: true,
			Message: "Email sent successfully.",
		}
	}

	message := strings.TrimSpace(request.URL.Query().Get("send_error"))
	if message == "" {
		message = "Failed to send email."
	}

	return &sendStatus{
		Name:    name,
		Success: false,
		Message: message,
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

	textOutput, err := email.Text()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	setHTMLHeader(writer)
	if err := pageTemplate.ExecuteTemplate(writer, templateName, previewViewData{
		Title:   titlePrefix + name,
		Theme:   themeName,
		Name:    name,
		Preview: previewPrefix + name + "/html?theme=" + themeName,
		Text:    textOutput,
	}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
