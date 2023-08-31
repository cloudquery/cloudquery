package waf_packages

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildWAFPackages(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var wafPackage cloudflare.WAFPackage
	if err := faker.FakeObject(&wafPackage); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWAFPackages(
		gomock.Any(),
		client.TestZoneID,
	).Return(
		[]cloudflare.WAFPackage{wafPackage},
		nil,
	)

	var wafGroup cloudflare.WAFGroup
	if err := faker.FakeObject(&wafGroup); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWAFGroups(
		gomock.Any(),
		client.TestZoneID,
		wafPackage.ID,
	).Return(
		[]cloudflare.WAFGroup{wafGroup},
		nil,
	)

	var wafRule cloudflare.WAFRule
	if err := faker.FakeObject(&wafRule); err != nil {
		t.Fatal(err)
	}

	wafRule.Group = struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		"fake-group-id",
		"fake-group-name",
	}

	mock.EXPECT().ListWAFRules(
		gomock.Any(),
		client.TestZoneID,
		wafPackage.ID,
	).Return(
		[]cloudflare.WAFRule{wafRule},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestWAFPackages(t *testing.T) {
	client.MockTestHelper(t, WAFPackages(), buildWAFPackages)
}
