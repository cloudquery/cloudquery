package cmd

import (
	"encoding/json"
	"io"
	"slices"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
)

type upgradeReportsJSON struct {
	Reports []upgradeReportJSON `json:"reports"`
}

type upgradeReportJSON struct {
	Source            upgradeSourceJSON             `json:"source"`
	Destination       *upgradeDestinationJSON       `json:"destination"`
	Verdict           string                        `json:"verdict"`
	Summary           string                        `json:"summary,omitempty"`
	Tables            []upgradeTableImpact          `json:"tables"`
	RemovedTables     []string                      `json:"removed_tables"`
	OutputComparisons []upgradeOutputComparisonJSON `json:"output_comparisons"`
	CoverageGaps      []string                      `json:"coverage_gaps"`
	Action            string                        `json:"action"`
}

type upgradeSourceJSON struct {
	Name        string                `json:"name"`
	FromVersion string                `json:"from_version"`
	ToVersion   string                `json:"to_version"`
	From        upgradeSourceSideJSON `json:"from"`
	To          upgradeSourceSideJSON `json:"to"`
}

type upgradeSourceSideJSON struct {
	Registry string `json:"registry"`
	Path     string `json:"path"`
	Version  string `json:"version"`
}

type upgradeDestinationJSON struct {
	Name        string `json:"name"`
	Registry    string `json:"registry"`
	Path        string `json:"path"`
	Version     string `json:"version,omitempty"`
	WriteMode   string `json:"write_mode"`
	PKMode      string `json:"pk_mode"`
	MigrateMode string `json:"migrate_mode"`
}

type upgradeOutputComparisonJSON struct {
	Table    string                `json:"table"`
	Column   string                `json:"column,omitempty"`
	OldType  string                `json:"old_type,omitempty"`
	NewType  string                `json:"new_type,omitempty"`
	Evidence []upgradeEvidenceJSON `json:"evidence"`
}

type upgradeEvidenceJSON struct {
	SyntheticValue string `json:"synthetic_value"`
	Before         string `json:"before"`
	After          string `json:"after"`
}

func renderUpgradeReportsJSON(w io.Writer, reports []upgradeReport) error {
	out := upgradeReportsJSON{Reports: make([]upgradeReportJSON, len(reports))}
	for i, report := range reports {
		out.Reports[i] = upgradeReportToJSON(report)
	}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	return encoder.Encode(out)
}

func upgradeReportToJSON(r upgradeReport) upgradeReportJSON {
	r = sortUpgradeReport(r)
	impacts := upgradeTableImpacts(r)
	comparisons := upgradeOutputComparisons(r.Findings)
	migrateMode := specs.MigrateModeSafe
	if r.Destination != nil {
		migrateMode = r.Destination.MigrateMode
	}

	verdict, _ := upgradeVerdict(r, impacts, comparisons, migrateMode)
	action, _ := upgradeAction(r, impacts, comparisons, migrateMode)
	out := upgradeReportJSON{
		Source: upgradeSourceJSON{
			Name:        r.SourceName,
			FromVersion: r.FromVersion,
			ToVersion:   r.ToVersion,
			From:        upgradeSourceSideJSON{Registry: r.FromOrigin.Registry.String(), Path: r.FromOrigin.Path, Version: r.FromVersion},
			To:          upgradeSourceSideJSON{Registry: r.ToOrigin.Registry.String(), Path: r.ToOrigin.Path, Version: r.ToVersion},
		},
		Verdict:           verdict,
		Tables:            append([]upgradeTableImpact{}, impacts...),
		RemovedTables:     upgradeRemovedTableNames(r.RemovedTables),
		OutputComparisons: upgradeOutputComparisonsToJSON(comparisons),
		CoverageGaps:      append([]string{}, upgradeCoverageGaps(r)...),
		Action:            action,
	}
	if verdict == "" {
		out.Verdict = upgradeNoChangesText
	} else {
		out.Summary = upgradeSummary(r, impacts, migrateMode)
	}
	if d := r.Destination; d != nil {
		out.Destination = &upgradeDestinationJSON{
			Name:        d.Name,
			Registry:    d.Registry.String(),
			Path:        d.Path,
			Version:     d.Version,
			WriteMode:   d.WriteMode.String(),
			PKMode:      d.PKMode.String(),
			MigrateMode: d.MigrateMode.String(),
		}
	}
	return out
}

func upgradeOutputComparisonsToJSON(comparisons []*pluginPb.AssessTables_TableFinding) []upgradeOutputComparisonJSON {
	out := []upgradeOutputComparisonJSON{}
	for _, finding := range comparisons {
		for _, column := range finding.Columns {
			if len(column.Evidence) > 0 {
				out = append(out, upgradeOutputComparisonJSON{
					Table:    finding.TableName,
					Column:   column.ColumnName,
					OldType:  column.OldType,
					NewType:  column.NewType,
					Evidence: upgradeEvidenceToJSON(column.Evidence),
				})
			}
		}
		if len(finding.Evidence) > 0 {
			out = append(out, upgradeOutputComparisonJSON{Table: finding.TableName, Evidence: upgradeEvidenceToJSON(finding.Evidence)})
		}
	}
	slices.SortStableFunc(out, func(a, b upgradeOutputComparisonJSON) int {
		return upgradeDiffersFirst(upgradeEvidenceJSONDiffers(a.Evidence), upgradeEvidenceJSONDiffers(b.Evidence))
	})
	return out
}

func upgradeEvidenceJSONDiffers(evidence []upgradeEvidenceJSON) bool {
	return slices.ContainsFunc(evidence, func(e upgradeEvidenceJSON) bool { return e.Before != e.After })
}

func upgradeEvidenceToJSON(evidence []*pluginPb.AssessTables_Evidence) []upgradeEvidenceJSON {
	out := make([]upgradeEvidenceJSON, len(evidence))
	for i, e := range evidence {
		out[i] = upgradeEvidenceJSON{SyntheticValue: e.SyntheticValue, Before: e.Before, After: e.After}
	}
	return out
}
