package resources

import (
	"bytes"
	"fmt"
	"path"
)

func (r *Resource) generateSchema(dir string) error {
	tpl, err := parse("resource")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, r.SubService+".go")
	return write(filePath, buff.Bytes())
}
