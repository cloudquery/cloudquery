package domains

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/golang/mock/gomock"
)

func buildDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockDomainClient(ctrl)

	var dlistItem domain.ListResponse
	if err := faker.FakeObject(&dlistItem); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDomains().Return([]domain.ListResponse{dlistItem}, nil)

	var d domain.Details
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetDomain(dlistItem.FQDN).Return(d, nil)

	var w domain.WebRedirection
	if err := faker.FakeObject(&w); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWebRedirections(d.FQDN).Return([]domain.WebRedirection{w}, nil)

	var l domain.LiveDNS
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetLiveDNS(d.FQDN).Return(l, nil)

	var g domain.GlueRecord
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListGlueRecords(d.FQDN).Return([]domain.GlueRecord{g}, nil)

	var k domain.DNSSECKey
	if err := faker.FakeObject(&k); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDNSSECKeys(d.FQDN).Return([]domain.DNSSECKey{k}, nil)

	return client.Services{
		DomainClient: mock,
	}
}

func TestDomains(t *testing.T) {
	client.MockTestHelper(t, Domains(), buildDomains)
}
