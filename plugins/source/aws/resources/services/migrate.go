package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var directory string

func migrate(dir string) {
	items, _ := os.ReadDir(dir)
	for _, it := range items {
		if strings.HasSuffix(it.Name(), "_fetch.go") || strings.HasSuffix(it.Name(), ".bk") {
			fmt.Fprintln(os.Stderr, "Skipping "+it.Name())
			continue
		}
		if strings.HasSuffix(it.Name(), ".go") && !strings.HasSuffix(it.Name(), "_test.go") {
			migrateFile(path.Join(dir, it.Name()))
		}
	}
}

func migrateFile(f string) {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	s := string(data)
	stopAt := strings.Index(s, "// ==========")
	if stopAt == -1 {
		fmt.Fprintln(os.Stderr, "Already migrated?")
		return
	}

	importEnd := strings.Index(s, ")")

	d := path.Dir(f)
	b := path.Base(f)
	old := path.Join(d, b+".bk")
	fetch := path.Join(d, b[:len(b)-3]+"_fetch.go")

	// write backup file
	err = os.WriteFile(old, []byte(s[:stopAt]), 0755)
	if err != nil {
		panic(err)
	}

	// write resolvers
	fetchContent := s[stopAt:]
	lines := strings.Split(fetchContent, "\n")
	firstLine := 0 // skip comment lines
	for _, line := range lines {
		if strings.HasPrefix(line, "//") {
			firstLine++
			continue
		}
		break
	}
	lines = lines[firstLine:]
	fetchContent = strings.Join(lines, "\n")
	fetchContent = s[:importEnd+1] + "\n" + fetchContent
	err = os.WriteFile(fetch, []byte(fetchContent), 0755)
	if err != nil {
		panic(err)
	}

	// delete original file (this will need to be regenerated)
	os.RemoveAll(f)
}

func main() {
	flag.StringVar(&directory, "d", "", "directory to migrate")
	flag.Parse()

	if directory == "" {
		fmt.Fprintln(os.Stderr, "Usage: -d <directory_to_migrate>")
		return
	}
	migrate(directory)
}
