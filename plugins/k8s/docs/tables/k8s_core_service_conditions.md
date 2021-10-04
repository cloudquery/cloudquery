
# Table: k8s_core_service_conditions
Condition contains details for one aspect of the current state of this API Resource. --- This struct is intended for direct use as an array at the field path .status.conditions
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of k8s_core_services table (FK)|
|type|text|type of condition in CamelCase or in foo.example.com/CamelCase. --- Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt) +required +kubebuilder:validation:Required +kubebuilder:validation:Pattern=`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$` +kubebuilder:validation:MaxLength=316|
|status|text|status of the condition, one of True, False, Unknown. +required +kubebuilder:validation:Required +kubebuilder:validation:Enum=True;False;Unknown|
|observed_generation|bigint|observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance. +optional +kubebuilder:validation:Minimum=0|
|reason|text|reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty. +required +kubebuilder:validation:Required +kubebuilder:validation:MaxLength=1024 +kubebuilder:validation:MinLength=1 +kubebuilder:validation:Pattern=`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`|
|message|text|message is a human readable message indicating details about the transition. This may be an empty string. +required +kubebuilder:validation:Required +kubebuilder:validation:MaxLength=32768|
