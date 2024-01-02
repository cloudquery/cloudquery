//go:build !windows

package specs

import _ "embed"

const noSuchFileErr = "open testdata/dir/no_such_file.yml: no such file or directory"

//go:embed testdata/application_default_credentials.json
var expectedApplicationDefaultCredentials string
