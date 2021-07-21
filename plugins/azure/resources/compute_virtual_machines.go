package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeVirtualMachines() *schema.Table {
	return &schema.Table{
		Name:         "azure_compute_virtual_machines",
		Description:  "VirtualMachine describes a Virtual Machine",
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
				Description: "Specifies the product of the image from the marketplace This is the same value as Offer under the imageReference element",
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
				Name:        "hardware_profile_vm_size",
				Description: "Specifies the size of the virtual machine <br><br> The enum data type is currently deprecated and will be removed by December 23rd 2023 <br><br> Recommended way to get the list of available sizes is using these APIs: <br><br> [List all available virtual machine sizes in an availability set](https://docsmicrosoftcom/rest/api/compute/availabilitysets/listavailablesizes) <br><br> [List all available virtual machine sizes in a region]( https://docsmicrosoftcom/en-us/rest/api/compute/resourceskus/list) <br><br> [List all available virtual machine sizes for resizing](https://docsmicrosoftcom/rest/api/compute/virtualmachines/listavailablesizes) For more information about virtual machine sizes, see [Sizes for virtual machines](https://docsmicrosoftcom/en-us/azure/virtual-machines/sizes) <br><br> The available VM sizes depend on region and availability set Possible values include: 'BasicA0', 'BasicA1', 'BasicA2', 'BasicA3', 'BasicA4', 'StandardA0', 'StandardA1', 'StandardA2', 'StandardA3', 'StandardA4', 'StandardA5', 'StandardA6', 'StandardA7', 'StandardA8', 'StandardA9', 'StandardA10', 'StandardA11', 'StandardA1V2', 'StandardA2V2', 'StandardA4V2', 'StandardA8V2', 'StandardA2mV2', 'StandardA4mV2', 'StandardA8mV2', 'StandardB1s', 'StandardB1ms', 'StandardB2s', 'StandardB2ms', 'StandardB4ms', 'StandardB8ms', 'StandardD1', 'StandardD2', 'StandardD3', 'StandardD4', 'StandardD11', 'StandardD12', 'StandardD13', 'StandardD14', 'StandardD1V2', 'StandardD2V2', 'StandardD3V2', 'StandardD4V2', 'StandardD5V2', 'StandardD2V3', 'StandardD4V3', 'StandardD8V3', 'StandardD16V3', 'StandardD32V3', 'StandardD64V3', 'StandardD2sV3', 'StandardD4sV3', 'StandardD8sV3', 'StandardD16sV3', 'StandardD32sV3', 'StandardD64sV3', 'StandardD11V2', 'StandardD12V2', 'StandardD13V2', 'StandardD14V2', 'StandardD15V2', 'StandardDS1', 'StandardDS2', 'StandardDS3', 'StandardDS4', 'StandardDS11', 'StandardDS12', 'StandardDS13', 'StandardDS14', 'StandardDS1V2', 'StandardDS2V2', 'StandardDS3V2', 'StandardDS4V2', 'StandardDS5V2', 'StandardDS11V2', 'StandardDS12V2', 'StandardDS13V2', 'StandardDS14V2', 'StandardDS15V2', 'StandardDS134V2', 'StandardDS132V2', 'StandardDS148V2', 'StandardDS144V2', 'StandardE2V3', 'StandardE4V3', 'StandardE8V3', 'StandardE16V3', 'StandardE32V3', 'StandardE64V3', 'StandardE2sV3', 'StandardE4sV3', 'StandardE8sV3', 'StandardE16sV3', 'StandardE32sV3', 'StandardE64sV3', 'StandardE3216V3', 'StandardE328sV3', 'StandardE6432sV3', 'StandardE6416sV3', 'StandardF1', 'StandardF2', 'StandardF4', 'StandardF8', 'StandardF16', 'StandardF1s', 'StandardF2s', 'StandardF4s', 'StandardF8s', 'StandardF16s', 'StandardF2sV2', 'StandardF4sV2', 'StandardF8sV2', 'StandardF16sV2', 'StandardF32sV2', 'StandardF64sV2', 'StandardF72sV2', 'StandardG1', 'StandardG2', 'StandardG3', 'StandardG4', 'StandardG5', 'StandardGS1', 'StandardGS2', 'StandardGS3', 'StandardGS4', 'StandardGS5', 'StandardGS48', 'StandardGS44', 'StandardGS516', 'StandardGS58', 'StandardH8', 'StandardH16', 'StandardH8m', 'StandardH16m', 'StandardH16r', 'StandardH16mr', 'StandardL4s', 'StandardL8s', 'StandardL16s', 'StandardL32s', 'StandardM64s', 'StandardM64ms', 'StandardM128s', 'StandardM128ms', 'StandardM6432ms', 'StandardM6416ms', 'StandardM12864ms', 'StandardM12832ms', 'StandardNC6', 'StandardNC12', 'StandardNC24', 'StandardNC24r', 'StandardNC6sV2', 'StandardNC12sV2', 'StandardNC24sV2', 'StandardNC24rsV2', 'StandardNC6sV3', 'StandardNC12sV3', 'StandardNC24sV3', 'StandardNC24rsV3', 'StandardND6s', 'StandardND12s', 'StandardND24s', 'StandardND24rs', 'StandardNV6', 'StandardNV12', 'StandardNV24'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.HardwareProfile.VMSize"),
			},
			{
				Name:        "storage_profile",
				Description: "Specifies the storage settings for the virtual machine disks",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachineStorageProfile,
			},
			{
				Name:        "additional_capabilities_ultra_ssd_enabled",
				Description: "The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS Managed disks with storage account type UltraSSD_LRS can be added to a virtual machine or virtual machine scale set only if this property is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.AdditionalCapabilities.UltraSSDEnabled"),
			},
			{
				Name:        "computer_name",
				Description: "Specifies the host OS name of the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.ComputerName"),
			},
			{
				Name:        "admin_username",
				Description: "Specifies the name of the administrator account <br><br> This property cannot be updated after the VM is created <br><br> **Windows-only restriction:** Cannot end in \"\" <br><br> **Disallowed values:** \"administrator\", \"admin\", \"user\", \"user1\", \"test\", \"user2\", \"test1\", \"user3\", \"admin1\", \"1\", \"123\", \"a\", \"actuser\", \"adm\", \"admin2\", \"aspnet\", \"backup\", \"console\", \"david\", \"guest\", \"john\", \"owner\", \"root\", \"server\", \"sql\", \"support\", \"support_388945a0\", \"sys\", \"test2\", \"test3\", \"user4\", \"user5\" <br><br> **Minimum-length (Linux):** 1  character <br><br> **Max-length (Linux):** 64 characters <br><br> **Max-length (Windows):** 20 characters  <br><br><li> For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)<br><li> For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AdminUsername"),
			},
			{
				Name:        "admin_password",
				Description: "Specifies the password of the administrator account <br><br> **Minimum-length (Windows):** 8 characters <br><br> **Minimum-length (Linux):** 6 characters <br><br> **Max-length (Windows):** 123 characters <br><br> **Max-length (Linux):** 72 characters <br><br> **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled <br> Has lower characters <br>Has upper characters <br> Has a digit <br> Has a special character (Regex match [\\W_]) <br><br> **Disallowed values:** \"abc@123\", \"P@$$w0rd\", \"P@ssw0rd\", \"P@ssword123\", \"Pa$$word\", \"pass@word1\", \"Password!\", \"Password1\", \"Password22\", \"iloveyou!\" <br><br> For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftocjson) <br><br> For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson#reset-root-password)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AdminPassword"),
			},
			{
				Name:        "custom_data",
				Description: "Specifies a base-64 encoded string of custom data The base-64 encoded string is decoded to a binary array that is saved as a file on the Virtual Machine The maximum length of the binary array is 65535 bytes <br><br> **Note: Do not pass any secrets or passwords in customData property** <br><br> This property cannot be updated after the VM is created <br><br> customData is passed to the VM to be saved as a file, for more information see [Custom Data on Azure VMs](https://azuremicrosoftcom/en-us/blog/custom-data-and-cloud-init-on-windows-azure/) <br><br> For using cloud-init for your Linux VM, see [Using cloud-init to customize a Linux VM during creation](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.CustomData"),
			},
			{
				Name:        "windows_configuration_provision_vm_agent",
				Description: "Indicates whether virtual machine agent should be provisioned on the virtual machine <br><br> When this property is not specified in the request body, default behavior is to set it to true  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.ProvisionVMAgent"),
			},
			{
				Name:        "windows_configuration_enable_automatic_updates",
				Description: "Indicates whether Automatic Updates is enabled for the Windows virtual machine Default value is true <br><br> For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.EnableAutomaticUpdates"),
			},
			{
				Name:        "windows_configuration_time_zone",
				Description: "Specifies the time zone of the virtual machine eg \"Pacific Standard Time\" <br><br> Possible values can be [TimeZoneInfoId](https://docsmicrosoftcom/en-us/dotnet/api/systemtimezoneinfoid?#System_TimeZoneInfo_Id) value from time zones returned by [TimeZoneInfoGetSystemTimeZones](https://docsmicrosoftcom/en-us/dotnet/api/systemtimezoneinfogetsystemtimezones)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.TimeZone"),
			},
			{
				Name:        "windows_configuration_additional_unattend_content",
				Description: "Specifies additional base-64 encoded XML formatted information that can be included in the Unattendxml file, which is used by Windows Setup",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachineWindowsConfigurationAdditionalUnattendContent,
			},
			{
				Name:        "windows_configuration_patch_settings_patch_mode",
				Description: "the property WindowsConfigurationenableAutomaticUpdates must be false<br /><br /> **AutomaticByOS** - The virtual machine will automatically be updated by the OS The property WindowsConfigurationenableAutomaticUpdates must be true <br /><br /> **AutomaticByPlatform** - the virtual machine will automatically updated by the platform The properties provisionVMAgent and WindowsConfigurationenableAutomaticUpdates must be true Possible values include: 'WindowsVMGuestPatchModeManual', 'WindowsVMGuestPatchModeAutomaticByOS', 'WindowsVMGuestPatchModeAutomaticByPlatform'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.PatchMode"),
			},
			{
				Name:        "windows_configuration_patch_settings_enable_hotpatching",
				Description: "Enables customers to patch their Azure VMs without requiring a reboot For enableHotpatching, the 'provisionVMAgent' must be set to true and 'patchMode' must be set to 'AutomaticByPlatform'",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.EnableHotpatching"),
			},
			{
				Name:        "linux_configuration_disable_password_authentication",
				Description: "Specifies whether password authentication should be disabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.DisablePasswordAuthentication"),
			},
			{
				Name:        "linux_configuration_ssh_public_keys",
				Description: "The list of SSH public keys used to authenticate with linux based VMs",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachineLinuxConfigurationSSHPublicKeys,
			},
			{
				Name:        "linux_configuration_provision_vm_agent",
				Description: "Indicates whether virtual machine agent should be provisioned on the virtual machine <br><br> When this property is not specified in the request body, default behavior is to set it to true  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.ProvisionVMAgent"),
			},
			{
				Name:        "linux_configuration_patch_settings_patch_mode",
				Description: "Specifies the mode of VM Guest Patching to IaaS virtual machine<br /><br /> Possible values are:<br /><br /> **ImageDefault** - The virtual machine's default patching configuration is used <br /><br /> **AutomaticByPlatform** - The virtual machine will be automatically updated by the platform The property provisionVMAgent must be true Possible values include: 'ImageDefault', 'AutomaticByPlatform'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.PatchSettings.PatchMode"),
			},
			{
				Name:        "allow_extension_operations",
				Description: "Specifies whether extension operations should be allowed on the virtual machine <br><br>This may only be set to False when no extensions are present on the virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AllowExtensionOperations"),
			},
			{
				Name:        "require_guest_provision_signal",
				Description: "Specifies whether the guest provision signal is required to infer provision success of the virtual machine  **Note: This property is for private testing only, and all customers must not set the property to false**",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.RequireGuestProvisionSignal"),
			},
			{
				Name:        "network_profile_network_interfaces",
				Description: "Specifies the list of resource Ids for the network interfaces associated with the virtual machine",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachineNetworkProfileNetworkInterfaces,
			},
			{
				Name:        "security_profile_uefi_settings_secure_boot_enabled",
				Description: "Specifies whether secure boot should be enabled on the virtual machine <br><br>Minimum api-version: 2020-12-01",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.UefiSettings.SecureBootEnabled"),
			},
			{
				Name:        "security_profile_uefi_settings_v_tpm_enabled",
				Description: "Specifies whether vTPM should be enabled on the virtual machine <br><br>Minimum api-version: 2020-12-01",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.UefiSettings.VTpmEnabled"),
			},
			{
				Name:        "security_profile_encryption_at_host",
				Description: "This property can be used by user in the request to enable or disable the Host Encryption for the virtual machine or virtual machine scale set This will enable the encryption for all the disks including Resource/Temp disk at host itself <br><br> Default: The Encryption at host will be disabled unless this property is set to true for the resource",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.EncryptionAtHost"),
			},
			{
				Name:        "security_profile_security_type",
				Description: "Specifies the SecurityType of the virtual machine It is set as TrustedLaunch to enable UefiSettings <br><br> Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch Possible values include: 'SecurityTypesTrustedLaunch'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.SecurityType"),
			},
			{
				Name:        "diagnostics_profile_boot_diagnostics_enabled",
				Description: "Whether boot diagnostics should be enabled on the Virtual Machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.Enabled"),
			},
			{
				Name:        "diagnostics_profile_boot_diagnostics_storage_uri",
				Description: "Uri of the storage account to use for placing the console output and screenshot <br><br>If storageUri is not specified while enabling boot diagnostics, managed storage will be used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.StorageURI"),
			},
			{
				Name:        "availability_set_id",
				Description: "Availability set id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.AvailabilitySet.ID"),
			},
			{
				Name:        "virtual_machine_scale_set_id",
				Description: "Virtual machine scale set id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.VirtualMachineScaleSet.ID"),
			},
			{
				Name:        "proximity_placement_group_id",
				Description: "Proximity placement group resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ProximityPlacementGroup.ID"),
			},
			{
				Name:        "priority",
				Description: "Specifies the priority for the virtual machine <br><br>Minimum api-version: 2019-03-01 Possible values include: 'Regular', 'Low', 'Spot'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.Priority"),
			},
			{
				Name:        "eviction_policy",
				Description: "Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01 <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview Possible values include: 'Deallocate', 'Delete'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.EvictionPolicy"),
			},
			{
				Name:        "billing_profile_max_price",
				Description: "Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS This price is in US Dollars <br><br> This price will be compared with the current Azure Spot price for the VM size Also, the prices are compared at the time of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current Azure Spot price <br><br> The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the maxPrice after creation of VM/VMSS <br><br> Possible values are: <br><br> - Any decimal value greater than zero Example: 001538 <br><br> -1 – indicates default price to be up-to on-demand <br><br> You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons Also, the default max price is -1 if it is not provided by you <br><br>Minimum api-version: 2019-03-01",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("VirtualMachineProperties.BillingProfile.MaxPrice"),
			},
			{
				Name:        "host_id",
				Description: "Host Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.Host.ID"),
			},
			{
				Name:        "host_group_id",
				Description: "Host group Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.HostGroup.ID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state, which only appears in the response",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ProvisioningState"),
			},
			{
				Name:        "instance_view",
				Description: "The virtual machine instance view",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeVirtualMachineInstanceView,
			},
			{
				Name:        "license_type",
				Description: "Specifies that the image or disk that is being used was licensed on-premises <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docsmicrosoftcom/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docsmicrosoftcom/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.LicenseType"),
			},
			{
				Name:        "vm_id",
				Description: "Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.VMID"),
			},
			{
				Name:        "extensions_time_budget",
				Description: "Specifies the time alloted for all extensions to start The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format The default value is 90 minutes (PT1H30M) <br><br> Minimum api-version: 2020-06-01",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ExtensionsTimeBudget"),
			},
			{
				Name:        "platform_fault_domain",
				Description: "Specifies the scale set logical fault domain into which the Virtual Machine will be created.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VirtualMachineProperties.PlatformFaultDomain"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal id of virtual machine identity This property will only be provided for a system assigned identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant id associated with the virtual machine This property will only be provided for a system assigned identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the virtual machine The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities The type 'None' will remove any identities from the virtual machine Possible values include: 'ResourceIdentityTypeSystemAssigned', 'ResourceIdentityTypeUserAssigned', 'ResourceIdentityTypeSystemAssignedUserAssigned', 'ResourceIdentityTypeNone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user identities associated with the Virtual Machine The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}'",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "zones",
				Description: "The virtual machine zones",
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
				Description: "The type of the extended location Possible values include: 'EdgeZone'",
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
				Name:        "azure_compute_virtual_machine_win_config_rm_listeners",
				Description: "WinRMListener describes Protocol and thumbprint of Windows Remote Management listener",
				Resolver:    fetchComputeVirtualMachineWinConfigRmListeners,
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
						Description: "Specifies the protocol of WinRM listener <br><br> Possible values are: <br>**http** <br><br> **https** Possible values include: 'HTTP', 'HTTPS'",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_url",
						Description: "This is the URL of a certificate that has been uploaded to Key Vault as a secret For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docsmicrosoftcom/azure/key-vault/key-vault-get-started/#add) In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  \"data\":\"<Base64-encoded-certificate>\",<br>  \"dataType\":\"pfx\",<br>  \"password\":\"<pfx-file-password>\"<br>}",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CertificateURL"),
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_secrets",
				Description: "VaultSecretGroup describes a set of certificates which are all in the same Key Vault",
				Resolver:    fetchComputeVirtualMachineSecrets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_machine_cq_id", "source_vault_id"}},
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
						Description: "Source vault Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceVault.ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_compute_virtual_machine_secret_vault_certificates",
						Description: "VaultCertificate describes a single certificate reference in a Key Vault, and where the certificate should reside on the VM",
						Resolver:    fetchComputeVirtualMachineSecretVaultCertificates,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_machine_secret_cq_id", "certificate_url"}},
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_secret_cq_id",
								Description: "Unique CloudQuery ID of azure_compute_virtual_machine_secrets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "certificate_url",
								Description: "This is the URL of a certificate that has been uploaded to Key Vault as a secret For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docsmicrosoftcom/azure/key-vault/key-vault-get-started/#add) In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  \"data\":\"<Base64-encoded-certificate>\",<br>  \"dataType\":\"pfx\",<br>  \"password\":\"<pfx-file-password>\"<br>}",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CertificateURL"),
							},
							{
								Name:        "certificate_store",
								Description: "UppercaseThumbprint&gt;crt for the X509 certificate file and &lt;UppercaseThumbprint&gt;prv for private key Both of these files are pem formatted",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_resources",
				Description: "VirtualMachineExtension describes a Virtual Machine Extension",
				Resolver:    fetchComputeVirtualMachineResources,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_machine_cq_id", "id"}},
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
						Name:     "virtual_machine_extension_properties",
						Type:     schema.TypeJSON,
						Resolver: resolveComputeVirtualMachineResourceVirtualMachineExtensionProperties,
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
			},
			{
				Name:        "azure_compute_virtual_machine_network_interfaces",
				Description: "NetworkInterfaceReference describes a network interface reference",
				Resolver:    fetchComputeVirtualMachineNetworkInterfaces,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_machine_cq_id", "id"}},
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
						Name:        "network_interface_reference_properties_primary",
						Description: "Specifies the primary network interface in case the virtual machine has more than 1 network interface",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("NetworkInterfaceReferenceProperties.Primary"),
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
func fetchComputeVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachines
	response, err := svc.ListAll(ctx, "false")
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func resolveComputeVirtualMachineStorageProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", resource.Item)
	}

	if p.StorageProfile == nil {
		return nil
	}
	data, err := json.Marshal(p.StorageProfile)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveComputeVirtualMachineWindowsConfigurationAdditionalUnattendContent(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", resource.Item)
	}

	if p.OsProfile == nil ||
		p.OsProfile.WindowsConfiguration == nil ||
		p.OsProfile.WindowsConfiguration.AdditionalUnattendContent == nil {
		return nil
	}
	data, err := json.Marshal(p.OsProfile.WindowsConfiguration.AdditionalUnattendContent)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveComputeVirtualMachineLinuxConfigurationSSHPublicKeys(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", resource.Item)
	}

	if p.OsProfile == nil ||
		p.OsProfile.LinuxConfiguration == nil ||
		p.OsProfile.LinuxConfiguration.SSH == nil ||
		p.OsProfile.LinuxConfiguration.SSH.PublicKeys == nil {
		return nil
	}

	result := make(map[string]string)
	for _, ssh := range *p.OsProfile.LinuxConfiguration.SSH.PublicKeys {
		result[*ssh.Path] = *ssh.KeyData
	}

	return resource.Set(c.Name, result)
}
func resolveComputeVirtualMachineNetworkProfileNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", resource.Item)
	}

	if p.NetworkProfile == nil || p.NetworkProfile.NetworkInterfaces == nil {
		return nil
	}
	data, err := json.Marshal(p.NetworkProfile.NetworkInterfaces)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func resolveComputeVirtualMachineInstanceView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", resource.Item)
	}

	if p.InstanceView == nil {
		return nil
	}
	data, err := json.Marshal(p.InstanceView)

	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func fetchComputeVirtualMachineWinConfigRmListeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}
	if p.OsProfile == nil ||
		p.OsProfile.WindowsConfiguration == nil ||
		p.OsProfile.WindowsConfiguration.WinRM == nil ||
		p.OsProfile.WindowsConfiguration.WinRM.Listeners == nil {
		return nil
	}

	res <- *p.OsProfile.WindowsConfiguration.WinRM.Listeners
	return nil
}
func fetchComputeVirtualMachineSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.OsProfile == nil || p.OsProfile.Secrets == nil {
		return nil
	}

	res <- *p.OsProfile.Secrets
	return nil
}
func fetchComputeVirtualMachineSecretVaultCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VaultSecretGroup)
	if !ok {
		return fmt.Errorf("expected to have compute.VaultSecretGroup but got %T", parent.Item)
	}

	if p.VaultCertificates == nil {
		return nil
	}

	res <- *p.VaultCertificates
	return nil
}
func fetchComputeVirtualMachineResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachineExtensions
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}
	details, err := client.ParseResourceID(*p.ID)
	if err != nil {
		return err
	}
	response, err := svc.List(ctx, details.ResourceGroup, *p.Name, "")
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	return nil
}
func resolveComputeVirtualMachineResourceVirtualMachineExtensionProperties(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachineExtension)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtension but got %T", resource.Item)
	}

	if p.VirtualMachineExtensionProperties == nil {
		return nil
	}
	data, err := json.Marshal(p.VirtualMachineExtensionProperties)

	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func fetchComputeVirtualMachineNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.NetworkProfile == nil ||
		p.VirtualMachineProperties.NetworkProfile.NetworkInterfaces == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.NetworkProfile.NetworkInterfaces
	return nil
}
