package transfer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildServersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockTransferClient(ctrl)

	var ls types.ListedServer
	require.NoError(t, faker.FakeObject(&ls))
	m.EXPECT().ListServers(
		gomock.Any(),
		&transfer.ListServersInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&transfer.ListServersOutput{Servers: []types.ListedServer{ls}},
		nil,
	)

	var ds types.DescribedServer
	require.NoError(t, faker.FakeObject(&ds))
	ds.ServerId = ls.ServerId
	ds.Arn = ls.Arn
	m.EXPECT().DescribeServer(
		gomock.Any(),
		&transfer.DescribeServerInput{ServerId: ls.ServerId},
		gomock.Any(),
	).Return(
		&transfer.DescribeServerOutput{Server: &ds},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&transfer.ListTagsForResourceInput{Arn: ds.Arn},
		gomock.Any(),
	).Return(
		&transfer.ListTagsForResourceOutput{Tags: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}}},
		nil,
	)

	return client.Services{
		Transfer: m,
	}
}

func TestServers(t *testing.T) {
	client.AwsMockTestHelper(t, Servers(), buildServersMock, client.TestOptions{})
}
