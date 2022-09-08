package services

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWafOverrides(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var wafOverride cloudflare.WAFOverride
	if err := faker.FakeData(&wafOverride); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWAFOverrides(
		gomock.Any(),
		client.TestZoneID,
	).Return(
		[]cloudflare.WAFOverride{wafOverride},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestWafOverrides(t *testing.T) {
	client.MockTestHelper(t, WafOverrides(), buildWafOverrides)
}
