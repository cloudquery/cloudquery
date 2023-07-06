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

func buildDHCPOptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	item := types.DhcpOptions{}
	require.NoError(t, faker.FakeObject(&item))

	m.EXPECT().DescribeDhcpOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeDhcpOptionsOutput{
			DhcpOptions: []types.DhcpOptions{item},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestDHCPOptions(t *testing.T) {
	client.AwsMockTestHelper(t, DHCPOptions(), buildDHCPOptions, client.TestOptions{})
}
