package services

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildWafs(t *testing.T, ctrl *gomock.Controller) client.Api {
	mock := mocks.NewMockApi(ctrl)

	var wafPackage cloudflare.WAFPackage
	if err := faker.FakeData(&wafPackage); err != nil {
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
	if err := faker.FakeData(&wafGroup); err != nil {
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
	skipFields := []string{"group"}
	if err := faker.FakeDataSkipFields(&wafRule, skipFields); err != nil {
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

	return mock
}

func TestWafs(t *testing.T) {
	client.CFMockTestHelper(t, Wafs(), buildWafs)
}
