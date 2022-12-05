package partitions

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

// Generate generates a partitions.go file
func Generate() error {
	fmt.Println("Generate partitions")

	parts, err := readPartitions()
	if err != nil {
		return err
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	//csr := caser.New()
	// write services.go file
	servicesTpl, err := template.New("partitions.go.tpl").
		Funcs(template.FuncMap{
			"ToCamel": strcase.ToCamel,
		}).
		ParseFS(templatesFS, "templates/*.go.tpl")
	if err != nil {
		return err
	}

	buff := new(bytes.Buffer)
	if err := servicesTpl.Execute(buff, parts); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	filePath := path.Join(path.Dir(filename), "../../client/partitions.go")
	return formatAndWriteFile(filePath, buff.Bytes())
}

func formatAndWriteFile(filePath string, content []byte) error {
	formattedContent, err := format.Source(content)
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
