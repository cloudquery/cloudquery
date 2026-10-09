package cmd

import (
	"cmp"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fatih/color"
)

var (
	upgradeRed         = color.New(color.FgRed)
	upgradeYellow      = color.New(color.FgYellow)
	upgradeGreen       = color.New(color.FgGreen)
	upgradeFaint       = color.New(color.Faint)
	upgradeDestination = color.New(color.Bold, color.FgCyan)
)

type upgradeReport struct {
	SourceName             string
	FromVersion            string
	ToVersion              string
	FromOrigin             upgradeSourceOrigin
	ToOrigin               upgradeSourceOrigin
	SourceUnknown          bool
	SourceGaps             []string
	RemovedTables          []upgradeRemovedTable
	UnmatchedTablePatterns []string
	Destination            *specs.Destination
	Tables                 map[string]upgradeTablePair
	Findings               []*pluginPb.AssessTables_TableFinding
}

type upgradeChangeKind string

const (
	upgradeColumnAdded   upgradeChangeKind = "added"
	upgradeColumnRemoved upgradeChangeKind = "removed"
	upgradeColumnChanged upgradeChangeKind = "changed"
)

type upgradeColumnChange struct {
	Kind          upgradeChangeKind `json:"kind"`
	Column        string            `json:"column"`
	OldType       string            `json:"old_type,omitempty"`
	NewType       string            `json:"new_type,omitempty"`
	OldSourceType string            `json:"old_source_type,omitempty"`
	NewSourceType string            `json:"new_source_type,omitempty"`
	Reason        string            `json:"reason,omitempty"`
}

type upgradeOutcomeResult string

const (
	upgradeOutcomeApplied  upgradeOutcomeResult = "applied"
	upgradeOutcomeFails    upgradeOutcomeResult = "fails"
	upgradeOutcomeDataLoss upgradeOutcomeResult = "data_loss"
)

type upgradeOutcome struct {
	Result upgradeOutcomeResult `json:"result"`
	Text   string               `json:"text"`
}

type upgradeTableImpact struct {
	Name     string                         `json:"name"`
	Category pluginPb.AssessTables_Category `json:"-"`
	NewTable bool                           `json:"new_table,omitempty"`
	Changes  []upgradeColumnChange          `json:"changes,omitempty"`
	Outcomes map[string]upgradeOutcome      `json:"outcomes,omitempty"`
}

var upgradeMigrateModes = []specs.MigrateMode{specs.MigrateModeSafe, specs.MigrateModeForced}

const (
	upgradeNoChangesText    = "No schema changes affect your selected tables."
	upgradeNoOutputDiffText = "NO OUTPUT DIFFERENCE DETECTED for equivalent test values"
	upgradeOutputDiffText   = "OUTPUT DIFFERENCE DETECTED for some test values"
	upgradeNotAssessedText  = "Actual source values and behavior were not assessed."
	upgradePreviewOnlyText  = "This check only previews the changes. It does not migrate, write, delete or upload anything."
	upgradeRemovedTableText = "removed table, the new source version no longer provides it"
	upgradeDataLossText     = "dropped and recreated, existing rows deleted"
	upgradeSafeModeFallback = "safe mode cannot apply these changes"
	upgradeMaxTypeWidth     = 40
)

