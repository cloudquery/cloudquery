package main

import (
	"context"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	ID          uint   `gorm:"primarykey"`
	ClusterName string `neo:"unique"`
	Name        string `neo:"unique"`
	Namespace   string `neo:"unique"`

	SpecVolumes                       []*PodVolume             `gorm:"constraint:OnDelete:CASCADE;"`
	SpecContainers                    []*PodContainer          `gorm:"constraint:OnDelete:CASCADE;"`
	SpecEphemeralContainers           []*PodEphemeralContainer `gorm:"constraint:OnDelete:CASCADE;"`
	SpecRestartPolicy                 string
	SpecTerminationGracePeriodSeconds *int64
	SpecActiveDeadlineSeconds         *int64
	SpecDNSPolicy                     string
	SpecServiceAccountName            string
	SpecAutomountServiceAccountToken  *bool
	SpecNodeName                      string
	SpecHostNetwork                   bool
	SpecHostPID                       bool
	SpecHostIPC                       bool
	SpecShareProcessNamespace         *bool

	SpecSecurityContextSELinuxOptionsUser  string
	SpecSecurityContextSELinuxOptionsRole  string
	SpecSecurityContextSELinuxOptionsType  string
	SpecSecurityContextSELinuxOptionsLevel string

	SpecSecurityContextWindowsOptionsGMSACredentialSpecName *string
	SpecSecurityContextWindowsOptionsGMSACredentialSpec     *string
	SpecSecurityContextWindowsOptionsRunAsUserName          *string
	SpecSecurityContextRunAsUser                            *int64
	SpecSecurityContextRunAsGroup                           *int64
	SpecSecurityContextRunAsNonRoot                         *bool
	SpecSecurityContextFSGroup                              *int64
	SpecSecurityContextSysctls                              []*PodSysctl `gorm:"constraint:OnDelete:CASCADE;"`
	SpecSecurityContextFSGroupChangePolicy                  *v1.PodFSGroupChangePolicy

	SpecSecurityContextSeccompProfileType             string
	SpecSecurityContextSeccompProfileLocalhostProfile *string
	SpecHostname                                      string
	SpecSubdomain                                     string

	SpecSchedulerName     string
	SpecTolerations       []*PodToleration `gorm:"constraint:OnDelete:CASCADE;"`
	SpecHostAliases       []*PodHostAlias  `gorm:"constraint:OnDelete:CASCADE;"`
	SpecPriorityClassName string
	SpecPriority          *int32

	SpecDNSConfigNameservers []*PodDNSConfigNameservers `gorm:"constraint:OnDelete:CASCADE;"`
	SpecDNSConfigSearches    []*PodDNSConfigSearches    `gorm:"constraint:OnDelete:CASCADE;"`
	SpecDNSConfigOptions     []*PodDNSConfigOption      `gorm:"constraint:OnDelete:CASCADE;"`
	SpecReadinessGates       []*PodReadinessGate        `gorm:"constraint:OnDelete:CASCADE;"`
	SpecRuntimeClassName     *string
	SpecEnableServiceLinks   *bool
	SpecPreemptionPolicy     *v1.PreemptionPolicy
	SpecSetHostnameAsFQDN    *bool

	StatusPhase             string
	StatusConditions        []*PodCondition `gorm:"constraint:OnDelete:CASCADE;"`
	StatusMessage           string
	StatusReason            string
	StatusNominatedNodeName string
	StatusHostIP            string
	StatusPodIP             string
	StatusPodIPs            []*PodIP `gorm:"constraint:OnDelete:CASCADE;"`

	StatusContainerStatuses []*PodContainerStatus `gorm:"constraint:OnDelete:CASCADE;"`
	StatusQOSClass          string
}

func (Pod) TableName() string {
	return "k8s_pods"
}

type PodVolume struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Name        string
}

func (PodVolume) TableName() string {
	return "k8s_pod_volumes"
}

type PodContainerCommand struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Value          string
}

func (PodContainerCommand) TableName() string {
	return "k8s_pod_container_commands"
}

type PodContainerArgs struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Value          string
}

