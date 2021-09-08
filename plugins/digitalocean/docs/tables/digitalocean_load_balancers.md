
# Table: digitalocean_load_balancers
LoadBalancer represents a DigitalOcean load balancer configuration. Tags can only be provided upon the creation of a Load Balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|uuid|A unique ID that can be used to identify and reference a load balancer.|
|name|text|A human-readable name for a load balancer instance.|
|ip|inet|An attribute containing the public-facing IP address of the load balancer.|
|size|text|The size of the load balancer. The available sizes are lb-small, lb-medium, or lb-large. You can resize load balancers after creation up to once per hour. You cannot resize a load balancer within the first hour of its creation|
|algorithm|text|The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either `round_robin` or `least_connections`.|
|status|text|A status string indicating the current state of the load balancer. This can be `new`, `active`, or `errored`.|
|created|text|A time value given in ISO8601 combined date and time format that represents when the load balancer was created.|
|health_check_protocol|text|The protocol used for health checks sent to the backend Droplets. The possible values are `http`, `https`, or `tcp`.|
|health_check_port|bigint|An integer representing the port on the backend Droplets on which the health check will attempt a connection.|
|health_check_path|text|The path on the backend Droplets to which the load balancer instance will send a request.|
|health_check_check_interval_seconds|bigint|The number of seconds between between two consecutive health checks.|
|health_check_response_timeout_seconds|bigint|The number of seconds the load balancer instance will wait for a response until marking a health check as failed.|
|health_check_healthy_threshold|bigint|The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool.|
|health_check_unhealthy_threshold|bigint|The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool.|
|sticky_sessions_type|text|An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`.|
|sticky_sessions_cookie_name|text|The name of the cookie sent to the client. This attribute is only returned when using `cookies` for the sticky sessions type.|
|sticky_sessions_cookie_ttl_seconds|bigint|The number of seconds until the cookie set by the load balancer expires. This attribute is only returned when using `cookies` for the sticky sessions type.|
|region_slug|text|A human-readable string that is used as a unique identifier for each region.|
|region_name|text|The display name of the region.  This will be a full name that is used in the control panel and other interfaces.|
|region_sizes|text[]|This attribute is set to an array which contains the identifying slugs for the sizes available in this region.|
|region_available|boolean|This is a boolean value that represents whether new Droplets can be created in this region.|
|region_features|text[]|This attribute is set to an array which contains features available in this region|
|droplet_ids|integer[]|An array containing the IDs of the Droplets assigned to the load balancer.|
|tag|text|The name of a Droplet tag corresponding to Droplets assigned to the load balancer.|
|tags|text[]||
|redirect_http_to_https|boolean|A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443.|
|enable_proxy_protocol|boolean|A boolean value indicating whether PROXY Protocol is in use.|
|enable_backend_keepalive|boolean|A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets.|
|vpc_uuid|text|A string specifying the UUID of the VPC to which the load balancer is assigned.|
