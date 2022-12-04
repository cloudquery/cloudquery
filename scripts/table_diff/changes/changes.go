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

func parseColumnChange(line string) (name string, dataType string) {
	match := columnRegex.FindStringSubmatch(line)
	if match == nil {
		return "", ""
	}
	result := make(map[string]string)
	for i, name := range columnRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result["name"], result["dataType"]
}

func getColumnChanges(file *gitdiff.File, table string) (changes []change) {
	addedColumns := make(map[string]string)
	deletedColumns := make(map[string]string)
	for _, fragment := range file.TextFragments {
		for _, line := range fragment.Lines {
			name, dataType := parseColumnChange(line.Line)
			if name == "" || dataType == "" {
				continue
			}
			switch line.Op {
			case gitdiff.OpAdd:
				addedColumns[name] = dataType
			case gitdiff.OpDelete:
				deletedColumns[name] = dataType
			}
		}
	}
	for deletedName, deletedDataType := range deletedColumns {
		if addedDataType, ok := addedColumns[deletedName]; ok {
			if addedDataType != deletedDataType {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %q: column type changed from %q to %q for %q", table, deletedDataType, addedDataType, deletedName),
					Breaking: true,
				})
			} else {
				changes = append(changes, change{
					Text:     fmt.Sprintf("Table %q: column order changed for %q", table, deletedName),
					Breaking: false,
				})
			}
		} else {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %q: column removed %q from table", table, deletedName),
				Breaking: true,
			})
		}
	}
	for addedName, addedDataType := range addedColumns {
		if _, ok := deletedColumns[addedName]; !ok {
			changes = append(changes, change{
				Text:     fmt.Sprintf("Table %q: column added with name %q and type %q", table, addedName, addedDataType),
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
		return strings.Compare(changes[i].Text, changes[j].Text) < 0
	})
	return changes
}

func getFileChanges(file *gitdiff.File) (changes []change, err error) {
	oldTableName := strings.TrimSuffix(filepath.Base(file.OldName), filepath.Ext(file.OldName))
	newTableName := strings.TrimSuffix(filepath.Base(file.NewName), filepath.Ext(file.NewName))

	switch {
	case file.IsDelete:
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %q was removed", oldTableName),
			Breaking: true,
		})
	case file.IsRename:
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %q was renamed to %q", oldTableName, newTableName),
			Breaking: true,
		})
	case file.IsNew:
		changes = append(changes, change{
			Text:     fmt.Sprintf("Table %q was added", newTableName),
			Breaking: false,
		})
	case file.IsCopy:
		return nil, fmt.Errorf("unhandled IsCopy table diff, %q -> %q", oldTableName, newTableName)
	}

	checkColumnChanges := !file.IsDelete && !file.IsNew
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