func (PodContainerArgs) TableName() string {
	return "k8s_pod_container_args"
}

type PodContainerPort struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Name           string
	HostPort       int32
	ContainerPort  int32
	Protocol       string
	HostIP         string
}

func (PodContainerPort) TableName() string {
	return "k8s_pod_container_ports"
}

type PodEnvFromSource struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Prefix         string

	ConfigMapRefOptional *bool

	SecretRefOptional *bool
}

func (PodEnvFromSource) TableName() string {
	return "k8s_pod_env_from_sources"
}

type PodEnvVar struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Name           string
	Value          string

	ValueFromFieldRefAPIVersion string
	ValueFromFieldRefFieldPath  string

	ValueFromResourceFieldRefContainerName string
	ValueFromResourceFieldRefResource      string

	ValueFromResourceFieldRefDivisorivalue int64
	ValueFromResourceFieldRefDivisoriscale int32

	ValueFromResourceFieldRefDivisors string

	ValueFromConfigMapKeyRefKey      string
	ValueFromConfigMapKeyRefOptional *bool

	ValueFromSecretKeyRefKey      string
	ValueFromSecretKeyRefOptional *bool
}

func (PodEnvVar) TableName() string {
	return "k8s_pod_env_vars"
}

type PodVolumeMount struct {
	ID               uint   `gorm:"primarykey"`
	PodContainerID   uint   `neo:"ignore"`
	ClusterName      string `gorm:"-"`
	Name             string
	ReadOnly         bool
	MountPath        string
	SubPath          string
	MountPropagation *v1.MountPropagationMode
	SubPathExpr      string
}

func (PodVolumeMount) TableName() string {
	return "k8s_pod_volume_mounts"
}

type PodVolumeDevice struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Name           string
	DevicePath     string
}

func (PodVolumeDevice) TableName() string {
	return "k8s_pod_volume_devices"
}

type PodExecActionCommand struct {
	ID              uint   `gorm:"primarykey"`
	PodExecActionID uint   `neo:"ignore"`
	ClusterName     string `gorm:"-"`
	Value           string
}

func (PodExecActionCommand) TableName() string {
	return "k8s_pod_exec_action_commands"
}

type PodHTTPHeader struct {
	ID                 uint   `gorm:"primarykey"`
	PodHTTPGetActionID uint   `neo:"ignore"`
	ClusterName        string `gorm:"-"`
	Name               string
	Value              string
}

func (PodHTTPHeader) TableName() string {
	return "k8s_pod_http_headers"
}

type PodCapabilitiesAdd struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Value          string
}

func (PodCapabilitiesAdd) TableName() string {
	return "k8s_pod_capabilities_add"
}

type PodCapabilitiesDrop struct {
	ID             uint   `gorm:"primarykey"`
	PodContainerID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Value          string
}

func (PodCapabilitiesDrop) TableName() string {
	return "k8s_pod_capabilities_drop"
}

