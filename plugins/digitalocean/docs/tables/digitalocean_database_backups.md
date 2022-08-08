
# Table: digitalocean_database_backups
DatabaseBackup represents a database backup.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of digitalocean_databases table (FK)|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format at which the backup was created.|
|size_gigabytes|float|The size of the database backup in GBs.|
