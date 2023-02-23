package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	services := []types.Service{}
	err := faker.FakeObject(&services)
	if err != nil {
		t.Fatal(err)
	}

	for _, languageCode := range servicesSupportedLanguageCodes {
		m.EXPECT().DescribeServices(gomock.Any(), &support.DescribeServicesInput{Language: aws.String(languageCode)}).Return(&support.DescribeServicesOutput{Services: services}, nil)
	}

	return client.Services{
		Support: m,
	}
}

func TestServices(t *testing.T) {
	client.AwsMockTestHelper(t, Services(), buildServices, client.TestOptions{})
}