type PodContainer struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Name        string
	Image       string
	Command     []*PodContainerCommand `gorm:"constraint:OnDelete:CASCADE;"`
	Args        []*PodContainerArgs    `gorm:"constraint:OnDelete:CASCADE;"`
	WorkingDir  string
	Ports       []*PodContainerPort `gorm:"constraint:OnDelete:CASCADE;"`
	EnvFrom     []*PodEnvFromSource `gorm:"constraint:OnDelete:CASCADE;"`
	Env         []*PodEnvVar        `gorm:"constraint:OnDelete:CASCADE;"`

	VolumeMounts  []*PodVolumeMount  `gorm:"constraint:OnDelete:CASCADE;"`
	VolumeDevices []*PodVolumeDevice `gorm:"constraint:OnDelete:CASCADE;"`

	LivenessProbeInitialDelaySeconds int32
	LivenessProbeTimeoutSeconds      int32
	LivenessProbePeriodSeconds       int32
	LivenessProbeSuccessThreshold    int32
	LivenessProbeFailureThreshold    int32

	ReadinessProbeInitialDelaySeconds int32
	ReadinessProbeTimeoutSeconds      int32
	ReadinessProbePeriodSeconds       int32
	ReadinessProbeSuccessThreshold    int32
	ReadinessProbeFailureThreshold    int32

	StartupProbeInitialDelaySeconds int32
	StartupProbeTimeoutSeconds      int32
	StartupProbePeriodSeconds       int32
	StartupProbeSuccessThreshold    int32
	StartupProbeFailureThreshold    int32

	LifecyclePostStartHTTPGetPath   string
	LifecyclePostStartHTTPGetPort   string
	LifecyclePostStartHTTPGetHost   string
	LifecyclePostStartHTTPGetScheme string

	LifecyclePostStartTCPSocketPort string
	LifecyclePostStartTCPSocketHost string

	LifecyclePreStopHTTPGetPath   string
	LifecyclePreStopHTTPGetPort   string
	LifecyclePreStopHTTPGetHost   string
	LifecyclePreStopHTTPGetScheme string

	LifecyclePreStopTCPSocketPort string
	LifecyclePreStopTCPSocketHost string
	TerminationMessagePath        string
	TerminationMessagePolicy      string
	ImagePullPolicy               string

	SecurityContextCapabilitiesAdd  []*PodCapabilitiesAdd  `gorm:"constraint:OnDelete:CASCADE;"`
	SecurityContextCapabilitiesDrop []*PodCapabilitiesDrop `gorm:"constraint:OnDelete:CASCADE;"`
	SecurityContextPrivileged       *bool

	SecurityContextSELinuxOptionsUser  string
	SecurityContextSELinuxOptionsRole  string
	SecurityContextSELinuxOptionsType  string
	SecurityContextSELinuxOptionsLevel string

	SecurityContextWindowsOptionsGMSACredentialSpecName *string
	SecurityContextWindowsOptionsGMSACredentialSpec     *string
	SecurityContextWindowsOptionsRunAsUserName          *string
	SecurityContextRunAsUser                            *int64
	SecurityContextRunAsGroup                           *int64
	SecurityContextRunAsNonRoot                         *bool
	SecurityContextReadOnlyRootFilesystem               *bool
	SecurityContextAllowPrivilegeEscalation             *bool
	SecurityContextProcMount                            *v1.ProcMountType

	SecurityContextSeccompProfileType             string
	SecurityContextSeccompProfileLocalhostProfile *string
	Stdin                                         bool
	StdinOnce                                     bool
	TTY                                           bool
}

func (PodContainer) TableName() string {
	return "k8s_pod_containers"
}

type PodEphemeralContainer struct {
	ID                  uint   `gorm:"primarykey"`
	PodID               uint   `neo:"ignore"`
	ClusterName         string `gorm:"-"`
	TargetContainerName string
}

func (PodEphemeralContainer) TableName() string {
	return "k8s_pod_ephemeral_containers"
}

type PodSysctl struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Name        string
	Value       string
}

func (PodSysctl) TableName() string {
	return "k8s_pod_sysctls"
}

type PodAffinityTermNamespaces struct {
	ID                uint   `gorm:"primarykey"`
	PodAffinityTermID uint   `neo:"ignore"`
	ClusterName       string `gorm:"-"`
	Value             string
}

func (PodAffinityTermNamespaces) TableName() string {
	return "k8s_pod_affinity_term_namespaces"
}

type PodToleration struct {
	ID                uint   `gorm:"primarykey"`
	PodID             uint   `neo:"ignore"`
	ClusterName       string `gorm:"-"`
	Key               string
	Operator          string
	Value             string
	Effect            string
	TolerationSeconds *int64
}

func (PodToleration) TableName() string {
	return "k8s_pod_tolerations"
}

type PodHostAliasHostnames struct {
	ID             uint   `gorm:"primarykey"`
	PodHostAliasID uint   `neo:"ignore"`
	ClusterName    string `gorm:"-"`
	Value          string
}

func (PodHostAliasHostnames) TableName() string {
	return "k8s_pod_host_alias_hostnames"
}

