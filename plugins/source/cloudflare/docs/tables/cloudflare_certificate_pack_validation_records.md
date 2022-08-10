
# Table: cloudflare_certificate_pack_validation_records
SSLValidationRecord displays Domain Control Validation tokens.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|certificate_pack_cq_id|uuid|Unique CloudQuery ID of cloudflare_certificate_packs table (FK)|
|cname_target|text||
|cname_name|text||
|txt_name|text||
|txt_value|text||
|http_url|text||
|http_body|text||
|emails|text[]||
