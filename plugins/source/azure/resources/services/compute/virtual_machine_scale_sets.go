package compute

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func VirtualMachineScaleSets() *schema.Table {
	return &schema.Table{
		Name:         "azure_compute_virtual_machine_scale_sets",
		Description:  "VirtualMachineScaleSet describes a Virtual Machine Scale Set",
		Resolver:     fetchComputeVirtualMachineScaleSets,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		// This table is in bad shape and need terraforms and test to verify that it works.
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "sku_name",
				Description: "The sku name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Specifies the tier of virtual machines in a scale set<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_capacity",
				Description: "Specifies the number of virtual machines in the scale set",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "plan_name",
				Description: "The plan ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.Name"),
			},
			{
				Name:        "plan_publisher",
				Description: "The publisher ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.Publisher"),
			},
			{
				Name:        "plan_product",
				Description: "Specifies the product of the image from the marketplace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.Product"),
			},
			{
				Name:        "plan_promotion_code",
				Description: "The promotion code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.PromotionCode"),
			},
			{
				Name:        "upgrade_policy",
				Description: "The upgrade policy",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "automatic_repairs_policy_enabled",
				Description: "Specifies whether automatic repairs should be enabled on the virtual machine scale set",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.AutomaticRepairsPolicy.Enabled"),
			},
			{
				Name:        "automatic_repairs_policy_grace_period",
				Description: "The amount of time for which automatic repairs are suspended due to a state change on VM",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.AutomaticRepairsPolicy.GracePeriod"),
			},
			{
				Name:        "os_profile_computer_name_prefix",
				Description: "Specifies the computer name prefix for all of the virtual machines in the scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.ComputerNamePrefix"),
			},
			{
				Name:        "os_profile_admin_username",
				Description: "Specifies the name of the administrator account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.AdminUsername"),
			},
			{
				Name:        "os_profile_admin_password",
				Description: "Specifies the password of the administrator account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.AdminPassword"),
			},
			{
				Name:        "os_profile_custom_data",
				Description: "Specifies a base-64 encoded string of custom data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.CustomData"),
			},
			{
				Name:        "os_profile_windows_configuration",
				Description: "Specifies Windows operating system settings on the virtual machine",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsOsProfileWindowsConfiguration,
			},
			{
				Name:        "os_profile_linux_configuration",
				Description: "Specifies the Linux operating system settings on the virtual machine",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsOsProfileLinuxConfiguration,
			},
			{
				Name:        "storage_profile",
				Description: "Specifies the storage settings for the virtual machine disks",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsStorageProfile,
			},
			{
				Name:        "network_profile",
				Description: "Specifies properties of the network interfaces of the virtual machines in the scale set",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsNetworkProfile,
			},
			{
				Name:        "security_profile",
				Description: "Specifies the Security related profile settings for the virtual machines in the scale set",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsSecurityProfile,
			},
			{
				Name:        "diagnostics_profile",
				Description: "Specifies the boot diagnostic settings state",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsDiagnosticsProfile,
			},
			{
				Name:        "extension_profile_extensions_time_budget",
				Description: "Specifies the time alloted for all extensions to start",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.ExtensionProfile.ExtensionsTimeBudget"),
			},
			{
				Name:        "license_type",
				Description: "Specifies that the image or disk that is being used was licensed on-premises",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.LicenseType"),
			},
			{
				Name:        "priority",
				Description: "Specifies the priority for the virtual machines in the scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.Priority"),
			},
			{
				Name:        "eviction_policy",
				Description: "Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.EvictionPolicy"),
			},
			{
				Name:        "billing_profile_max_price",
				Description: "Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.BillingProfile.MaxPrice"),
			},
			{
				Name:        "scheduled_events_profile",
				Description: "Specifies Scheduled Event related configurations",
				Type:        schema.TypeJSON,
				Resolver:    resolveVirtualMachineScaleSetsScheduledEventsProfile,
			},
			{
				Name:        "user_data",
				Description: "UserData for the virtual machines in the scale set, which must be base-64 encoded",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.VirtualMachineProfile.UserData"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state, which only appears in the response",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.ProvisioningState"),
			},
			{
				Name:        "overprovision",
				Description: "Specifies whether the Virtual Machine Scale Set should be overprovisioned",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.Overprovision"),
			},
			{
				Name:        "do_not_run_extensions_on_overprovisioned_vms",
				Description: "When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.DoNotRunExtensionsOnOverprovisionedVMs"),
			},
			{
				Name:        "unique_id",
				Description: "Specifies the ID which uniquely identifies a Virtual Machine Scale Set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.UniqueID"),
			},
			{
				Name:        "single_placement_group",
				Description: "When true this limits the scale set to a single placement group, of max size 100 virtual machines",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.SinglePlacementGroup"),
			},
			{
				Name:        "zone_balance",
				Description: "Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.ZoneBalance"),
			},
			{
				Name:        "platform_fault_domain_count",
				Description: "Fault Domain count for each placement group",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.PlatformFaultDomainCount"),
			},
			{
				Name:        "proximity_placement_group_id",
				Description: "Proximity placement group resource id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.ProximityPlacementGroup.ID"),
			},
			{
				Name:        "host_group_id",
				Description: "Host group resource id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.HostGroup.ID"),
			},
			{
				Name:        "additional_capabilities_ultra_ssd_enabled",
				Description: "The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.AdditionalCapabilities.UltraSSDEnabled"),
			},
			{
				Name:        "scale_in_policy_rules",
				Description: "The rules to be followed when scaling-in a virtual machine scale set",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.ScaleInPolicy.Rules"),
			},
			{
				Name:        "orchestration_mode",
				Description: "Specifies the orchestration mode for the virtual machine scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineScaleSetProperties.OrchestrationMode"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal id of virtual machine scale set identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant id associated with the virtual machine scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the virtual machine scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user identities associated with the virtual machine scale set",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "zones",
				Description: "The virtual machine scale set zones",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "extended_location_name",
				Description: "The name of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_compute_virtual_machine_scale_set_os_profile_secrets",
				Description: "VaultSecretGroup describes a set of certificates which are all in the same Key Vault",
				Resolver:    fetchComputeVirtualMachineScaleSetOsProfileSecrets,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_scale_set_cq_id",
						Description: "Unique CloudQuery ID of azure_compute_virtual_machine_scale_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "source_vault_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceVault.ID"),
					},
					{
						Name:        "vault_certificates",
						Description: "The list of key vault references in SourceVault which contain certificates",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_scale_set_extensions",
				Description: "VirtualMachineScaleSetExtension describes a Virtual Machine Scale Set Extension",
				Resolver:    fetchComputeVirtualMachineScaleSetExtensions,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_scale_set_cq_id",
						Description: "Unique CloudQuery ID of azure_compute_virtual_machine_scale_sets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
						Resolver:    ResolveComputeVirtualMachineScaleSetExtensionType,
					},
					{
						Name:        "extension_type",
						Description: "The type of the extension",
						Type:        schema.TypeString,
						Resolver:    ResolveComputeVirtualMachineScaleSetExtensionExtensionType,
					},
					{
						Name:        "name",
						Description: "The name of the extension",
						Type:        schema.TypeString,
					},
					{
						Name:        "force_update_tag",
						Description: "If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.ForceUpdateTag"),
					},
					{
						Name:        "publisher",
						Description: "The name of the extension handler publisher",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.Publisher"),
					},
					{
						Name:        "type_handler_version",
						Description: "Specifies the version of the script handler",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.TypeHandlerVersion"),
					},
					{
						Name:        "auto_upgrade_minor_version",
						Description: "Indicates whether the extension should use a newer minor version if one is available at deployment time",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.AutoUpgradeMinorVersion"),
					},
					{
						Name:        "enable_automatic_upgrade",
						Description: "Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.EnableAutomaticUpgrade"),
					},
					{
						Name:        "settings",
						Description: "Json formatted public settings for the extension",
						Type:        schema.TypeJSON,
						Resolver:    resolveVirtualMachineScaleSetExtensionsSettings,
					},
					{
						Name:        "protected_settings",
						Description: "The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all",
						Type:        schema.TypeJSON,
						Resolver:    resolveVirtualMachineScaleSetExtensionsProtectedSettings,
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state, which only appears in the response",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.ProvisioningState"),
					},
					{
						Name:        "provision_after_extensions",
						Description: "Collection of extension names after which this extension needs to be provisioned",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VirtualMachineScaleSetExtensionProperties.ProvisionAfterExtensions"),
					},
					{
						Name:        "id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchComputeVirtualMachineScaleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachineScaleSets
	response, err := svc.ListAll(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func resolveVirtualMachineScaleSetsOsProfileWindowsConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.WindowsConfiguration == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.WindowsConfiguration)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetsOsProfileLinuxConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.LinuxConfiguration == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.LinuxConfiguration)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetsStorageProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.StorageProfile == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.StorageProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetsNetworkProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.NetworkProfile == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.NetworkProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetsSecurityProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.SecurityProfile == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.SecurityProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetsDiagnosticsProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.DiagnosticsProfile == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.DiagnosticsProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetsScheduledEventsProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.ScheduledEventsProfile == nil {
		return nil
	}

	data, err := json.Marshal(*p.VirtualMachineScaleSetProperties.VirtualMachineProfile.ScheduledEventsProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func fetchComputeVirtualMachineScaleSetOsProfileSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.Secrets == nil {
		return nil
	}

	res <- *p.VirtualMachineScaleSetProperties.VirtualMachineProfile.OsProfile.Secrets
	return nil
}
func fetchComputeVirtualMachineScaleSetExtensions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(compute.VirtualMachineScaleSet)
	if p.VirtualMachineScaleSetProperties == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile == nil ||
		p.VirtualMachineScaleSetProperties.VirtualMachineProfile.ExtensionProfile == nil {
		return nil
	}

	res <- *p.VirtualMachineScaleSetProperties.VirtualMachineProfile.ExtensionProfile.Extensions
	return nil
}
func ResolveComputeVirtualMachineScaleSetExtensionType(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSetExtension)
	if p.VirtualMachineScaleSetExtensionProperties == nil ||
		p.VirtualMachineScaleSetExtensionProperties.Type == nil {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, p.VirtualMachineScaleSetExtensionProperties.Type))
}
func ResolveComputeVirtualMachineScaleSetExtensionExtensionType(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSetExtension)
	if p.Type == nil {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, p.Type))
}
func resolveVirtualMachineScaleSetExtensionsSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSetExtension)
	if p.VirtualMachineScaleSetExtensionProperties == nil ||
		p.VirtualMachineScaleSetExtensionProperties.Settings == nil {
		return nil
	}

	data, err := json.Marshal(p.VirtualMachineScaleSetExtensionProperties.Settings)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveVirtualMachineScaleSetExtensionsProtectedSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineScaleSetExtension)
	if p.VirtualMachineScaleSetExtensionProperties == nil ||
		p.VirtualMachineScaleSetExtensionProperties.ProtectedSettings == nil {
		return nil
	}

	data, err := json.Marshal(p.VirtualMachineScaleSetExtensionProperties.ProtectedSettings)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
