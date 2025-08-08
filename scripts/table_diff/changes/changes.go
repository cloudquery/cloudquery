package changes

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var (
	columnRegex = regexp.MustCompile(`^\|(?P<name>.*)\|(?P<dataType>.*)\|`)
	pkRegex     = regexp.MustCompile(`^The composite primary key for this table is \(([^)]+)\)\.`)

	// There is a different message for single PKs
	singlePKRegex = regexp.MustCompile(`^The primary key for this table is ([^)]+)\.`)
)

type change struct {
	Text     string `json:"text"`
	Breaking bool   `json:"breaking"`
}

type columnType int

const (
	columnTypePK columnType = 1 << iota
	columnTypeIncremental
)

type column struct {
	dataType    string
	dataTypeRaw string
	columnType  columnType
}

func (c column) pk() bool {
	return c.columnType&columnTypePK != 0
}

func (c column) incremental() bool {
	return c.columnType&columnTypeIncremental != 0
}

func backtickStrings(strs ...string) []any {
	backticked := make([]any, len(strs))
	for i, s := range strs {
		backticked[i] = "`" + s + "`"
	}
	return backticked
}

func parseColumnChange(line string) (name string, col column) {
	match := columnRegex.FindStringSubmatch(line)
	if match == nil {
		return "", column{}
	}
	result := make(map[string]string)
	for i, name := range columnRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	if strings.Contains(result["name"], " (PK)") {
		col.columnType |= columnTypePK
	}
	if strings.Contains(result["name"], " (Incremental Key)") {
		col.columnType |= columnTypeIncremental
	}
	cleanName := strings.Split(result["name"], " (")[0]
	col.dataTypeRaw = result["dataType"]
	col.dataType = strings.Trim(col.dataTypeRaw, "`")
	return cleanName, col
}

func parsePKChange(line string) (names []string) {
	matchMulti := pkRegex.FindStringSubmatch(line)
	matchSingle := singlePKRegex.FindStringSubmatch(line)
	if len(matchMulti) == 2 {
		for _, part := range strings.Split(matchMulti[1], ", ") {
			names = append(names, strings.Trim(part, "*"))
		}
	}
	if len(matchSingle) == 2 {
		for _, part := range strings.Split(matchSingle[1], ", ") {
			names = append(names, strings.Trim(part, "*"))
		}
	}
	return names
}

