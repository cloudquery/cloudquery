# Amazon Kinesis Firehose Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "firehose")}/>

This destination plugin lets you sync data from a CloudQuery source to a Amazon Kinesis Firehose in various formats such as CSV, JSON.


## Authentication

Authentication is handled by the AWS SDK. Credentials and configurations are sourced from the environment. Credentials are sourced in the following order:

1. Environment variables.
  Static Credentials (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN)
  Web Identity Token (AWS_WEB_IDENTITY_TOKEN_FILE)
2. Shared configuration files.
  SDK defaults to `credentials` file under `.aws` folder that is placed in the home folder on your computer.
  SDK defaults to `config` file under `.aws` folder that is placed in the home folder on your computer.
3. If your application uses an ECS task definition or RunTask API operation, IAM role for tasks.
4. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.


## Example


The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "firehose"
  path: "cloudquery/firehose"
  version: "VERSION_DESTINATION_FIREHOSE"
  write_mode: "append" # this plugin only supports 'append' mode
  # batch_size: 500 # optional
  # batch_size_bytes: 5242880 # optional
  spec:
    stream_arn: "arn:aws:firehose:us-east-1:111122223333:deliverystream/TestRedshiftStream"
```

The Amazon Kinesis Firehose destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes). 

It is important to note that Amazon Kinesis Firehose has the following limitations that cannot be changed:
  - The maximum size of a record sent to Kinesis Data Firehose, before base64-encoding, is 1,000 KiB.
  - The `PutRecordBatch` operation can take up to 500 records per batch or 4 MiB per batch, whichever is smaller.



## Firehose Spec

- `stream_arn` (string) (required)

  Kinesis Firehose delivery stream where data will be sent.

