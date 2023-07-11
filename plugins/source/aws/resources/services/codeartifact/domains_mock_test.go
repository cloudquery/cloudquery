package codeartifact

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodeartifactClient(ctrl)

	domainSummary := types.DomainSummary{}
	require.NoError(t, faker.FakeObject(&domainSummary))

	domain := types.DomainDescription{}
	require.NoError(t, faker.FakeObject(&domain))

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
	tag := types.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		},
		nil,
	)
	return client.Services{Codeartifact: m}
}

func TestDomains(t *testing.T) {
	client.AwsMockTestHelper(t, Domains(), buildDomains, client.TestOptions{})
}
