package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var parameterGroups docdb.DescribeCertificatesOutput
	if err := faker.FakeObject(&parameterGroups); err != nil {
		t.Fatal(err)
	}
	parameterGroups.Marker = nil
	m.EXPECT().DescribeCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameterGroups,
		nil,
	)
	return services
}

func TestCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, Certificates(), buildCertificatesMock, client.TestOptions{})
}