func renderUpgradeReport(w io.Writer, r upgradeReport) error {
	r = sortUpgradeReport(r)
	impacts := upgradeTableImpacts(r)
	comparisons := upgradeOutputComparisons(r.Findings)
	gaps := upgradeCoverageGaps(r)
	migrateMode := specs.MigrateModeSafe
	if r.Destination != nil {
		migrateMode = r.Destination.MigrateMode
	}

	var b strings.Builder
	fromLabel, toLabel := upgradeSourceLabels(r.FromOrigin, r.FromVersion, r.ToOrigin, r.ToVersion)
	header := fmt.Sprintf("%s %s → %s", r.SourceName, fromLabel, toLabel)
	if r.Destination == nil {
		b.WriteString(bold.Sprint(header) + "\n")
	} else {
		b.WriteString(bold.Sprint(header+" | ") + upgradeDestination.Sprint(r.Destination.VersionString()) + "\n")
		settings := fmt.Sprintf("write_mode: %s | pk_mode: %s", r.Destination.WriteMode, r.Destination.PKMode)
		if upgradeHasMigrateModes(r.Findings) {
			settings += " | migrate_mode: " + r.Destination.MigrateMode.String()
		}
		b.WriteString(upgradeFaint.Sprint(settings) + "\n")
	}

	verdict, verdictColor := upgradeVerdict(r, impacts, comparisons, migrateMode)
	if verdict == "" {
		b.WriteString(upgradeGreen.Sprint(upgradeNoChangesText) + "\n")
		writeUpgradeCoverageGaps(&b, gaps)
		b.WriteString("\n")
		_, err := io.WriteString(w, b.String())
		return err
	}
	if summary := upgradeSummary(r, impacts, migrateMode); summary != "" {
		verdict += " — " + summary
	}
	b.WriteString(verdictColor.Sprint(verdict) + "\n")

	writeUpgradeChanges(&b, impacts, upgradeRemovedTableNames(r.RemovedTables))
	writeUpgradeOutputComparisons(&b, comparisons, r.Tables)
	writeUpgradeNextSync(&b, impacts, migrateMode)
	writeUpgradeCoverageGaps(&b, gaps)

	action, actionColor := upgradeAction(r, impacts, comparisons, migrateMode)
	b.WriteString("\n" + color.New(color.Bold, actionColor).Sprint("Action: "+action) + "\n")
	b.WriteString(upgradeFaint.Sprint(upgradePreviewOnlyText) + "\n\n")
	_, err := io.WriteString(w, b.String())
	return err
}

const (
	upgradeExitNoAction     = 0
	upgradeExitActionNeeded = 3
	upgradeExitUnknown      = 4
)

func upgradeExitCode(r upgradeReport) int {
	actionNeeded := len(r.RemovedTables) > 0 || upgradeOutputDiffers(upgradeOutputComparisons(r.Findings)) ||
		slices.ContainsFunc(r.Findings, func(finding *pluginPb.AssessTables_TableFinding) bool {
			return finding.Category == pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED || finding.Category == pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED
		})
	switch {
	case r.SourceUnknown || upgradeNotFullyAssessedTables(r) > 0:
		return upgradeExitUnknown
	case actionNeeded:
		return upgradeExitActionNeeded
	}
	return upgradeExitNoAction
}

func sortUpgradeReport(r upgradeReport) upgradeReport {
	r.RemovedTables = slices.SortedFunc(slices.Values(r.RemovedTables), func(a, b upgradeRemovedTable) int {
		return strings.Compare(a.Name, b.Name)
	})
	r.Findings = slices.SortedFunc(slices.Values(r.Findings), func(a, b *pluginPb.AssessTables_TableFinding) int {
		return strings.Compare(a.TableName, b.TableName)
	})
	return r
}

func upgradeHasMigrateModes(findings []*pluginPb.AssessTables_TableFinding) bool {
	return slices.ContainsFunc(findings, func(finding *pluginPb.AssessTables_TableFinding) bool {
		return finding.Category == pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED ||
			finding.Category == pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE ||
			finding.SafeModeBehavior != "" || finding.ForcedModeBehavior != ""
	})
}

func upgradeTableImpacts(r upgradeReport) []upgradeTableImpact {
	hasMigrateModes := upgradeHasMigrateModes(r.Findings)
	var impacts []upgradeTableImpact
	for _, finding := range r.Findings {
		if finding.Category == pluginPb.AssessTables_CATEGORY_NO_CHANGE || finding.Category == pluginPb.AssessTables_CATEGORY_UNKNOWN {
			continue
		}
		table := r.Tables[finding.TableName]
		impact := upgradeTableImpact{
			Name:     finding.TableName,
			Category: finding.Category,
			NewTable: table.From == nil && table.To != nil,
			Changes:  upgradeColumnChanges(finding, table),
		}
		if hasMigrateModes {
			impact.Outcomes = upgradeOutcomes(impact, finding)
		}
		impacts = append(impacts, impact)
	}
	return impacts
}

