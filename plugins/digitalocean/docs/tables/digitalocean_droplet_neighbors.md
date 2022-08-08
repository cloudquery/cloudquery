
# Table: digitalocean_droplet_neighbors
Droplets that are co-located on the same physical hardware
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|droplet_cq_id|uuid|Unique CloudQuery ID of digitalocean_droplets table (FK)|
|droplet_id|bigint|Unique identifier of the droplet associated with the neighbor droplet.|
|neighbor_id|bigint|Droplet neighbor identifier that exists on same the same physical hardware as the droplet.|
