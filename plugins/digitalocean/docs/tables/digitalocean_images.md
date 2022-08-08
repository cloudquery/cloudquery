
# Table: digitalocean_images
Image represents a DigitalOcean Image
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|bigint|A unique number that can be used to identify and reference a specific image.|
|name|text|The display name that has been given to an image.  This is what is shown in the control panel and is generally a descriptive title for the image in question.|
|type|text|Describes the kind of image. It may be one of "snapshot", "backup", or "custom". This specifies whether an image is a user-generated Droplet snapshot, automatically created Droplet backup, or a user-provided virtual machine image.|
|distribution|text|The name of a custom image's distribution. Currently, the valid values are  "Arch Linux", "CentOS", "CoreOS", "Debian", "Fedora", "Fedora Atomic",  "FreeBSD", "Gentoo", "openSUSE", "RancherOS", "Ubuntu", and "Unknown".  Any other value will be accepted but ignored, and "Unknown" will be used in its place.|
|slug|text|A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.|
|public|boolean|This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.|
|regions|text[]|This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.|
|min_disk_size|bigint|The minimum disk size in GB required for a Droplet to use this image.|
|size_giga_bytes|float|The size of the image in gigabytes.|
|created|text|A time value given in ISO8601 combined date and time format that represents when the image was created.|
|description|text|An optional free-form text field to describe an image.|
|tags|text[]|A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.|
|status|text|A status string indicating the state of a custom image. This may be `NEW`,  `available`, `pending`, `deleted`, or `retired`.|
|error_message|text|A string containing information about errors that may occur when importing  a custom image.|
