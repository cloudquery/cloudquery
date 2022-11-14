package timestream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildTimestreamDatabasesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockTimestreamwriteClient(ctrl)
	database := types.Database{}
	err := faker.FakeObject(&database)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListDatabases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&timestreamwrite.ListDatabasesOutput{Databases: []types.Database{database}}, nil)

	table := types.Table{}
	err = faker.FakeObject(&table)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&timestreamwrite.ListTablesOutput{Tables: []types.Table{table}}, nil)

	var tags []types.Tag
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&timestreamwrite.ListTagsForResourceOutput{Tags: tags}, nil)

	return client.Services{Timestreamwrite: m}
}

func TestTimestreamDatabases(t *testing.T) {
	client.AwsMockTestHelper(t, Databases(), buildTimestreamDatabasesMock, client.TestOptions{})
}
