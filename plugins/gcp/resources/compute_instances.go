package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeInstances() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_instances",
		Resolver:     fetchComputeInstances,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "advanced_machine_features_enable_nested_virtualization",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AdvancedMachineFeatures.EnableNestedVirtualization"),
			},
			{
				Name: "can_ip_forward",
				Type: schema.TypeBool,
			},
			{
				Name:     "confidential_instance_config_enable_confidential_compute",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfidentialInstanceConfig.EnableConfidentialCompute"),
			},
			{
				Name: "cpu_platform",
				Type: schema.TypeString,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "deletion_protection",
				Type: schema.TypeBool,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "display_device_enable_display",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisplayDevice.EnableDisplay"),
			},
			{
				Name: "fingerprint",
				Type: schema.TypeString,
			},
			{
				Name:     "guest_accelerators",
				Type:     schema.TypeJSON,
				Resolver: resolveComputeInstanceGuestAccelerators,
			},
			{
				Name: "hostname",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "label_fingerprint",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "last_start_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "last_stop_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "last_suspended_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "machine_type",
				Type: schema.TypeString,
			},
			{
				Name:     "metadata_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Metadata.Fingerprint"),
			},
			{
				Name:     "metadata_items",
				Type:     schema.TypeJSON,
				Resolver: resolveComputeInstanceMetadataItems,
			},
			{
				Name:     "metadata_kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Metadata.Kind"),
			},
			{
				Name: "min_cpu_platform",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "post_key_revocation_action_type",
				Type: schema.TypeString,
			},
			{
				Name: "private_ipv6_google_access",
				Type: schema.TypeString,
			},
			{
				Name:     "reservation_affinity_consume_reservation_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservationAffinity.ConsumeReservationType"),
			},
			{
				Name:     "reservation_affinity_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservationAffinity.Key"),
			},
			{
				Name:     "reservation_affinity_values",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReservationAffinity.Values"),
			},
			{
				Name: "resource_policies",
				Type: schema.TypeStringArray,
			},
			{
				Name: "satisfies_pzs",
				Type: schema.TypeBool,
			},
			{
				Name:     "scheduling_automatic_restart",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Scheduling.AutomaticRestart"),
			},
			{
				Name:     "scheduling_location_hint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scheduling.LocationHint"),
			},
			{
				Name:     "scheduling_min_node_cpus",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Scheduling.MinNodeCpus"),
			},
			{
				Name:     "scheduling_on_host_maintenance",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Scheduling.OnHostMaintenance"),
			},
			{
				Name:     "scheduling_preemptible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Scheduling.Preemptible"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name:     "shielded_instance_config_enable_integrity_monitoring",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ShieldedInstanceConfig.EnableIntegrityMonitoring"),
			},
			{
				Name:     "shielded_instance_config_enable_secure_boot",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ShieldedInstanceConfig.EnableSecureBoot"),
			},
			{
				Name:     "shielded_instance_config_enable_vtpm",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ShieldedInstanceConfig.EnableVtpm"),
			},
			{
				Name:     "shielded_instance_integrity_policy_update_auto_learn_policy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ShieldedInstanceIntegrityPolicy.UpdateAutoLearnPolicy"),
			},
			{
				Name: "start_restricted",
				Type: schema.TypeBool,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "status_message",
				Type: schema.TypeString,
			},
			{
				Name:     "tags_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tags.Fingerprint"),
			},
			{
				Name:     "tags_items",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags.Items"),
			},
			{
				Name: "zone",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_instance_disks",
				Resolver: fetchComputeInstanceDisks,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "auto_delete",
						Type: schema.TypeBool,
					},
					{
						Name: "boot",
						Type: schema.TypeBool,
					},
					{
						Name: "device_name",
						Type: schema.TypeString,
					},
					{
						Name:     "disk_encryption_key_kms_key_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DiskEncryptionKey.KmsKeyName"),
					},
					{
						Name:     "disk_encryption_key_kms_key_service_account",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DiskEncryptionKey.KmsKeyServiceAccount"),
					},
					{
						Name:     "disk_encryption_key_raw_key",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DiskEncryptionKey.RawKey"),
					},
					{
						Name:     "disk_encryption_key_sha256",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DiskEncryptionKey.Sha256"),
					},
					{
						Name: "disk_size_gb",
						Type: schema.TypeBigInt,
					},
					{
						Name:     "guest_os_features",
						Type:     schema.TypeStringArray,
						Resolver: resolveComputeInstanceDiskGuestOsFeatures,
					},
					{
						Name: "index",
						Type: schema.TypeBigInt,
					},
					{
						Name:     "description",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.Description"),
					},
					{
						Name:     "disk_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.DiskName"),
					},
					{
						Name:     "initialized_disk_size_gb",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("InitializeParams.DiskSizeGb"),
					},
					{
						Name:     "disk_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.DiskType"),
					},
					{
						Name:     "labels",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("InitializeParams.Labels"),
					},
					{
						Name:     "on_update_action",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.OnUpdateAction"),
					},
					{
						Name:     "provisioned_iops",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("InitializeParams.ProvisionedIops"),
					},
					{
						Name:     "resource_policies",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("InitializeParams.ResourcePolicies"),
					},
					{
						Name:     "source_image",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceImage"),
					},
					{
						Name:     "source_image_encryption_key_kms_key_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceImageEncryptionKey.KmsKeyName"),
					},
					{
						Name:     "source_image_encryption_key_kms_key_service_account",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceImageEncryptionKey.KmsKeyServiceAccount"),
					},
					{
						Name:     "source_image_encryption_key_raw_key",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceImageEncryptionKey.RawKey"),
					},
					{
						Name:     "source_image_encryption_key_sha256",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceImageEncryptionKey.Sha256"),
					},
					{
						Name:     "source_snapshot",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceSnapshot"),
					},
					{
						Name:     "source_snapshot_encryption_key_kms_key_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.KmsKeyName"),
					},
					{
						Name:     "source_snapshot_encryption_key_kms_key_service_account",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.KmsKeyServiceAccount"),
					},
					{
						Name:     "source_snapshot_encryption_key_raw_key",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.RawKey"),
					},
					{
						Name:     "source_snapshot_encryption_key_sha256",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("InitializeParams.SourceSnapshotEncryptionKey.Sha256"),
					},
					{
						Name: "interface",
						Type: schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name: "licenses",
						Type: schema.TypeStringArray,
					},
					{
						Name: "mode",
						Type: schema.TypeString,
					},
					{
						Name:     "shielded_instance_initial_state_pk_content",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ShieldedInstanceInitialState.Pk.Content"),
					},
					{
						Name:     "shielded_instance_initial_state_pk_file_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ShieldedInstanceInitialState.Pk.FileType"),
					},
					{
						Name: "source",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "gcp_compute_instance_network_interfaces",
				Resolver: fetchComputeInstanceNetworkInterfaces,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "fingerprint",
						Type: schema.TypeString,
					},
					{
						Name: "ipv6_address",
						Type: schema.TypeString,
					},
					{
						Name: "kind",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "network",
						Type: schema.TypeString,
					},
					{
						Name:     "network_ip",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("NetworkIP"),
					},
					{
						Name: "nic_type",
						Type: schema.TypeString,
					},
					{
						Name: "subnetwork",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "gcp_compute_instance_network_interface_access_configs",
						Resolver: fetchComputeInstanceNetworkInterfaceAccessConfigs,
						Columns: []schema.Column{
							{
								Name:     "instance_network_interface_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "kind",
								Type: schema.TypeString,
							},
							{
								Name: "name",
								Type: schema.TypeString,
							},
							{
								Name:     "nat_ip",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("NatIP"),
							},
							{
								Name: "network_tier",
								Type: schema.TypeString,
							},
							{
								Name: "public_ptr_domain_name",
								Type: schema.TypeString,
							},
							{
								Name: "set_public_ptr",
								Type: schema.TypeBool,
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "gcp_compute_instance_network_interface_alias_ip_ranges",
						Resolver: fetchComputeInstanceNetworkInterfaceAliasIpRanges,
						Columns: []schema.Column{
							{
								Name:     "instance_network_interface_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "ip_cidr_range",
								Type: schema.TypeString,
							},
							{
								Name: "subnetwork_range_name",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "gcp_compute_instance_scheduling_node_affinities",
				Resolver: fetchComputeInstanceSchedulingNodeAffinities,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "key",
						Type: schema.TypeString,
					},
					{
						Name: "operator",
						Type: schema.TypeString,
					},
					{
						Name: "values",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "gcp_compute_instance_service_accounts",
				Resolver: fetchComputeInstanceServiceAccounts,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "email",
						Type: schema.TypeString,
					},
					{
						Name: "scopes",
						Type: schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Instances.AggregatedList(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
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
func resolveComputeInstanceGuestAccelerators(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*compute.Instance)
	res := map[string]int64{}
	for _, v := range r.GuestAccelerators {
		res[v.AcceleratorType] = v.AcceleratorCount
	}
	resource.Set("guest_accelerators", res)
	return nil
}
func resolveComputeInstanceMetadataItems(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*compute.Instance)
	res := map[string]string{}
	if r.Metadata != nil {
		for _, v := range r.Metadata.Items {
			res[v.Key] = *v.Value
		}
	}
	resource.Set("metadata_items", res)
	return nil
}
func fetchComputeInstanceDisks(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Instance)
	res <- r.Disks
	return nil
}
func resolveComputeInstanceDiskGuestOsFeatures(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*compute.AttachedDisk)
	res := make([]string, len(r.GuestOsFeatures))
	for i, v := range r.GuestOsFeatures {
		res[i] = v.Type
	}
	resource.Set("guest_os_features", res)
	return nil
}

func fetchComputeInstanceNetworkInterfaces(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Instance)
	res <- r.NetworkInterfaces
	return nil
}
func fetchComputeInstanceNetworkInterfaceAccessConfigs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.NetworkInterface)
	res <- r.AccessConfigs
	return nil
}
func fetchComputeInstanceNetworkInterfaceAliasIpRanges(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.NetworkInterface)
	res <- r.AliasIpRanges
	return nil
}
func fetchComputeInstanceSchedulingNodeAffinities(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Instance)
	if r.Scheduling != nil {
		res <- r.Scheduling.NodeAffinities
	}
	return nil
}
func fetchComputeInstanceServiceAccounts(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(*compute.Instance)
	if r.Scheduling != nil {
		res <- r.ServiceAccounts
	}
	return nil
}
