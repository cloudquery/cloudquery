package main

//go:generate mockgen --package=mocks --destination=mocks/interface.go "k8s.io/client-go/kubernetes" Interface

//go:generate mockgen --package=mocks --destination=mocks/discovery/client.go "k8s.io/client-go/discovery" DiscoveryInterface

//go:generate mockgen --package=v1 --destination=mocks/admissionregistration/v1/client.go "k8s.io/client-go/kubernetes/typed/admissionregistration/v1" AdmissionregistrationV1Interface
//go:generate mockgen --package=v1 --destination=mocks/admissionregistration/v1/mutating_webhook_configuration.go "k8s.io/client-go/kubernetes/typed/admissionregistration/v1" MutatingWebhookConfigurationsGetter,MutatingWebhookConfigurationInterface
//go:generate mockgen --package=v1 --destination=mocks/admissionregistration/v1/validating_webhook_configuration.go "k8s.io/client-go/kubernetes/typed/admissionregistration/v1" ValidatingWebhookConfigurationsGetter,ValidatingWebhookConfigurationInterface

//go:generate mockgen --package=mocks --destination=mocks/admissionregistration/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1" AdmissionregistrationV1beta1Interface

//go:generate mockgen --package=mocks --destination=mocks/apiserverinternal/v1alpha1/client.go "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1" InternalV1alpha1Interface

//go:generate mockgen --package=mocks --destination=mocks/apps/v1/client.go "k8s.io/client-go/kubernetes/typed/apps/v1" AppsV1Interface
//go:generate mockgen --package=mocks --destination=mocks/apps/v1/daemonset.go "k8s.io/client-go/kubernetes/typed/apps/v1" DaemonSetsGetter,DaemonSetInterface
//go:generate mockgen --package=mocks --destination=mocks/apps/v1/deployment.go "k8s.io/client-go/kubernetes/typed/apps/v1" DeploymentsGetter,DeploymentInterface
//go:generate mockgen --package=mocks --destination=mocks/apps/v1/replicaset.go "k8s.io/client-go/kubernetes/typed/apps/v1" ReplicaSetsGetter,ReplicaSetInterface
//go:generate mockgen --package=mocks --destination=mocks/apps/v1/statefulset.go "k8s.io/client-go/kubernetes/typed/apps/v1" StatefulSetsGetter,StatefulSetInterface

//go:generate mockgen --package=mocks --destination=mocks/apps/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/apps/v1beta1" AppsV1beta1Interface

//go:generate mockgen --package=mocks --destination=mocks/apps/v1beta2/client.go "k8s.io/client-go/kubernetes/typed/apps/v1beta2" AppsV1beta2Interface

//go:generate mockgen --package=mocks --destination=mocks/authentication/v1/client.go "k8s.io/client-go/kubernetes/typed/authentication/v1" AuthenticationV1Interface

//go:generate mockgen --package=mocks --destination=mocks/authentication/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/authentication/v1beta1" AuthenticationV1beta1Interface

//go:generate mockgen --package=mocks --destination=mocks/authorization/v1/client.go "k8s.io/client-go/kubernetes/typed/authorization/v1" AuthorizationV1Interface

//go:generate mockgen --package=mocks --destination=mocks/authorization/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/authorization/v1beta1" AuthorizationV1beta1Interface

//go:generate mockgen --package=mocks --destination=mocks/autoscaling/v1/client.go "k8s.io/client-go/kubernetes/typed/autoscaling/v1" AutoscalingV1Interface
//go:generate mockgen --package=mocks --destination=mocks/autoscaling/v1/hpa.go "k8s.io/client-go/kubernetes/typed/autoscaling/v1" HorizontalPodAutoscalersGetter,HorizontalPodAutoscalerInterface

//go:generate mockgen --package=mocks --destination=mocks/autoscaling/v2/client.go "k8s.io/client-go/kubernetes/typed/autoscaling/v2" AutoscalingV2Interface

//go:generate mockgen --package=mocks --destination=mocks/autoscaling/v2beta1/client.go "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1" AutoscalingV2beta1Interface

//go:generate mockgen --package=mocks --destination=mocks/autoscaling/v2beta2/client.go "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2" AutoscalingV2beta2Interface

//go:generate mockgen --package=v1 --destination=mocks/batch/v1/client.go "k8s.io/client-go/kubernetes/typed/batch/v1" BatchV1Interface
//go:generate mockgen --package=v1 --destination=mocks/batch/v1/cronjobs.go "k8s.io/client-go/kubernetes/typed/batch/v1" CronJobsGetter,CronJobInterface
//go:generate mockgen --package=v1 --destination=mocks/batch/v1/jobs.go "k8s.io/client-go/kubernetes/typed/batch/v1" JobsGetter,JobInterface

