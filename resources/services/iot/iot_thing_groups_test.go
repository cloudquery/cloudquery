//go:build mock
// +build mock

package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotThingGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	groupsOutput := iot.ListThingGroupsOutput{}
	err := faker.FakeData(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListThingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeThingGroupOutput{}
	err = faker.FakeData(&groupOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupOutput, nil)

	thingsInThingGroupOutput := iot.ListThingsInThingGroupOutput{}
	err = faker.FakeData(&thingsInThingGroupOutput)
	if err != nil {
		t.Fatal(err)
	}
	thingsInThingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&thingsInThingGroupOutput, nil)

	p := iot.ListAttachedPoliciesOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	p.NextMarker = nil
	m.EXPECT().ListAttachedPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		IOT: m,
	}
}

func TestIotThingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, IotThingGroups(), buildIotThingGroupsMock, client.TestOptions{})
}
