
# Table: digitalocean_snapshots
Snapshot represents a DigitalOcean Snapshot
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|The unique identifier for the snapshot.|
|name|text|A human-readable name for the snapshot.|
|resource_id|text|The unique identifier for the resource that the snapshot originated from.|
|resource_type|text|The type of resource that the snapshot originated from.|
|regions|text[]|An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.|
|min_disk_size|bigint|The minimum size in GB required for a volume or Droplet to use this snapshot.|
|size_giga_bytes|float|The billable size of the snapshot in gigabytes.|
|created|text|A time value given in ISO8601 combined date and time format that represents when the snapshot was created.|
|tags|text[]|An array of Tags the snapshot has been tagged with.|
