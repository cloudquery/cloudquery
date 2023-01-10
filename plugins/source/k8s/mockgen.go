package main

//go:generate mockgen --destination=mocks/admissionregistration/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1" AdmissionregistrationV1Interface
//go:generate mockgen --destination=mocks/admissionregistration/v1/mutating_webhook_configuration.go --package=v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1" MutatingWebhookConfigurationsGetter,MutatingWebhookConfigurationInterface
//go:generate mockgen --destination=mocks/admissionregistration/v1/validating_webhook_configuration.go --package=v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1" ValidatingWebhookConfigurationsGetter,ValidatingWebhookConfigurationInterface

//go:generate mockgen --destination=mocks/admissionregistration/v1beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1" AdmissionregistrationV1beta1Interface

//go:generate mockgen --destination=mocks/apiserverinternal/v1alpha1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1" InternalV1alpha1Interface

//go:generate mockgen --destination=mocks/apps/v1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1" AppsV1Interface
//go:generate mockgen --destination=mocks/apps/v1/daemonset.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1" DaemonSetsGetter,DaemonSetInterface
//go:generate mockgen --destination=mocks/apps/v1/deployment.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1" DeploymentsGetter,DeploymentInterface
//go:generate mockgen --destination=mocks/apps/v1/replicaset.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1" ReplicaSetsGetter,ReplicaSetInterface
//go:generate mockgen --destination=mocks/apps/v1/statefulset.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1" StatefulSetsGetter,StatefulSetInterface

//go:generate mockgen --destination=mocks/apps/v1beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1beta1" AppsV1beta1Interface

//go:generate mockgen --destination=mocks/apps/v1beta2/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/apps/v1beta2" AppsV1beta2Interface

//go:generate mockgen --destination=mocks/authentication/v1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/authentication/v1" AuthenticationV1Interface

//go:generate mockgen --destination=mocks/authentication/v1beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/authentication/v1beta1" AuthenticationV1beta1Interface

//go:generate mockgen --destination=mocks/authorization/v1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/authorization/v1" AuthorizationV1Interface

//go:generate mockgen --destination=mocks/authorization/v1beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/authorization/v1beta1" AuthorizationV1beta1Interface

//go:generate mockgen --destination=mocks/autoscaling/v1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/autoscaling/v1" AutoscalingV1Interface
//go:generate mockgen --destination=mocks/autoscaling/v1/hpa.go --package=mocks "k8s.io/client-go/kubernetes/typed/autoscaling/v1" HorizontalPodAutoscalersGetter,HorizontalPodAutoscalerInterface

//go:generate mockgen --destination=mocks/autoscaling/v2/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/autoscaling/v2" AutoscalingV2Interface

//go:generate mockgen --destination=mocks/autoscaling/v2beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1" AutoscalingV2beta1Interface

//go:generate mockgen --destination=mocks/autoscaling/v2beta2/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2" AutoscalingV2beta2Interface

//go:generate mockgen --destination=mocks/batch/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/batch/v1" BatchV1Interface
//go:generate mockgen --destination=mocks/batch/v1/cronjobs.go --package=v1 "k8s.io/client-go/kubernetes/typed/batch/v1" CronJobsGetter,CronJobInterface
//go:generate mockgen --destination=mocks/batch/v1/jobs.go --package=v1 "k8s.io/client-go/kubernetes/typed/batch/v1" JobsGetter,JobInterface

//go:generate mockgen --destination=mocks/batch/v1beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/batch/v1beta1" BatchV1beta1Interface

//go:generate mockgen --destination=mocks/certificates/v1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/certificates/v1" CertificatesV1Interface
//go:generate mockgen --destination=mocks/certificates/v1/csrs.go --package=mocks "k8s.io/client-go/kubernetes/typed/certificates/v1" CertificateSigningRequestsGetter,CertificateSigningRequestInterface

//go:generate mockgen --destination=mocks/certificates/v1beta1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/certificates/v1beta1" CertificatesV1beta1Interface

//go:generate mockgen --destination=mocks/coordination/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/coordination/v1" CoordinationV1Interface
//go:generate mockgen --destination=mocks/coordination/v1/lease.go --package=v1 "k8s.io/client-go/kubernetes/typed/coordination/v1" LeasesGetter,LeaseInterface

//go:generate mockgen --destination=mocks/core/v1/client.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" CoreV1Interface
//go:generate mockgen --destination=mocks/core/v1/componentstatus.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" ComponentStatusesGetter,ComponentStatusInterface
//go:generate mockgen --destination=mocks/core/v1/configmaps.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" ConfigMapsGetter,ConfigMapInterface
//go:generate mockgen --destination=mocks/core/v1/endpoints.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" EndpointsGetter,EndpointsInterface
//go:generate mockgen --destination=mocks/core/v1/events.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" EventsGetter,EventInterface
//go:generate mockgen --destination=mocks/core/v1/limitrange.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" LimitRangesGetter,LimitRangeInterface
//go:generate mockgen --destination=mocks/core/v1/namespace.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" NamespacesGetter,NamespaceInterface
//go:generate mockgen --destination=mocks/core/v1/node.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" NodesGetter,NodeInterface
//go:generate mockgen --destination=mocks/core/v1/persistentvolume.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" PersistentVolumesGetter,PersistentVolumeInterface
//go:generate mockgen --destination=mocks/core/v1/persistentvolumeclaim.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" PersistentVolumeClaimsGetter,PersistentVolumeClaimInterface
//go:generate mockgen --destination=mocks/core/v1/pod.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" PodsGetter,PodInterface
//go:generate mockgen --destination=mocks/core/v1/podtemplate.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" PodTemplatesGetter,PodTemplateInterface
//go:generate mockgen --destination=mocks/core/v1/replicationcontroller.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" ReplicationControllersGetter,ReplicationControllerInterface
//go:generate mockgen --destination=mocks/core/v1/resourceauota.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" ResourceQuotasGetter,ResourceQuotaInterface
//go:generate mockgen --destination=mocks/core/v1/secret.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" SecretsGetter,SecretInterface
//go:generate mockgen --destination=mocks/core/v1/service.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" ServicesGetter,ServiceInterface
//go:generate mockgen --destination=mocks/core/v1/serviceaccount.go --package=mocks "k8s.io/client-go/kubernetes/typed/core/v1" ServiceAccountsGetter,ServiceAccountInterface

