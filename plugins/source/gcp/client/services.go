package client

import (
	"context"

	billing "cloud.google.com/go/billing/apiv1"
	compute "cloud.google.com/go/compute/apiv1"
	container "cloud.google.com/go/container/apiv1"
	domains "cloud.google.com/go/domains/apiv1beta1"
	functions "cloud.google.com/go/functions/apiv1"
	kms "cloud.google.com/go/kms/apiv1"
	logging "cloud.google.com/go/logging/apiv2"
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	redis "cloud.google.com/go/redis/apiv1"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	run "cloud.google.com/go/run/apiv2"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	"cloud.google.com/go/storage"
	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/dns/v1"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

type GcpService string

type Services struct {
	Bigquery                      *bigquery.Service
	BillingCloudBillingClient     *billing.CloudBillingClient
	BillingCloudCatalogClient     *billing.CloudCatalogClient
	FunctionsCloudFunctionsClient *functions.CloudFunctionsClient

	ComputeAddressesClient         *compute.AddressesClient
	ComputeAutoscalersClient       *compute.AutoscalersClient
	ComputeBackendServicesClient   *compute.BackendServicesClient
	ComputeDiskTypesClient         *compute.DiskTypesClient
	ComputeDisksClient             *compute.DisksClient
	ComputeFirewallsClient         *compute.FirewallsClient
	ComputeForwardingRulesClient   *compute.ForwardingRulesClient
	ComputeImagesClient            *compute.ImagesClient
	ComputeInstancesClient         *compute.InstancesClient
	ComputeInstanceGroupsClient    *compute.InstanceGroupsClient
	ComputeInterconnectsClient     *compute.InterconnectsClient
	ComputeNetworksClient          *compute.NetworksClient
	ComputeProjectsClient          *compute.ProjectsClient
	ComputeSslCertificatesClient   *compute.SslCertificatesClient
	ComputeSslPoliciesClient       *compute.SslPoliciesClient
	ComputeSubnetworksClient       *compute.SubnetworksClient
	ComputeTargetHttpProxiesClient *compute.TargetHttpProxiesClient
	ComputeTargetSslProxiesClient  *compute.TargetSslProxiesClient
	ComputeUrlMapsClient           *compute.UrlMapsClient
	ComputeVpnGatewaysClient       *compute.VpnGatewaysClient

	RunServicesClient *run.ServicesClient

	ContainerClusterManagerClient *container.ClusterManagerClient
	Dns                           *dns.Service

	DomainsClient                 *domains.Client
	Iam                           *iam.Service
	KmsKeyManagementClient        *kms.KeyManagementClient
	LoggingConfigClient           *logging.ConfigClient
	LoggingMetricsClient          *logging.MetricsClient
	MonitoringAlertPolicyClient   *monitoring.AlertPolicyClient
	RedisCloudRedisClient         *redis.CloudRedisClient
	ResourcemanagerProjectsClient *resourcemanager.ProjectsClient
	ResourcemanagerFoldersClient  *resourcemanager.FoldersClient
	ServiceusageClient            *serviceusage.Client
	SecretmanagerClient           *secretmanager.Client
	SqlService                    *sql.Service
	StorageClient                 *storage.Client
}

const (
	BigQueryService             GcpService = "bigquery.googleapis.com"
	CloudBillingService         GcpService = "cloudbilling.googleapis.com"
	CloudFunctionsService       GcpService = "cloudfunctions.googleapis.com"
	CloudKmsService             GcpService = "cloudkms.googleapis.com"
	CloudResourceManagerService GcpService = "cloudresourcemanager.googleapis.com"
	ComputeService              GcpService = "compute.googleapis.com"
	DnsService                  GcpService = "dns.googleapis.com"
	DomainsService              GcpService = "domains.googleapis.com"
	IamService                  GcpService = "iam.googleapis.com"
	LoggingService              GcpService = "logging.googleapis.com"
	RedisService                GcpService = "redis.googleapis.com"
	MonitoringService           GcpService = "monitoring.googleapis.com"
	SqlAdminService             GcpService = "sqladmin.googleapis.com"
	StorageService              GcpService = "storage-api.googleapis.com"
)

func initServices(ctx context.Context, options []option.ClientOption) (*Services, error) {
	svcs := Services{}
	var err error
	svcs.Bigquery, err = bigquery.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.BillingCloudBillingClient, err = billing.NewCloudBillingClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.FunctionsCloudFunctionsClient, err = functions.NewCloudFunctionsClient(ctx, options...)
	// cloudFunctionsSvc, err := cloudfunctions.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.RunServicesClient, err = run.NewServicesClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.ComputeAddressesClient, err = compute.NewAddressesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeAutoscalersClient, err = compute.NewAutoscalersRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeBackendServicesClient, err = compute.NewBackendServicesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeDiskTypesClient, err = compute.NewDiskTypesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeDisksClient, err = compute.NewDisksRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeFirewallsClient, err = compute.NewFirewallsRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeForwardingRulesClient, err = compute.NewForwardingRulesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeImagesClient, err = compute.NewImagesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	svcs.ComputeInstancesClient, err = compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, err
	}

	svcs.ContainerClusterManagerClient, err = container.NewClusterManagerClient(ctx)
	// containerSvc, err := container.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.DomainsClient, err = domains.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.Dns, err = dns.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.Iam, err = iam.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.KmsKeyManagementClient, err = kms.NewKeyManagementClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.LoggingConfigClient, err = logging.NewConfigClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.LoggingMetricsClient, err = logging.NewMetricsClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.MonitoringAlertPolicyClient, err = monitoring.NewAlertPolicyClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.RedisCloudRedisClient, err = redis.NewCloudRedisClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.ResourcemanagerProjectsClient, err = resourcemanager.NewProjectsClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ResourcemanagerFoldersClient, err = resourcemanager.NewFoldersClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.SecretmanagerClient, err = secretmanager.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ServiceusageClient, err = serviceusage.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.SqlService, err = sql.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.StorageClient, err = storage.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &svcs, nil
}
