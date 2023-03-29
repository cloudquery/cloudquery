package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildStackSet(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudformationClient(ctrl)

	var stack types.StackSet
	if err := faker.FakeObject(&stack); err != nil {
		t.Fatal(err)
	}

	var stackSummary types.StackSetSummary
	if err := faker.FakeObject(&stackSummary); err != nil {
		t.Fatal(err)
	}

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
	if err := faker.FakeObject(&stackSetOperationSummary); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListStackSetOperations(
		gomock.Any(),
		&cloudformation.ListStackSetOperationsInput{StackSetName: stackSummary.StackSetName, CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.ListStackSetOperationsOutput{Summaries: []types.StackSetOperationSummary{stackSetOperationSummary}},
		nil,
	)

	var stackSetOperation types.StackSetOperation
	if err := faker.FakeObject(&stackSetOperation); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeStackSetOperation(
		gomock.Any(),
		&cloudformation.DescribeStackSetOperationInput{StackSetName: stackSummary.StackSetName, OperationId: stackSetOperationSummary.OperationId, CallAs: types.CallAsDelegatedAdmin},
		gomock.Any(),
	).Return(
		&cloudformation.DescribeStackSetOperationOutput{StackSetOperation: &stackSetOperation},
		nil,
	)

	var stackSetOperationResultSummary types.StackSetOperationResultSummary
	if err := faker.FakeObject(&stackSetOperationResultSummary); err != nil {
		t.Fatal(err)
	}

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
