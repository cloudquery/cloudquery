// Code generated by codegen; DO NOT EDIT.

package elasticsearch

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:                "aws_elasticsearch_domains",
		Resolver:            fetchElasticsearchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer("es"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElasticsearchDomainTags,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "elasticsearch_cluster_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ElasticsearchClusterConfig"),
			},
			{
				Name:     "access_policies",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessPolicies"),
			},
			{
				Name:     "advanced_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdvancedOptions"),
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
				Name:     "change_progress_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ChangeProgressDetails"),
			},
			{
				Name:     "cognito_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CognitoOptions"),
			},
			{
				Name:     "created",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "deleted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Deleted"),
			},
			{
				Name:     "domain_endpoint_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainEndpointOptions"),
			},
			{
				Name:     "ebs_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EBSOptions"),
			},
			{
				Name:     "elasticsearch_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ElasticsearchVersion"),
			},
			{
				Name:     "encryption_at_rest_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionAtRestOptions"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Endpoints"),
			},
			{
				Name:     "log_publishing_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogPublishingOptions"),
			},
			{
				Name:     "node_to_node_encryption_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodeToNodeEncryptionOptions"),
			},
			{
				Name:     "processing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Processing"),
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
				Name:     "upgrade_processing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("UpgradeProcessing"),
			},
			{
				Name:     "vpc_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VPCOptions"),
			},
		},
	}
}
