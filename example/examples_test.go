package example

import (
	"testing"

	"github.com/gzuidhof/myrtle"
)

type sender interface {
	Send(subject, htmlBody, textBody string) error
}

type mockSender struct {
	count int
}

func (mock *mockSender) Send(htmlBody, textBody string) error {
	mock.count++
	if htmlBody == "" || textBody == "" {
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
		{name: "password-reset", build: func() (*myrtle.Email, error) { return PasswordResetEmail() }},
		{name: "account-deletion-confirmation", build: func() (*myrtle.Email, error) { return AccountDeletionConfirmationEmail() }},
		{name: "report", build: func() (*myrtle.Email, error) { return WeeklyReportEmail() }},
		{name: "onboarding-checklist", build: func() (*myrtle.Email, error) { return OnboardingChecklistEmail() }},
		{name: "billing-receipt", build: func() (*myrtle.Email, error) { return BillingReceiptEmail() }},
		{name: "incident-notice", build: func() (*myrtle.Email, error) { return IncidentNoticeEmail() }},
		{name: "feature-digest", build: func() (*myrtle.Email, error) { return FeatureDigestEmail() }},
		{name: "weekly-operations-brief", build: func() (*myrtle.Email, error) { return WeeklyOperationsBriefEmail() }},
		{name: "columns-complex", build: func() (*myrtle.Email, error) { return ColumnsComplexEmail() }},
		{name: "bar-chart", build: func() (*myrtle.Email, error) { return HorizontalBarChartEmail() }},
		{name: "vertical-bar-chart", build: func() (*myrtle.Email, error) { return VerticalBarChartEmail() }},
		{name: "vertical-bar-chart-ticks", build: func() (*myrtle.Email, error) { return VerticalBarChartTicksEmail() }},
		{name: "inset-modes", build: func() (*myrtle.Email, error) { return InsetModesEmail() }},
		{name: "invoice-summary", build: func() (*myrtle.Email, error) { return InvoiceSummaryEmail() }},
		{name: "activity-empty-state", build: func() (*myrtle.Email, error) { return ActivityEmptyStateEmail() }},
		{name: "dark-mode-styles", build: func() (*myrtle.Email, error) { return DarkModeStylesEmail() }},
		{name: "monster", build: func() (*myrtle.Email, error) { return MonsterEmail() }},
		{name: "monster-dark-mode", build: func() (*myrtle.Email, error) { return MonsterDarkModeEmail() }},
		{name: "monster-rtl", build: func() (*myrtle.Email, error) { return MonsterRTLEmail() }},
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

		textBody, err := email.Text()
		if err != nil {
			t.Fatalf("%s text render failed: %v", constructor.name, err)
		}

		if err := mock.Send(htmlBody, textBody); err != nil {
			t.Fatalf("%s send failed: %v", constructor.name, err)
		}
	}

	if mock.count != len(constructors) {
		t.Fatalf("expected %d sends got %d", len(constructors), mock.count)
	}
}
