package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotBillingGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	groupsOutput := iot.ListBillingGroupsOutput{}
	err := faker.FakeData(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListBillingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeBillingGroupOutput{}
	err = faker.FakeData(&groupOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupOutput, nil)

	thingsInBillingGroupOutput := iot.ListThingsInBillingGroupOutput{}
	err = faker.FakeData(&thingsInBillingGroupOutput)
	if err != nil {
		t.Fatal(err)
	}
	thingsInBillingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&thingsInBillingGroupOutput, nil)

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

func TestIotBillingGroups(t *testing.T) {
	client.AwsMockTestHelper(t, IotBillingGroups(), buildIotBillingGroupsMock, client.TestOptions{})
}
