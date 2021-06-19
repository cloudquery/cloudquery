<p align="center">
<a href="https://cloudquery.io">
<img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</a>
</p>

CloudQuery transforms your cloud infrastructure into queryable SQL for easy monitoring, governance and security.

### What is CloudQuery and why use it?

CloudQuery pulls, normalize, expose and monitor your cloud infrastructure and SaaS apps as SQL database.
This abstracts various scattered APIs enabling you to define security, governance, cost and compliance policies with SQL.

CloudQuery can be easily extended to more resources and SaaS providers (open an [Issue](https://github.com/cloudquery/cloudquery/issues)).

CloudQuery comes with built-in policy packs such as: [AWS CIS](#running-policy-packs) (more is coming!).

Think about CloudQuery as a compliance-as-code tool inspired by tools like [osquery](https://github.com/osquery/osquery)
and [terraform](https://github.com/hashicorp/terraform), cool right?

### Links

- Homepage: https://cloudquery.io
- Releases: https://github.com/cloudquery/cloudquery/releases
- Documentation: https://docs.cloudquery.io
- Hub (Provider and schema docs): https://hub.cloudquery.io/

### Supported providers (Actively expanding)

Checkout https://hub.cloudquery.io

If you want us to add a new provider or resource please open an [Issue](https://github.com/cloudquery/cloudquery/issues).

See [docs](https://docs.cloudquery.io/developers/developing-new-provider) for developing new provider.

## Download & install

You can download the precompiled binary from [releases](https://github.com/cloudquery/cloudquery/releases), or using CLI:

```shell script
export OS=Darwin # Possible values: Linux,Windows,Darwin
curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_x86_64 -o cloudquery
chmod a+x cloudquery
./cloudquery --help

# if you want to download a specific version and not latest use the following endpoint
export VERSION= # specifiy a version
curl -L https://github.com/cloudquery/cloudquery/releases/download/${VERSION}/cloudquery_${OS}_x86_64 -o cloudquery
```

Homebrew

```shell script
brew install cloudquery/tap/cloudquery
# After initial install you can upgrade the version via:
brew upgrade cloudquery
```

## Quick Start

### Running

First generate a `config.hcl` file that will describe which resources you want cloudquery to pull, normalize
and transform resources to the specified SQL database by running the following command:

```shell script
cloudquery init aws # choose one or more from: [aws azure gcp okta]
# cloudquery init gcp azure # This will generate a config containing gcp and azure providers
# cloudquery init --help # Show all possible auto generated configs and flags
```

Once your `config.hcl` is generated run the following command to fetch the resources:

```shell script
# you can spawn a local postgresql with docker
# docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres
cloudquery fetch --dsn "host=localhost user=postgres password=pass DB.name=postgres port=5432"
# cloudquery fetch --help # Show all possible fetch flags
```

Using `psql -h localhost -p 5432 -U postgres -d postgres`

```shell script
postgres=# \dt
                                    List of relations
 Schema |                            Name                             | Type  |  Owner
--------+-------------------------------------------------------------+-------+----------
 public | aws_autoscaling_launch_configuration_block_device_mapping   | table | postgres
 public | aws_autoscaling_launch_configurations                       | table | postgres
```

Run the following example queries from `psql` shell

List ec2_images

```sql
SELECT * FROM aws_ec2_images;
```

Find all public facing AWS load balancers

```sql
SELECT * FROM aws_elbv2_load_balancers WHERE scheme = 'internet-facing';
```

#### Running policy packs

cloudquery comes with some ready compliance policy pack which you can use as is or modify to fit your use-case.

Currently, cloudquery support [AWS CIS](https://d0.awsstatic.com/whitepapers/compliance/AWS_CIS_Foundations_Benchmark.pdf)
policy pack (it is under active development, so it doesn't cover the whole spec yet).

To run AWS CIS pack enter the following commands (make sure you fetched all the resources beforehand by the `fetch` command):

```shell script
./cloudquery policy --path=<PATH_TO_POLICY_FILE> --output=<PATH_TO_OUTPUT_POLICY_RESULT> --dsn "host=localhost user=postgres password=pass DB.name=postgres port=5432"
```

You can also create your own policy file. E.g.:

```yaml
views:
  - name: "my_custom_view"
    query: >
      CREATE VIEW my_custom_view AS ...
queries:
  - name: "Find thing that violates policy"
    query: >
      SELECT account_id, arn FROM ...
```

The `policy` command uses the policy file path `./policy.yml` by default, but this can be overridden via the `--path` flag, or the `CQ_POLICY_PATH` environment variable.

Full Documentation, resources and SQL schema definitions are available [here](https://hub.cloudquery.io).

### Providers Authentication

See additional documentation for each provider at [https://hub.cloudquery.io](https://hub.cloudquery.io).

## Compile and run

```
go build .
./cloudquery # --help to see all options
```

## Running on AWS (Lambda, Terraform)

Checkout [cloudquery/terraform-aws-cloudquery](https://github.com/cloudquery/terraform-aws-cloudquery)

## License

By contributing to cloudquery you agree that your contributions will be licensed as defined on the LICENSE file.

## Contribution

Feel free to open Pull-Request for small fixes and changes. For bigger changes and new providers please open an issue first to prevent double work and discuss relevant stuff.
