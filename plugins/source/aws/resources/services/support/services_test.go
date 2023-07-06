package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	services := []types.Service{}
	require.NoError(t, faker.FakeObject(&services))

	for _, languageCode := range servicesSupportedLanguageCodes {
		m.EXPECT().DescribeServices(gomock.Any(), &support.DescribeServicesInput{Language: aws.String(languageCode)}, gomock.Any()).Return(&support.DescribeServicesOutput{Services: services}, nil)
	}

	return client.Services{
		Support: m,
	}
}

func TestServices(t *testing.T) {
	client.AwsMockTestHelper(t, Services(), buildServices, client.TestOptions{})
}
