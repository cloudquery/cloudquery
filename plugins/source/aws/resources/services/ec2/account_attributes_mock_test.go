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

func buildAccountAttributesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	ac := types.AccountAttribute{}
	require.NoError(t, faker.FakeObject(&ac))

	m.EXPECT().DescribeAccountAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeAccountAttributesOutput{
			AccountAttributes: []types.AccountAttribute{ac},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestAccountAttributes(t *testing.T) {
	client.AwsMockTestHelper(t, AccountAttributes(), buildAccountAttributesMock, client.TestOptions{})
}
