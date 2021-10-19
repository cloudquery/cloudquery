package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElasticSearchDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticSearch(ctrl)

	var di types.DomainInfo
	if err := faker.FakeData(&di); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListDomainNames(gomock.Any(), &elasticsearchservice.ListDomainNamesInput{}, gomock.Any()).Return(
		&elasticsearchservice.ListDomainNamesOutput{DomainNames: []types.DomainInfo{di}}, nil)

	var ds types.ElasticsearchDomainStatus
	if err := faker.FakeData(&ds); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeElasticsearchDomain(
		gomock.Any(),
		&elasticsearchservice.DescribeElasticsearchDomainInput{DomainName: di.DomainName},
		gomock.Any(),
	).Return(&elasticsearchservice.DescribeElasticsearchDomainOutput{DomainStatus: &ds}, nil)

	var tags elasticsearchservice.ListTagsOutput
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{ElasticSearch: m}
}

func TestElasticSearchDomains(t *testing.T) {
	awsTestHelper(t, ElasticsearchDomains(), buildElasticSearchDomains, TestOptions{})
}
