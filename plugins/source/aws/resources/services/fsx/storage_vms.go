package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource storage_vms --config storage_vms.hcl --output .
func StorageVms() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_storage_vms",
		Description:  "Describes the Amazon FSx for NetApp ONTAP storage virtual machine (SVM) configuration",
		Resolver:     fetchFsxStorageVms,
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
				Name:        "ad_cfg_net_bios_name",
				Description: "The NetBIOS name of the Active Directory computer object that is joined to your SVM",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ActiveDirectoryConfiguration.NetBiosName"),
			},
			{
				Name:        "ad_cfg_dns_ips",
				Description: "A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ActiveDirectoryConfiguration.SelfManagedActiveDirectoryConfiguration.DnsIps"),
			},
			{
				Name:        "ad_cfg_domain_name",
				Description: "The fully qualified domain name of the self-managed AD directory",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ActiveDirectoryConfiguration.SelfManagedActiveDirectoryConfiguration.DomainName"),
			},
			{
				Name:        "ad_cfg_file_system_administrators_group",
				Description: "The name of the domain group whose members have administrative privileges for the FSx file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ActiveDirectoryConfiguration.SelfManagedActiveDirectoryConfiguration.FileSystemAdministratorsGroup"),
			},
			{
				Name:        "ad_cfg_organizational_unit_distinguished_name",
				Description: "The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ActiveDirectoryConfiguration.SelfManagedActiveDirectoryConfiguration.OrganizationalUnitDistinguishedName"),
			},
			{
				Name:        "ad_cfg_user_name",
				Description: "The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ActiveDirectoryConfiguration.SelfManagedActiveDirectoryConfiguration.UserName"),
			},
			{
				Name:        "creation_time",
				Description: "The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "endpoints_iscsi_dns_name",
				Description: "The Domain Name Service (DNS) name for the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoints.Iscsi.DNSName"),
			},
			{
				Name:        "endpoints_iscsi_ip_addresses",
				Description: "The SVM endpoint's IP addresses",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Endpoints.Iscsi.IpAddresses"),
			},
			{
				Name:        "endpoints_management_dns_name",
				Description: "The Domain Name Service (DNS) name for the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoints.Management.DNSName"),
			},
			{
				Name:        "endpoints_management_ip_addresses",
				Description: "The SVM endpoint's IP addresses",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Endpoints.Management.IpAddresses"),
			},
			{
				Name:        "endpoints_nfs_dns_name",
				Description: "The Domain Name Service (DNS) name for the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoints.Nfs.DNSName"),
			},
			{
				Name:        "endpoints_nfs_ip_addresses",
				Description: "The SVM endpoint's IP addresses",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Endpoints.Nfs.IpAddresses"),
			},
			{
				Name:        "endpoints_smb_dns_name",
				Description: "The Domain Name Service (DNS) name for the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Endpoints.Smb.DNSName"),
			},
			{
				Name:        "endpoints_smb_ip_addresses",
				Description: "The SVM endpoint's IP addresses",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Endpoints.Smb.IpAddresses"),
			},
			{
				Name:        "file_system_id",
				Description: "The globally unique ID of the file system, assigned by Amazon FSx",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "Describes the SVM's lifecycle status",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle_transition_reason_message",
				Description: "A detailed error message",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LifecycleTransitionReason.Message"),
			},
			{
				Name:        "name",
				Description: "The name of the SVM, if provisioned",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for a given resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceARN"),
			},
			{
				Name:        "root_volume_security_style",
				Description: "The security style of the root volume of the SVM",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The SVM's system generated unique ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StorageVirtualMachineId"),
			},
			{
				Name:        "subtype",
				Description: "Describes the SVM's subtype",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A list of Tag values, with a maximum of 50 elements",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "uuid",
				Description: "The SVM's UUID (universally unique identifier)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UUID"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFsxStorageVms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeStorageVirtualMachinesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.StorageVirtualMachines
	}
	return nil
}
