
# Table: digitalocean_firewall_inbound_rules
InboundRule represents a DigitalOcean Firewall inbound rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firewall_cq_id|uuid|Unique CloudQuery ID of digitalocean_firewalls table (FK)|
|protocol|text||
|port_range|text||
|sources_addresses|text[]||
|sources_tags|text[]||
|sources_droplet_ids|integer[]||
|sources_load_balancer_uid_s|text[]||
|sources_kubernetes_ids|text[]||
