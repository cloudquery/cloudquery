package athena

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWorkGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAthenaClient(ctrl)

	workGroupsOutput := athena.ListWorkGroupsOutput{}
	require.NoError(t, faker.FakeObject(&workGroupsOutput))

	workGroupsOutput.NextToken = nil
	m.EXPECT().ListWorkGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&workGroupsOutput, nil)

	workGroup := athena.GetWorkGroupOutput{}
	require.NoError(t, faker.FakeObject(&workGroup))

	m.EXPECT().GetWorkGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&workGroup, nil)

	namedQueriesOutput := athena.ListNamedQueriesOutput{}
	require.NoError(t, faker.FakeObject(&namedQueriesOutput))

	namedQueriesOutput.NextToken = nil
	m.EXPECT().ListNamedQueries(gomock.Any(), gomock.Any(), gomock.Any()).Return(&namedQueriesOutput, nil)

	queryExecutionsOutput := athena.ListQueryExecutionsOutput{}
	require.NoError(t, faker.FakeObject(&queryExecutionsOutput))

	queryExecutionsOutput.NextToken = nil
	m.EXPECT().ListQueryExecutions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryExecutionsOutput, nil)

	preparedStatementsOutput := athena.ListPreparedStatementsOutput{}
	require.NoError(t, faker.FakeObject(&preparedStatementsOutput))

	preparedStatementsOutput.NextToken = nil
	m.EXPECT().ListPreparedStatements(gomock.Any(), gomock.Any(), gomock.Any()).Return(&preparedStatementsOutput, nil)

	tags := athena.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	preparedStatement := athena.GetPreparedStatementOutput{}
	require.NoError(t, faker.FakeObject(&preparedStatement))

	m.EXPECT().GetPreparedStatement(gomock.Any(), gomock.Any(), gomock.Any()).Return(&preparedStatement, nil)

	namedQuery := athena.GetNamedQueryOutput{}
	require.NoError(t, faker.FakeObject(&namedQuery))

	m.EXPECT().GetNamedQuery(gomock.Any(), gomock.Any(), gomock.Any()).Return(&namedQuery, nil)
	queryExecution := athena.GetQueryExecutionOutput{}
	require.NoError(t, faker.FakeObject(&queryExecution))

	m.EXPECT().GetQueryExecution(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryExecution, nil)

	return client.Services{
		Athena: m,
	}
}

func TestWorkGroups(t *testing.T) {
	client.AwsMockTestHelper(t, WorkGroups(), buildWorkGroups, client.TestOptions{})
}
