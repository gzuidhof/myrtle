package customblock

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

//go:embed *.html.tmpl
var templatesFS embed.FS

// FeatureFlagRolloutKind is the block kind identifier used by the feature rollout custom block.
const FeatureFlagRolloutKind theme.BlockKind = "feature_flag_rollout"

// FeatureFlagRollout describes rollout status data rendered by the custom feature flag block.
type FeatureFlagRollout struct {
	FlagName         string
	Environment      string
	RolloutPercent   int
	ErrorBudgetUsed  string
	P95LatencyDelta  string
	AutoRollback     string
	Status           string
	Owner            string
	ChangeID         string
	OpenFlagURL      string
	RollbackNowURL   string
	IncidentBoardURL string
}

func NewFeatureFlagRolloutBlock(data FeatureFlagRollout) myrtle.Block {
	return myrtle.NewCustomBlock(FeatureFlagRolloutKind, data, renderHTML, renderText)
}

var featureFlagRolloutTemplate = template.Must(template.New("feature-flag-rollout").ParseFS(
	templatesFS,
	"feature_flag_rollout.html.tmpl",
))

const featureFlagRolloutTemplateName = "feature_flag_rollout.html.tmpl"

func renderHTML(data FeatureFlagRollout, values theme.Values) (string, error) {
	normalized := normalize(data)
	payload := struct {
		Data   FeatureFlagRollout
		Values theme.Values
	}{
		Data:   normalized,
		Values: values,
	}

	var output bytes.Buffer
	if err := featureFlagRolloutTemplate.ExecuteTemplate(&output, featureFlagRolloutTemplateName, payload); err != nil {
		return "", err
	}

	return strings.TrimSpace(output.String()), nil
}

func renderText(data FeatureFlagRollout, _ myrtle.RenderContext) (string, error) {
	normalized := normalize(data)

	main := fmt.Sprintf(`Feature flag rollout
--------------------
- Flag: %s (%s)
- Status: %s
- Rollout: %d%%
- Error budget used: %s
- P95 latency delta: %s
- Auto rollback: %s
- Owner: %s
- Change ID: %s
- Open flag (%s)`,
		normalized.FlagName,
		normalized.Environment,
		normalized.Status,
		normalized.RolloutPercent,
		normalized.ErrorBudgetUsed,
		normalized.P95LatencyDelta,
		normalized.AutoRollback,
		normalized.Owner,
		normalized.ChangeID,
		normalized.OpenFlagURL,
	)

	parts := []string{main}
	if normalized.RollbackNowURL != "" {
		parts = append(parts, fmt.Sprintf("- Rollback now (%s)", normalized.RollbackNowURL))
	}
	if normalized.IncidentBoardURL != "" {
		parts = append(parts, fmt.Sprintf("- Incident board (%s)", normalized.IncidentBoardURL))
	}

	return strings.Join(parts, "\n"), nil
}

func normalize(value FeatureFlagRollout) FeatureFlagRollout {
	normalized := value

	normalized.FlagName = trimOrDefault(normalized.FlagName, "checkout.v2")
	normalized.Environment = trimOrDefault(normalized.Environment, "production")
	normalized.RolloutPercent = clamp(normalized.RolloutPercent, 0, 100)
	normalized.ErrorBudgetUsed = trimOrDefault(normalized.ErrorBudgetUsed, "0%")
	normalized.P95LatencyDelta = trimOrDefault(normalized.P95LatencyDelta, "+0ms")
	normalized.AutoRollback = trimOrDefault(normalized.AutoRollback, "enabled")
	normalized.Status = normalizeStatus(normalized.Status)
	normalized.Owner = trimOrDefault(normalized.Owner, "unknown")
	normalized.ChangeID = trimOrDefault(normalized.ChangeID, "n/a")
	normalized.OpenFlagURL = trimOrDefault(normalized.OpenFlagURL, "https://example.com/flags/checkout.v2")
	normalized.RollbackNowURL = strings.TrimSpace(normalized.RollbackNowURL)
	normalized.IncidentBoardURL = strings.TrimSpace(normalized.IncidentBoardURL)

	return normalized
}

func trimOrDefault(value, fallback string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return fallback
	}

	return trimmed
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}

func normalizeStatus(value string) string {
	status := strings.TrimSpace(strings.ToLower(value))
	if status == "at-risk" || status == "rolled-back" {
		return status
	}

	return "healthy"
}
