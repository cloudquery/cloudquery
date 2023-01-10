package certificates

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/go-gandi/go-gandi/certificate"
	"github.com/golang/mock/gomock"
)

func buildPackages(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCertificateClient(ctrl)

	var p certificate.Package
	if err := faker.FakeObject(&p); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListPackages().Return([]certificate.Package{p}, nil)

	return client.Services{
		CertificateClient: mock,
	}
}

func TestPackages(t *testing.T) {
	client.MockTestHelper(t, CertificatePackages(), buildPackages)
}
