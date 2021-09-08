
# Table: digitalocean_firewall_pending_changes
PendingChange represents a DigitalOcean Firewall status details.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firewall_cq_id|uuid|Unique CloudQuery ID of digitalocean_firewalls table (FK)|
|droplet_id|bigint||
|removing|boolean||
|status|text||
