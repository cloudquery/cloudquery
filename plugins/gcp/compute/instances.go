package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Instance struct {
	_                                                   interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                                                  uint        `gorm:"primarykey"`
	ProjectID                                           string      `neo:"unique"`
	CanIpForward                                        bool
	ConfidentialInstanceConfigEnableConfidentialCompute bool
	CpuPlatform                                         string
	CreationTimestamp                                   string
	DeletionProtection                                  bool
	Description                                         string
	Disks                                               []*InstanceAttachedDisk `gorm:"constraint:OnDelete:CASCADE;"`
	DisplayDeviceEnableDisplay                          bool
	Fingerprint                                         string
	GuestAccelerators                                   []*InstanceAcceleratorConfig `gorm:"constraint:OnDelete:CASCADE;"`
	Hostname                                            string
	ResourceID                                          uint64 `neo:"unique"`
	Kind                                                string
	LabelFingerprint                                    string
	//Labels                            []*InstanceLabels `gorm:"constraint:OnDelete:CASCADE;"`
	LastStartTimestamp      string
	LastStopTimestamp       string
	LastSuspendedTimestamp  string
	MachineType             string
	Metadata                []*InstanceMetadataItem `gorm:"constraint:OnDelete:CASCADE;"`
	MinCpuPlatform          string
	Name                    string
	NetworkInterfaces       []*InstanceNetworkInterface `gorm:"constraint:OnDelete:CASCADE;"`
	PrivateIpv6GoogleAccess string

	ReservationAffinityConsumeReservationType string
	ReservationAffinityKey                    string
	ReservationAffinityValues                 []*InstanceReservationAffinityValue `gorm:"constraint:OnDelete:CASCADE;"`
	Policies                                  []*InstancePolicy                   `gorm:"constraint:OnDelete:CASCADE;"`

	SchedulingAutomaticRestart  *bool
	SchedulingMinNodeCpus       int64
	SchedulingNodeAffinities    []*InstanceSchedulingNodeAffinity `gorm:"constraint:OnDelete:CASCADE;"`
	SchedulingOnHostMaintenance string
	SchedulingPreemptible       bool

	SelfLink        string
	ServiceAccounts []*InstanceServiceAccount `gorm:"constraint:OnDelete:CASCADE;"`

	ShieldedInstanceConfigEnableIntegrityMonitoring      bool
	ShieldedInstanceConfigEnableSecureBoot               bool
	ShieldedInstanceConfigEnableVtpm                     bool
	ShieldedInstanceIntegrityPolicyUpdateAutoLearnPolicy bool

	StartRestricted bool
	Status          string
	StatusMessage   string
	Tags            []*InstanceTag `gorm:"constraint:OnDelete:CASCADE;"`
	Zone            string
}

func (Instance) TableName() string {
	return "gcp_compute_instances"
}

type InstanceAttachedDisk struct {
	ID                    uint   `gorm:"primarykey"`
	InstanceID            uint   `neo:"ignore"`
	ProjectID             string `gorm:"-"`
	AutoDelete            bool
	Boot                  bool
	DeviceName            string
	DiskSizeGb            int64
	GuestOsFeatures       []*InstanceGuestOsFeature `gorm:"constraint:OnDelete:CASCADE;"`
	Index                 int64
	InitializeDescription string
	InitializeDiskName    string
	InitializeDiskSizeGb  int64
	InitializeDiskType    string
	//InitializeLabels           []*InstanceAttachedDiskInitializeLabel `gorm:"constraint:OnDelete:CASCADE;"`
	InitializeOnUpdateAction string
	InitializePolicies       []*InstanceAttachedDiskInitializePolicy `gorm:"constraint:OnDelete:CASCADE;"`
	InitializeSourceImage    string
	InitializeSourceSnapshot string
	Interface                string
	Kind                     string
	Licenses                 []*InstanceAttachedDiskLicense `gorm:"constraint:OnDelete:CASCADE;"`
	Mode                     string
	Source                   string
	Type                     string
}

func (InstanceAttachedDisk) TableName() string {
	return "gcp_compute_instance_attached_disk"
}

