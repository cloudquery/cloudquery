//go:build mock
// +build mock

package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDynamodbTablesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamoDBClient(ctrl)
	services := client.Services{
		DynamoDB: m,
	}
	var tableName string
	if err := faker.FakeData(&tableName); err != nil {
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
	if err := faker.FakeData(descOutput.Table); err != nil {
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
	if err := faker.FakeData(&repOutput.TableAutoScalingDescription.Replicas); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTableReplicaAutoScaling(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		repOutput,
		nil,
	)

	cbOutput := &dynamodb.DescribeContinuousBackupsOutput{
		ContinuousBackupsDescription: &types.ContinuousBackupsDescription{},
	}
	if err := faker.FakeData(&cbOutput.ContinuousBackupsDescription); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeContinuousBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cbOutput,
		nil,
	)

	tags := &dynamodb.ListTagsOfResourceOutput{}
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsOfResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestDynamodbTables(t *testing.T) {
	client.AwsMockTestHelper(t, DynamodbTables(), buildDynamodbTablesMock, client.TestOptions{})
}
