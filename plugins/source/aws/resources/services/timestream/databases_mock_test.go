package timestream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildTimestreamDatabasesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockTimestreamwriteClient(ctrl)
	database := types.Database{}
	require.NoError(t, faker.FakeObject(&database))

	m.EXPECT().ListDatabases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&timestreamwrite.ListDatabasesOutput{Databases: []types.Database{database}}, nil)

	table := types.Table{}
	require.NoError(t, faker.FakeObject(&table))

	m.EXPECT().ListTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&timestreamwrite.ListTablesOutput{Tables: []types.Table{table}}, nil)

	var tags []types.Tag
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&timestreamwrite.ListTagsForResourceOutput{Tags: tags}, nil)

	return client.Services{Timestreamwrite: m}
}

func TestTimestreamDatabases(t *testing.T) {
	client.AwsMockTestHelper(t, Databases(), buildTimestreamDatabasesMock, client.TestOptions{})
}
