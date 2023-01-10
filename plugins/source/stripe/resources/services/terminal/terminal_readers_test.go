package terminal_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/terminal"
)

func TestTerminalReaders(t *testing.T) {
	client.MockTestHelper(t, terminal.TerminalReaders(), client.TestOptions{})
}
