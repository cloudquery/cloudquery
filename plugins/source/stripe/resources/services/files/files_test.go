package files_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/files"
)

func TestFiles(t *testing.T) {
	client.MockTestHelper(t, files.Files(), client.TestOptions{})
}