type PodHostAlias struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	IP          string
	Hostnames   []*PodHostAliasHostnames `gorm:"constraint:OnDelete:CASCADE;"`
}

func (PodHostAlias) TableName() string {
	return "k8s_pod_host_aliass"
}

type PodDNSConfigNameservers struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Value       string
}

func (PodDNSConfigNameservers) TableName() string {
	return "k8s_pod_dns_config_nameservers"
}

type PodDNSConfigSearches struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Value       string
}

func (PodDNSConfigSearches) TableName() string {
	return "k8s_pod_dns_config_searches"
}

type PodDNSConfigOption struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Name        string
	Value       *string
}

func (PodDNSConfigOption) TableName() string {
	return "k8s_pod_dns_config_options"
}

type PodReadinessGate struct {
	ID            uint   `gorm:"primarykey"`
	PodID         uint   `neo:"ignore"`
	ClusterName   string `gorm:"-"`
	ConditionType string
}

func (PodReadinessGate) TableName() string {
	return "k8s_pod_readiness_gates"
}

type PodCondition struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Type        string
	Status      string

	Reason  string
	Message string
}

func (PodCondition) TableName() string {
	return "k8s_pod_conditions"
}

type PodIP struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	IP          string
}

func (PodIP) TableName() string {
	return "k8s_pod_ips"
}

