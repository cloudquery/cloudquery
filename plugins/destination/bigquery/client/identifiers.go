package client

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	maxTableNameBytes           = 1024
	forbiddenTableNameChars     = "/\\.`?#%:*"
	forbiddenQuotedIdentifierCh = "`\\"
)

var unquotedColumnNamePattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

func validateTableName(name string) error {
	switch {
	case name == "":
		return errors.New("invalid table name: name is empty")
	case len(name) > maxTableNameBytes:
		return fmt.Errorf("invalid table name %q: longer than %d bytes", name, maxTableNameBytes)
	case !utf8.ValidString(name):
		return fmt.Errorf("invalid table name %q: not valid UTF-8", name)
	case strings.ContainsAny(name, forbiddenTableNameChars) || strings.ContainsFunc(name, unicode.IsControl):
		return fmt.Errorf("invalid table name %q: must not contain control characters or any of %s", name, forbiddenTableNameChars)
	}
	return nil
}

func validateUnquotedColumnName(name string) error {
	if !unquotedColumnNamePattern.MatchString(name) {
		return fmt.Errorf("invalid column name %q: must contain only letters, digits and underscores, and must not start with a digit", name)
	}
	return nil
}

func validateQuotedIdentifier(kind, value string) error {
	if value == "" || strings.ContainsAny(value, forbiddenQuotedIdentifierCh) || strings.ContainsFunc(value, unicode.IsControl) {
		return fmt.Errorf("invalid %s %q: must not be empty or contain backticks, backslashes or control characters", kind, value)
	}
	return nil
}

func (t TableConfig) validateIdentifiers() error {
	if err := validateTableName(t.SourceTableName); err != nil {
		return fmt.Errorf("source_table_name: %w", err)
	}
	if err := validateTableName(t.TargetTableName); err != nil {
		return fmt.Errorf("target_table_name: %w", err)
	}
	for _, column := range t.EmbedColumns {
		if err := validateUnquotedColumnName(column); err != nil {
			return fmt.Errorf("embed_columns: %w", err)
		}
	}
	for _, column := range t.MetadataColumns {
		if err := validateUnquotedColumnName(column); err != nil {
			return fmt.Errorf("metadata_columns: %w", err)
		}
	}
	return nil
}
