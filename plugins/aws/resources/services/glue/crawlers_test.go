package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCrawlers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var crawler glue.GetCrawlersOutput
	if err := faker.FakeData(&crawler); err != nil {
		t.Fatal(err)
	}
	crawler.NextToken = nil
	m.EXPECT().GetCrawlers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&crawler, nil)

	var tags glue.GetTagsOutput
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{
		Glue: m,
	}
}

func TestCrawlers(t *testing.T) {
	client.AwsMockTestHelper(t, Crawlers(), buildCrawlers, client.TestOptions{})
}
