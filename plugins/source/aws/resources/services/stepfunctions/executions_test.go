package stepfunctions

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildExecutions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSfnClient(ctrl)
	eli := types.ExecutionListItem{}
	err := faker.FakeObject(&eli)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListExecutions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListExecutionsOutput{
			Executions: []types.ExecutionListItem{eli},
		}, nil)

	execOut := sfn.DescribeExecutionOutput{}
	if err := faker.FakeObject(&execOut); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeExecution(gomock.Any(), gomock.Any(), gomock.Any()).Return(&execOut, nil)

	mrli := types.MapRunListItem{}
	if err = faker.FakeObject(&mrli); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListMapRuns(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListMapRunsOutput{
			MapRuns: []types.MapRunListItem{mrli},
		}, nil)

	mapRunOut := sfn.DescribeMapRunOutput{}
	if err := faker.FakeObject(&mapRunOut); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeMapRun(gomock.Any(), gomock.Any(), gomock.Any()).Return(&mapRunOut, nil)

	return client.Services{
		Sfn: m,
	}
}

func TestExecutions(t *testing.T) {
	client.AwsMockTestHelper(t, Executions(), buildExecutions, client.TestOptions{})
}
