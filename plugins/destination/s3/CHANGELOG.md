# Changelog

## [6.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v6.0.0...plugins-destination-s3-v6.0.1) (2024-04-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.39.0 ([#17710](https://github.com/cloudquery/cloudquery/issues/17710)) ([e6b3986](https://github.com/cloudquery/cloudquery/commit/e6b39865d674cefb5b001a1c97a25779246087b5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.39.1 ([#17760](https://github.com/cloudquery/cloudquery/issues/17760)) ([7f6faad](https://github.com/cloudquery/cloudquery/commit/7f6faad99e6678d17d449d0da18e0340a2537848))

## [6.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.2.0...plugins-destination-s3-v6.0.0) (2024-04-17)


### ⚠ BREAKING CHANGES

* Set S3 `ContentType` for each object based on format set in `spec` while also allowing users to override and set a custom value ([#17680](https://github.com/cloudquery/cloudquery/issues/17680))

### Features

* Set S3 `ContentType` for each object based on format set in `spec` while also allowing users to override and set a custom value ([#17680](https://github.com/cloudquery/cloudquery/issues/17680)) ([7609be6](https://github.com/cloudquery/cloudquery/commit/7609be605dd864bd66b173afb1ae04abd2a8b0a8)), closes [#17679](https://github.com/cloudquery/cloudquery/issues/17679)

## [5.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.1.3...plugins-destination-s3-v5.2.0) (2024-04-17)


### Features

* Add custom JSON schema errors ([#17669](https://github.com/cloudquery/cloudquery/issues/17669)) ([a3ad426](https://github.com/cloudquery/cloudquery/commit/a3ad426db818fa51bba1107362d6036a2e6a7078))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/codegen to v0.3.14 ([#17658](https://github.com/cloudquery/cloudquery/issues/17658)) ([478eb9c](https://github.com/cloudquery/cloudquery/commit/478eb9c03f764322402703b3975b71b7086a5dea))
* **deps:** Update module github.com/cloudquery/codegen to v0.3.15 ([#17659](https://github.com/cloudquery/cloudquery/issues/17659)) ([58586d0](https://github.com/cloudquery/cloudquery/commit/58586d012a8f4f38b0a693dcbd46d2340bb72a61))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.38.1 ([#17610](https://github.com/cloudquery/cloudquery/issues/17610)) ([a12d17b](https://github.com/cloudquery/cloudquery/commit/a12d17b6f93ef5379b0c11d1338f02dad28f1914))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.38.2 ([#17656](https://github.com/cloudquery/cloudquery/issues/17656)) ([058910b](https://github.com/cloudquery/cloudquery/commit/058910bcb37a6130deb55720a4a1afaec123a319))

## [5.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.1.2...plugins-destination-s3-v5.1.3) (2024-04-09)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#17541](https://github.com/cloudquery/cloudquery/issues/17541)) ([fcb16a4](https://github.com/cloudquery/cloudquery/commit/fcb16a4ea455d7dd8eb62a7adcb7efa62a1f2e77))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.14 ([#17525](https://github.com/cloudquery/cloudquery/issues/17525)) ([4839d53](https://github.com/cloudquery/cloudquery/commit/4839d53643176cc23f69da3680d697b5257a1b8c))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.15 ([#17527](https://github.com/cloudquery/cloudquery/issues/17527)) ([c868fe7](https://github.com/cloudquery/cloudquery/commit/c868fe7375c0e8188fc504b737615a8d11743ccf))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.16 ([#17557](https://github.com/cloudquery/cloudquery/issues/17557)) ([c4a4c02](https://github.com/cloudquery/cloudquery/commit/c4a4c02da469035365198d6ca142a5ba2e266b96))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.4 ([#17485](https://github.com/cloudquery/cloudquery/issues/17485)) ([f370de4](https://github.com/cloudquery/cloudquery/commit/f370de449e61244398e6f413b973cbfa15c019a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.5 ([#17526](https://github.com/cloudquery/cloudquery/issues/17526)) ([554c499](https://github.com/cloudquery/cloudquery/commit/554c499eb9bc9f98f6f3dc4be23fd02049f48dcd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.37.0 ([#17532](https://github.com/cloudquery/cloudquery/issues/17532)) ([8080970](https://github.com/cloudquery/cloudquery/commit/8080970f40d22b6bc9db4c359780c744b476bb02))

## [5.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.1.1...plugins-destination-s3-v5.1.2) (2024-04-02)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#17428](https://github.com/cloudquery/cloudquery/issues/17428)) ([8f855f2](https://github.com/cloudquery/cloudquery/commit/8f855f2d7376962d8223fe1c85593b3c4d12fb8a))
* **deps:** Update module github.com/cloudquery/codegen to v0.3.13 ([#17444](https://github.com/cloudquery/cloudquery/issues/17444)) ([da276fe](https://github.com/cloudquery/cloudquery/commit/da276fe64c46ec0a5f182c83ebc32a90d55f5d50))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.12 ([#17443](https://github.com/cloudquery/cloudquery/issues/17443)) ([737b1e4](https://github.com/cloudquery/cloudquery/commit/737b1e49a251912f4819fdf95fb76681f5d32343))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.13 ([#17460](https://github.com/cloudquery/cloudquery/issues/17460)) ([728497a](https://github.com/cloudquery/cloudquery/commit/728497a6c817522c84923cfd851af78f8c7ab437))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.1 ([#17372](https://github.com/cloudquery/cloudquery/issues/17372)) ([aaf6187](https://github.com/cloudquery/cloudquery/commit/aaf61873ae5d2e01ea5f3b8b319e4f79afb7b29c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.2 ([#17450](https://github.com/cloudquery/cloudquery/issues/17450)) ([2947506](https://github.com/cloudquery/cloudquery/commit/294750650269f8191c6dfff060c4d3a546405763))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.3 ([#17456](https://github.com/cloudquery/cloudquery/issues/17456)) ([020865a](https://github.com/cloudquery/cloudquery/commit/020865a6fde8c896947a844021f0cd7daeb01b06))

## [5.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.1.0...plugins-destination-s3-v5.1.1) (2024-03-26)


### Bug Fixes

* Update docs for new path variable ([#17338](https://github.com/cloudquery/cloudquery/issues/17338)) ([78d001a](https://github.com/cloudquery/cloudquery/commit/78d001a54531e30029894cb38e30e9f849a8da57))

## [5.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.0.4...plugins-destination-s3-v5.1.0) (2024-03-26)


### Features

* S3 expose `{{SyncID}}` as a path variable ([#17286](https://github.com/cloudquery/cloudquery/issues/17286)) ([e7d0fa6](https://github.com/cloudquery/cloudquery/commit/e7d0fa6eb684ee2dd534c1acc94c2cf9fb849184))


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#17314](https://github.com/cloudquery/cloudquery/issues/17314)) ([849fe09](https://github.com/cloudquery/cloudquery/commit/849fe0936786be8642e9d0a0ecc31bd350159774))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.35.0 ([#17299](https://github.com/cloudquery/cloudquery/issues/17299)) ([524ba20](https://github.com/cloudquery/cloudquery/commit/524ba202801c2ae1eb59a5b462a5efc62d1b4000))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.36.0 ([#17325](https://github.com/cloudquery/cloudquery/issues/17325)) ([eb1b4b8](https://github.com/cloudquery/cloudquery/commit/eb1b4b8b963917b87ff644318cfec9745471d50a))

## [5.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.0.3...plugins-destination-s3-v5.0.4) (2024-03-22)


### Bug Fixes

* **deps:** Update github.com/cloudquery/jsonschema digest to 92878fa ([#16718](https://github.com/cloudquery/cloudquery/issues/16718)) ([7fe8588](https://github.com/cloudquery/cloudquery/commit/7fe858818fe1f88fcca6304c873a4614767a57b9))
* Update `_authentication.md` docs ([#17295](https://github.com/cloudquery/cloudquery/issues/17295)) ([99951c2](https://github.com/cloudquery/cloudquery/commit/99951c2e35496e785868faba63162d2700122fd4))

## [5.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.0.2...plugins-destination-s3-v5.0.3) (2024-03-19)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#17221](https://github.com/cloudquery/cloudquery/issues/17221)) ([8400636](https://github.com/cloudquery/cloudquery/commit/840063602eeeb957cfdca93dac043bb2c8138477))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.0 ([#17203](https://github.com/cloudquery/cloudquery/issues/17203)) ([4b128b6](https://github.com/cloudquery/cloudquery/commit/4b128b6722dea883d66458f2f3c831184926353d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.1 ([#17220](https://github.com/cloudquery/cloudquery/issues/17220)) ([08d4950](https://github.com/cloudquery/cloudquery/commit/08d49504aee10f6883e1bd4f7e1102a274c8ee81))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.2 ([#17229](https://github.com/cloudquery/cloudquery/issues/17229)) ([41ed721](https://github.com/cloudquery/cloudquery/commit/41ed721cfa435a4937f3022501dd4d45a3a880b0))
* **deps:** Update module google.golang.org/protobuf to v1.33.0 [SECURITY] ([#17148](https://github.com/cloudquery/cloudquery/issues/17148)) ([170d618](https://github.com/cloudquery/cloudquery/commit/170d6185b9836e4a60df6919a46be92e98f6caa8))

## [5.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.0.1...plugins-destination-s3-v5.0.2) (2024-03-12)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#17097](https://github.com/cloudquery/cloudquery/issues/17097)) ([7d53574](https://github.com/cloudquery/cloudquery/commit/7d535747d8e84b46ab71c16a3c5ef66102b0020a))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.1 ([#17044](https://github.com/cloudquery/cloudquery/issues/17044)) ([d3592e7](https://github.com/cloudquery/cloudquery/commit/d3592e7f3ae600655778eb508aeccfa4e5b74e8c))

## [5.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v5.0.0...plugins-destination-s3-v5.0.1) (2024-03-05)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 814bf88 ([#16977](https://github.com/cloudquery/cloudquery/issues/16977)) ([d4d0e81](https://github.com/cloudquery/cloudquery/commit/d4d0e8138ec10e2c27eb0bf83e88905e838570d0))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to df926f6 ([#16980](https://github.com/cloudquery/cloudquery/issues/16980)) ([4684a2b](https://github.com/cloudquery/cloudquery/commit/4684a2b84b9c0f3c9dfd214b2cda517a46e8a0fb))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to df926f6 ([#16981](https://github.com/cloudquery/cloudquery/issues/16981)) ([4d6cef9](https://github.com/cloudquery/cloudquery/commit/4d6cef9134401b9a6fcd60e70683f1992e526c4d))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.11 ([#16987](https://github.com/cloudquery/cloudquery/issues/16987)) ([05d2f54](https://github.com/cloudquery/cloudquery/commit/05d2f547bf8de4036582e262f2e7c167932eb5a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.0 ([#16957](https://github.com/cloudquery/cloudquery/issues/16957)) ([8ffe2fe](https://github.com/cloudquery/cloudquery/commit/8ffe2fe13a11144cc4f463b01ede1d59c49fcc96))

## [5.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.10.1...plugins-destination-s3-v5.0.0) (2024-02-27)


### ⚠ BREAKING CHANGES

* Properly replace JSON values when using `athena: true` (https://github.com/cloudquery/cloudquery/pull/16942). Previously, `json` columns would have been sanitized & then marshaled twice, resulting in `base64` encoded bytes value. Now, the `json` columns have a proper object value.

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.31.0 ([#16899](https://github.com/cloudquery/cloudquery/issues/16899)) ([2fac27a](https://github.com/cloudquery/cloudquery/commit/2fac27a2e3e789f6152b643c0af1c97ee95c4745))
* Properly replace JSON values when using `athena: true` (https://github.com/cloudquery/cloudquery/pull/16942) ([336eac1](https://github.com/cloudquery/cloudquery/commit/336eac1050aefff0697ee15e1d8b8cce7f8b73a9))

## [4.10.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.10.0...plugins-destination-s3-v4.10.1) (2024-02-27)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#16855](https://github.com/cloudquery/cloudquery/issues/16855)) ([7ced4c8](https://github.com/cloudquery/cloudquery/commit/7ced4c81c90fa76a9be3a2e82efd2d062a6647ee))

## [4.10.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.9.1...plugins-destination-s3-v4.10.0) (2024-02-21)


### Features

* Add kms keys to s3 destination ([#16785](https://github.com/cloudquery/cloudquery/issues/16785)) ([ca23861](https://github.com/cloudquery/cloudquery/commit/ca2386182763839bfdec179c8f04f66953b9a7f7))

## [4.9.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.9.0...plugins-destination-s3-v4.9.1) (2024-02-20)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#16675](https://github.com/cloudquery/cloudquery/issues/16675)) ([c468dda](https://github.com/cloudquery/cloudquery/commit/c468dda51d44c9f8a871605ac8cdcc1cf70ebdef))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.30.0 ([#16669](https://github.com/cloudquery/cloudquery/issues/16669)) ([44b9729](https://github.com/cloudquery/cloudquery/commit/44b9729fa5d7590f65b9073ce4a1fc18a529117e))
* Ensure all writers have a logger ([#16683](https://github.com/cloudquery/cloudquery/issues/16683)) ([c063679](https://github.com/cloudquery/cloudquery/commit/c06367923e2edae62c855733ba4fdd2b3f84e496))

## [4.9.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.9...plugins-destination-s3-v4.9.0) (2024-02-14)


### Features

* Add JSON schema to `s3` destination plugin ([#16465](https://github.com/cloudquery/cloudquery/issues/16465)) ([45a35db](https://github.com/cloudquery/cloudquery/commit/45a35db2f907e86183ae4eed19eaa78791bcbed6))


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#16323](https://github.com/cloudquery/cloudquery/issues/16323)) ([289c4b3](https://github.com/cloudquery/cloudquery/commit/289c4b3a7c2c111bcd7febb5bff9692f3bb96df1))
* **deps:** Update github.com/cloudquery/jsonschema digest to d771afd ([#16483](https://github.com/cloudquery/cloudquery/issues/16483)) ([dcaa994](https://github.com/cloudquery/cloudquery/commit/dcaa9949df43919c0745e05308ce97bf409c4d77))
* **deps:** Update golang.org/x/exp digest to 1b97071 ([#16419](https://github.com/cloudquery/cloudquery/issues/16419)) ([6d77cd1](https://github.com/cloudquery/cloudquery/commit/6d77cd19b6fc648a4ddb12025c22127e960036a4))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 1f4bbc5 ([#16421](https://github.com/cloudquery/cloudquery/issues/16421)) ([9489931](https://github.com/cloudquery/cloudquery/commit/9489931c1b64bf1f7d5da51997944ee54370215b))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 1f4bbc5 ([#16422](https://github.com/cloudquery/cloudquery/issues/16422)) ([74e98fc](https://github.com/cloudquery/cloudquery/commit/74e98fcbde6c6e11baf98284aef0341c597d4817))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.10 ([#16445](https://github.com/cloudquery/cloudquery/issues/16445)) ([9933075](https://github.com/cloudquery/cloudquery/commit/9933075da8c26966c4cd119c30e7a4b5063be9ae))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.6 ([#16356](https://github.com/cloudquery/cloudquery/issues/16356)) ([72298db](https://github.com/cloudquery/cloudquery/commit/72298db12595f15bd432f3b836dc48beff4aacd5))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.7 ([#16366](https://github.com/cloudquery/cloudquery/issues/16366)) ([8459232](https://github.com/cloudquery/cloudquery/commit/8459232dc357ebbe372230c13cef3c678763bef3))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.8 ([#16373](https://github.com/cloudquery/cloudquery/issues/16373)) ([3d0d7f6](https://github.com/cloudquery/cloudquery/commit/3d0d7f69ffd57959d359768ffa4013a97f8499e4))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.9 ([#16424](https://github.com/cloudquery/cloudquery/issues/16424)) ([1b73a7e](https://github.com/cloudquery/cloudquery/commit/1b73a7e8065ad0d540550177ed0d4becd4a536d7))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.0 ([#16237](https://github.com/cloudquery/cloudquery/issues/16237)) ([3fcdab0](https://github.com/cloudquery/cloudquery/commit/3fcdab08816ad9de7bb4eecab59c7be1bda3d00c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.1 ([#16296](https://github.com/cloudquery/cloudquery/issues/16296)) ([ab4a0da](https://github.com/cloudquery/cloudquery/commit/ab4a0dace0a870755fd22d92c6e9c999351f594e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.2 ([#16342](https://github.com/cloudquery/cloudquery/issues/16342)) ([f3eb857](https://github.com/cloudquery/cloudquery/commit/f3eb85729e5db16c2530b31d6d276934866d5ef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.28.0 ([#16362](https://github.com/cloudquery/cloudquery/issues/16362)) ([9166b6b](https://github.com/cloudquery/cloudquery/commit/9166b6b603d0d56a30c2e5072c4f2da5c0c837b5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.0 ([#16395](https://github.com/cloudquery/cloudquery/issues/16395)) ([fb1102e](https://github.com/cloudquery/cloudquery/commit/fb1102eac8af4b3722b82b882187fdf322546513))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.1 ([#16430](https://github.com/cloudquery/cloudquery/issues/16430)) ([738e89f](https://github.com/cloudquery/cloudquery/commit/738e89f2c969a8a3f1698a8686aeaddb358e7a23))
* Use `write_mode` instead of `write-mode` in docs ([#16480](https://github.com/cloudquery/cloudquery/issues/16480)) ([ab2efee](https://github.com/cloudquery/cloudquery/commit/ab2efeeccb7417b75fb1eeb2da266e26adcf7e92))

## [4.8.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.8...plugins-destination-s3-v4.8.9) (2024-01-23)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#16210](https://github.com/cloudquery/cloudquery/issues/16210)) ([dee650e](https://github.com/cloudquery/cloudquery/commit/dee650eafdda5af3526fdc50f4bc1b0f256267bf))

## [4.8.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.7...plugins-destination-s3-v4.8.8) (2024-01-16)


### Bug Fixes

* **deps:** Update aws-sdk-go-v2 monorepo ([#16109](https://github.com/cloudquery/cloudquery/issues/16109)) ([5004149](https://github.com/cloudquery/cloudquery/commit/5004149714f6d56536e910e0868ae14cafb53df6))
* **deps:** Update github.com/apache/arrow/go/v15 digest to 6d44906 ([#16115](https://github.com/cloudquery/cloudquery/issues/16115)) ([8b0ae62](https://github.com/cloudquery/cloudquery/commit/8b0ae6266d19a10fe84102837802358f0b9bb1bc))
* **deps:** Update github.com/apache/arrow/go/v15 digest to 7e703aa ([#16134](https://github.com/cloudquery/cloudquery/issues/16134)) ([72d5eb3](https://github.com/cloudquery/cloudquery/commit/72d5eb35644ce78d775790b0298a0c7690788d28))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.5 ([#16144](https://github.com/cloudquery/cloudquery/issues/16144)) ([e4076f5](https://github.com/cloudquery/cloudquery/commit/e4076f501af25f7f155a75c3ad9393dc9f8f64e5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.1 ([#16069](https://github.com/cloudquery/cloudquery/issues/16069)) ([edda65c](https://github.com/cloudquery/cloudquery/commit/edda65c238b2cb78a7a2078b62557a7d8d822e49))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.2 ([#16130](https://github.com/cloudquery/cloudquery/issues/16130)) ([7ae6f41](https://github.com/cloudquery/cloudquery/commit/7ae6f41957edb3446ff3175857aaf3dcea2cf5bc))

## [4.8.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.6...plugins-destination-s3-v4.8.7) (2024-01-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.4 ([#15941](https://github.com/cloudquery/cloudquery/issues/15941)) ([557886c](https://github.com/cloudquery/cloudquery/commit/557886ca01b5561212cc1c44534b6afd61dac211))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.0 ([#15932](https://github.com/cloudquery/cloudquery/issues/15932)) ([2292b5a](https://github.com/cloudquery/cloudquery/commit/2292b5a2aa5936f2529238a05708de0b3bde9a35))

## [4.8.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.5...plugins-destination-s3-v4.8.6) (2024-01-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 7c3480e ([#15904](https://github.com/cloudquery/cloudquery/issues/15904)) ([a3ec012](https://github.com/cloudquery/cloudquery/commit/a3ec01203183e5c94630beae86434519e87e225d))
* **deps:** Update github.com/gomarkdown/markdown digest to 1d6d208 ([#15907](https://github.com/cloudquery/cloudquery/issues/15907)) ([86d29a9](https://github.com/cloudquery/cloudquery/commit/86d29a900e6c9dbcad09f5b0c4b0615aee59a2ae))
* **deps:** Update golang.org/x/exp digest to 02704c9 ([#15909](https://github.com/cloudquery/cloudquery/issues/15909)) ([dfe32d2](https://github.com/cloudquery/cloudquery/commit/dfe32d2557dcac0fb6dc741c9df4edccdcb07076))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 995d672 ([#15911](https://github.com/cloudquery/cloudquery/issues/15911)) ([18ac2b8](https://github.com/cloudquery/cloudquery/commit/18ac2b806d798e0a9052cc10e8442557ab1c4253))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.3 ([#15900](https://github.com/cloudquery/cloudquery/issues/15900)) ([500197f](https://github.com/cloudquery/cloudquery/commit/500197f536944546560ee8643852efa6524729f2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.24.0 ([#15863](https://github.com/cloudquery/cloudquery/issues/15863)) ([47d7899](https://github.com/cloudquery/cloudquery/commit/47d78994370f083912b6d4329f12d5cef9c255d5))

## [4.8.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.4...plugins-destination-s3-v4.8.5) (2023-12-28)


### Bug Fixes

* **deps:** Update `github.com/apache/arrow/go` to `v15` ([#15754](https://github.com/cloudquery/cloudquery/issues/15754)) ([bd962eb](https://github.com/cloudquery/cloudquery/commit/bd962eb1093cf09e928e2a0e7782288ec4020ec4))
* **deps:** Update aws-sdk-go-v2 monorepo ([#15786](https://github.com/cloudquery/cloudquery/issues/15786)) ([1d1be2a](https://github.com/cloudquery/cloudquery/commit/1d1be2a1a9658448dfea936a85680a55777a2026))
* **deps:** Update github.com/apache/arrow/go/v15 digest to bcaeaa8 ([#15791](https://github.com/cloudquery/cloudquery/issues/15791)) ([89dc812](https://github.com/cloudquery/cloudquery/commit/89dc81201529de2a1fc1ecce5efa74d6f363e57b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.0 ([#15651](https://github.com/cloudquery/cloudquery/issues/15651)) ([6e96125](https://github.com/cloudquery/cloudquery/commit/6e96125a9d9c75616483952edb7a9e402818b264))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.1 ([#15699](https://github.com/cloudquery/cloudquery/issues/15699)) ([67c10c3](https://github.com/cloudquery/cloudquery/commit/67c10c38a04dcdd1512bf6dc739f89bc11baa888))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.2 ([#15724](https://github.com/cloudquery/cloudquery/issues/15724)) ([ad750b1](https://github.com/cloudquery/cloudquery/commit/ad750b1530af06353f2225c7d3397af580093687))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.23.0 ([#15803](https://github.com/cloudquery/cloudquery/issues/15803)) ([b6f9373](https://github.com/cloudquery/cloudquery/commit/b6f937385020c63ce59b2bc60402752b6c239c6c))
* **deps:** Update module golang.org/x/crypto to v0.17.0 [SECURITY] ([#15730](https://github.com/cloudquery/cloudquery/issues/15730)) ([718be50](https://github.com/cloudquery/cloudquery/commit/718be502014ff36aa50cde3a83453b3d6ce15a83))

## [4.8.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.3...plugins-destination-s3-v4.8.4) (2023-12-12)


### Bug Fixes

* **deps:** Update AWS modules ([#15147](https://github.com/cloudquery/cloudquery/issues/15147)) ([037ff43](https://github.com/cloudquery/cloudquery/commit/037ff4314196b9de83d50b61cdc8ce915754b808))
* **deps:** Update AWS modules ([#15365](https://github.com/cloudquery/cloudquery/issues/15365)) ([ad949ec](https://github.com/cloudquery/cloudquery/commit/ad949eca7629dac3ba6143c90814faccae23796a))
* **deps:** Update AWS modules ([#15454](https://github.com/cloudquery/cloudquery/issues/15454)) ([73349c1](https://github.com/cloudquery/cloudquery/commit/73349c1ba1ca56ba4c6eeadc649be93572472e6e))
* **deps:** Update AWS modules ([#15528](https://github.com/cloudquery/cloudquery/issues/15528)) ([2d4b7eb](https://github.com/cloudquery/cloudquery/commit/2d4b7eb6f529a417801bfabd37894f148abe6e00))
* **deps:** Update AWS modules ([#15583](https://github.com/cloudquery/cloudquery/issues/15583)) ([996f222](https://github.com/cloudquery/cloudquery/commit/996f222f33334a16ee2900f663e58bbf1e154909))
* **deps:** Update github.com/gomarkdown/markdown digest to a660076 ([#15517](https://github.com/cloudquery/cloudquery/issues/15517)) ([fa1334c](https://github.com/cloudquery/cloudquery/commit/fa1334c5ce0e157834b0cd676b38af26510fbe43))
* **deps:** Update golang.org/x/exp digest to 6522937 ([#15518](https://github.com/cloudquery/cloudquery/issues/15518)) ([69f9a06](https://github.com/cloudquery/cloudquery/commit/69f9a06754b2feb7c73bd5a80d42fd191c7fdb17))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 3a041ad ([#15520](https://github.com/cloudquery/cloudquery/issues/15520)) ([b2a322a](https://github.com/cloudquery/cloudquery/commit/b2a322a5ec5c1945af5a655c759493a879a9be09))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.1 ([#15536](https://github.com/cloudquery/cloudquery/issues/15536)) ([2751670](https://github.com/cloudquery/cloudquery/commit/27516700723f996d5c585e05a0a08c2d32524298))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.0 ([#15132](https://github.com/cloudquery/cloudquery/issues/15132)) ([81ee138](https://github.com/cloudquery/cloudquery/commit/81ee138ff86c4b92c3ec93208e0a7e05af2b0036))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.1 ([#15263](https://github.com/cloudquery/cloudquery/issues/15263)) ([a9a39ef](https://github.com/cloudquery/cloudquery/commit/a9a39efe8112a564f21c06ba7627fe6c7ced4cdf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.2 ([#15325](https://github.com/cloudquery/cloudquery/issues/15325)) ([77f2db5](https://github.com/cloudquery/cloudquery/commit/77f2db52634bad6e56d970d55172b08d823b97c9))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.3 ([#15355](https://github.com/cloudquery/cloudquery/issues/15355)) ([d8455e5](https://github.com/cloudquery/cloudquery/commit/d8455e5ca1059746c7aced395e9bc150ea495591))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.0 ([#15509](https://github.com/cloudquery/cloudquery/issues/15509)) ([41c689d](https://github.com/cloudquery/cloudquery/commit/41c689d0835487a8d924bb11c989c231f5e3df7c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.1 ([#15539](https://github.com/cloudquery/cloudquery/issues/15539)) ([a298555](https://github.com/cloudquery/cloudquery/commit/a298555343fc7ad483361c2f19c3d39693dab882))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.20.0 ([#15574](https://github.com/cloudquery/cloudquery/issues/15574)) ([317dca4](https://github.com/cloudquery/cloudquery/commit/317dca4182478d6f3789082ae563d9e8bd417d20))

## [4.8.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.2...plugins-destination-s3-v4.8.3) (2023-11-03)


### Bug Fixes

* Athena Compatibility  ([#15123](https://github.com/cloudquery/cloudquery/issues/15123)) ([017941c](https://github.com/cloudquery/cloudquery/commit/017941c7ad6edb229795f7e6e9fff0789f977448))
* **deps:** Update AWS modules ([#15038](https://github.com/cloudquery/cloudquery/issues/15038)) ([f81b943](https://github.com/cloudquery/cloudquery/commit/f81b9431f35c3ddabae61dacac9109cf1144d596))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.2 ([#15034](https://github.com/cloudquery/cloudquery/issues/15034)) ([45c2caa](https://github.com/cloudquery/cloudquery/commit/45c2caa345aa33199ad1592bf378a5a839612c6f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.0 ([#15064](https://github.com/cloudquery/cloudquery/issues/15064)) ([9c2db8c](https://github.com/cloudquery/cloudquery/commit/9c2db8cedaec682a89b444db29e8c0fb45989408))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.1 ([#15075](https://github.com/cloudquery/cloudquery/issues/15075)) ([151769e](https://github.com/cloudquery/cloudquery/commit/151769e7c02028a04ef0ed280951c000ebb1f9c2))
* **deps:** Update module github.com/docker/docker to v24 [SECURITY] ([#15060](https://github.com/cloudquery/cloudquery/issues/15060)) ([41acd0e](https://github.com/cloudquery/cloudquery/commit/41acd0e4ac63221e90cca89a7137a8685692267d))
* **deps:** Update module google.golang.org/grpc to v1.58.3 [SECURITY] ([#14940](https://github.com/cloudquery/cloudquery/issues/14940)) ([e1addea](https://github.com/cloudquery/cloudquery/commit/e1addeaf58ad965e545a3e068860609dadcffa10))

## [4.8.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.1...plugins-destination-s3-v4.8.2) (2023-10-24)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.2.0 ([#14844](https://github.com/cloudquery/cloudquery/issues/14844)) ([f034696](https://github.com/cloudquery/cloudquery/commit/f0346967836a7fc6afffdded045257bb3caac79b))

## [4.8.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.8.0...plugins-destination-s3-v4.8.1) (2023-10-23)


### Bug Fixes

* **deps:** Update AWS modules ([#14800](https://github.com/cloudquery/cloudquery/issues/14800)) ([aced9e5](https://github.com/cloudquery/cloudquery/commit/aced9e59b17cfaa60b9f75eb88f8b2f125bc4484))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to f46436f ([#14803](https://github.com/cloudquery/cloudquery/issues/14803)) ([f5248d7](https://github.com/cloudquery/cloudquery/commit/f5248d749398ded6a50903e09ecabbb996e94a34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.2 ([#14662](https://github.com/cloudquery/cloudquery/issues/14662)) ([e274fe4](https://github.com/cloudquery/cloudquery/commit/e274fe419f6cacdf62547cd7134f40916e5ddd96))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.3 ([#14679](https://github.com/cloudquery/cloudquery/issues/14679)) ([0513c19](https://github.com/cloudquery/cloudquery/commit/0513c193919f4555d41f22ba2ff66efaaf5fca67))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.16.1 ([#14721](https://github.com/cloudquery/cloudquery/issues/14721)) ([1c7ee1d](https://github.com/cloudquery/cloudquery/commit/1c7ee1dc99d7a9cb3358a83e8d827d59be78cefa))
* Set plugin metadata ([#14715](https://github.com/cloudquery/cloudquery/issues/14715)) ([39935e2](https://github.com/cloudquery/cloudquery/commit/39935e2531c4edbd960d5db91e1027b13d7c0a4f))
* Update plugin-SDK to v4.16.0 ([#14702](https://github.com/cloudquery/cloudquery/issues/14702)) ([0dcb545](https://github.com/cloudquery/cloudquery/commit/0dcb5455a71eaa7d28193b1b2fbcdd184dfad2ab))

## [4.8.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.7.3...plugins-destination-s3-v4.8.0) (2023-10-18)


### Features

* Enable Users to Disable TLS Verification when using Custom Endpoint   ([#14192](https://github.com/cloudquery/cloudquery/issues/14192)) ([1762b3a](https://github.com/cloudquery/cloudquery/commit/1762b3a3197d2441c7293c8fa2ef3792e7ed49c5))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to d401686 ([#14459](https://github.com/cloudquery/cloudquery/issues/14459)) ([7ce40f8](https://github.com/cloudquery/cloudquery/commit/7ce40f8dcb1e408c385e877e56b5bb78906b10d2))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to dbcb149 ([#14537](https://github.com/cloudquery/cloudquery/issues/14537)) ([68686f4](https://github.com/cloudquery/cloudquery/commit/68686f4e7636db02bddd961e3d75b60d5218ca85))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.6 ([#14475](https://github.com/cloudquery/cloudquery/issues/14475)) ([83fe7ca](https://github.com/cloudquery/cloudquery/commit/83fe7ca2f5fa83bd3219ddde8fe44fcf1d447480))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.8 ([#14503](https://github.com/cloudquery/cloudquery/issues/14503)) ([4056593](https://github.com/cloudquery/cloudquery/commit/40565937cfc12b33048980b55e91a9a60a62bd47))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.9 ([#14627](https://github.com/cloudquery/cloudquery/issues/14627)) ([c1d244c](https://github.com/cloudquery/cloudquery/commit/c1d244c95199141ac39a713a3f0577b2fb3bf736))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.3.0 ([#14635](https://github.com/cloudquery/cloudquery/issues/14635)) ([00b380c](https://github.com/cloudquery/cloudquery/commit/00b380c10be1642f737f871ba5588888ed5dd180))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.0 ([#14639](https://github.com/cloudquery/cloudquery/issues/14639)) ([f139c0e](https://github.com/cloudquery/cloudquery/commit/f139c0e9369ef92a3cd874003db40b48e229ab58))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.2 ([#14378](https://github.com/cloudquery/cloudquery/issues/14378)) ([a2e0c46](https://github.com/cloudquery/cloudquery/commit/a2e0c4615af4aa205fa082d3f196ea2dc5ce2445))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.3 ([#14436](https://github.com/cloudquery/cloudquery/issues/14436)) ([d529e2d](https://github.com/cloudquery/cloudquery/commit/d529e2d22da93a234492c4165e7eed1257c5767f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.4 ([#14489](https://github.com/cloudquery/cloudquery/issues/14489)) ([9bb45dc](https://github.com/cloudquery/cloudquery/commit/9bb45dc2dacc2c7a6fbd47538b954f731741809b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.5 ([#14516](https://github.com/cloudquery/cloudquery/issues/14516)) ([2d905bf](https://github.com/cloudquery/cloudquery/commit/2d905bf9ea81556282c8ca27dcc6334606a2e83b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.14.0 ([#14577](https://github.com/cloudquery/cloudquery/issues/14577)) ([223c4c1](https://github.com/cloudquery/cloudquery/commit/223c4c1df6c432d7f1bf67a48114e417282bcd0f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.0 ([#14622](https://github.com/cloudquery/cloudquery/issues/14622)) ([b497a6b](https://github.com/cloudquery/cloudquery/commit/b497a6bc5645854bd25d4083fd91ec549a7f274f))
* **deps:** Update module golang.org/x/net to v0.17.0 [SECURITY] ([#14500](https://github.com/cloudquery/cloudquery/issues/14500)) ([9e603d5](https://github.com/cloudquery/cloudquery/commit/9e603d50d28033ed5bf451e569abc7c25014dbfb))

## [4.7.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.7.2...plugins-destination-s3-v4.7.3) (2023-10-04)


### Bug Fixes

* **deps:** Update AWS modules ([#14196](https://github.com/cloudquery/cloudquery/issues/14196)) ([1be729f](https://github.com/cloudquery/cloudquery/commit/1be729fa696d68fa2df88659870fc5623db8e70c))
* **deps:** Update github.com/apache/arrow/go/v14 digest to 00efb06 ([#14202](https://github.com/cloudquery/cloudquery/issues/14202)) ([fc8cc62](https://github.com/cloudquery/cloudquery/commit/fc8cc62ed70db157612e88678c123ba6a34b3b3c))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 7ded38b ([#14246](https://github.com/cloudquery/cloudquery/issues/14246)) ([005891e](https://github.com/cloudquery/cloudquery/commit/005891e1892b41235ddb3b102f4bb6dafd48949a))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.15 ([#14269](https://github.com/cloudquery/cloudquery/issues/14269)) ([b05bb30](https://github.com/cloudquery/cloudquery/commit/b05bb30bf9b359d199856e67a420be86851e44ff))

## [4.7.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.7.1...plugins-destination-s3-v4.7.2) (2023-09-28)


### Bug Fixes

* **deps:** Update AWS modules ([#13872](https://github.com/cloudquery/cloudquery/issues/13872)) ([bef95c0](https://github.com/cloudquery/cloudquery/commit/bef95c08837745bfec7010c948f8e2e41626d802))
* **deps:** Update AWS modules ([#14019](https://github.com/cloudquery/cloudquery/issues/14019)) ([dcd5d1c](https://github.com/cloudquery/cloudquery/commit/dcd5d1ce168993090aa4d2302ea8b39766bc8046))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to ffb7089 ([#13879](https://github.com/cloudquery/cloudquery/issues/13879)) ([f95ced5](https://github.com/cloudquery/cloudquery/commit/f95ced5daa2b123bd71ddff75bd76b3b008790c1))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.12 ([#13949](https://github.com/cloudquery/cloudquery/issues/13949)) ([7bbe086](https://github.com/cloudquery/cloudquery/commit/7bbe086c926223ef499757e6cf196cda21cc8c44))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.13 ([#14037](https://github.com/cloudquery/cloudquery/issues/14037)) ([ca1baf0](https://github.com/cloudquery/cloudquery/commit/ca1baf0d7b653e95ce7dfcc1d15a1a7255bb6630))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.0 ([#13978](https://github.com/cloudquery/cloudquery/issues/13978)) ([2efdf55](https://github.com/cloudquery/cloudquery/commit/2efdf55aed94a14c35c51632ff61ed454caaf5a5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.2 ([#13988](https://github.com/cloudquery/cloudquery/issues/13988)) ([aebaddf](https://github.com/cloudquery/cloudquery/commit/aebaddfc5ca0d7574b8cd72e9e074ec612472dbe))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.0 ([#14031](https://github.com/cloudquery/cloudquery/issues/14031)) ([ac7cdc4](https://github.com/cloudquery/cloudquery/commit/ac7cdc4f7d71599dad89b3170bb7bda676984228))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.1 ([#14063](https://github.com/cloudquery/cloudquery/issues/14063)) ([5a0ff7b](https://github.com/cloudquery/cloudquery/commit/5a0ff7b67890478c371385b379e0a8ef0c2f4865))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.9.0 ([#13960](https://github.com/cloudquery/cloudquery/issues/13960)) ([f074076](https://github.com/cloudquery/cloudquery/commit/f074076a21dc0b8cadfdc3adb9731473d24d28b1))

## [4.7.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.7.0...plugins-destination-s3-v4.7.1) (2023-09-12)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 483f6b2 ([#13780](https://github.com/cloudquery/cloudquery/issues/13780)) ([8d31b44](https://github.com/cloudquery/cloudquery/commit/8d31b44f787f42d47f186cdcc4a5739a3a370a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.0 ([#13625](https://github.com/cloudquery/cloudquery/issues/13625)) ([bb5463f](https://github.com/cloudquery/cloudquery/commit/bb5463fb5919f50f1327eecae884b2ab99fb8b34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.1 ([#13713](https://github.com/cloudquery/cloudquery/issues/13713)) ([73004dc](https://github.com/cloudquery/cloudquery/commit/73004dcabd05bf474d8b5960b8c747a894b98560))

## [4.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.6.4...plugins-destination-s3-v4.7.0) (2023-09-05)


### Features

* Add `use_path_style` ([#13551](https://github.com/cloudquery/cloudquery/issues/13551)) ([0c1b19d](https://github.com/cloudquery/cloudquery/commit/0c1b19d9547d27cbad274212fb8fc9adb9b4191c)), closes [#13260](https://github.com/cloudquery/cloudquery/issues/13260)


### Bug Fixes

* **deps:** Update AWS modules ([#13595](https://github.com/cloudquery/cloudquery/issues/13595)) ([ccb092a](https://github.com/cloudquery/cloudquery/commit/ccb092af5ba9e1b86b256f881db7411dee73fb96))
* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))
* **deps:** Update github.com/apache/arrow/go/v14 digest to a526ba6 ([#13562](https://github.com/cloudquery/cloudquery/issues/13562)) ([248672b](https://github.com/cloudquery/cloudquery/commit/248672beb020828cde1cb608d5c1ed6d656c777b))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to cd3d411 ([#13598](https://github.com/cloudquery/cloudquery/issues/13598)) ([f22bfa6](https://github.com/cloudquery/cloudquery/commit/f22bfa6b2d4fd0caeacf0726ccd307db38f8860c))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.10 ([#13592](https://github.com/cloudquery/cloudquery/issues/13592)) ([cd957e6](https://github.com/cloudquery/cloudquery/commit/cd957e6338ee7be8844fb228cc9f575d8cb17051))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.11 ([#13613](https://github.com/cloudquery/cloudquery/issues/13613)) ([6d87e00](https://github.com/cloudquery/cloudquery/commit/6d87e00400b7b0fcea867b2f636d47f136fbb7b9))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.8 ([#13554](https://github.com/cloudquery/cloudquery/issues/13554)) ([6cecf5d](https://github.com/cloudquery/cloudquery/commit/6cecf5dfd52cd9fe86dc60cd3fc235e380750ffe))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.9 ([#13573](https://github.com/cloudquery/cloudquery/issues/13573)) ([6dd21d5](https://github.com/cloudquery/cloudquery/commit/6dd21d50a0710d929cdf12add95202ff8373e008))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.0 ([#13492](https://github.com/cloudquery/cloudquery/issues/13492)) ([c305876](https://github.com/cloudquery/cloudquery/commit/c305876e3d92944aa6c1a26547f786fdc5b50e23))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.4 ([#13519](https://github.com/cloudquery/cloudquery/issues/13519)) ([9d25165](https://github.com/cloudquery/cloudquery/commit/9d25165820703844c6de96688d939aa5033608ae))

## [4.6.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.6.3...plugins-destination-s3-v4.6.4) (2023-08-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.7 ([#13432](https://github.com/cloudquery/cloudquery/issues/13432)) ([f607482](https://github.com/cloudquery/cloudquery/commit/f60748231dffc91d5fb9eaab5f17d2a8e07666ce))

## [4.6.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.6.2...plugins-destination-s3-v4.6.3) (2023-08-29)


### Bug Fixes

* **deps:** Update `github.com/cloudquery/arrow/go/v13` to `github.com/apache/arrow/go/v14` ([#13341](https://github.com/cloudquery/cloudquery/issues/13341)) ([feb8f87](https://github.com/cloudquery/cloudquery/commit/feb8f87d8d761eb9c49ce84329ad0397f730a918))
* **deps:** Update AWS modules ([#13244](https://github.com/cloudquery/cloudquery/issues/13244)) ([bbe0207](https://github.com/cloudquery/cloudquery/commit/bbe02073701ce165c0bc35ee42d3f789d2840993))
* **deps:** Update AWS modules ([#13339](https://github.com/cloudquery/cloudquery/issues/13339)) ([73d421d](https://github.com/cloudquery/cloudquery/commit/73d421da2437ed69c589e11cfa9bf0eebe7305d2))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 5b83d4f ([#13203](https://github.com/cloudquery/cloudquery/issues/13203)) ([b0a4b8c](https://github.com/cloudquery/cloudquery/commit/b0a4b8ccf7c429bf5a6ed88866865212015b68e4))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.3 ([#13119](https://github.com/cloudquery/cloudquery/issues/13119)) ([5f3e77d](https://github.com/cloudquery/cloudquery/commit/5f3e77df82d1824877b972faaab54518185b2e64))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.4 ([#13261](https://github.com/cloudquery/cloudquery/issues/13261)) ([38dae68](https://github.com/cloudquery/cloudquery/commit/38dae68c9bb08bc78751f9363c50d07d1504a70c))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.5 ([#13284](https://github.com/cloudquery/cloudquery/issues/13284)) ([9ffba02](https://github.com/cloudquery/cloudquery/commit/9ffba02b0f58744d481b7c3b4b32c00ee0f6a9f7))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.1 ([#13195](https://github.com/cloudquery/cloudquery/issues/13195)) ([a184c37](https://github.com/cloudquery/cloudquery/commit/a184c3786ad49df8564344773e9b96f617ef87a1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.5 ([#13285](https://github.com/cloudquery/cloudquery/issues/13285)) ([e076abd](https://github.com/cloudquery/cloudquery/commit/e076abd9d67813a29ced0c1b7b1664fd728b9ba8))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.6 ([#13345](https://github.com/cloudquery/cloudquery/issues/13345)) ([a995a05](https://github.com/cloudquery/cloudquery/commit/a995a0598a209e0fe3ba09f4ced2a052dc14b67a))

## [4.6.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.6.1...plugins-destination-s3-v4.6.2) (2023-08-15)


### Bug Fixes

* **deps:** Update AWS modules ([#13010](https://github.com/cloudquery/cloudquery/issues/13010)) ([96bf61d](https://github.com/cloudquery/cloudquery/commit/96bf61d8335b4b1eb5c65db5d47a9450609ee443))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e9683e1 ([#13015](https://github.com/cloudquery/cloudquery/issues/13015)) ([6557696](https://github.com/cloudquery/cloudquery/commit/65576966d3bd14297499a5b85d3b4fc2c7918df3))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.2 ([#12929](https://github.com/cloudquery/cloudquery/issues/12929)) ([45433e1](https://github.com/cloudquery/cloudquery/commit/45433e11a8dd98aaf436f9346ba8c9d89d87b014))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.4.0 ([#12850](https://github.com/cloudquery/cloudquery/issues/12850)) ([0861200](https://github.com/cloudquery/cloudquery/commit/086120054b45213947e95be954ba6164b9cf6587))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.0 ([#13068](https://github.com/cloudquery/cloudquery/issues/13068)) ([7bb0e4b](https://github.com/cloudquery/cloudquery/commit/7bb0e4ba654971726e16a6a501393e3831170307))

## [4.6.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.6.0...plugins-destination-s3-v4.6.1) (2023-08-08)


### Bug Fixes

* **deps:** Update AWS modules ([#12587](https://github.com/cloudquery/cloudquery/issues/12587)) ([f8e1996](https://github.com/cloudquery/cloudquery/commit/f8e1996df09b24f1e0e6690f483cc2f421115d82))
* **deps:** Update AWS modules ([#12773](https://github.com/cloudquery/cloudquery/issues/12773)) ([1b4376c](https://github.com/cloudquery/cloudquery/commit/1b4376c72eaee5cbabad177e525ebb7417ed3d54))
* **deps:** Update github.com/apache/arrow/go/v13 digest to 112f949 ([#12659](https://github.com/cloudquery/cloudquery/issues/12659)) ([48d73a9](https://github.com/cloudquery/cloudquery/commit/48d73a93e678994f43171c363f5a75c29547b0b9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 3452eb0 ([#12595](https://github.com/cloudquery/cloudquery/issues/12595)) ([c1c0949](https://github.com/cloudquery/cloudquery/commit/c1c09490b17f2e64435e05d745890cdb8b22310d))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f53878d ([#12778](https://github.com/cloudquery/cloudquery/issues/12778)) ([6f5d58e](https://github.com/cloudquery/cloudquery/commit/6f5d58e3b84d3c76b1d1a3d6c5a488f77995a057))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.1 ([#12800](https://github.com/cloudquery/cloudquery/issues/12800)) ([ef33dff](https://github.com/cloudquery/cloudquery/commit/ef33dffcb92e508d0939b54eab4ea62b5689492f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.4 ([#12718](https://github.com/cloudquery/cloudquery/issues/12718)) ([f059a15](https://github.com/cloudquery/cloudquery/commit/f059a159a2ee406ab2b0a33792c244cd217025a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.5 ([#12731](https://github.com/cloudquery/cloudquery/issues/12731)) ([d267239](https://github.com/cloudquery/cloudquery/commit/d267239aa3aca5f94bd36a8db1ec0d9f7dc0865f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.6 ([#12799](https://github.com/cloudquery/cloudquery/issues/12799)) ([fb0e0d7](https://github.com/cloudquery/cloudquery/commit/fb0e0d75ab010f421c834e58d93676de76fcb423))

## [4.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.5.2...plugins-destination-s3-v4.6.0) (2023-07-25)


### Features

* Support compression ([#12411](https://github.com/cloudquery/cloudquery/issues/12411)) ([03fae6d](https://github.com/cloudquery/cloudquery/commit/03fae6d803a4e2ac00ff83d9e90a4169b247c346))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.0.2 ([#12388](https://github.com/cloudquery/cloudquery/issues/12388)) ([af71bb1](https://github.com/cloudquery/cloudquery/commit/af71bb1d35837b0e4cc028525b5408a48d252d02))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.0.3 ([#12422](https://github.com/cloudquery/cloudquery/issues/12422)) ([96c6f27](https://github.com/cloudquery/cloudquery/commit/96c6f27aa019fdbfe3f4cbb3e4c9e74cb0ebec43))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.1.0 ([#12433](https://github.com/cloudquery/cloudquery/issues/12433)) ([6f91a34](https://github.com/cloudquery/cloudquery/commit/6f91a34d571d2361069079cb17f0556ec11d1130))

## [4.5.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.5.1...plugins-destination-s3-v4.5.2) (2023-07-18)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))

## [4.5.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.5.0...plugins-destination-s3-v4.5.1) (2023-07-18)


### Bug Fixes

* **deps:** Update AWS modules ([#12216](https://github.com/cloudquery/cloudquery/issues/12216)) ([d2e77a7](https://github.com/cloudquery/cloudquery/commit/d2e77a7db3da96eef586ff04243e541d68e937dd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.71 ([#12076](https://github.com/cloudquery/cloudquery/issues/12076)) ([4d3f9d8](https://github.com/cloudquery/cloudquery/commit/4d3f9d8db2409bb5905951f7a0f024312891a66e))
* **deps:** Update module github.com/cloudquery/filetypes/v4 to v4.0.1 ([#12140](https://github.com/cloudquery/cloudquery/issues/12140)) ([7321d97](https://github.com/cloudquery/cloudquery/commit/7321d97f12292755ae99e943ac42369a46f6cddf))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.0 ([#12174](https://github.com/cloudquery/cloudquery/issues/12174)) ([80f0289](https://github.com/cloudquery/cloudquery/commit/80f02892a4cf876c4bf4dd4fd9367afb3770ad26))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.1 ([#12185](https://github.com/cloudquery/cloudquery/issues/12185)) ([cfaff16](https://github.com/cloudquery/cloudquery/commit/cfaff16d89800235b6e3015eeb6957d5783d1393))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.0 ([#12256](https://github.com/cloudquery/cloudquery/issues/12256)) ([eaec331](https://github.com/cloudquery/cloudquery/commit/eaec33165345ad51fdb6ddbffbf8a1199ebd6384))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.1 ([#12272](https://github.com/cloudquery/cloudquery/issues/12272)) ([557ca69](https://github.com/cloudquery/cloudquery/commit/557ca69a7dee9dabb80e6afb6f41f205fd8a80d8))
* **deps:** Upgrade destination plugins to SDK v4.0.0 release ([#12137](https://github.com/cloudquery/cloudquery/issues/12137)) ([bf48760](https://github.com/cloudquery/cloudquery/commit/bf48760eef9fe7ce24d73f54bd25da72287a2ed4))

## [4.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.4.2...plugins-destination-s3-v4.5.0) (2023-07-06)


### Features

* **s3:** Upgrade to SDK v4 ([#11700](https://github.com/cloudquery/cloudquery/issues/11700)) ([409729a](https://github.com/cloudquery/cloudquery/commit/409729ac7225d960b98bcad160b8b1ecacbd87e0))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.18.1 ([#11652](https://github.com/cloudquery/cloudquery/issues/11652)) ([4230b52](https://github.com/cloudquery/cloudquery/commit/4230b52a19e91b84fc38348291c371c6c8a735af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.27 ([#11653](https://github.com/cloudquery/cloudquery/issues/11653)) ([4b45408](https://github.com/cloudquery/cloudquery/commit/4b454088055dcbd265e6cbb09420f7dae66865b5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.70 ([#11731](https://github.com/cloudquery/cloudquery/issues/11731)) ([e29d9bb](https://github.com/cloudquery/cloudquery/commit/e29d9bb7b04f077282a06cd45854a753ca90846f))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.2.0 ([#11549](https://github.com/cloudquery/cloudquery/issues/11549)) ([2772f76](https://github.com/cloudquery/cloudquery/commit/2772f7613b35ff909a0fe73c8fb5eb3e051efd62))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.5.0 ([#11850](https://github.com/cloudquery/cloudquery/issues/11850)) ([3255857](https://github.com/cloudquery/cloudquery/commit/3255857938bf16862d52491f5c2a8a0fa53faef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.6 ([#11473](https://github.com/cloudquery/cloudquery/issues/11473)) ([7272133](https://github.com/cloudquery/cloudquery/commit/72721336632e127dd37de4541f2f503bf4f73fb6))
* Update SDK for filetypes dests ([#11990](https://github.com/cloudquery/cloudquery/issues/11990)) ([4897e00](https://github.com/cloudquery/cloudquery/commit/4897e00e1f5068b50463af8de37c7e39dc3cea1a))
* Use configtype.Duration ([#11939](https://github.com/cloudquery/cloudquery/issues/11939)) ([84f8915](https://github.com/cloudquery/cloudquery/commit/84f8915cb68c1ccf9175254505b138de8bc749b7))

## [4.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.4.1...plugins-destination-s3-v4.4.2) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.1.5 ([#11246](https://github.com/cloudquery/cloudquery/issues/11246)) ([a615fc5](https://github.com/cloudquery/cloudquery/commit/a615fc5bd89eb2b26b0c43b88a2080e68a8f545a))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.1.6 ([#11257](https://github.com/cloudquery/cloudquery/issues/11257)) ([eace5ef](https://github.com/cloudquery/cloudquery/commit/eace5ef5669db44a1a4c73241185f9cd6cc405bc))

## [4.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.4.0...plugins-destination-s3-v4.4.1) (2023-06-06)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.1.1 ([#11162](https://github.com/cloudquery/cloudquery/issues/11162)) ([70982b6](https://github.com/cloudquery/cloudquery/commit/70982b6e587735648ccb36690ae7857d85708fc3))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.1.2 ([#11197](https://github.com/cloudquery/cloudquery/issues/11197)) ([409a4dd](https://github.com/cloudquery/cloudquery/commit/409a4dddc4e18da06dcfd493853b28c2a31bcf35))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.1.3 ([#11200](https://github.com/cloudquery/cloudquery/issues/11200)) ([ecd7f68](https://github.com/cloudquery/cloudquery/commit/ecd7f688d56ac999d32027b42e77cb98a0607422))
* **deps:** Update module github.com/cloudquery/filetypes/v3 to v3.1.4 ([#11204](https://github.com/cloudquery/cloudquery/issues/11204)) ([4d83a82](https://github.com/cloudquery/cloudquery/commit/4d83a8235bb780c35706114d1ace51ef885fd9af))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11150](https://github.com/cloudquery/cloudquery/issues/11150)) ([dc00994](https://github.com/cloudquery/cloudquery/commit/dc00994e32936af7e9893c93561d0f9df225a929))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.6.7 ([#11043](https://github.com/cloudquery/cloudquery/issues/11043)) ([3c6d885](https://github.com/cloudquery/cloudquery/commit/3c6d885c3d201b0b39cbc1406c6e54a57ec5ed5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.7.0 ([#11113](https://github.com/cloudquery/cloudquery/issues/11113)) ([487bf87](https://github.com/cloudquery/cloudquery/commit/487bf871afe360cb8d9d592dfea48837d6e7cf27))

## [4.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.3.1...plugins-destination-s3-v4.4.0) (2023-05-25)


### Features

* Update to `github.com/cloudquery/filetypes/v3` `v3.1.0` ([#10942](https://github.com/cloudquery/cloudquery/issues/10942)) ([40ca741](https://github.com/cloudquery/cloudquery/commit/40ca7415fb4149481b6e601c73c5f2019f3353aa))

## [4.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.3.0...plugins-destination-s3-v4.3.1) (2023-05-25)


### Bug Fixes

* **s3:** Update SDK, filetypes to 3.0.1, arrow to latest cqmain ([#10921](https://github.com/cloudquery/cloudquery/issues/10921)) ([cb8faf3](https://github.com/cloudquery/cloudquery/commit/cb8faf3d859032c02890683382903ac36052a8f9))

## [4.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-s3-v4.2.0...plugins-destination-s3-v4.3.0) (2023-05-16)


### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* **s3:** Support `{{FORMAT}}` placeholder ([#10694](https://github.com/cloudquery/cloudquery/issues/10694)) ([f28795d](https://github.com/cloudquery/cloudquery/commit/f28795d933c445c236343197e338ae0239c4a574))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.25 ([#10786](https://github.com/cloudquery/cloudquery/issues/10786)) ([caca1a4](https://github.com/cloudquery/cloudquery/commit/caca1a41e298c06afb6f474b8fd911c4544a2eec))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.67 ([#10788](https://github.com/cloudquery/cloudquery/issues/10788)) ([fd660b2](https://github.com/cloudquery/cloudquery/commit/fd660b25463256ffc4350c2b795bf5138e03fbdb))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

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
