# Changelog

## [4.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.1.0...plugins-destination-s3-v4.2.0) (2023-05-09)


### Features

* **s3:** Add support for custom S3 endpoint ([#10589](https://github.com/cloudquery/cloudquery/issues/10589)) ([48c6379](https://github.com/cloudquery/cloudquery/commit/48c63791fcc1cdf665548403a0dcb6c4ef846bc3))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.23 ([#10576](https://github.com/cloudquery/cloudquery/issues/10576)) ([eeb13d5](https://github.com/cloudquery/cloudquery/commit/eeb13d5b1b6b6fcb32764c8711bfbb79da35f9a8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.65 ([#10582](https://github.com/cloudquery/cloudquery/issues/10582)) ([4ed90e3](https://github.com/cloudquery/cloudquery/commit/4ed90e3aa7454e54f956144da544d9fe6532cf1f))

## [4.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.0.0...plugins-destination-s3-v4.1.0) (2023-05-02)


### Features

* **s3:** Add `test_write` option to allow skipping write test ([#10287](https://github.com/cloudquery/cloudquery/issues/10287)) ([48f1a2f](https://github.com/cloudquery/cloudquery/commit/48f1a2f53b4308f4af581b3b28d57a07027154a6)), closes [#9839](https://github.com/cloudquery/cloudquery/issues/9839)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/filetypes/v2 to v2.0.4 ([#10451](https://github.com/cloudquery/cloudquery/issues/10451)) ([6b6867e](https://github.com/cloudquery/cloudquery/commit/6b6867e91556b69a1471a1b43585cce6820c5cd3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v3.1.2...plugins-destination-s3-v4.0.0) (2023-04-25)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes for CSV or JSON output formats, however the Parquet output changes for UUID columns, which now have dashes, and timestamps, which now uses the default Arrow time format (e.g. `2023-01-02 12:23:45`). If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([33b5382](https://github.com/cloudquery/cloudquery/commit/33b5382930a95a7dcbfee357aa83a80f6e066010))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.8 ([#9781](https://github.com/cloudquery/cloudquery/issues/9781)) ([69bb790](https://github.com/cloudquery/cloudquery/commit/69bb790afbeac9ff01a41e71c8f631fb60fe64d1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.20 ([#9782](https://github.com/cloudquery/cloudquery/issues/9782)) ([1febd5b](https://github.com/cloudquery/cloudquery/commit/1febd5bbd944459a2fcbe380eb90385ecccfb079))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.21 ([#10127](https://github.com/cloudquery/cloudquery/issues/10127)) ([3bcde69](https://github.com/cloudquery/cloudquery/commit/3bcde697c5f927fa4eab52ea4293f1f7724812d1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.61 ([#9791](https://github.com/cloudquery/cloudquery/issues/9791)) ([f9dcef8](https://github.com/cloudquery/cloudquery/commit/f9dcef81bb81da123b6820ef2c4b204325e64203))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.62 ([#10129](https://github.com/cloudquery/cloudquery/issues/10129)) ([13f8670](https://github.com/cloudquery/cloudquery/commit/13f867006cd17c92bc1b18022ab3a210266258d8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.63 ([#10267](https://github.com/cloudquery/cloudquery/issues/10267)) ([7a8a4c7](https://github.com/cloudquery/cloudquery/commit/7a8a4c787bf2849b799014f51d32bec42942d16d))
* **deps:** Update module github.com/cloudquery/filetypes/v2 to v2.0.3 ([#10277](https://github.com/cloudquery/cloudquery/issues/10277)) ([1988c5a](https://github.com/cloudquery/cloudquery/commit/1988c5a38a32a10bc65a47d092045cc8d5b02394))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10260](https://github.com/cloudquery/cloudquery/issues/10260)) ([53cbd9a](https://github.com/cloudquery/cloudquery/commit/53cbd9acd3e2fded9c002909e478010ae8371fe4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [3.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v3.1.1...plugins-destination-s3-v3.1.2) (2023-04-04)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 7e19111 ([#9561](https://github.com/cloudquery/cloudquery/issues/9561)) ([dab17b9](https://github.com/cloudquery/cloudquery/commit/dab17b9c73e93aeafa06a4643db5b932a5a463c9))
* **deps:** Update github.com/xitongsys/parquet-go-source digest to fbbcdea ([#9566](https://github.com/cloudquery/cloudquery/issues/9566)) ([4256350](https://github.com/cloudquery/cloudquery/commit/4256350bd8a727963fc244dccc5f13cd97b0e5cd))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.60 ([#9616](https://github.com/cloudquery/cloudquery/issues/9616)) ([d155d28](https://github.com/cloudquery/cloudquery/commit/d155d28f4956be7b2e32ed163f62b4e05432cf6f))
* **deps:** Update module github.com/cloudquery/filetypes to v1.6.2 ([#9659](https://github.com/cloudquery/cloudquery/issues/9659)) ([11d3160](https://github.com/cloudquery/cloudquery/commit/11d3160ac65294eafe76de038f939f5aa06fb247))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [3.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v3.1.0...plugins-destination-s3-v3.1.1) (2023-03-28)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.7 ([#9425](https://github.com/cloudquery/cloudquery/issues/9425)) ([c8a4ab1](https://github.com/cloudquery/cloudquery/commit/c8a4ab1aaf52a1ae68f816b26b6bf7c47910501e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.19 ([#9426](https://github.com/cloudquery/cloudquery/issues/9426)) ([2017697](https://github.com/cloudquery/cloudquery/commit/2017697a59970f61c79e713054e8d3e4e482c453))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.13.18 ([#9427](https://github.com/cloudquery/cloudquery/issues/9427)) ([b2ef029](https://github.com/cloudquery/cloudquery/commit/b2ef0292574d3fa03b7cba8d8a6d25031210079a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.59 ([#9429](https://github.com/cloudquery/cloudquery/issues/9429)) ([71c69a1](https://github.com/cloudquery/cloudquery/commit/71c69a110732f30c61e490360dfe0320fe5e211f))

## [3.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v3.0.2...plugins-destination-s3-v3.1.0) (2023-03-21)


### Features

* **s3:** Update filetypes to v1.6.0 ([#9149](https://github.com/cloudquery/cloudquery/issues/9149)) ([708a971](https://github.com/cloudquery/cloudquery/commit/708a971481bc68c8c9fedafeeb5ab84fbb9041b6))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.18 ([#9227](https://github.com/cloudquery/cloudquery/issues/9227)) ([f630ecc](https://github.com/cloudquery/cloudquery/commit/f630ecc28c19e8388626c823954dca9f561e3920))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.58 ([#9229](https://github.com/cloudquery/cloudquery/issues/9229)) ([f8654b4](https://github.com/cloudquery/cloudquery/commit/f8654b4deaaa1a38c5f653a382c1eb6cff6cec74))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [3.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v3.0.1...plugins-destination-s3-v3.0.2) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.5.1 ([#8951](https://github.com/cloudquery/cloudquery/issues/8951)) ([197559e](https://github.com/cloudquery/cloudquery/commit/197559e71e5d01b5a7fa194008ffc6e3cd22705f))

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v3.0.0...plugins-destination-s3-v3.0.1) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.6 ([#8882](https://github.com/cloudquery/cloudquery/issues/8882)) ([5fa0031](https://github.com/cloudquery/cloudquery/commit/5fa0031ff61a92ff1fc086c1fd8b201a5417af36))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.16 ([#8883](https://github.com/cloudquery/cloudquery/issues/8883)) ([82ffe4d](https://github.com/cloudquery/cloudquery/commit/82ffe4d5aada3b0d3a174fa7a7722ce1a3719993))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/ec2/imds to v1.12.24 ([#8885](https://github.com/cloudquery/cloudquery/issues/8885)) ([674fec4](https://github.com/cloudquery/cloudquery/commit/674fec4c02af4d39613d064ef7d88be62e0a160a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.56 ([#8886](https://github.com/cloudquery/cloudquery/issues/8886)) ([8a3db4b](https://github.com/cloudquery/cloudquery/commit/8a3db4b90501b32fbcc87e5800e2f34fa0b299b7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/ini to v1.3.31 ([#8889](https://github.com/cloudquery/cloudquery/issues/8889)) ([f8fdb07](https://github.com/cloudquery/cloudquery/commit/f8fdb074573c9fcf394f0f0969156beaaf0ef592))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.22 ([#8890](https://github.com/cloudquery/cloudquery/issues/8890)) ([3c5b412](https://github.com/cloudquery/cloudquery/commit/3c5b41286590308a47207460c93f132e28c8e0a3))
* Use the correct path for test file ([#8939](https://github.com/cloudquery/cloudquery/issues/8939)) ([8e4be40](https://github.com/cloudquery/cloudquery/commit/8e4be40ba80613498eb0a7f00d5809d0e9260d94))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.3.0...plugins-destination-s3-v3.0.0) (2023-03-09)


### ⚠ BREAKING CHANGES

* **dest-s3:** Make region required and don't make getbucketlocation ([#8843](https://github.com/cloudquery/cloudquery/issues/8843))

### Bug Fixes

* **dest-s3:** Make region required and don't make getbucketlocation ([#8843](https://github.com/cloudquery/cloudquery/issues/8843)) ([716aba3](https://github.com/cloudquery/cloudquery/commit/716aba36cad94cb6839bb94633b25043daca0d45))

## [2.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.2.3...plugins-destination-s3-v2.3.0) (2023-03-07)


### Features

* **deps:** Update to filetypes v1.5.0 with arrow support ([#8739](https://github.com/cloudquery/cloudquery/issues/8739)) ([1870d4b](https://github.com/cloudquery/cloudquery/commit/1870d4b1fa2a93fad0fcb8b58abdb20c636e11e4))


### Bug Fixes

* **deps:** Update golang.org/x/xerrors digest to 04be3eb ([#8561](https://github.com/cloudquery/cloudquery/issues/8561)) ([39ccfcd](https://github.com/cloudquery/cloudquery/commit/39ccfcd9a293509e67c31a668c843f2f799a5a38))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/pierrec/lz4/v4 to v4.1.17 ([#8623](https://github.com/cloudquery/cloudquery/issues/8623)) ([fd968d8](https://github.com/cloudquery/cloudquery/commit/fd968d8938e8b603f2e9f6405eac2409ac41636b))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [2.2.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.2.2...plugins-destination-s3-v2.2.3) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.5 ([#8446](https://github.com/cloudquery/cloudquery/issues/8446)) ([e86922b](https://github.com/cloudquery/cloudquery/commit/e86922b62e01d609bcdbacc6afdc2e51febeb7f0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.15 ([#8447](https://github.com/cloudquery/cloudquery/issues/8447)) ([98cb352](https://github.com/cloudquery/cloudquery/commit/98cb352834ea715bcb9365b2c124dc98eb9474db))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/ec2/imds to v1.12.23 ([#8449](https://github.com/cloudquery/cloudquery/issues/8449)) ([c59f43e](https://github.com/cloudquery/cloudquery/commit/c59f43e23944c0ffb4f9762bd3efe70a41e4731f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.55 ([#8450](https://github.com/cloudquery/cloudquery/issues/8450)) ([416a435](https://github.com/cloudquery/cloudquery/commit/416a435304cbef7c228b6ee1bc90ec9d1197ae1c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/ini to v1.3.30 ([#8453](https://github.com/cloudquery/cloudquery/issues/8453)) ([912401b](https://github.com/cloudquery/cloudquery/commit/912401b0b64ff41ad864403ab0cc3f280a0a6355))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.21 ([#8454](https://github.com/cloudquery/cloudquery/issues/8454)) ([7820d00](https://github.com/cloudquery/cloudquery/commit/7820d00414bebb5890beb2ac26326ce0d5a44199))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [2.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.2.1...plugins-destination-s3-v2.2.2) (2023-02-21)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.13 ([#8231](https://github.com/cloudquery/cloudquery/issues/8231)) ([1eb436d](https://github.com/cloudquery/cloudquery/commit/1eb436d4db2f467419413c250c9fd1252d0a2fa5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.53 ([#8233](https://github.com/cloudquery/cloudquery/issues/8233)) ([3bc3b86](https://github.com/cloudquery/cloudquery/commit/3bc3b8613a2e59fea4e0838d3b751e4da12b8379))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.20 ([#8234](https://github.com/cloudquery/cloudquery/issues/8234)) ([6516f73](https://github.com/cloudquery/cloudquery/commit/6516f735ac2edb576afbe168bf56f9d5b25eef71))
* **deps:** Update module github.com/cloudquery/filetypes to v1.4.2 ([#8218](https://github.com/cloudquery/cloudquery/issues/8218)) ([9e656c2](https://github.com/cloudquery/cloudquery/commit/9e656c2f204951b54839547df5d8360de3ba4778))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))
* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [2.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.2.0...plugins-destination-s3-v2.2.1) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.4.1 ([#8024](https://github.com/cloudquery/cloudquery/issues/8024)) ([380476f](https://github.com/cloudquery/cloudquery/commit/380476fc11cb5cab576f320baa12e215bb148f86))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [2.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.1.0...plugins-destination-s3-v2.2.0) (2023-02-14)


### Features

* Add support for date placeholders in S3 plugin ([#7981](https://github.com/cloudquery/cloudquery/issues/7981)) ([7d3c25f](https://github.com/cloudquery/cloudquery/commit/7d3c25fccd829595d1725097a8c6e034353b8ec5))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v2.0.0...plugins-destination-s3-v2.1.0) (2023-02-07)


### Features

* **s3:** Update filetypes to support parquet ([#7726](https://github.com/cloudquery/cloudquery/issues/7726)) ([2faaff2](https://github.com/cloudquery/cloudquery/commit/2faaff24452c9c8c8a71263ebe4e147f5d0b5ccf))


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.4 ([#7652](https://github.com/cloudquery/cloudquery/issues/7652)) ([2196050](https://github.com/cloudquery/cloudquery/commit/2196050848b7abdafa9174af97151d0dbdf629c4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.12 ([#7653](https://github.com/cloudquery/cloudquery/issues/7653)) ([59daf42](https://github.com/cloudquery/cloudquery/commit/59daf423f2992c89db3db542c000286800d4ca61))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/ec2/imds to v1.12.22 ([#7655](https://github.com/cloudquery/cloudquery/issues/7655)) ([4e56621](https://github.com/cloudquery/cloudquery/commit/4e56621f73f515874c15eddb6da8b349d0889d6c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.51 ([#7656](https://github.com/cloudquery/cloudquery/issues/7656)) ([43a0c59](https://github.com/cloudquery/cloudquery/commit/43a0c59ca701281fa558c7a73a7673e019ad3ad6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/ini to v1.3.29 ([#7659](https://github.com/cloudquery/cloudquery/issues/7659)) ([60f15d7](https://github.com/cloudquery/cloudquery/commit/60f15d7cadfb3323c9b072869e252cdc7dfb0aab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.19 ([#7660](https://github.com/cloudquery/cloudquery/issues/7660)) ([9035012](https://github.com/cloudquery/cloudquery/commit/9035012d6ac2d41bdbdf0e2bf6f025f1bbac058b))
* **deps:** Update module github.com/cloudquery/filetypes to v1.3.1 ([#7274](https://github.com/cloudquery/cloudquery/issues/7274)) ([d0b6df8](https://github.com/cloudquery/cloudquery/commit/d0b6df81915bb4d623f0580516f600144c78340d))
* **deps:** Update module github.com/cloudquery/filetypes to v1.3.2 ([#7598](https://github.com/cloudquery/cloudquery/issues/7598)) ([0a7a1a8](https://github.com/cloudquery/cloudquery/commit/0a7a1a839e78e8b4f8e30c284d43d9901d626af9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))
* **destinations:** Unmarshal spec error messages ([#7463](https://github.com/cloudquery/cloudquery/issues/7463)) ([85450ad](https://github.com/cloudquery/cloudquery/commit/85450adcc6a73e230a70ef2f56aff3d93dada185))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.1.1...plugins-destination-s3-v2.0.0) (2023-01-31)


### ⚠ BREAKING CHANGES

* Add ability to override CSV File options (headers and delimiters) ([#6958](https://github.com/cloudquery/cloudquery/issues/6958))
* **s3:** Clean inputs ([#7116](https://github.com/cloudquery/cloudquery/issues/7116))

### Features

* Add ability to override CSV File options (headers and delimiters) ([#6958](https://github.com/cloudquery/cloudquery/issues/6958)) ([d03819c](https://github.com/cloudquery/cloudquery/commit/d03819ce1439e5f0509eb128da5c6ed75acf416b))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.49 ([#7309](https://github.com/cloudquery/cloudquery/issues/7309)) ([16da39d](https://github.com/cloudquery/cloudquery/commit/16da39d4bd8a6851329cbd25c2d80801b1872663))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))
* **s3:** Clean inputs ([#7116](https://github.com/cloudquery/cloudquery/issues/7116)) ([0132fb4](https://github.com/cloudquery/cloudquery/commit/0132fb4782399f25fcc42eddca6174eebae61dcf))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.1.0...plugins-destination-s3-v1.1.1) (2023-01-24)


### Bug Fixes

* **s3:** Use reflection for sanitizeJSONKeys to cover more cases ([#7104](https://github.com/cloudquery/cloudquery/issues/7104)) ([a6608fe](https://github.com/cloudquery/cloudquery/commit/a6608feeec42d299b0beb4e7d00a6b27e8e71966))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.0.5...plugins-destination-s3-v1.1.0) (2023-01-24)


### Features

* Add support for {{TABLE}} and {{UUID}} placeholders in S3 plugin path ([#6951](https://github.com/cloudquery/cloudquery/issues/6951)) ([4dee50a](https://github.com/cloudquery/cloudquery/commit/4dee50ae128f2290f0bc1aefee9368b2c762642c))
* **s3:** Add Athena flag to S3 destination ([#7079](https://github.com/cloudquery/cloudquery/issues/7079)) ([eef5823](https://github.com/cloudquery/cloudquery/commit/eef5823fc629decadbed8a742a078e8847cd146b))
* **s3:** Automatically discover bucket location ([#6793](https://github.com/cloudquery/cloudquery/issues/6793)) ([1fea90c](https://github.com/cloudquery/cloudquery/commit/1fea90ccaa1441809c025d14c0b0fd178a253455))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.0.4...plugins-destination-s3-v1.0.5) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.0.6 ([#6858](https://github.com/cloudquery/cloudquery/issues/6858)) ([129f91d](https://github.com/cloudquery/cloudquery/commit/129f91d0a3e05dff5d790887d3b419efde68670c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.0.3...plugins-destination-s3-v1.0.4) (2023-01-11)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.0.5 ([#6681](https://github.com/cloudquery/cloudquery/issues/6681)) ([32adfa2](https://github.com/cloudquery/cloudquery/commit/32adfa259912f24f555fbb49b45ac697bdb4c9b3))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.0.2...plugins-destination-s3-v1.0.3) (2023-01-10)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.8 ([#6393](https://github.com/cloudquery/cloudquery/issues/6393)) ([ffba44f](https://github.com/cloudquery/cloudquery/commit/ffba44f1318eb401d2b7ce2fa91c155d8925d90d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.47 ([#6395](https://github.com/cloudquery/cloudquery/issues/6395)) ([71ec9b9](https://github.com/cloudquery/cloudquery/commit/71ec9b99328ae4b7b0739a0c22258a805b586948))
* **deps:** Update module github.com/cloudquery/filetypes to v1.0.2 ([#6340](https://github.com/cloudquery/cloudquery/issues/6340)) ([d0d867b](https://github.com/cloudquery/cloudquery/commit/d0d867b0cd6c8b2968133d62e99b3abc498e9a17))
* **deps:** Update module github.com/cloudquery/filetypes to v1.0.3 ([#6523](https://github.com/cloudquery/cloudquery/issues/6523)) ([5378f3b](https://github.com/cloudquery/cloudquery/commit/5378f3be6d9d0ee3eb899244e1c2800326477a53))
* **deps:** Update module github.com/cloudquery/filetypes to v1.0.4 ([#6565](https://github.com/cloudquery/cloudquery/issues/6565)) ([80c7c06](https://github.com/cloudquery/cloudquery/commit/80c7c069d2d078d2635707dfbb000221d788e354))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **deps:** Update plugin-sdk to v1.21.0 for destinations ([#6419](https://github.com/cloudquery/cloudquery/issues/6419)) ([f3b989f](https://github.com/cloudquery/cloudquery/commit/f3b989f7cbe335481dc01ad2a56cf7eff48e01d5))
* Return error if read is called when `no_rotate` is false ([#6263](https://github.com/cloudquery/cloudquery/issues/6263)) ([c475be7](https://github.com/cloudquery/cloudquery/commit/c475be71b66c63761dabdecb1f99b65f94e77549))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.0.1...plugins-destination-s3-v1.0.2) (2023-01-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.0.1 ([#6264](https://github.com/cloudquery/cloudquery/issues/6264)) ([da3a1f0](https://github.com/cloudquery/cloudquery/commit/da3a1f0135370e3086bdaed357588955cb0094e8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6222](https://github.com/cloudquery/cloudquery/issues/6222)) ([5ba0d6d](https://github.com/cloudquery/cloudquery/commit/5ba0d6dcdefa9575c361ba7a6cdd86bf985e40c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))
* **deps:** Update module github.com/thoas/go-funk to v0.9.3 ([#6179](https://github.com/cloudquery/cloudquery/issues/6179)) ([e6d89ec](https://github.com/cloudquery/cloudquery/commit/e6d89ec1848f4ca2484ec0f0d7ea5ccaf74c14f5))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v1.0.0...plugins-destination-s3-v1.0.1) (2022-12-30)


### Bug Fixes

* Update file,gcs,s3 to sdk 1.16.0 ([#6146](https://github.com/cloudquery/cloudquery/issues/6146)) ([9ee9384](https://github.com/cloudquery/cloudquery/commit/9ee938400d1bc28ac353b0e80f12d9094e348b29))

## 1.0.0 (2022-12-29)


### Features

* Add S3 destination ([#6108](https://github.com/cloudquery/cloudquery/issues/6108)) ([ef86871](https://github.com/cloudquery/cloudquery/commit/ef8687103f5eebdcda5a22edee2415063e535bed))

## Changelog
