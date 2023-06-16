package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildConfigDeliveryChannels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)

	var dc types.DeliveryChannel
	if err := faker.FakeObject(&dc); err != nil {
		t.Fatal(err)
	}
	var dcs types.DeliveryChannelStatus
	if err := faker.FakeObject(&dcs); err != nil {
		t.Fatal(err)
	}

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
