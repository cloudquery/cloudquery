package container

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func TestContainerRegistries(t *testing.T) {
	client.AzureMockTestHelper(t, ContainerRegistries(), buildContainerRegistries, client.TestOptions{})
}

func buildContainerRegistries(t *testing.T, ctrl *gomock.Controller) services.Services {
	reg := mocks.NewMockContainerRegistriesClient(ctrl)
	rep := mocks.NewMockContainerReplicationsClient(ctrl)

	// we resort to manually creating this data because faker fails to do it (recursive member)
	registry := containerregistry.Registry{}
	err := faker.FakeDataSkipFields(&registry, []string{"RegistryProperties"})
	if err != nil {
		t.Fatal(err)
	}
	id := client.FakeResourceGroup
	registry.ID = &id
	registry.RegistryProperties = fakeContainerRegistryProperties(t)
	reg.EXPECT().List(gomock.Any()).Return(
		containerregistry.NewRegistryListResultPage(
			containerregistry.RegistryListResult{Value: &[]containerregistry.Registry{registry}},
			func(context.Context, containerregistry.RegistryListResult) (containerregistry.RegistryListResult, error) {
				return containerregistry.RegistryListResult{}, nil
			},
		), nil,
	)

	// we resort to manually creating this data because faker fails to do it (recursive member)
	replication := containerregistry.Replication{}
	err = faker.FakeDataSkipFields(&replication, []string{"ReplicationProperties"})
	if err != nil {
		t.Fatal(err)
	}
	replication.ReplicationProperties = fakeContainerReplicationProperties(t)
	rep.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		containerregistry.NewReplicationListResultPage(
			containerregistry.ReplicationListResult{Value: &[]containerregistry.Replication{replication}},
			func(context.Context, containerregistry.ReplicationListResult) (containerregistry.ReplicationListResult, error) {
				return containerregistry.ReplicationListResult{}, nil
			},
		), nil,
	)
	return services.Services{
		ContainerRegistry: services.ContainerRegistryClient{Registries: reg, Replications: rep},
	}
}

func fakeContainerRegistryProperties(t *testing.T) *containerregistry.RegistryProperties {
	var mcp containerregistry.RegistryProperties
	err := faker.FakeDataSkipFields(&mcp, []string{"ProvisioningState"})
	if err != nil {
		t.Fatal(err)
	}
	mcp.ProvisioningState = "test"
	cidr := faker.IPv4() + "/24"
	(*mcp.NetworkRuleSet.IPRules)[0].IPAddressOrRange = &cidr
	return &mcp
}

func fakeContainerReplicationProperties(t *testing.T) *containerregistry.ReplicationProperties {
	var mcp containerregistry.ReplicationProperties
	err := faker.FakeDataSkipFields(&mcp, []string{"ProvisioningState"})
	if err != nil {
		t.Fatal(err)
	}
	mcp.ProvisioningState = "test"
	return &mcp
}
