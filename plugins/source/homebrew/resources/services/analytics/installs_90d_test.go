package analytics

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
)

func TestInstalls90d(t *testing.T) {
	client.MockTestHelper(t, Installs90Days(), buildInstalls, client.TestOptions{})
}
