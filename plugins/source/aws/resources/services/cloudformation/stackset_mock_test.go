package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildStackSet(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudformationClient(ctrl)

	var stack types.StackSet
	require.NoError(t, faker.FakeObject(&stack))

	var stackSummary types.StackSetSummary
	require.NoError(t, faker.FakeObject(&stackSummary))

	mock.EXPECT().ListStackSets(
		gomock.Any(),
		&cloudformation.ListStackSetsInput{CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.ListStackSetsOutput{Summaries: []types.StackSetSummary{stackSummary}},
		nil,
	)

	mock.EXPECT().DescribeStackSet(
		gomock.Any(),
		&cloudformation.DescribeStackSetInput{StackSetName: stackSummary.StackSetName, CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.DescribeStackSetOutput{StackSet: &stack},
		nil,
	)

	var stackSetOperationSummary types.StackSetOperationSummary
	require.NoError(t, faker.FakeObject(&stackSetOperationSummary))

	mock.EXPECT().ListStackSetOperations(
		gomock.Any(),
		&cloudformation.ListStackSetOperationsInput{StackSetName: stackSummary.StackSetName, CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.ListStackSetOperationsOutput{Summaries: []types.StackSetOperationSummary{stackSetOperationSummary}},
		nil,
	)

	var stackSetOperation types.StackSetOperation
	require.NoError(t, faker.FakeObject(&stackSetOperation))

	mock.EXPECT().DescribeStackSetOperation(
		gomock.Any(),
		&cloudformation.DescribeStackSetOperationInput{StackSetName: stackSummary.StackSetName, OperationId: stackSetOperationSummary.OperationId, CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.DescribeStackSetOperationOutput{StackSetOperation: &stackSetOperation},
		nil,
	)

	var stackSetOperationResultSummary types.StackSetOperationResultSummary
	require.NoError(t, faker.FakeObject(&stackSetOperationResultSummary))

	mock.EXPECT().ListStackSetOperationResults(
		gomock.Any(),
		&cloudformation.ListStackSetOperationResultsInput{StackSetName: stackSummary.StackSetName, OperationId: stackSetOperationSummary.OperationId, CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.ListStackSetOperationResultsOutput{Summaries: []types.StackSetOperationResultSummary{stackSetOperationResultSummary}},
		nil,
	)

	return client.Services{Cloudformation: mock}
}

func TestCloudformationStackSet(t *testing.T) {
	client.AwsMockTestHelper(t, StackSets(), buildStackSet, client.TestOptions{})
}
