// When executed, this file modifies the mocks in the client/mocks directory
// and adds assertions to all of them to check that the Region parameter is being set
// in all calls to the AWS SDK.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

var reOptions = regexp.MustCompile(`func \(m \*.+\) (?P<func>[^\(]+).+, (?P<arg>\S+) \.\.\.func\(\*(?P<name>[^.]+).Options\)`)

var tmpl = `
	// Assertion inserted by client/mockgen/main.go
	o := &%s.Options{}
	for _, f := range %s {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to %s")
	}
`

func AddAssertions(mockFile string) error {
	f, err := os.Open(mockFile)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("")
	}
	content := string(b)

	if strings.Contains(content, "// Assertion inserted by") {
		log.Println(mockFile, "already converted, skipping.")
		return nil
	}

	lines := strings.Split(content, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		matches := reOptions.FindAllStringSubmatchIndex(line, 1)
		if len(matches) == 0 {
			continue
		}
		m := matches[0]
		funcName := line[m[2]:m[3]]
		argName := line[m[4]:m[5]]
		packageName := line[m[6]:m[7]]
		// fmt.Println(funcName, argName, packageName)
		v := fmt.Sprintf(tmpl, packageName, argName, funcName)
		lines = append(lines[:i+1], append([]string{v}, lines[i+1:]...)...)
	}

	newContent := strings.Join(lines, "\n")
	s, _ := f.Stat()
	return os.WriteFile(mockFile, []byte(newContent), s.Mode().Perm())
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	mocksDir := path.Join(currentDir, "../mocks")
	entries, err := os.ReadDir(mocksDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		err := AddAssertions(path.Join(mocksDir, e.Name()))
		if err != nil {
			log.Fatalf("failed to modify mock file: %v", err)
		}
	}
}
