package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIotBillingGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	groupsOutput := iot.ListBillingGroupsOutput{}
	err := faker.FakeObject(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListBillingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeBillingGroupOutput{}
	err = faker.FakeObject(&groupOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupOutput, nil)

	thingsInBillingGroupOutput := iot.ListThingsInBillingGroupOutput{}
	err = faker.FakeObject(&thingsInBillingGroupOutput)
	if err != nil {
		t.Fatal(err)
	}
	thingsInBillingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&thingsInBillingGroupOutput, nil)

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

func TestIotBillingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, BillingGroups(), buildIotBillingGroupsMock, client.TestOptions{})
}
