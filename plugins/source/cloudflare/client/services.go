package client

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_client.go . Api
type Api interface {
	Accounts(ctx context.Context, params cloudflare.AccountsListParams) ([]cloudflare.Account, cloudflare.ResultInfo, error)
	AccountMembers(ctx context.Context, accountID string, pageOpts cloudflare.PaginationOptions) ([]cloudflare.AccountMember, cloudflare.ResultInfo, error)

	ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (r cloudflare.ZonesResponse, err error)

	DNSRecords(ctx context.Context, zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)

	ListWAFPackages(ctx context.Context, zoneID string) ([]cloudflare.WAFPackage, error)
	ListWAFGroups(ctx context.Context, zoneID, packageID string) ([]cloudflare.WAFGroup, error)
	ListWAFRules(ctx context.Context, zoneID, packageID string) ([]cloudflare.WAFRule, error)

	ListWorkerScripts(ctx context.Context) (cloudflare.WorkerListResponse, error)
	ListWorkerRoutes(ctx context.Context, zoneID string) (cloudflare.WorkerRoutesResponse, error)
	ListWorkerCronTriggers(ctx context.Context, accountID, scriptName string) ([]cloudflare.WorkerCronTrigger, error)
	ListWorkersSecrets(ctx context.Context, script string) (cloudflare.WorkersListSecretsResponse, error)

	ListCertificatePacks(ctx context.Context, zoneID string) ([]cloudflare.CertificatePack, error)

	ZoneLevelAccessGroups(ctx context.Context, zoneID string, pageOpts cloudflare.PaginationOptions) ([]cloudflare.AccessGroup, cloudflare.ResultInfo, error)
}
