
# Table: azure_front_door_routing_rules
Routing rules represent specifications for traffic to treat and where to send it, along with health probe information.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_cq_id|uuid|Unique CloudQuery ID of azure_front_doors table (FK)|
|resource_state|text|Resource state|
|frontend_endpoints|text[]|Frontend endpoints associated with the rule|
|accepted_protocols|text[]|Protocol schemes to match for the rule|
|patterns_to_match|text[]|The route patterns of the rule|
|enabled_state|text|Whether the rule is enabled|
|route_configuration|jsonb|A reference to the routing configuration|
|rules_engine_id|text|ID of a specific Rules Engine Configuration to apply to the route|
|web_application_firewall_policy_link_id|text|ID of the Web Application Firewall policy for each routing rule (if applicable)|
|name|text|Resource name|
|type|text|Resource type|
|id|text|Resource ID|