type PodContainerStatus struct {
	ID          uint   `gorm:"primarykey"`
	PodID       uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`
	Name        string

	StateWaitingReason  string
	StateWaitingMessage string

	StateTerminatedExitCode int32
	StateTerminatedSignal   int32
	StateTerminatedReason   string
	StateTerminatedMessage  string

	StateTerminatedContainerID string

	LastTerminationStateWaitingReason  string
	LastTerminationStateWaitingMessage string

	LastTerminationStateTerminatedExitCode int32
	LastTerminationStateTerminatedSignal   int32
	LastTerminationStateTerminatedReason   string
	LastTerminationStateTerminatedMessage  string

	LastTerminationStateTerminatedContainerID string
	Ready                                     bool
	RestartCount                              int32
	Image                                     string
	ImageID                                   string
	ContainerID                               string
	Started                                   *bool
}

func (PodContainerStatus) TableName() string {
	return "k8s_pod_container_statuses"
}

func (p *Provider) transformPods(values []v1.Pod) []*Pod {
	var tValues []*Pod
	for _, value := range values {
		tValue := Pod{
			ClusterName:                       p.clusterName,
			Name:                              value.Name,
			Namespace:                         value.Namespace,
			SpecVolumes:                       p.transformPodVolumes(value.Spec.Volumes),
			SpecContainers:                    p.transformPodContainers(value.Spec.Containers),
			SpecEphemeralContainers:           p.transformPodEphemeralContainers(value.Spec.EphemeralContainers),
			SpecRestartPolicy:                 string(value.Spec.RestartPolicy),
			SpecTerminationGracePeriodSeconds: value.Spec.TerminationGracePeriodSeconds,
			SpecActiveDeadlineSeconds:         value.Spec.ActiveDeadlineSeconds,
			SpecDNSPolicy:                     string(value.Spec.DNSPolicy),
			SpecServiceAccountName:            value.Spec.ServiceAccountName,
			SpecAutomountServiceAccountToken:  value.Spec.AutomountServiceAccountToken,
			SpecNodeName:                      value.Spec.NodeName,
			SpecHostNetwork:                   value.Spec.HostNetwork,
			SpecHostPID:                       value.Spec.HostPID,
			SpecHostIPC:                       value.Spec.HostIPC,
			SpecShareProcessNamespace:         value.Spec.ShareProcessNamespace,

			SpecHostname:           value.Spec.Hostname,
			SpecSubdomain:          value.Spec.Subdomain,
			SpecSchedulerName:      value.Spec.SchedulerName,
			SpecTolerations:        p.transformPodTolerations(value.Spec.Tolerations),
			SpecHostAliases:        p.transformPodHostAliass(value.Spec.HostAliases),
			SpecPriorityClassName:  value.Spec.PriorityClassName,
			SpecPriority:           value.Spec.Priority,
			SpecReadinessGates:     p.transformPodReadinessGates(value.Spec.ReadinessGates),
			SpecRuntimeClassName:   value.Spec.RuntimeClassName,
			SpecEnableServiceLinks: value.Spec.EnableServiceLinks,
			SpecPreemptionPolicy:   value.Spec.PreemptionPolicy,
			SpecSetHostnameAsFQDN:  value.Spec.SetHostnameAsFQDN,

			StatusPhase:             string(value.Status.Phase),
			StatusConditions:        p.transformPodConditions(value.Status.Conditions),
			StatusMessage:           value.Status.Message,
			StatusReason:            value.Status.Reason,
			StatusNominatedNodeName: value.Status.NominatedNodeName,
			StatusHostIP:            value.Status.HostIP,
			StatusPodIP:             value.Status.PodIP,
			StatusPodIPs:            p.transformPodIPs(value.Status.PodIPs),
			StatusContainerStatuses: p.transformPodContainerStatuss(value.Status.ContainerStatuses),
			StatusQOSClass:          string(value.Status.QOSClass),
		}

		if value.Spec.DNSConfig != nil {
			tValue.SpecDNSConfigNameservers = p.transformPodDNSConfigNameservers(value.Spec.DNSConfig.Nameservers)
			tValue.SpecDNSConfigSearches = p.transformPodDNSConfigSearches(value.Spec.DNSConfig.Searches)
			tValue.SpecDNSConfigOptions = p.transformPodDNSConfigOptions(value.Spec.DNSConfig.Options)
		}

		if value.Spec.SecurityContext != nil {
			if value.Spec.SecurityContext.SELinuxOptions != nil {
				tValue.SpecSecurityContextSELinuxOptionsUser = value.Spec.SecurityContext.SELinuxOptions.User
				tValue.SpecSecurityContextSELinuxOptionsRole = value.Spec.SecurityContext.SELinuxOptions.Role
				tValue.SpecSecurityContextSELinuxOptionsType = value.Spec.SecurityContext.SELinuxOptions.Type
				tValue.SpecSecurityContextSELinuxOptionsLevel = value.Spec.SecurityContext.SELinuxOptions.Level
			}

			if value.Spec.SecurityContext.WindowsOptions != nil {
				tValue.SpecSecurityContextWindowsOptionsGMSACredentialSpecName = value.Spec.SecurityContext.WindowsOptions.GMSACredentialSpecName
				tValue.SpecSecurityContextWindowsOptionsGMSACredentialSpec = value.Spec.SecurityContext.WindowsOptions.GMSACredentialSpec
				tValue.SpecSecurityContextWindowsOptionsRunAsUserName = value.Spec.SecurityContext.WindowsOptions.RunAsUserName
			}

			if value.Spec.SecurityContext.SeccompProfile != nil {
				tValue.SpecSecurityContextSeccompProfileType = string(value.Spec.SecurityContext.SeccompProfile.Type)
				tValue.SpecSecurityContextSeccompProfileLocalhostProfile = value.Spec.SecurityContext.SeccompProfile.LocalhostProfile
			}

			tValue.SpecSecurityContextRunAsUser = value.Spec.SecurityContext.RunAsUser
			tValue.SpecSecurityContextRunAsGroup = value.Spec.SecurityContext.RunAsGroup
			tValue.SpecSecurityContextRunAsNonRoot = value.Spec.SecurityContext.RunAsNonRoot
			tValue.SpecSecurityContextFSGroup = value.Spec.SecurityContext.FSGroup
			tValue.SpecSecurityContextSysctls = p.transformPodSysctls(value.Spec.SecurityContext.Sysctls)
			tValue.SpecSecurityContextFSGroupChangePolicy = value.Spec.SecurityContext.FSGroupChangePolicy

		}

		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodVolumes(values []v1.Volume) []*PodVolume {
	var tValues []*PodVolume
	for _, value := range values {
		tValue := PodVolume{
			ClusterName: p.clusterName,
			Name:        value.Name,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (p *Provider) transformPodContainerCommand(values []string) []*PodContainerCommand {
	var tValues []*PodContainerCommand
	for _, v := range values {
		tValues = append(tValues, &PodContainerCommand{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodContainerArgs(values []string) []*PodContainerArgs {
	var tValues []*PodContainerArgs
	for _, v := range values {
		tValues = append(tValues, &PodContainerArgs{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodContainerPorts(values []v1.ContainerPort) []*PodContainerPort {
	var tValues []*PodContainerPort
	for _, value := range values {
		tValue := PodContainerPort{
			ClusterName:   p.clusterName,
			Name:          value.Name,
			HostPort:      value.HostPort,
			ContainerPort: value.ContainerPort,
			Protocol:      string(value.Protocol),
			HostIP:        value.HostIP,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodEnvFromSources(values []v1.EnvFromSource) []*PodEnvFromSource {
	var tValues []*PodEnvFromSource
	for _, value := range values {
		tValue := PodEnvFromSource{
			ClusterName: p.clusterName,
			Prefix:      value.Prefix,
		}
		if value.ConfigMapRef != nil {
			tValue.ConfigMapRefOptional = value.ConfigMapRef.Optional
		}
		if value.SecretRef != nil {
			tValue.SecretRefOptional = value.SecretRef.Optional
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodEnvVars(values []v1.EnvVar) []*PodEnvVar {
	var tValues []*PodEnvVar
	for _, value := range values {
		tValue := PodEnvVar{
			ClusterName: p.clusterName,
			Name:        value.Name,
			Value:       value.Value,
		}
		if value.ValueFrom != nil {

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodVolumeMounts(values []v1.VolumeMount) []*PodVolumeMount {
	var tValues []*PodVolumeMount
	for _, value := range values {
		tValue := PodVolumeMount{
			ClusterName:      p.clusterName,
			Name:             value.Name,
			ReadOnly:         value.ReadOnly,
			MountPath:        value.MountPath,
			SubPath:          value.SubPath,
			MountPropagation: value.MountPropagation,
			SubPathExpr:      value.SubPathExpr,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodVolumeDevices(values []v1.VolumeDevice) []*PodVolumeDevice {
	var tValues []*PodVolumeDevice
	for _, value := range values {
		tValue := PodVolumeDevice{
			ClusterName: p.clusterName,
			Name:        value.Name,
			DevicePath:  value.DevicePath,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodHTTPHeaders(values []v1.HTTPHeader) []*PodHTTPHeader {
	var tValues []*PodHTTPHeader
	for _, value := range values {
		tValue := PodHTTPHeader{
			ClusterName: p.clusterName,
			Name:        value.Name,
			Value:       value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (p *Provider) transformPodExecActionCommand(values []string) []*PodExecActionCommand {
	var tValues []*PodExecActionCommand
	for _, v := range values {
		tValues = append(tValues, &PodExecActionCommand{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodCapabilitiesAdd(values []string) []*PodCapabilitiesAdd {
	var tValues []*PodCapabilitiesAdd
	for _, v := range values {
		tValues = append(tValues, &PodCapabilitiesAdd{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodCapabilitiesDrop(values []string) []*PodCapabilitiesDrop {
	var tValues []*PodCapabilitiesDrop
	for _, v := range values {
		tValues = append(tValues, &PodCapabilitiesDrop{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodContainers(values []v1.Container) []*PodContainer {
	var tValues []*PodContainer
	for _, value := range values {
		tValue := PodContainer{
			ClusterName: p.clusterName,
			Name:        value.Name,
			Image:       value.Image,
			Command:     p.transformPodContainerCommand(value.Command),
			Args:        p.transformPodContainerArgs(value.Args),
			WorkingDir:  value.WorkingDir,
			Ports:       p.transformPodContainerPorts(value.Ports),
			EnvFrom:     p.transformPodEnvFromSources(value.EnvFrom),
			Env:         p.transformPodEnvVars(value.Env),

			VolumeMounts:             p.transformPodVolumeMounts(value.VolumeMounts),
			VolumeDevices:            p.transformPodVolumeDevices(value.VolumeDevices),
			TerminationMessagePath:   value.TerminationMessagePath,
			TerminationMessagePolicy: string(value.TerminationMessagePolicy),
			ImagePullPolicy:          string(value.ImagePullPolicy),
			Stdin:                    value.Stdin,
			StdinOnce:                value.StdinOnce,
			TTY:                      value.TTY,
		}
		if value.LivenessProbe != nil {
			tValue.LivenessProbeInitialDelaySeconds = value.LivenessProbe.InitialDelaySeconds
			tValue.LivenessProbeTimeoutSeconds = value.LivenessProbe.TimeoutSeconds
			tValue.LivenessProbePeriodSeconds = value.LivenessProbe.PeriodSeconds
			tValue.LivenessProbeSuccessThreshold = value.LivenessProbe.SuccessThreshold
			tValue.LivenessProbeFailureThreshold = value.LivenessProbe.FailureThreshold
		}
		if value.ReadinessProbe != nil {

			tValue.ReadinessProbeInitialDelaySeconds = value.ReadinessProbe.InitialDelaySeconds
			tValue.ReadinessProbeTimeoutSeconds = value.ReadinessProbe.TimeoutSeconds
			tValue.ReadinessProbePeriodSeconds = value.ReadinessProbe.PeriodSeconds
			tValue.ReadinessProbeSuccessThreshold = value.ReadinessProbe.SuccessThreshold
			tValue.ReadinessProbeFailureThreshold = value.ReadinessProbe.FailureThreshold
		}
		if value.StartupProbe != nil {

			tValue.StartupProbeInitialDelaySeconds = value.StartupProbe.InitialDelaySeconds
			tValue.StartupProbeTimeoutSeconds = value.StartupProbe.TimeoutSeconds
			tValue.StartupProbePeriodSeconds = value.StartupProbe.PeriodSeconds
			tValue.StartupProbeSuccessThreshold = value.StartupProbe.SuccessThreshold
			tValue.StartupProbeFailureThreshold = value.StartupProbe.FailureThreshold
		}

		if value.SecurityContext != nil {
			tValue.SecurityContextPrivileged = value.SecurityContext.Privileged
			tValue.SecurityContextRunAsUser = value.SecurityContext.RunAsUser
			tValue.SecurityContextRunAsGroup = value.SecurityContext.RunAsGroup
			tValue.SecurityContextRunAsNonRoot = value.SecurityContext.RunAsNonRoot
			tValue.SecurityContextReadOnlyRootFilesystem = value.SecurityContext.ReadOnlyRootFilesystem
			tValue.SecurityContextAllowPrivilegeEscalation = value.SecurityContext.AllowPrivilegeEscalation
			tValue.SecurityContextProcMount = value.SecurityContext.ProcMount
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodEphemeralContainers(values []v1.EphemeralContainer) []*PodEphemeralContainer {
	var tValues []*PodEphemeralContainer
	for _, value := range values {
		tValue := PodEphemeralContainer{
			ClusterName:         p.clusterName,
			TargetContainerName: value.TargetContainerName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodSysctls(values []v1.Sysctl) []*PodSysctl {
	var tValues []*PodSysctl
	for _, value := range values {
		tValue := PodSysctl{
			ClusterName: p.clusterName,
			Name:        value.Name,
			Value:       value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodAffinityTermNamespaces(values []string) []*PodAffinityTermNamespaces {
	var tValues []*PodAffinityTermNamespaces
	for _, v := range values {
		tValues = append(tValues, &PodAffinityTermNamespaces{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodTolerations(values []v1.Toleration) []*PodToleration {
	var tValues []*PodToleration
	for _, value := range values {
		tValue := PodToleration{
			ClusterName:       p.clusterName,
			Key:               value.Key,
			Operator:          string(value.Operator),
			Value:             value.Value,
			Effect:            string(value.Effect),
			TolerationSeconds: value.TolerationSeconds,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (p *Provider) transformPodHostAliasHostnames(values []string) []*PodHostAliasHostnames {
	var tValues []*PodHostAliasHostnames
	for _, v := range values {
		tValues = append(tValues, &PodHostAliasHostnames{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodHostAliass(values []v1.HostAlias) []*PodHostAlias {
	var tValues []*PodHostAlias
	for _, value := range values {
		tValue := PodHostAlias{
			ClusterName: p.clusterName,
			IP:          value.IP,
			Hostnames:   p.transformPodHostAliasHostnames(value.Hostnames),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (p *Provider) transformPodDNSConfigNameservers(values []string) []*PodDNSConfigNameservers {
	var tValues []*PodDNSConfigNameservers
	for _, v := range values {
		tValues = append(tValues, &PodDNSConfigNameservers{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodDNSConfigSearches(values []string) []*PodDNSConfigSearches {
	var tValues []*PodDNSConfigSearches
	for _, v := range values {
		tValues = append(tValues, &PodDNSConfigSearches{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformPodDNSConfigOptions(values []v1.PodDNSConfigOption) []*PodDNSConfigOption {
	var tValues []*PodDNSConfigOption
	for _, value := range values {
		tValue := PodDNSConfigOption{
			ClusterName: p.clusterName,
			Name:        value.Name,
			Value:       value.Value,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodReadinessGates(values []v1.PodReadinessGate) []*PodReadinessGate {
	var tValues []*PodReadinessGate
	for _, value := range values {
		tValue := PodReadinessGate{
			ClusterName:   p.clusterName,
			ConditionType: string(value.ConditionType),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodConditions(values []v1.PodCondition) []*PodCondition {
	var tValues []*PodCondition
	for _, value := range values {
		tValue := PodCondition{
			ClusterName: p.clusterName,
			Type:        string(value.Type),
			Status:      string(value.Status),

			Reason:  value.Reason,
			Message: value.Message,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodIPs(values []v1.PodIP) []*PodIP {
	var tValues []*PodIP
	for _, value := range values {
		tValue := PodIP{
			ClusterName: p.clusterName,
			IP:          value.IP,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformPodContainerStatuss(values []v1.ContainerStatus) []*PodContainerStatus {
	var tValues []*PodContainerStatus
	for _, value := range values {
		tValue := PodContainerStatus{
			ClusterName: p.clusterName,
			Name:        value.Name,

			Ready:        value.Ready,
			RestartCount: value.RestartCount,
			Image:        value.Image,
			ImageID:      value.ImageID,
			ContainerID:  value.ContainerID,
			Started:      value.Started,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type PodConfig struct {
	Filter string
}

var podTables = []interface{}{
	&Pod{},
	&PodVolume{},
	&PodContainer{},
	&PodContainerCommand{},
	&PodContainerArgs{},
	&PodContainerPort{},
	&PodEnvFromSource{},
	&PodEnvVar{},
	&PodVolumeMount{},
	&PodVolumeDevice{},
	&PodExecActionCommand{},
	&PodHTTPHeader{},
	&PodExecActionCommand{},
	&PodCapabilitiesAdd{},
	&PodCapabilitiesDrop{},
	&PodEphemeralContainer{},
	&PodSysctl{},
	&PodAffinityTermNamespaces{},
	&PodAffinityTermNamespaces{},
	&PodToleration{},
	&PodHostAlias{},
	&PodHostAliasHostnames{},
	&PodDNSConfigNameservers{},
	&PodDNSConfigSearches{},
	&PodDNSConfigOption{},
	&PodReadinessGate{},
	&PodCondition{},
	&PodIP{},
	&PodContainerStatus{},
}

func (p *Provider) pods(_ interface{}) error {
	//var config PodConfig
	ctx := context.Background()
	//err := mapstructure.Decode(gConfig, &config)
	//if err != nil {
	//	return err
	//}

	output, err := p.client.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	p.db.Where("cluster_name", p.clusterName).Delete(podTables...)
	p.db.ChunkedCreate(p.transformPods(output.Items))
	p.Logger.Info("Fetched resources", "resource", "k8s.services", "count", len(output.Items))

	return nil
}
