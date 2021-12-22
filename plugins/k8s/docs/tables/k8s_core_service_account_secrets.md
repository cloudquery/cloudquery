
# Table: k8s_core_service_account_secrets
ObjectReference contains enough information to let you inspect or modify the referred object. --- New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.  1
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_account_cq_id|uuid|Unique CloudQuery ID of k8s_core_service_accounts table (FK)|
|kind|text|Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds|
|namespace|text|Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/|
|name|text|Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names|
|uid|text|UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids|
|api_version|text|API version of the referent.|
|resource_version|text|Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency|
|field_path|text|If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod)|
