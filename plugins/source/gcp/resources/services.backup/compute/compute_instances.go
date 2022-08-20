package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeInstances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_instances",
		Description: "Represents an Instance resource  An instance is a virtual machine that is hosted on Google Cloud Platform For more information, read Virtual Machine Instances",
		Resolver:    fetchComputeInstances,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "advanced_machine_features_enable_nested_virtualization",
				Description: "Whether to enable nested virtualization or not (default is false)",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AdvancedMachineFeatures.EnableNestedVirtualization"),
			},
			{
				Name:        "can_ip_forward",
				Description: "Allows this instance to send and receive packets with non-matching destination or source IPs This is required if you plan to use this instance to forward routes For more information, see Enabling IP Forwarding",
				Type:        schema.TypeBool,
			},
			{
				Name:        "confidential_instance_config_enable_confidential_compute",
				Description: "Defines whether the instance should have confidential compute enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ConfidentialInstanceConfig.EnableConfidentialCompute"),
			},
			{
				Name:        "cpu_platform",
				Description: "The CPU platform used by this instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "deletion_protection",
				Description: "Whether the resource should be protected against deletion",
				Type:        schema.TypeBool,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "display_device_enable_display",
				Description: "Defines whether the instance has Display enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DisplayDevice.EnableDisplay"),
			},
			{
				Name:        "fingerprint",
				Description: "Specifies a fingerprint for this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "guest_accelerators",
				Description: "A list of the type and count of accelerator cards attached to the instance",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeInstanceGuestAccelerators,
			},
			{
				Name:        "hostname",
				Description: "Specifies the hostname of the instance The specified hostname must be RFC1035 compliant If hostname is not specified, the default hostname is [INSTANCE_NAME]c[PROJECT_ID]internal when using the global DNS, and [INSTANCE_NAME][ZONE]c[PROJECT_ID]internal when using zonal DNS",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#instance for instances",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_fingerprint",
				Description: "A fingerprint for the labels being applied to this image",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Labels for this resource",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "last_start_timestamp",
				Description: "Last start timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_stop_timestamp",
				Description: "Last stop timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_suspended_timestamp",
				Description: "Last suspended timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "machine_type",
				Description: "Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type",
				Type:        schema.TypeString,
			},
			{
				Name:        "metadata_fingerprint",
				Description: "Specifies a fingerprint for this request",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Fingerprint"),
			},
			{
				Name:        "metadata_items",
				Description: "Array of key/value pairs The total size of all keys and values must be less than 512 KB",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeInstanceMetadataItems,
			},
			{
				Name:        "metadata_kind",
				Description: "Type of the resource Always compute#metadata for metadata",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Metadata.Kind"),
			},
			{
				Name:        "min_cpu_platform",
				Description: "Specifies a minimum CPU platform for the VM instance Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: \"Intel Haswell\" or minCpuPlatform: \"Intel Sandy Bridge\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ipv6_google_access",
				Description: "The private IPv6 google access type for the VM If not specified, use  INHERIT_FROM_SUBNETWORK as default",
				Type:        schema.TypeString,
			},
			{
				Name:        "reservation_affinity_consume_reservation_type",
				Description: "Specifies the type of reservation from which this instance can consume resources: ANY_RESERVATION (default), SPECIFIC_RESERVATION, or NO_RESERVATION",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReservationAffinity.ConsumeReservationType"),
			},
			{
				Name:        "reservation_affinity_key",
				Description: "Corresponds to the label key of a reservation resource To target a SPECIFIC_RESERVATION by name, specify googleapiscom/reservation-name as the key and specify the name of your reservation as its value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ReservationAffinity.Key"),
			},
			{
				Name:          "reservation_affinity_values",
				Description:   "Corresponds to the label values of a reservation resource",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("ReservationAffinity.Values"),
			},
			{
				Name:          "resource_policies",
				Description:   "Resource policies applied to this instance",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "satisfies_pzs",
				Description: "Reserved for future use",
				Type:        schema.TypeBool,
			},
			{
				Name:        "scheduling_automatic_restart",
				Description: "Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user) You can only set the automatic restart option for standard instances Preemptible instances cannot be automatically restarted  By default, this is set to true so an instance is automatically restarted if it is terminated by Compute Engine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Scheduling.AutomaticRestart"),
			},
			{
				Name:        "scheduling_location_hint",
				Description: "An opaque location hint used to place the instance close to other resources This field is for use by internal tools that use the public API",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Scheduling.LocationHint"),
			},
			{
				Name:        "scheduling_min_node_cpus",
				Description: "The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Scheduling.MinNodeCpus"),
			},
			{
				Name:        "scheduling_on_host_maintenance",
				Description: "Defines the maintenance behavior for this instance For standard instances, the default behavior is MIGRATE For preemptible instances, the default and only possible behavior is TERMINATE For more information, see Setting Instance Scheduling Options",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Scheduling.OnHostMaintenance"),
			},
			{
				Name:        "scheduling_preemptible",
				Description: "Defines whether the instance is preemptible This can only be set during instance creation or while the instance is stopped and therefore, in a `TERMINATED` state See Instance Life Cycle for more information on the possible instance states",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Scheduling.Preemptible"),
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "shielded_instance_config_enable_integrity_monitoring",
				Description: "Defines whether the instance has integrity monitoring enabled Enabled by default",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ShieldedInstanceConfig.EnableIntegrityMonitoring"),
			},
			{
				Name:        "shielded_instance_config_enable_secure_boot",
				Description: "Defines whether the instance has Secure Boot enabled Disabled by default",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ShieldedInstanceConfig.EnableSecureBoot"),
			},
			{
				Name:        "shielded_instance_config_enable_vtpm",
				Description: "Defines whether the instance has the vTPM enabled Enabled by default",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ShieldedInstanceConfig.EnableVtpm"),
			},
			{
				Name:        "shielded_instance_integrity_policy_update_auto_learn_policy",
				Description: "Updates the integrity policy baseline using the measurements from the VM instance's most recent boot",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ShieldedInstanceIntegrityPolicy.UpdateAutoLearnPolicy"),
			},
			{
				Name:        "start_restricted",
				Description: "Whether a VM has been restricted for start because Compute Engine has detected suspicious activity",
				Type:        schema.TypeBool,
			},
			{
				Name:        "status",
				Description: "The status of the instance One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED For more information about the status of the instance, see  Instance life cycle",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "An optional, human-readable explanation of the status",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags_fingerprint",
				Description: "Specifies a fingerprint for this request, which is essentially a hash of the tags' contents and used for optimistic locking The fingerprint is initially generated by Compute Engine and changes after every request",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Tags.Fingerprint"),
			},
			{
				Name:        "tags_items",
				Description: "An array of tags Each tag must be 1-63 characters long, and comply with RFC1035",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Tags.Items"),
			},
			{
				Name:        "zone",
				Description: "URL of the zone where the instance resides You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_instance_disks",
				Description: "An instance-attached disk resource",
				Resolver:    fetchComputeInstanceDisks,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_compute_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "auto_delete",
						Description: "Specifies whether the disk will be auto-deleted when the instance is deleted (but not when the disk is detached from the instance)",
						Type:        schema.TypeBool,
					},
					{
						Name:        "boot",
						Description: "Indicates that this is a boot disk The virtual machine will use the first partition of the disk for its root filesystem",
						Type:        schema.TypeBool,
					},
					{
						Name:        "device_name",
						Description: "Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance This name can be used to reference the device for mounting, resizing, and so on, from within the instance  If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine This field is only applicable for persistent disks",
						Type:        schema.TypeString,
					},
					{
						Name:        "disk_encryption_key_kms_key_name",
						Description: "The name of the encryption key that is stored in Google Cloud KMS",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskEncryptionKey.KmsKeyName"),
					},
					{
						Name:        "disk_encryption_key_kms_key_service_account",
						Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskEncryptionKey.KmsKeyServiceAccount"),
					},
					{
						Name:        "disk_encryption_key_raw_key",
						Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskEncryptionKey.RawKey"),
					},
					{
						Name:        "disk_encryption_key_sha256",
						Description: "[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskEncryptionKey.Sha256"),
					},
					{
						Name:        "disk_size_gb",
						Description: "The size of the disk in GB",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "guest_os_features",
						Description: "A list of features to enable on the guest operating system Applicable only for bootable images Read  Enabling guest operating system features to see a list of available options",
						Type:        schema.TypeStringArray,
						Resolver:    resolveComputeInstanceDiskGuestOsFeatures,
					},
					{
						Name:        "index",
						Description: "A zero-based index to this disk, where 0 is reserved for the boot disk If you have many disks attached to an instance, each disk would have a unique index number",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "description",
						Description: "An optional description Provide this property when creating the disk",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.Description"),
					},
					{
						Name:        "disk_name",
						Description: "Specifies the disk name If not specified, the default is to use the name of the instance If a disk with the same name already exists in the given region, the existing disk is attached to the new instance and the new disk is not created",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.DiskName"),
					},
					{
						Name:        "initialized_disk_size_gb",
						Description: "Specifies the size of the disk in base-2 GB The size must be at least 10 GB If you specify a sourceImage, which is required for boot disks, the default size is the size of the sourceImage If you do not specify a sourceImage, the default disk size is 500 GB",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("InitializeParams.DiskSizeGb"),
					},
					{
						Name:        "disk_type",
						Description: "Specifies the disk type to use to create the instance",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.DiskType"),
					},
					{
						Name:          "labels",
						Description:   "Labels to apply to this disk These can be later modified by the diskssetLabels method This field is only applicable for persistent disks",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("InitializeParams.Labels"),
					},
					{
						Name:        "on_update_action",
						Description: "Specifies which action to take on instance update with this disk Default is to use the existing disk",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.OnUpdateAction"),
					},
					{
						Name:        "provisioned_iops",
						Description: "Indicates how many IOPS must be provisioned for the disk",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("InitializeParams.ProvisionedIops"),
					},
					{
						Name:          "resource_policies",
						Description:   "Resource policies applied to this disk for automatic snapshot creations Specified using the full or partial URL For instance template, specify only the resource policy name",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
						Resolver:      schema.PathResolver("InitializeParams.ResourcePolicies"),
					},
					{
						Name:        "source_image",
						Description: "The source image to create this disk",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceImage"),
					},
					{
						Name:        "source_image_encryption_key_kms_key_name",
						Description: "The name of the encryption key that is stored in Google Cloud KMS",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceImageEncryptionKey.KmsKeyName"),
					},
					{
						Name:        "source_image_encryption_key_kms_key_service_account",
						Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceImageEncryptionKey.KmsKeyServiceAccount"),
					},
					{
						Name:        "source_image_encryption_key_raw_key",
						Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceImageEncryptionKey.RawKey"),
					},
					{
						Name:        "source_image_encryption_key_sha256",
						Description: "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceImageEncryptionKey.Sha256"),
					},
					{
						Name:        "source_snapshot",
						Description: "The source snapshot to create this disk",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceSnapshot"),
					},
					{
						Name:        "source_snapshot_encryption_key_kms_key_name",
						Description: "The name of the encryption key that is stored in Google Cloud KMS",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.KmsKeyName"),
					},
					{
						Name:        "source_snapshot_encryption_key_kms_key_service_account",
						Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.KmsKeyServiceAccount"),
					},
					{
						Name:        "source_snapshot_encryption_key_raw_key",
						Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.RawKey"),
					},
					{
						Name:        "source_snapshot_encryption_key_sha256",
						Description: "[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.Sha256"),
					},
					{
						Name:        "interface",
						Description: "Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME The default is SCSI Persistent disks must always use SCSI and the request will fail if you attempt to attach a persistent disk in any other format than SCSI Local SSDs can use either NVME or SCSI For performance characteristics of SCSI over NVMe, see Local SSD performance",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "Type of the resource Always compute#attachedDisk for attached disks",
						Type:        schema.TypeString,
					},
					{
						Name:        "licenses",
						Description: "Any valid publicly visible licenses",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "mode",
						Description: "The mode in which to attach this disk, either READ_WRITE or READ_ONLY If not specified, the default is to attach the disk in READ_WRITE mode",
						Type:        schema.TypeString,
					},
					{
						Name:        "shielded_instance_initial_state_pk_content",
						Description: "The raw content in the secure keys file",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ShieldedInstanceInitialState.Pk.Content"),
					},
					{
						Name:        "shielded_instance_initial_state_pk_file_type",
						Description: "The file type of source file",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ShieldedInstanceInitialState.Pk.FileType"),
					},
					{
						Name:        "source",
						Description: "The source snapshot to create this disk",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Specifies the type of the disk, either SCRATCH or PERSISTENT If not specified, the default is PERSISTENT",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_compute_instance_network_interfaces",
				Description: "A network interface resource attached to an instance",
				Resolver:    fetchComputeInstanceNetworkInterfaces,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_compute_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "fingerprint",
						Description: "Fingerprint hash of contents stored in this network interface This field will be ignored when inserting an Instance or adding a NetworkInterface An up-to-date fingerprint must be provided in order to update the NetworkInterface The request will fail with error 400 Bad Request if the fingerprint is not provided, or 412 Precondition Failed if the fingerprint is out of date",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv6_address",
						Description: "An IPv6 internal network address for this network interface",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "Type of the resource Always compute#networkInterface for network interfaces",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the network interface, which is generated by the server For network devices, these are eth0, eth1, etc",
						Type:        schema.TypeString,
					},
					{
						Name:        "network",
						Description: "URL of the network resource for this instance When creating an instance, if neither the network nor the subnetwork is specified, the default network global/networks/default is used; if the network is not specified but the subnetwork is specified, the network is inferred  If you specify this property, you can specify the network as a full or partial URL For example, the following are all valid URLs: - https://wwwgoogleapis",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_ip",
						Description: "An IPv4 internal IP address to assign to the instance for this network interface If not specified by the user, an unused internal IP is assigned by the system",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NetworkIP"),
					},
					{
						Name:        "nic_type",
						Description: "The type of vNIC to be used on this interface This may be gVNIC or VirtioNet",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnetwork",
						Description: "The URL of the Subnetwork resource for this instance If the network resource is in legacy mode, do not specify this field If the network is in auto subnet mode, specifying the subnetwork is optional If the network is in custom subnet mode, specifying the subnetwork is required If you specify this field, you can specify the subnetwork as a full or partial URL For example, the following are all valid URLs: - https://wwwgoogleapis",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "gcp_compute_instance_network_interface_access_configs",
						Description:   "An access configuration attached to an instance's network interface Only one access config per instance is supported",
						Resolver:      fetchComputeInstanceNetworkInterfaceAccessConfigs,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_cq_id",
								Description: "Unique ID of gcp_compute_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "instance_network_interface_name",
								Type:     schema.TypeString,
								Resolver: schema.ParentResourceFieldResolver("name"),
							},
							{
								Name:        "kind",
								Description: "Type of the resource Always compute#accessConfig for access configs",
								Type:        schema.TypeString,
							},
							{
								Name:        "name",
								Description: "The name of this access configuration The default and recommended name is External NAT, but you can use any arbitrary string, such as My external IP or Network Access",
								Type:        schema.TypeString,
							},
							{
								Name:        "nat_ip",
								Description: "An external IP address associated with this instance Specify an unused static external IP address available to the project or leave this field undefined to use an IP from a shared ephemeral IP address pool If you specify a static external IP address, it must live in the same region as the zone of the instance",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("NatIP"),
							},
							{
								Name:        "network_tier",
								Description: "This signifies the networking tier used for configuring this access configuration and can only take the following values: PREMIUM, STANDARD  If an AccessConfig is specified without a valid external IP address, an ephemeral IP will be created with this networkTier  If an AccessConfig with a valid external IP address is specified, it must match that of the networkTier associated with the Address resource owning that IP",
								Type:        schema.TypeString,
							},
							{
								Name:        "public_ptr_domain_name",
								Description: "The DNS domain name for the public PTR record You can set this field only if the `setPublicPtr` field is enabled",
								Type:        schema.TypeString,
							},
							{
								Name:        "set_public_ptr",
								Description: "Specifies whether a public DNS 'PTR' record should be created to map the external IP address of the instance to a DNS domain name",
								Type:        schema.TypeBool,
							},
							{
								Name:        "type",
								Description: "The type of configuration The default and only option is ONE_TO_ONE_NAT",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "gcp_compute_instance_network_interface_alias_ip_ranges",
						Description:   "An alias IP range attached to an instance's network interface",
						Resolver:      fetchComputeInstanceNetworkInterfaceAliasIpRanges,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_cq_id",
								Description: "Unique ID of gcp_compute_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "instance_network_interface_name",
								Type:     schema.TypeString,
								Resolver: schema.ParentResourceFieldResolver("name"),
							},
							{
								Name:        "ip_cidr_range",
								Description: "The IP alias ranges to allocate for this interface This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces This range may be a single IP address (such as 10234), a netmask (such as /24) or a CIDR-formatted string (such as 10120/24)",
								Type:        schema.TypeString,
							},
							{
								Name:        "subnetwork_range_name",
								Description: "The name of a subnetwork secondary IP range from which to allocate an IP alias range If not specified, the primary range of the subnetwork is used",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:          "gcp_compute_instance_scheduling_node_affinities",
				Description:   "Node Affinity: the configuration of desired nodes onto which this Instance could be scheduled",
				Resolver:      fetchComputeInstanceSchedulingNodeAffinities,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_compute_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "key",
						Description: "Corresponds to the label key of Node resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "Defines the operation of node selection Valid operators are IN for affinity and NOT_IN for anti-affinity",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "Corresponds to the label values of Node resource",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "gcp_compute_instance_service_accounts",
				Description: "A service account",

				Resolver: fetchComputeInstanceServiceAccounts,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique ID of gcp_compute_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "instance_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "email",
						Description: "Email address of the service account",
						Type:        schema.TypeString,
					},
					{
						Name:          "scopes",
						Description:   "The list of scopes to be made available for this service account",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.Instances.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var instances []*compute.Instance
		for _, items := range output.Items {
			instances = append(instances, items.Instances...)
		}
		res <- instances
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeInstanceGuestAccelerators(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Instance)
	res := map[string]int64{}
	for _, v := range r.GuestAccelerators {
		res[v.AcceleratorType] = v.AcceleratorCount
	}
	return errors.WithStack(resource.Set("guest_accelerators", res))
}
func resolveComputeInstanceMetadataItems(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Instance)
	res := map[string]string{}
	if r.Metadata != nil {
		for _, v := range r.Metadata.Items {
			res[v.Key] = *v.Value
		}
	}
	return errors.WithStack(resource.Set("metadata_items", res))
}
func fetchComputeInstanceDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Instance)
	res <- r.Disks
	return nil
}
func resolveComputeInstanceDiskGuestOsFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.AttachedDisk)
	res := make([]string, len(r.GuestOsFeatures))
	for i, v := range r.GuestOsFeatures {
		res[i] = v.Type
	}
	return errors.WithStack(resource.Set("guest_os_features", res))
}
func fetchComputeInstanceNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Instance)
	res <- r.NetworkInterfaces
	return nil
}
func fetchComputeInstanceNetworkInterfaceAccessConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.NetworkInterface)
	res <- r.AccessConfigs
	return nil
}
func fetchComputeInstanceNetworkInterfaceAliasIpRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.NetworkInterface)
	res <- r.AliasIpRanges
	return nil
}
func fetchComputeInstanceSchedulingNodeAffinities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Instance)
	if r.Scheduling != nil {
		res <- r.Scheduling.NodeAffinities
	}
	return nil
}
func fetchComputeInstanceServiceAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.Instance)
	res <- r.ServiceAccounts
	return nil
}
