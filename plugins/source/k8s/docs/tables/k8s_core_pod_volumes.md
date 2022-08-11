
# Table: k8s_core_pod_volumes
Volume represents a named volume in a pod that may be accessed by any container in the pod.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pod_cq_id|uuid|Unique CloudQuery ID of k8s_core_pods table (FK)|
|name|text|Volume's name. Must be a DNS_LABEL and unique within the pod.|
|host_path|jsonb|Pre-existing file or directory on the host machine that is directly exposed to the container.|
|empty_dir|jsonb|Temporary directory that shares a pod's lifetime.|
|gce_persistent_disk|jsonb|GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod.|
|aws_elastic_block_store|jsonb|AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod.|
|secret|jsonb|A secret that should populate this volume.|
|nfs|jsonb|NFS mount on the host that shares a pod's lifetime|
|iscsi|jsonb|ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod.|
|glusterfs|jsonb|Glusterfs mount on the host that shares a pod's lifetime.|
|persistent_volume_claim|jsonb|Persistent volume claim.|
|rbd|jsonb|Rados Block Device mount on the host that shares a pod's lifetime.|
|flex_volume|jsonb|Generic volume resource that is provisioned/attached using an exec based plugin.|
|cinder|jsonb|Cinder volume attached and mounted on kubelets host machine.|
|ceph_fs|jsonb|Ceph FS mount on the host that shares a pod's lifetime.|
|flocker|jsonb|Flocker volume attached to a kubelet's host machine.|
|downward_api|jsonb|Optional: mode bits to use on created files by default|
|fc|jsonb|Fibre Channel resource that is attached to a kubelet's host machine.|
|azure_file|jsonb|Azure File Service mount on the host and bind mount to the pod.|
|config_map|jsonb|configMap that should populate this volume|
|vsphere_volume|jsonb|vSphere volume attached and mounted on kubelets host machine.|
|quobyte|jsonb|Quobyte mount on the host that shares a pod's lifetime.|
|azure_disk|jsonb|The Name of the data disk in the blob storage|
|photon_persistent_disk|jsonb|PhotonController persistent disk attached and mounted on kubelets host machine.|
|projected|jsonb|Items for all in one resources secrets, configmaps, and downward API.|
|portworx_volume|jsonb|Portworx volume attached and mounted on kubelets host machine.|
|scale_io|jsonb|ScaleIO persistent volume attached and mounted on Kubernetes nodes.|
|storage_os|jsonb|StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.|
|csi|jsonb|CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).|
|ephemeral|jsonb|Ephemeral represents a volume that is handled by a cluster storage driver.|
