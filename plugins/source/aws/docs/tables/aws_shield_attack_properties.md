
# Table: aws_shield_attack_properties
Details of a Shield event
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|attack_cq_id|uuid|Unique CloudQuery ID of aws_shield_attacks table (FK)|
|attack_layer|text|The type of Shield event that was observed|
|attack_property_identifier|text|Defines the Shield event property information that is provided|
|top_contributors|jsonb|Contributor objects for the top five contributors to a Shield event|
|total|bigint|The total contributions made to this Shield event by all contributors|
|unit|text|The unit used for the ContributorValue property|
