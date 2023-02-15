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

func buildStateMachines(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSfnClient(ctrl)
	im := types.StateMachineListItem{}
	err := faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListStateMachines(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListStateMachinesOutput{
			StateMachines: []types.StateMachineListItem{im},
		}, nil)

	out := &sfn.DescribeStateMachineOutput{}
	err = faker.FakeObject(&out)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeStateMachine(gomock.Any(), gomock.Any(), gomock.Any()).Return(out, nil)

	tag := types.Tag{}
	tagerr := faker.FakeObject(&tag)
	if tagerr != nil {
		t.Fatal(tagerr)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		}, nil)

	eli := types.ExecutionListItem{}
	if err := faker.FakeObject(&eli); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListExecutions(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&sfn.ListExecutionsOutput{
			Executions: []types.ExecutionListItem{eli},
		}, nil)

	execOut := sfn.DescribeExecutionOutput{}
	if err := faker.FakeObject(&execOut); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeExecution(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&execOut, nil)

	mrli := types.MapRunListItem{}
	if err = faker.FakeObject(&mrli); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListMapRuns(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
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

func TestStateMachines(t *testing.T) {
	client.AwsMockTestHelper(t, StateMachines(), buildStateMachines, client.TestOptions{})
}
