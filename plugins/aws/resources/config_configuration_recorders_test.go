package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildConfigConfigurationRecorders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigServiceClient(ctrl)

	cr := types.ConfigurationRecorder{}
	err := faker.FakeData(&cr)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationRecorders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigurationRecordersOutput{
			ConfigurationRecorders: []types.ConfigurationRecorder{cr},
		}, nil)

	return client.Services{
		ConfigService: m,
	}
}

func TestConfigConfigurationRecorders(t *testing.T) {
	awsTestHelper(t, ConfigConfigurationRecorders(), buildConfigConfigurationRecorders, TestOptions{})
}
