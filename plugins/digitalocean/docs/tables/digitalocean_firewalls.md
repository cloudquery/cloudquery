
# Table: digitalocean_firewalls
Firewall represents a DigitalOcean Firewall configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|uuid|A unique ID that can be used to identify and reference a firewall.|
|name|text|A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).|
|status|text|A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".|
|droplet_ids|integer[]|An array containing the IDs of the Droplets assigned to the firewall.|
|tags|text[]||
|created|text|A time value given in ISO8601 combined date and time format that represents when the firewall was created.|
