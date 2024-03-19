# Changelog

## [5.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.4.3...plugins-destination-duckdb-v5.5.0) (2024-03-19)


### Features

* Automatically handle dropping Unique clause ([#17206](https://github.com/cloudquery/cloudquery/issues/17206)) ([8940f8a](https://github.com/cloudquery/cloudquery/commit/8940f8abd016ec863a573a3375f50fb82002e983))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.0 ([#17203](https://github.com/cloudquery/cloudquery/issues/17203)) ([4b128b6](https://github.com/cloudquery/cloudquery/commit/4b128b6722dea883d66458f2f3c831184926353d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.1 ([#17220](https://github.com/cloudquery/cloudquery/issues/17220)) ([08d4950](https://github.com/cloudquery/cloudquery/commit/08d49504aee10f6883e1bd4f7e1102a274c8ee81))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.34.2 ([#17229](https://github.com/cloudquery/cloudquery/issues/17229)) ([41ed721](https://github.com/cloudquery/cloudquery/commit/41ed721cfa435a4937f3022501dd4d45a3a880b0))
* **deps:** Update module google.golang.org/protobuf to v1.33.0 [SECURITY] ([#17134](https://github.com/cloudquery/cloudquery/issues/17134)) ([8b03c33](https://github.com/cloudquery/cloudquery/commit/8b03c3392627edde7682b08e0c5e6691638c6ba8))

## [5.4.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.4.2...plugins-destination-duckdb-v5.4.3) (2024-03-12)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.1 ([#17044](https://github.com/cloudquery/cloudquery/issues/17044)) ([d3592e7](https://github.com/cloudquery/cloudquery/commit/d3592e7f3ae600655778eb508aeccfa4e5b74e8c))

## [5.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.4.1...plugins-destination-duckdb-v5.4.2) (2024-03-05)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 814bf88 ([#16977](https://github.com/cloudquery/cloudquery/issues/16977)) ([d4d0e81](https://github.com/cloudquery/cloudquery/commit/d4d0e8138ec10e2c27eb0bf83e88905e838570d0))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to df926f6 ([#16980](https://github.com/cloudquery/cloudquery/issues/16980)) ([4684a2b](https://github.com/cloudquery/cloudquery/commit/4684a2b84b9c0f3c9dfd214b2cda517a46e8a0fb))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to df926f6 ([#16981](https://github.com/cloudquery/cloudquery/issues/16981)) ([4d6cef9](https://github.com/cloudquery/cloudquery/commit/4d6cef9134401b9a6fcd60e70683f1992e526c4d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.31.0 ([#16899](https://github.com/cloudquery/cloudquery/issues/16899)) ([2fac27a](https://github.com/cloudquery/cloudquery/commit/2fac27a2e3e789f6152b643c0af1c97ee95c4745))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.32.0 ([#16957](https://github.com/cloudquery/cloudquery/issues/16957)) ([8ffe2fe](https://github.com/cloudquery/cloudquery/commit/8ffe2fe13a11144cc4f463b01ede1d59c49fcc96))

## [5.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.4.0...plugins-destination-duckdb-v5.4.1) (2024-02-26)


### Bug Fixes

* Downgrade DuckDB library version ([#16870](https://github.com/cloudquery/cloudquery/issues/16870)) ([09305ef](https://github.com/cloudquery/cloudquery/commit/09305ef59adc2683ec2471e328e5e45f58d9e5eb))

## [5.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.3.0...plugins-destination-duckdb-v5.4.0) (2024-02-22)


### Features

* Use Appender interface ([#16668](https://github.com/cloudquery/cloudquery/issues/16668)) ([8192d6d](https://github.com/cloudquery/cloudquery/commit/8192d6d99898692e9232d29d9ac71f5d2c6d8d4a))

## [5.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.2.1...plugins-destination-duckdb-v5.3.0) (2024-02-21)


### Features

* Add binary and large type handling ([#16772](https://github.com/cloudquery/cloudquery/issues/16772)) ([7a2a4d6](https://github.com/cloudquery/cloudquery/commit/7a2a4d6e201c6890e057a8ecebe009aa94e6c671))


### Bug Fixes

* Sanitize ID in queries, better debug logging ([#16727](https://github.com/cloudquery/cloudquery/issues/16727)) ([05f0dcd](https://github.com/cloudquery/cloudquery/commit/05f0dcdb9b3a799729a0d9d4e950b100fb104838))

## [5.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.2.0...plugins-destination-duckdb-v5.2.1) (2024-02-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.30.0 ([#16669](https://github.com/cloudquery/cloudquery/issues/16669)) ([44b9729](https://github.com/cloudquery/cloudquery/commit/44b9729fa5d7590f65b9073ce4a1fc18a529117e))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.6.0 ([#16649](https://github.com/cloudquery/cloudquery/issues/16649)) ([5529419](https://github.com/cloudquery/cloudquery/commit/55294190b3b8dc5c4f5a90e328244fa1cad51a57))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.6.1 ([#16674](https://github.com/cloudquery/cloudquery/issues/16674)) ([926bee8](https://github.com/cloudquery/cloudquery/commit/926bee8e7ac48e44c1cad58e80d219d9435068ee))
* Ensure all writers have a logger ([#16683](https://github.com/cloudquery/cloudquery/issues/16683)) ([c063679](https://github.com/cloudquery/cloudquery/commit/c06367923e2edae62c855733ba4fdd2b3f84e496))

## [5.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.1.1...plugins-destination-duckdb-v5.2.0) (2024-02-16)


### Features

* Add user agent if MotherDuck connection ([#16651](https://github.com/cloudquery/cloudquery/issues/16651)) ([fdd0f21](https://github.com/cloudquery/cloudquery/commit/fdd0f21d67b38a8a799d09a29da57e66bc47ba19))

## [5.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.1.0...plugins-destination-duckdb-v5.1.1) (2024-02-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/jsonschema digest to d771afd ([#16483](https://github.com/cloudquery/cloudquery/issues/16483)) ([dcaa994](https://github.com/cloudquery/cloudquery/commit/dcaa9949df43919c0745e05308ce97bf409c4d77))

## [5.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.12...plugins-destination-duckdb-v5.1.0) (2024-02-01)


### Features

* Add JSON schema to DuckDB destination plugin ([#16372](https://github.com/cloudquery/cloudquery/issues/16372)) ([1ace351](https://github.com/cloudquery/cloudquery/commit/1ace3516342dd8a9e663b02b47e3edf28811de34))


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 1b97071 ([#16419](https://github.com/cloudquery/cloudquery/issues/16419)) ([6d77cd1](https://github.com/cloudquery/cloudquery/commit/6d77cd19b6fc648a4ddb12025c22127e960036a4))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 1f4bbc5 ([#16421](https://github.com/cloudquery/cloudquery/issues/16421)) ([9489931](https://github.com/cloudquery/cloudquery/commit/9489931c1b64bf1f7d5da51997944ee54370215b))
* **deps:** Update google.golang.org/genproto/googleapis/rpc digest to 1f4bbc5 ([#16422](https://github.com/cloudquery/cloudquery/issues/16422)) ([74e98fc](https://github.com/cloudquery/cloudquery/commit/74e98fcbde6c6e11baf98284aef0341c597d4817))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.28.0 ([#16362](https://github.com/cloudquery/cloudquery/issues/16362)) ([9166b6b](https://github.com/cloudquery/cloudquery/commit/9166b6b603d0d56a30c2e5072c4f2da5c0c837b5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.0 ([#16395](https://github.com/cloudquery/cloudquery/issues/16395)) ([fb1102e](https://github.com/cloudquery/cloudquery/commit/fb1102eac8af4b3722b82b882187fdf322546513))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.29.1 ([#16430](https://github.com/cloudquery/cloudquery/issues/16430)) ([738e89f](https://github.com/cloudquery/cloudquery/commit/738e89f2c969a8a3f1698a8686aeaddb358e7a23))

## [5.0.12](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.11...plugins-destination-duckdb-v5.0.12) (2024-01-29)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.0 ([#16237](https://github.com/cloudquery/cloudquery/issues/16237)) ([3fcdab0](https://github.com/cloudquery/cloudquery/commit/3fcdab08816ad9de7bb4eecab59c7be1bda3d00c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.27.1 ([#16296](https://github.com/cloudquery/cloudquery/issues/16296)) ([ab4a0da](https://github.com/cloudquery/cloudquery/commit/ab4a0dace0a870755fd22d92c6e9c999351f594e))
* Don't use `CGO` unless specifically selected to ([#16329](https://github.com/cloudquery/cloudquery/issues/16329)) ([597267a](https://github.com/cloudquery/cloudquery/commit/597267a221f57c7d381ece16faff59949fc0e717))

## [5.0.11](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.10...plugins-destination-duckdb-v5.0.11) (2024-01-16)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 6d44906 ([#16115](https://github.com/cloudquery/cloudquery/issues/16115)) ([8b0ae62](https://github.com/cloudquery/cloudquery/commit/8b0ae6266d19a10fe84102837802358f0b9bb1bc))
* **deps:** Update github.com/apache/arrow/go/v15 digest to 7e703aa ([#16134](https://github.com/cloudquery/cloudquery/issues/16134)) ([72d5eb3](https://github.com/cloudquery/cloudquery/commit/72d5eb35644ce78d775790b0298a0c7690788d28))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.1 ([#16069](https://github.com/cloudquery/cloudquery/issues/16069)) ([edda65c](https://github.com/cloudquery/cloudquery/commit/edda65c238b2cb78a7a2078b62557a7d8d822e49))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.2 ([#16130](https://github.com/cloudquery/cloudquery/issues/16130)) ([7ae6f41](https://github.com/cloudquery/cloudquery/commit/7ae6f41957edb3446ff3175857aaf3dcea2cf5bc))

## [5.0.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.9...plugins-destination-duckdb-v5.0.10) (2024-01-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.25.0 ([#15932](https://github.com/cloudquery/cloudquery/issues/15932)) ([2292b5a](https://github.com/cloudquery/cloudquery/commit/2292b5a2aa5936f2529238a05708de0b3bde9a35))

## [5.0.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.8...plugins-destination-duckdb-v5.0.9) (2024-01-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v15 digest to 7c3480e ([#15904](https://github.com/cloudquery/cloudquery/issues/15904)) ([a3ec012](https://github.com/cloudquery/cloudquery/commit/a3ec01203183e5c94630beae86434519e87e225d))
* **deps:** Update github.com/gomarkdown/markdown digest to 1d6d208 ([#15907](https://github.com/cloudquery/cloudquery/issues/15907)) ([86d29a9](https://github.com/cloudquery/cloudquery/commit/86d29a900e6c9dbcad09f5b0c4b0615aee59a2ae))
* **deps:** Update golang.org/x/exp digest to 02704c9 ([#15909](https://github.com/cloudquery/cloudquery/issues/15909)) ([dfe32d2](https://github.com/cloudquery/cloudquery/commit/dfe32d2557dcac0fb6dc741c9df4edccdcb07076))
* **deps:** Update google.golang.org/genproto/googleapis/api digest to 995d672 ([#15911](https://github.com/cloudquery/cloudquery/issues/15911)) ([18ac2b8](https://github.com/cloudquery/cloudquery/commit/18ac2b806d798e0a9052cc10e8442557ab1c4253))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.24.0 ([#15863](https://github.com/cloudquery/cloudquery/issues/15863)) ([47d7899](https://github.com/cloudquery/cloudquery/commit/47d78994370f083912b6d4329f12d5cef9c255d5))

## [5.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.7...plugins-destination-duckdb-v5.0.8) (2023-12-28)


### Bug Fixes

* **deps:** Update `github.com/apache/arrow/go` to `v15` ([#15754](https://github.com/cloudquery/cloudquery/issues/15754)) ([bd962eb](https://github.com/cloudquery/cloudquery/commit/bd962eb1093cf09e928e2a0e7782288ec4020ec4))
* **deps:** Update github.com/apache/arrow/go/v15 digest to bcaeaa8 ([#15791](https://github.com/cloudquery/cloudquery/issues/15791)) ([89dc812](https://github.com/cloudquery/cloudquery/commit/89dc81201529de2a1fc1ecce5efa74d6f363e57b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.23.0 ([#15803](https://github.com/cloudquery/cloudquery/issues/15803)) ([b6f9373](https://github.com/cloudquery/cloudquery/commit/b6f937385020c63ce59b2bc60402752b6c239c6c))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.5.5 ([#15775](https://github.com/cloudquery/cloudquery/issues/15775)) ([0fdd168](https://github.com/cloudquery/cloudquery/commit/0fdd168a2437bc4a3e0e85905c20421b39cd5818))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.5.6 ([#15790](https://github.com/cloudquery/cloudquery/issues/15790)) ([e19781f](https://github.com/cloudquery/cloudquery/commit/e19781ffab1908d48096d07beceb8c322ad61d36))

## [5.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.6...plugins-destination-duckdb-v5.0.7) (2023-12-19)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.0 ([#15651](https://github.com/cloudquery/cloudquery/issues/15651)) ([6e96125](https://github.com/cloudquery/cloudquery/commit/6e96125a9d9c75616483952edb7a9e402818b264))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.1 ([#15699](https://github.com/cloudquery/cloudquery/issues/15699)) ([67c10c3](https://github.com/cloudquery/cloudquery/commit/67c10c38a04dcdd1512bf6dc739f89bc11baa888))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.21.2 ([#15724](https://github.com/cloudquery/cloudquery/issues/15724)) ([ad750b1](https://github.com/cloudquery/cloudquery/commit/ad750b1530af06353f2225c7d3397af580093687))
* **deps:** Update module golang.org/x/crypto to v0.17.0 [SECURITY] ([#15730](https://github.com/cloudquery/cloudquery/issues/15730)) ([718be50](https://github.com/cloudquery/cloudquery/commit/718be502014ff36aa50cde3a83453b3d6ce15a83))

## [5.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.5...plugins-destination-duckdb-v5.0.6) (2023-12-12)


### Bug Fixes

* **deps:** Update github.com/gomarkdown/markdown digest to a660076 ([#15517](https://github.com/cloudquery/cloudquery/issues/15517)) ([fa1334c](https://github.com/cloudquery/cloudquery/commit/fa1334c5ce0e157834b0cd676b38af26510fbe43))
* **deps:** Update golang.org/x/exp digest to 6522937 ([#15518](https://github.com/cloudquery/cloudquery/issues/15518)) ([69f9a06](https://github.com/cloudquery/cloudquery/commit/69f9a06754b2feb7c73bd5a80d42fd191c7fdb17))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.0 ([#15132](https://github.com/cloudquery/cloudquery/issues/15132)) ([81ee138](https://github.com/cloudquery/cloudquery/commit/81ee138ff86c4b92c3ec93208e0a7e05af2b0036))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.1 ([#15263](https://github.com/cloudquery/cloudquery/issues/15263)) ([a9a39ef](https://github.com/cloudquery/cloudquery/commit/a9a39efe8112a564f21c06ba7627fe6c7ced4cdf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.2 ([#15325](https://github.com/cloudquery/cloudquery/issues/15325)) ([77f2db5](https://github.com/cloudquery/cloudquery/commit/77f2db52634bad6e56d970d55172b08d823b97c9))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.18.3 ([#15355](https://github.com/cloudquery/cloudquery/issues/15355)) ([d8455e5](https://github.com/cloudquery/cloudquery/commit/d8455e5ca1059746c7aced395e9bc150ea495591))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.0 ([#15509](https://github.com/cloudquery/cloudquery/issues/15509)) ([41c689d](https://github.com/cloudquery/cloudquery/commit/41c689d0835487a8d924bb11c989c231f5e3df7c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.19.1 ([#15539](https://github.com/cloudquery/cloudquery/issues/15539)) ([a298555](https://github.com/cloudquery/cloudquery/commit/a298555343fc7ad483361c2f19c3d39693dab882))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.20.0 ([#15574](https://github.com/cloudquery/cloudquery/issues/15574)) ([317dca4](https://github.com/cloudquery/cloudquery/commit/317dca4182478d6f3789082ae563d9e8bd417d20))
* **deps:** Update module github.com/docker/docker to v24 [SECURITY] ([#15060](https://github.com/cloudquery/cloudquery/issues/15060)) ([41acd0e](https://github.com/cloudquery/cloudquery/commit/41acd0e4ac63221e90cca89a7137a8685692267d))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.5.2 ([#15280](https://github.com/cloudquery/cloudquery/issues/15280)) ([0d43010](https://github.com/cloudquery/cloudquery/commit/0d43010f754a612ef4badfd4e2cb798c604ffd20))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.5.3 ([#15436](https://github.com/cloudquery/cloudquery/issues/15436)) ([f26f618](https://github.com/cloudquery/cloudquery/commit/f26f6183e0abae6ab5b13b74dcb2c6182fa7a683))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.5.4 ([#15498](https://github.com/cloudquery/cloudquery/issues/15498)) ([d018c6d](https://github.com/cloudquery/cloudquery/commit/d018c6d56d85a0254c8a2c5ab219acd7098f8eee))

## [5.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.4...plugins-destination-duckdb-v5.0.5) (2023-11-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.2 ([#15034](https://github.com/cloudquery/cloudquery/issues/15034)) ([45c2caa](https://github.com/cloudquery/cloudquery/commit/45c2caa345aa33199ad1592bf378a5a839612c6f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.0 ([#15064](https://github.com/cloudquery/cloudquery/issues/15064)) ([9c2db8c](https://github.com/cloudquery/cloudquery/commit/9c2db8cedaec682a89b444db29e8c0fb45989408))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.17.1 ([#15075](https://github.com/cloudquery/cloudquery/issues/15075)) ([151769e](https://github.com/cloudquery/cloudquery/commit/151769e7c02028a04ef0ed280951c000ebb1f9c2))
* **deps:** Update module google.golang.org/grpc to v1.58.3 [SECURITY] ([#14940](https://github.com/cloudquery/cloudquery/issues/14940)) ([e1addea](https://github.com/cloudquery/cloudquery/commit/e1addeaf58ad965e545a3e068860609dadcffa10))

## [5.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.3...plugins-destination-duckdb-v5.0.4) (2023-10-23)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to f46436f ([#14803](https://github.com/cloudquery/cloudquery/issues/14803)) ([f5248d7](https://github.com/cloudquery/cloudquery/commit/f5248d749398ded6a50903e09ecabbb996e94a34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.16.1 ([#14721](https://github.com/cloudquery/cloudquery/issues/14721)) ([1c7ee1d](https://github.com/cloudquery/cloudquery/commit/1c7ee1dc99d7a9cb3358a83e8d827d59be78cefa))
* Set plugin metadata ([#14715](https://github.com/cloudquery/cloudquery/issues/14715)) ([39935e2](https://github.com/cloudquery/cloudquery/commit/39935e2531c4edbd960d5db91e1027b13d7c0a4f))
* Update plugin-SDK to v4.16.0 ([#14702](https://github.com/cloudquery/cloudquery/issues/14702)) ([0dcb545](https://github.com/cloudquery/cloudquery/commit/0dcb5455a71eaa7d28193b1b2fbcdd184dfad2ab))

## [5.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.2...plugins-destination-duckdb-v5.0.3) (2023-10-19)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.3 ([#14679](https://github.com/cloudquery/cloudquery/issues/14679)) ([0513c19](https://github.com/cloudquery/cloudquery/commit/0513c193919f4555d41f22ba2ff66efaaf5fca67))

## [5.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.1...plugins-destination-duckdb-v5.0.2) (2023-10-18)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.2 ([#14662](https://github.com/cloudquery/cloudquery/issues/14662)) ([e274fe4](https://github.com/cloudquery/cloudquery/commit/e274fe419f6cacdf62547cd7134f40916e5ddd96))

## [5.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v5.0.0...plugins-destination-duckdb-v5.0.1) (2023-10-18)


### Bug Fixes

* Compile DuckDB, SQLite and Snowflake plugins for running under Alpine Linux  ([#14612](https://github.com/cloudquery/cloudquery/issues/14612)) ([c47cd1c](https://github.com/cloudquery/cloudquery/commit/c47cd1cc8bb014d654097087e81a9d658ea8f1dc))
* Conditional static linking for linux/amd64 for DuckDB, SQLite and Snowflake ([#14626](https://github.com/cloudquery/cloudquery/issues/14626)) ([0ba7e3f](https://github.com/cloudquery/cloudquery/commit/0ba7e3f938c11fbc0c9c3f9dc05fce677d311f47))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.9 ([#14627](https://github.com/cloudquery/cloudquery/issues/14627)) ([c1d244c](https://github.com/cloudquery/cloudquery/commit/c1d244c95199141ac39a713a3f0577b2fb3bf736))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.3.0 ([#14635](https://github.com/cloudquery/cloudquery/issues/14635)) ([00b380c](https://github.com/cloudquery/cloudquery/commit/00b380c10be1642f737f871ba5588888ed5dd180))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.4.0 ([#14639](https://github.com/cloudquery/cloudquery/issues/14639)) ([f139c0e](https://github.com/cloudquery/cloudquery/commit/f139c0e9369ef92a3cd874003db40b48e229ab58))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.14.0 ([#14577](https://github.com/cloudquery/cloudquery/issues/14577)) ([223c4c1](https://github.com/cloudquery/cloudquery/commit/223c4c1df6c432d7f1bf67a48114e417282bcd0f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.15.0 ([#14622](https://github.com/cloudquery/cloudquery/issues/14622)) ([b497a6b](https://github.com/cloudquery/cloudquery/commit/b497a6bc5645854bd25d4083fd91ec549a7f274f))
* Remove unnecessary type registration ([#14573](https://github.com/cloudquery/cloudquery/issues/14573)) ([5e3aad2](https://github.com/cloudquery/cloudquery/commit/5e3aad25b7011e706f91018ebceb3f24e1a87f52))

## [5.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.10...plugins-destination-duckdb-v5.0.0) (2023-10-16)


### ⚠ BREAKING CHANGES

* Upgrade to DuckDB v0.9.1. Because this changes the underlying storage layer of DuckDB, it is considered a breaking change. Previous versions of the CloudQuery DuckDB plugin used DuckDB version v0.8.x. The storage of DuckDB is not yet stable; newer versions of DuckDB cannot read old database files and vice versa. The storage will be stabilized when version 1.0 releases. For now, we recommend that you load the database file in a supported version of DuckDB, and use the EXPORT DATABASE command followed by IMPORT DATABASE on the current version of DuckDB. See the storage page for more information: https://duckdb.org/internals/storage

### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to d401686 ([#14459](https://github.com/cloudquery/cloudquery/issues/14459)) ([7ce40f8](https://github.com/cloudquery/cloudquery/commit/7ce40f8dcb1e408c385e877e56b5bb78906b10d2))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to dbcb149 ([#14537](https://github.com/cloudquery/cloudquery/issues/14537)) ([68686f4](https://github.com/cloudquery/cloudquery/commit/68686f4e7636db02bddd961e3d75b60d5218ca85))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.6 ([#14475](https://github.com/cloudquery/cloudquery/issues/14475)) ([83fe7ca](https://github.com/cloudquery/cloudquery/commit/83fe7ca2f5fa83bd3219ddde8fe44fcf1d447480))
* **deps:** Update module github.com/cloudquery/cloudquery-api-go to v1.2.8 ([#14503](https://github.com/cloudquery/cloudquery/issues/14503)) ([4056593](https://github.com/cloudquery/cloudquery/commit/40565937cfc12b33048980b55e91a9a60a62bd47))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.2 ([#14378](https://github.com/cloudquery/cloudquery/issues/14378)) ([a2e0c46](https://github.com/cloudquery/cloudquery/commit/a2e0c4615af4aa205fa082d3f196ea2dc5ce2445))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.3 ([#14436](https://github.com/cloudquery/cloudquery/issues/14436)) ([d529e2d](https://github.com/cloudquery/cloudquery/commit/d529e2d22da93a234492c4165e7eed1257c5767f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.4 ([#14489](https://github.com/cloudquery/cloudquery/issues/14489)) ([9bb45dc](https://github.com/cloudquery/cloudquery/commit/9bb45dc2dacc2c7a6fbd47538b954f731741809b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.12.5 ([#14516](https://github.com/cloudquery/cloudquery/issues/14516)) ([2d905bf](https://github.com/cloudquery/cloudquery/commit/2d905bf9ea81556282c8ca27dcc6334606a2e83b))
* **deps:** Update module golang.org/x/net to v0.17.0 [SECURITY] ([#14500](https://github.com/cloudquery/cloudquery/issues/14500)) ([9e603d5](https://github.com/cloudquery/cloudquery/commit/9e603d50d28033ed5bf451e569abc7c25014dbfb))
* Upgrade to DuckDB v0.9.1. Because this changes the underlying storage layer of DuckDB, it is considered a breaking change. Previous versions of the CloudQuery DuckDB plugin used DuckDB version v0.8.x. The storage of DuckDB is not yet stable; newer versions of DuckDB cannot read old database files and vice versa. The storage will be stabilized when version 1.0 releases. For now, we recommend that you load the database file in a supported version of DuckDB, and use the EXPORT DATABASE command followed by IMPORT DATABASE on the current version of DuckDB. See the storage page for more information: https://duckdb.org/internals/storage ([109aa9b](https://github.com/cloudquery/cloudquery/commit/109aa9b8e6d5d08561c08c62430aea26b4b4ebaa))

## [4.2.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.9...plugins-destination-duckdb-v4.2.10) (2023-10-05)


### Bug Fixes

* Disable Windows publishing ([#14372](https://github.com/cloudquery/cloudquery/issues/14372)) ([1dce3b6](https://github.com/cloudquery/cloudquery/commit/1dce3b6a2b201978c469f4e32b8614d7e106c5af))

## [4.2.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.8...plugins-destination-duckdb-v4.2.9) (2023-10-04)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v14 digest to 00efb06 ([#14202](https://github.com/cloudquery/cloudquery/issues/14202)) ([fc8cc62](https://github.com/cloudquery/cloudquery/commit/fc8cc62ed70db157612e88678c123ba6a34b3b3c))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 7ded38b ([#14246](https://github.com/cloudquery/cloudquery/issues/14246)) ([005891e](https://github.com/cloudquery/cloudquery/commit/005891e1892b41235ddb3b102f4bb6dafd48949a))

## [4.2.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.7...plugins-destination-duckdb-v4.2.8) (2023-09-28)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to ffb7089 ([#13879](https://github.com/cloudquery/cloudquery/issues/13879)) ([f95ced5](https://github.com/cloudquery/cloudquery/commit/f95ced5daa2b123bd71ddff75bd76b3b008790c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.0 ([#13978](https://github.com/cloudquery/cloudquery/issues/13978)) ([2efdf55](https://github.com/cloudquery/cloudquery/commit/2efdf55aed94a14c35c51632ff61ed454caaf5a5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.2 ([#13988](https://github.com/cloudquery/cloudquery/issues/13988)) ([aebaddf](https://github.com/cloudquery/cloudquery/commit/aebaddfc5ca0d7574b8cd72e9e074ec612472dbe))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.0 ([#14031](https://github.com/cloudquery/cloudquery/issues/14031)) ([ac7cdc4](https://github.com/cloudquery/cloudquery/commit/ac7cdc4f7d71599dad89b3170bb7bda676984228))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.1 ([#14063](https://github.com/cloudquery/cloudquery/issues/14063)) ([5a0ff7b](https://github.com/cloudquery/cloudquery/commit/5a0ff7b67890478c371385b379e0a8ef0c2f4865))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.8.0 ([#13950](https://github.com/cloudquery/cloudquery/issues/13950)) ([15b0b69](https://github.com/cloudquery/cloudquery/commit/15b0b6925932613ed2915a3255b3466f21a5c7bf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.9.0 ([#13960](https://github.com/cloudquery/cloudquery/issues/13960)) ([f074076](https://github.com/cloudquery/cloudquery/commit/f074076a21dc0b8cadfdc3adb9731473d24d28b1))

## [4.2.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.6...plugins-destination-duckdb-v4.2.7) (2023-09-12)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 483f6b2 ([#13780](https://github.com/cloudquery/cloudquery/issues/13780)) ([8d31b44](https://github.com/cloudquery/cloudquery/commit/8d31b44f787f42d47f186cdcc4a5739a3a370a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.0 ([#13625](https://github.com/cloudquery/cloudquery/issues/13625)) ([bb5463f](https://github.com/cloudquery/cloudquery/commit/bb5463fb5919f50f1327eecae884b2ab99fb8b34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.1 ([#13713](https://github.com/cloudquery/cloudquery/issues/13713)) ([73004dc](https://github.com/cloudquery/cloudquery/commit/73004dcabd05bf474d8b5960b8c747a894b98560))

## [4.2.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.5...plugins-destination-duckdb-v4.2.6) (2023-09-05)


### Bug Fixes

* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))
* **deps:** Update github.com/apache/arrow/go/v14 digest to a526ba6 ([#13562](https://github.com/cloudquery/cloudquery/issues/13562)) ([248672b](https://github.com/cloudquery/cloudquery/commit/248672beb020828cde1cb608d5c1ed6d656c777b))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to cd3d411 ([#13598](https://github.com/cloudquery/cloudquery/issues/13598)) ([f22bfa6](https://github.com/cloudquery/cloudquery/commit/f22bfa6b2d4fd0caeacf0726ccd307db38f8860c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.0 ([#13492](https://github.com/cloudquery/cloudquery/issues/13492)) ([c305876](https://github.com/cloudquery/cloudquery/commit/c305876e3d92944aa6c1a26547f786fdc5b50e23))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.4 ([#13519](https://github.com/cloudquery/cloudquery/issues/13519)) ([9d25165](https://github.com/cloudquery/cloudquery/commit/9d25165820703844c6de96688d939aa5033608ae))

## [4.2.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.4...plugins-destination-duckdb-v4.2.5) (2023-08-29)


### Bug Fixes

* Properly handle sliced struct arrays ([#13388](https://github.com/cloudquery/cloudquery/issues/13388)) ([43ee769](https://github.com/cloudquery/cloudquery/commit/43ee769954a107b045e50c083f2e357fab383b31))

## [4.2.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.3...plugins-destination-duckdb-v4.2.4) (2023-08-29)


### Bug Fixes

* **deps:** Update `github.com/cloudquery/arrow/go/v13` to `github.com/apache/arrow/go/v14` ([#13341](https://github.com/cloudquery/cloudquery/issues/13341)) ([feb8f87](https://github.com/cloudquery/cloudquery/commit/feb8f87d8d761eb9c49ce84329ad0397f730a918))
* **deps:** Update `github.com/cloudquery/plugin-sdk/v4` to v4.5.5 ([#13280](https://github.com/cloudquery/cloudquery/issues/13280)) ([2d9abfb](https://github.com/cloudquery/cloudquery/commit/2d9abfb42b0840a2f353594b89080ed51aa719ad))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 5b83d4f ([#13203](https://github.com/cloudquery/cloudquery/issues/13203)) ([b0a4b8c](https://github.com/cloudquery/cloudquery/commit/b0a4b8ccf7c429bf5a6ed88866865212015b68e4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.1 ([#13195](https://github.com/cloudquery/cloudquery/issues/13195)) ([a184c37](https://github.com/cloudquery/cloudquery/commit/a184c3786ad49df8564344773e9b96f617ef87a1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.6 ([#13345](https://github.com/cloudquery/cloudquery/issues/13345)) ([a995a05](https://github.com/cloudquery/cloudquery/commit/a995a0598a209e0fe3ba09f4ced2a052dc14b67a))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.4.4 ([#13286](https://github.com/cloudquery/cloudquery/issues/13286)) ([0978184](https://github.com/cloudquery/cloudquery/commit/09781844e264b4c18db7bf3c5f316c8f8eadc296))

## [4.2.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.2...plugins-destination-duckdb-v4.2.3) (2023-08-15)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e9683e1 ([#13015](https://github.com/cloudquery/cloudquery/issues/13015)) ([6557696](https://github.com/cloudquery/cloudquery/commit/65576966d3bd14297499a5b85d3b4fc2c7918df3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.4.0 ([#12850](https://github.com/cloudquery/cloudquery/issues/12850)) ([0861200](https://github.com/cloudquery/cloudquery/commit/086120054b45213947e95be954ba6164b9cf6587))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.0 ([#13068](https://github.com/cloudquery/cloudquery/issues/13068)) ([7bb0e4b](https://github.com/cloudquery/cloudquery/commit/7bb0e4ba654971726e16a6a501393e3831170307))

## [4.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.1...plugins-destination-duckdb-v4.2.2) (2023-08-08)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f53878d ([#12778](https://github.com/cloudquery/cloudquery/issues/12778)) ([6f5d58e](https://github.com/cloudquery/cloudquery/commit/6f5d58e3b84d3c76b1d1a3d6c5a488f77995a057))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.4 ([#12718](https://github.com/cloudquery/cloudquery/issues/12718)) ([f059a15](https://github.com/cloudquery/cloudquery/commit/f059a159a2ee406ab2b0a33792c244cd217025a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.5 ([#12731](https://github.com/cloudquery/cloudquery/issues/12731)) ([d267239](https://github.com/cloudquery/cloudquery/commit/d267239aa3aca5f94bd36a8db1ec0d9f7dc0865f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.6 ([#12799](https://github.com/cloudquery/cloudquery/issues/12799)) ([fb0e0d7](https://github.com/cloudquery/cloudquery/commit/fb0e0d75ab010f421c834e58d93676de76fcb423))

## [4.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.2.0...plugins-destination-duckdb-v4.2.1) (2023-08-01)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 112f949 ([#12659](https://github.com/cloudquery/cloudquery/issues/12659)) ([48d73a9](https://github.com/cloudquery/cloudquery/commit/48d73a93e678994f43171c363f5a75c29547b0b9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 3452eb0 ([#12595](https://github.com/cloudquery/cloudquery/issues/12595)) ([c1c0949](https://github.com/cloudquery/cloudquery/commit/c1c09490b17f2e64435e05d745890cdb8b22310d))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.4.2 ([#12657](https://github.com/cloudquery/cloudquery/issues/12657)) ([feebab9](https://github.com/cloudquery/cloudquery/commit/feebab9b107f9d3e6f2278b9af9e40fec05e575f))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.4.3 ([#12672](https://github.com/cloudquery/cloudquery/issues/12672)) ([877133f](https://github.com/cloudquery/cloudquery/commit/877133f20042ec9a027b13dfb633135602a75770))

## [4.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.1.4...plugins-destination-duckdb-v4.2.0) (2023-07-27)


### Features

* Support date types ([#12538](https://github.com/cloudquery/cloudquery/issues/12538)) ([abb9ec0](https://github.com/cloudquery/cloudquery/commit/abb9ec06c23fe0204707667824f46a1b5435b8f5))

## [4.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.1.3...plugins-destination-duckdb-v4.1.4) (2023-07-25)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))
* **migration:** Make it clear that migration can be done manually and not only via `migrate_mode: forced` ([#12390](https://github.com/cloudquery/cloudquery/issues/12390)) ([33d39cf](https://github.com/cloudquery/cloudquery/commit/33d39cfa87243660241518e23fd5d845ce56d9da))

## [4.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.1.2...plugins-destination-duckdb-v4.1.3) (2023-07-18)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.0 ([#12174](https://github.com/cloudquery/cloudquery/issues/12174)) ([80f0289](https://github.com/cloudquery/cloudquery/commit/80f02892a4cf876c4bf4dd4fd9367afb3770ad26))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.1 ([#12185](https://github.com/cloudquery/cloudquery/issues/12185)) ([cfaff16](https://github.com/cloudquery/cloudquery/commit/cfaff16d89800235b6e3015eeb6957d5783d1393))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.0 ([#12256](https://github.com/cloudquery/cloudquery/issues/12256)) ([eaec331](https://github.com/cloudquery/cloudquery/commit/eaec33165345ad51fdb6ddbffbf8a1199ebd6384))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.1 ([#12272](https://github.com/cloudquery/cloudquery/issues/12272)) ([557ca69](https://github.com/cloudquery/cloudquery/commit/557ca69a7dee9dabb80e6afb6f41f205fd8a80d8))
* **deps:** Upgrade destination plugins to SDK v4.0.0 release ([#12145](https://github.com/cloudquery/cloudquery/issues/12145)) ([09172d3](https://github.com/cloudquery/cloudquery/commit/09172d35baddc68a0267fdb6491e361ed8835285))

## [4.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.1.1...plugins-destination-duckdb-v4.1.2) (2023-07-10)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **duckdb:** Support V0/V1 syncs ([#12106](https://github.com/cloudquery/cloudquery/issues/12106)) ([eb2e1e4](https://github.com/cloudquery/cloudquery/commit/eb2e1e4f16df0c7ffdd34944ecb2d2b10c0f9993)), closes [#12105](https://github.com/cloudquery/cloudquery/issues/12105)
* **postgresql:** Rerun release please ([#12002](https://github.com/cloudquery/cloudquery/issues/12002)) ([9d12843](https://github.com/cloudquery/cloudquery/commit/9d12843462d1019d26bc239f8f928bf5f62940cf))

## [4.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.1.0...plugins-destination-duckdb-v4.1.1) (2023-07-03)


### Bug Fixes

* Close writers ([#11887](https://github.com/cloudquery/cloudquery/issues/11887)) ([26fad6c](https://github.com/cloudquery/cloudquery/commit/26fad6c7cf041abecdd82ebf4d8894e8b1ef13b4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))

## [4.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.0.3...plugins-destination-duckdb-v4.1.0) (2023-07-02)


### Features

* **duckdb:** Update to SDK v4.1.0-rc1 ([#11802](https://github.com/cloudquery/cloudquery/issues/11802)) ([8631dbf](https://github.com/cloudquery/cloudquery/commit/8631dbfbffeb2700806fad037f9b242c380381a2))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))
* **duckdb:** Update SDK to v4.4.0-rc1 ([#11834](https://github.com/cloudquery/cloudquery/issues/11834)) ([33d6a40](https://github.com/cloudquery/cloudquery/commit/33d6a40a1be0691f548a77e7727a02c8e43cdf9b))
* **duckdb:** Update to SDK v4.4.0-rc1 ([#11851](https://github.com/cloudquery/cloudquery/issues/11851)) ([aade378](https://github.com/cloudquery/cloudquery/commit/aade3781c769926b7c23e9ba59c3585a8f60b2fd))
* Update destinations to v4.4.2-rc1 ([#11872](https://github.com/cloudquery/cloudquery/issues/11872)) ([bef90db](https://github.com/cloudquery/cloudquery/commit/bef90db73d0d808ae8013cf5c926e91b63c3cd5f))

## [4.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.0.2...plugins-destination-duckdb-v4.0.3) (2023-06-21)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.6 ([#11473](https://github.com/cloudquery/cloudquery/issues/11473)) ([7272133](https://github.com/cloudquery/cloudquery/commit/72721336632e127dd37de4541f2f503bf4f73fb6))
* Update DuckDB driver to v1.4.1 ([#11692](https://github.com/cloudquery/cloudquery/issues/11692)) ([0ce23c1](https://github.com/cloudquery/cloudquery/commit/0ce23c117760a121be8b43ae04fdf6c9170c21af))
* Write buffered parquet files ([#11546](https://github.com/cloudquery/cloudquery/issues/11546)) ([77699d2](https://github.com/cloudquery/cloudquery/commit/77699d23dd2b2425aa999e3439105e466d74404e))

## [4.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.0.1...plugins-destination-duckdb-v4.0.2) (2023-06-13)


### Bug Fixes

* Don't update `UNIQUE` columns (DuckDB contraint) ([#11516](https://github.com/cloudquery/cloudquery/issues/11516)) ([697a9d9](https://github.com/cloudquery/cloudquery/commit/697a9d9a19fe05e5f93d34e45497fbaf3dab1b00))

## [4.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.0.0...plugins-destination-duckdb-v4.0.1) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v3.0.0...plugins-destination-duckdb-v4.0.0) (2023-06-06)


### ⚠ BREAKING CHANGES

* **types:** Support writing Apache Arrow nested types:
    * Structs as DuckDB structs
    * Maps as DuckDB maps

### Features

* **types:** Proper support for nested types ([#11196](https://github.com/cloudquery/cloudquery/issues/11196)) ([7c6a3e2](https://github.com/cloudquery/cloudquery/commit/7c6a3e2d24632ffd0ac3eacd6bd65b89394ebe0f))


### Bug Fixes

* **deps:** Update `github.com/cloudquery/plugin-sdk/v3` to `v3.10.0` ([#11116](https://github.com/cloudquery/cloudquery/issues/11116)) ([bba7c4e](https://github.com/cloudquery/cloudquery/commit/bba7c4ef9368741ed00e5c04bdc2bc9a1de9a521))
* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11150](https://github.com/cloudquery/cloudquery/issues/11150)) ([dc00994](https://github.com/cloudquery/cloudquery/commit/dc00994e32936af7e9893c93561d0f9df225a929))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **race:** Fix issue caused by [Over-Eager Unique Constraint Checking](https://duckdb.org/docs/sql/indexes#over-eager-unique-constraint-checking) in DuckDB overwrite ([#11215](https://github.com/cloudquery/cloudquery/issues/11215)) ([c0b9f0a](https://github.com/cloudquery/cloudquery/commit/c0b9f0af37b70b7a386ae4b1c20e52794db87681))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v2.0.1...plugins-destination-duckdb-v3.0.0) (2023-05-24)


### ⚠ BREAKING CHANGES

* **duckdb:** Move DuckDB write to Parquet ([#10874](https://github.com/cloudquery/cloudquery/issues/10874))

### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* **duckdb:** Migrate to SDK V3 ([#10874](https://github.com/cloudquery/cloudquery/issues/10874)) ([84e6631](https://github.com/cloudquery/cloudquery/commit/84e663193b5cecdeb56f9a5debcd4ff59e1c49bb))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v2.0.0...plugins-destination-duckdb-v2.0.1) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.6...plugins-destination-duckdb-v2.0.0) (2023-04-26)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([e38eae6](https://github.com/cloudquery/cloudquery/commit/e38eae6bffbdd34f5959ff3cd7124b789ed2dd26))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.5...plugins-destination-duckdb-v1.0.6) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.4...plugins-destination-duckdb-v1.0.5) (2023-04-09)


### Bug Fixes

* **duckdb:** Fix multiple pks in delete from ([#9775](https://github.com/cloudquery/cloudquery/issues/9775)) ([706217f](https://github.com/cloudquery/cloudquery/commit/706217fef50703ee228f4df20d95eb55d934fb86))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.3...plugins-destination-duckdb-v1.0.4) (2023-04-09)


### Bug Fixes

* **duckdb:** Support multiple PKs ([#9772](https://github.com/cloudquery/cloudquery/issues/9772)) ([d94d87d](https://github.com/cloudquery/cloudquery/commit/d94d87da16ab6df52dd3705c5f6c60cd151a26fc))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.2...plugins-destination-duckdb-v1.0.3) (2023-04-04)


### Bug Fixes

* **deps:** Update ghcr.io/gythialy/golang-cross Docker tag to v1.20.2 ([#9599](https://github.com/cloudquery/cloudquery/issues/9599)) ([46ce2dc](https://github.com/cloudquery/cloudquery/commit/46ce2dc3165e87f4017608c883c7e4ede3d8b19d))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.2.2 ([#9608](https://github.com/cloudquery/cloudquery/issues/9608)) ([4dfbc9e](https://github.com/cloudquery/cloudquery/commit/4dfbc9e5e9b892c053b878fa60568045459a17d1))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.1...plugins-destination-duckdb-v1.0.2) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.0...plugins-destination-duckdb-v1.0.1) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## 1.0.0 (2023-03-09)


### Features

* DuckDB destination ([#8758](https://github.com/cloudquery/cloudquery/issues/8758)) ([2ed9c37](https://github.com/cloudquery/cloudquery/commit/2ed9c37708d7df595ce633d3b13099c6074086c6))
