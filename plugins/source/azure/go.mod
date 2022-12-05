module github.com/cloudquery/cloudquery/plugins/source/azure

go 1.19

require (
	github.com/Azure/azure-sdk-for-go v61.6.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions v1.0.0
	github.com/Azure/go-autorest/autorest v0.11.28
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/cloudquery/plugin-sdk v1.11.0
	github.com/gertd/go-pluralize v0.2.1
	github.com/gofrs/uuid v4.3.1+incompatible
	github.com/golang/mock v1.6.0
	github.com/iancoleman/strcase v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.28.0
	github.com/stretchr/testify v1.8.1
	github.com/tombuildsstuff/giovanni v0.20.0
	golang.org/x/exp v0.0.0-20221126150942-6ab00d035af9
)

require github.com/kylelemons/godebug v1.1.0 // indirect

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.1.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/addons/armaddons v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood v0.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement v0.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation v0.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commerce/armcommerce v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory v1.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager v0.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter v0.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub v0.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation v0.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fluidrelay/armfluidrelay v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice v0.1.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor v0.8.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi v0.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx v1.0.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep v0.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiembedded/armpowerbiembedded v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector v0.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy v0.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm v0.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity v0.9.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops v0.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine v0.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/visualstudio/armvisualstudio v0.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot v1.0.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads v0.2.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.5 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.15.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/thoas/go-funk v0.9.2 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/tools v0.2.0 // indirect
	google.golang.org/genproto v0.0.0-20221111202108-142d8a6fa32e // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
