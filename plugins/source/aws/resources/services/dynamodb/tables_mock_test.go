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

func buildDynamodbTablesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var tableName string
	if err := faker.FakeObject(&tableName); err != nil {
		t.Fatal(err)
	}
	listOutput := &dynamodb.ListTablesOutput{
		TableNames: []string{tableName},
	}
	m.EXPECT().ListTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)

	descOutput := &dynamodb.DescribeTableOutput{
		Table: &types.TableDescription{
			TableName: &tableName,
		},
	}
	if err := faker.FakeObject(descOutput.Table); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTable(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		descOutput,
		nil,
	)

	repOutput := &dynamodb.DescribeTableReplicaAutoScalingOutput{
		TableAutoScalingDescription: &types.TableAutoScalingDescription{
			TableName:   &tableName,
			TableStatus: types.TableStatusActive,
		},
	}
	if err := faker.FakeObject(&repOutput.TableAutoScalingDescription.Replicas); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTableReplicaAutoScaling(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		repOutput,
		nil,
	)

	cbOutput := &dynamodb.DescribeContinuousBackupsOutput{
		ContinuousBackupsDescription: &types.ContinuousBackupsDescription{},
	}
	if err := faker.FakeObject(&cbOutput.ContinuousBackupsDescription); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeContinuousBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cbOutput,
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

func TestDynamodbTables(t *testing.T) {
	client.AwsMockTestHelper(t, Tables(), buildDynamodbTablesMock, client.TestOptions{})
}
