package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIotBillingGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	groupsOutput := iot.ListBillingGroupsOutput{}
	require.NoError(t, faker.FakeObject(&groupsOutput))
	groupsOutput.NextToken = nil
	m.EXPECT().ListBillingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeBillingGroupOutput{}
	require.NoError(t, faker.FakeObject(&groupOutput))
	m.EXPECT().DescribeBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupOutput, nil)

	thingsInBillingGroupOutput := iot.ListThingsInBillingGroupOutput{}
	require.NoError(t, faker.FakeObject(&thingsInBillingGroupOutput))
	thingsInBillingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&thingsInBillingGroupOutput, nil)

	tags := iot.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotBillingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, BillingGroups(), buildIotBillingGroupsMock, client.TestOptions{})
}
