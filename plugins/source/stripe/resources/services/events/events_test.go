package events_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/events"
)

func TestEvents(t *testing.T) {
	client.MockTestHelper(t, events.Events(), client.TestOptions{})
}
