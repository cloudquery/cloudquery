package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRegionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	r := types.Region{}
	require.NoError(t, faker.FakeObject(&r))

	r.OptInStatus = aws.String("opted-in")
	r.RegionName = aws.String("us-east-1")
	m.EXPECT().DescribeRegions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeRegionsOutput{
			Regions: []types.Region{r},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestRegions(t *testing.T) {
	client.AwsMockTestHelper(t, Regions(), buildRegionsMock, client.TestOptions{})
}
