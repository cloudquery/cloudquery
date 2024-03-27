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

	ListDNSRecords(ctx context.Context, rc *cloudflare.ResourceContainer, rr cloudflare.ListDNSRecordsParams) ([]cloudflare.DNSRecord, *cloudflare.ResultInfo, error)

	ListRulesets(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListRulesetsParams) ([]cloudflare.Ruleset, error)

	ListWAFPackages(ctx context.Context, zoneID string) ([]cloudflare.WAFPackage, error)
	ListWAFGroups(ctx context.Context, zoneID, packageID string) ([]cloudflare.WAFGroup, error)
	ListWAFRules(ctx context.Context, zoneID, packageID string) ([]cloudflare.WAFRule, error)

	ListWorkers(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListWorkersParams) (cloudflare.WorkerListResponse, *cloudflare.ResultInfo, error)
	ListWorkerRoutes(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListWorkerRoutesParams) (cloudflare.WorkerRoutesResponse, error)
	ListWorkerCronTriggers(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListWorkerCronTriggersParams) ([]cloudflare.WorkerCronTrigger, error)
	ListWorkersSecrets(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListWorkersSecretsParams) (cloudflare.WorkersListSecretsResponse, error)

	ListCertificatePacks(ctx context.Context, zoneID string) ([]cloudflare.CertificatePack, error)

	ListAccessGroups(ctx context.Context, rc *cloudflare.ResourceContainer, pageOpts cloudflare.ListAccessGroupsParams) ([]cloudflare.AccessGroup, *cloudflare.ResultInfo, error)
	ListAccessApplications(ctx context.Context, rc *cloudflare.ResourceContainer, params cloudflare.ListAccessApplicationsParams) ([]cloudflare.AccessApplication, *cloudflare.ResultInfo, error)

	ListWAFOverrides(ctx context.Context, zoneID string) ([]cloudflare.WAFOverride, error)
	ListImages(ctx context.Context, rc *cloudflare.ResourceContainer, pageOpts cloudflare.ListImagesParams) ([]cloudflare.Image, error)
}
