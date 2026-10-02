package cmd

import (
	"cmp"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fatih/color"
)

var (
	upgradeRed    = color.New(color.FgRed)
	upgradeYellow = color.New(color.FgYellow)
	upgradeGreen  = color.New(color.FgGreen)
)

type upgradeReport struct {
	SourceName    string
	FromVersion   string
	ToVersion     string
	SourceUnknown bool
	SourceGaps    []string
	RemovedTables []string
	Destination   *specs.Destination
	Tables        map[string]upgradeTablePair
	Findings      []*pluginPb.AssessTables_TableFinding
}

type upgradeCategoryText struct {
	verdict string
	action  string
	color   *color.Color
}

var upgradeCategoriesBySeverity = []pluginPb.AssessTables_Category{
	pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
	pluginPb.AssessTables_CATEGORY_TABLE_REMOVED,
	pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
	pluginPb.AssessTables_CATEGORY_UNKNOWN,
	pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
}

var upgradeCategoryTexts = map[pluginPb.AssessTables_Category]upgradeCategoryText{
	pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED: {"REVIEW REQUIRED", "plan a manual migration or rebuild.", upgradeRed},
	pluginPb.AssessTables_CATEGORY_TABLE_REMOVED:             {"SELECTED TABLES REMOVED", "remove explicit selections and update dependent consumers.", upgradeYellow},
	pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED:       {"FILE SCHEMA CHANGED", "review readers that combine old and new files.", upgradeYellow},
	pluginPb.AssessTables_CATEGORY_UNKNOWN:                   {"UNKNOWN", "review the source changelog for what this check could not assess.", upgradeYellow},
	pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE:  {"AUTOMATICALLY MIGRATABLE", "use safe migration.", upgradeYellow},
}

func renderUpgradeReport(w io.Writer, r upgradeReport) error {
	var b strings.Builder
	header := fmt.Sprintf("%s %s → %s", r.SourceName, r.FromVersion, r.ToVersion)
	if r.Destination != nil {
		header += " | " + r.Destination.VersionString()
	}
	b.WriteString(header + "\n")
	if r.Destination != nil {
		fmt.Fprintf(&b, "write_mode: %s | pk_mode: %s\n", r.Destination.WriteMode, r.Destination.PKMode)
	}

	verdicts := r.verdicts()
	gaps := r.coverageGaps()
	if len(verdicts) == 0 {
		b.WriteString(upgradeGreen.Sprint("No schema changes affect your selected tables.") + "\n")
		writeUpgradeCoverageGaps(&b, gaps)
		b.WriteString("\n")
		_, err := io.WriteString(w, b.String())
		return err
	}

	for i, verdict := range verdicts {
		if i > 0 {
			b.WriteString("\n")
		}
		writeUpgradeHeading(&b, verdict.Category, verdict.Tables, verdict.ChangedColumns)
		switch verdict.Category {
		case pluginPb.AssessTables_CATEGORY_TABLE_REMOVED:
			b.WriteString("\n")
			for _, table := range verdict.Tables {
				b.WriteString(bold.Sprint(table) + "\n")
			}
			b.WriteString("The new source version no longer provides these tables.\n")
		case pluginPb.AssessTables_CATEGORY_UNKNOWN:
		default:
			for _, finding := range verdict.Findings {
				writeUpgradeTableFinding(&b, r.Destination.Name, finding, r.Tables[finding.TableName])
			}
		}
	}
	writeUpgradeCoverageGaps(&b, gaps)

	b.WriteString("\n")
	for _, verdict := range verdicts {
		fmt.Fprintf(&b, "Action: %s\n", upgradeCategoryTexts[verdict.Category].action)
	}
	b.WriteString("This check only previews the changes. It does not migrate, write, delete or upload anything.\n\n")
	_, err := io.WriteString(w, b.String())
	return err
}

type upgradeVerdict struct {
	Category       pluginPb.AssessTables_Category
	Tables         []string
	ChangedColumns int
	Findings       []*pluginPb.AssessTables_TableFinding
}

func (r upgradeReport) sortedFindings() []*pluginPb.AssessTables_TableFinding {
	findings := slices.Clone(r.Findings)
	slices.SortFunc(findings, func(a, b *pluginPb.AssessTables_TableFinding) int {
		return strings.Compare(a.TableName, b.TableName)
	})
	return findings
}

func (r upgradeReport) verdicts() []upgradeVerdict {
	findingsByCategory := make(map[pluginPb.AssessTables_Category][]*pluginPb.AssessTables_TableFinding)
	for _, finding := range r.sortedFindings() {
		findingsByCategory[finding.Category] = append(findingsByCategory[finding.Category], finding)
	}

	var verdicts []upgradeVerdict
	for _, category := range upgradeCategoriesBySeverity {
		if category == pluginPb.AssessTables_CATEGORY_TABLE_REMOVED {
			if len(r.RemovedTables) > 0 {
				verdicts = append(verdicts, upgradeVerdict{Category: category, Tables: r.RemovedTables})
			}
			continue
		}
		findings := findingsByCategory[category]
		unknownSource := category == pluginPb.AssessTables_CATEGORY_UNKNOWN && r.SourceUnknown
		if len(findings) == 0 && !unknownSource {
			continue
		}
		verdict := upgradeVerdict{Category: category, Findings: findings}
		for _, finding := range findings {
			verdict.Tables = append(verdict.Tables, finding.TableName)
			verdict.ChangedColumns += len(changedColumnFindings(finding))
		}
		verdicts = append(verdicts, verdict)
	}
	return verdicts
}

