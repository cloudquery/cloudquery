# Changelog

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.14...plugins-destination-clickhouse-v4.0.0) (2024-03-19)


### ⚠ BREAKING CHANGES

* Use documented Apache Arrow type conversion ([#17208](https://github.com/cloudquery/cloudquery/issues/17208)). The following type conversions were made to match the ClickHouse [type convention](https://clickhouse.com/docs/en/sql-reference/formats#data-format-arrow):
    * `date64` is mapped to [`DateTime`](https://clickhouse.com/docs/en/sql-reference/data-types/datetime)
    * `time32` & `time64` are mapped to [`DateTime64`](https://clickhouse.com/docs/en/sql-reference/data-types/datetime64) with proper precision

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.0 ([#17203](https://github.com/cloudquery/cloudquery/issues/17203)) ([4b128b6](https://github.com/cloudquery/cloudquery/commit/4b128b6722dea883d66458f2f3c831184926353d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.1 ([#17220](https://github.com/cloudquery/cloudquery/issues/17220)) ([08d4950](https://github.com/cloudquery/cloudquery/commit/08d49504aee10f6883e1bd4f7e1102a274c8ee81))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.2 ([#17229](https://github.com/cloudquery/cloudquery/issues/17229)) ([41ed721](https://github.com/cloudquery/cloudquery/commit/41ed721cfa435a4937f3022501dd4d45a3a880b0))
* **deps:** Update module google.golang.org/protobuf to v1.33.0 [SECURITY] ([#17133](https://github.com/cloudquery/cloudquery/issues/17133)) ([22461f4](https://github.com/cloudquery/cloudquery/commit/22461f41407d3bc410e965141abc453f1e9d78ca))
* Use documented Apache Arrow type conversion ([#17208](https://github.com/cloudquery/cloudquery/issues/17208)) ([2c24352](https://github.com/cloudquery/cloudquery/commit/2c243523ef874d43ee7452cfdf0d62f6894e78c1))

## [3.4.14](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.13...plugins-destination-clickhouse-v3.4.14) (2024-03-12)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.1 ([#17044](https://github.com/cloudquery/cloudquery/issues/17044)) ([d3592e7](https://github.com/cloudquery/cloudquery/commit/d3592e7f3ae600655778eb508aeccfa4e5b74e8c))

## [3.4.13](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.12...plugins-destination-clickhouse-v3.4.13) (2024-03-05)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 814bf88 ([#16977](https://github.com/cloudquery/cloudquery/issues/16977)) ([d4d0e81](https://github.com/cloudquery/cloudquery/commit/d4d0e8138ec10e2c27eb0bf83e88905e838570d0))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to df926f6 ([#16980](https://github.com/cloudquery/cloudquery/issues/16980)) ([4684a2b](https://github.com/cloudquery/cloudquery/commit/4684a2b84b9c0f3c9dfd214b2cda517a46e8a0fb))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to df926f6 ([#16981](https://github.com/cloudquery/cloudquery/issues/16981)) ([4d6cef9](https://github.com/cloudquery/cloudquery/commit/4d6cef9134401b9a6fcd60e70683f1992e526c4d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.31.0 ([#16899](https://github.com/cloudquery/cloudquery/issues/16899)) ([2fac27a](https://github.com/cloudquery/cloudquery/commit/2fac27a2e3e789f6152b643c0af1c97ee95c4745))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.0 ([#16957](https://github.com/cloudquery/cloudquery/issues/16957)) ([8ffe2fe](https://github.com/cloudquery/cloudquery/commit/8ffe2fe13a11144cc4f463b01ede1d59c49fcc96))

## [3.4.12](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.11...plugins-destination-clickhouse-v3.4.12) (2024-02-27)


### Bug Fixes

* Use Logger for Client rather than the one that is passed in ([#16788](https://github.com/cloudquery/cloudquery/issues/16788)) ([72c19b5](https://github.com/cloudquery/cloudquery/commit/72c19b51840f1ea47067042593e1f651f1801ca4))

## [3.4.11](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.10...plugins-destination-clickhouse-v3.4.11) (2024-02-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.30.0 ([#16669](https://github.com/cloudquery/cloudquery/issues/16669)) ([44b9729](https://github.com/cloudquery/cloudquery/commit/44b9729fa5d7590f65b9073ce4a1fc18a529117e))

## [3.4.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.9...plugins-destination-clickhouse-v3.4.10) (2024-02-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/jsonschema digest to d771afd ([#16483](https://github.com/cloudquery/cloudquery/issues/16483)) ([dcaa994](https://github.com/cloudquery/cloudquery/commit/dcaa9949df43919c0745e05308ce97bf409c4d77))
* **deps:** Update golang.org/x/exp digest to 1b97071 ([#16419](https://github.com/cloudquery/cloudquery/issues/16419)) ([6d77cd1](https://github.com/cloudquery/cloudquery/commit/6d77cd19b6fc648a4ddb12025c22127e960036a4))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 1f4bbc5 ([#16421](https://github.com/cloudquery/cloudquery/issues/16421)) ([9489931](https://github.com/cloudquery/cloudquery/commit/9489931c1b64bf1f7d5da51997944ee54370215b))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 1f4bbc5 ([#16422](https://github.com/cloudquery/cloudquery/issues/16422)) ([74e98fc](https://github.com/cloudquery/cloudquery/commit/74e98fcbde6c6e11baf98284aef0341c597d4817))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.1 ([#16296](https://github.com/cloudquery/cloudquery/issues/16296)) ([ab4a0da](https://github.com/cloudquery/cloudquery/commit/ab4a0dace0a870755fd22d92c6e9c999351f594e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.2 ([#16342](https://github.com/cloudquery/cloudquery/issues/16342)) ([f3eb857](https://github.com/cloudquery/cloudquery/commit/f3eb85729e5db16c2530b31d6d276934866d5ef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.28.0 ([#16362](https://github.com/cloudquery/cloudquery/issues/16362)) ([9166b6b](https://github.com/cloudquery/cloudquery/commit/9166b6b603d0d56a30c2e5072c4f2da5c0c837b5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.0 ([#16395](https://github.com/cloudquery/cloudquery/issues/16395)) ([fb1102e](https://github.com/cloudquery/cloudquery/commit/fb1102eac8af4b3722b82b882187fdf322546513))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.1 ([#16430](https://github.com/cloudquery/cloudquery/issues/16430)) ([738e89f](https://github.com/cloudquery/cloudquery/commit/738e89f2c969a8a3f1698a8686aeaddb358e7a23))
* **deps:** Upgrade plugin-sdk to `v4.27.0` ([#16260](https://github.com/cloudquery/cloudquery/issues/16260)) ([dd7b02f](https://github.com/cloudquery/cloudquery/commit/dd7b02f0ede8be64aeeb72cb881f5406273fc245))

## [3.4.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.8...plugins-destination-clickhouse-v3.4.9) (2024-01-16)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 6d44906 ([#16115](https://github.com/cloudquery/cloudquery/issues/16115)) ([8b0ae62](https://github.com/cloudquery/cloudquery/commit/8b0ae6266d19a10fe84102837802358f0b9bb1bc))
* **deps:** Update github.com/apache/arrow/go/v15 digest to 7e703aa ([#16134](https://github.com/cloudquery/cloudquery/issues/16134)) ([72d5eb3](https://github.com/cloudquery/cloudquery/commit/72d5eb35644ce78d775790b0298a0c7690788d28))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.1 ([#16069](https://github.com/cloudquery/cloudquery/issues/16069)) ([edda65c](https://github.com/cloudquery/cloudquery/commit/edda65c238b2cb78a7a2078b62557a7d8d822e49))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.2 ([#16130](https://github.com/cloudquery/cloudquery/issues/16130)) ([7ae6f41](https://github.com/cloudquery/cloudquery/commit/7ae6f41957edb3446ff3175857aaf3dcea2cf5bc))

## [3.4.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.7...plugins-destination-clickhouse-v3.4.8) (2024-01-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.0 ([#15932](https://github.com/cloudquery/cloudquery/issues/15932)) ([2292b5a](https://github.com/cloudquery/cloudquery/commit/2292b5a2aa5936f2529238a05708de0b3bde9a35))

## [3.4.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.6...plugins-destination-clickhouse-v3.4.7) (2024-01-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 7c3480e ([#15904](https://github.com/cloudquery/cloudquery/issues/15904)) ([a3ec012](https://github.com/cloudquery/cloudquery/commit/a3ec01203183e5c94630beae86434519e87e225d))
* **deps:** Update github.com/gomarkdown/markdown digest to 1d6d208 ([#15907](https://github.com/cloudquery/cloudquery/issues/15907)) ([86d29a9](https://github.com/cloudquery/cloudquery/commit/86d29a900e6c9dbcad09f5b0c4b0615aee59a2ae))
* **deps:** Update golang.org/x/exp digest to 02704c9 ([#15909](https://github.com/cloudquery/cloudquery/issues/15909)) ([dfe32d2](https://github.com/cloudquery/cloudquery/commit/dfe32d2557dcac0fb6dc741c9df4edccdcb07076))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 995d672 ([#15911](https://github.com/cloudquery/cloudquery/issues/15911)) ([18ac2b8](https://github.com/cloudquery/cloudquery/commit/18ac2b806d798e0a9052cc10e8442557ab1c4253))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.24.0 ([#15863](https://github.com/cloudquery/cloudquery/issues/15863)) ([47d7899](https://github.com/cloudquery/cloudquery/commit/47d78994370f083912b6d4329f12d5cef9c255d5))

## [3.4.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.5...plugins-destination-clickhouse-v3.4.6) (2023-12-28)


### Bug Fixes

* **deps:** Update `github.com/apache/arrow/go` to `v15` ([#15754](https://github.com/cloudquery/cloudquery/issues/15754)) ([bd962eb](https://github.com/cloudquery/cloudquery/commit/bd962eb1093cf09e928e2a0e7782288ec4020ec4))
* **deps:** Update github.com/apache/arrow/go/v15 digest to bcaeaa8 ([#15791](https://github.com/cloudquery/cloudquery/issues/15791)) ([89dc812](https://github.com/cloudquery/cloudquery/commit/89dc81201529de2a1fc1ecce5efa74d6f363e57b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.23.0 ([#15803](https://github.com/cloudquery/cloudquery/issues/15803)) ([b6f9373](https://github.com/cloudquery/cloudquery/commit/b6f937385020c63ce59b2bc60402752b6c239c6c))

## [3.4.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.4...plugins-destination-clickhouse-v3.4.5) (2023-12-19)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.0 ([#15651](https://github.com/cloudquery/cloudquery/issues/15651)) ([6e96125](https://github.com/cloudquery/cloudquery/commit/6e96125a9d9c75616483952edb7a9e402818b264))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.1 ([#15699](https://github.com/cloudquery/cloudquery/issues/15699)) ([67c10c3](https://github.com/cloudquery/cloudquery/commit/67c10c38a04dcdd1512bf6dc739f89bc11baa888))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.2 ([#15724](https://github.com/cloudquery/cloudquery/issues/15724)) ([ad750b1](https://github.com/cloudquery/cloudquery/commit/ad750b1530af06353f2225c7d3397af580093687))
* **deps:** Update module golang.org/x/crypto to v0.17.0 [SECURITY] ([#15730](https://github.com/cloudquery/cloudquery/issues/15730)) ([718be50](https://github.com/cloudquery/cloudquery/commit/718be502014ff36aa50cde3a83453b3d6ce15a83))

## [3.4.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.3...plugins-destination-clickhouse-v3.4.4) (2023-12-12)


### Bug Fixes

* **deps:** Update github.com/gomarkdown/markdown digest to a660076 ([#15517](https://github.com/cloudquery/cloudquery/issues/15517)) ([fa1334c](https://github.com/cloudquery/cloudquery/commit/fa1334c5ce0e157834b0cd676b38af26510fbe43))
* **deps:** Update golang.org/x/exp digest to 6522937 ([#15518](https://github.com/cloudquery/cloudquery/issues/15518)) ([69f9a06](https://github.com/cloudquery/cloudquery/commit/69f9a06754b2feb7c73bd5a80d42fd191c7fdb17))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 3a041ad ([#15520](https://github.com/cloudquery/cloudquery/issues/15520)) ([b2a322a](https://github.com/cloudquery/cloudquery/commit/b2a322a5ec5c1945af5a655c759493a879a9be09))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.0 ([#15132](https://github.com/cloudquery/cloudquery/issues/15132)) ([81ee138](https://github.com/cloudquery/cloudquery/commit/81ee138ff86c4b92c3ec93208e0a7e05af2b0036))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.1 ([#15263](https://github.com/cloudquery/cloudquery/issues/15263)) ([a9a39ef](https://github.com/cloudquery/cloudquery/commit/a9a39efe8112a564f21c06ba7627fe6c7ced4cdf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.2 ([#15325](https://github.com/cloudquery/cloudquery/issues/15325)) ([77f2db5](https://github.com/cloudquery/cloudquery/commit/77f2db52634bad6e56d970d55172b08d823b97c9))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.3 ([#15355](https://github.com/cloudquery/cloudquery/issues/15355)) ([d8455e5](https://github.com/cloudquery/cloudquery/commit/d8455e5ca1059746c7aced395e9bc150ea495591))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.0 ([#15509](https://github.com/cloudquery/cloudquery/issues/15509)) ([41c689d](https://github.com/cloudquery/cloudquery/commit/41c689d0835487a8d924bb11c989c231f5e3df7c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.1 ([#15539](https://github.com/cloudquery/cloudquery/issues/15539)) ([a298555](https://github.com/cloudquery/cloudquery/commit/a298555343fc7ad483361c2f19c3d39693dab882))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.20.0 ([#15574](https://github.com/cloudquery/cloudquery/issues/15574)) ([317dca4](https://github.com/cloudquery/cloudquery/commit/317dca4182478d6f3789082ae563d9e8bd417d20))

## [3.4.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.2...plugins-destination-clickhouse-v3.4.3) (2023-11-13)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.2 ([#15034](https://github.com/cloudquery/cloudquery/issues/15034)) ([45c2caa](https://github.com/cloudquery/cloudquery/commit/45c2caa345aa33199ad1592bf378a5a839612c6f))
* **deps:** Update module github.com/cloudquery/codegen to v0.3.12 ([#15084](https://github.com/cloudquery/cloudquery/issues/15084)) ([ff308d5](https://github.com/cloudquery/cloudquery/commit/ff308d5f0696417f037d8f11cd5f398e1d24ac39))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.0 ([#15064](https://github.com/cloudquery/cloudquery/issues/15064)) ([9c2db8c](https://github.com/cloudquery/cloudquery/commit/9c2db8cedaec682a89b444db29e8c0fb45989408))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.1 ([#15075](https://github.com/cloudquery/cloudquery/issues/15075)) ([151769e](https://github.com/cloudquery/cloudquery/commit/151769e7c02028a04ef0ed280951c000ebb1f9c2))
* **deps:** Update module github.com/docker/docker to v24 [SECURITY] ([#15060](https://github.com/cloudquery/cloudquery/issues/15060)) ([41acd0e](https://github.com/cloudquery/cloudquery/commit/41acd0e4ac63221e90cca89a7137a8685692267d))
* **deps:** Update module google.golang.org/grpc to v1.58.3 [SECURITY] ([#14940](https://github.com/cloudquery/cloudquery/issues/14940)) ([e1addea](https://github.com/cloudquery/cloudquery/commit/e1addeaf58ad965e545a3e068860609dadcffa10))

## [3.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.1...plugins-destination-clickhouse-v3.4.2) (2023-10-24)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/codegen to v0.3.11 ([#14870](https://github.com/cloudquery/cloudquery/issues/14870)) ([4fa917d](https://github.com/cloudquery/cloudquery/commit/4fa917d5085b6d99e7818413e507c3fbb32be523))

## [3.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.4.0...plugins-destination-clickhouse-v3.4.1) (2023-10-23)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to f46436f ([#14803](https://github.com/cloudquery/cloudquery/issues/14803)) ([f5248d7](https://github.com/cloudquery/cloudquery/commit/f5248d749398ded6a50903e09ecabbb996e94a34))
* **deps:** Update module github.com/cloudquery/codegen to v0.3.10 ([#14773](https://github.com/cloudquery/cloudquery/issues/14773)) ([98f3e2c](https://github.com/cloudquery/cloudquery/commit/98f3e2c73c94b65f6ae30a55663b6445ebf1146a))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.3 ([#14679](https://github.com/cloudquery/cloudquery/issues/14679)) ([0513c19](https://github.com/cloudquery/cloudquery/commit/0513c193919f4555d41f22ba2ff66efaaf5fca67))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.16.1 ([#14721](https://github.com/cloudquery/cloudquery/issues/14721)) ([1c7ee1d](https://github.com/cloudquery/cloudquery/commit/1c7ee1dc99d7a9cb3358a83e8d827d59be78cefa))
* Set plugin metadata ([#14715](https://github.com/cloudquery/cloudquery/issues/14715)) ([39935e2](https://github.com/cloudquery/cloudquery/commit/39935e2531c4edbd960d5db91e1027b13d7c0a4f))
* Update `engine.parameters` validation ([#14719](https://github.com/cloudquery/cloudquery/issues/14719)) ([715bfe8](https://github.com/cloudquery/cloudquery/commit/715bfe8cbc4cb0400dea454a7ca89a373f8b6f34))
* Update plugin-SDK to v4.16.0 ([#14702](https://github.com/cloudquery/cloudquery/issues/14702)) ([0dcb545](https://github.com/cloudquery/cloudquery/commit/0dcb5455a71eaa7d28193b1b2fbcdd184dfad2ab))

## [3.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.9...plugins-destination-clickhouse-v3.4.0) (2023-10-18)


### Features

* Add JSON schema for plugin spec ([#14623](https://github.com/cloudquery/cloudquery/issues/14623)) ([9bd1185](https://github.com/cloudquery/cloudquery/commit/9bd11856105fd9f2fdc41fa500f2a3e72313e674)), closes [#14616](https://github.com/cloudquery/cloudquery/issues/14616)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to d401686 ([#14459](https://github.com/cloudquery/cloudquery/issues/14459)) ([7ce40f8](https://github.com/cloudquery/cloudquery/commit/7ce40f8dcb1e408c385e877e56b5bb78906b10d2))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to dbcb149 ([#14537](https://github.com/cloudquery/cloudquery/issues/14537)) ([68686f4](https://github.com/cloudquery/cloudquery/commit/68686f4e7636db02bddd961e3d75b60d5218ca85))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.6 ([#14475](https://github.com/cloudquery/cloudquery/issues/14475)) ([83fe7ca](https://github.com/cloudquery/cloudquery/commit/83fe7ca2f5fa83bd3219ddde8fe44fcf1d447480))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.8 ([#14503](https://github.com/cloudquery/cloudquery/issues/14503)) ([4056593](https://github.com/cloudquery/cloudquery/commit/40565937cfc12b33048980b55e91a9a60a62bd47))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.9 ([#14627](https://github.com/cloudquery/cloudquery/issues/14627)) ([c1d244c](https://github.com/cloudquery/cloudquery/commit/c1d244c95199141ac39a713a3f0577b2fb3bf736))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.3.0 ([#14635](https://github.com/cloudquery/cloudquery/issues/14635)) ([00b380c](https://github.com/cloudquery/cloudquery/commit/00b380c10be1642f737f871ba5588888ed5dd180))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.0 ([#14639](https://github.com/cloudquery/cloudquery/issues/14639)) ([f139c0e](https://github.com/cloudquery/cloudquery/commit/f139c0e9369ef92a3cd874003db40b48e229ab58))
* **deps:** Update module github.com/cloudquery/codegen to v0.3.9 ([#14660](https://github.com/cloudquery/cloudquery/issues/14660)) ([68ab0bb](https://github.com/cloudquery/cloudquery/commit/68ab0bb4092f554538aebf892081735fcacb11e7))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.2 ([#14378](https://github.com/cloudquery/cloudquery/issues/14378)) ([a2e0c46](https://github.com/cloudquery/cloudquery/commit/a2e0c4615af4aa205fa082d3f196ea2dc5ce2445))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.3 ([#14436](https://github.com/cloudquery/cloudquery/issues/14436)) ([d529e2d](https://github.com/cloudquery/cloudquery/commit/d529e2d22da93a234492c4165e7eed1257c5767f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.4 ([#14489](https://github.com/cloudquery/cloudquery/issues/14489)) ([9bb45dc](https://github.com/cloudquery/cloudquery/commit/9bb45dc2dacc2c7a6fbd47538b954f731741809b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.5 ([#14516](https://github.com/cloudquery/cloudquery/issues/14516)) ([2d905bf](https://github.com/cloudquery/cloudquery/commit/2d905bf9ea81556282c8ca27dcc6334606a2e83b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.14.0 ([#14577](https://github.com/cloudquery/cloudquery/issues/14577)) ([223c4c1](https://github.com/cloudquery/cloudquery/commit/223c4c1df6c432d7f1bf67a48114e417282bcd0f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.2 ([#14662](https://github.com/cloudquery/cloudquery/issues/14662)) ([e274fe4](https://github.com/cloudquery/cloudquery/commit/e274fe419f6cacdf62547cd7134f40916e5ddd96))
* **deps:** Update module golang.org/x/net to v0.17.0 [SECURITY] ([#14500](https://github.com/cloudquery/cloudquery/issues/14500)) ([9e603d5](https://github.com/cloudquery/cloudquery/commit/9e603d50d28033ed5bf451e569abc7c25014dbfb))

## [3.3.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.8...plugins-destination-clickhouse-v3.3.9) (2023-10-04)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v14 digest to 00efb06 ([#14202](https://github.com/cloudquery/cloudquery/issues/14202)) ([fc8cc62](https://github.com/cloudquery/cloudquery/commit/fc8cc62ed70db157612e88678c123ba6a34b3b3c))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 7ded38b ([#14246](https://github.com/cloudquery/cloudquery/issues/14246)) ([005891e](https://github.com/cloudquery/cloudquery/commit/005891e1892b41235ddb3b102f4bb6dafd48949a))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.0 ([#14281](https://github.com/cloudquery/cloudquery/issues/14281)) ([85835a9](https://github.com/cloudquery/cloudquery/commit/85835a938bfa58d1b0d320ae17aff5fe7f6cfef2))

## [3.3.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.7...plugins-destination-clickhouse-v3.3.8) (2023-09-28)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to ffb7089 ([#13879](https://github.com/cloudquery/cloudquery/issues/13879)) ([f95ced5](https://github.com/cloudquery/cloudquery/commit/f95ced5daa2b123bd71ddff75bd76b3b008790c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.0 ([#13978](https://github.com/cloudquery/cloudquery/issues/13978)) ([2efdf55](https://github.com/cloudquery/cloudquery/commit/2efdf55aed94a14c35c51632ff61ed454caaf5a5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.2 ([#13988](https://github.com/cloudquery/cloudquery/issues/13988)) ([aebaddf](https://github.com/cloudquery/cloudquery/commit/aebaddfc5ca0d7574b8cd72e9e074ec612472dbe))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.0 ([#14031](https://github.com/cloudquery/cloudquery/issues/14031)) ([ac7cdc4](https://github.com/cloudquery/cloudquery/commit/ac7cdc4f7d71599dad89b3170bb7bda676984228))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.1 ([#14063](https://github.com/cloudquery/cloudquery/issues/14063)) ([5a0ff7b](https://github.com/cloudquery/cloudquery/commit/5a0ff7b67890478c371385b379e0a8ef0c2f4865))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.8.0 ([#13950](https://github.com/cloudquery/cloudquery/issues/13950)) ([15b0b69](https://github.com/cloudquery/cloudquery/commit/15b0b6925932613ed2915a3255b3466f21a5c7bf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.9.0 ([#13960](https://github.com/cloudquery/cloudquery/issues/13960)) ([f074076](https://github.com/cloudquery/cloudquery/commit/f074076a21dc0b8cadfdc3adb9731473d24d28b1))

## [3.3.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.6...plugins-destination-clickhouse-v3.3.7) (2023-09-12)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 483f6b2 ([#13780](https://github.com/cloudquery/cloudquery/issues/13780)) ([8d31b44](https://github.com/cloudquery/cloudquery/commit/8d31b44f787f42d47f186cdcc4a5739a3a370a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.0 ([#13625](https://github.com/cloudquery/cloudquery/issues/13625)) ([bb5463f](https://github.com/cloudquery/cloudquery/commit/bb5463fb5919f50f1327eecae884b2ab99fb8b34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.1 ([#13713](https://github.com/cloudquery/cloudquery/issues/13713)) ([73004dc](https://github.com/cloudquery/cloudquery/commit/73004dcabd05bf474d8b5960b8c747a894b98560))

## [3.3.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.5...plugins-destination-clickhouse-v3.3.6) (2023-09-05)


### Bug Fixes

* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))
* **deps:** Update github.com/apache/arrow/go/v14 digest to a526ba6 ([#13562](https://github.com/cloudquery/cloudquery/issues/13562)) ([248672b](https://github.com/cloudquery/cloudquery/commit/248672beb020828cde1cb608d5c1ed6d656c777b))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to cd3d411 ([#13598](https://github.com/cloudquery/cloudquery/issues/13598)) ([f22bfa6](https://github.com/cloudquery/cloudquery/commit/f22bfa6b2d4fd0caeacf0726ccd307db38f8860c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.0 ([#13492](https://github.com/cloudquery/cloudquery/issues/13492)) ([c305876](https://github.com/cloudquery/cloudquery/commit/c305876e3d92944aa6c1a26547f786fdc5b50e23))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.4 ([#13519](https://github.com/cloudquery/cloudquery/issues/13519)) ([9d25165](https://github.com/cloudquery/cloudquery/commit/9d25165820703844c6de96688d939aa5033608ae))

## [3.3.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.4...plugins-destination-clickhouse-v3.3.5) (2023-08-29)


### Bug Fixes

* **deps:** Update `github.com/cloudquery/arrow/go/v13` to `github.com/apache/arrow/go/v14` ([#13341](https://github.com/cloudquery/cloudquery/issues/13341)) ([feb8f87](https://github.com/cloudquery/cloudquery/commit/feb8f87d8d761eb9c49ce84329ad0397f730a918))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 5b83d4f ([#13203](https://github.com/cloudquery/cloudquery/issues/13203)) ([b0a4b8c](https://github.com/cloudquery/cloudquery/commit/b0a4b8ccf7c429bf5a6ed88866865212015b68e4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.1 ([#13195](https://github.com/cloudquery/cloudquery/issues/13195)) ([a184c37](https://github.com/cloudquery/cloudquery/commit/a184c3786ad49df8564344773e9b96f617ef87a1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.5 ([#13285](https://github.com/cloudquery/cloudquery/issues/13285)) ([e076abd](https://github.com/cloudquery/cloudquery/commit/e076abd9d67813a29ced0c1b7b1664fd728b9ba8))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.6 ([#13345](https://github.com/cloudquery/cloudquery/issues/13345)) ([a995a05](https://github.com/cloudquery/cloudquery/commit/a995a0598a209e0fe3ba09f4ced2a052dc14b67a))

## [3.3.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.3...plugins-destination-clickhouse-v3.3.4) (2023-08-15)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e9683e1 ([#13015](https://github.com/cloudquery/cloudquery/issues/13015)) ([6557696](https://github.com/cloudquery/cloudquery/commit/65576966d3bd14297499a5b85d3b4fc2c7918df3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.4.0 ([#12850](https://github.com/cloudquery/cloudquery/issues/12850)) ([0861200](https://github.com/cloudquery/cloudquery/commit/086120054b45213947e95be954ba6164b9cf6587))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.0 ([#13068](https://github.com/cloudquery/cloudquery/issues/13068)) ([7bb0e4b](https://github.com/cloudquery/cloudquery/commit/7bb0e4ba654971726e16a6a501393e3831170307))

## [3.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.2...plugins-destination-clickhouse-v3.3.3) (2023-08-08)


### Bug Fixes

* **ci:** Return structs & maps tests ([#12682](https://github.com/cloudquery/cloudquery/issues/12682)) ([8694a4d](https://github.com/cloudquery/cloudquery/commit/8694a4d50318e1ed88754b00092c8394dce44536))
* **deps:** Update github.com/apache/arrow/go/v13 digest to 112f949 ([#12659](https://github.com/cloudquery/cloudquery/issues/12659)) ([48d73a9](https://github.com/cloudquery/cloudquery/commit/48d73a93e678994f43171c363f5a75c29547b0b9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 3452eb0 ([#12595](https://github.com/cloudquery/cloudquery/issues/12595)) ([c1c0949](https://github.com/cloudquery/cloudquery/commit/c1c09490b17f2e64435e05d745890cdb8b22310d))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f53878d ([#12778](https://github.com/cloudquery/cloudquery/issues/12778)) ([6f5d58e](https://github.com/cloudquery/cloudquery/commit/6f5d58e3b84d3c76b1d1a3d6c5a488f77995a057))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.5 ([#12731](https://github.com/cloudquery/cloudquery/issues/12731)) ([d267239](https://github.com/cloudquery/cloudquery/commit/d267239aa3aca5f94bd36a8db1ec0d9f7dc0865f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.6 ([#12799](https://github.com/cloudquery/cloudquery/issues/12799)) ([fb0e0d7](https://github.com/cloudquery/cloudquery/commit/fb0e0d75ab010f421c834e58d93676de76fcb423))

## [3.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.1...plugins-destination-clickhouse-v3.3.2) (2023-07-25)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))
* **migration:** Make it clear that migration can be done manually and not only via `migrate_mode: forced` ([#12390](https://github.com/cloudquery/cloudquery/issues/12390)) ([33d39cf](https://github.com/cloudquery/cloudquery/commit/33d39cfa87243660241518e23fd5d845ce56d9da))

## [3.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.0...plugins-destination-clickhouse-v3.3.1) (2023-07-18)


### Bug Fixes

* **clickhouse:** Update to SDK v4.2.1 ([#12155](https://github.com/cloudquery/cloudquery/issues/12155)) ([bde9102](https://github.com/cloudquery/cloudquery/commit/bde91021801587138918d137d715558f0a3a2d5f))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))

## [3.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.2.1...plugins-destination-clickhouse-v3.3.0) (2023-07-14)


### Features

* **clickhouse:** Migrate to SDK v4 ([#11794](https://github.com/cloudquery/cloudquery/issues/11794)) ([4bb2d8a](https://github.com/cloudquery/cloudquery/commit/4bb2d8a5327b0e0bc521770e580d28a5f6121d6f)), closes [#11747](https://github.com/cloudquery/cloudquery/issues/11747)
* Use chunks while preparing batch ([#11658](https://github.com/cloudquery/cloudquery/issues/11658)) ([8294e7d](https://github.com/cloudquery/cloudquery/commit/8294e7d28afe6b10ff208757244be3adab30e6fb))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.5.0 ([#11850](https://github.com/cloudquery/cloudquery/issues/11850)) ([3255857](https://github.com/cloudquery/cloudquery/commit/3255857938bf16862d52491f5c2a8a0fa53faef0))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.6 ([#11473](https://github.com/cloudquery/cloudquery/issues/11473)) ([7272133](https://github.com/cloudquery/cloudquery/commit/72721336632e127dd37de4541f2f503bf4f73fb6))
* Use `configtype.Duration` ([#11940](https://github.com/cloudquery/cloudquery/issues/11940)) ([9876bcb](https://github.com/cloudquery/cloudquery/commit/9876bcb0adb3b0ad5c664783d842f729a09c30c4))

## [3.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.2.0...plugins-destination-clickhouse-v3.2.1) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))

## [3.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.1.1...plugins-destination-clickhouse-v3.2.0) (2023-06-06)


### Features

* **nested:** Simplify constructing nested ClickHouse value ([#11205](https://github.com/cloudquery/cloudquery/issues/11205)) ([dd6b275](https://github.com/cloudquery/cloudquery/commit/dd6b2751c723714eeb69c2b152266dfbc07152d8))


### Bug Fixes

* **deps:** Update `github.com/cloudquery/plugin-sdk/v3` to `v3.8.1` ([#11078](https://github.com/cloudquery/cloudquery/issues/11078)) ([0a1764f](https://github.com/cloudquery/cloudquery/commit/0a1764ff0cceb650d914ad5cb0a34e66d33baf3d))
* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))

## [3.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.1.0...plugins-destination-clickhouse-v3.1.1) (2023-05-18)


### Bug Fixes

* Upgrade to plugin-sdk v3.5.2 (Fixes delete-stale for incremental tables) ([#10854](https://github.com/cloudquery/cloudquery/issues/10854)) ([96c17c7](https://github.com/cloudquery/cloudquery/commit/96c17c7ea4a2a455cc3fa52728d818c87e0cff33))

## [3.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.0.0...plugins-destination-clickhouse-v3.1.0) (2023-05-18)


### Features

* **clickhouse:** Migrate to `github.com/cloudquery/plugin-sdk/v3` ([#10807](https://github.com/cloudquery/cloudquery/issues/10807)) ([e0edde7](https://github.com/cloudquery/cloudquery/commit/e0edde7e86956c5fce12cf47bbbe9394d4043652)), closes [#10714](https://github.com/cloudquery/cloudquery/issues/10714)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v2.0.1...plugins-destination-clickhouse-v3.0.0) (2023-05-09)


### ⚠ BREAKING CHANGES

* This change enables [`allow_nullable_key`](https://clickhouse.com/docs/en/operations/settings/settings#allow-nullable-key) for tables ([#10284](https://github.com/cloudquery/cloudquery/issues/10284)).

### Features

* Allow nullable columns in primary keys ([#10284](https://github.com/cloudquery/cloudquery/issues/10284)) ([26e7d0e](https://github.com/cloudquery/cloudquery/commit/26e7d0e9deb5eb2709e3df605a008e8dba4d6552))
* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#10284](https://github.com/cloudquery/cloudquery/issues/10284)) ([26e7d0e](https://github.com/cloudquery/cloudquery/commit/26e7d0e9deb5eb2709e3df605a008e8dba4d6552))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v2.0.0...plugins-destination-clickhouse-v2.0.1) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.3...plugins-destination-clickhouse-v2.0.0) (2023-04-04)


### ⚠ BREAKING CHANGES

* **clickhouse:** Stop reading `ca_cert` value as file path. See [file variable substitution](/docs/advanced-topics/environment-variable-substitution#file-variable-substitution-example) for how to read this value from a file.

### Features

* **clickhouse:** Read only plain `ca_cert` value ([#9495](https://github.com/cloudquery/cloudquery/issues/9495)) ([dcffd50](https://github.com/cloudquery/cloudquery/commit/dcffd506b847ec3634c05fdd4e841764f3434b91))


### Bug Fixes

* **clickhouse:** Check `ca_cert` append result ([#9505](https://github.com/cloudquery/cloudquery/issues/9505)) ([eea1b11](https://github.com/cloudquery/cloudquery/commit/eea1b11151560be38c5413e839c372d5c6eb64a4))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))
* Ignore primary key option when migrating tables ([3a0c68b](https://github.com/cloudquery/cloudquery/commit/3a0c68b59b8b15b3b7b7fa3bb7584b0ad9c5782b))

## [1.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.2...plugins-destination-clickhouse-v1.3.3) (2023-03-23)


### Bug Fixes

* **clickhouse:** Add `ON CLUSTER` for `DROP TABLE` statement ([#9377](https://github.com/cloudquery/cloudquery/issues/9377)) ([76a74ff](https://github.com/cloudquery/cloudquery/commit/76a74ffb3479bc7c086b33020d665a28bdf75db5))

## [1.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.1...plugins-destination-clickhouse-v1.3.2) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.0...plugins-destination-clickhouse-v1.3.1) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.2.1...plugins-destination-clickhouse-v1.3.0) (2023-03-10)


### Features

* **clickhouse:** Add table engine option ([#8844](https://github.com/cloudquery/cloudquery/issues/8844)) ([447b29c](https://github.com/cloudquery/cloudquery/commit/447b29c172129a2c4a24fb81053a54fb9c6d8103)), closes [#8667](https://github.com/cloudquery/cloudquery/issues/8667)

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.2.0...plugins-destination-clickhouse-v1.2.1) (2023-03-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.3...plugins-destination-clickhouse-v1.2.0) (2023-03-01)


### Features

* **clickhouse:** Support distributed DDL ([#8663](https://github.com/cloudquery/cloudquery/issues/8663)) ([c46705f](https://github.com/cloudquery/cloudquery/commit/c46705f02cec99fd573ed1d1721921c58d1f4cab)), closes [#8654](https://github.com/cloudquery/cloudquery/issues/8654)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to c95f2b4 ([#8560](https://github.com/cloudquery/cloudquery/issues/8560)) ([9c3bd5b](https://github.com/cloudquery/cloudquery/commit/9c3bd5b68f9741a360fde6c54bf3f5f3efe06d9e))
* **deps:** Update module github.com/andybalholm/brotli to v1.0.5 ([#8570](https://github.com/cloudquery/cloudquery/issues/8570)) ([1251c4d](https://github.com/cloudquery/cloudquery/commit/1251c4dd228cee7d34af4e9ec8df1e9ccfb41e3e))
* **deps:** Update module github.com/ClickHouse/ch-go to v0.53.0 ([#8652](https://github.com/cloudquery/cloudquery/issues/8652)) ([a016609](https://github.com/cloudquery/cloudquery/commit/a0166095c8b57330d1ba292848f7df3c09728032))
* **deps:** Update module github.com/ClickHouse/clickhouse-go/v2 to v2.6.5 ([#8568](https://github.com/cloudquery/cloudquery/issues/8568)) ([d553b70](https://github.com/cloudquery/cloudquery/commit/d553b700a05bb0c0d8a59f74f454b0c46371a6b7))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [1.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.2...plugins-destination-clickhouse-v1.1.3) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [1.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.1...plugins-destination-clickhouse-v1.1.2) (2023-02-23)


### Bug Fixes

* **clickhouse:** Bump minimum ClickHouse version ([#8406](https://github.com/cloudquery/cloudquery/issues/8406)) ([a5890b2](https://github.com/cloudquery/cloudquery/commit/a5890b2d304b06b3460612c52980ac60dfcf6058))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.0...plugins-destination-clickhouse-v1.1.1) (2023-02-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **migration:** Don't add trailing comma to create table ([#8402](https://github.com/cloudquery/cloudquery/issues/8402)) ([1182b21](https://github.com/cloudquery/cloudquery/commit/1182b2194cd87d9f61c35aac3acd5ba25ec352da))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.0.0...plugins-destination-clickhouse-v1.1.0) (2023-02-21)


### Features

* **deps:** Update ClickHouse plugin-sdk to v1.38.2 ([#8255](https://github.com/cloudquery/cloudquery/issues/8255)) ([ddbc004](https://github.com/cloudquery/cloudquery/commit/ddbc004eb7e65c756929afd84758d756aba7549b))


### Bug Fixes

* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## 1.0.0 (2023-02-14)


### Features

* **clickhouse:** ClickHouse destination ([#6982](https://github.com/cloudquery/cloudquery/issues/6982)) ([09411e4](https://github.com/cloudquery/cloudquery/commit/09411e45c01609b986f0cc6f42554096ae5558dc)), closes [#6254](https://github.com/cloudquery/cloudquery/issues/6254)
