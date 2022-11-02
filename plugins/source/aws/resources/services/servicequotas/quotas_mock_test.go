package servicequotas

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildQuotas(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockServiceQuotasClient(ctrl)
	services := servicequotas.ListServicesOutput{}
	err := faker.FakeObject(&services)
	if err != nil {
		t.Fatal(err)
	}
	services.NextToken = nil
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&services, nil)

	quotas := servicequotas.ListServiceQuotasOutput{}
	err = faker.FakeObject(&quotas)
	if err != nil {
		t.Fatal(err)
	}

	quotas.NextToken = nil
	m.EXPECT().ListServiceQuotas(gomock.Any(), gomock.Any(), gomock.Any()).Return(&quotas, nil).AnyTimes()
	return client.Services{
		ServiceQuotas: m,
	}
}

func TestQuotas(t *testing.T) {
	client.AwsMockTestHelper(t, Quotas(), buildQuotas, client.TestOptions{})
}
