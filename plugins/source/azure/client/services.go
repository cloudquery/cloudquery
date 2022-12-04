package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

type Services struct {
	ArmadvisorOperationEntities        *armadvisor.OperationsClient
	ArmadvisorConfigData               *armadvisor.ConfigurationsClient
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
	ArmlogicWorkflows                  *armlogic.WorkflowsClient
	ArmmariadbServers                  *armmariadb.ServersClient
	ArmmysqlServers                    *armmysql.ServersClient
	ArmnetworkVirtualNetworks          *armnetwork.VirtualNetworksClient
	ArmnetworkSecurityGroups           *armnetwork.SecurityGroupsClient
	ArmnetworkInterfaces               *armnetwork.InterfacesClient
	ArmnetworkWatchers                 *armnetwork.WatchersClient
	ArmpostgresqlServers               *armpostgresql.ServersClient
	ArmredisCaches                     *armredis.Client
}

func InitServices(subscriptionId string, azCred azcore.TokenCredential) (Services, error) {
	var services Services

	// ArmadvisorOperationEntities, err := armadvisor.NewOperationsClient(subscriptionId, azCred, nil)
	// if err != nil {
		// return services, err
	// }
	// services.ArmadvisorOperationEntities = ArmadvisorOperationEntities

	ArmadvisorConfigData, err := armadvisor.NewConfigurationsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmadvisorConfigData = ArmadvisorConfigData

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

	ArmlogicWorkflows, err := armlogic.NewWorkflowsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmlogicWorkflows = ArmlogicWorkflows

	ArmmariadbServers, err := armmariadb.NewServersClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmmariadbServers = ArmmariadbServers

	ArmmysqlServers, err := armmysql.NewServersClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmmysqlServers = ArmmysqlServers

	ArmnetworkVirtualNetworks, err := armnetwork.NewVirtualNetworksClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmnetworkVirtualNetworks = ArmnetworkVirtualNetworks

	ArmnetworkSecurityGroups, err := armnetwork.NewSecurityGroupsClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmnetworkSecurityGroups = ArmnetworkSecurityGroups

	ArmnetworkInterfaces, err := armnetwork.NewInterfacesClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmnetworkInterfaces = ArmnetworkInterfaces

	ArmnetworkWatchers, err := armnetwork.NewWatchersClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmnetworkWatchers = ArmnetworkWatchers

	ArmpostgresqlServers, err := armpostgresql.NewServersClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmpostgresqlServers = ArmpostgresqlServers

	ArmredisCaches, err := armredis.NewClient(subscriptionId, azCred, nil)
	if err != nil {
		return services, err
	}
	services.ArmredisCaches = ArmredisCaches

	return services, nil
}
