
# Table: digitalocean_droplet_networks_v4
NetworkV4 represents a DigitalOcean IPv4 Network.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|droplet_cq_id|uuid|Unique CloudQuery ID of digitalocean_droplets table (FK)|
|ip_address|inet|The IP address of the IPv4 network interface.|
|netmask|inet|The netmask of the IPv4 network interface.|
|gateway|inet|The gateway of the specified IPv4 network interface.  For private interfaces, a gateway is not provided. This is denoted by returning `nil` as its value. |
|type|text|The type of the IPv4 network interface.|
