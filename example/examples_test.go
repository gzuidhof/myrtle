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
	constructors := []struct {
		name  string
		build func() (*myrtle.Email, error)
	}{
		{name: "welcome", build: func() (*myrtle.Email, error) { return WelcomeEmail() }},
		{name: "security", build: func() (*myrtle.Email, error) { return SecurityCodeEmail() }},
		{name: "password-reset", build: func() (*myrtle.Email, error) { return PasswordResetEmail() }},
		{name: "report", build: func() (*myrtle.Email, error) { return WeeklyReportEmail() }},
		{name: "common-blocks", build: func() (*myrtle.Email, error) { return CommonBlocksEmail() }},
		{name: "onboarding-checklist", build: func() (*myrtle.Email, error) { return OnboardingChecklistEmail() }},
		{name: "billing-receipt", build: func() (*myrtle.Email, error) { return BillingReceiptEmail() }},
		{name: "incident-notice", build: func() (*myrtle.Email, error) { return IncidentNoticeEmail() }},
		{name: "feature-digest", build: func() (*myrtle.Email, error) { return FeatureDigestEmail() }},
		{name: "high-impact", build: func() (*myrtle.Email, error) { return HighImpactEmail() }},
		{name: "columns-complex", build: func() (*myrtle.Email, error) { return ColumnsComplexEmail() }},
		{name: "bar-chart", build: func() (*myrtle.Email, error) { return BarChartEmail() }},
		{name: "product-launch", build: func() (*myrtle.Email, error) { return ProductLaunchEmail() }},
		{name: "invoice-summary", build: func() (*myrtle.Email, error) { return InvoiceSummaryEmail() }},
		{name: "activity-empty-state", build: func() (*myrtle.Email, error) { return ActivityEmptyStateEmail() }},
		{name: "monster", build: func() (*myrtle.Email, error) { return MonsterEmail() }},
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
