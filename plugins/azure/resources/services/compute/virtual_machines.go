package compute

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeVirtualMachines() *schema.Table {
	return &schema.Table{
		Name:         "azure_compute_virtual_machines",
		Description:  "VirtualMachine describes a Virtual Machine.",
		Resolver:     fetchComputeVirtualMachines,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:          "plan_name",
				Description:   "The plan ID.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Plan.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "plan_publisher",
				Description:   "The publisher ID.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Plan.Publisher"),
				IgnoreInTests: true,
			},
			{
				Name:          "plan_product",
				Description:   "Specifies the product of the image from the marketplace",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Plan.Product"),
				IgnoreInTests: true,
			},
			{
				Name:          "plan_promotion_code",
				Description:   "The promotion code.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Plan.PromotionCode"),
				IgnoreInTests: true,
			},
			{
				Name:        "hardware_profile_vm_size",
				Description: "Specifies the size of the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.HardwareProfile.VMSize"),
			},
			{
				Name:        "storage_profile",
				Description: "Specifies the storage settings for the virtual machine disks.",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachinesStorageProfile,
			},
			{
				Name:          "additional_capabilities_ultra_ssd_enabled",
				Description:   "The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.AdditionalCapabilities.UltraSSDEnabled"),
				IgnoreInTests: true,
			},
			{
				Name:        "computer_name",
				Description: "Specifies the host OS name of the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.ComputerName"),
			},
			{
				Name:        "admin_username",
				Description: "Specifies the name of the administrator account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AdminUsername"),
			},
			{
				Name:          "admin_password",
				Description:   "Specifies the password of the administrator account",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.AdminPassword"),
				IgnoreInTests: true,
			},
			{
				Name:          "custom_data",
				Description:   "Specifies a base-64 encoded string of custom data",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.CustomData"),
				IgnoreInTests: true,
			},
			{
				Name:        "windows_configuration_provision_vm_agent",
				Description: "Indicates whether virtual machine agent should be provisioned on the virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.ProvisionVMAgent"),
			},
			{
				Name:        "windows_configuration_enable_automatic_updates",
				Description: "Indicates whether Automatic Updates is enabled for the Windows virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.EnableAutomaticUpdates"),
			},
			{
				Name:          "windows_configuration_time_zone",
				Description:   "Specifies the time zone of the virtual machine",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.TimeZone"),
				IgnoreInTests: true,
			},
			{
				Name:          "windows_configuration_additional_unattend_content",
				Description:   "Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.",
				Type:          schema.TypeJSON,
				Resolver:      resolveComputeVirtualMachinesWindowsConfigurationAdditionalUnattendContent,
				IgnoreInTests: true,
			},
			{
				Name:        "windows_configuration_patch_settings_patch_mode",
				Description: "You  control the application of patches to a virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.PatchMode"),
			},
			{
				Name:          "windows_configuration_patch_settings_enable_hotpatching",
				Description:   "Enables customers to patch their Azure VMs without requiring a reboot",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.EnableHotpatching"),
				IgnoreInTests: true,
			},
			{
				Name:          "windows_configuration_patch_settings_assessment_mode",
				Description:   "The platform will trigger periodic patch assessments",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.AssessmentMode"),
				IgnoreInTests: true,
			},
			{
				Name:          "linux_configuration_disable_password_authentication",
				Description:   "Specifies whether password authentication should be disabled.",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.DisablePasswordAuthentication"),
				IgnoreInTests: true,
			},
			{
				Name:          "linux_configuration_ssh_public_keys",
				Description:   "The list of SSH public keys used to authenticate with linux based VMs.",
				Type:          schema.TypeJSON,
				Resolver:      resolveComputeVirtualMachinesLinuxConfigurationSshPublicKeys,
				IgnoreInTests: true,
			},
			{
				Name:          "linux_configuration_provision_vm_agent",
				Description:   "Indicates whether virtual machine agent should be provisioned on the virtual machine",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.ProvisionVMAgent"),
				IgnoreInTests: true,
			},
			{
				Name:        "linux_configuration_patch_settings_patch_mode",
				Description: "The virtual machine's default patching configuration is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.PatchSettings.PatchMode"),
			},
			{
				Name:        "linux_configuration_patch_settings_assessment_mode",
				Description: "You control the timing of patch assessments on a virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.PatchSettings.AssessmentMode"),
			},
			{
				Name:        "allow_extension_operations",
				Description: "Specifies whether extension operations should be allowed on the virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AllowExtensionOperations"),
			},
			{
				Name:        "require_guest_provision_signal",
				Description: "Specifies whether the guest provision signal is required to infer provision success of the virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.RequireGuestProvisionSignal"),
			},
			{
				Name:        "network_profile_network_interfaces",
				Description: "Specifies the list of resource Ids for the network interfaces associated with the virtual machine.",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachinesNetworkProfileNetworkInterfaces,
			},
			{
				Name:        "network_profile_network_api_version",
				Description: "specifies the Microsoft.Network API version used when creating networking resources in the Network Interface Configurations",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.NetworkProfile.NetworkAPIVersion"),
			},
			{
				Name:          "network_profile_network_interface_configurations",
				Description:   "Specifies the networking configurations that will be used to create the virtual machine networking resources.",
				Type:          schema.TypeJSON,
				Resolver:      resolveComputeVirtualMachinesNetworkProfileNetworkInterfaceConfigurations,
				IgnoreInTests: true,
			},
			{
				Name:          "security_profile_uefi_settings_secure_boot_enabled",
				Description:   "Specifies whether secure boot should be enabled on the virtual machine",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.SecurityProfile.UefiSettings.SecureBootEnabled"),
				IgnoreInTests: true,
			},
			{
				Name:          "security_profile_uefi_settings_v_tpm_enabled",
				Description:   "Specifies whether vTPM should be enabled on the virtual machine",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.SecurityProfile.UefiSettings.VTpmEnabled"),
				IgnoreInTests: true,
			},
			{
				Name:          "security_profile_encryption_at_host",
				Description:   "This property can be used by user in the request to enable or disable the Host Encryption for the virtual machine or virtual machine scale set",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualMachineProperties.SecurityProfile.EncryptionAtHost"),
				IgnoreInTests: true,
			},
			{
				Name:        "security_profile_security_type",
				Description: "Specifies the SecurityType of the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.SecurityType"),
			},
			{
				Name:        "diagnostics_profile_boot_diagnostics_enabled",
				Description: "Whether boot diagnostics should be enabled on the Virtual Machine.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.Enabled"),
			},
			{
				Name:          "diagnostics_profile_boot_diagnostics_storage_uri",
				Description:   "Uri of the storage account to use for placing the console output and screenshot",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.StorageURI"),
				IgnoreInTests: true,
			},
			{
				Name:          "availability_set_id",
				Description:   "Resource Id",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.AvailabilitySet.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "virtual_machine_scale_set_id",
				Description:   "Resource Id",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.VirtualMachineScaleSet.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "proximity_placement_group_id",
				Description:   "Resource Id",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.ProximityPlacementGroup.ID"),
				IgnoreInTests: true,
			},
			{
				Name:        "priority",
				Description: "Specifies the priority for the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.Priority"),
			},
			{
				Name:        "eviction_policy",
				Description: "Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.EvictionPolicy"),
			},
			{
				Name:          "billing_profile_max_price",
				Description:   "Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS",
				Type:          schema.TypeFloat,
				Resolver:      schema.PathResolver("VirtualMachineProperties.BillingProfile.MaxPrice"),
				IgnoreInTests: true,
			},
			{
				Name:          "host_id",
				Description:   "Resource Id",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.Host.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "host_group_id",
				Description:   "Resource Id",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.HostGroup.ID"),
				IgnoreInTests: true,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state, which only appears in the response.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ProvisioningState"),
			},
			{
				Name:          "instance_view",
				Description:   "The virtual machine instance view.",
				Type:          schema.TypeJSON,
				Resolver:      resolveComputeVirtualMachinesInstanceView,
				IgnoreInTests: true,
			},
			{
				Name:          "license_type",
				Description:   "Specifies that the image or disk that is being used was licensed on-premises",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.LicenseType"),
				IgnoreInTests: true,
			},
			{
				Name:        "vm_id",
				Description: "Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.VMID"),
			},
			{
				Name:          "extensions_time_budget",
				Description:   "Specifies the time alloted for all extensions to start",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.ExtensionsTimeBudget"),
				IgnoreInTests: true,
			},
			{
				Name:          "platform_fault_domain",
				Description:   "Specifies the scale set logical fault domain into which the Virtual Machine will be created.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("VirtualMachineProperties.PlatformFaultDomain"),
				IgnoreInTests: true,
			},
			{
				Name:          "scheduled_events_profile",
				Description:   "Specifies Scheduled Event related configurations.",
				Type:          schema.TypeJSON,
				Resolver:      resolveComputeVirtualMachinesScheduledEventsProfile,
				IgnoreInTests: true,
			},
			{
				Name:          "user_data",
				Description:   "UserData for the VM, which must be base-64 encoded",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualMachineProperties.UserData"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_principal_id",
				Description:   "The principal id of virtual machine identity",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_tenant_id",
				Description:   "The tenant id associated with the virtual machine",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:          "identity_user_assigned_identities",
				Description:   "The list of user identities associated with the Virtual Machine",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Identity.UserAssignedIdentities"),
				IgnoreInTests: true,
			},
			{
				Name:          "zones",
				Description:   "The virtual machine zones.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "extended_location_name",
				Description:   "The name of the extended location.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Name"),
				IgnoreInTests: true,
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
				Name:          "tags",
				Description:   "Resource tags",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_compute_virtual_machine_win_config_rm_listeners",
				Description:   "WinRMListener describes Protocol and thumbprint of Windows Remote Management listener",
				Resolver:      fetchComputeVirtualMachineWinConfigRmListeners,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_cq_id",
						Description: "Unique CloudQuery ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "virtual_machine_id",
						Description: "ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "protocol",
						Description: "Specifies the protocol of WinRM listener",
						Type:        schema.TypeString,
					},
					{
						Name:          "certificate_url",
						Description:   "This is the URL of a certificate that has been uploaded to Key Vault as a secret",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("CertificateURL"),
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "azure_compute_virtual_machine_secrets",
				Description:   "VaultSecretGroup describes a set of certificates which are all in the same Key Vault.",
				Resolver:      fetchComputeVirtualMachineSecrets,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_cq_id",
						Description: "Unique CloudQuery ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "virtual_machine_id",
						Description: "ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "source_vault_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceVault.ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_compute_virtual_machine_secret_vault_certificates",
						Description: "VaultCertificate describes a single certificate reference in a Key Vault, and where the certificate should reside on the VM.",
						Resolver:    fetchComputeVirtualMachineSecretVaultCertificates,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_secret_cq_id",
								Description: "Unique CloudQuery ID of azure_compute_virtual_machine_secrets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "certificate_url",
								Description: "This is the URL of a certificate that has been uploaded to Key Vault as a secret",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CertificateURL"),
							},
							{
								Name:          "certificate_store",
								Description:   "For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate should be added",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
						},
					},
				},
			},
			{
				Name:          "azure_compute_virtual_machine_resources",
				Description:   "VirtualMachineExtension describes a Virtual Machine Extension.",
				Resolver:      fetchComputeVirtualMachineResources,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_cq_id",
						Description: "Unique CloudQuery ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "virtual_machine_id",
						Description: "ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:          "force_update_tag",
						Description:   "How the extension handler should be forced to update even if the extension configuration has not changed.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualMachineExtensionProperties.ForceUpdateTag"),
						IgnoreInTests: true,
					},
					{
						Name:        "publisher",
						Description: "The name of the extension handler publisher.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.Publisher"),
					},
					{
						Name:        "type_handler_version",
						Description: "Specifies the version of the script handler.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.TypeHandlerVersion"),
					},
					{
						Name:        "auto_upgrade_minor_version",
						Description: "Indicates whether the extension should use a newer minor version if one is available at deployment time",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.AutoUpgradeMinorVersion"),
					},
					{
						Name:        "enable_automatic_upgrade",
						Description: "Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.EnableAutomaticUpgrade"),
					},
					{
						Name:        "settings",
						Description: "Json formatted public settings for the extension.",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeVirtualMachineResourcesSettings,
					},
					{
						Name:          "protected_settings",
						Description:   "The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.",
						Type:          schema.TypeJSON,
						Resolver:      resolveComputeVirtualMachineResourcesProtectedSettings,
						IgnoreInTests: true,
					},
					{
						Name:        "extension_type",
						Description: "Type of the extension",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.Type"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state, which only appears in the response.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.ProvisioningState"),
					},
					{
						Name:          "instance_view",
						Description:   "The virtual machine extension instance view.",
						Type:          schema.TypeJSON,
						Resolver:      resolveComputeVirtualMachineResourcesInstanceView,
						IgnoreInTests: true,
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
						Name:          "tags",
						Description:   "Resource tags",
						Type:          schema.TypeJSON,
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

func fetchComputeVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachines
	response, err := svc.ListAll(ctx, "false")
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
func resolveComputeVirtualMachinesStorageProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachine)
	data, err := json.Marshal(p.StorageProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}

