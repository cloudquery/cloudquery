
# Table: cloudflare_certificate_pack_certificates
CertificatePackCertificate is the base structure of a TLS certificate that is contained within a certificate pack.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|certificate_pack_cq_id|uuid|Unique CloudQuery ID of cloudflare_certificate_packs table (FK)|
|id|text||
|hosts|text[]||
|issuer|text||
|signature|text||
|status|text||
|bundle_method|text||
|geo_restrictions_label|text||
|zone_id|text||
|uploaded_on|timestamp without time zone||
|modified_on|timestamp without time zone||
|expires_on|timestamp without time zone||
|priority|bigint||
