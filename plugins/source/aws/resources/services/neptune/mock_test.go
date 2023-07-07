package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNeptuneDBClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNeptuneClient(ctrl)
	l := types.DBCluster{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&neptune.DescribeDBClustersOutput{
			DBClusters: []types.DBCluster{l},
		}, nil)
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: l.DBClusterArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{
		Neptune: m,
	}
}

func buildNeptuneDBInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNeptuneClient(ctrl)
	l := types.DBInstance{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&neptune.DescribeDBInstancesOutput{
			DBInstances: []types.DBInstance{l},
		}, nil)
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: l.DBInstanceArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	return client.Services{
		Neptune: m,
	}
}

func buildNeptuneDBSubnetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNeptuneClient(ctrl)
	l := types.DBSubnetGroup{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeDBSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&neptune.DescribeDBSubnetGroupsOutput{
			DBSubnetGroups: []types.DBSubnetGroup{l},
		}, nil)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: l.DBSubnetGroupArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{
		Neptune: m,
	}
}

func TestNeptuneInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildNeptuneDBInstances, client.TestOptions{})
}
func TestNeptuneClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildNeptuneDBClusters, client.TestOptions{})
}
func TestNeptuneSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubnetGroups(), buildNeptuneDBSubnetGroups, client.TestOptions{})
}
