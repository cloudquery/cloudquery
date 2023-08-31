package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRemediationConfigurations(t *testing.T, m *mocks.MockConfigserviceClient) client.Services {
	l := types.RemediationConfiguration{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeRemediationConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeRemediationConfigurationsOutput{
			RemediationConfigurations: []types.RemediationConfiguration{l},
		}, nil)
	return client.Services{
		Configservice: m,
	}
}
