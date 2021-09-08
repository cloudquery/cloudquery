
# Table: digitalocean_load_balancer_droplets
Droplets that are co-located on the same physical hardware
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of digitalocean_load_balancers table (FK)|
|droplet_id|bigint|Unique identifier of Droplet assigned to the load balancer.|
|load_balancer_id|uuid|The unique identifier for the load balancer.|
