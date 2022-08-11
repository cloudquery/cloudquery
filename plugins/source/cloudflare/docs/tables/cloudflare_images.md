
# Table: cloudflare_images
Image represents a Cloudflare Image.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|id|text|Image unique identifier|
|filename|text|Image file name|
|metadata|jsonb|User modifiable key-value store. Can be used for keeping references to another system of record for managing images. Metadata must not exceed 1024 bytes.|
|require_signed_url_s|boolean||
|variants|jsonb|Object specifying available variants for an image.|
|uploaded|timestamp without time zone|When the media item was uploaded.|
