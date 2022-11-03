package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildConfigurationSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	cs := sesv2.GetConfigurationSetOutput{}
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}

	ed := types.EventDestination{}
	if err := faker.FakeObject(&ed); err != nil {
		t.Fatal(err)
	}

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
