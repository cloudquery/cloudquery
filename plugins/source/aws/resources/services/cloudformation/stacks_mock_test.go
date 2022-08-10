package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildStacks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudFormationClient(ctrl)

	var stack types.Stack
	if err := faker.FakeData(&stack); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeStacks(
		gomock.Any(),
		&cloudformation.DescribeStacksInput{},
		gomock.Any(),
	).Return(
		&cloudformation.DescribeStacksOutput{Stacks: []types.Stack{stack}},
		nil,
	)

	var resource types.StackResourceSummary
	if err := faker.FakeData(&resource); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListStackResources(
		gomock.Any(),
		&cloudformation.ListStackResourcesInput{StackName: stack.StackName},
		gomock.Any(),
	).Return(
		&cloudformation.ListStackResourcesOutput{StackResourceSummaries: []types.StackResourceSummary{resource}},
		nil,
	)

	return client.Services{Cloudformation: mock}
}

func TestCloudformationStacks(t *testing.T) {
	client.AwsMockTestHelper(t, Stacks(), buildStacks, client.TestOptions{})
}
