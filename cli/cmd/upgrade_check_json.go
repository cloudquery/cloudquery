package cmd

import (
	"encoding/json"
	"io"
	"strings"

	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
)

type upgradeReportsJSON struct {
	Reports []upgradeReportJSON `json:"reports"`
}

type upgradeReportJSON struct {
	Source       upgradeSourceJSON       `json:"source"`
	Destination  *upgradeDestinationJSON `json:"destination"`
	Verdicts     []upgradeVerdictJSON    `json:"verdicts"`
	Tables       []upgradeTableJSON      `json:"tables"`
	CoverageGaps []string                `json:"coverage_gaps"`
}

type upgradeSourceJSON struct {
	Name        string `json:"name"`
	FromVersion string `json:"from_version"`
	ToVersion   string `json:"to_version"`
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

type upgradeVerdictJSON struct {
	Category       string   `json:"category"`
	Tables         []string `json:"tables"`
	ChangedColumns int      `json:"changed_columns"`
	Action         string   `json:"action"`
}

type upgradeTableJSON struct {
	Name                     string                `json:"name"`
	Category                 string                `json:"category"`
	SafeModeBehavior         string                `json:"safe_mode_behavior,omitempty"`
	ForcedModeBehavior       string                `json:"forced_mode_behavior,omitempty"`
	Columns                  []upgradeColumnJSON   `json:"columns,omitempty"`
	Evidence                 []upgradeEvidenceJSON `json:"evidence,omitempty"`
	CoverageIncomplete       bool                  `json:"coverage_incomplete"`
	CoverageIncompleteReason string                `json:"coverage_incomplete_reason,omitempty"`
}

type upgradeColumnJSON struct {
	Name               string                `json:"name"`
	Category           string                `json:"category"`
	SourceType         upgradeTypeChangeJSON `json:"source_type"`
	DestinationType    upgradeTypeChangeJSON `json:"destination_type"`
	SafeModeBehavior   string                `json:"safe_mode_behavior,omitempty"`
	ForcedModeBehavior string                `json:"forced_mode_behavior,omitempty"`
	Evidence           []upgradeEvidenceJSON `json:"evidence,omitempty"`
}

type upgradeTypeChangeJSON struct {
	Old string `json:"old,omitempty"`
	New string `json:"new,omitempty"`
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
	out := upgradeReportJSON{
		Source:       upgradeSourceJSON{Name: r.SourceName, FromVersion: r.FromVersion, ToVersion: r.ToVersion},
		Verdicts:     []upgradeVerdictJSON{},
		Tables:       []upgradeTableJSON{},
		CoverageGaps: append([]string{}, r.coverageGaps()...),
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
	for _, verdict := range r.verdicts() {
		out.Verdicts = append(out.Verdicts, upgradeVerdictJSON{
			Category:       upgradeCategoryName(verdict.Category),
			Tables:         append([]string{}, verdict.Tables...),
			ChangedColumns: verdict.ChangedColumns,
			Action:         upgradeCategoryTexts[verdict.Category].action,
		})
	}
	for _, finding := range r.sortedFindings() {
		out.Tables = append(out.Tables, upgradeTableFindingToJSON(finding, r.Tables[finding.TableName]))
	}
	return out
}

func upgradeTableFindingToJSON(finding *pluginPb.AssessTables_TableFinding, table upgradeTablePair) upgradeTableJSON {
	out := upgradeTableJSON{
		Name:                     finding.TableName,
		Category:                 upgradeCategoryName(finding.Category),
		SafeModeBehavior:         finding.SafeModeBehavior,
		ForcedModeBehavior:       finding.ForcedModeBehavior,
		Evidence:                 upgradeEvidenceToJSON(finding.Evidence),
		CoverageIncomplete:       finding.CoverageIncomplete,
		CoverageIncompleteReason: finding.CoverageIncompleteReason,
	}
	for _, column := range changedColumnFindings(finding) {
		out.Columns = append(out.Columns, upgradeColumnJSON{
			Name:     column.ColumnName,
			Category: upgradeCategoryName(column.Category),
			SourceType: upgradeTypeChangeJSON{
				Old: sourceColumnType(table.From, column.ColumnName),
				New: sourceColumnType(table.To, column.ColumnName),
			},
			DestinationType:    upgradeTypeChangeJSON{Old: column.OldType, New: column.NewType},
			SafeModeBehavior:   column.SafeModeBehavior,
			ForcedModeBehavior: column.ForcedModeBehavior,
			Evidence:           upgradeEvidenceToJSON(column.Evidence),
		})
	}
	return out
}

func upgradeEvidenceToJSON(evidence []*pluginPb.AssessTables_Evidence) []upgradeEvidenceJSON {
	out := make([]upgradeEvidenceJSON, len(evidence))
	for i, e := range evidence {
		out[i] = upgradeEvidenceJSON{SyntheticValue: e.SyntheticValue, Before: e.Before, After: e.After}
	}
	return out
}

func upgradeCategoryName(category pluginPb.AssessTables_Category) string {
	return strings.ToLower(strings.TrimPrefix(category.String(), "CATEGORY_"))
}
