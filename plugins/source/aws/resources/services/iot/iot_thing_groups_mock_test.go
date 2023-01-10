package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIotThingGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	groupsOutput := iot.ListThingGroupsOutput{}
	err := faker.FakeObject(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListThingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeThingGroupOutput{}
	err = faker.FakeObject(&groupOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupOutput, nil)

	thingsInThingGroupOutput := iot.ListThingsInThingGroupOutput{}
	err = faker.FakeObject(&thingsInThingGroupOutput)
	if err != nil {
		t.Fatal(err)
	}
	thingsInThingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&thingsInThingGroupOutput, nil)

	p := iot.ListAttachedPoliciesOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	p.NextMarker = nil
	m.EXPECT().ListAttachedPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
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