type InstanceGuestOsFeature struct {
	ID                     uint   `gorm:"primarykey"`
	InstanceAttachedDiskID uint   `neo:"ignore"`
	ProjectID              string `gorm:"-"`
	Type                   string
}

func (InstanceGuestOsFeature) TableName() string {
	return "gcp_compute_instance_guest_os_features"
}

type InstanceAttachedDiskInitializePolicy struct {
	ID                     uint   `gorm:"primarykey"`
	InstanceAttachedDiskID uint   `neo:"ignore"`
	ProjectID              string `gorm:"-"`
	Value                  string
}

func (InstanceAttachedDiskInitializePolicy) TableName() string {
	return "gcp_compute_instance_attached_disk_initialize_policies"
}

type InstanceAttachedDiskLicense struct {
	ID                     uint   `gorm:"primarykey"`
	InstanceAttachedDiskID uint   `neo:"ignore"`
	ProjectID              string `gorm:"-"`
	Value                  string
}

func (InstanceAttachedDiskLicense) TableName() string {
	return "gcp_compute_instance_attached_disk_licenses"
}

type InstanceAcceleratorConfig struct {
	ID               uint   `gorm:"primarykey"`
	InstanceID       uint   `neo:"ignore"`
	ProjectID        string `gorm:"-"`
	AcceleratorCount int64
	AcceleratorType  string
}

func (InstanceAcceleratorConfig) TableName() string {
	return "gcp_compute_instance_accelerator_configs"
}

type InstanceMetadataItem struct {
	ID          uint   `gorm:"primarykey"`
	InstanceID  uint   `neo:"ignore"`
	ProjectID   string `gorm:"-"`
	Fingerprint string
	Key         string
	Value       *string
	Kind        string
}

func (InstanceMetadataItem) TableName() string {
	return "gcp_compute_instance_metadata_items"
}

type InstanceNetworkInterface struct {
	ID            uint                    `gorm:"primarykey"`
	InstanceID    uint                    `neo:"ignore"`
	ProjectID     string                  `gorm:"-"`
	AccessConfigs []*InstanceAccessConfig `gorm:"constraint:OnDelete:CASCADE;"`
	AliasIpRanges []*InstanceAliasIpRange `gorm:"constraint:OnDelete:CASCADE;"`
	Fingerprint   string
	Ipv6Address   string
	Kind          string
	Name          string
	Network       string
	NetworkIP     string
	Subnetwork    string
}

func (InstanceNetworkInterface) TableName() string {
	return "gcp_compute_instance_network_interfaces"
}

