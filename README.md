# cloudquery

cloudquery exposes your cloud configuration and metadata as sql tables,
providing powerful analysis and monitoring without writing code.

### Links
* Homepage: https://cloudquery.run
* Releases: https://github.com/cloudquery/cloudquery/releases
* Documentation: https://docs.cloudquery.run

## Download & install

You can download the precompiled binary from [releases](https://github.com/cloudquery/cloudquery/releases), or using CLI:

```bash
export VERSION=v0.1.0
export OS=Darwin # Possible values: Linux,Windows,Darwin
curl -L https://github.com/cloudquery/cloudquery/releases/download/${VERSION}/cloudquery_${OS}_x86_64 -o cloudquery
chmod a+x cloudquery
./cloudquery --help
```

## Quick Start

Currently, cloudquery only supports AWS (Azure, GCP DO are on the roadmap).
The number of AWS resources is actively expanding.

You should be authenticated with an AWS account with correct permission with either option (see full [documentation](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html)):
 * `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
 * `~/.aws/credentials` created via `aws configure`
 
 Run the following commands:
 
```bash
cp example.config.yml config.yml
# uncomment resource of interest in config.yml
./cloudquery
 ```

If you uncommented images (it may take up to 30 seconds to fetch all images),
you can now run the following basic query in your sqlite3 console (`sqlite3 ./cloudquery.db`):

```sqlite
SELECT * FROM aws_ec2_images;
```

Full Documentation, resources and SQL schema definitions are available [here](https://docs.cloudquery.run)

## License

By contributing to cloudquery you agree that your contributions will be licensed as defined on the LICENSE file.

## Compile and run

```
go build .
./cloudquery # --help to see all options
```

## Roadmap

cloudquery currently support AWS only, Azure, GCP and DO are on the near roadmap and we are actively expanding number of supported resource with AWS.
