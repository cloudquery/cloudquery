
# Table: digitalocean_droplet_networks_v6
NetworkV6 represents a DigitalOcean IPv6 network.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|droplet_cq_id|uuid|Unique CloudQuery ID of digitalocean_droplets table (FK)|
|ip_address|inet|The IP address of the IPv6 network interface.|
|netmask|bigint|The netmask of the IPv6 network interface.|
|gateway|inet|The gateway of the specified IPv6 network interface.|
|type|text|The type of the IPv6 network interface.  **Note**: IPv6 private  networking is not currently supported. |
