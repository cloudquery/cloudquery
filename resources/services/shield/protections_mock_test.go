package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildProtections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	protection := shield.ListProtectionsOutput{}
	err := faker.FakeData(&protection)
	if err != nil {
		t.Fatal(err)
	}
	protection.NextToken = nil
	m.EXPECT().ListProtections(gomock.Any(), gomock.Any(), gomock.Any()).Return(&protection, nil)

	tags := shield.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)
	return client.Services{
		Shield: m,
	}
}

func TestProtections(t *testing.T) {
	client.AwsMockTestHelper(t, Protections(), buildProtections, client.TestOptions{})
}
