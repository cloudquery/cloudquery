package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func RdsInstances() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_instances",
		Resolver:     fetchRdsInstances,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "allocated_storage",
				Type: schema.TypeInt,
			},
			{
				Name: "auto_minor_version_upgrade",
				Type: schema.TypeBool,
			},
			{
				Name: "availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "aws_backup_recovery_point_arn",
				Type: schema.TypeString,
			},
			{
				Name: "backup_retention_period",
				Type: schema.TypeInt,
			},
			{
				Name:     "ca_certificate_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CACertificateIdentifier"),
			},
			{
				Name: "character_set_name",
				Type: schema.TypeString,
			},
			{
				Name: "copy_tags_to_snapshot",
				Type: schema.TypeBool,
			},
			{
				Name: "customer_owned_ip_enabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceArn"),
			},
			{
				Name:     "db_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceClass"),
			},
			{
				Name:     "db_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceIdentifier"),
			},
			{
				Name:     "db_instance_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceStatus"),
			},
			{
				Name:     "db_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBName"),
			},
			{
				Name:     "db_subnet_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup.DBSubnetGroupArn"),
			},
			{
				Name:     "db_subnet_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup.DBSubnetGroupDescription"),
			},
			{
				Name:     "db_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup.DBSubnetGroupName"),
			},
			{
				Name:     "db_subnet_group_subnet_group_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup.SubnetGroupStatus"),
			},
			{
				Name:     "db_subnet_group_vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup.VpcId"),
			},
			{
				Name: "db_instance_port",
				Type: schema.TypeInt,
			},
			{
				Name: "dbi_resource_id",
				Type: schema.TypeString,
			},
			{
				Name: "deletion_protection",
				Type: schema.TypeBool,
			},
			{
				Name: "enabled_cloudwatch_logs_exports",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "endpoint_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint.Address"),
			},
			{
				Name:     "endpoint_hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint.HostedZoneId"),
			},
			{
				Name:     "endpoint_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Endpoint.Port"),
			},
			{
				Name: "engine",
				Type: schema.TypeString,
			},
			{
				Name: "engine_version",
				Type: schema.TypeString,
			},
			{
				Name: "enhanced_monitoring_resource_arn",
				Type: schema.TypeString,
			},
			{
				Name:     "iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name: "instance_create_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "iops",
				Type: schema.TypeInt,
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "latest_restorable_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "license_model",
				Type: schema.TypeString,
			},
			{
				Name:     "listener_endpoint_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ListenerEndpoint.Address"),
			},
			{
				Name:     "listener_endpoint_hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ListenerEndpoint.HostedZoneId"),
			},
			{
				Name:     "listener_endpoint_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ListenerEndpoint.Port"),
			},
			{
				Name: "master_username",
				Type: schema.TypeString,
			},
			{
				Name: "max_allocated_storage",
				Type: schema.TypeInt,
			},
			{
				Name: "monitoring_interval",
				Type: schema.TypeInt,
			},
			{
				Name: "monitoring_role_arn",
				Type: schema.TypeString,
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name: "nchar_character_set_name",
				Type: schema.TypeString,
			},
			{
				Name:     "pending_modified_values_allocated_storage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingModifiedValues.AllocatedStorage"),
			},
			{
				Name:     "pending_modified_values_backup_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingModifiedValues.BackupRetentionPeriod"),
			},
			{
				Name:     "pending_modified_values_ca_certificate_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.CACertificateIdentifier"),
			},
			{
				Name:     "pending_modified_values_db_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.DBInstanceClass"),
			},
			{
				Name:     "pending_modified_values_db_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.DBInstanceIdentifier"),
			},
			{
				Name:     "pending_modified_values_db_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.DBSubnetGroupName"),
			},
			{
				Name:     "pending_modified_values_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.EngineVersion"),
			},
			{
				Name:     "pending_modified_values_iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PendingModifiedValues.IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:     "pending_modified_values_iops",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingModifiedValues.Iops"),
			},
			{
				Name:     "pending_modified_values_license_model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.LicenseModel"),
			},
			{
				Name:     "pending_modified_values_master_user_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.MasterUserPassword"),
			},
			{
				Name:     "pending_modified_values_multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PendingModifiedValues.MultiAZ"),
			},
			{
				Name:     "pending_cloudwatch_logs_types_to_disable",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable"),
			},
			{
				Name:     "pending_cloudwatch_logs_types_to_enable",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable"),
			},
			{
				Name:     "pending_modified_values_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingModifiedValues.Port"),
			},
			{
				Name:     "pending_modified_values_storage_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PendingModifiedValues.StorageType"),
			},
			{
				Name: "performance_insights_enabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "performance_insights_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PerformanceInsightsKMSKeyId"),
			},
			{
				Name: "performance_insights_retention_period",
				Type: schema.TypeInt,
			},
			{
				Name: "preferred_backup_window",
				Type: schema.TypeString,
			},
			{
				Name: "preferred_maintenance_window",
				Type: schema.TypeString,
			},
			{
				Name: "promotion_tier",
				Type: schema.TypeInt,
			},
			{
				Name: "publicly_accessible",
				Type: schema.TypeBool,
			},
			{
				Name:     "read_replica_db_cluster_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReadReplicaDBClusterIdentifiers"),
			},
			{
				Name:     "read_replica_db_instance_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReadReplicaDBInstanceIdentifiers"),
			},
			{
				Name:     "read_replica_source_db_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReadReplicaSourceDBInstanceIdentifier"),
			},
			{
				Name: "replica_mode",
				Type: schema.TypeString,
			},
			{
				Name: "secondary_availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "storage_encrypted",
				Type: schema.TypeBool,
			},
			{
				Name: "storage_type",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstanceTags,
			},
			{
				Name: "tde_credential_arn",
				Type: schema.TypeString,
			},
			{
				Name: "timezone",
				Type: schema.TypeString,
			},
			{
				Name:     "aws_rds_instance_pending_modified_values_processor_features",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstancePendingModifiedValuesProcessorFeatures,
			},
			{
				Name:     "aws_rds_instance_processor_features",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsInstanceProcessorFeatures,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_rds_instance_associated_roles",
				Resolver: fetchRdsInstanceAssociatedRoles,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "feature_name",
						Type: schema.TypeString,
					},
					{
						Name: "role_arn",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_db_instance_automated_backups_replications",
				Resolver: fetchRdsInstanceDbInstanceAutomatedBackupsReplications,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "db_instance_automated_backups_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DBInstanceAutomatedBackupsArn"),
					},
				},
			},
			{
				Name:     "aws_rds_instance_db_parameter_groups",
				Resolver: fetchRdsInstanceDbParameterGroups,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "db_parameter_group_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DBParameterGroupName"),
					},
					{
						Name: "parameter_apply_status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_db_security_groups",
				Resolver: fetchRdsInstanceDbSecurityGroups,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "db_security_group_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DBSecurityGroupName"),
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_db_subnet_group_subnets",
				Resolver: fetchRdsInstanceDbSubnetGroupSubnets,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "subnet_availability_zone_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name: "subnet_identifier",
						Type: schema.TypeString,
					},
					{
						Name:     "subnet_outpost_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetOutpost.Arn"),
					},
					{
						Name: "subnet_status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_domain_memberships",
				Resolver: fetchRdsInstanceDomainMemberships,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "domain",
						Type: schema.TypeString,
					},
					{
						Name:     "fqdn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("FQDN"),
					},
					{
						Name:     "iam_role_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("IAMRoleName"),
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_option_group_memberships",
				Resolver: fetchRdsInstanceOptionGroupMemberships,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "option_group_name",
						Type: schema.TypeString,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_status_infos",
				Resolver: fetchRdsInstanceStatusInfos,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "message",
						Type: schema.TypeString,
					},
					{
						Name: "normal",
						Type: schema.TypeBool,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "status_type",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_rds_instance_vpc_security_groups",
				Resolver: fetchRdsInstanceVpcSecurityGroups,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_security_group_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config rds.DescribeDBInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeDBInstances(ctx, &config, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.DBInstances
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveRdsInstanceTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	tags := map[string]*string{}
	for _, t := range r.TagList {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func fetchRdsInstanceAssociatedRoles(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.AssociatedRoles
	return nil
}
func fetchRdsInstanceDbInstanceAutomatedBackupsReplications(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBInstanceAutomatedBackupsReplications
	return nil
}
func fetchRdsInstanceDbParameterGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBParameterGroups
	return nil
}
func fetchRdsInstanceDbSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBSecurityGroups
	return nil
}
func fetchRdsInstanceDbSubnetGroupSubnets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DBSubnetGroup.Subnets
	return nil
}
func fetchRdsInstanceDomainMemberships(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.DomainMemberships
	return nil
}
func fetchRdsInstanceOptionGroupMemberships(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.OptionGroupMemberships
	return nil
}
func resolveRdsInstancePendingModifiedValuesProcessorFeatures(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	pendingProcessorFeatures := map[string]*string{}
	for _, t := range r.PendingModifiedValues.ProcessorFeatures {
		pendingProcessorFeatures[*t.Name] = t.Value
	}
	resource.Set("aws_rds_instance_pending_modified_values_processor_features", pendingProcessorFeatures)
	return nil
}
func resolveRdsInstanceProcessorFeatures(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	processorFeatures := map[string]*string{}
	for _, t := range r.ProcessorFeatures {
		processorFeatures[*t.Name] = t.Value
	}
	resource.Set("aws_rds_instance_processor_features", processorFeatures)
	return nil
}
func fetchRdsInstanceStatusInfos(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.StatusInfos
	return nil
}
func fetchRdsInstanceVpcSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.DBInstance)
	if !ok {
		return fmt.Errorf("not instance")
	}
	res <- instance.VpcSecurityGroups
	return nil
}
