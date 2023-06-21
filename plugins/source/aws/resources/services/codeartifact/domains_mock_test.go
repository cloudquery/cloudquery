package codeartifact

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodeartifactClient(ctrl)

	domainSummary := types.DomainSummary{}
	if err := faker.FakeObject(&domainSummary); err != nil {
		t.Fatal(err)
	}
	domain := types.DomainDescription{}
	if err := faker.FakeObject(&domain); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListDomains(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.ListDomainsOutput{
			Domains:   []types.DomainSummary{domainSummary},
			NextToken: nil,
		},
		nil,
	)

	m.EXPECT().DescribeDomain(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.DescribeDomainOutput{
			Domain: &domain,
		},
		nil,
	)

	return client.Services{Codeartifact: m}
}

func TestDomains(t *testing.T) {
	client.AwsMockTestHelper(t, Domains(), buildDomains, client.TestOptions{})
}
