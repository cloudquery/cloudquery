<p align="center">
<a href="https://cloudquery.io">
<img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</a>
</p>

The open-source cloud asset inventory powered by SQL.

![BuildStatus](https://img.shields.io/github/workflow/status/cloudquery/cloudquery/test?style=flat-square)
![License](https://img.shields.io/github/license/cloudquery/cloudquery?style=flat-square)

CloudQuery extracts, transforms, and loads your cloud assets into [normalized](https://hub.cloudquery.io) PostgreSQL tables. CloudQuery enables you to assess, audit, and monitor the configurations of your cloud assets.

CloudQuery key use-cases and features:

- **Search**: Use standard SQL to find any asset based on any configuration or relation to other assets.
- **Visualize**: Connect CloudQuery standard PostgreSQL database to your favorite BI/Visualization tool such as Grafana, QuickSight, etc...
- **Policy-as-Code**: Codify your security & compliance rules with SQL as the query engine.

### Links

- Homepage: https://cloudquery.io
- Releases: https://github.com/cloudquery/cloudquery/releases
- Documentation: https://docs.cloudquery.io
- Hub (Provider and schema docs): https://hub.cloudquery.io/

### Supported providers (Actively expanding)

Check out https://hub.cloudquery.io.

If you want us to add a new provider or resource please open an [Issue](https://github.com/cloudquery/cloudquery/issues).

See [docs](https://docs.cloudquery.io/docs/developers/developing-new-provider) for developing new provider.

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
and transform to the specified SQL database by running the following command:

```shell script
cloudquery init aws # choose one or more from: [aws azure gcp okta]
# cloudquery init gcp azure # This will generate a config containing gcp and azure providers
# cloudquery init --help # Show all possible auto generated configs and flags
```

Once your `config.hcl` is generated run the following command to fetch the resources:

```shell script
# you can spawn a local postgresql with docker
# docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres
cloudquery fetch --dsn "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
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

#### Running CloudQuery Policies

CloudQuery comes with out-of-the-box [policies](https://hub.cloudquery.io/policies) which you can use as is or modify to fit your use-case.

For example, to run [AWS CIS](https://hub.cloudquery.io/policies/cloudquery/aws-cis-1.20/latest) policies enter the following commands (make sure you fetched all the resources beforehand by the `fetch` command):

```shell script
./cloudquery policy run aws-cis-1.2.0
```

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

## Hiring

If you are into Go, Backend, Cloud, GCP, AWS - ping us at jobs [at] our domain

## Contribution

Feel free to open Pull-Request for small fixes and changes. For bigger changes and new providers please open an issue first to prevent double work and discuss relevant stuff.
