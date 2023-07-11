package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildConfigurationSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	cs := sesv2.GetConfigurationSetOutput{}
	require.NoError(t, faker.FakeObject(&cs))

	ed := types.EventDestination{}
	require.NoError(t, faker.FakeObject(&ed))

	sesClient.EXPECT().ListConfigurationSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sesv2.ListConfigurationSetsOutput{ConfigurationSets: []string{*cs.ConfigurationSetName}},
		nil,
	)
	sesClient.EXPECT().GetConfigurationSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cs,
		nil,
	)

	sesClient.EXPECT().GetConfigurationSetEventDestinations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sesv2.GetConfigurationSetEventDestinationsOutput{EventDestinations: []types.EventDestination{ed}},
		nil,
	)

	return client.Services{
		Sesv2: sesClient,
	}
}

func TestConfigurationSets(t *testing.T) {
	client.AwsMockTestHelper(t, ConfigurationSets(), buildConfigurationSets, client.TestOptions{})
}
