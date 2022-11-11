package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildProtectionGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	pp := shield.ListProtectionGroupsOutput{}
	err := faker.FakeObject(&pp)
	if err != nil {
		t.Fatal(err)
	}
	pp.NextToken = nil
	m.EXPECT().ListProtectionGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&pp, nil)

	tags := shield.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{
		Shield: m,
	}
}

func TestProtectionGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ProtectionGroups(), buildProtectionGroups, client.TestOptions{})
}
