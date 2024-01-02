//go:build windows

package specs

import _ "embed"

const noSuchFileErr = "open testdata/dir/no_such_file.yml: The system cannot find the file specified."

//go:embed testdata/application_default_credentials.json
var expectedApplicationDefaultCredentials string

func init() {
	expectedApplicationDefaultCredentials = strings.ReplaceAll(expectedApplicationDefaultCredentials, "\n", "\r\n")
}
