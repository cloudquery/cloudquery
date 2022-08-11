package athena

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWorkGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAthenaClient(ctrl)

	workGroupsOutput := athena.ListWorkGroupsOutput{}
	err := faker.FakeData(&workGroupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	workGroupsOutput.NextToken = nil
	m.EXPECT().ListWorkGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&workGroupsOutput, nil)

	workGroup := athena.GetWorkGroupOutput{}
	err = faker.FakeData(&workGroup)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetWorkGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&workGroup, nil)

	namedQueriesOutput := athena.ListNamedQueriesOutput{}
	err = faker.FakeData(&namedQueriesOutput)
	if err != nil {
		t.Fatal(err)
	}
	namedQueriesOutput.NextToken = nil
	m.EXPECT().ListNamedQueries(gomock.Any(), gomock.Any(), gomock.Any()).Return(&namedQueriesOutput, nil)

	queryExecutionsOutput := athena.ListQueryExecutionsOutput{}
	err = faker.FakeData(&queryExecutionsOutput)
	if err != nil {
		t.Fatal(err)
	}
	queryExecutionsOutput.NextToken = nil
	m.EXPECT().ListQueryExecutions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryExecutionsOutput, nil)

	preparedStatementsOutput := athena.ListPreparedStatementsOutput{}
	err = faker.FakeData(&preparedStatementsOutput)
	if err != nil {
		t.Fatal(err)
	}
	preparedStatementsOutput.NextToken = nil
	m.EXPECT().ListPreparedStatements(gomock.Any(), gomock.Any(), gomock.Any()).Return(&preparedStatementsOutput, nil)

	tags := athena.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	preparedStatement := athena.GetPreparedStatementOutput{}
	err = faker.FakeData(&preparedStatement)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetPreparedStatement(gomock.Any(), gomock.Any(), gomock.Any()).Return(&preparedStatement, nil)

	namedQuery := athena.GetNamedQueryOutput{}
	err = faker.FakeData(&namedQuery)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetNamedQuery(gomock.Any(), gomock.Any(), gomock.Any()).Return(&namedQuery, nil)
	queryExecution := athena.GetQueryExecutionOutput{}
	err = faker.FakeData(&queryExecution)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetQueryExecution(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryExecution, nil)

	return client.Services{
		Athena: m,
	}
}

func TestWorkGroups(t *testing.T) {
	client.AwsMockTestHelper(t, WorkGroups(), buildWorkGroups, client.TestOptions{})
}