func upgradeColumnChanges(finding *pluginPb.AssessTables_TableFinding, table upgradeTablePair) []upgradeColumnChange {
	var changes []upgradeColumnChange
	for _, column := range finding.Columns {
		if column.Category == pluginPb.AssessTables_CATEGORY_NO_CHANGE {
			continue
		}
		oldColumn, newColumn := upgradeTableColumn(table.From, column.ColumnName), upgradeTableColumn(table.To, column.ColumnName)
		change := upgradeColumnChange{
			Kind:          upgradeColumnChanged,
			Column:        column.ColumnName,
			OldType:       column.OldType,
			NewType:       column.NewType,
			OldSourceType: upgradeSourceTypeName(oldColumn),
			NewSourceType: upgradeSourceTypeName(newColumn),
			Reason:        upgradeColumnChangeReason(oldColumn, newColumn),
		}
		switch {
		case column.OldType == "":
			change.Kind = upgradeColumnAdded
		case column.NewType == "":
			change.Kind = upgradeColumnRemoved
		}
		changes = append(changes, change)
	}
	return changes
}

func upgradeTableColumn(table *schema.Table, columnName string) *schema.Column {
	if table == nil {
		return nil
	}
	return table.Columns.Get(columnName)
}

func upgradeSourceTypeName(column *schema.Column) string {
	if column == nil {
		return ""
	}
	return upgradeArrowTypeName(column.Type)
}

func upgradeArrowTypeName(dataType arrow.DataType) string {
	switch dataType := dataType.(type) {
	case *arrow.MapType:
		return dataType.String()
	case *arrow.StringType, *arrow.LargeStringType, *arrow.StringViewType:
		return "string"
	case arrow.ListLikeType:
		return "list<" + upgradeArrowTypeName(dataType.Elem()) + ">"
	}
	return dataType.String()
}

func upgradeColumnChangeReason(oldColumn, newColumn *schema.Column) string {
	switch {
	case oldColumn == nil && newColumn != nil && newColumn.PrimaryKey:
		return "part of the primary key"
	case oldColumn != nil && newColumn == nil && oldColumn.PrimaryKey:
		return "was part of the primary key"
	case oldColumn == nil || newColumn == nil:
		return ""
	case !oldColumn.PrimaryKey && newColumn.PrimaryKey:
		return "now part of the primary key"
	case oldColumn.PrimaryKey && !newColumn.PrimaryKey:
		return "primary key removed"
	case !oldColumn.NotNull && newColumn.NotNull:
		return "now NOT NULL"
	}
	return ""
}

func upgradeOutcomes(impact upgradeTableImpact, finding *pluginPb.AssessTables_TableFinding) map[string]upgradeOutcome {
	switch impact.Category {
	case pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED:
		return map[string]upgradeOutcome{
			specs.MigrateModeSafe.String():   {upgradeOutcomeFails, "fails: " + upgradeSafeModeFailure(impact.Changes, finding)},
			specs.MigrateModeForced.String(): {upgradeOutcomeDataLoss, upgradeDataLossText},
		}
	case pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE:
		applied := upgradeOutcome{upgradeOutcomeApplied, "migrated in place"}
		if impact.NewTable {
			applied.Text = "created"
		}
		return map[string]upgradeOutcome{specs.MigrateModeSafe.String(): applied, specs.MigrateModeForced.String(): applied}
	}
	return nil
}

const (
	upgradeSafeModePrimaryKey = "safe mode cannot change a primary key"
	upgradeSafeModeColumnType = "safe mode cannot change a column type"
	upgradeSafeModeNotNull    = "safe mode cannot make a column NOT NULL"
)

func upgradeSafeModeBlocker(change upgradeColumnChange) string {
	switch {
	case strings.Contains(change.Reason, "primary key"):
		return upgradeSafeModePrimaryKey
	case change.Kind == upgradeColumnChanged && change.OldType != change.NewType:
		return upgradeSafeModeColumnType
	case change.Reason == "now NOT NULL":
		return upgradeSafeModeNotNull
	}
	return ""
}

func upgradeSafeModeFailure(changes []upgradeColumnChange, finding *pluginPb.AssessTables_TableFinding) string {
	for _, blocker := range []string{upgradeSafeModePrimaryKey, upgradeSafeModeColumnType, upgradeSafeModeNotNull} {
		if slices.ContainsFunc(changes, func(change upgradeColumnChange) bool { return upgradeSafeModeBlocker(change) == blocker }) {
			return blocker
		}
	}
	return cmp.Or(finding.SafeModeBehavior, upgradeSafeModeFallback)
}

func upgradeOutputComparisons(findings []*pluginPb.AssessTables_TableFinding) []*pluginPb.AssessTables_TableFinding {
	var comparisons []*pluginPb.AssessTables_TableFinding
	for _, finding := range findings {
		hasEvidence := len(finding.Evidence) > 0 || slices.ContainsFunc(finding.Columns, func(column *pluginPb.AssessTables_ColumnFinding) bool {
			return len(column.Evidence) > 0
		})
		if finding.Category == pluginPb.AssessTables_CATEGORY_NO_CHANGE && hasEvidence {
			comparisons = append(comparisons, finding)
		}
	}
	return comparisons
}

