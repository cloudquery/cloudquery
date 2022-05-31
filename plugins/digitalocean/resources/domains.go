package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_domains",
		Description:  "Domain represents a DigitalOcean domain",
		Resolver:     fetchDomains,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"name"}},
		Columns: []schema.Column{
			{
				Name:        "name",
				Description: "The name of the domain itself. This should follow the standard domain format of domain.TLD. For instance, `example.com` is a valid domain name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ttl",
				Description: "This value is the time to live for the records on this domain, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("TTL"),
			},
			{
				Name:        "zone_file",
				Description: "This attribute contains the complete contents of the zone file for the selected domain. Individual domain record resources should be used to get more granular control over records. However, this attribute can also be used to get information about the SOA record, which is created automatically and is not accessible as an individual record resource.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_domain_records",
				Description: "DomainRecord represents a DigitalOcean DomainRecord",
				Resolver:    fetchDomainRecords,
				Columns: []schema.Column{
					{
						Name:        "domain_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_domains table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "A unique identifier for each domain record.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "type",
						Description: "The type of the DNS record. For example: A, CNAME, TXT, ...",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The host name, alias, or service being defined by the record.",
						Type:        schema.TypeString,
					},
					{
						Name:        "data",
						Description: "Variable data depending on record type. For example, the \"data\" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.",
						Type:        schema.TypeString,
					},
					{
						Name:        "priority",
						Description: "The priority for SRV and MX records.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "port",
						Description: "The port for SRV records.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "ttl",
						Description: "This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("TTL"),
					},
					{
						Name:        "weight",
						Description: "The weight for SRV records.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "flags",
						Description: "An unsigned integer between 0-255 used for CAA records.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "tag",
						Description: "The parameter tag for CAA records. Valid values are \"issue\", \"issuewild\", or \"iodef\"",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		domains, resp, err := svc.DoClient.Domains.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- domains
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
func fetchDomainRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	domain := parent.Item.(godo.Domain)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		records, resp, err := svc.DoClient.Domains.Records(ctx, domain.Name, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- records
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
