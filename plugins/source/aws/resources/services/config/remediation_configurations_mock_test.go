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

func buildRemediationConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)
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

func TestRemediationConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, RemediationConfigurations(), buildRemediationConfigurations, client.TestOptions{})
}
