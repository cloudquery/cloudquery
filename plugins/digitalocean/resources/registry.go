package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_registry",
		Description:  "Registry represents a registry.",
		Resolver:     fetchRegistries,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"name"}},
		Columns: []schema.Column{
			{
				Name:        "name",
				Description: "A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.",
				Type:        schema.TypeString,
			},
			{
				Name:        "storage_usage_bytes",
				Description: "The amount of storage used in the registry in bytes.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "storage_usage_bytes_updated_at",
				Description: "The time at which the storage usage was updated. Storage usage is calculated asynchronously, and may not immediately reflect pushes to the registry.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "created_at",
				Description: "A time value given in ISO8601 combined date and time format that represents when the registry was created.",
				Type:        schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "digitalocean_registry_repositories",
				Description:   "Repository represents a repository",
				Resolver:      fetchRegistryRepositories,
				DeleteFilter:  client.DeleteFilter,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "registry_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_registry table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "registry_name",
						Description: "The name of the container registry.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the repository.",
						Type:        schema.TypeString,
					},
					{
						Name:        "latest_tag_registry_name",
						Description: "The name of the container registry.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LatestTag.RegistryName"),
					},
					{
						Name:        "latest_tag_repository",
						Description: "The name of the repository.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LatestTag.Repository"),
					},
					{
						Name:        "latest_tag",
						Description: "The name of the tag.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LatestTag.Tag"),
					},
					{
						Name:        "latest_tag_manifest_digest",
						Description: "The digest of the manifest associated with the tag.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LatestTag.ManifestDigest"),
					},
					{
						Name:        "latest_tag_compressed_size_bytes",
						Description: "The compressed size of the tag in bytes.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LatestTag.CompressedSizeBytes"),
					},
					{
						Name:        "latest_tag_size_bytes",
						Description: "The uncompressed size of the tag in bytes (this size is calculated asynchronously so it may not be immediately available).",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LatestTag.SizeBytes"),
					},
					{
						Name:        "latest_tag_updated_at",
						Description: "The time the tag was last updated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("LatestTag.UpdatedAt"),
					},
					{
						Name:        "tag_count",
						Description: "The number of tags in the repository.",
						Type:        schema.TypeBigInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	registry, _, err := svc.DoClient.Registry.Get(ctx)
	if err != nil {
		return err
	}
	res <- registry
	return nil
}
func fetchRegistryRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	registry := parent.Item.(*godo.Registry)

	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		certs, resp, err := svc.DoClient.Registry.ListRepositories(ctx, registry.Name, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- certs
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
