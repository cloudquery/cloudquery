package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2LaunchTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	lt := types.LaunchTemplate{}
	ltv := types.LaunchTemplateVersion{}
	require.NoError(t, faker.FakeObject(&lt))

	require.NoError(t, faker.FakeObject(&ltv))

	m.EXPECT().DescribeLaunchTemplates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeLaunchTemplatesOutput{
			LaunchTemplates: []types.LaunchTemplate{lt},
		}, nil)

	m.EXPECT().DescribeLaunchTemplateVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeLaunchTemplateVersionsOutput{
			LaunchTemplateVersions: []types.LaunchTemplateVersion{ltv},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2LaunchTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, LaunchTemplates(), buildEc2LaunchTemplates, client.TestOptions{})
}
