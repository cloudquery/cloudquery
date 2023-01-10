package radar_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/radar"
)

func TestRadarEarlyFraudWarnings(t *testing.T) {
	client.MockTestHelper(t, radar.RadarEarlyFraudWarnings(), client.TestOptions{})
}
