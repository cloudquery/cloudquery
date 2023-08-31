package efs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEfsAccessPointsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEfsClient(ctrl)
	l := types.AccessPointDescription{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeAccessPoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&efs.DescribeAccessPointsOutput{
			AccessPoints: []types.AccessPointDescription{l},
		}, nil)

	return client.Services{
		Efs: m,
	}
}

func TestEfsAccesspoints(t *testing.T) {
	client.AwsMockTestHelper(t, AccessPoints(), buildEfsAccessPointsMock, client.TestOptions{})
}
