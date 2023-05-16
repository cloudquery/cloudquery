# Changelog

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v3.0.0...plugins-destination-azblob-v3.0.1) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/filetypes/v2 to v2.0.4 ([#10451](https://github.com/cloudquery/cloudquery/issues/10451)) ([6b6867e](https://github.com/cloudquery/cloudquery/commit/6b6867e91556b69a1471a1b43585cce6820c5cd3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.3.1...plugins-destination-azblob-v3.0.0) (2023-04-25)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes for CSV or JSON output formats, however the Parquet output changes for UUID columns, which now have dashes, and timestamps, which now uses the default Arrow time format (e.g. `2023-01-02 12:23:45`). If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([c86c36a](https://github.com/cloudquery/cloudquery/commit/c86c36a7fca58d600074b5408a6a7f0aeaebc26e))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes/v2 to v2.0.3 ([#10230](https://github.com/cloudquery/cloudquery/issues/10230)) ([730fa51](https://github.com/cloudquery/cloudquery/commit/730fa51814f6800069c138daf074d49b7dc3958b))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [2.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.3.0...plugins-destination-azblob-v2.3.1) (2023-04-04)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 7e19111 ([#9561](https://github.com/cloudquery/cloudquery/issues/9561)) ([dab17b9](https://github.com/cloudquery/cloudquery/commit/dab17b9c73e93aeafa06a4643db5b932a5a463c9))
* **deps:** Update github.com/xitongsys/parquet-go-source digest to fbbcdea ([#9566](https://github.com/cloudquery/cloudquery/issues/9566)) ([4256350](https://github.com/cloudquery/cloudquery/commit/4256350bd8a727963fc244dccc5f13cd97b0e5cd))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/filetypes to v1.6.2 ([#9659](https://github.com/cloudquery/cloudquery/issues/9659)) ([11d3160](https://github.com/cloudquery/cloudquery/commit/11d3160ac65294eafe76de038f939f5aa06fb247))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [2.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.2.1...plugins-destination-azblob-v2.3.0) (2023-03-21)


### Features

* **azblob:** Update to filetypes v1.6.0 ([#9152](https://github.com/cloudquery/cloudquery/issues/9152)) ([737b3a6](https://github.com/cloudquery/cloudquery/commit/737b3a6b4c0781112db94349e35a6f7012d92ea9))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [2.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.2.0...plugins-destination-azblob-v2.2.1) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.5.1 ([#8951](https://github.com/cloudquery/cloudquery/issues/8951)) ([197559e](https://github.com/cloudquery/cloudquery/commit/197559e71e5d01b5a7fa194008ffc6e3cd22705f))

## [2.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.1.3...plugins-destination-azblob-v2.2.0) (2023-03-08)


### Features

* **azblob:** Update filetypes to v1.5.0 with arrow support. ([#8741](https://github.com/cloudquery/cloudquery/issues/8741)) ([ee52b5a](https://github.com/cloudquery/cloudquery/commit/ee52b5afedb0cf6b54091757ae06291886ade1f8))


### Bug Fixes

* **deps:** Update golang.org/x/xerrors digest to 04be3eb ([#8561](https://github.com/cloudquery/cloudquery/issues/8561)) ([39ccfcd](https://github.com/cloudquery/cloudquery/commit/39ccfcd9a293509e67c31a668c843f2f799a5a38))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.3.1 ([#8567](https://github.com/cloudquery/cloudquery/issues/8567)) ([d17e2aa](https://github.com/cloudquery/cloudquery/commit/d17e2aa63e8258ecc3c6815431222d1c5f0a06bf))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/pierrec/lz4/v4 to v4.1.17 ([#8623](https://github.com/cloudquery/cloudquery/issues/8623)) ([fd968d8](https://github.com/cloudquery/cloudquery/commit/fd968d8938e8b603f2e9f6405eac2409ac41636b))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [2.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.1.2...plugins-destination-azblob-v2.1.3) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [2.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.1.1...plugins-destination-azblob-v2.1.2) (2023-02-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.4.2 ([#8218](https://github.com/cloudquery/cloudquery/issues/8218)) ([9e656c2](https://github.com/cloudquery/cloudquery/commit/9e656c2f204951b54839547df5d8360de3ba4778))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))
* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [2.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.1.0...plugins-destination-azblob-v2.1.1) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.4.1 ([#8024](https://github.com/cloudquery/cloudquery/issues/8024)) ([380476f](https://github.com/cloudquery/cloudquery/commit/380476fc11cb5cab576f320baa12e215bb148f86))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v2.0.0...plugins-destination-azblob-v2.1.0) (2023-02-07)


### Features

* **destination-azblob:** Update filetypes to support Parquet ([#7729](https://github.com/cloudquery/cloudquery/issues/7729)) ([efe4dbf](https://github.com/cloudquery/cloudquery/commit/efe4dbf3b64d940d45b22717fe562f164625c8f2))


### Bug Fixes

* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azidentity to v1.2.1 ([#7540](https://github.com/cloudquery/cloudquery/issues/7540)) ([3b5c838](https://github.com/cloudquery/cloudquery/commit/3b5c83832064d729ad1097728f7d12aedbbb9400))
* **deps:** Update module github.com/cloudquery/filetypes to v1.3.1 ([#7274](https://github.com/cloudquery/cloudquery/issues/7274)) ([d0b6df8](https://github.com/cloudquery/cloudquery/commit/d0b6df81915bb4d623f0580516f600144c78340d))
* **deps:** Update module github.com/cloudquery/filetypes to v1.3.2 ([#7598](https://github.com/cloudquery/cloudquery/issues/7598)) ([0a7a1a8](https://github.com/cloudquery/cloudquery/commit/0a7a1a839e78e8b4f8e30c284d43d9901d626af9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))
* **deps:** Update module github.com/golang-jwt/jwt/v4 to v4.4.3 ([#7543](https://github.com/cloudquery/cloudquery/issues/7543)) ([0607454](https://github.com/cloudquery/cloudquery/commit/060745428eda5839be801c153c2f7261fcc54abd))
* **destinations:** Unmarshal spec error messages ([#7463](https://github.com/cloudquery/cloudquery/issues/7463)) ([85450ad](https://github.com/cloudquery/cloudquery/commit/85450adcc6a73e230a70ef2f56aff3d93dada185))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v1.0.5...plugins-destination-azblob-v2.0.0) (2023-01-31)


### ⚠ BREAKING CHANGES

* Add ability to override CSV File options (headers and delimiters) ([#6958](https://github.com/cloudquery/cloudquery/issues/6958))

### Features

* Add ability to override CSV File options (headers and delimiters) ([#6958](https://github.com/cloudquery/cloudquery/issues/6958)) ([d03819c](https://github.com/cloudquery/cloudquery/commit/d03819ce1439e5f0509eb128da5c6ed75acf416b))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v1.0.4...plugins-destination-azblob-v1.0.5) (2023-01-24)


### Bug Fixes

* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.3.0 ([#6955](https://github.com/cloudquery/cloudquery/issues/6955)) ([66bfdf3](https://github.com/cloudquery/cloudquery/commit/66bfdf30343f6d8fc1b8cac9675631095ebfc01e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v1.0.3...plugins-destination-azblob-v1.0.4) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.0.6 ([#6858](https://github.com/cloudquery/cloudquery/issues/6858)) ([129f91d](https://github.com/cloudquery/cloudquery/commit/129f91d0a3e05dff5d790887d3b419efde68670c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v1.0.2...plugins-destination-azblob-v1.0.3) (2023-01-11)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes to v1.0.5 ([#6681](https://github.com/cloudquery/cloudquery/issues/6681)) ([32adfa2](https://github.com/cloudquery/cloudquery/commit/32adfa259912f24f555fbb49b45ac697bdb4c9b3))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v1.0.1...plugins-destination-azblob-v1.0.2) (2023-01-10)


### Bug Fixes

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

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-azblob-v1.0.0...plugins-destination-azblob-v1.0.1) (2023-01-03)


### Bug Fixes

* **deps:** Update github.com/pkg/browser digest to 681adbf ([#6193](https://github.com/cloudquery/cloudquery/issues/6193)) ([f0fda44](https://github.com/cloudquery/cloudquery/commit/f0fda4489b6267f63602e831273e893a508d1da2))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/internal to v1.1.2 ([#6205](https://github.com/cloudquery/cloudquery/issues/6205)) ([154fa6f](https://github.com/cloudquery/cloudquery/commit/154fa6fccf41278ac395a1c8287634c4d65926d1))
* **deps:** Update module github.com/cloudquery/filetypes to v1.0.1 ([#6264](https://github.com/cloudquery/cloudquery/issues/6264)) ([da3a1f0](https://github.com/cloudquery/cloudquery/commit/da3a1f0135370e3086bdaed357588955cb0094e8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6222](https://github.com/cloudquery/cloudquery/issues/6222)) ([5ba0d6d](https://github.com/cloudquery/cloudquery/commit/5ba0d6dcdefa9575c361ba7a6cdd86bf985e40c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))
* **deps:** Update module github.com/thoas/go-funk to v0.9.3 ([#6179](https://github.com/cloudquery/cloudquery/issues/6179)) ([e6d89ec](https://github.com/cloudquery/cloudquery/commit/e6d89ec1848f4ca2484ec0f0d7ea5ccaf74c14f5))

## 1.0.0 (2022-12-31)


### Features

* Azure Blob Storage destination ([#6115](https://github.com/cloudquery/cloudquery/issues/6115)) ([d9a62c7](https://github.com/cloudquery/cloudquery/commit/d9a62c75359f1a0c717c60d1d8c863bb31a71a09))

## Changelog
