package client

import (
	"context"

	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/cloudfunctions/v1"
	kms "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/dns/v1"
	domains "google.golang.org/api/domains/v1beta1"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/logging/v2"
	"google.golang.org/api/monitoring/v3"
	sql "google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/api/storage/v1"
)

type Services struct {
	Kms             *kms.Service
	Storage         *storage.Service
	Sql             *sql.Service
	Iam             *iam.Service
	Crm             *cloudresourcemanager.Service
	CloudFunctions  *cloudfunctions.Service
	Domain          *domains.Service
	Compute         *compute.Service
	BigQuery        *bigquery.Service
	Dns             *dns.Service
	Logging         *logging.Service
	Monitoring      *monitoring.Service
	ResourceManager *cloudresourcemanager.Service
}

func initServices(ctx context.Context) (*Services, error) {

	kmsSvc, err := kms.NewService(ctx)
	if err != nil {
		return nil, err
	}
	storageSvc, err := storage.NewService(ctx)
	if err != nil {
		return nil, err
	}
	sqlSvc, err := sql.NewService(ctx)
	if err != nil {
		return nil, err
	}
	iamSvc, err := iam.NewService(ctx)
	if err != nil {
		return nil, err
	}
	crmSvc, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return nil, err
	}
	cfSvc, err := cloudfunctions.NewService(ctx)
	if err != nil {
		return nil, err
	}
	domainSvc, err := domains.NewService(ctx)
	if err != nil {
		return nil, err
	}
	computeSvc, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}
	bigquerySvc, err := bigquery.NewService(ctx)
	if err != nil {
		return nil, err
	}
	dnsSvc, err := dns.NewService(ctx)
	if err != nil {
		return nil, err
	}
	loggingSvc, err := logging.NewService(ctx)
	if err != nil {
		return nil, err
	}
	monitoringSvc, err := monitoring.NewService(ctx)
	if err != nil {
		return nil, err
	}
	resourceManagerSvc, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Services{
		Kms:             kmsSvc,
		Storage:         storageSvc,
		Sql:             sqlSvc,
		Iam:             iamSvc,
		Crm:             crmSvc,
		CloudFunctions:  cfSvc,
		Domain:          domainSvc,
		Compute:         computeSvc,
		BigQuery:        bigquerySvc,
		Dns:             dnsSvc,
		Logging:         loggingSvc,
		Monitoring:      monitoringSvc,
		ResourceManager: resourceManagerSvc,
	}, nil
}
