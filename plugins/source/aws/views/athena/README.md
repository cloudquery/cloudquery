# Athena Resources View

You can create an `aws_resources` view for Athena, similar to the [equivalent Postgres version](../resources.sql). This allows you to query resources across all tables by tag or ARN.

There are currently two methods to create this view:

1. a query that outputs a `CREATE VIEW` statement you can run in the Athena console (easiest, lowest effort)
2. a Go script that will create the view for you. This can be run as a binary or as a Lambda function (can be automated to run periodically)

Note that in both cases, the view works best when the underlying format is **parquet**. JSON and CSV formats are not supported.

## Option 1: Standalone SQL script

### Usage

Execute the [`resources.sql`](./resources.sql) script in Athena. The output of this execution will be a `CREATE OR REPLACE VIEW` statement.

Copy the full `CREATE OR REPLACE VIEW` output statement and execute it in Athena.

## Option 2: **Athena Resources View Creation Tool**
   
### Usage

The script can be run as a standalone script or as an AWS Lambda function. In either case, you will need a working installation of Go.

### Standalone binary

Build the binary:

```shell
go build -o athena_resources_view
``` 

Show the help options:

```shell
$ ./athena_resources_view --help
Usage of ./athena_resources_view:
  -catalog string
      Catalog name (default "awsdatacatalog")
  -database string
      Database name
  -output string
      Query output path (e.g. s3://bucket/path)
  -region string
      View name (default: aws_resources) (default "us-east-1")
  -view-name string
      View name (default: aws_resources) (default "aws_resources")
```

Run the script (it will use the AWS credentials from your environment):

```shell
$ ./athena_resources_view -database athena-example -output 's3://cloudquery-athena-example/output/'
```

### Lambda function

1. Build the binary:

  ```shell
  GOOS=linux GOARCH=amd64 go build -o main main.go
  ```

2. Zip the binary. On Linux or Mac, you can use the `zip` command:

  ```shell
  zip main.zip main
  ```
  
  On Windows, follow the instructions [here](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html#golang-package-windows) to create a zip file using the `build-lambda-zip` tool.

3. Create the execution role:

  ```shell
  aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
  ```

  Make a note of your role ARN. You will need this to create your function.

4. Add permissions to the role to allow it to execute Athena queries. You can use the following policy document as a template:

  ```json
  {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Effect": "Allow",
          "Action": [
            "athena:StartQueryExecution",
            "athena:GetQueryExecution",
            "athena:GetQueryResults",
            "glue:GetDatabases",
            "glue:GetDatabase",
            "glue:GetTables",
            "glue:GetTable",
            "glue:UpdateTable"
          ],
          "Resource": "*"
        },
        {
          "Effect": "Allow",
          "Action": [
            "s3:GetObject",
            "s3:ListBucket",
            "s3:PutObject",
            "s3:GetBucketLocation"
          ],
          "Resource": [
            "arn:aws:s3:::cloudquery-athena-example",
            "arn:aws:s3:::cloudquery-athena-example/output/*"
          ]
        }
      ]
  }
  ```

  Save this as `athena-policy.json`, make appropriate edits for your environment (e.g. replacing `cloudquery-athena-example/output` with the bucket and directory where query results should be written), and then run:

  ```shell
  aws iam put-role-policy --role-name lambda-ex --policy-name athena-policy --policy-document file://athena-policy.json
  ```

5. Create the function (replace `<your-role-arn>` with the ARN of the role you created in step 2):

  ```shell
  aws lambda create-function --function-name athena-resources-view --zip-file fileb://main.zip --handler main --runtime go1.x --role <your-role-arn>
  ```
  
  Athena queries can take a while to run, so we should increase the timeout. The default is 3 seconds, but we can increase this to 5 minutes:

  ```shell
  aws lambda update-function-configuration --function-name athena-resources-view --timeout 300
  ```
  
6. Finally, run the function. This might be easier from the console, but here is an example of how to do it from the command line (you will need to modify the values in the payload for your environment):

  ```shell
  aws lambda invoke --cli-binary-format raw-in-base64-out --function-name athena-resources-view --invocation-type Event --payload '{"catalog": "awsdatacatalog", "database": "athena-example", "output": "s3://cloudquery-athena-example/output", "view": "aws_resources", "region": "us-east-1", "extra_columns": []}' response.json
  ```
  
  The above command uses the following JSON payload, which you should adapt for your environment:
  ```json
  {
      "catalog": "awsdatacatalog",
      "database": "athena-example",
      "output": "s3://cloudquery-athena-example/output",
      "view": "aws_resources",
      "region": "us-east-1",
      "extra_columns": []
  }
  ```
  
  If any query errors occur, you should be able to see them in the "Recent Queries" tab in the Athena console. Also check the logs for the function itself for any clues.

With the Lambda created, you are free to then schedule it to run on a regular basis, or after a CloudQuery run.
