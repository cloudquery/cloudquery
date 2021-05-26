package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildConfigConformancePack(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigServiceClient(ctrl)

	cp := types.ConformancePackDetail{}
	err := faker.FakeData(&cp)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConformancePacks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConformancePacksOutput{
			ConformancePackDetails: []types.ConformancePackDetail{cp},
		}, nil)

	return client.Services{
		ConfigService: m,
	}
}

func TestConfigConformancePack(t *testing.T) {
	awsTestHelper(t, ConfigConformancePack(), buildConfigConformancePack)
}
