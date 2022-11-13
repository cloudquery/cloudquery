package quicksight

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var lo quicksight.ListGroupsOutput
	if err := faker.FakeObject(&lo); err != nil {
		t.Fatal(err)
	}
	lo.NextToken = nil
	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lo, nil)

	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	var gm quicksight.ListGroupMembershipsOutput
	if err := faker.FakeObject(&gm); err != nil {
		t.Fatal(err)
	}
	gm.NextToken = nil
	m.EXPECT().ListGroupMemberships(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gm, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildGroupsMock, client.TestOptions{})
}
