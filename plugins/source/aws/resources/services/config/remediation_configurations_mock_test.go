package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildRemediationConfigurations(t *testing.T, m *mocks.MockConfigserviceClient) client.Services {
	l := types.RemediationConfiguration{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeRemediationConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeRemediationConfigurationsOutput{
			RemediationConfigurations: []types.RemediationConfiguration{l},
		}, nil)
	return client.Services{
		Configservice: m,
	}
}
