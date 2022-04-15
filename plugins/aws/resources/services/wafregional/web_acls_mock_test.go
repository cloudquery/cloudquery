package wafregional

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWebACLsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafRegionalClient(ctrl)

	var acl types.WebACL
	if err := faker.FakeData(&acl); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListWebACLs(
		gomock.Any(),
		&wafregional.ListWebACLsInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListWebACLsOutput{
			WebACLs: []types.WebACLSummary{{WebACLId: acl.WebACLId}},
		},
		nil,
	)

	m.EXPECT().GetWebACL(
		gomock.Any(),
		&wafregional.GetWebACLInput{WebACLId: acl.WebACLId},
		gomock.Any(),
	).Return(
		&wafregional.GetWebACLOutput{WebACL: &acl},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: acl.WebACLArn,
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{},
		nil,
	)

	return client.Services{WafRegional: m}
}

func TestWebACLs(t *testing.T) {
	client.AwsMockTestHelper(t, WebAcls(), buildWebACLsMock, client.TestOptions{})
}