//go:generate mockgen --package=mocks --destination=mocks/batch/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/batch/v1beta1" BatchV1beta1Interface

//go:generate mockgen --package=mocks --destination=mocks/certificates/v1/client.go "k8s.io/client-go/kubernetes/typed/certificates/v1" CertificatesV1Interface
//go:generate mockgen --package=mocks --destination=mocks/certificates/v1/csrs.go "k8s.io/client-go/kubernetes/typed/certificates/v1" CertificateSigningRequestsGetter,CertificateSigningRequestInterface

//go:generate mockgen --package=mocks --destination=mocks/certificates/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/certificates/v1beta1" CertificatesV1beta1Interface

//go:generate mockgen --package=v1 --destination=mocks/coordination/v1/client.go "k8s.io/client-go/kubernetes/typed/coordination/v1" CoordinationV1Interface
//go:generate mockgen --package=v1 --destination=mocks/coordination/v1/lease.go "k8s.io/client-go/kubernetes/typed/coordination/v1" LeasesGetter,LeaseInterface

//go:generate mockgen --package=mocks --destination=mocks/core/v1/client.go "k8s.io/client-go/kubernetes/typed/core/v1" CoreV1Interface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/componentstatus.go "k8s.io/client-go/kubernetes/typed/core/v1" ComponentStatusesGetter,ComponentStatusInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/configmaps.go "k8s.io/client-go/kubernetes/typed/core/v1" ConfigMapsGetter,ConfigMapInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/endpoints.go "k8s.io/client-go/kubernetes/typed/core/v1" EndpointsGetter,EndpointsInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/events.go "k8s.io/client-go/kubernetes/typed/core/v1" EventsGetter,EventInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/limitrange.go "k8s.io/client-go/kubernetes/typed/core/v1" LimitRangesGetter,LimitRangeInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/namespace.go "k8s.io/client-go/kubernetes/typed/core/v1" NamespacesGetter,NamespaceInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/node.go "k8s.io/client-go/kubernetes/typed/core/v1" NodesGetter,NodeInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/persistentvolume.go "k8s.io/client-go/kubernetes/typed/core/v1" PersistentVolumesGetter,PersistentVolumeInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/persistentvolumeclaim.go "k8s.io/client-go/kubernetes/typed/core/v1" PersistentVolumeClaimsGetter,PersistentVolumeClaimInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/pod.go "k8s.io/client-go/kubernetes/typed/core/v1" PodsGetter,PodInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/podtemplate.go "k8s.io/client-go/kubernetes/typed/core/v1" PodTemplatesGetter,PodTemplateInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/replicationcontroller.go "k8s.io/client-go/kubernetes/typed/core/v1" ReplicationControllersGetter,ReplicationControllerInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/resourceauota.go "k8s.io/client-go/kubernetes/typed/core/v1" ResourceQuotasGetter,ResourceQuotaInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/secret.go "k8s.io/client-go/kubernetes/typed/core/v1" SecretsGetter,SecretInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/service.go "k8s.io/client-go/kubernetes/typed/core/v1" ServicesGetter,ServiceInterface
//go:generate mockgen --package=mocks --destination=mocks/core/v1/serviceaccount.go "k8s.io/client-go/kubernetes/typed/core/v1" ServiceAccountsGetter,ServiceAccountInterface

//go:generate mockgen --package=v1 --destination=mocks/discovery/v1/client.go "k8s.io/client-go/kubernetes/typed/discovery/v1" DiscoveryV1Interface
//go:generate mockgen --package=v1 --destination=mocks/discovery/v1/endpointslice.go "k8s.io/client-go/kubernetes/typed/discovery/v1" EndpointSlicesGetter,EndpointSliceInterface

//go:generate mockgen --package=v1 --destination=mocks/discovery/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/discovery/v1beta1" DiscoveryV1beta1Interface

//go:generate mockgen --package=v1 --destination=mocks/discovery/v1/client.go "k8s.io/client-go/kubernetes/typed/events/v1" EventsV1Interface
//go:generate mockgen --package=v1 --destination=mocks/discovery/v1/event.go "k8s.io/client-go/kubernetes/typed/events/v1" EventsGetter,EventInterface

//go:generate mockgen --package=v1 --destination=mocks/discovery/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/events/v1beta1" EventsV1beta1Interface

