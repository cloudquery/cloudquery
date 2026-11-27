package client

import (
	"fmt"
	"strings"
	"time"
)

const (
	maxTableNameBytes          = 255
	forbiddenTableNameChars    = `*,?/\#"<>| :`
	forbiddenTableNamePrefixes = "_-+"
	partitionSeparator         = "-"
	partitionSuffixLayout      = time.DateOnly
)

func validateTableName(name string) error {
	switch {
	case name == "":
		return fmt.Errorf("invalid table name %q: must not be empty", name)
	case name == "." || name == "..":
		return fmt.Errorf("invalid table name %q: must not be a dot path segment", name)
	case len(name) > maxTableNameBytes:
		return fmt.Errorf("invalid table name %q: must not be longer than %d bytes", name, maxTableNameBytes)
	case strings.ContainsAny(name[:1], forbiddenTableNamePrefixes):
		return fmt.Errorf("invalid table name %q: must not start with any of %q", name, forbiddenTableNamePrefixes)
	case strings.ContainsAny(name, forbiddenTableNameChars):
		return fmt.Errorf("invalid table name %q: must not contain any of %q", name, forbiddenTableNameChars)
	}
	return nil
}

func isTableIndex(tableName, indexName string) bool {
	if indexName == tableName {
		return true
	}
	suffix, found := strings.CutPrefix(indexName, tableName+partitionSeparator)
	if !found || len(suffix) != len(partitionSuffixLayout) {
		return false
	}
	_, err := time.Parse(partitionSuffixLayout, suffix)
	return err == nil
}

func tableIndices(tableName string, indexNames []string) []string {
	var owned []string
	for _, indexName := range indexNames {
		if isTableIndex(tableName, indexName) {
			owned = append(owned, indexName)
		}
	}
	return owned
}
