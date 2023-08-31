package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRdsDbProxiesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	proxy := types.DBProxy{}
	require.NoError(t, faker.FakeObject(&proxy))

	tags := rds.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().DescribeDBProxies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBProxiesOutput{
			DBProxies: []types.DBProxy{proxy},
		}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Rds: m,
	}
}

func TestRdsDbProxues(t *testing.T) {
	client.AwsMockTestHelper(t, DbProxies(), buildRdsDbProxiesMock, client.TestOptions{})
}
