package zones

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildZones(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var zonesResp cloudflare.ZonesResponse
	if err := faker.FakeObject(&zonesResp); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListZonesContext(
		gomock.Any(),
		gomock.Any(),
	).Return(
		zonesResp,
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestZones(t *testing.T) {
	client.MockTestHelper(t, Zones(), buildZones)
}
