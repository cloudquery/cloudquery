package emr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSecurityConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary types.SecurityConfigurationSummary
	require.NoError(t, faker.FakeObject(&summary))

	mock.EXPECT().ListSecurityConfigurations(gomock.Any(), &emr.ListSecurityConfigurationsInput{}, gomock.Any()).Return(
		&emr.ListSecurityConfigurationsOutput{SecurityConfigurations: []types.SecurityConfigurationSummary{summary}},
		nil,
	)

	configString := "{}"
	mock.EXPECT().DescribeSecurityConfiguration(gomock.Any(), &emr.DescribeSecurityConfigurationInput{Name: summary.Name}, gomock.Any()).Return(
		&emr.DescribeSecurityConfigurationOutput{CreationDateTime: summary.CreationDateTime, Name: summary.Name, SecurityConfiguration: &configString},
		nil,
	)
	return client.Services{Emr: mock}
}

func TestSecurityConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, SecurityConfigurations(), buildSecurityConfigurations, client.TestOptions{})
}