func (r upgradeReport) coverageGaps() []string {
	gaps := slices.Clone(r.SourceGaps)
	for _, finding := range r.sortedFindings() {
		if finding.CoverageIncomplete {
			gaps = append(gaps, fmt.Sprintf("%s: %s", finding.TableName, cmp.Or(finding.CoverageIncompleteReason, "coverage incomplete")))
		}
	}
	return gaps
}

func writeUpgradeHeading(b *strings.Builder, category pluginPb.AssessTables_Category, tableNames []string, changedColumns int) {
	text := upgradeCategoryTexts[category]
	heading := text.verdict
	switch {
	case len(tableNames) == 1:
		heading += " — " + tableNames[0]
	case len(tableNames) > 1 && changedColumns > 0:
		heading += fmt.Sprintf(" — %d tables / %d columns", len(tableNames), changedColumns)
	case len(tableNames) > 1:
		heading += fmt.Sprintf(" — %d tables", len(tableNames))
	}
	b.WriteString(text.color.Sprint(heading) + "\n")
}

func writeUpgradeTableFinding(b *strings.Builder, destinationName string, finding *pluginPb.AssessTables_TableFinding, table upgradeTablePair) {
	columns := changedColumnFindings(finding)
	for _, column := range columns {
		fmt.Fprintf(b, "\n%s\n", bold.Sprintf("%s.%s", finding.TableName, column.ColumnName))
		writeUpgradeLine(b, "Source", upgradeTypeChange(sourceColumnType(table.From, column.ColumnName), sourceColumnType(table.To, column.ColumnName)))
		writeUpgradeLine(b, destinationName, upgradeTypeChange(column.OldType, column.NewType))
		writeUpgradeBehavior(b, column.SafeModeBehavior, column.ForcedModeBehavior, column.Evidence)
	}
	safeMode, forcedMode := finding.SafeModeBehavior, finding.ForcedModeBehavior
	columnsShowTableBehavior := len(columns) > 0 && !slices.ContainsFunc(columns, func(column *pluginPb.AssessTables_ColumnFinding) bool {
		return column.SafeModeBehavior != safeMode || column.ForcedModeBehavior != forcedMode
	})
	if columnsShowTableBehavior {
		safeMode, forcedMode = "", ""
	}
	if len(columns) > 0 && safeMode == "" && forcedMode == "" && len(finding.Evidence) == 0 {
		return
	}
	fmt.Fprintf(b, "\n%s\n", bold.Sprint(finding.TableName))
	writeUpgradeBehavior(b, safeMode, forcedMode, finding.Evidence)
}

func writeUpgradeBehavior(b *strings.Builder, safeMode, forcedMode string, evidence []*pluginPb.AssessTables_Evidence) {
	writeUpgradeLine(b, "Safe mode", safeMode)
	writeUpgradeLine(b, "Forced mode", forcedMode)
	for _, e := range evidence {
		writeUpgradeLine(b, "Synthetic value", e.SyntheticValue)
		writeUpgradeLine(b, "Before", e.Before)
		writeUpgradeLine(b, "After", e.After)
	}
}

func writeUpgradeLine(b *strings.Builder, label, value string) {
	if value != "" {
		fmt.Fprintf(b, "  %-12s %s\n", label+":", value)
	}
}

func writeUpgradeCoverageGaps(b *strings.Builder, gaps []string) {
	if len(gaps) == 0 {
		return
	}
	b.WriteString("\nCoverage gaps:\n")
	for _, gap := range gaps {
		b.WriteString("  " + gap + "\n")
	}
}

func changedColumnFindings(finding *pluginPb.AssessTables_TableFinding) []*pluginPb.AssessTables_ColumnFinding {
	return slices.DeleteFunc(slices.Clone(finding.Columns), func(column *pluginPb.AssessTables_ColumnFinding) bool {
		return column.Category == pluginPb.AssessTables_CATEGORY_NO_CHANGE
	})
}

func upgradeTypeChange(oldType, newType string) string {
	if oldType == "" && newType == "" {
		return ""
	}
	return fmt.Sprintf("%s → %s", upgradeRed.Sprint(cmp.Or(oldType, "none")), upgradeGreen.Sprint(cmp.Or(newType, "none")))
}

func sourceColumnType(table *schema.Table, columnName string) string {
	if table == nil {
		return ""
	}
	column := table.Columns.Get(columnName)
	if column == nil {
		return ""
	}
	return column.Type.String()
}

func upgradeTablesByName(tables []upgradeTableSchemas) (map[string]upgradeTablePair, error) {
	byName := make(map[string]upgradeTablePair, len(tables))
	for _, table := range tables {
		from, err := upgradeTableFromBytes(table.From)
		if err != nil {
			return nil, fmt.Errorf("failed to decode table %s: %w", table.Name, err)
		}
		to, err := upgradeTableFromBytes(table.To)
		if err != nil {
			return nil, fmt.Errorf("failed to decode table %s: %w", table.Name, err)
		}
		pair := upgradeTablePair{Name: table.Name, From: from, To: to}
		switch {
		case to != nil:
			pair.Name = to.Name
		case from != nil:
			pair.Name = from.Name
		}
		byName[pair.Name] = pair
	}
	return byName, nil
}

func upgradeTableFromBytes(b []byte) (*schema.Table, error) {
	if len(b) == 0 {
		return nil, nil
	}
	sc, err := pluginPb.NewSchemaFromBytes(b)
	if err != nil {
		return nil, err
	}
	return schema.NewTableFromArrowSchema(sc)
}
