package stats

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
)

func buildServiceStatsMock(t *testing.T, ctrl *gomock.Controller) services.FastlyClient {
	m := mocks.NewMockFastlyClient(ctrl)
	m.EXPECT().GetStatsJSON(gomock.Any(), gomock.Any()).DoAndReturn(func(a, resp any) error {
		b, err := os.ReadFile("testdata/stats.json")
		if err != nil {
			return err
		}
		return json.Unmarshal(b, resp)
	})
	return m
}

func TestStatsServices(t *testing.T) {
	createdAt := time.Now().Add(-time.Hour * 24 * 20)
	client.MockTestHelper(t, StatsServices(), buildServiceStatsMock, client.TestOptions{
		Service: &fastly.Service{
			CreatedAt: &createdAt,
		},
		Region: "test-region",
	})
}
