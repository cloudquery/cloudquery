package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRoute53Domains(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRoute53domainsClient(ctrl)

	var ds types.DomainSummary
	require.NoError(t, faker.FakeObject(&ds.DomainName))

	mock.EXPECT().ListDomains(gomock.Any(), &route53domains.ListDomainsInput{}, gomock.Any()).Return(
		&route53domains.ListDomainsOutput{Domains: []types.DomainSummary{ds}},
		nil,
	)

	var detail route53domains.GetDomainDetailOutput
	require.NoError(t, faker.FakeObject(&detail))

	detail.DomainName = ds.DomainName
	mock.EXPECT().GetDomainDetail(gomock.Any(), &route53domains.GetDomainDetailInput{DomainName: ds.DomainName}, gomock.Any()).Return(
		&detail, nil,
	)

	var tagsOut route53domains.ListTagsForDomainOutput
	require.NoError(t, faker.FakeObject(&tagsOut))

	mock.EXPECT().ListTagsForDomain(gomock.Any(), &route53domains.ListTagsForDomainInput{DomainName: ds.DomainName}, gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		Route53domains: mock,
	}
}

func TestRoute53Domains(t *testing.T) {
	client.AwsMockTestHelper(t, Domains(), buildRoute53Domains, client.TestOptions{})
}
