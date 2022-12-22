package elasticsearch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElasticSearchDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticsearchserviceClient(ctrl)

	var info types.DomainInfo
	if err := faker.FakeObject(&info); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListDomainNames(gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.ListDomainNamesOutput{
			DomainNames: []types.DomainInfo{info},
		},
		nil,
	)

	var ds types.ElasticsearchDomainStatus
	if err := faker.FakeObject(&ds); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeElasticsearchDomain(gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.DescribeElasticsearchDomainOutput{
			DomainStatus: &ds,
		},
		nil,
	)

	var principal types.AuthorizedPrincipal
	if err := faker.FakeObject(&principal); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListVpcEndpointAccess(gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.ListVpcEndpointAccessOutput{
			AuthorizedPrincipalList: []types.AuthorizedPrincipal{principal},
		},
		nil,
	)

	var tags elasticsearchservice.ListTagsOutput
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{Elasticsearchservice: m}
}

func TestElasticSearchDomains(t *testing.T) {
	client.AwsMockTestHelper(t, Domains(), buildElasticSearchDomains, client.TestOptions{})
}
