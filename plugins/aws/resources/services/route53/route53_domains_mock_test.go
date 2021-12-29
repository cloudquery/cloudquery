// +build mock

package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRoute53Domains(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRoute53DomainsClient(ctrl)

	var ds types.DomainSummary
	if err := faker.FakeData(&ds.DomainName); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDomains(gomock.Any(), &route53domains.ListDomainsInput{}, gomock.Any()).Return(
		&route53domains.ListDomainsOutput{Domains: []types.DomainSummary{ds}},
		nil,
	)

	var detail route53domains.GetDomainDetailOutput
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}
	detail.DomainName = ds.DomainName
	mock.EXPECT().GetDomainDetail(gomock.Any(), &route53domains.GetDomainDetailInput{DomainName: ds.DomainName}, gomock.Any()).Return(
		&detail, nil,
	)

	var tagsOut route53domains.ListTagsForDomainOutput
	if err := faker.FakeData(&tagsOut); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTagsForDomain(gomock.Any(), &route53domains.ListTagsForDomainInput{DomainName: ds.DomainName}, gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		Route53Domains: mock,
	}
}

func TestRoute53Domains(t *testing.T) {
	client.AwsMockTestHelper(t, Route53Domains(), buildRoute53Domains, client.TestOptions{})
}
