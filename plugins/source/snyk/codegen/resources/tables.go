package resources

import (
	"bytes"
	"fmt"
	"path"
	"runtime"

	"golang.org/x/exp/slices"
)

func tables(resources []*Resource) error {
	tpl, err := parse("tables")
	if err != nil {
		return err
	}

	sorted := make([]*Resource, len(resources))
	copy(sorted, resources)
	slices.SortFunc(sorted, func(a, b *Resource) bool { return a.Name < b.Name })

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, sorted); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	filePath := path.Join(path.Dir(filename), "../../resources/plugin/tables.go")

	return write(filePath, buff.Bytes())
}
