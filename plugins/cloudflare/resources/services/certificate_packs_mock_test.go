package services

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCertificatePacks(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var certPack cloudflare.CertificatePack
	if err := faker.FakeData(&certPack); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListCertificatePacks(
		gomock.Any(),
		client.TestZoneID,
	).Return(
		[]cloudflare.CertificatePack{certPack},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestCertificatePacks(t *testing.T) {
	client.CFMockTestHelper(t, CertificatePacks(), buildCertificatePacks)
}
