package client

import (
	"context"

	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/cloudfunctions/v1"
	kms "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/dns/v1"
	domains "google.golang.org/api/domains/v1beta1"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/logging/v2"
	"google.golang.org/api/monitoring/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/run/v1"
	"google.golang.org/api/secretmanager/v1"
	"google.golang.org/api/serviceusage/v1"
	sql "google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/api/storage/v1"
)

type GcpService string

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
	MonitoringService           GcpService = "monitoring.googleapis.com"
	SqlAdminService             GcpService = "sqladmin.googleapis.com"
	StorageService              GcpService = "storage-api.googleapis.com"
)

type Services struct {
	BigQuery        *bigquery.Service
	CloudBilling    *cloudbilling.APIService
	CloudFunctions  *cloudfunctions.Service
	CloudRun        *run.APIService
	Compute         *compute.Service
	Container       *container.Service
	Dns             *dns.Service
	Domain          *domains.Service
	Iam             *iam.Service
	Kms             *kms.Service
	Logging         *logging.Service
	Monitoring      *monitoring.Service
	ResourceManager *cloudresourcemanager.Service
	ServiceUsage    *serviceusage.Service
	SecretManager   *secretmanager.Service
	Sql             *sql.Service
	Storage         *storage.Service
}

func initServices(ctx context.Context, options []option.ClientOption) (*Services, error) {
	bigQuerySvc, err := bigquery.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	cloudBillingSvc, err := cloudbilling.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	cloudFunctionsSvc, err := cloudfunctions.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	cloudRunSvc, err := run.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	computeSvc, err := compute.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	containerSvc, err := container.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	domainSvc, err := domains.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	dnsSvc, err := dns.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	iamSvc, err := iam.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	kmsSvc, err := kms.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	loggingSvc, err := logging.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	monitoringSvc, err := monitoring.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	resourceManagerSvc, err := cloudresourcemanager.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	secretManagerSvc, err := secretmanager.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	serviceUsageManagerSvc, err := serviceusage.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	sqlSvc, err := sql.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	storageSvc, err := storage.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &Services{
		BigQuery:        bigQuerySvc,
		CloudBilling:    cloudBillingSvc,
		CloudFunctions:  cloudFunctionsSvc,
		CloudRun:        cloudRunSvc,
		Compute:         computeSvc,
		Container:       containerSvc,
		Dns:             dnsSvc,
		Domain:          domainSvc,
		Iam:             iamSvc,
		Kms:             kmsSvc,
		Logging:         loggingSvc,
		Monitoring:      monitoringSvc,
		ResourceManager: resourceManagerSvc,
		SecretManager:   secretManagerSvc,
		ServiceUsage:    serviceUsageManagerSvc,
		Sql:             sqlSvc,
		Storage:         storageSvc,
	}, nil
}
