package dns_records

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildDNSRecords(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var record cloudflare.DNSRecord
	if err := faker.FakeObject(&record); err != nil {
		t.Fatal(err)
	}

	record.Meta = map[string]any{
		"foo": "bar",
	}

	record.Data = map[string]any{
		"foo": "bar",
	}

	mock.EXPECT().ListDNSRecords(
		gomock.Any(),
		cloudflare.ZoneIdentifier(client.TestZoneID),
		gomock.Any(),
	).Return(
		[]cloudflare.DNSRecord{record},
		&cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestDNSRecords(t *testing.T) {
	client.MockTestHelper(t, DNSRecords(), buildDNSRecords)
}