//go:generate mockgen --package=v1 --destination=mocks/extensions/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/extensions/v1beta1" ExtensionsV1beta1Interface

//go:generate mockgen --package=v1 --destination=mocks/flowcontrol/v1alpha1/client.go "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1" FlowcontrolV1alpha1Interface

//go:generate mockgen --package=v1 --destination=mocks/flowcontrol/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1" FlowcontrolV1beta1Interface

//go:generate mockgen --package=v1 --destination=mocks/flowcontrol/v1beta2/client.go "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2" FlowcontrolV1beta2Interface

//go:generate mockgen --package=v1 --destination=mocks/networking/v1/client.go "k8s.io/client-go/kubernetes/typed/networking/v1" NetworkingV1Interface
//go:generate mockgen --package=v1 --destination=mocks/networking/v1/ingress.go "k8s.io/client-go/kubernetes/typed/networking/v1" IngressesGetter,IngressInterface
//go:generate mockgen --package=v1 --destination=mocks/networking/v1/ingressclass.go "k8s.io/client-go/kubernetes/typed/networking/v1" IngressClassesGetter,IngressClassInterface
//go:generate mockgen --package=v1 --destination=mocks/networking/v1/networkpolicy.go "k8s.io/client-go/kubernetes/typed/networking/v1" NetworkPoliciesGetter,NetworkPolicyInterface

//go:generate mockgen --package=v1alpha1 --destination=mocks/networking/v1alpha1/client.go "k8s.io/client-go/kubernetes/typed/networking/v1alpha1" NetworkingV1alpha1Interface

//go:generate mockgen --package=v1beta1 --destination=mocks/networking/v1beta1/client.go "k8s.io/client-go/kubernetes/typed/networking/v1beta1" NetworkingV1beta1Interface

//go:generate mockgen --package=v1 --destination=mocks/node/v1/client.go "k8s.io/client-go/kubernetes/typed/node/v1" NodeV1Interface
//go:generate mockgen --package=v1 --destination=mocks/node/v1/runtime.go "k8s.io/client-go/kubernetes/typed/node/v1" RuntimeClassesGetter,RuntimeClassInterface

//go:generate mockgen --package=v1alpha1 --destination=mocks/node/v1alpha1/client.go "k8s.io/client-go/kubernetes/typed/node/v1alpha1" NodeV1alpha1Interface

//go:generate mockgen --package=v1 --destination=mocks/rbac/v1/client.go "k8s.io/client-go/kubernetes/typed/rbac/v1" RbacV1Interface
//go:generate mockgen --package=v1 --destination=mocks/rbac/v1/role.go "k8s.io/client-go/kubernetes/typed/rbac/v1" RolesGetter,RoleInterface
//go:generate mockgen --package=v1 --destination=mocks/rbac/v1/clusterrole.go "k8s.io/client-go/kubernetes/typed/rbac/v1" ClusterRolesGetter,ClusterRoleInterface
//go:generate mockgen --package=v1 --destination=mocks/rbac/v1/clusterrolebinding.go "k8s.io/client-go/kubernetes/typed/rbac/v1" ClusterRoleBindingsGetter,ClusterRoleBindingInterface
//go:generate mockgen --package=v1 --destination=mocks/rbac/v1/rolebinding.go "k8s.io/client-go/kubernetes/typed/rbac/v1" RoleBindingsGetter,RoleBindingInterface

//go:generate mockgen --package=v1 --destination=mocks/storage/v1/client.go "k8s.io/client-go/kubernetes/typed/storage/v1" StorageV1Interface
//go:generate mockgen --package=v1 --destination=mocks/storage/v1/csidriver.go "k8s.io/client-go/kubernetes/typed/storage/v1" CSIDriversGetter,CSIDriverInterface
//go:generate mockgen --package=v1 --destination=mocks/storage/v1/csinode.go "k8s.io/client-go/kubernetes/typed/storage/v1" CSINodesGetter,CSINodeInterface
//go:generate mockgen --package=v1 --destination=mocks/storage/v1/csistoragecapacity.go "k8s.io/client-go/kubernetes/typed/storage/v1" CSIStorageCapacitiesGetter,CSIStorageCapacityInterface
//go:generate mockgen --package=v1 --destination=mocks/storage/v1/storageclass.go "k8s.io/client-go/kubernetes/typed/storage/v1" StorageClassesGetter,StorageClassInterface
//go:generate mockgen --package=v1 --destination=mocks/storage/v1/volumeattachment.go "k8s.io/client-go/kubernetes/typed/storage/v1" VolumeAttachmentsGetter,VolumeAttachmentInterface
