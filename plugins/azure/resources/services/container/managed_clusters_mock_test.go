package container

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestContainerManagedClusters(t *testing.T) {
	client.AzureMockTestHelper(t, ContainerManagedClusters(), buildContainerManagedClusters, client.TestOptions{})
}

func buildContainerManagedClusters(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockManagedClustersClient(ctrl)

	// we resort to manually creating this data because faker fails to do it (recursive member)
	mc := containerservice.ManagedCluster{
		ManagedClusterProperties: fakeManagedClusterProperties(t),
	}
	require.Nil(t, faker.FakeData(&mc.Identity))
	require.Nil(t, faker.FakeData(&mc.Sku))
	require.Nil(t, faker.FakeData(&mc.ExtendedLocation))
	require.Nil(t, faker.FakeData(&mc.ID))
	require.Nil(t, faker.FakeData(&mc.Name))
	require.Nil(t, faker.FakeData(&mc.Type))
	require.Nil(t, faker.FakeData(&mc.Location))
	require.Nil(t, faker.FakeData(&mc.Tags))
	m.EXPECT().List(gomock.Any()).Return(
		containerservice.NewManagedClusterListResultPage(
			containerservice.ManagedClusterListResult{Value: &[]containerservice.ManagedCluster{mc}},
			func(context.Context, containerservice.ManagedClusterListResult) (containerservice.ManagedClusterListResult, error) {
				return containerservice.ManagedClusterListResult{}, nil
			},
		), nil,
	)
	return services.Services{
		Container: services.ContainerServiceClient{ManagedClusters: m},
	}
}

func fakeManagedClusterProperties(t *testing.T) *containerservice.ManagedClusterProperties {
	var mcp containerservice.ManagedClusterProperties
	require.Nil(t, faker.FakeData(&mcp.ProvisioningState))
	require.Nil(t, faker.FakeData(&mcp.PowerState))
	require.Nil(t, faker.FakeData(&mcp.MaxAgentPools))
	require.Nil(t, faker.FakeData(&mcp.KubernetesVersion))
	require.Nil(t, faker.FakeData(&mcp.DNSPrefix))
	require.Nil(t, faker.FakeData(&mcp.FqdnSubdomain))
	require.Nil(t, faker.FakeData(&mcp.Fqdn))
	require.Nil(t, faker.FakeData(&mcp.PrivateFQDN))
	require.Nil(t, faker.FakeData(&mcp.AzurePortalFQDN))
	require.Nil(t, faker.FakeData(&mcp.AgentPoolProfiles))
	require.Nil(t, faker.FakeData(&mcp.LinuxProfile))
	require.Nil(t, faker.FakeData(&mcp.WindowsProfile))
	require.Nil(t, faker.FakeData(&mcp.ServicePrincipalProfile))
	require.Nil(t, faker.FakeData(&mcp.AddonProfiles))
	require.Nil(t, faker.FakeData(&mcp.NodeResourceGroup))
	require.Nil(t, faker.FakeData(&mcp.EnableRBAC))
	require.Nil(t, faker.FakeData(&mcp.EnablePodSecurityPolicy))
	require.Nil(t, faker.FakeData(&mcp.NetworkProfile))
	require.Nil(t, faker.FakeData(&mcp.AadProfile))
	require.Nil(t, faker.FakeData(&mcp.AutoUpgradeProfile))
	require.Nil(t, faker.FakeData(&mcp.AutoScalerProfile))
	require.Nil(t, faker.FakeData(&mcp.APIServerAccessProfile))
	require.Nil(t, faker.FakeData(&mcp.DiskEncryptionSetID))
	require.Nil(t, faker.FakeData(&mcp.DisableLocalAccounts))
	require.Nil(t, faker.FakeData(&mcp.HTTPProxyConfig))
	require.Nil(t, faker.FakeData(&mcp.IdentityProfile))
	require.Nil(t, faker.FakeData(&mcp.PrivateLinkResources))
	require.Nil(t, faker.FakeData(&mcp.DisableLocalAccounts))

	var pip containerservice.ManagedClusterPodIdentityProfile
	require.Nil(t, faker.FakeData(&pip.Enabled))
	require.Nil(t, faker.FakeData(&pip.AllowNetworkPluginKubenet))
	require.Nil(t, faker.FakeData(&pip.UserAssignedIdentityExceptions))

	var pi containerservice.ManagedClusterPodIdentity
	require.Nil(t, faker.FakeData(&pi.Name))
	require.Nil(t, faker.FakeData(&pi.Namespace))
	require.Nil(t, faker.FakeData(&pi.BindingSelector))
	require.Nil(t, faker.FakeData(&pi.Identity))
	require.Nil(t, faker.FakeData(&pi.ProvisioningState))

	pip.UserAssignedIdentities = &[]containerservice.ManagedClusterPodIdentity{pi}
	mcp.PodIdentityProfile = &pip
	return &mcp
}
