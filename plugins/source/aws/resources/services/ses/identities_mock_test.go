package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIdentities(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	ei := types.IdentityInfo{}
	require.NoError(t, faker.FakeObject(&ei))

	o := sesv2.GetEmailIdentityOutput{}
	require.NoError(t, faker.FakeObject(&o))

	o.IdentityType = ei.IdentityType

	sesClient.EXPECT().ListEmailIdentities(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sesv2.ListEmailIdentitiesOutput{EmailIdentities: []types.IdentityInfo{ei}},
		nil,
	)
	sesClient.EXPECT().GetEmailIdentity(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	return client.Services{
		Sesv2: sesClient,
	}
}

func TestIdentities(t *testing.T) {
	client.AwsMockTestHelper(t, Identities(), buildIdentities, client.TestOptions{})
}
