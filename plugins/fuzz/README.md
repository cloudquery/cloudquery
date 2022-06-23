<p align="center">
<a href="https://cloudquery.io">
<img alt="cloudquery logo" width=75% src="https://github.com/cloudquery/cloudquery/raw/main/docs/images/logo.png" />
</a>
</p>

CloudQuery Test Provider ![BuildStatus](https://img.shields.io/github/workflow/status/cloudquery/cq-provider-fuzz/test?style=flat-square) ![License](https://img.shields.io/github/license/cloudquery/cloudquery?style=flat-square)
==================================

This [CloudQuery](https://github.com/cloudquery/cloudquery) provider is only for testing purposes.

## Fuzzing Test

Set the following environment variables to configure the fuzzing parameters:

```bash
# These are optional, the values below are the defaults
export CQ_FUZZ_NUMBER_OF_RESOURCES=200
export CQ_FUZZ_NUMBER_OF_RELATIONS=10
export CQ_FUZZ_RELATION_DEPTH=2
export CQ_FUZZ_MIN_FETCH_DELAY_MILLISECONDS=100
export CQ_FUZZ_MAX_FETCH_DELAY_MILLISECONDS=60000
```

## What is CloudQuery

CloudQuery pulls, normalize, expose and monitor your cloud infrastructure and SaaS apps as a SQL database.
This abstracts various scattered APIs enabling you to define security,governance,cost and compliance policies with SQL

cloudquery can be easily extended to more resources and SaaS providers (open an [Issue](https://github.com/cloudquery/cloudquery/issues)).

cloudquery comes with built-in policy packs such as: [AWS CIS](#running-policy-packs) (more is coming!).

Think about cloudquery as a compliance-as-code tool inspired by tools like [osquery](https://github.com/osquery/osquery)
and [terraform](https://github.com/hashicorp/terraform), cool right?

### Links
* Homepage: https://cloudquery.io
* Releases: https://github.com/cloudquery/cloudquery/releases
* Documentation: https://docs.cloudquery.io
* Schema explorer (schemaspy): https://schema.cloudquery.io/
* Database Configuration: https://docs.cloudquery.io/database-configuration