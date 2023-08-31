package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDataRepoAssociationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var a types.DataRepositoryAssociation
	require.NoError(t, faker.FakeObject(&a))
	m.EXPECT().DescribeDataRepositoryAssociations(
		gomock.Any(),
		&fsx.DescribeDataRepositoryAssociationsInput{MaxResults: aws.Int32(25)},
		gomock.Any(),
	).Return(
		&fsx.DescribeDataRepositoryAssociationsOutput{Associations: []types.DataRepositoryAssociation{a}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestDataRepoAssociations(t *testing.T) {
	client.AwsMockTestHelper(t, DataRepositoryAssociations(), buildDataRepoAssociationsMock, client.TestOptions{})
}
