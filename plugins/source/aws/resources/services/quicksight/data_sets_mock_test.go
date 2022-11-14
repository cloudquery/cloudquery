package quicksight

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDataSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var ld quicksight.ListDataSetsOutput
	if err := faker.FakeObject(&ld); err != nil {
		t.Fatal(err)
	}
	ld.NextToken = nil
	m.EXPECT().ListDataSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ld, nil)

	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	var ig quicksight.ListIngestionsOutput
	if err := faker.FakeObject(&ig); err != nil {
		t.Fatal(err)
	}
	ig.NextToken = nil
	m.EXPECT().ListIngestions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ig, nil)

	var to2 quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to2); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to2, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightDataSets(t *testing.T) {
	client.AwsMockTestHelper(t, DataSets(), buildDataSetsMock, client.TestOptions{})
}
