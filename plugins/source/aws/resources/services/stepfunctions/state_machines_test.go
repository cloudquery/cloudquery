package stepfunctions

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildStateMachines(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSfnClient(ctrl)
	im := types.StateMachineListItem{}
	require.NoError(t, faker.FakeObject(&im))

	m.EXPECT().ListStateMachines(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListStateMachinesOutput{
			StateMachines: []types.StateMachineListItem{im},
		}, nil)

	out := &sfn.DescribeStateMachineOutput{}
	require.NoError(t, faker.FakeObject(&out))

	m.EXPECT().DescribeStateMachine(gomock.Any(), gomock.Any(), gomock.Any()).Return(out, nil)

	tag := types.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		}, nil)

	eli := types.ExecutionListItem{}
	require.NoError(t, faker.FakeObject(&eli))

	m.EXPECT().ListExecutions(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&sfn.ListExecutionsOutput{
			Executions: []types.ExecutionListItem{eli},
		}, nil)

	execOut := sfn.DescribeExecutionOutput{}
	require.NoError(t, faker.FakeObject(&execOut))

	m.EXPECT().DescribeExecution(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&execOut, nil)

	mrli := types.MapRunListItem{}
	require.NoError(t, faker.FakeObject(&mrli))
	m.EXPECT().ListMapRuns(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&sfn.ListMapRunsOutput{
			MapRuns: []types.MapRunListItem{mrli},
		}, nil)

	mapRunOut := sfn.DescribeMapRunOutput{}
	require.NoError(t, faker.FakeObject(&mapRunOut))

	m.EXPECT().DescribeMapRun(gomock.Any(), gomock.Any(), gomock.Any()).Return(&mapRunOut, nil)
	return client.Services{
		Sfn: m,
	}
}

func TestStateMachines(t *testing.T) {
	client.AwsMockTestHelper(t, StateMachines(), buildStateMachines, client.TestOptions{})
}