type InstanceAccessConfig struct {
	ID                         uint   `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint   `neo:"ignore"`
	ProjectID                  string `gorm:"-"`
	Kind                       string
	Name                       string
	NatIP                      string
	NetworkTier                string
	PublicPtrDomainName        string
	SetPublicPtr               bool
	Type                       string
}

func (InstanceAccessConfig) TableName() string {
	return "gcp_compute_instance_access_configs"
}

type InstanceAliasIpRange struct {
	ID                         uint   `gorm:"primarykey"`
	InstanceNetworkInterfaceID uint   `neo:"ignore"`
	ProjectID                  string `gorm:"-"`
	IpCidrRange                string
	SubnetworkRangeName        string
}

func (InstanceAliasIpRange) TableName() string {
	return "gcp_compute_instance_alias_ip_ranges"
}

type InstanceReservationAffinityValue struct {
	ID         uint   `gorm:"primarykey"`
	InstanceID uint   `neo:"ignore"`
	ProjectID  string `gorm:"-"`
	Value      string
}

func (InstanceReservationAffinityValue) TableName() string {
	return "gcp_compute_instance_reservation_affinity_values"
}

type InstancePolicy struct {
	ID         uint   `gorm:"primarykey"`
	InstanceID uint   `neo:"ignore"`
	ProjectID  string `gorm:"-"`
	Value      string
}

func (InstancePolicy) TableName() string {
	return "gcp_compute_instance_policy"
}

type InstanceSchedulingNodeAffinity struct {
	ID         uint   `gorm:"primarykey"`
	InstanceID uint   `neo:"ignore"`
	ProjectID  string `gorm:"-"`
	Key        string
	Operator   string
	Value      string
}

func (InstanceSchedulingNodeAffinity) TableName() string {
	return "gcp_compute_instance_scheduling_node_affinities"
}

type InstanceServiceAccount struct {
	ID         uint   `gorm:"primarykey"`
	InstanceID uint   `neo:"ignore"`
	ProjectID  string `gorm:"-"`
	Email      string
	Scope      string
}

func (InstanceServiceAccount) TableName() string {
	return "gcp_compute_instance_service_accounts"
}

type InstanceTag struct {
	ID          uint   `gorm:"primarykey"`
	InstanceID  uint   `neo:"ignore"`
	ProjectID   string `gorm:"-"`
	Fingerprint string
	Value       string `gorm:"constraint:OnDelete:CASCADE;"`
}

func (InstanceTag) TableName() string {
	return "gcp_compute_instance_tags"
}

func (c *Client) transformInstanceGuestOsFeatures(values []*compute.GuestOsFeature) []*InstanceGuestOsFeature {
	var tValues []*InstanceGuestOsFeature
	for _, v := range values {
		tValues = append(tValues, &InstanceGuestOsFeature{
			ProjectID: c.projectID,
			Type:      v.Type,
		})
	}
	return tValues
}

func (c *Client) transformInstanceAttachedDiskLicenses(values []string) []*InstanceAttachedDiskLicense {
	var tValues []*InstanceAttachedDiskLicense
	for _, v := range values {
		tValues = append(tValues, &InstanceAttachedDiskLicense{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformInstanceAttachedDisk(value *compute.AttachedDisk) *InstanceAttachedDisk {
	return &InstanceAttachedDisk{
		ProjectID:       c.projectID,
		AutoDelete:      value.AutoDelete,
		Boot:            value.Boot,
		DeviceName:      value.DeviceName,
		DiskSizeGb:      value.DiskSizeGb,
		GuestOsFeatures: c.transformInstanceGuestOsFeatures(value.GuestOsFeatures),
		Index:           value.Index,
		Interface:       value.Interface,
		Kind:            value.Kind,
		Licenses:        c.transformInstanceAttachedDiskLicenses(value.Licenses),
		Mode:            value.Mode,
		Source:          value.Source,
		Type:            value.Type,
	}
}

func (c *Client) transformInstanceAttachedDisks(values []*compute.AttachedDisk) []*InstanceAttachedDisk {
	var tValues []*InstanceAttachedDisk
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAttachedDisk(v))
	}
	return tValues
}

func (c *Client) transformInstanceAcceleratorConfig(value *compute.AcceleratorConfig) *InstanceAcceleratorConfig {
	return &InstanceAcceleratorConfig{
		ProjectID:        c.projectID,
		AcceleratorCount: value.AcceleratorCount,
		AcceleratorType:  value.AcceleratorType,
	}
}

func (c *Client) transformInstanceAcceleratorConfigs(values []*compute.AcceleratorConfig) []*InstanceAcceleratorConfig {
	var tValues []*InstanceAcceleratorConfig
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAcceleratorConfig(v))
	}
	return tValues
}

func (c *Client) transformInstanceMetadataItems(value *compute.Metadata) []*InstanceMetadataItem {
	var tValues []*InstanceMetadataItem
	for _, v := range value.Items {
		tValues = append(tValues, &InstanceMetadataItem{
			ProjectID:   c.projectID,
			Fingerprint: value.Fingerprint,
			Key:         v.Key,
			Value:       v.Value,
			Kind:        value.Kind,
		})
	}
	return tValues
}

func (c *Client) transformInstanceAccessConfig(value *compute.AccessConfig) *InstanceAccessConfig {
	return &InstanceAccessConfig{
		ProjectID:           c.projectID,
		Kind:                value.Kind,
		Name:                value.Name,
		NatIP:               value.NatIP,
		NetworkTier:         value.NetworkTier,
		PublicPtrDomainName: value.PublicPtrDomainName,
		SetPublicPtr:        value.SetPublicPtr,
		Type:                value.Type,
	}
}

func (c *Client) transformInstanceAccessConfigs(values []*compute.AccessConfig) []*InstanceAccessConfig {
	var tValues []*InstanceAccessConfig
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAccessConfig(v))
	}
	return tValues
}

func (c *Client) transformInstanceAliasIpRange(value *compute.AliasIpRange) *InstanceAliasIpRange {
	return &InstanceAliasIpRange{
		ProjectID:           c.projectID,
		IpCidrRange:         value.IpCidrRange,
		SubnetworkRangeName: value.SubnetworkRangeName,
	}
}

func (c *Client) transformInstanceAliasIpRanges(values []*compute.AliasIpRange) []*InstanceAliasIpRange {
	var tValues []*InstanceAliasIpRange
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceAliasIpRange(v))
	}
	return tValues
}

func (c *Client) transformInstanceNetworkInterface(value *compute.NetworkInterface) *InstanceNetworkInterface {
	return &InstanceNetworkInterface{
		ProjectID:     c.projectID,
		AccessConfigs: c.transformInstanceAccessConfigs(value.AccessConfigs),
		AliasIpRanges: c.transformInstanceAliasIpRanges(value.AliasIpRanges),
		Fingerprint:   value.Fingerprint,
		Ipv6Address:   value.Ipv6Address,
		Kind:          value.Kind,
		Name:          value.Name,
		Network:       value.Network,
		NetworkIP:     value.NetworkIP,
		Subnetwork:    value.Subnetwork,
	}
}

func (c *Client) transformInstanceNetworkInterfaces(values []*compute.NetworkInterface) []*InstanceNetworkInterface {
	var tValues []*InstanceNetworkInterface
	for _, v := range values {
		tValues = append(tValues, c.transformInstanceNetworkInterface(v))
	}
	return tValues
}

func (c *Client) transformInstanceReservationAffinityValues(values []string) []*InstanceReservationAffinityValue {
	var tValues []*InstanceReservationAffinityValue
	for _, v := range values {
		tValues = append(tValues, &InstanceReservationAffinityValue{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformInstanceResourcePolicies(values []string) []*InstancePolicy {
	var tValues []*InstancePolicy
	for _, v := range values {
		tValues = append(tValues, &InstancePolicy{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformInstanceSchedulingNodeAffinities(values []*compute.SchedulingNodeAffinity) []*InstanceSchedulingNodeAffinity {
	var tValues []*InstanceSchedulingNodeAffinity
	for _, v := range values {
		for _, v1 := range v.Values {
			tValues = append(tValues, &InstanceSchedulingNodeAffinity{
				ProjectID: c.projectID,
				Key:       v.Key,
				Operator:  v.Operator,
				Value:     v1,
			})
		}
	}
	return tValues
}

func (c *Client) transformInstanceServiceAccounts(values []*compute.ServiceAccount) []*InstanceServiceAccount {
	var tValues []*InstanceServiceAccount
	for _, v := range values {
		for _, scope := range v.Scopes {
			tValues = append(tValues, &InstanceServiceAccount{
				ProjectID: c.projectID,
				Email:     v.Email,
				Scope:     scope,
			})
		}
	}
	return tValues
}

func (c *Client) transformTags(value *compute.Tags) []*InstanceTag {
	var tValues []*InstanceTag
	for _, v := range value.Items {
		tValues = append(tValues, &InstanceTag{
			ProjectID:   c.projectID,
			Fingerprint: value.Fingerprint,
			Value:       v,
		})
	}
	return tValues
}

func (c *Client) transformMetadataItems(value *compute.Metadata) []*InstanceMetadataItem {
	var tValues []*InstanceMetadataItem
	for _, v := range value.Items {
		tValues = append(tValues, &InstanceMetadataItem{
			ProjectID:   c.projectID,
			Fingerprint: value.Fingerprint,
			Key:         v.Key,
			Value:       v.Value,
			Kind:        value.Kind,
		})
	}
	return tValues
}

func (c *Client) transformInstance(value *compute.Instance) *Instance {
	res := Instance{
		ProjectID:    c.projectID,
		CanIpForward: value.CanIpForward,

		CpuPlatform:             value.CpuPlatform,
		CreationTimestamp:       value.CreationTimestamp,
		DeletionProtection:      value.DeletionProtection,
		Description:             value.Description,
		Disks:                   c.transformInstanceAttachedDisks(value.Disks),
		Fingerprint:             value.Fingerprint,
		GuestAccelerators:       c.transformInstanceAcceleratorConfigs(value.GuestAccelerators),
		Hostname:                value.Hostname,
		ResourceID:              value.Id,
		Kind:                    value.Kind,
		LabelFingerprint:        value.LabelFingerprint,
		LastStartTimestamp:      value.LastStartTimestamp,
		LastStopTimestamp:       value.LastStopTimestamp,
		LastSuspendedTimestamp:  value.LastSuspendedTimestamp,
		MachineType:             value.MachineType,
		Metadata:                c.transformMetadataItems(value.Metadata),
		MinCpuPlatform:          value.MinCpuPlatform,
		Name:                    value.Name,
		NetworkInterfaces:       c.transformInstanceNetworkInterfaces(value.NetworkInterfaces),
		PrivateIpv6GoogleAccess: value.PrivateIpv6GoogleAccess,

		Policies: c.transformInstanceResourcePolicies(value.ResourcePolicies),

		SelfLink:        value.SelfLink,
		ServiceAccounts: c.transformInstanceServiceAccounts(value.ServiceAccounts),

		StartRestricted: value.StartRestricted,
		Status:          value.Status,
		StatusMessage:   value.StatusMessage,
		Tags:            c.transformTags(value.Tags),
		Zone:            value.Zone,
	}

	if value.DisplayDevice != nil {
		res.DisplayDeviceEnableDisplay = value.DisplayDevice.EnableDisplay
	}

	if value.ConfidentialInstanceConfig != nil {
		res.ConfidentialInstanceConfigEnableConfidentialCompute = value.ConfidentialInstanceConfig.EnableConfidentialCompute
	}

	if value.ReservationAffinity != nil {
		res.ReservationAffinityKey = value.ReservationAffinity.Key
		res.ReservationAffinityConsumeReservationType = value.ReservationAffinity.ConsumeReservationType
		res.ReservationAffinityValues = c.transformInstanceReservationAffinityValues(value.ReservationAffinity.Values)
	}
	if value.Scheduling != nil {
		res.SchedulingAutomaticRestart = value.Scheduling.AutomaticRestart
		res.SchedulingMinNodeCpus = value.Scheduling.MinNodeCpus
		res.SchedulingOnHostMaintenance = value.Scheduling.OnHostMaintenance
		res.SchedulingNodeAffinities = c.transformInstanceSchedulingNodeAffinities(value.Scheduling.NodeAffinities)
	}

	if value.ShieldedInstanceConfig != nil {
		res.ShieldedInstanceConfigEnableIntegrityMonitoring = value.ShieldedInstanceConfig.EnableIntegrityMonitoring
		res.ShieldedInstanceConfigEnableSecureBoot = value.ShieldedInstanceConfig.EnableSecureBoot
		res.ShieldedInstanceConfigEnableVtpm = value.ShieldedInstanceConfig.EnableVtpm
		res.ShieldedInstanceIntegrityPolicyUpdateAutoLearnPolicy = value.ShieldedInstanceIntegrityPolicy.UpdateAutoLearnPolicy
	}

	return &res
}

func (c *Client) transformInstances(values []*compute.Instance) []*Instance {
	var tValues []*Instance
	for _, v := range values {
		tValues = append(tValues, c.transformInstance(v))
	}
	return tValues
}

type InstanceConfig struct {
	Filter string
}

var InstanceTables = []interface{}{
	&Instance{},
	&InstanceAttachedDisk{},
	&InstanceAttachedDiskLicense{},
	&InstanceGuestOsFeature{},
	&InstanceAcceleratorConfig{},
	&InstanceMetadataItem{},
	&InstanceNetworkInterface{},
	&InstanceAccessConfig{},
	&InstanceAliasIpRange{},
	&InstanceSchedulingNodeAffinity{},
	&InstanceServiceAccount{},
	&InstanceTag{},
}

func (c *Client) instances(gConfig interface{}) error {
	var config InstanceConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	nextPageToken := ""
	for {
		call := c.svc.Instances.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id", c.projectID).Delete(InstanceTables...)
		var tValues []*Instance
		for _, items := range output.Items {
			tValues = append(tValues, c.transformInstances(items.Instances)...)
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.instances"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
