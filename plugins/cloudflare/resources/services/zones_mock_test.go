package services

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildZones(t *testing.T, ctrl *gomock.Controller) client.Api {
	mock := mocks.NewMockApi(ctrl)

	var zone cloudflare.Zone
	if err := faker.FakeData(&zone); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListZonesContext(
		gomock.Any(),
		gomock.Any(),
	).Return(
		[]cloudflare.Zone{zone},
		nil,
	)

	var record cloudflare.DNSRecord
	skipFields := []string{"Meta", "Data"}
	if err := faker.FakeDataSkipFields(&record, skipFields); err != nil {
		t.Fatal(err)
	}

	record.Meta = map[string]interface{}{
		"foo": "bar",
	}

	record.Data = map[string]interface{}{
		"foo": "bar",
	}

	mock.EXPECT().DNSRecords(
		gomock.Any(),
		zone.ID,
		gomock.Any(),
	).Return(
		[]cloudflare.DNSRecord{record},
		nil,
	)

	return mock
}

func TestZones(t *testing.T) {
	client.CFMockTestHelper(t, Zones(), buildZones)
}
