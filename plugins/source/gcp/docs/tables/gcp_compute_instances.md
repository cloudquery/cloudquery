
# Table: gcp_compute_instances
Represents an Instance resource  An instance is a virtual machine that is hosted on Google Cloud Platform For more information, read Virtual Machine Instances
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|advanced_machine_features_enable_nested_virtualization|boolean|Whether to enable nested virtualization or not (default is false)|
|can_ip_forward|boolean|Allows this instance to send and receive packets with non-matching destination or source IPs This is required if you plan to use this instance to forward routes For more information, see Enabling IP Forwarding|
|confidential_instance_config_enable_confidential_compute|boolean|Defines whether the instance should have confidential compute enabled|
|cpu_platform|text|The CPU platform used by this instance|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|deletion_protection|boolean|Whether the resource should be protected against deletion|
|description|text|An optional description of this resource Provide this property when you create the resource|
|display_device_enable_display|boolean|Defines whether the instance has Display enabled|
|fingerprint|text|Specifies a fingerprint for this resource|
|guest_accelerators|jsonb|A list of the type and count of accelerator cards attached to the instance|
|hostname|text|Specifies the hostname of the instance The specified hostname must be RFC1035 compliant If hostname is not specified, the default hostname is [INSTANCE_NAME]c[PROJECT_ID]internal when using the global DNS, and [INSTANCE_NAME][ZONE]c[PROJECT_ID]internal when using zonal DNS|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#instance for instances|
|label_fingerprint|text|A fingerprint for the labels being applied to this image|
|labels|jsonb|Labels for this resource|
|last_start_timestamp|text|Last start timestamp in RFC3339 text format|
|last_stop_timestamp|text|Last stop timestamp in RFC3339 text format|
|last_suspended_timestamp|text|Last suspended timestamp in RFC3339 text format|
|machine_type|text|Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type|
|metadata_fingerprint|text|Specifies a fingerprint for this request|
|metadata_items|jsonb|Array of key/value pairs The total size of all keys and values must be less than 512 KB|
|metadata_kind|text|Type of the resource Always compute#metadata for metadata|
|min_cpu_platform|text|Specifies a minimum CPU platform for the VM instance Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge"|
|name|text|Name of the resource Provided by the client when the resource is created|
|private_ipv6_google_access|text|The private IPv6 google access type for the VM If not specified, use  INHERIT_FROM_SUBNETWORK as default|
|reservation_affinity_consume_reservation_type|text|Specifies the type of reservation from which this instance can consume resources: ANY_RESERVATION (default), SPECIFIC_RESERVATION, or NO_RESERVATION|
|reservation_affinity_key|text|Corresponds to the label key of a reservation resource To target a SPECIFIC_RESERVATION by name, specify googleapiscom/reservation-name as the key and specify the name of your reservation as its value|
|reservation_affinity_values|text[]|Corresponds to the label values of a reservation resource|
|resource_policies|text[]|Resource policies applied to this instance|
|satisfies_pzs|boolean|Reserved for future use|
|scheduling_automatic_restart|boolean|Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user) You can only set the automatic restart option for standard instances Preemptible instances cannot be automatically restarted  By default, this is set to true so an instance is automatically restarted if it is terminated by Compute Engine|
|scheduling_location_hint|text|An opaque location hint used to place the instance close to other resources This field is for use by internal tools that use the public API|
|scheduling_min_node_cpus|bigint|The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node|
|scheduling_on_host_maintenance|text|Defines the maintenance behavior for this instance For standard instances, the default behavior is MIGRATE For preemptible instances, the default and only possible behavior is TERMINATE For more information, see Setting Instance Scheduling Options|
|scheduling_preemptible|boolean|Defines whether the instance is preemptible This can only be set during instance creation or while the instance is stopped and therefore, in a `TERMINATED` state See Instance Life Cycle for more information on the possible instance states|
|self_link|text|Server-defined URL for this resource|
|shielded_instance_config_enable_integrity_monitoring|boolean|Defines whether the instance has integrity monitoring enabled Enabled by default|
|shielded_instance_config_enable_secure_boot|boolean|Defines whether the instance has Secure Boot enabled Disabled by default|
|shielded_instance_config_enable_vtpm|boolean|Defines whether the instance has the vTPM enabled Enabled by default|
|shielded_instance_integrity_policy_update_auto_learn_policy|boolean|Updates the integrity policy baseline using the measurements from the VM instance's most recent boot|
|start_restricted|boolean|Whether a VM has been restricted for start because Compute Engine has detected suspicious activity|
|status|text|The status of the instance One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED For more information about the status of the instance, see  Instance life cycle|
|status_message|text|An optional, human-readable explanation of the status|
|tags_fingerprint|text|Specifies a fingerprint for this request, which is essentially a hash of the tags' contents and used for optimistic locking The fingerprint is initially generated by Compute Engine and changes after every request|
|tags_items|text[]|An array of tags Each tag must be 1-63 characters long, and comply with RFC1035|
|zone|text|URL of the zone where the instance resides You must specify this field as part of the HTTP request URL It is not settable as a field in the request body|
