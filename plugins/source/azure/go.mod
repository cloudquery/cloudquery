module github.com/cloudquery/cloudquery/plugins/source/azure

go 1.19

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.4.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.2
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation v0.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation v0.7.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4 v4.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware v0.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2 v2.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2 v2.3.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2 v2.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3 v3.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics v0.6.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub v0.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic v0.6.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric v0.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor v0.8.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2 v2.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v2 v2.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2 v2.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3 v3.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy v0.6.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions v1.1.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity v0.9.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine v0.7.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3 v3.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse v0.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads v0.2.0
	github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue v0.1.0
	github.com/apache/arrow/go/v13 v13.0.0-20230630125530-5a06b2ec2a8e
	github.com/cloudquery/plugin-sdk/v4 v4.8.1-rc1
	github.com/cockroachdb/cockroachdb-parser v0.0.0-20230515042840-c9c144eab71a
	github.com/gorilla/mux v1.8.0
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mjibson/sqlfmt v0.5.0
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
	golang.org/x/sync v0.1.0
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230710001530-a2a76ebbb85f

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.2.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.9.0 // indirect
	github.com/biogo/store v0.0.0-20201120204734-aad293a2328f // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/cloudquery/plugin-pb-go v1.6.0 // indirect
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect
	github.com/cockroachdb/errors v1.9.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/dave/dst v0.27.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/flatbuffers v23.1.21+incompatible // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/petermattis/goid v0.0.0-20211229010228-4d14c490ee36 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pierrre/geohash v1.0.0 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twpayne/go-geom v1.4.2 // indirect
	github.com/twpayne/go-kml v1.5.2 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20230525234025-438c736192d0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230525234020-1aefcd67740a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
