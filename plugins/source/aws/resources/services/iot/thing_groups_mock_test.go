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

func buildIotThingGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	groupsOutput := iot.ListThingGroupsOutput{}
	require.NoError(t, faker.FakeObject(&groupsOutput))
	groupsOutput.NextToken = nil
	m.EXPECT().ListThingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeThingGroupOutput{}
	require.NoError(t, faker.FakeObject(&groupOutput))
	m.EXPECT().DescribeThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupOutput, nil)

	thingsInThingGroupOutput := iot.ListThingsInThingGroupOutput{}
	require.NoError(t, faker.FakeObject(&thingsInThingGroupOutput))
	thingsInThingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&thingsInThingGroupOutput, nil)

	p := iot.ListAttachedPoliciesOutput{}
	require.NoError(t, faker.FakeObject(&p))
	p.NextMarker = nil
	m.EXPECT().ListAttachedPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotThingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ThingGroups(), buildIotThingGroupsMock, client.TestOptions{})
}
