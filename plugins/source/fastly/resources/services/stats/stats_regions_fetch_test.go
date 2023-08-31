package stats

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
)

func buildRegionsMock(t *testing.T, ctrl *gomock.Controller) services.FastlyClient {
	m := mocks.NewMockFastlyClient(ctrl)
	f := &fastly.RegionsResponse{}
	err := faker.FakeObject(&f)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRegions().Return(f, nil)
	return m
}

func TestStatsRegions(t *testing.T) {
	client.MockTestHelper(t, StatsRegions(), buildRegionsMock, client.TestOptions{})
}
