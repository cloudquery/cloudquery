package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ElasticsearchDomains() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticsearch_domains",
		Description: "The current status of an Elasticsearch domain.",
		Resolver:    fetchElasticsearchDomains,
		Multiplex:   client.ServiceAccountRegionMultiplexer("es"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElasticsearchDomainTags,
			},
			{
				Name:        "arn",
				Description: "The Amazon resource name (ARN) of an Elasticsearch domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:            "id",
				Description:     "The unique identifier for the specified Elasticsearch domain. ",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("DomainId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The name of an Elasticsearch domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DomainName"),
			},
			{
				Name:     "elasticsearch_cluster_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ElasticsearchClusterConfig"),
			},
			{
				Name:        "access_policies",
				Description: "IAM access policy as a JSON-formatted string.",
				Type:        schema.TypeString,
			},
			{
				Name:        "advanced_options",
				Description: "Specifies the status of the AdvancedOptions",
				Type:        schema.TypeJSON,
			},
			{
				Name:     "advanced_security_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdvancedSecurityOptions"),
			},
			{
				Name:     "auto_tune_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoTuneOptions"),
			},
			{
				Name:     "cognito_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CognitoOptions"),
			},
			{
				Name:        "created",
				Description: "The domain creation status",
				Type:        schema.TypeBool,
			},
			{
				Name:        "deleted",
				Description: "The domain deletion status",
				Type:        schema.TypeBool,
			},
			{
				Name:          "domain_endpoint_options",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("DomainEndpointOptions"),
				IgnoreInTests: true,
			},
			{
				Name:     "ebs_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EBSOptions"),
			},
			{
				Name: "elasticsearch_version",
				Type: schema.TypeString,
			},
			{
				Name:     "encryption_at_rest_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionAtRestOptions"),
			},
			{
				Name:        "endpoint",
				Description: "The Elasticsearch domain endpoint that you use to submit index and search requests.",
				Type:        schema.TypeString,
			},
			{
				Name:          "endpoints",
				Description:   "Map containing the Elasticsearch domain endpoints used to submit index and search requests",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "log_publishing_options",
				Description:   "Log publishing options for the given domain.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:     "node_to_node_encryption_options",
				Type:     schema.TypeJSON,
			},
			{
				Name:        "processing",
				Description: "The status of the Elasticsearch domain configuration",
				Type:        schema.TypeBool,
			},
			{
				Name:     "service_software_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceSoftwareOptions"),
			},
			{
				Name:     "snapshot_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SnapshotOptions"),
			},
			{
				Name:        "upgrade_processing",
				Description: "The status of an Elasticsearch domain version upgrade",
				Type:        schema.TypeBool,
			},
			{
				Name:          "vpc_options",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("VPCOptions"),
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchElasticsearchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	optsFunc := func(options *elasticsearchservice.Options) {
		options.Region = c.Region
	}
	svc := c.Services().ElasticSearch
	out, err := svc.ListDomainNames(ctx, &elasticsearchservice.ListDomainNamesInput{}, optsFunc)
	if err != nil {
		return err
	}
	for _, info := range out.DomainNames {
		domainOutput, err := svc.DescribeElasticsearchDomain(ctx, &elasticsearchservice.DescribeElasticsearchDomainInput{DomainName: info.DomainName}, optsFunc)
		if err != nil {
			return nil
		}
		res <- domainOutput.DomainStatus
	}
	return nil
}
func resolveElasticsearchDomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ElasticSearch
	domain := resource.Item.(*types.ElasticsearchDomainStatus)
	tagsOutput, err := svc.ListTags(ctx, &elasticsearchservice.ListTagsInput{
		ARN: domain.ARN,
	}, func(o *elasticsearchservice.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}
	if len(tagsOutput.TagList) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.TagList {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