func upgradeCoverageGaps(r upgradeReport) []string {
	gaps := slices.Clone(r.SourceGaps)
	for _, finding := range r.Findings {
		if finding.IncompleteCoverageReason != "" {
			gaps = append(gaps, fmt.Sprintf("%s: %s", finding.TableName, finding.IncompleteCoverageReason))
		}
	}
	return gaps
}

func upgradeImpactsWithOutcome(impacts []upgradeTableImpact, migrateMode specs.MigrateMode, result upgradeOutcomeResult) []upgradeTableImpact {
	return slices.DeleteFunc(slices.Clone(impacts), func(impact upgradeTableImpact) bool {
		return impact.Outcomes[migrateMode.String()].Result != result
	})
}

func upgradeHasCategory(impacts []upgradeTableImpact, category pluginPb.AssessTables_Category) bool {
	return slices.ContainsFunc(impacts, func(impact upgradeTableImpact) bool { return impact.Category == category })
}

func upgradeNotFullyAssessedTables(r upgradeReport) int {
	notAssessed := 0
	for _, finding := range r.Findings {
		if finding.Category == pluginPb.AssessTables_CATEGORY_UNKNOWN || finding.IncompleteCoverageReason != "" {
			notAssessed++
		}
	}
	return notAssessed
}

func upgradeOutputDiffers(comparisons []*pluginPb.AssessTables_TableFinding) bool {
	differs := func(e *pluginPb.AssessTables_Evidence) bool { return e.Before != e.After }
	return slices.ContainsFunc(comparisons, func(finding *pluginPb.AssessTables_TableFinding) bool {
		return slices.ContainsFunc(finding.Evidence, differs) || slices.ContainsFunc(finding.Columns, func(column *pluginPb.AssessTables_ColumnFinding) bool {
			return slices.ContainsFunc(column.Evidence, differs)
		})
	})
}

func upgradeVerdict(r upgradeReport, impacts []upgradeTableImpact, comparisons []*pluginPb.AssessTables_TableFinding, migrateMode specs.MigrateMode) (string, *color.Color) {
	needsReview := len(upgradeImpactsWithOutcome(impacts, migrateMode, upgradeOutcomeFails)) > 0 ||
		len(upgradeImpactsWithOutcome(impacts, migrateMode, upgradeOutcomeDataLoss)) > 0
	switch {
	case needsReview:
		return "REVIEW REQUIRED", upgradeRed
	case len(r.RemovedTables) > 0:
		return "SELECTED TABLES REMOVED", upgradeYellow
	case upgradeHasCategory(impacts, pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED):
		return "FILE SCHEMA CHANGED", upgradeYellow
	case upgradeOutputDiffers(comparisons):
		return "FILE OUTPUT CHANGED", upgradeYellow
	case r.SourceUnknown || upgradeNotFullyAssessedTables(r) > 0:
		return "UNKNOWN", upgradeYellow
	case len(impacts) > 0:
		return "AUTOMATICALLY MIGRATABLE", upgradeYellow
	case len(comparisons) > 0:
		return upgradeNoOutputDiffText, upgradeGreen
	}
	return "", nil
}

func upgradeSummary(r upgradeReport, impacts []upgradeTableImpact, migrateMode specs.MigrateMode) string {
	var failing, lossy, changed, added int
	for _, impact := range impacts {
		switch {
		case impact.Outcomes[migrateMode.String()].Result == upgradeOutcomeFails:
			failing++
		case impact.Outcomes[migrateMode.String()].Result == upgradeOutcomeDataLoss:
			lossy++
		case impact.NewTable:
			added++
		default:
			changed++
		}
	}
	var parts []string
	for _, part := range []struct {
		count            int
		singular, plural string
	}{
		{failing, "table needs a manual migration", "tables need a manual migration"},
		{lossy, "table will be recreated, deleting existing rows", "tables will be recreated, deleting existing rows"},
		{changed, "changed table", "changed tables"},
		{added, "new table", "new tables"},
		{len(r.RemovedTables), "removed table", "removed tables"},
		{upgradeNotFullyAssessedTables(r), "table not fully assessed", "tables not fully assessed"},
	} {
		switch {
		case part.count == 1:
			parts = append(parts, "1 "+part.singular)
		case part.count > 1:
			parts = append(parts, fmt.Sprintf("%d %s", part.count, part.plural))
		}
	}
	return strings.Join(parts, ", ")
}

