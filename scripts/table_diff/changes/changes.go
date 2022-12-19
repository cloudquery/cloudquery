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
)

type change struct {
	Text     string `json:"text"`
	Breaking bool   `json:"breaking"`
}

type column struct {
	dataType string
	pk       bool
}

func backtickStrings(strings ...string) []interface{} {
	backticked := make([]interface{}, len(strings))
	for i, s := range strings {
		backticked[i] = fmt.Sprintf("`%s`", s)
	}
	return backticked
}

func parseColumnChange(line string) (name string, dataType string, pk bool) {
	match := columnRegex.FindStringSubmatch(line)
	if match == nil {
		return "", "", false
	}
	result := make(map[string]string)
	for i, name := range columnRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	cleanName := strings.TrimSuffix(result["name"], " (PK)")
	return cleanName, result["dataType"], result["name"] != cleanName
}

func getColumnChanges(file *gitdiff.File, table string) (changes []change) {
	addedColumns := make(map[string]column)
	deletedColumns := make(map[string]column)
	for _, fragment := range file.TextFragments {
		for _, line := range fragment.Lines {
			name, dataType, pk := parseColumnChange(line.Line)
			if name == "" || dataType == "" {
				continue
			}
			column := column{dataType: dataType, pk: pk}
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
			if deletedColumn.dataType == addedColumn.dataType && deletedColumn.pk == addedColumn.pk {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: column order changed for %s", backtickStrings(table, deletedName)...),
					Breaking: false,
				})
				continue
			}

			if addedColumn.dataType != deletedColumn.dataType {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: column type changed from %s to %s for %s", backtickStrings(table, deletedColumn.dataType, addedColumn.dataType, deletedName)...),
					Breaking: false,
				})
			}

			if addedColumn.pk && !deletedColumn.pk {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: primary key constraint added to column %s", backtickStrings(table, deletedName)...),
					Breaking: false,
				})
			}

			if !addedColumn.pk && deletedColumn.pk {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %s: primary key constraint removed from column %s", backtickStrings(table, deletedName)...),
					Breaking: false,
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
			if addedColumn.pk {
				name = fmt.Sprintf("%s (PK)", name)
			}
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %s: column added with name %s and type %s", backtickStrings(table, name, addedColumn.dataType)...),
				Breaking: false,
			})
		}
	}

	sort.SliceStable(changes, func(i, j int) bool {
		iBreaking := changes[i].Breaking
		jBreaking := changes[j].Breaking
		if iBreaking && !jBreaking {
			return true
		}
		if !iBreaking && jBreaking {
			return false
		}
		return changes[i].Text < changes[j].Text
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