//go:generate mockgen --destination=mocks/discovery/client.go --package=mocks "k8s.io/client-go/discovery" DiscoveryInterface

//go:generate mockgen --destination=mocks/discovery/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/discovery/v1" DiscoveryV1Interface
//go:generate mockgen --destination=mocks/discovery/v1/endpointslice.go --package=v1 "k8s.io/client-go/kubernetes/typed/discovery/v1" EndpointSlicesGetter,EndpointSliceInterface
//go:generate mockgen --destination=mocks/discovery/v1/event.go --package=v1 "k8s.io/client-go/kubernetes/typed/events/v1" EventsGetter,EventInterface,EventsV1Interface

//go:generate mockgen --destination=mocks/discovery/v1beta1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1" DiscoveryV1beta1Interface
//go:generate mockgen --destination=mocks/discovery/v1beta1/event.go --package=v1 "k8s.io/client-go/kubernetes/typed/events/v1beta1" EventsV1beta1Interface

//go:generate mockgen --destination=mocks/extensions/v1beta1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1" ExtensionsV1beta1Interface

//go:generate mockgen --destination=mocks/flowcontrol/v1alpha1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1" FlowcontrolV1alpha1Interface
//go:generate mockgen --destination=mocks/flowcontrol/v1beta1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1" FlowcontrolV1beta1Interface
//go:generate mockgen --destination=mocks/flowcontrol/v1beta2/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2" FlowcontrolV1beta2Interface

//go:generate mockgen --destination=mocks/interface.go --package=mocks "k8s.io/client-go/kubernetes" Interface

//go:generate mockgen --destination=mocks/networking/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/networking/v1" NetworkingV1Interface
//go:generate mockgen --destination=mocks/networking/v1/ingress.go --package=v1 "k8s.io/client-go/kubernetes/typed/networking/v1" IngressesGetter,IngressInterface
//go:generate mockgen --destination=mocks/networking/v1/ingressclass.go --package=v1 "k8s.io/client-go/kubernetes/typed/networking/v1" IngressClassesGetter,IngressClassInterface
//go:generate mockgen --destination=mocks/networking/v1/networkpolicy.go --package=v1 "k8s.io/client-go/kubernetes/typed/networking/v1" NetworkPoliciesGetter,NetworkPolicyInterface
//go:generate mockgen --destination=mocks/networking/v1alpha1/client.go --package=v1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1" NetworkingV1alpha1Interface
//go:generate mockgen --destination=mocks/networking/v1beta1/client.go --package=v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1" NetworkingV1beta1Interface

//go:generate mockgen --destination=mocks/node/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/node/v1" NodeV1Interface
//go:generate mockgen --destination=mocks/node/v1/runtime.go --package=v1 "k8s.io/client-go/kubernetes/typed/node/v1" RuntimeClassesGetter,RuntimeClassInterface
//go:generate mockgen --destination=mocks/node/v1alpha1/client.go --package=v1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1" NodeV1alpha1Interface

//go:generate mockgen --destination=mocks/rbac/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/rbac/v1" RbacV1Interface
//go:generate mockgen --destination=mocks/rbac/v1/clusterrole.go --package=v1 "k8s.io/client-go/kubernetes/typed/rbac/v1" ClusterRolesGetter,ClusterRoleInterface
//go:generate mockgen --destination=mocks/rbac/v1/clusterrolebinding.go --package=v1 "k8s.io/client-go/kubernetes/typed/rbac/v1" ClusterRoleBindingsGetter,ClusterRoleBindingInterface
//go:generate mockgen --destination=mocks/rbac/v1/role.go --package=v1 "k8s.io/client-go/kubernetes/typed/rbac/v1" RolesGetter,RoleInterface
//go:generate mockgen --destination=mocks/rbac/v1/rolebinding.go --package=v1 "k8s.io/client-go/kubernetes/typed/rbac/v1" RoleBindingsGetter,RoleBindingInterface

//go:generate mockgen --destination=mocks/storage/v1/client.go --package=v1 "k8s.io/client-go/kubernetes/typed/storage/v1" StorageV1Interface
//go:generate mockgen --destination=mocks/storage/v1/csidriver.go --package=v1 "k8s.io/client-go/kubernetes/typed/storage/v1" CSIDriversGetter,CSIDriverInterface
//go:generate mockgen --destination=mocks/storage/v1/csinode.go --package=v1 "k8s.io/client-go/kubernetes/typed/storage/v1" CSINodesGetter,CSINodeInterface
//go:generate mockgen --destination=mocks/storage/v1/csistoragecapacity.go --package=v1 "k8s.io/client-go/kubernetes/typed/storage/v1" CSIStorageCapacitiesGetter,CSIStorageCapacityInterface
//go:generate mockgen --destination=mocks/storage/v1/storageclass.go --package=v1 "k8s.io/client-go/kubernetes/typed/storage/v1" StorageClassesGetter,StorageClassInterface
//go:generate mockgen --destination=mocks/storage/v1/volumeattachment.go --package=v1 "k8s.io/client-go/kubernetes/typed/storage/v1" VolumeAttachmentsGetter,VolumeAttachmentInterface
