package commands_tests

import (
	"os"
	"os/exec"
	"path"
	"runtime"
	"testing"
)

type integrationTest struct {
	dir string
}

func newIntegrationTest(t *testing.T) *integrationTest {
	dir := t.TempDir()
	return &integrationTest{dir: dir}
}

func (i *integrationTest) runCommand(t *testing.T, args ...string) (string, error) {
	t.Helper()

	command := exec.Command(getCLIFile(t), args...)
	command.Dir = i.dir

	output, err := command.CombinedOutput()

	return string(output), err
}

func getCLIFile(t *testing.T) string {
	t.Helper()
	_, filename, _, _ := runtime.Caller(1)
	binary := "cli"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	f, _ := os.Open(path.Join(path.Dir(filename), "..", binary))
	if f == nil {
		t.Fatal("cli file not found, please run go build in the cli directory before running tests")
	}
	return f.Name()
}
