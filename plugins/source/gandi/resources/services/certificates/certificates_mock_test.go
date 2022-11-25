package certificates

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/go-gandi/go-gandi/certificate"
	"github.com/golang/mock/gomock"
)

func buildCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCertificateClient(ctrl)

	var cert certificate.CertificateType
	if err := faker.FakeObject(&cert); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListCertificates().Return([]certificate.CertificateType{cert}, nil)

	return client.Services{
		CertificateClient: mock,
	}
}

func TestCertificates(t *testing.T) {
	client.MockTestHelper(t, Certificates(), buildCertificates)
}
