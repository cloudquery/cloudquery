package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

type Services struct {
	ArmbatchAccounts                   *armbatch.AccountClient
	ArmcdnProfiles                     *armcdn.ProfilesClient
	ArmcdnEndpoints                    *armcdn.EndpointsClient
	ArmcomputeDisks                    *armcompute.DisksClient
	ArmcomputeVirtualMachines          *armcompute.VirtualMachinesClient
	ArmcomputeVirtualMachineScaleSets  *armcompute.VirtualMachineScaleSetsClient
	ArmcontainerregistryRegistries     *armcontainerregistry.RegistriesClient
	ArmcontainerserviceManagedClusters *armcontainerservice.ManagedClustersClient
	ArmcosmosAccounts                  *armcosmos.DatabaseAccountsClient
	ArmdatalakestoreAccounts           *armdatalakestore.AccountsClient
	ArmdatalakeanalyticsAccounts       *armdatalakeanalytics.AccountsClient
	ArmeventhubNamespaces              *armeventhub.NamespacesClient
	ArmeventhubNetworkRuleSets         *armeventhub.NamespacesClient
	ArmfrontdoorDoors                  *armfrontdoor.FrontDoorsClient
	ArmkeyvaultVaults                  *armkeyvault.VaultsClient
}

func InitServices(subscriptionId string, azCred azcore.TokenCredential) (Services, error) {
	var services Services

	ArmbatchAccounts, err := armbatch.NewAccountClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmbatchAccounts = ArmbatchAccounts

	ArmcdnProfiles, err := armcdn.NewProfilesClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcdnProfiles = ArmcdnProfiles

	ArmcdnEndpoints, err := armcdn.NewEndpointsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcdnEndpoints = ArmcdnEndpoints

	ArmcomputeDisks, err := armcompute.NewDisksClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcomputeDisks = ArmcomputeDisks

	ArmcomputeVirtualMachines, err := armcompute.NewVirtualMachinesClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcomputeVirtualMachines = ArmcomputeVirtualMachines

	ArmcomputeVirtualMachineScaleSets, err := armcompute.NewVirtualMachineScaleSetsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcomputeVirtualMachineScaleSets = ArmcomputeVirtualMachineScaleSets

	ArmcontainerregistryRegistries, err := armcontainerregistry.NewRegistriesClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcontainerregistryRegistries = ArmcontainerregistryRegistries

	ArmcontainerserviceManagedClusters, err := armcontainerservice.NewManagedClustersClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcontainerserviceManagedClusters = ArmcontainerserviceManagedClusters

	ArmcosmosAccounts, err := armcosmos.NewDatabaseAccountsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmcosmosAccounts = ArmcosmosAccounts

	ArmdatalakestoreAccounts, err := armdatalakestore.NewAccountsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmdatalakestoreAccounts = ArmdatalakestoreAccounts

	ArmdatalakeanalyticsAccounts, err := armdatalakeanalytics.NewAccountsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmdatalakeanalyticsAccounts = ArmdatalakeanalyticsAccounts

	ArmeventhubNamespaces, err := armeventhub.NewNamespacesClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmeventhubNamespaces = ArmeventhubNamespaces

	ArmeventhubNetworkRuleSets, err := armeventhub.NewNamespacesClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmeventhubNetworkRuleSets = ArmeventhubNetworkRuleSets

	ArmfrontdoorDoors, err := armfrontdoor.NewFrontDoorsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmfrontdoorDoors = ArmfrontdoorDoors

	ArmkeyvaultVaults, err := armkeyvault.NewVaultsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmkeyvaultVaults = ArmkeyvaultVaults

	return services, nil
}
