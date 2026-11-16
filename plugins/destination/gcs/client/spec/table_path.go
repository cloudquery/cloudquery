package spec

import (
	"fmt"
	"path"
	"strings"
)

func validateTableName(table string) error {
	if table == "" || table == "." || table == ".." || strings.ContainsAny(table, "/\\\x00") {
		return fmt.Errorf("table name %q cannot be used in a path", table)
	}
	return nil
}

func staticPathPrefix(template string) string {
	prefix, _, _ := strings.Cut(template, "{{")
	cleaned := path.Clean(prefix)
	if prefix == "" || cleaned == "." {
		return ""
	}
	if strings.HasSuffix(prefix, "/") && cleaned != "/" {
		return cleaned + "/"
	}
	return cleaned
}

func validatePathPrefix(name string, template string) error {
	prefix := staticPathPrefix(template)
	escapes := !strings.HasPrefix(name, prefix)
	if prefix == "" {
		escapes = name == ".." || strings.HasPrefix(name, "../")
	}
	if escapes {
		return fmt.Errorf("path %q is outside of the configured path %q", name, template)
	}
	return nil
}
