package servicequotas

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockServicequotasClient(ctrl)

	services := servicequotas.ListServicesOutput{}
	require.NoError(t, faker.FakeObject(&services))
	services.NextToken = nil
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&services, nil)

	quotas := servicequotas.ListServiceQuotasOutput{}
	require.NoError(t, faker.FakeObject(&quotas))

	quotas.NextToken = nil
	m.EXPECT().ListServiceQuotas(gomock.Any(), gomock.Any(), gomock.Any()).Return(&quotas, nil).AnyTimes()

	return client.Services{
		Servicequotas: m,
	}
}

func TestQuotas(t *testing.T) {
	client.AwsMockTestHelper(t, Services(), buildServices, client.TestOptions{})
}
