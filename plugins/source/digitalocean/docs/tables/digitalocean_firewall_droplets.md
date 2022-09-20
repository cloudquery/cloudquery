
# Table: digitalocean_firewall_droplets
IDs of the Droplets assigned to the firewall
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firewall_cq_id|uuid|Unique CloudQuery ID of digitalocean_firewalls table (FK)|
|droplet_id|bigint|Unique identifier of Droplet assigned to the firewall.|
|firewall_id|uuid|The unique identifier for the firewall.|
