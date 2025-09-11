```yaml
kind: destination
spec:
  name: "firehose"
  path: "cloudquery/firehose"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_FIREHOSE"
  write_mode: "append" # this plugin only supports 'append' mode
  send_sync_summary: true
  spec:
    # Required parameters e.g. arn:aws:firehose:us-east-1:111122223333:deliverystream/TestRedshiftStream
    stream_arn: "${FIREHOSE_STREAM_ARN}"
    # Optional parameters
    # max_retries: 5
    # max_record_size_bytes: 1024000 # optional
    # max_batch_records: 500 # optional
    # max_batch_size_bytes: 4194000 # optional
```
