package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCorePods(t *testing.T, ctrl *gomock.Controller) client.Services {
	pods := mocks.NewMockPodsClient(ctrl)
	pods.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.PodList{Items: []corev1.Pod{fakePod(t)}}, nil,
	)
	return client.Services{
		Pods: pods,
	}
}

func fakeThroughPointers(t *testing.T, ptrs []interface{}) {
	for i, ptr := range ptrs {
		if err := faker.FakeData(ptr); err != nil {
			t.Fatalf("%v %v", i, ptr)
		}
	}
}

func fakeVolume(t *testing.T) corev1.Volume {
	// faker chokes on volume.VolumeSource.Ephemeral
	var volume corev1.Volume
	fakeThroughPointers(t, []interface{}{
		&volume.Name,
		&volume.VolumeSource.HostPath,
		&volume.VolumeSource.EmptyDir,
		&volume.VolumeSource.GCEPersistentDisk,
		&volume.VolumeSource.AWSElasticBlockStore,
		&volume.VolumeSource.GitRepo,
		&volume.VolumeSource.Secret,
		&volume.VolumeSource.NFS,
		&volume.VolumeSource.ISCSI,
		&volume.VolumeSource.Glusterfs,
		&volume.VolumeSource.PersistentVolumeClaim,
		&volume.VolumeSource.RBD,
		&volume.VolumeSource.FlexVolume,
		&volume.VolumeSource.Cinder,
		&volume.VolumeSource.CephFS,
		&volume.VolumeSource.Flocker,
		&volume.VolumeSource.DownwardAPI,
		&volume.VolumeSource.FC,
		&volume.VolumeSource.AzureFile,
		&volume.VolumeSource.ConfigMap,
		&volume.VolumeSource.VsphereVolume,
		&volume.VolumeSource.Quobyte,
		&volume.VolumeSource.AzureDisk,
		&volume.VolumeSource.PhotonPersistentDisk,
		&volume.VolumeSource.Projected,
		&volume.VolumeSource.PortworxVolume,
		&volume.VolumeSource.ScaleIO,
		&volume.VolumeSource.StorageOS,
		&volume.VolumeSource.CSI,
		// &volume.VolumeSource.Ephemeral,
	})
	volume.Ephemeral = &corev1.EphemeralVolumeSource{}
	return volume
}

func fakeContainer(t *testing.T) corev1.Container {
	var c corev1.Container
	fakeThroughPointers(t, []interface{}{
		&c.Name,
		&c.Image,
		&c.Command,
		&c.Args,
		&c.WorkingDir,
		&c.Ports,
		&c.EnvFrom,
		&c.Env,
		// &c.Resources,
		&c.VolumeMounts,
		&c.VolumeDevices,
		// &c.LivenessProbe,
		// &c.ReadinessProbe,
		// &c.StartupProbe,
		// &c.Lifecycle,
		&c.TerminationMessagePath,
		&c.TerminationMessagePolicy,
		&c.ImagePullPolicy,
		&c.SecurityContext,
	})
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	c.Resources.Limits = rl
	c.Resources.Requests = rl
	c.LivenessProbe = &corev1.Probe{}
	c.ReadinessProbe = &corev1.Probe{}
	c.StartupProbe = &corev1.Probe{}
	c.Lifecycle = &corev1.Lifecycle{}
	return c
}

func fakeEphemeralContainer(t *testing.T) corev1.EphemeralContainer {
	var c corev1.EphemeralContainer
	fakeThroughPointers(t, []interface{}{
		&c.TargetContainerName,
		&c.Name,
		&c.Image,
		&c.Command,
		&c.Args,
		&c.WorkingDir,
		&c.Ports,
		&c.EnvFrom,
		&c.Env,
		// &c.Resources,
		&c.VolumeMounts,
		&c.VolumeDevices,
		// &c.LivenessProbe,
		// &c.ReadinessProbe,
		// &c.StartupProbe,
		// &c.Lifecycle,
		&c.TerminationMessagePath,
		&c.TerminationMessagePolicy,
		&c.ImagePullPolicy,
		&c.SecurityContext,
	})
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	c.Resources.Limits = rl
	c.Resources.Requests = rl
	c.LivenessProbe = &corev1.Probe{}
	c.ReadinessProbe = &corev1.Probe{}
	c.StartupProbe = &corev1.Probe{}
	c.Lifecycle = &corev1.Lifecycle{}
	return c
}

func fakePod(t *testing.T) corev1.Pod {
	var pod corev1.Pod
	pod.Spec.Volumes = []corev1.Volume{fakeVolume(t)}
	fakeThroughPointers(t, []interface{}{
		&pod.TypeMeta,
		&pod.ObjectMeta,

		// &pod.Spec.InitContainers,
		// &pod.Spec.Containers,
		// &pod.Spec.EphemeralContainers,
		&pod.Spec.RestartPolicy,
		&pod.Spec.TerminationGracePeriodSeconds,
		&pod.Spec.ActiveDeadlineSeconds,
		&pod.Spec.DNSPolicy,
		&pod.Spec.NodeSelector,
		&pod.Spec.ServiceAccountName,
		&pod.Spec.AutomountServiceAccountToken,
		&pod.Spec.NodeName,
		&pod.Spec.HostNetwork,
		&pod.Spec.HostPID,
		&pod.Spec.HostIPC,
		&pod.Spec.ShareProcessNamespace,
		&pod.Spec.SecurityContext,
		&pod.Spec.ImagePullSecrets,
		&pod.Spec.Hostname,
		&pod.Spec.Subdomain,
		&pod.Spec.Affinity,
		&pod.Spec.SchedulerName,
		&pod.Spec.Tolerations,
		&pod.Spec.HostAliases,
		&pod.Spec.PriorityClassName,
		&pod.Spec.Priority,
		&pod.Spec.DNSConfig,
		&pod.Spec.ReadinessGates,
		&pod.Spec.RuntimeClassName,
		&pod.Spec.EnableServiceLinks,
		&pod.Spec.PreemptionPolicy,
		// &pod.Spec.Overhead,
		&pod.Spec.TopologySpreadConstraints,
		&pod.Spec.SetHostnameAsFQDN,

		&pod.Status,
	})
	rl := make(corev1.ResourceList)
	rl["name"] = *apiresource.NewQuantity(1024*1024, apiresource.BinarySI)
	pod.Spec.Overhead = rl

	pod.Spec.InitContainers = []corev1.Container{fakeContainer(t)}
	pod.Spec.Containers = []corev1.Container{fakeContainer(t)}
	pod.Spec.EphemeralContainers = []corev1.EphemeralContainer{fakeEphemeralContainer(t)}
	return pod
}

func TestCorePods(t *testing.T) {
	k8sTestHelper(t, CorePods(), createCorePods)
}