func upgradeAction(r upgradeReport, impacts []upgradeTableImpact, comparisons []*pluginPb.AssessTables_TableFinding, migrateMode specs.MigrateMode) (string, color.Attribute) {
	if failing := upgradeImpactsWithOutcome(impacts, migrateMode, upgradeOutcomeFails); len(failing) > 0 {
		possessive := "its"
		if len(failing) > 1 {
			possessive = "their"
		}
		return fmt.Sprintf("migrate %s manually before upgrading, or switch to migrate_mode: forced and accept losing %s rows.", upgradeTableList(failing), possessive), color.FgRed
	}
	if lossy := upgradeImpactsWithOutcome(impacts, migrateMode, upgradeOutcomeDataLoss); len(lossy) > 0 {
		return fmt.Sprintf("the next sync drops and recreates %s and deletes existing rows; back up any data you need before upgrading.", upgradeTableList(lossy)), color.FgRed
	}
	switch {
	case len(r.RemovedTables) > 0:
		return upgradeRemovedTablesAction(r), color.FgYellow
	case upgradeHasCategory(impacts, pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED), upgradeOutputDiffers(comparisons):
		return "review readers that combine old and new files.", color.FgYellow
	case r.SourceUnknown || upgradeNotFullyAssessedTables(r) > 0:
		return "review the source changelog for what this check could not assess.", color.FgYellow
	case len(impacts) > 0 && migrateMode == specs.MigrateModeForced:
		return "no action needed; the next sync applies these changes.", color.FgGreen
	case len(impacts) > 0:
		return "use safe migration.", color.FgGreen
	case len(comparisons) > 0:
		return "no action needed; the output is the same for equivalent values.", color.FgGreen
	}
	return "no action needed.", color.FgGreen
}

func upgradeRemovedTablesAction(r upgradeReport) string {
	var selections []string
	for _, table := range r.RemovedTables {
		if table.SelectedBy == upgradeSelectedByName {
			selections = append(selections, table.Name)
		}
	}
	selections = append(selections, r.UnmatchedTablePatterns...)
	if len(selections) == 0 {
		return "update dependent consumers; the removed tables stop syncing, and their existing data in the destination is kept but no longer updated."
	}
	return fmt.Sprintf("remove explicit selections (%s) and update dependent consumers.", upgradeShortList(selections))
}

func upgradeShortList(names []string) string {
	if len(names) > 3 {
		return fmt.Sprintf("%s and %d more", strings.Join(names[:3], ", "), len(names)-3)
	}
	return strings.Join(names, ", ")
}

func upgradeRemovedTableNames(tables []upgradeRemovedTable) []string {
	names := make([]string, len(tables))
	for i, table := range tables {
		names[i] = table.Name
	}
	return names
}

func upgradeTableList(impacts []upgradeTableImpact) string {
	if len(impacts) > 3 {
		return fmt.Sprintf("these %d tables", len(impacts))
	}
	names := make([]string, len(impacts))
	for i, impact := range impacts {
		names[i] = impact.Name
	}
	return strings.Join(names, ", ")
}

func writeUpgradeChanges(b *strings.Builder, impacts []upgradeTableImpact, removedTables []string) {
	if len(impacts) == 0 && len(removedTables) == 0 {
		return
	}
	tableWidth, columnWidth, typeWidth := 0, 0, 0
	for _, impact := range impacts {
		tableWidth = max(tableWidth, len(impact.Name))
		for _, change := range impact.Changes {
			columnWidth = max(columnWidth, len(change.Column))
			typeWidth = max(typeWidth, min(len([]rune(upgradeChangeType(change))), upgradeMaxTypeWidth))
		}
	}
	for _, table := range removedTables {
		tableWidth = max(tableWidth, len(table))
	}

	b.WriteString("\nChanges\n")
	for _, impact := range impacts {
		if impact.NewTable {
			fmt.Fprintf(b, "  %s%s%s\n", bold.Sprint(impact.Name), strings.Repeat(" ", tableWidth-len(impact.Name)+3), upgradeGreen.Sprint("new table"))
			continue
		}
		b.WriteString("  " + bold.Sprint(impact.Name) + "\n")
		writeUpgradeColumnChanges(b, impact, columnWidth, typeWidth)
	}
	for _, table := range removedTables {
		fmt.Fprintf(b, "  %s%s%s\n", bold.Sprint(table), strings.Repeat(" ", tableWidth-len(table)+3), upgradeRed.Sprint(upgradeRemovedTableText))
	}
}

