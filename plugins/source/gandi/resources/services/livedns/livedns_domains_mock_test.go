package livedns

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/golang/mock/gomock"
)

func buildDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLiveDNSClient(ctrl)

	var d livedns.Domain
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListDomains().Return([]livedns.Domain{d}, nil)

	var s livedns.Snapshot
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListSnapshots(d.FQDN).Return([]livedns.Snapshot{s}, nil)

	return client.Services{
		LiveDNSClient: mock,
	}
}

func TestDomains(t *testing.T) {
	client.MockTestHelper(t, LiveDNSDomains(), buildDomains)
}
