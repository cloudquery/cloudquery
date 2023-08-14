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

func buildConfigDeliveryChannels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)

	var dc types.DeliveryChannel
	require.NoError(t, faker.FakeObject(&dc))

	var dcs types.DeliveryChannelStatus
	require.NoError(t, faker.FakeObject(&dcs))

	m.EXPECT().DescribeDeliveryChannels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeDeliveryChannelsOutput{
			DeliveryChannels: []types.DeliveryChannel{dc},
		}, nil)
	m.EXPECT().DescribeDeliveryChannelStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeDeliveryChannelStatusOutput{
			DeliveryChannelsStatus: []types.DeliveryChannelStatus{dcs},
		}, nil)

	return client.Services{
		Configservice: m,
	}
}

func TestConfigDeliveryChannels(t *testing.T) {
	client.AwsMockTestHelper(t, DeliveryChannels(), buildConfigDeliveryChannels, client.TestOptions{})
}
