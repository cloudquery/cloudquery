
# Table: digitalocean_firewall_outbound_rules
OutboundRule represents a DigitalOcean Firewall outbound rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firewall_cq_id|uuid|Unique CloudQuery ID of digitalocean_firewalls table (FK)|
|protocol|text||
|port_range|text||
|destinations_addresses|text[]||
|destinations_tags|text[]||
|destinations_droplet_ids|integer[]||
|destinations_load_balancer_uid_s|text[]||
|destinations_kubernetes_ids|text[]||
