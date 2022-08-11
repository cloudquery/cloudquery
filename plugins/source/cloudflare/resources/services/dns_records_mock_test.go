package services

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDNSRecords(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

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
		client.TestZoneID,
		gomock.Any(),
	).Return(
		[]cloudflare.DNSRecord{record},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestDNSRecords(t *testing.T) {
	client.CFMockTestHelper(t, DNSRecords(), buildDNSRecords)
}
