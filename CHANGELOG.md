# Changelog

All notable changes to CloudQuery will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
## [Unreleased] 

### :rocket: Added
* Added core migrations implementation
* Added fetch summary saving to `fetches` table


<!-- 
## Unreleased
### Added
### Changed
### Fixed
### Breaking Changes
-->

## [v0.19.0] - 2022-01-10
### Breaking Changes
* Policy command updated and spec changed [#369](https://github.com/cloudquery/cloudquery/pull/369)
### Fixed
* Fixed empty policy bug [#399](https://github.com/cloudquery/cloudquery/pull/399).
* Fixed lambda json conversion [#397](https://github.com/cloudquery/cloudquery/pull/397).
* Removed confusing error message [#391](https://github.com/cloudquery/cloudquery/pull/391).
* Respected absolute file path in policies [#395](https://github.com/cloudquery/cloudquery/pull/395).
* Fixed isLevel for logger [#385](https://github.com/cloudquery/cloudquery/pull/385).
* Fixed pathing for hub to use real source path [#394](https://github.com/cloudquery/cloudquery/pull/394).
* CreateDatabase: check for err in correct place [#389](https://github.com/cloudquery/cloudquery/pull/389).
* Prevented reporting of errors to sentry twice [#386](https://github.com/cloudquery/cloudquery/pull/386).

### :gear: Changed
* Removed stack traces from sentry [#387](https://github.com/cloudquery/cloudquery/pull/87).
* Sentry send stack trace only on panic [#390](https://github.com/cloudquery/cloudquery/pull/390).



## [v0.18.0]- 2022-01-03
### 🚀 Added
* On cancel show error [#371](https://github.com/cloudquery/cloudquery/pull/371)
### 💥 Breaking Changes
* Upgrade to sdk [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/releases/tag/v0.6.1)
### :gear: Changed
* remove the need for json hcl2json convertor [#375](https://github.com/cloudquery/cloudquery/pull/375)
* removed gen config [#370](https://github.com/cloudquery/cloudquery/pull/370)
### :spider: Fixed
* Table upgrades with history mode enabled [#381](https://github.com/cloudquery/cloudquery/pull/381).

## [v0.17.4]- 2021-12-23

### 💥 Breaking Changes
* Removed old terraform deployment from core package, new deployment located [here](https://github.com/cloudquery/terraform-aws-cloudquery) [#357](https://github.com/cloudquery/cloudquery/pull/357).

### :rocket: Added
* Drift: Use correct ID for EMR clusters [#360](https://github.com/cloudquery/cloudquery/pull/360).
* Policy: added more logging to policy execution [#341](https://github.com/cloudquery/cloudquery/pull/341).
* Added hash of config to telemetry [#358](https://github.com/cloudquery/cloudquery/pull/359).

### :spider: Fixed
* Fixed Sentry issues [#347](https://github.com/cloudquery/cloudquery/pull/347).


### :gear: Changed
* Changed how we classify errors for sentry reducing errors sent, so only critical errors are report [#350](https://github.com/cloudquery/cloudquery/pull/350).
* Disable sentry module reporting [#351](https://github.com/cloudquery/cloudquery/pull/351).
* Made `source` attribute optional in CloudQuery config [#352](https://github.com/cloudquery/cloudquery/pull/352).
* Improved misleading help messaeg in cloudquery init [#359](https://github.com/cloudquery/cloudquery/pull/359).


## [v0.17.3]- 2021-12-16

### :spider: Fixed
* Report panics to Sentry [#347](https://github.com/cloudquery/cloudquery/pull/347).

## [v0.17.2] - 2021-12-16

### :spider: Fixed
* Panic on `cloudquery fetch`

## [v0.17.1] - 2021-12-15

### :rocket: Added
* Added [#210](https://github.com/cloudquery/cloudquery/issues/210) contribution [guide](https://github.com/cloudquery/cloudquery/blob/main/.github/CONTRIBUTING.md) [#331](https://github.com/cloudquery/cloudquery/pull/331).
* Added new provider update available notification [#336](https://github.com/cloudquery/cloudquery/pull/336) fixes [#299](https://github.com/cloudquery/cloudquery/issues/299).
* Added notification if an update to CQ core is available [#338](https://github.com/cloudquery/cloudquery/pull/338).
* Added sentry for crash error reporting to improve stability [#342](https://github.com/cloudquery/cloudquery/pull/342).

### :gear: Changed
* Telemetry: collect hash of MAC + Hostname [#339](https://github.com/cloudquery/cloudquery/pull/339).

### :spider: Fixed
* Provider download routine added before to policy run command [#335](https://github.com/cloudquery/cloudquery/pull/335) fixes [#316](https://github.com/cloudquery/cloudquery/issues/316).
* Fixed [#303](https://github.com/cloudquery/cloudquery/issues/303) UUID output in policies [#332](https://github.com/cloudquery/cloudquery/pull/332).
* Fixed Telemetry error counting, changed `debug-telemetry` flag to only set open-telelmetry client to debug mode [#340](https://github.com/cloudquery/cloudquery/pull/340)


## [v0.17.0] - 2021-12-06

### 💥 Breaking Changes
* `policy run` flag `--subpath` has been removed to execute sub policy pass it as second argument i.e `policy run <policy_name> <subpath>`

### :rocket: Added
* Added `policy describe <policy_name>` subcommand, allowing to see all policies and sub-policies available and execution paths 
* Added support for CloudQuery History **Alpha** for more info see [docs](https://docs.cloudquery.io/cli/history/overview)
* Exposed diagnostic counts on fetch for telemetry [#319](https://github.com/cloudquery/cloudquery/pull/319)

### :spider: Fixed
* Fixed resource fetch summary total fetched resources wouldn't sum correctly [#326](https://github.com/cloudquery/cloudquery/pull/326)
* Provider fetch failure cancels out other provider fetches [#325](https://github.com/cloudquery/cloudquery/pull/325)

### :gear: Changed
* Upgraded to SDK Version [v0.5.3](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).
* Read persisted telemetry ID if exists [#313](https://github.com/cloudquery/cloudquery/pull/313)
* Cleanup init command [#320](https://github.com/cloudquery/cloudquery/pull/320)
* Improve logging for policy execution errors [#323](https://github.com/cloudquery/cloudquery/pull/323)
* Updated drift aws configuration for new version [#329](https://github.com/cloudquery/cloudquery/pull/329)

## [v0.16.2] - 2021-11-29

### :rocket: Added
* Added support for telemetry, to gain better insight on usage to improve features and tool performance. For additional info see [docs](https://docs.cloudquery.io/docs/cli/telemetry) [#280](https://github.com/cloudquery/cloudquery/pull/280).
* Added support for executing policy in policy [#302](https://github.com/cloudquery/cloudquery/issues/302)

### :spider: Fixed
* Fixed Policy Not Found unclear message [#306](https://github.com/cloudquery/cloudquery/issues/306)
* Fixed Logging Statements Output [#305](https://github.com/cloudquery/cloudquery/issues/305)

## [v0.16.1] - 2021-11-22

### :spider: Fixed
* Fix fetch failure on providers that don't support upgrade [#295](https://github.com/cloudquery/cloudquery/pull/295)

## [v0.16.0] - 2021-11-19

### :rocket: Added
* Added support for [Terraform Drift detection](https://www.cloudquery.io/blog/announcing-cloudquery-terraform-drift-detection).
* Allow regex patterns for drift configuration (both local files and s3 bucket + keys  [#281](https://github.com/cloudquery/cloudquery/issues/281)
* Run provider upgrades before fetch [#283](https://github.com/cloudquery/cloudquery/pull/283)
* Support running policies from configuration [#269](https://github.com/cloudquery/cloudquery/pull/269) 
* Added a changelog :rocket:

### :spider: Fixed
* Fixed Confusing Error when config.hcl doesn't exist [#277](https://github.com/cloudquery/cloudquery/issues/277)

## [0.15.11] - 2021-11-18

Base version at which changelog was introduced.

