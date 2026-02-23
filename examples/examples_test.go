package example

import (
	"testing"

	"github.com/gzuidhof/myrtle"
)

type sender interface {
	Send(subject, htmlBody, markdownBody string) error
}

type mockSender struct {
	count int
}

func (mock *mockSender) Send(htmlBody, markdownBody string) error {
	mock.count++
	if htmlBody == "" || markdownBody == "" {
		return errInvalidEmail
	}
	return nil
}

type emailError string

func (value emailError) Error() string {
	return string(value)
}

const errInvalidEmail = emailError("invalid email content")

func TestRenderAndSendExamples(t *testing.T) {
	t.Parallel()

	constructors := []struct {
		name  string
		build func() (*myrtle.Email, error)
	}{
		{name: "welcome", build: func() (*myrtle.Email, error) { return WelcomeEmail() }},
		{name: "security", build: func() (*myrtle.Email, error) { return SecurityCodeEmail() }},
		{name: "report", build: func() (*myrtle.Email, error) { return WeeklyReportEmail() }},
	}

	mock := &mockSender{}

	for _, constructor := range constructors {
		email, err := constructor.build()
		if err != nil {
			t.Fatalf("%s build failed: %v", constructor.name, err)
		}

		htmlBody, err := email.HTML()
		if err != nil {
			t.Fatalf("%s html render failed: %v", constructor.name, err)
		}

		markdownBody, err := email.Text()
		if err != nil {
			t.Fatalf("%s markdown render failed: %v", constructor.name, err)
		}

		if err := mock.Send(htmlBody, markdownBody); err != nil {
			t.Fatalf("%s send failed: %v", constructor.name, err)
		}
	}

	if mock.count != len(constructors) {
		t.Fatalf("expected %d sends got %d", len(constructors), mock.count)
	}
}
