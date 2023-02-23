package analytics

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
)

func TestInstalls365d(t *testing.T) {
	client.MockTestHelper(t, Installs365Days(), buildInstalls, client.TestOptions{})
}
