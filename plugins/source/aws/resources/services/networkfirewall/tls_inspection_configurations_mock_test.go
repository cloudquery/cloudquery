package networkfirewall

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildTLSInspectionConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	ticm := types.TLSInspectionConfigurationMetadata{}
	require.NoError(t, faker.FakeObject(&ticm))

	m.EXPECT().ListTLSInspectionConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListTLSInspectionConfigurationsOutput{
			TLSInspectionConfigurations: []types.TLSInspectionConfigurationMetadata{ticm},
		}, nil)

	tico := networkfirewall.DescribeTLSInspectionConfigurationOutput{}
	require.NoError(t, faker.FakeObject(&tico))

	m.EXPECT().DescribeTLSInspectionConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tico, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestTLSInspectionConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, TLSInspectionConfigurations(), buildTLSInspectionConfigurationsMock, client.TestOptions{})
}
