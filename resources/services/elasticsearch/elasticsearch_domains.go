package elasticsearch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticsearchDomains() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticsearch_domains",
		Description:  "The current status of an Elasticsearch domain.",
		Resolver:     fetchElasticsearchDomains,
		Multiplex:    client.ServiceAccountRegionMultiplexer("es"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
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
				Name:        "id",
				Description: "The unique identifier for the specified Elasticsearch domain.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DomainId"),
			},
			{
				Name:        "name",
				Description: "The name of an Elasticsearch domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DomainName"),
			},
			{
				Name:        "cluster_cold_storage_options_enabled",
				Description: "True to enable cold storage for an Elasticsearch domain.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.ColdStorageOptions.Enabled"),
			},
			{
				Name:        "cluster_dedicated_master_count",
				Description: "Total number of dedicated master nodes, active and on standby, for the cluster.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.DedicatedMasterCount"),
			},
			{
				Name:        "cluster_dedicated_master_enabled",
				Description: "A boolean value to indicate whether a dedicated master node is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.DedicatedMasterEnabled"),
			},
			{
				Name:        "cluster_dedicated_master_type",
				Description: "The instance type for a dedicated master node.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.DedicatedMasterType"),
			},
			{
				Name:        "cluster_instance_count",
				Description: "The number of instances in the specified domain cluster.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.InstanceCount"),
			},
			{
				Name:        "cluster_instance_type",
				Description: "The instance type for an Elasticsearch cluster",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.InstanceType"),
			},
			{
				Name:        "cluster_warm_count",
				Description: "The number of warm nodes in the cluster.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.WarmCount"),
			},
			{
				Name:        "cluster_warm_enabled",
				Description: "True to enable warm storage.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.WarmEnabled"),
			},
			{
				Name:        "cluster_warm_type",
				Description: "The instance type for the Elasticsearch cluster's warm nodes.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.WarmType"),
			},
			{
				Name:        "cluster_zone_awareness_config_availability_zone_count",
				Description: "An integer value to indicate the number of availability zones for a domain when zone awareness is enabled",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.ZoneAwarenessConfig.AvailabilityZoneCount"),
			},
			{
				Name:        "cluster_zone_awareness_enabled",
				Description: "A boolean value to indicate whether zone awareness is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ElasticsearchClusterConfig.ZoneAwarenessEnabled"),
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
				Name:        "advanced_security_enabled",
				Description: "True if advanced security is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.Enabled"),
			},
			{
				Name:        "advanced_security_internal_user_database_enabled",
				Description: "True if the internal user database is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.InternalUserDatabaseEnabled"),
			},
			{
				Name:        "advanced_security_saml_enabled",
				Description: "True if SAML is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.SAMLOptions.Enabled"),
			},
			{
				Name:        "advanced_security_saml_idp_entity_id",
				Description: "The unique Entity ID of the application in SAML Identity Provider.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.SAMLOptions.Idp.EntityId"),
			},
			{
				Name:        "advanced_security_saml_roles_key",
				Description: "The Metadata of the SAML application in XML format.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.SAMLOptions.Idp.MetadataContent"),
			},
			{
				Name:        "advanced_security_options_saml_options_roles_key",
				Description: "The key used for matching the SAML Roles attribute.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.SAMLOptions.RolesKey"),
			},
			{
				Name:        "advanced_security_saml_session_timeout_minutes",
				Description: "The duration, in minutes, after which a user session becomes inactive.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.SAMLOptions.SessionTimeoutMinutes"),
			},
			{
				Name:        "advanced_security_saml_subject_key",
				Description: "The key used for matching the SAML Subject attribute.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AdvancedSecurityOptions.SAMLOptions.SubjectKey"),
			},
			{
				Name:        "auto_tune_error_message",
				Description: "Specifies the error message while enabling or disabling the Auto-Tune.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoTuneOptions.ErrorMessage"),
			},
			{
				Name:        "auto_tune_options_state",
				Description: "Specifies the AutoTuneState for the Elasticsearch domain.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoTuneOptions.State"),
			},
			{
				Name:        "cognito_enabled",
				Description: "Specifies the option to enable Cognito for Kibana authentication.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CognitoOptions.Enabled"),
			},
			{
				Name:        "cognito_identity_pool_id",
				Description: "Specifies the Cognito identity pool ID for Kibana authentication.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CognitoOptions.IdentityPoolId"),
			},
			{
				Name:        "cognito_role_arn",
				Description: "Specifies the role ARN that provides Elasticsearch permissions for accessing Cognito resources.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CognitoOptions.RoleArn"),
			},
			{
				Name:        "cognito_user_pool_id",
				Description: "Specifies the Cognito user pool ID for Kibana authentication.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CognitoOptions.UserPoolId"),
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
				Name:        "domain_endpoint_custom",
				Description: "Specify the fully qualified domain for your custom endpoint.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DomainEndpointOptions.CustomEndpoint"),
			},
			{
				Name:        "domain_endpoint_custom_certificate_arn",
				Description: "Specify ACM certificate ARN for your custom endpoint.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DomainEndpointOptions.CustomEndpointCertificateArn"),
			},
			{
				Name:        "domain_endpoint_custom_enabled",
				Description: "Specify if custom endpoint should be enabled for the Elasticsearch domain.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DomainEndpointOptions.CustomEndpointEnabled"),
			},
			{
				Name:        "domain_endpoint_enforce_https",
				Description: "Specify if only HTTPS endpoint should be enabled for the Elasticsearch domain.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DomainEndpointOptions.EnforceHTTPS"),
			},
			{
				Name:        "domain_endpoint_tls_security_policy",
				Description: "Specify the TLS security policy that needs to be applied to the HTTPS endpoint of Elasticsearch domain.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DomainEndpointOptions.TLSSecurityPolicy"),
			},
			{
				Name:        "ebs_enabled",
				Description: "Specifies whether EBS-based storage is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EBSOptions.EBSEnabled"),
			},
			{
				Name:        "ebs_iops",
				Description: "Specifies the IOPD for a Provisioned IOPS EBS volume (SSD).",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EBSOptions.Iops"),
			},
			{
				Name:        "ebs_volume_size",
				Description: "Integer to specify the size of an EBS volume.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EBSOptions.VolumeSize"),
			},
			{
				Name:        "ebs_volume_type",
				Description: "Specifies the volume type for EBS-based storage.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EBSOptions.VolumeType"),
			},
			{
				Name: "elasticsearch_version",
				Type: schema.TypeString,
			},
			{
				Name:        "encryption_at_rest_enabled",
				Description: "Specifies the option to enable Encryption At Rest.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EncryptionAtRestOptions.Enabled"),
			},
			{
				Name:        "encryption_at_rest_kms_key_id",
				Description: "Specifies the KMS Key ID for Encryption At Rest options.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionAtRestOptions.KmsKeyId"),
			},
			{
				Name:        "endpoint",
				Description: "The Elasticsearch domain endpoint that you use to submit index and search requests.",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoints",
				Description: "Map containing the Elasticsearch domain endpoints used to submit index and search requests",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "log_publishing_options",
				Description: "Log publishing options for the given domain.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "node_to_node_encryption_enabled",
				Description: "Specify true to enable node-to-node encryption.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NodeToNodeEncryptionOptions.Enabled"),
			},
			{
				Name:        "processing",
				Description: "The status of the Elasticsearch domain configuration",
				Type:        schema.TypeBool,
			},
			{
				Name:        "service_software_automated_update_date",
				Description: "Timestamp, in Epoch time, until which you can manually request a service software update",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.AutomatedUpdateDate"),
			},
			{
				Name:        "service_software_cancellable",
				Description: "True if you are able to cancel your service software version update",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.Cancellable"),
			},
			{
				Name:        "service_software_current_version",
				Description: "The current service software version that is present on the domain.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.CurrentVersion"),
			},
			{
				Name:        "service_software_description",
				Description: "The description of the UpdateStatus.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.Description"),
			},
			{
				Name:        "service_software_new_version",
				Description: "The new service software version if one is available.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.NewVersion"),
			},
			{
				Name:        "service_software_optional_deployment",
				Description: "True if a service software is never automatically updated",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.OptionalDeployment"),
			},
			{
				Name:        "service_software_update_available",
				Description: "True if you are able to update you service software version",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.UpdateAvailable"),
			},
			{
				Name:        "service_software_update_status",
				Description: "The status of your service software update",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceSoftwareOptions.UpdateStatus"),
			},
			{
				Name:        "snapshot_options_automated_snapshot_start_hour",
				Description: "Specifies the time, in UTC format, when the service takes a daily automated snapshot of the specified Elasticsearch domain",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SnapshotOptions.AutomatedSnapshotStartHour"),
			},
			{
				Name:        "upgrade_processing",
				Description: "The status of an Elasticsearch domain version upgrade",
				Type:        schema.TypeBool,
			},
			{
				Name:        "vpc_availability_zones",
				Description: "The availability zones for the Elasticsearch domain",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VPCOptions.AvailabilityZones"),
			},
			{
				Name:        "vpc_security_group_ids",
				Description: "Specifies the security groups for VPC endpoint.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VPCOptions.SecurityGroupIds"),
			},
			{
				Name:        "vpc_subnet_ids",
				Description: "Specifies the subnets for VPC endpoint.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VPCOptions.SubnetIds"),
			},
			{
				Name:        "vpc_vpc_id",
				Description: "The VPC Id for the Elasticsearch domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VPCOptions.VPCId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticsearchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	optsFunc := func(options *elasticsearchservice.Options) { options.Region = c.Region }
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
	domain, ok := resource.Item.(*types.ElasticsearchDomainStatus)
	if !ok {
		return fmt.Errorf("expected to have *types.ElasticsearchDomainStatus but got %T", resource.Item)
	}
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
