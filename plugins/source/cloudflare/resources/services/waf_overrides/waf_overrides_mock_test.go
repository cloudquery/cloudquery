package waf_overrides

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildWAFOverrides(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var wafOverride cloudflare.WAFOverride
	if err := faker.FakeObject(&wafOverride); err != nil {
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
	client.MockTestHelper(t, WAFOverrides(), buildWAFOverrides)
}
