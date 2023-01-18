package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

func parseDocsTables() map[string]bool {
	tablesMap := make(map[string]bool)
	tablesReadmes, err := doublestar.Glob(os.DirFS("../../"), "plugins/source/**/docs/tables/README.md", doublestar.WithFailOnPatternNotExist(), doublestar.WithFilesOnly())
	if err != nil {
		panic(err)
	}

	for _, readme := range tablesReadmes {
		content, err := ioutil.ReadFile("../../" + readme)
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			pos1 := strings.Index(line, "[")
			if pos1 == -1 {
				continue
			}
			pos2 := strings.Index(line, "]")
			table := line[pos1+1 : pos2]
			tablesMap[table] = true
		}
	}

	return tablesMap
}

func parseCodeTables() map[string]string {
	tablesMap := make(map[string]string)
	tableFiles, err := doublestar.Glob(os.DirFS("../../"), "plugins/source/**/resources/services/**/*.go", doublestar.WithFailOnPatternNotExist(), doublestar.WithFilesOnly())
	if err != nil {
		panic(err)
	}
	for _, tableFile := range tableFiles {
		if strings.HasSuffix(tableFile, "_test.go") {
			continue
		}
		content, err := ioutil.ReadFile("../../" + tableFile)
		if err != nil {
			panic(err)
		}
		contentString := string(content)
		if !strings.Contains(contentString, "*schema.Table") {
			continue
		}

		tableNameRegex := regexp.MustCompile(`schema.Table[\s\S]+?Name\:.*?"(.*?)",`)
		tableName := tableNameRegex.FindStringSubmatch(contentString)
		tablesMap[tableName[1]] = tableFile

	}
	return tablesMap
}

func main() {
	tablesFromReadmes := parseDocsTables()
	tablesFromCode := parseCodeTables()

	for table, file := range tablesFromCode {
		if _, ok := tablesFromReadmes[table]; !ok {
			fmt.Printf("- Table `%s` is declared in code but missing from README. Table declaration file: `%s`\n", table, file)
		}
	}

	for table := range tablesFromReadmes {
		if _, ok := tablesFromCode[table]; !ok {
			fmt.Printf("- Table `%s` is in README but not declared in code\n", table)
		}
	}
}
