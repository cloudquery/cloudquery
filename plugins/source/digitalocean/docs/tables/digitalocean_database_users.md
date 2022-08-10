
# Table: digitalocean_database_users
DatabaseUser represents a user in the database
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of digitalocean_databases table (FK)|
|name|text|The name of a database user.|
|role|text|A string representing the database user's role. The value will be either "primary" or "normal". |
|my_sql_settings_auth_plugin|text|A string specifying the authentication method to be used for connections to the MySQL user account. The valid values are `mysql_native_password` or `caching_sha2_password`. If excluded when creating a new user, the default for the version of MySQL in use will be used. As of MySQL 8.0, the default is `caching_sha2_password`. |
