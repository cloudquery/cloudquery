package mwaa

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildMwaaEnvironments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockMwaaClient(ctrl)
	g := types.Environment{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().ListEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mwaa.ListEnvironmentsOutput{
			Environments: []string{*g.Name},
		}, nil)
	m.EXPECT().GetEnvironment(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mwaa.GetEnvironmentOutput{
			Environment: &g,
		}, nil)
	return client.Services{
		Mwaa: m,
	}
}

func TestMwaaEnvironments(t *testing.T) {
	client.AwsMockTestHelper(t, Environments(), buildMwaaEnvironments, client.TestOptions{})
}
