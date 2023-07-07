# Table: azure_redis_caches

This table shows data for Azure Redis Caches.

https://learn.microsoft.com/en-us/rest/api/redis/redis/list-by-subscription?tabs=HTTP#redisresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|identity|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|