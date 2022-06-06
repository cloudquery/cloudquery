package test_integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	setupCommands [][]string
	actCommand    []string
}

const (
	dbUser     = "postgres"
	dbPort     = "5432"
	dbName     = "postgres"
	dbHost     = "localhost"
	dbPassword = "pass"
)

var mainFile = getMainFile()

var testCases = map[string]testCase{
	"policy-output": {
		setupCommands: [][]string{
			{"dropdb", "-h", dbHost, "-p", dbPort, "-U", dbUser, dbName},
			{"createdb", "-h", dbHost, "-p", dbPort, "-U", dbUser, dbName},
			// Created using `pg_dump -h 127.0.0.1 -U postgres postgres > test_integration/fixtures/fetch_data/aws_s3.pgsql` after fetching only `s3*` resources, and sanitizing the output
			{"psql", fmt.Sprintf("host=%s port=%s dbname=%s user=%s", dbHost, dbPort, dbName, dbUser), "-f", getFixtureFilePath("fetch_data/aws_s3.pgsql")},
			{"go", "run", mainFile, "init", "aws"},
		},
		actCommand: []string{"go", "run", mainFile, "policy", "run", "aws"},
	},
}

func TestIntegrationCommands(t *testing.T) {
	for testName, args := range testCases {
		t.Run(testName, func(t *testing.T) {
			cwd, err := getCwd()
			require.NoError(t, err)
			defer os.RemoveAll(cwd)
			for _, setupCommand := range args.setupCommands {
				_, err := runCommand(cwd, setupCommand...)
				if err != "" {
					fmt.Println(err)
				}
			}
			out, _ := runCommand(cwd, args.actCommand...)
			cupaloy.SnapshotT(t, out)
		})
	}
}

func getTestFileDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func getMainFile() string {
	return path.Join(getTestFileDir(), "..", "main.go")
}

func getFixuresDir() string {
	return path.Join(getTestFileDir(), "fixtures")
}

func getFixtureFilePath(fixtureName string) string {
	return path.Join(getFixuresDir(), fixtureName)
}

func getCwd() (string, error) {
	tempDir := path.Join(getTestFileDir(), "..", "tmp")
	os.Mkdir(tempDir, 0755)
	dir, err := ioutil.TempDir(tempDir, "test_integration")
	return dir, err
}

func runCommand(cwd string, args ...string) (out string, err string) {
	command := exec.Command(args[0], args[1:]...)
	command.Env = append(os.Environ(), "PGPASSWORD="+dbPassword)
	command.Dir = cwd
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	command.Run()
	return stdout.String(), stderr.String()
}