func getColumnChanges(file *gitdiff.File, table string) (changes []change) {
	addedColumns := make(map[string]column)
	deletedColumns := make(map[string]column)
	var addedPK, deletedPK []string
	for _, fragment := range file.TextFragments {
		for _, line := range fragment.Lines {
			pkChanges := parsePKChange(line.Line)
			if len(pkChanges) > 0 {
				switch line.Op {
				case gitdiff.OpAdd:
					addedPK = pkChanges
				case gitdiff.OpDelete:
					deletedPK = pkChanges
				}
				continue
			}
			name, col := parseColumnChange(line.Line)
			if name == "" || col.dataType == "" {
				continue
			}
			switch line.Op {
			case gitdiff.OpAdd:
				addedColumns[name] = col
			case gitdiff.OpDelete:
				deletedColumns[name] = col
			}
		}
	}
	for name, deleted := range deletedColumns {
		added, ok := addedColumns[name]
		if !ok {
			if name == "_cq_source_name" || name == "_cq_sync_time" {
				// Ignore removal of these columns for SDK v4 migration; they are now
				// owned by the CLI as an optional transformation.
				continue
			}
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column %s removed from table", backtickStrings(table, name)...),
				Breaking: true,
			})
			continue
		}

		if deleted.dataType != added.dataType {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column type changed from %s to %s for %s", backtickStrings(table, deleted.dataType, added.dataType, name)...),
				Breaking: true,
			})
			continue
		}

		if deleted.columnType == added.columnType {
			// we ignore ordering changes
			continue
		}

		if added.pk() && !deleted.pk() && !(len(addedPK) == 1 && addedPK[0] == "_cq_id" && len(deletedPK) > 0) {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: primary key constraint added to column %s", backtickStrings(table, name)...),
				Breaking: true,
			})
		}

		if !added.pk() && deleted.pk() && !(len(addedPK) == 1 && addedPK[0] == "_cq_id" && len(deletedPK) > 0) {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: primary key constraint removed from column %s", backtickStrings(table, name)...),
				Breaking: true,
			})
		}

		if added.incremental() && !deleted.incremental() {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column %s added to cursor for incremental syncs", backtickStrings(table, name)...),
				Breaking: true,
			})
		}

		if !added.incremental() && deleted.incremental() {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column %s removed from cursor for incremental syncs", backtickStrings(table, name)...),
				Breaking: true,
			})
		}
	}

	for name, added := range addedColumns {
		if _, ok := deletedColumns[name]; ok {
			continue
		}
		if added.pk() {
			name += " (PK)"
		}
		if added.incremental() {
			name += " (Incremental Key)"
		}
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %s: column added with name %s and type %s", backtickStrings(table, name, added.dataType)...),
			Breaking: added.pk(),
		})
	}

	// check PK:
	// Only if all the Columns are the same before and after the change should
	// we consider this a "primary key order" change
	ordering := func(a, b string) bool { return a < b }
	diff := cmp.Diff(addedPK, deletedPK, cmpopts.SortSlices(ordering))
	if len(addedPK) > 0 && diff == "" {
		changes = append(changes, change{
			Text: fmt.Sprintf("Table %s: primary key order changed from %s to %s",
				backtickStrings(
					table,
					strings.Join(deletedPK, ", "),
					strings.Join(addedPK, ", "),
				)...,
			),
			Breaking: true,
		})
	}

	sort.SliceStable(changes, func(i, j int) bool {
		chI := changes[i]
		chJ := changes[j]
		switch {
		case chI.Breaking && !chJ.Breaking:
			return true
		case !chI.Breaking && chJ.Breaking:
			return false
		default:
			return chI.Text < chJ.Text
		}
	})

	if len(addedPK) == 1 && addedPK[0] == "_cq_id" && len(deletedPK) > 0 {
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %s: all existing primary key constraints have been removed and a primary key new constraint has been added to `_cq_id`", backtickStrings(table)...),
			Breaking: true,
		})
	}

	return changes
}

func getFileChanges(file *gitdiff.File) (changes []change, err error) {
	oldTableName := strings.TrimSuffix(filepath.Base(file.OldName), filepath.Ext(file.OldName))
	newTableName := strings.TrimSuffix(filepath.Base(file.NewName), filepath.Ext(file.NewName))

	switch {
	case file.IsDelete:
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %s was removed", backtickStrings(oldTableName)...),
			Breaking: true,
		})
	case file.IsRename && oldTableName != newTableName:
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %s was renamed to %s", backtickStrings(oldTableName, newTableName)...),
			Breaking: true,
		})
	case file.IsNew:
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %s was added", backtickStrings(newTableName)...),
			Breaking: false,
		})
	case file.IsCopy:
		return nil, fmt.Errorf("unhandled IsCopy table diff, %s -> %s", backtickStrings(oldTableName, newTableName)...)
	}

	checkColumnChanges := !file.IsDelete && !file.IsNew
	// Don't report column changes for deleted or new tables to avoid noise
	if checkColumnChanges {
		changes = append(changes, getColumnChanges(file, newTableName)...)
	}

	return changes, nil
}

//nolint:revive
func GetChanges(files []*gitdiff.File) (changes []change, err error) {
	changes = make([]change, 0)
	for _, file := range files {
		fileChanges, err := getFileChanges(file)
		if err != nil {
			return nil, err
		}
		changes = append(changes, fileChanges...)
	}

	return changes, nil
}
