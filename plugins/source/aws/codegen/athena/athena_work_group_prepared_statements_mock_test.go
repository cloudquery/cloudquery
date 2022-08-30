// Code generated by codegen; DO NOT EDIT.

package athena

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

func buildAthenaWorkGroupPreparedStatements(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAthenaClient(ctrl)

	var item types.PreparedStatementSummary
	if err := faker.FakeData(&item); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListPreparedStatements(
		gomock.Any(),
		&athena.ListPreparedStatementsInput{},
		gomock.Any(),
	).Return(
		&athena.ListPreparedStatementsOutput{
			PreparedStatements: []types.PreparedStatementSummary{item},
		},
		nil,
	)

	var detail types.PreparedStatement
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}

	detail.StatementName = item.StatementName

	mock.EXPECT().GetPreparedStatement(
		gomock.Any(),
		&athena.GetPreparedStatementInput{

			StatementName: item.StatementName,
		},
		gomock.Any(),
	).Return(
		&athena.GetPreparedStatementOutput{
			PreparedStatement: &detail,
		},
		nil,
	)

	return client.Services{
		Athena: mock,
	}
}

func TestAthenaWorkGroupPreparedStatements(t *testing.T) {
	client.AwsMockTestHelper(t, AthenaWorkGroupPreparedStatements(), buildAthenaWorkGroupPreparedStatements, client.TestOptions{})
}
