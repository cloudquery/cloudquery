# Changelog

All notable changes to CloudQuery will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
### [0.23.5](https://github.com/cloudquery/cloudquery/compare/v0.23.4...v0.23.5) (2022-05-22)


### Features

* Classify some policy download errors as USER ([#742](https://github.com/cloudquery/cloudquery/issues/742)) ([8224e60](https://github.com/cloudquery/cloudquery/commit/8224e60d3a76d3b3f181d3b32b9153a63b04816a))


### Bug Fixes

* Classify "no policies in config" as USER error ([#743](https://github.com/cloudquery/cloudquery/issues/743)) ([4cbc03e](https://github.com/cloudquery/cloudquery/commit/4cbc03e22f5a0bbfa33812b407e65704727a88fd))
* Don't attempt to download provider in re-attach mode ([#748](https://github.com/cloudquery/cloudquery/issues/748)) ([59973b8](https://github.com/cloudquery/cloudquery/commit/59973b84826599915f7b76fc8d8b16626dd26c74))
* FetchId column regression ([#745](https://github.com/cloudquery/cloudquery/issues/745)) ([585d395](https://github.com/cloudquery/cloudquery/commit/585d39589ef6c27ae2aab5d224fc00a2387d7628))
* Handle DeadlineExceeded errors ([#741](https://github.com/cloudquery/cloudquery/issues/741)) ([0167ce4](https://github.com/cloudquery/cloudquery/commit/0167ce4158d4795fc3a4b0f6661c19ae197c20c9))
* Handle Outputting Policies With Selectors ([a3ecfc9](https://github.com/cloudquery/cloudquery/commit/a3ecfc9166170e1bb77011befd11a5fbe1c86007))
* Space trimming in telemetry file ([#734](https://github.com/cloudquery/cloudquery/issues/734)) ([16c4cfc](https://github.com/cloudquery/cloudquery/commit/16c4cfce7e15f4474af3ab5d7e0cdb3698d2d08e))

### [0.23.4](https://github.com/cloudquery/cloudquery/compare/v0.23.3...v0.23.4) (2022-05-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.4 ([#725](https://github.com/cloudquery/cloudquery/issues/725)) ([69afca7](https://github.com/cloudquery/cloudquery/commit/69afca7dd34200ef0fda2341293e3fb46ee75faa))
* Don't generate telemetry-random-id in current dir ([#729](https://github.com/cloudquery/cloudquery/issues/729)) ([5eb493b](https://github.com/cloudquery/cloudquery/commit/5eb493b7215dc488d515106beff1a863a384b002))
* Panic on nil fetch response ([#728](https://github.com/cloudquery/cloudquery/issues/728)) ([8118554](https://github.com/cloudquery/cloudquery/commit/811855475622955dcfb323298292bde958f4372d))
* Panic on nil fetch result ([#730](https://github.com/cloudquery/cloudquery/issues/730)) ([7f224d0](https://github.com/cloudquery/cloudquery/commit/7f224d0371ed0014948fb6c572adef20bdb16094))
* squash redact errors ([#727](https://github.com/cloudquery/cloudquery/issues/727)) ([bccf7b7](https://github.com/cloudquery/cloudquery/commit/bccf7b71094eef4552e9227e4290aeec9a47896f))

### [0.23.3](https://github.com/cloudquery/cloudquery/compare/v0.23.2...v0.23.3) (2022-05-17)


### Features

* Add global log id ([#714](https://github.com/cloudquery/cloudquery/issues/714)) ([cece150](https://github.com/cloudquery/cloudquery/commit/cece150a78c83365a36cb3c295de8218ae959995))
* Policy Output ([#664](https://github.com/cloudquery/cloudquery/issues/664)) ([31f7e19](https://github.com/cloudquery/cloudquery/commit/31f7e19463da541b5ec13e18f4faf6d91dcfe6b0))
* Resource list enhancements ([#706](https://github.com/cloudquery/cloudquery/issues/706)) ([1952a27](https://github.com/cloudquery/cloudquery/commit/1952a27f212e109bac7bc74761cf193478aa1289))
* Use database id as unique id ([#705](https://github.com/cloudquery/cloudquery/issues/705)) ([dc00381](https://github.com/cloudquery/cloudquery/commit/dc0038158924b48ac41cbe57f7140084f2059ec3))


### Bug Fixes

* Add missing descriptions ([#700](https://github.com/cloudquery/cloudquery/issues/700)) ([c3c288c](https://github.com/cloudquery/cloudquery/commit/c3c288c62ff134109b2f35ec1a73b6cdd63c2d72))
* Classify not found policies and improve errors ([#697](https://github.com/cloudquery/cloudquery/issues/697)) ([413a2cf](https://github.com/cloudquery/cloudquery/commit/413a2cfe757f6a29ebc2fdb2db07b99b1fa9c4a1))
* Classify policy parse errors as User ([#716](https://github.com/cloudquery/cloudquery/issues/716)) ([f5947bf](https://github.com/cloudquery/cloudquery/commit/f5947bf443631454d41c2764c45bb32e5cfc2058))
* Classify subdir not found error ([#701](https://github.com/cloudquery/cloudquery/issues/701)) ([1a30732](https://github.com/cloudquery/cloudquery/commit/1a307321dab4c75c6697b20f8756d7282689a5cf))
* Completion issue ([#703](https://github.com/cloudquery/cloudquery/issues/703)) ([21c7bfe](https://github.com/cloudquery/cloudquery/commit/21c7bfeeb7afee4f1da7b8492e7be3a4c92b2bca))
* Handle empty policy directory ([#699](https://github.com/cloudquery/cloudquery/issues/699)) ([6acd308](https://github.com/cloudquery/cloudquery/commit/6acd3087cb3a81d990c77f351b969445e12d2bfd))
* Remove empty keys from init config ([#696](https://github.com/cloudquery/cloudquery/issues/696)) ([0e8dda1](https://github.com/cloudquery/cloudquery/commit/0e8dda1aecf5ac8ca785f1f9d4912b412b040ae8))
* Remove lambda support ([#710](https://github.com/cloudquery/cloudquery/issues/710)) ([5254f34](https://github.com/cloudquery/cloudquery/commit/5254f34f30f96b27d82e627a2be6c302bcb174af))
* Remove unused lambda dependency ([#717](https://github.com/cloudquery/cloudquery/issues/717)) ([7c78974](https://github.com/cloudquery/cloudquery/commit/7c78974668ad4144c7d9ded285cb4290fb0b01e6))
* Set ID For all Versions ([#724](https://github.com/cloudquery/cloudquery/issues/724)) ([ac46d2a](https://github.com/cloudquery/cloudquery/commit/ac46d2ad77bc8987e693028211a034bfe70cb06f))

### [0.23.2](https://github.com/cloudquery/cloudquery/compare/v0.23.1...v0.23.2) (2022-05-11)


### Bug Fixes

* **deps:** Bump github.com/hashicorp/go-getter from 1.5.10 to 1.5.11 ([#691](https://github.com/cloudquery/cloudquery/issues/691)) ([2ef215e](https://github.com/cloudquery/cloudquery/commit/2ef215e70af2de6243e2fd424c6785a920a8bfb2))

### [0.23.1](https://github.com/cloudquery/cloudquery/compare/v0.23.0...v0.23.1) (2022-05-11)


### Features

* DSN credentials ([#670](https://github.com/cloudquery/cloudquery/issues/670)) ([35e27d0](https://github.com/cloudquery/cloudquery/commit/35e27d03bb4d1102c93b04b981ed435720171386))


### Bug Fixes

* Handle nil policy run response ([#688](https://github.com/cloudquery/cloudquery/issues/688)) ([bd3e3bd](https://github.com/cloudquery/cloudquery/commit/bd3e3bd36e7a531f0fdb56378c658a9822b1166e))
* Run detectors in order ([#690](https://github.com/cloudquery/cloudquery/issues/690)) ([a39b2b6](https://github.com/cloudquery/cloudquery/commit/a39b2b6c878d41bcd78e81e84daf1ee95f05d125))

## [0.23.0](https://github.com/cloudquery/cloudquery/compare/v0.22.10...v0.23.0) (2022-05-10)


### Features

* Change to rudder ([#650](https://github.com/cloudquery/cloudquery/issues/650)) ([8f3f4c1](https://github.com/cloudquery/cloudquery/commit/8f3f4c14be4b7f95b7c673b1de6d4c2153556f93))
* Track db installations ([#652](https://github.com/cloudquery/cloudquery/issues/652)) ([e38acb7](https://github.com/cloudquery/cloudquery/commit/e38acb7d70297f764b1683dffe8389d908636369))


### Bug Fixes

* Bug where policy_run always fails ([#667](https://github.com/cloudquery/cloudquery/issues/667)) ([402266e](https://github.com/cloudquery/cloudquery/commit/402266ec8995bcd36d58093a2072efa795d89a1b))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.2 ([#637](https://github.com/cloudquery/cloudquery/issues/637)) ([55a60a9](https://github.com/cloudquery/cloudquery/commit/55a60a95328e4b5db00a5689ce5da5aed46dcbe5))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.3 ([#658](https://github.com/cloudquery/cloudquery/issues/658)) ([351cce5](https://github.com/cloudquery/cloudquery/commit/351cce50ecfa03d1be6dfd1d3fd7e268368a8aeb))
* Enable logging config through HCL ([#604](https://github.com/cloudquery/cloudquery/issues/604)) ([51bd06c](https://github.com/cloudquery/cloudquery/commit/51bd06c83f2a371b9e18969faf6edb1967b62e62))
* Encode json ([#641](https://github.com/cloudquery/cloudquery/issues/641)) ([1c04e45](https://github.com/cloudquery/cloudquery/commit/1c04e4515a9d865b92475200c9959253735ca9cd))
* panic on sync failure ([#676](https://github.com/cloudquery/cloudquery/issues/676)) ([27d574f](https://github.com/cloudquery/cloudquery/commit/27d574f6262417071c615675ec22b586317c50aa))
* **policy:** Add missing GitHub getter ([#613](https://github.com/cloudquery/cloudquery/issues/613)) ([e3fc361](https://github.com/cloudquery/cloudquery/commit/e3fc361c12139c58de14e42ab7ba89f2a967508a))
* **policy:** Use firebase instead of GitHub API to get latest version ([#618](https://github.com/cloudquery/cloudquery/issues/618)) ([455ed23](https://github.com/cloudquery/cloudquery/commit/455ed23ca3f0d075028385359a47436b8b05ead9))
* Sync support optional provider args ([#642](https://github.com/cloudquery/cloudquery/issues/642)) ([5eac023](https://github.com/cloudquery/cloudquery/commit/5eac02321222f6a50b95308274cc631402ab213a))
* Validate db version before proceeding ([#653](https://github.com/cloudquery/cloudquery/issues/653)) ([5af7f61](https://github.com/cloudquery/cloudquery/commit/5af7f615c580e94d319e2ad99b470ead9afd18f2))


### Miscellaneous Chores

* Release 0.23.0 ([#674](https://github.com/cloudquery/cloudquery/issues/674)) ([d4a2502](https://github.com/cloudquery/cloudquery/commit/d4a250288832b28104ae7e5497fbe6dc9a8f1231))

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
