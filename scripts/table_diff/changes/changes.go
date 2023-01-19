package changes

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

var (
	columnRegex = regexp.MustCompile(`^\|(?P<name>.*)\|(?P<dataType>.*)\|`)
	pkRegex     = regexp.MustCompile(`^The composite primary key for this table is \(([^)]+)\)\.`)
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
	dataType   string
	columnType columnType
}

func (c column) pk() bool {
	return c.columnType&columnTypePK != 0
}

func (c column) incremental() bool {
	return c.columnType&columnTypeIncremental != 0
}

func backtickStrings(strings ...string) []interface{} {
	backticked := make([]interface{}, len(strings))
	for i, s := range strings {
		backticked[i] = fmt.Sprintf("`%s`", s)
	}
	return backticked
}

func parseColumnChange(line string) (name string, dataType string, columnType columnType) {
	match := columnRegex.FindStringSubmatch(line)
	if match == nil {
		return "", "", columnType
	}
	result := make(map[string]string)
	for i, name := range columnRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	if strings.Contains(result["name"], " (PK)") {
		columnType |= columnTypePK
	}
	if strings.Contains(result["name"], " (Incremental Key)") {
		columnType |= columnTypeIncremental
	}
	cleanName := strings.Split(result["name"], " (")[0]
	return cleanName, result["dataType"], columnType
}

func parsePKChange(line string) (names []string) {
	match := pkRegex.FindStringSubmatch(line)
	if len(match) != 2 {
		return nil
	}
	for _, part := range strings.Split(match[1], ", ") {
		names = append(names, strings.Trim(part, "*"))
	}
	return
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
			name, dataType, columnType := parseColumnChange(line.Line)
			if name == "" || dataType == "" {
				continue
			}
			column := column{dataType: dataType, columnType: columnType}
			switch line.Op {
			case gitdiff.OpAdd:
				addedColumns[name] = column
			case gitdiff.OpDelete:
				deletedColumns[name] = column
			}
		}
	}
	for deletedName, deletedColumn := range deletedColumns {
		if addedColumn, ok := addedColumns[deletedName]; ok {
			if deletedColumn.dataType == addedColumn.dataType && deletedColumn.columnType == addedColumn.columnType {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: column order changed for %s", backtickStrings(table, deletedName)...),
					Breaking: false,
				})
				continue
			}

			if addedColumn.dataType != deletedColumn.dataType {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: column type changed from %s to %s for %s", backtickStrings(table, deletedColumn.dataType, addedColumn.dataType, deletedName)...),
					Breaking: true,
				})
			}

			if addedColumn.pk() && !deletedColumn.pk() {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: primary key constraint added to column %s", backtickStrings(table, deletedName)...),
					Breaking: false,
				})
			}

			if !addedColumn.pk() && deletedColumn.pk() {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: primary key constraint removed from column %s", backtickStrings(table, deletedName)...),
					Breaking: false,
				})
			}

			if addedColumn.incremental() && !deletedColumn.incremental() {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: column %s added to cursor for incremental syncs", backtickStrings(table, deletedName)...),
					Breaking: true,
				})
			}

			if !addedColumn.incremental() && deletedColumn.incremental() {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: column %s removed from cursor for incremental syncs", backtickStrings(table, deletedName)...),
					Breaking: true,
				})
			}
		} else {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column removed %s from table", backtickStrings(table, deletedName)...),
				Breaking: true,
			})
		}
	}
	for addedName, addedColumn := range addedColumns {
		if _, ok := deletedColumns[addedName]; !ok {
			name := addedName
			if addedColumn.pk() {
				name += " (PK)"
			}
			if addedColumn.incremental() {
				name += " (Incremental Key)"
			}
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column added with name %s and type %s", backtickStrings(table, name, addedColumn.dataType)...),
				Breaking: addedColumn.pk(),
			})
		}
	}

	// check PK:
	if len(addedPK) > 0 && len(addedPK) == len(deletedPK) {
		// if they are unequal the pk added/removed is correct.
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
	case file.IsRename:
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
