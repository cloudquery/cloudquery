package client

import (
	"context"

	apikeys "cloud.google.com/go/apikeys/apiv2"
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
	ApikeysClient                 *apikeys.Client
	BigqueryService               *bigquery.Service
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

func initServices(ctx context.Context, options []option.ClientOption) (*Services, error) {
	options = append(options, option.WithTelemetryDisabled())
	svcs := Services{}
	var err error

	svcs.ApikeysClient, err = apikeys.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.BigqueryService, err = bigquery.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.BillingCloudBillingClient, err = billing.NewCloudBillingClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.BillingCloudCatalogClient, err = billing.NewCloudCatalogClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.FunctionsCloudFunctionsClient, err = functions.NewCloudFunctionsClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.RunServicesClient, err = run.NewServicesClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.ComputeAddressesClient, err = compute.NewAddressesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeAutoscalersClient, err = compute.NewAutoscalersRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeBackendServicesClient, err = compute.NewBackendServicesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeDiskTypesClient, err = compute.NewDiskTypesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeDisksClient, err = compute.NewDisksRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeFirewallsClient, err = compute.NewFirewallsRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeForwardingRulesClient, err = compute.NewForwardingRulesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeImagesClient, err = compute.NewImagesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeInstancesClient, err = compute.NewInstancesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeInstanceGroupsClient, err = compute.NewInstanceGroupsRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeInterconnectsClient, err = compute.NewInterconnectsRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeNetworksClient, err = compute.NewNetworksRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeProjectsClient, err = compute.NewProjectsRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeSslCertificatesClient, err = compute.NewSslCertificatesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeSslPoliciesClient, err = compute.NewSslPoliciesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeSubnetworksClient, err = compute.NewSubnetworksRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeTargetHttpProxiesClient, err = compute.NewTargetHttpProxiesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeTargetSslProxiesClient, err = compute.NewTargetSslProxiesRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeUrlMapsClient, err = compute.NewUrlMapsRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	svcs.ComputeVpnGatewaysClient, err = compute.NewVpnGatewaysRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	svcs.ContainerClusterManagerClient, err = container.NewClusterManagerClient(ctx, options...)
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
