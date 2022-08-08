
# Table: digitalocean_projects
Project represents a DigitalOcean Project configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|The unique universal identifier of this project.|
|owner_uuid|text|The unique universal identifier of the project owner.|
|owner_id|bigint|The integer id of the project owner.|
|name|text|The human-readable name for the project. The maximum length is 175 characters and the name must be unique.|
|description|text|The description of the project. The maximum length is 255 characters.|
|purpose|text|The purpose of the project. The maximum length is 255 characters. It can have one of the following values:  - Just trying out DigitalOcean - Class project / Educational purposes - Website or blog - Web Application - Service or API - Mobile Application - Machine learning / AI / Data processing - IoT - Operational / Developer tooling  If another value for purpose is specified, for example, "your custom purpose", your purpose will be stored as `Other: your custom purpose`. |
|environment|text|The environment of the project's resources.|
|is_default|boolean|If true, all resources will be added to this project if no project is specified.|
|created_at|text|A time value given in ISO8601 combined date and time format that represents when the project was created.|
|updated_at|text|A time value given in ISO8601 combined date and time format that represents when the project was updated.|
