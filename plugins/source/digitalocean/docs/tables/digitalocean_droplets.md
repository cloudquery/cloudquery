
# Table: digitalocean_droplets
Droplet represents a DigitalOcean Droplet
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|bigint|A unique identifier for each Droplet instance. This is automatically generated upon Droplet creation.|
|name|text|The human-readable name set for the Droplet instance.|
|memory|bigint|Memory of the Droplet in megabytes.|
|vcpus|bigint|The number of virtual CPUs.|
|disk|bigint|The size of the Droplet's disk in gigabytes.|
|region_slug|text|A human-readable string that is used as a unique identifier for each region.|
|region_name|text|The display name of the region.  This will be a full name that is used in the control panel and other interfaces.|
|region_sizes|text[]|This attribute is set to an array which contains the identifying slugs for the sizes available in this region.|
|region_available|boolean|This is a boolean value that represents whether new Droplets can be created in this region.|
|region_features|text[]|This attribute is set to an array which contains features available in this region|
|image_id|bigint|A unique number that can be used to identify and reference a specific image.|
|image_name|text|The display name that has been given to an image.  This is what is shown in the control panel and is generally a descriptive title for the image in question.|
|image_type|text|Describes the kind of image. It may be one of "snapshot", "backup", or "custom". This specifies whether an image is a user-generated Droplet snapshot, automatically created Droplet backup, or a user-provided virtual machine image.|
|image_distribution|text|The name of a custom image's distribution. Currently, the valid values are  "Arch Linux", "CentOS", "CoreOS", "Debian", "Fedora", "Fedora Atomic",  "FreeBSD", "Gentoo", "openSUSE", "RancherOS", "Ubuntu", and "Unknown".  Any other value will be accepted but ignored, and "Unknown" will be used in its place.|
|image_slug|text|A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.|
|image_public|boolean|This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.|
|image_regions|text[]|This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.|
|image_min_disk_size|bigint|The minimum disk size in GB required for a Droplet to use this image.|
|image_size_giga_bytes|float|The size of the image in gigabytes.|
|image_created|text|A time value given in ISO8601 combined date and time format that represents when the image was created.|
|image_description|text|An optional free-form text field to describe an image.|
|image_tags|text[]|A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.|
|image_status|text|A status string indicating the state of a custom image. This may be `NEW`,  `available`, `pending`, `deleted`, or `retired`.|
|image_error_message|text|A string containing information about errors that may occur when importing  a custom image.|
|size_memory|bigint|The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.|
|size_vcpus|bigint||
|size_disk|bigint|The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.|
|size_price_monthly|float||
|size_price_hourly|float|This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.|
|size_regions|text[]||
|size_available|boolean|This is a boolean value that represents whether new Droplets can be created with this size.|
|size_transfer|float|The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.|
|size_description|text||
|size_slug|text|The unique slug identifier for the size of this Droplet.|
|backup_ids|integer[]|An array of backup IDs of any backups that have been taken of the Droplet instance.  Droplet backups are enabled at the time of the instance creation.|
|next_backup_window_start_time|timestamp without time zone|A time value given in ISO8601 combined date and time format specifying the start of the Droplet's backup window.|
|next_backup_window_end_time|timestamp without time zone|A time value given in ISO8601 combined date and time format specifying the end of the Droplet's backup window.|
|snapshot_ids|integer[]|An array of snapshot IDs of any snapshots created from the Droplet instance.|
|features|text[]|An array of features enabled on this Droplet.|
|locked|boolean|A boolean value indicating whether the Droplet has been locked, preventing actions by users.|
|status|text|A status string indicating the state of the Droplet instance. This may be "new", "active", "off", or "archive".|
|created|text|A time value given in ISO8601 combined date and time format that represents when the Droplet was created.|
|kernel_id|bigint|A unique number used to identify and reference a specific kernel.|
|kernel_name|text|The display name of the kernel. This is shown in the web UI and is generally a descriptive title for the kernel in question.|
|kernel_version|text|A standard kernel version string representing the version, patch, and release information.|
|tags|text[]|An array of Tags the Droplet has been tagged with.|
|volume_ids|text[]|A flat array including the unique identifier for each Block Storage volume attached to the Droplet.|
|vpc_uuid|text|A string specifying the UUID of the VPC to which the Droplet is assigned.|
