package cloudtrail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildCloudtrailChannelsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	channel := types.Channel{}
	if err := faker.FakeObject(&channel); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListChannels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.ListChannelsOutput{
			Channels: []types.Channel{channel},
		},
		nil,
	)

	channelOutput := &cloudtrail.GetChannelOutput{}
	if err := faker.FakeObject(&channelOutput); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetChannel(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		channelOutput,
		nil,
	)

	return services
}

func TestCloudtrailChannels(t *testing.T) {
	client.AwsMockTestHelper(t, Channels(), buildCloudtrailChannelsMock, client.TestOptions{})
}
