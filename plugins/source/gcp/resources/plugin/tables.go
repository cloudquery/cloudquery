package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/aiplatform"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/apigateway"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/apikeys"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/appengine"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/artifactregistry"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/baremetalsolution"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/beyondcorp"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/bigquery"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/bigtableadmin"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/billing"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/binaryauthorization"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/certificatemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/clouddeploy"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/clouderrorreporting"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudiot"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudresourcemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudscheduler"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudsupport"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/container"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/containeranalysis"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/dns"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/domains"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/functions"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/iam"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/kms"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/livestream"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/logging"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/monitoring"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/networkconnectivity"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/redis"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/resourcemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/run"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/secretmanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/securitycenter"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/serviceusage"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/translate"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/videotranscoder"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/vision"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/vmmigration"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/vpcaccess"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/websecurityscanner"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/workflows"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PluginAutoGeneratedTables() schema.Tables {
	tables := []*schema.Table{
		aiplatform.DatasetLocations(),
		aiplatform.EndpointLocations(),
		aiplatform.FeaturestoreLocations(),
		aiplatform.IndexendpointLocations(),
		aiplatform.IndexLocations(),
		aiplatform.JobLocations(),
		aiplatform.MetadataLocations(),
		aiplatform.ModelLocations(),
		aiplatform.Operations(),
		aiplatform.PipelineLocations(),
		aiplatform.SpecialistpoolLocations(),
		aiplatform.TensorboardLocations(),
		aiplatform.VizierLocations(),
		apigateway.Apis(),
		apigateway.Gateways(),
		apikeys.Keys(),
		appengine.Apps(),
		appengine.AuthorizedCertificates(),
		appengine.AuthorizedDomains(),
		appengine.DomainMappings(),
		appengine.FirewallIngressRules(),
		appengine.Services(),
		artifactregistry.Locations(),
		baremetalsolution.Instances(),
		baremetalsolution.Networks(),
		baremetalsolution.NfsShares(),
		baremetalsolution.Volumes(),
		batch.Jobs(),
		batch.TaskGroups(),
		beyondcorp.AppConnections(),
		beyondcorp.AppConnectors(),
		beyondcorp.AppGateways(),
		beyondcorp.ClientConnectorServices(),
		beyondcorp.ClientGateways(),
		bigquery.Datasets(),
		bigtableadmin.Instances(),
		billing.BillingAccounts(),
		billing.Services(),
		binaryauthorization.Assertors(),
		certificatemanager.CertificateIssuanceConfigs(),
		certificatemanager.CertificateMaps(),
		certificatemanager.Certificates(),
		certificatemanager.DnsAuthorizations(),
		clouddeploy.DeliveryPipelines(),
		clouddeploy.Targets(),
		clouderrorreporting.ErrorGroupStats(),
		cloudiot.DeviceRegistries(),
		cloudresourcemanager.Organizations(),
		cloudscheduler.Locations(),
		cloudsupport.Cases(),
		compute.Addresses(),
		compute.Autoscalers(),
		compute.BackendServices(),
		compute.Disks(),
		compute.DiskTypes(),
		compute.ExternalVpnGateways(),
		compute.Firewalls(),
		compute.ForwardingRules(),
		compute.Images(),
		compute.InstanceGroups(),
		compute.Instances(),
		compute.InterconnectAttachments(),
		compute.InterconnectLocations(),
		compute.InterconnectRemoteLocations(),
		compute.Interconnects(),
		compute.Networks(),
		compute.Projects(),
		compute.Routers(),
		compute.Routes(),
		compute.SslCertificates(),
		compute.SslPolicies(),
		compute.Subnetworks(),
		compute.TargetHttpProxies(),
		compute.TargetSslProxies(),
		compute.TargetVpnGateways(),
		compute.UrlMaps(),
		compute.VpnGateways(),
		compute.VpnTunnels(),
		compute.Zones(),
		container.Clusters(),
		containeranalysis.Occurrences(),
		dns.ManagedZones(),
		dns.Policies(),
		domains.Registrations(),
		functions.Functions(),
		iam.DenyPolicies(),
		iam.Roles(),
		iam.ServiceAccounts(),
		kms.Locations(),
		livestream.Channels(),
		livestream.Inputs(),
		logging.Metrics(),
		logging.Sinks(),
		monitoring.AlertPolicies(),
		networkconnectivity.Locations(),
		redis.Instances(),
		resourcemanager.Folders(),
		resourcemanager.OrganizationTagKeys(),
		resourcemanager.ProjectPolicies(),
		resourcemanager.Projects(),
		resourcemanager.ProjectTagBindings(),
		resourcemanager.ProjectTagKeys(),
		run.Locations(),
		secretmanager.Secrets(),
		securitycenter.FolderFindings(),
		securitycenter.OrganizationFindings(),
		securitycenter.ProjectFindings(),
		services.Projects(),
		serviceusage.Services(),
		sql.Instances(),
		storage.Buckets(),
		translate.Glossaries(),
		videotranscoder.Jobs(),
		videotranscoder.JobTemplates(),
		vision.Products(),
		vmmigration.Groups(),
		vmmigration.Sources(),
		vmmigration.TargetProjects(),
		vpcaccess.Locations(),
		websecurityscanner.ScanConfigs(),
		workflows.Workflows(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}
