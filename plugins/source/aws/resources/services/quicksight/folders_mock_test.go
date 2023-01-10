package quicksight

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildFoldersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var lo quicksight.ListFoldersOutput
	if err := faker.FakeObject(&lo); err != nil {
		t.Fatal(err)
	}
	lo.NextToken = nil
	m.EXPECT().ListFolders(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lo, nil)

	var qs quicksight.DescribeFolderOutput
	if err := faker.FakeObject(&qs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeFolder(gomock.Any(), gomock.Any(), gomock.Any()).Return(&qs, nil)

	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightFolders(t *testing.T) {
	client.AwsMockTestHelper(t, Folders(), buildFoldersMock, client.TestOptions{})
}