func writeUpgradeColumnChanges(b *strings.Builder, impact upgradeTableImpact, columnWidth, typeWidth int) {
	for _, change := range impact.Changes {
		marker, description := "~", cmp.Or(change.Reason, "type changed")
		switch change.Kind {
		case upgradeColumnAdded:
			marker, description = "+", upgradeJoinNonEmpty("new column", change.Reason)
		case upgradeColumnRemoved:
			marker, description = "-", upgradeJoinNonEmpty("column removed", change.Reason)
		}
		if sourceType := upgradeChangeType(upgradeColumnChange{OldType: change.OldSourceType, NewType: change.NewSourceType}); sourceType != "" {
			description += " (" + sourceType + ")"
		}
		lineColor := upgradeChangeColor(impact, change)
		changeType := upgradeChangeType(change)
		line := fmt.Sprintf("%s %-*s   %s%s   %s", marker, columnWidth, change.Column, changeType, strings.Repeat(" ", max(typeWidth-len([]rune(changeType)), 0)), description)
		b.WriteString("    " + lineColor.Sprint(line) + "\n")
	}
}

func upgradeChangeColor(impact upgradeTableImpact, change upgradeColumnChange) *color.Color {
	switch {
	case impact.Outcomes == nil:
		return upgradeYellow
	case impact.Category == pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED && upgradeSafeModeBlocker(change) != "":
		return upgradeRed
	case change.Kind == upgradeColumnRemoved:
		return upgradeYellow
	}
	return upgradeGreen
}

func upgradeChangeType(change upgradeColumnChange) string {
	if change.OldType == "" || change.NewType == "" || change.OldType == change.NewType {
		return cmp.Or(change.NewType, change.OldType)
	}
	return change.OldType + " → " + change.NewType
}

func upgradeJoinNonEmpty(parts ...string) string {
	return strings.Join(slices.DeleteFunc(parts, func(part string) bool { return part == "" }), ", ")
}

type upgradeOutputComparisonGroup struct {
	typeChange string
	names      []string
	evidence   []*pluginPb.AssessTables_Evidence
}

func groupUpgradeOutputComparisons(comparisons []*pluginPb.AssessTables_TableFinding, tables map[string]upgradeTablePair) []*upgradeOutputComparisonGroup {
	var groups []*upgradeOutputComparisonGroup
	add := func(typeChange, name string, evidence []*pluginPb.AssessTables_Evidence) {
		i := slices.IndexFunc(groups, func(group *upgradeOutputComparisonGroup) bool { return group.typeChange == typeChange })
		if i < 0 {
			groups = append(groups, &upgradeOutputComparisonGroup{typeChange: typeChange})
			i = len(groups) - 1
		}
		groups[i].names = append(groups[i].names, name)
		groups[i].evidence = append(groups[i].evidence, evidence...)
	}
	for _, finding := range comparisons {
		for _, column := range finding.Columns {
			if len(column.Evidence) > 0 {
				table := tables[finding.TableName]
				oldType := cmp.Or(upgradeSourceTypeName(upgradeTableColumn(table.From, column.ColumnName)), column.OldType)
				newType := cmp.Or(upgradeSourceTypeName(upgradeTableColumn(table.To, column.ColumnName)), column.NewType)
				add(upgradeChangeType(upgradeColumnChange{OldType: oldType, NewType: newType}), finding.TableName+"."+column.ColumnName, column.Evidence)
			}
		}
		if len(finding.Evidence) > 0 {
			add("", finding.TableName, finding.Evidence)
		}
	}
	return groups
}

