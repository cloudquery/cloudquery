package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource filesystems --config filesystems.hcl --output .
func Filesystems() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_filesystems",
		Description:  "A description of a specific Amazon FSx file system",
		Resolver:     fetchFsxFilesystems,
		Multiplex:    client.ServiceAccountRegionMultiplexer("fsx"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "creation_time",
				Description: "The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "dns_name",
				Description: "The Domain Name System (DNS) name for the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DNSName"),
			},
			{
				Name:        "failure_details_message",
				Description: "A message describing any failures that occurred during file system creation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDetails.Message"),
			},
			{
				Name:        "id",
				Description: "The system-generated, unique 17-digit ID of the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemId"),
			},
			{
				Name:        "type",
				Description: "The type of Amazon FSx file system, which can be LUSTRE, WINDOWS, ONTAP, or OPENZFS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemType"),
			},
			{
				Name:        "version",
				Description: "The Lustre version of the Amazon FSx for Lustre file system, either 2.10 or 2.12",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemTypeVersion"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the Key Management Service (KMS) key used to encrypt Amazon FSx file system data",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the file system",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_interface_ids",
				Description: "The IDs of the elastic network interfaces from which a specific file system is accessible",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "owner_id",
				Description: "The Amazon Web Services account that created the file system",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the file system resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceARN"),
			},
			{
				Name:        "storage_capacity",
				Description: "The storage capacity of the file system in gibibytes (GiB)",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "storage_type",
				Description: "The type of storage the file system is using",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "Specifies the IDs of the subnets that the file system is accessible from",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "The tags to associate with the file system",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the primary virtual private cloud (VPC) for the file system",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_fsx_filesystem_lustre_configuration",
				Description: "The configuration for the Amazon FSx for Lustre file system",
				Resolver:    schema.PathTableResolver("LustreConfiguration"),
				Columns: []schema.Column{
					{
						Name:        "filesystem_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_filesystems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "copy_tags_to_backups",
						Description: "A boolean flag indicating whether tags on the file system are copied to backups",
						Type:        schema.TypeBool,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "A recurring daily time, in the format HH:MM",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_compression_type",
						Description: "The data compression configuration for the file system",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_repo_cfg_auto_import_policy",
						Description: "Describes the file system's linked S3 data repository's AutoImportPolicy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataRepositoryConfiguration.AutoImportPolicy"),
					},
					{
						Name:        "data_repo_cfg_export_path",
						Description: "The export path to the Amazon S3 bucket (and prefix) that you are using to store new and changed Lustre file system files in S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataRepositoryConfiguration.ExportPath"),
					},
					{
						Name:        "data_repo_cfg_failure_details_message",
						Description: "A detailed error message",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataRepositoryConfiguration.FailureDetails.Message"),
					},
					{
						Name:        "data_repo_cfg_import_path",
						Description: "The import path to the Amazon S3 bucket (and optional prefix) that you're using as the data repository for your FSx for Lustre file system, for example s3://import-bucket/optional-prefix",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataRepositoryConfiguration.ImportPath"),
					},
					{
						Name:        "data_repo_cfg_imported_file_chunk_size",
						Description: "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataRepositoryConfiguration.ImportedFileChunkSize"),
					},
					{
						Name:        "data_repo_cfg_lifecycle",
						Description: "Describes the state of the file system's S3 durable data repository, if it is configured with an S3 repository",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataRepositoryConfiguration.Lifecycle"),
					},
					{
						Name:        "deployment_type",
						Description: "The deployment type of the FSx for Lustre file system",
						Type:        schema.TypeString,
					},
					{
						Name:        "drive_cache_type",
						Description: "The type of drive cache used by PERSISTENT_1 file systems that are provisioned with HDD storage devices",
						Type:        schema.TypeString,
					},
					{
						Name:        "log_configuration_level",
						Description: "The data repository events that are logged by Amazon FSx",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LogConfiguration.Level"),
					},
					{
						Name:        "log_configuration_destination",
						Description: "The Amazon Resource Name (ARN) that specifies the destination of the logs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LogConfiguration.Destination"),
					},
					{
						Name:        "mount_name",
						Description: "You use the MountName value when mounting the file system",
						Type:        schema.TypeString,
					},
					{
						Name:        "per_unit_storage_throughput",
						Description: "Per unit storage throughput represents the megabytes per second of read or write throughput per 1 tebibyte of storage provisioned",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "root_squash_configuration_no_squash_nids",
						Description: "When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("RootSquashConfiguration.NoSquashNids"),
					},
					{
						Name:        "root_squash_configuration_root_squash",
						Description: "You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RootSquashConfiguration.RootSquash"),
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_fsx_filesystem_ontap_configuration",
				Description: "Configuration for the FSx for NetApp ONTAP file system",
				Resolver:    schema.PathTableResolver("OntapConfiguration"),
				Columns: []schema.Column{
					{
						Name:        "filesystem_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_filesystems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "A recurring daily time, in the format HH:MM",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_type",
						Description: "Specifies the FSx for ONTAP file system deployment type in use in the file system",
						Type:        schema.TypeString,
					},
					{
						Name:        "disk_iops_configuration_iops",
						Description: "The total number of SSD IOPS provisioned for the file system",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Iops"),
					},
					{
						Name:        "disk_iops_configuration_mode",
						Description: "Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Mode"),
					},
					{
						Name:        "endpoint_ip_address_range",
						Description: "(Multi-AZ only) The IP address range in which the endpoints to access your file system are created",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoints_intercluster_dns_name",
						Description: "The Domain Name Service (DNS) name for the file system",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoints.Intercluster.DNSName"),
					},
					{
						Name:        "endpoints_intercluster_ip_addresses",
						Description: "IP addresses of the file system endpoint",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Endpoints.Intercluster.IpAddresses"),
					},
					{
						Name:        "endpoints_management_dns_name",
						Description: "The Domain Name Service (DNS) name for the file system",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Endpoints.Management.DNSName"),
					},
					{
						Name:        "endpoints_management_ip_addresses",
						Description: "IP addresses of the file system endpoint",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Endpoints.Management.IpAddresses"),
					},
					{
						Name:        "preferred_subnet_id",
						Description: "The ID for a subnet",
						Type:        schema.TypeString,
					},
					{
						Name:        "route_table_ids",
						Description: "(Multi-AZ only) The VPC route tables in which your file system's endpoints are created",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "throughput_capacity",
						Description: "The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps)",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "A recurring weekly time, in the format D:HH:MM",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_fsx_filesystem_open_zfs_configuration",
				Description: "The configuration for the Amazon FSx for OpenZFS file system",
				Resolver:    schema.PathTableResolver("OpenZFSConfiguration"),
				Columns: []schema.Column{
					{
						Name:        "filesystem_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_filesystems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "copy_tags_to_backups",
						Description: "A Boolean value indicating whether tags on the file system should be copied to backups",
						Type:        schema.TypeBool,
					},
					{
						Name:        "copy_tags_to_volumes",
						Description: "A Boolean value indicating whether tags for the volume should be copied to snapshots",
						Type:        schema.TypeBool,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "A recurring daily time, in the format HH:MM",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_type",
						Description: "Specifies the file-system deployment type",
						Type:        schema.TypeString,
					},
					{
						Name:        "disk_iops_configuration_iops",
						Description: "The total number of SSD IOPS provisioned for the file system",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Iops"),
					},
					{
						Name:        "disk_iops_configuration_mode",
						Description: "Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskIopsConfiguration.Mode"),
					},
					{
						Name:        "root_volume_id",
						Description: "The ID of the root volume of the OpenZFS file system",
						Type:        schema.TypeString,
					},
					{
						Name:        "throughput_capacity",
						Description: "The throughput of an Amazon FSx file system, measured in megabytes per second (MBps)",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "A recurring weekly time, in the format D:HH:MM",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_fsx_filesystem_windows_configuration",
				Description: "The configuration for this Microsoft Windows file system",
				Resolver:    schema.PathTableResolver("WindowsConfiguration"),
				Columns: []schema.Column{
					{
						Name:        "filesystem_cq_id",
						Description: "Unique CloudQuery ID of aws_fsx_filesystems table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "active_directory_id",
						Description: "The ID for an existing Amazon Web Services Managed Microsoft Active Directory instance that the file system is joined to",
						Type:        schema.TypeString,
					},
					{
						Name:        "aliases",
						Description: "An array of one or more DNS aliases that are currently associated with the Amazon FSx file system",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "file_access_audit_log_level",
						Description: "Sets which attempt type is logged by Amazon FSx for file and folder accesses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuditLogConfiguration.FileAccessAuditLogLevel"),
					},
					{
						Name:        "file_share_access_audit_log_level",
						Description: "Sets which attempt type is logged by Amazon FSx for file share accesses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuditLogConfiguration.FileShareAccessAuditLogLevel"),
					},
					{
						Name:        "audit_log_destination",
						Description: "The Amazon Resource Name (ARN) for the destination of the audit logs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuditLogConfiguration.AuditLogDestination"),
					},
					{
						Name:        "automatic_backup_retention_days",
						Description: "The number of days to retain automatic backups",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "copy_tags_to_backups",
						Description: "A boolean flag indicating whether tags on the file system should be copied to backups",
						Type:        schema.TypeBool,
					},
					{
						Name:        "daily_automatic_backup_start_time",
						Description: "The preferred time to take daily automatic backups, in the UTC time zone",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_type",
						Description: "Specifies the file system deployment type, valid values are the following:  * MULTI_AZ_1 - Specifies a high availability file system that is configured for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability, and supports SSD and HDD storage",
						Type:        schema.TypeString,
					},
					{
						Name:        "maintenance_operations_in_progress",
						Description: "The list of maintenance operations in progress for this file system",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "preferred_file_server_ip",
						Description: "For MULTI_AZ_1 deployment types, the IP address of the primary, or preferred, file server",
						Type:        schema.TypeString,
					},
					{
						Name:        "preferred_subnet_id",
						Description: "For MULTI_AZ_1 deployment types, it specifies the ID of the subnet where the preferred file server is located",
						Type:        schema.TypeString,
					},
					{
						Name:        "remote_administration_endpoint",
						Description: "For MULTI_AZ_1 deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell",
						Type:        schema.TypeString,
					},
					{
						Name:        "self_managed_ad_config_dns_ips",
						Description: "A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SelfManagedActiveDirectoryConfiguration.DnsIps"),
					},
					{
						Name:        "self_managed_ad_config_domain_name",
						Description: "The fully qualified domain name of the self-managed AD directory",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SelfManagedActiveDirectoryConfiguration.DomainName"),
					},
					{
						Name:        "self_managed_ad_config_file_system_administrators_group",
						Description: "The name of the domain group whose members have administrative privileges for the FSx file system",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SelfManagedActiveDirectoryConfiguration.FileSystemAdministratorsGroup"),
					},
					{
						Name:        "self_managed_ad_config_organizational_unit_distinguished_name",
						Description: "The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SelfManagedActiveDirectoryConfiguration.OrganizationalUnitDistinguishedName"),
					},
					{
						Name:        "self_managed_ad_config_user_name",
						Description: "The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SelfManagedActiveDirectoryConfiguration.UserName"),
					},
					{
						Name:        "throughput_capacity",
						Description: "The throughput of the Amazon FSx file system, measured in megabytes per second",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "weekly_maintenance_start_time",
						Description: "The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone",
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

func fetchFsxFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeFileSystemsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.FileSystems
	}
	return nil
}
