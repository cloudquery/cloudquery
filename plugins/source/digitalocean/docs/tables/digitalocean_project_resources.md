
# Table: digitalocean_project_resources
ProjectResource is the projects API's representation of a resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_cq_id|uuid|Unique CloudQuery ID of digitalocean_projects table (FK)|
|urn|text|The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.|
|assigned_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the project was created.|
|links_self|text|The links object contains the self object, which contains the resource relationship.|
|status|text|The status of assigning and fetching the resources.|
