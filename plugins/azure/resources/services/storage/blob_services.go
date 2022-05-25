package storage

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func StorageBlobServices() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_blob_services",
		Description: "Azure storage blob service",
		Resolver:    fetchStorageBlobServices,
		Columns: []schema.Column{
			{
				Name:        "account_cq_id",
				Description: "Unique CloudQuery ID of azure_storage_accounts table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:          "default_service_version",
				Description:   "DefaultServiceVersion indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.DefaultServiceVersion"),
				IgnoreInTests: true,
			},
			{
				Name:        "delete_retention_policy_enabled",
				Description: "Indicates whether DeleteRetentionPolicy is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BlobServicePropertiesProperties.DeleteRetentionPolicy.Enabled"),
			},
			{
				Name:          "delete_retention_policy_days",
				Description:   "Indicates the number of days that the deleted item should be retained.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.DeleteRetentionPolicy.Days"),
				IgnoreInTests: true,
			},
			{
				Name:        "is_versioning_enabled",
				Description: "Versioning is enabled if set to true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BlobServicePropertiesProperties.IsVersioningEnabled"),
			},
			{
				Name:          "automatic_snapshot_policy_enabled",
				Description:   "Deprecated in favor of isVersioningEnabled property",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.AutomaticSnapshotPolicyEnabled"),
				IgnoreInTests: true,
			},
			{
				Name:        "change_feed_enabled",
				Description: "Indicates whether change feed event logging is enabled for the Blob service",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BlobServicePropertiesProperties.ChangeFeed.Enabled"),
			},
			{
				Name:          "change_feed_retention_in_days",
				Description:   "Indicates the duration of changeFeed retention in days.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.ChangeFeed.RetentionInDays"),
				IgnoreInTests: true,
			},
			{
				Name:          "restore_policy_enabled",
				Description:   "Blob restore is enabled if set to true",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.RestorePolicy.Enabled"),
				IgnoreInTests: true,
			},
			{
				Name:          "restore_policy_days",
				Description:   "how long this blob can be restored It should be great than zero and less than DeleteRetentionPolicydays",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.RestorePolicy.Days"),
				IgnoreInTests: true,
			},
			{
				Name:     "restore_policy_last_enabled_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("BlobServicePropertiesProperties.RestorePolicy.LastEnabledTime.Time"),
			},
			{
				Name:     "restore_policy_min_restore_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("BlobServicePropertiesProperties.RestorePolicy.MinRestoreTime.Time"),
			},
			{
				Name:          "container_delete_retention_policy_enabled",
				Description:   "Indicates whether DeleteRetentionPolicy is enabled",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.ContainerDeleteRetentionPolicy.Enabled"),
				IgnoreInTests: true,
			},
			{
				Name:          "container_delete_retention_policy_days",
				Description:   "Indicates the number of days that the deleted item should be retained.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.ContainerDeleteRetentionPolicy.Days"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_access_time_tracking_policy_enable",
				Description:   "When set to true last access time based tracking is enabled",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.LastAccessTimeTrackingPolicy.Enable"),
				IgnoreInTests: true,
			},
			{
				Name:        "last_access_time_tracking_policy_name",
				Description: "Name of the policy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BlobServicePropertiesProperties.LastAccessTimeTrackingPolicy.Name"),
			},
			{
				Name:          "last_access_time_tracking_policy_tracking_granularity_in_days",
				Description:   "The field specifies blob object tracking granularity in days.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.LastAccessTimeTrackingPolicy.TrackingGranularityInDays"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_access_time_tracking_policy_blob_type",
				Description:   "An array of predefined supported blob types.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("BlobServicePropertiesProperties.LastAccessTimeTrackingPolicy.BlobType"),
				IgnoreInTests: true,
			},
			{
				Name:        "sku_name",
				Description: "Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Possible values include: 'Standard', 'Premium'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_storage_blob_service_cors_rules",
				Description:   "CorsRule specifies a CORS rule for the Blob service",
				Resolver:      fetchStorageBlobServiceCorsRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "blob_service_cq_id",
						Description: "Unique CloudQuery ID of azure_storage_blob_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "blob_service_id",
						Description: "Fully qualified resource ID of blob service",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "allowed_origins",
						Description: "A list of origin domains that will be allowed via CORS, or \"*\" to allow all domains",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_methods",
						Description: "A list of HTTP methods that are allowed to be executed by the origin",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "max_age_in_seconds",
						Description: "The number of seconds that the client/browser should cache a preflight response",
						Type:        schema.TypeInt,
					},
					{
						Name:        "exposed_headers",
						Description: "A list of response headers to expose to CORS clients",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_headers",
						Description: "A list of headers allowed to be part of the cross-origin request",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageBlobServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Storage.BlobServices
	account, ok := parent.Item.(storage.Account)
	if !ok {
		return fmt.Errorf("not a storage.Account: %T", parent.Item)
	}

	if !isBlobSupported(&account) {
		return nil
	}

	resource, err := client.ParseResourceID(*account.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.List(ctx, resource.ResourceGroup, *account.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	if result.Value == nil {
		return nil
	}
	res <- *result.Value
	return nil
}

func fetchStorageBlobServiceCorsRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	blob, ok := parent.Item.(storage.BlobServiceProperties)
	if !ok {
		return fmt.Errorf("not a storage.BlobServiceProperties: %T", parent.Item)
	}
	if blob.BlobServicePropertiesProperties == nil ||
		blob.BlobServicePropertiesProperties.Cors == nil ||
		blob.BlobServicePropertiesProperties.Cors.CorsRules == nil {
		return nil
	}
	res <- *blob.BlobServicePropertiesProperties.Cors.CorsRules
	return nil
}
