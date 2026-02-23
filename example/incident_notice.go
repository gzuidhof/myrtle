package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	"github.com/gzuidhof/myrtle/theme/flat"
)

func IncidentNoticeEmail() (*myrtle.Email, error) {
	return IncidentNoticeEmailWithTheme(flat.New())
}

func IncidentNoticeEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = flat.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
	).
		AddHeading("Incident update: delayed webhook delivery").
		Preheader("Status page updated at 09:42 UTC").
		Product("Myrtle Status", "https://status.example.com").
		Logo("/assets/logo.png", "Myrtle Status").
		AddCallout(myrtle.CalloutTypeWarning, "Investigating", "Webhook deliveries are delayed for some regions. Messages are queued and safe.").
		AddText("Our team is actively investigating and will share updates every 30 minutes.").
		AddList([]string{"Initial detection: 09:07 UTC", "Mitigation started: 09:18 UTC", "Next update: 10:00 UTC"}, false).
		AddText("Track live updates on the status page:").
		AddButton("Open status page", "https://status.example.com/incidents/abc-123").
		Build(), nil
}