func writeUpgradeOutputComparisons(b *strings.Builder, comparisons []*pluginPb.AssessTables_TableFinding, tables map[string]upgradeTablePair) {
	if len(comparisons) == 0 {
		return
	}
	b.WriteString("\nOutput comparison\n")
	for _, group := range groupUpgradeOutputComparisons(comparisons, tables) {
		heading := bold.Sprint(upgradeNameList(group.names))
		if group.typeChange != "" {
			heading = group.typeChange + "   " + heading
		}
		b.WriteString("  " + heading + "\n")
		identical := 0
		var identicalValues []string
		for _, e := range group.evidence {
			if e.Before != e.After {
				writeUpgradeEvidence(b, []*pluginPb.AssessTables_Evidence{e})
				continue
			}
			if !slices.Contains(identicalValues, e.SyntheticValue) {
				identicalValues = append(identicalValues, e.SyntheticValue)
			}
			if identical++; identical == 1 {
				writeUpgradeEvidence(b, []*pluginPb.AssessTables_Evidence{e})
			}
		}
		if len(identicalValues) > 1 {
			fmt.Fprintf(b, "    %d equivalent test values: identical output\n", len(identicalValues))
		}
	}
	if upgradeOutputDiffers(comparisons) {
		b.WriteString("  " + upgradeYellow.Sprint(upgradeOutputDiffText) + "\n")
	} else {
		b.WriteString("  " + upgradeGreen.Sprint(upgradeNoOutputDiffText) + "\n")
	}
	b.WriteString("  " + upgradeNotAssessedText + "\n")
}

func upgradeNameList(names []string) string {
	if len(names) == 1 {
		return names[0]
	}
	shown := names
	suffix := ""
	if len(names) > 3 {
		shown, suffix = names[:3], ", …"
	}
	return fmt.Sprintf("%d columns: %s%s", len(names), strings.Join(shown, ", "), suffix)
}

func writeUpgradeEvidence(b *strings.Builder, evidence []*pluginPb.AssessTables_Evidence) {
	for _, e := range evidence {
		writeUpgradeEvidenceLine(b, "Synthetic value", e.SyntheticValue)
		writeUpgradeEvidenceLine(b, "Before", e.Before)
		writeUpgradeEvidenceLine(b, "After", e.After)
	}
}

func writeUpgradeEvidenceLine(b *strings.Builder, label, value string) {
	if value != "" {
		fmt.Fprintf(b, "    %s %s\n", upgradeFaint.Sprintf("%-16s", label+":"), value)
	}
}

func writeUpgradeNextSync(b *strings.Builder, impacts []upgradeTableImpact, configured specs.MigrateMode) {
	withOutcomes := slices.DeleteFunc(slices.Clone(impacts), func(impact upgradeTableImpact) bool { return impact.Outcomes == nil })
	if len(withOutcomes) == 0 {
		return
	}
	tableWidth := 0
	for _, impact := range withOutcomes {
		tableWidth = max(tableWidth, len(impact.Name))
	}

	modes := []specs.MigrateMode{configured}
	for _, mode := range upgradeMigrateModes {
		if mode != configured {
			modes = append(modes, mode)
		}
	}

	b.WriteString("\nNext sync\n")
	for _, mode := range modes {
		if mode != configured && upgradeOutcomesMatch(withOutcomes, mode, configured) {
			fmt.Fprintf(b, "  %s\n", bold.Sprintf("migrate_mode: %s — same as %s", mode, configured))
			continue
		}
		label := "migrate_mode: " + mode.String()
		if mode == configured {
			label += " (your config)"
		}
		b.WriteString("  " + bold.Sprint(label) + "\n")
		for _, impact := range withOutcomes {
			outcome := impact.Outcomes[mode.String()]
			marker, lineColor := "✓", upgradeGreen
			switch outcome.Result {
			case upgradeOutcomeFails:
				marker, lineColor = "✗", upgradeRed
			case upgradeOutcomeDataLoss:
				marker, lineColor = "!", upgradeYellow
			}
			b.WriteString("    " + lineColor.Sprintf("%s %-*s   %s", marker, tableWidth, impact.Name, outcome.Text) + "\n")
		}
	}
}

func upgradeOutcomesMatch(impacts []upgradeTableImpact, a, b specs.MigrateMode) bool {
	return !slices.ContainsFunc(impacts, func(impact upgradeTableImpact) bool {
		return impact.Outcomes[a.String()] != impact.Outcomes[b.String()]
	})
}

func writeUpgradeCoverageGaps(b *strings.Builder, gaps []string) {
	if len(gaps) == 0 {
		return
	}
	b.WriteString("\nCoverage gaps\n")
	for _, gap := range gaps {
		b.WriteString("  " + upgradeYellow.Sprint(gap) + "\n")
	}
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
