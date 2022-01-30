package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
)

func buildConfigConfigurationRecorders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigServiceClient(ctrl)
	l := types.ConfigurationRecorder{}
	if err := faker.FakeData(&l); err != nil {
		t.Fatal(err)
	}
	sl := types.ConfigurationRecorderStatus{}
	if err := faker.FakeData(&sl); err != nil {
		t.Fatal(err)
	}
	sl.Name = l.Name
	m.EXPECT().DescribeConfigurationRecorderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigurationRecorderStatusOutput{
			ConfigurationRecordersStatus: []types.ConfigurationRecorderStatus{sl},
		}, nil)
	m.EXPECT().DescribeConfigurationRecorders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigurationRecordersOutput{
			ConfigurationRecorders: []types.ConfigurationRecorder{l},
		}, nil)
	return client.Services{
		ConfigService: m,
	}
}

func TestConfigConfigurationRecorders(t *testing.T) {
	client.AwsMockTestHelper(t, ConfigConfigurationRecorders(), buildConfigConfigurationRecorders, client.TestOptions{})
}
