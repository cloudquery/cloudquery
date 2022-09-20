package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/digitalocean/godo"
)

//go:generate mockgen -package=mocks -destination=./mocks/account_service.go . AccountService
type AccountService interface {
	Get(context.Context) (*godo.Account, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/cdn_service.go . CdnService
type CdnService interface {
	List(context.Context, *godo.ListOptions) ([]godo.CDN, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/billing_history_service.go . BillingHistoryService
type BillingHistoryService interface {
	List(context.Context, *godo.ListOptions) (*godo.BillingHistory, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/monitoring_service.go . MonitoringService
type MonitoringService interface {
	ListAlertPolicies(context.Context, *godo.ListOptions) ([]godo.AlertPolicy, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/balance_service.go . BalanceService
type BalanceService interface {
	Get(context.Context) (*godo.Balance, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/certificates_service.go . CertificatesService
type CertificatesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Certificate, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/databases_service.go . DatabasesService
type DatabasesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Database, *godo.Response, error)
	ListBackups(context.Context, string, *godo.ListOptions) ([]godo.DatabaseBackup, *godo.Response, error)
	ListReplicas(context.Context, string, *godo.ListOptions) ([]godo.DatabaseReplica, *godo.Response, error)
	GetFirewallRules(context.Context, string) ([]godo.DatabaseFirewallRule, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/domains_service.go . DomainsService
type DomainsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error)
	Records(context.Context, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/droplets_service.go . DropletsService
type DropletsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Droplet, *godo.Response, error)
	Neighbors(context.Context, int) ([]godo.Droplet, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/firewalls_service.go . FirewallsService
type FirewallsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Firewall, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/floating_ips_service.go . FloatingIpsService
type FloatingIpsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.FloatingIP, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/images_service.go . ImagesService
type ImagesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Image, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/keys_service.go . KeysService
type KeysService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Key, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/load_balancers_service.go . LoadBalancersService
type LoadBalancersService interface {
	List(context.Context, *godo.ListOptions) ([]godo.LoadBalancer, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/projcets_service.go . ProjectsService
type ProjectsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Project, *godo.Response, error)
	ListResources(context.Context, string, *godo.ListOptions) ([]godo.ProjectResource, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/regions_service.go . RegionsService
type RegionsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Region, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/registry_service.go . RegistryService
type RegistryService interface {
	Get(context.Context) (*godo.Registry, *godo.Response, error)
	ListRepositories(context.Context, string, *godo.ListOptions) ([]*godo.Repository, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/sizes_service.go . SizesService
type SizesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Size, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/snapshots_service.go . SnapshotsService
type SnapshotsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Snapshot, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/storage_service.go . StorageService
type StorageService interface {
	ListVolumes(context.Context, *godo.ListVolumeParams) ([]godo.Volume, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/vpc_service.go . VpcsService
type VpcsService interface {
	List(context.Context, *godo.ListOptions) ([]*godo.VPC, *godo.Response, error)
	ListMembers(context.Context, string, *godo.VPCListMembersRequest, *godo.ListOptions) ([]*godo.VPCMember, *godo.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/spaces_service.go . SpacesService
type SpacesService interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
}
