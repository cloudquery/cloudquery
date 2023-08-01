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

func buildIpams(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	i := types.Ipam{}
	require.NoError(t, faker.FakeObject(&i))

	m.EXPECT().DescribeIpams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeIpamsOutput{
			Ipams: []types.Ipam{i},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestIpams(t *testing.T) {
	client.AwsMockTestHelper(t, IPAMs(), buildIpams, client.TestOptions{})
}
