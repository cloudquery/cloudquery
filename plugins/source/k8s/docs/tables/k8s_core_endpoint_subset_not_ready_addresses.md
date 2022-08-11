
# Table: k8s_core_endpoint_subset_not_ready_addresses
EndpointAddress is a tuple that describes single IP address.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|endpoint_subset_cq_id|uuid|Unique CloudQuery ID of k8s_core_endpoint_subsets table (FK)|
|ip|inet|The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms|
|hostname|text|The Hostname of this endpoint|
|node_name|text|Optional: Node hosting this endpoint|
|target_ref_kind|text|Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds|
|target_ref_namespace|text|Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/|
|target_ref_name|text|Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names|
|target_ref_uid|text|UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids|
|target_ref_api_version|text|API version of the referent.|
|target_ref_resource_version|text|Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency|
|target_ref_field_path|text|If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod)|
