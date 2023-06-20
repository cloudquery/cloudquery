package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildOriginRequestPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	cp := cloudfrontTypes.OriginRequestPolicySummary{}
	if err := faker.FakeObject(&cp); err != nil {
		t.Fatal(err)
	}

	cloudfrontOutput := &cloudfront.ListOriginRequestPoliciesOutput{
		OriginRequestPolicyList: &cloudfrontTypes.OriginRequestPolicyList{
			Items: []cloudfrontTypes.OriginRequestPolicySummary{cp},
		},
	}
	m.EXPECT().ListOriginRequestPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)
	return services
}

func TestCloudfrontOriginRequestPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, OriginRequestPolicies(), buildOriginRequestPoliciesMock, client.TestOptions{})
}
