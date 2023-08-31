package eks

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEksClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEksClient(ctrl)
	l := eks.DescribeClusterOutput{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListClustersOutput{
			Clusters: []string{"test-cluster"},
		}, nil)
	m.EXPECT().DescribeCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&l, nil)

	fp := types.FargateProfile{}
	require.NoError(t, faker.FakeObject(&fp))

	m.EXPECT().ListFargateProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListFargateProfilesOutput{
			FargateProfileNames: []string{"test-profile"},
		}, nil)
	m.EXPECT().DescribeFargateProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.DescribeFargateProfileOutput{FargateProfile: &fp}, nil)

	ng := types.Nodegroup{}
	require.NoError(t, faker.FakeObject(&ng))

	m.EXPECT().ListNodegroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListNodegroupsOutput{
			Nodegroups: []string{"test-nodegroup"},
		}, nil)
	m.EXPECT().DescribeNodegroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.DescribeNodegroupOutput{
			Nodegroup: &ng,
		}, nil)

	ao := types.Addon{}
	require.NoError(t, faker.FakeObject(&ao))

	m.EXPECT().ListAddons(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListAddonsOutput{
			Addons: []string{"test-nodegroup"},
		}, nil)
	m.EXPECT().DescribeAddon(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.DescribeAddonOutput{
			Addon: &ao,
		}, nil)

	ipc := types.IdentityProviderConfig{}
	require.NoError(t, faker.FakeObject(&ipc))
	ipc.Type = aws.String("oidc")
	m.EXPECT().ListIdentityProviderConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.ListIdentityProviderConfigsOutput{
			IdentityProviderConfigs: []types.IdentityProviderConfig{ipc},
		}, nil)
	oipc := types.OidcIdentityProviderConfig{}
	require.NoError(t, faker.FakeObject(&oipc))

	m.EXPECT().DescribeIdentityProviderConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eks.DescribeIdentityProviderConfigOutput{
			IdentityProviderConfig: &types.IdentityProviderConfigResponse{
				Oidc: &oipc,
			},
		}, nil)

	return client.Services{
		Eks: m,
	}
}

func TestEksClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildEksClusters, client.TestOptions{})
}
