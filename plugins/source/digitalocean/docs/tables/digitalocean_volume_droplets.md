
# Table: digitalocean_volume_droplets
Droplets that are co-located on the same physical hardware
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|volume_cq_id|uuid|Unique CloudQuery ID of digitalocean_volumes table (FK)|
|droplet_id|bigint|Unique identifier of Droplet the volume is attached to.|
|volume_id|text|The unique identifier for the block storage volume.|
