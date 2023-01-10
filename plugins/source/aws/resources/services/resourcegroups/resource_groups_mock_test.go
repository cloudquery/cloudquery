package resourcegroups

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildResourceGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockResourcegroupsClient(ctrl)
	gId := types.GroupIdentifier{}
	err := faker.FakeObject(&gId)
	if err != nil {
		t.Fatal(err)
	}

	groupResponse := types.Group{}
	err = faker.FakeObject(&groupResponse)
	if err != nil {
		t.Fatal(err)
	}

	tagsResponse := resourcegroups.GetTagsOutput{}
	err = faker.FakeObject(&tagsResponse)
	if err != nil {
		t.Fatal(err)
	}

	query := types.GroupQuery{}
	err = faker.FakeObject(&query)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.ListGroupsOutput{
			GroupIdentifiers: []types.GroupIdentifier{gId},
		}, nil)
	m.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.GetGroupOutput{
			Group: &groupResponse,
		}, nil)
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsResponse, nil)
	m.EXPECT().GetGroupQuery(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.GetGroupQueryOutput{
			GroupQuery: &query,
		}, nil)

	return client.Services{
		Resourcegroups: m,
	}
}

func TestResourceGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceGroups(), buildResourceGroupsMock, client.TestOptions{})
}
