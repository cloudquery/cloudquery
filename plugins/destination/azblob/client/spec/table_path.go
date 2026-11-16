package spec

import (
	"fmt"
	"strings"
)

func validateTableName(table string) error {
	if table == "" || table == "." || table == ".." || strings.ContainsAny(table, "/\\\x00") {
		return fmt.Errorf("table name %q cannot be used in a path", table)
	}
	return nil
}

func (s *Spec) BlobName(table string) (string, error) {
	if err := validateTableName(table); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s.%s%s", s.Path, table, s.Format, s.FileSpec.Compression.Extension()), nil
}
