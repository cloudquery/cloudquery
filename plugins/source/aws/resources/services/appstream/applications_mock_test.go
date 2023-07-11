package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAppstreamApplicationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	application := types.Application{}
	require.NoError(t, faker.FakeObject(&application))

	applicationFleetAssociation := types.ApplicationFleetAssociation{}
	require.NoError(t, faker.FakeObject(&applicationFleetAssociation))

	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeApplicationsOutput{
			Applications: []types.Application{application},
		}, nil)

	m.EXPECT().DescribeApplicationFleetAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeApplicationFleetAssociationsOutput{
			ApplicationFleetAssociations: []types.ApplicationFleetAssociation{applicationFleetAssociation},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamApplications(t *testing.T) {
	client.AwsMockTestHelper(t, Applications(), buildAppstreamApplicationsMock, client.TestOptions{})
}
