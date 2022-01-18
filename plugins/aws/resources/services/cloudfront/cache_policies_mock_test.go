//go:build mock
// +build mock

package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudfrontCachePoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	cp := cloudfrontTypes.CachePolicySummary{}
	if err := faker.FakeData(&cp); err != nil {
		t.Fatal(err)
	}

	cloudfrontOutput := &cloudfront.ListCachePoliciesOutput{
		CachePolicyList: &cloudfrontTypes.CachePolicyList{
			Items: []cloudfrontTypes.CachePolicySummary{cp},
		},
	}
	m.EXPECT().ListCachePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)
	return services
}

func TestCloudfrontCachePolicies(t *testing.T) {
	client.AwsMockTestHelper(t, CloudfrontCachePolicies(), buildCloudfrontCachePoliciesMock, client.TestOptions{})
}
