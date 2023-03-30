package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDynamodbGlobalTablesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var globalTable types.GlobalTable
	if err := faker.FakeObject(&globalTable); err != nil {
		t.Fatal(err)
	}
	listOutput := &dynamodb.ListGlobalTablesOutput{
		GlobalTables: []types.GlobalTable{globalTable},
	}
	m.EXPECT().ListGlobalTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)
	var gtd types.GlobalTableDescription
	if err := faker.FakeObject(&gtd); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeGlobalTable(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodb.DescribeGlobalTableOutput{
			GlobalTableDescription: &gtd,
		},
		nil,
	)

	tags := &dynamodb.ListTagsOfResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsOfResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestDynamodbGlobalTables(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalTables(), buildDynamodbGlobalTablesMock, client.TestOptions{})
}