func resolveComputeVirtualMachinesWindowsConfigurationAdditionalUnattendContent(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachine)
	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.OsProfile == nil ||
		p.VirtualMachineProperties.OsProfile.WindowsConfiguration == nil {
		return nil
	}

	data, err := json.Marshal(p.VirtualMachineProperties.OsProfile.WindowsConfiguration.AdditionalUnattendContent)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachinesLinuxConfigurationSshPublicKeys(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachine)
	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.OsProfile == nil ||
		p.VirtualMachineProperties.OsProfile.LinuxConfiguration == nil ||
		p.VirtualMachineProperties.OsProfile.LinuxConfiguration.SSH == nil {
		return nil
	}

	data, err := json.Marshal(p.VirtualMachineProperties.OsProfile.LinuxConfiguration.SSH.PublicKeys)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachinesNetworkProfileNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachine)
	if p.NetworkProfile == nil {
		return nil
	}

	data, err := json.Marshal(p.NetworkProfile.NetworkInterfaces)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachinesNetworkProfileNetworkInterfaceConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachine)
	if p.NetworkProfile == nil {
		return nil
	}

	data, err := json.Marshal(p.NetworkProfile.NetworkInterfaceConfigurations)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachinesInstanceView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachines
	p := resource.Item.(compute.VirtualMachine)
	details, err := client.ParseResourceID(*p.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.InstanceView(ctx, details.ResourceGroup, *p.Name)
	if err != nil {
		return diag.WrapError(err)
	}

	data, err := json.Marshal(response)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachinesScheduledEventsProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachine)
	if p.ScheduledEventsProfile == nil {
		return nil
	}

	data, err := json.Marshal(p.ScheduledEventsProfile)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func fetchComputeVirtualMachineWinConfigRmListeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(compute.VirtualMachine)
	if p.OsProfile == nil ||
		p.OsProfile.WindowsConfiguration == nil ||
		p.OsProfile.WindowsConfiguration.WinRM == nil ||
		p.OsProfile.WindowsConfiguration.WinRM.Listeners == nil {
		return nil
	}

	res <- *p.OsProfile.WindowsConfiguration.WinRM.Listeners
	return nil
}
func fetchComputeVirtualMachineSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(compute.VirtualMachine)
	if p.OsProfile == nil || p.OsProfile.Secrets == nil {
		return nil
	}

	res <- *p.OsProfile.Secrets
	return nil
}
func fetchComputeVirtualMachineSecretVaultCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(compute.VaultSecretGroup)
	if p.VaultCertificates == nil {
		return nil
	}

	res <- *p.VaultCertificates
	return nil
}
func fetchComputeVirtualMachineResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachineExtensions
	p := parent.Item.(compute.VirtualMachine)
	details, err := client.ParseResourceID(*p.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.List(ctx, details.ResourceGroup, *p.Name, "")
	if err != nil {
		return diag.WrapError(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	return nil
}
func resolveComputeVirtualMachineResourcesSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineExtension)
	data, err := json.Marshal(p.Settings)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachineResourcesProtectedSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineExtension)
	data, err := json.Marshal(p.ProtectedSettings)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveComputeVirtualMachineResourcesInstanceView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(compute.VirtualMachineExtension)
	data, err := json.Marshal(p.InstanceView)
	if err != nil {
		return diag.WrapError(err)
	}

	return diag.WrapError(resource.Set(c.Name, data))
}
