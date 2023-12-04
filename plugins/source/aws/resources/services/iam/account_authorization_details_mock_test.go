package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIamAccountAuthDetails(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)

	details := &iam.GetAccountAuthorizationDetailsOutput{}
	require.NoError(t, faker.FakeObject(details))
	details.Marker = nil
	details.IsTruncated = false
	m.EXPECT().GetAccountAuthorizationDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(details, nil)
	return client.Services{
		Iam: m,
	}
}
func TestIamAccountAuthDetails(t *testing.T) {
	client.AwsMockTestHelper(t, AccountAuthorizationDetails(), buildIamAccountAuthDetails, client.TestOptions{})
}
