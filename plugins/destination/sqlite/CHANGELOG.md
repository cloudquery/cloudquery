# Changelog

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.6...plugins-destination-sqlite-v1.2.0) (2023-01-24)


### Features

* **sqlite-migrate:** Support PK changes to schema ([#7006](https://github.com/cloudquery/cloudquery/issues/7006)) ([dddd852](https://github.com/cloudquery/cloudquery/commit/dddd85282f0e008d3815a9cf1bd221d696b49291))
* **sqlite:** Collect and report migration errors before starting the migration ([#6759](https://github.com/cloudquery/cloudquery/issues/6759)) ([a80e9d9](https://github.com/cloudquery/cloudquery/commit/a80e9d9fe23b9cf98d70fe1366f6d76f9f540f3e))
* **sqlite:** Support force migration ([#6763](https://github.com/cloudquery/cloudquery/issues/6763)) ([19bba77](https://github.com/cloudquery/cloudquery/commit/19bba778420e215deca02856ebe155b6be985219))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))

## [1.1.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.5...plugins-destination-sqlite-v1.1.6) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))
* **sqlite-migrate:** Use `int` when getting `pk` column information ([#6848](https://github.com/cloudquery/cloudquery/issues/6848)) ([8ed1d94](https://github.com/cloudquery/cloudquery/commit/8ed1d946061b58eb80ccb65201fa26266671f11d))
* **sqlite:** Set module in logs to `sqlite-dest` instead of `pg-dest` ([#6764](https://github.com/cloudquery/cloudquery/issues/6764)) ([6cfda91](https://github.com/cloudquery/cloudquery/commit/6cfda91c6d1d62688ba0d56ea8119136842b334a))

## [1.1.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.4...plugins-destination-sqlite-v1.1.5) (2023-01-10)


### Bug Fixes

* Correct error message in migration ([#6612](https://github.com/cloudquery/cloudquery/issues/6612)) ([434fe4d](https://github.com/cloudquery/cloudquery/commit/434fe4d15746277e903edd3ce5635fe2323ed413))

## [1.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.3...plugins-destination-sqlite-v1.1.4) (2023-01-10)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([#6553](https://github.com/cloudquery/cloudquery/issues/6553)) ([392b848](https://github.com/cloudquery/cloudquery/commit/392b848b3124f9cf28f6234fdb9a43d671069879))
* **deps:** Update plugin-sdk to v1.21.0 for destinations ([#6419](https://github.com/cloudquery/cloudquery/issues/6419)) ([f3b989f](https://github.com/cloudquery/cloudquery/commit/f3b989f7cbe335481dc01ad2a56cf7eff48e01d5))
* Fix SQLite migration logic ([#6372](https://github.com/cloudquery/cloudquery/issues/6372)) ([5f8a197](https://github.com/cloudquery/cloudquery/commit/5f8a1973482a3e9cfcd0b7cb4720f727731ba15e))

## [1.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.2...plugins-destination-sqlite-v1.1.3) (2023-01-03)


### Bug Fixes

* **deps:** Update ghcr.io/gythialy/golang-cross Docker tag to v1.19.4 ([#6174](https://github.com/cloudquery/cloudquery/issues/6174)) ([8fe3f41](https://github.com/cloudquery/cloudquery/commit/8fe3f416bb59e205f8bb4fa9367428bb9e7ed2cb))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))

## [1.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.1...plugins-destination-sqlite-v1.1.2) (2022-12-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.1.0...plugins-destination-sqlite-v1.1.1) (2022-12-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.14...plugins-destination-sqlite-v1.1.0) (2022-12-23)


### Features

* **destinations:** Migrate to managed batching SDK ([#5805](https://github.com/cloudquery/cloudquery/issues/5805)) ([2f130c1](https://github.com/cloudquery/cloudquery/commit/2f130c12c6e83ccd8a2d036ab5c47b55e2fb5280))

## [1.0.14](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.13...plugins-destination-sqlite-v1.0.14) (2022-12-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.2 ([#5583](https://github.com/cloudquery/cloudquery/issues/5583)) ([d721c4e](https://github.com/cloudquery/cloudquery/commit/d721c4e06b8a97b5373215aca0e4ed64942ac489))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.3 ([#5639](https://github.com/cloudquery/cloudquery/issues/5639)) ([6452d0e](https://github.com/cloudquery/cloudquery/commit/6452d0ed5a44abad9d7530af6e79cde6504d0c4c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.4 ([#5649](https://github.com/cloudquery/cloudquery/issues/5649)) ([b4aa889](https://github.com/cloudquery/cloudquery/commit/b4aa889e396db3b0887d1684e4bc07da6050af43))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.5 ([#5661](https://github.com/cloudquery/cloudquery/issues/5661)) ([b354b8a](https://github.com/cloudquery/cloudquery/commit/b354b8a3683fa2bc918c1002afac487427d65a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.6 ([#5790](https://github.com/cloudquery/cloudquery/issues/5790)) ([8e2663c](https://github.com/cloudquery/cloudquery/commit/8e2663c17c3347afd5e53f665462adc3e709c96c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.7 ([#5797](https://github.com/cloudquery/cloudquery/issues/5797)) ([15da529](https://github.com/cloudquery/cloudquery/commit/15da5294786fa2656228ca5bbc48ef1fc44e486b))

## [1.0.13](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.12...plugins-destination-sqlite-v1.0.13) (2022-12-13)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.1 ([#5458](https://github.com/cloudquery/cloudquery/issues/5458)) ([58b7432](https://github.com/cloudquery/cloudquery/commit/58b74321cd253c9a843c8c103f324abb93952195))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.2 ([#5497](https://github.com/cloudquery/cloudquery/issues/5497)) ([c1876cf](https://github.com/cloudquery/cloudquery/commit/c1876cf793b43d825a25fb3c9ba4996e4b09964f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.0 ([#5539](https://github.com/cloudquery/cloudquery/issues/5539)) ([fb71293](https://github.com/cloudquery/cloudquery/commit/fb71293d5cfe1b2ef32ba83d604ac3c48e662bce))

## [1.0.12](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.11...plugins-destination-sqlite-v1.0.12) (2022-12-06)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.0 ([#5416](https://github.com/cloudquery/cloudquery/issues/5416)) ([2e7ca35](https://github.com/cloudquery/cloudquery/commit/2e7ca35922fdb14fd717f582aaaa9693dae2ef4c))

## [1.0.11](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.10...plugins-destination-sqlite-v1.0.11) (2022-12-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.10.0 ([#5153](https://github.com/cloudquery/cloudquery/issues/5153)) ([ea1f77e](https://github.com/cloudquery/cloudquery/commit/ea1f77e910f430287600e74cedd7d3f4ae79eb18))
* Handling of NULL bytes in Postgresql Text fields ([#5249](https://github.com/cloudquery/cloudquery/issues/5249)) ([936c311](https://github.com/cloudquery/cloudquery/commit/936c311e6cd5cc76e2c10b2f991e85de6e1fadb4))

## [1.0.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.9...plugins-destination-sqlite-v1.0.10) (2022-11-29)


### Bug Fixes

* DeleteStale from relational tables as well ([#5143](https://github.com/cloudquery/cloudquery/issues/5143)) ([23aa159](https://github.com/cloudquery/cloudquery/commit/23aa1597a2db275df85c15a52bcd2986f19d9621))
* **deps:** Update plugin-sdk for sqlite to v1.8.1 ([#5044](https://github.com/cloudquery/cloudquery/issues/5044)) ([7c085c6](https://github.com/cloudquery/cloudquery/commit/7c085c616e82967be9aecbd8d90ca3400dc89112))
* **deps:** Update plugin-sdk for sqlite to v1.9.0 ([#5087](https://github.com/cloudquery/cloudquery/issues/5087)) ([9f43e46](https://github.com/cloudquery/cloudquery/commit/9f43e46112ca00b06805c4791e7739aba41af0a6))

## [1.0.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.8...plugins-destination-sqlite-v1.0.9) (2022-11-23)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.7.0 ([#4915](https://github.com/cloudquery/cloudquery/issues/4915)) ([6e5268f](https://github.com/cloudquery/cloudquery/commit/6e5268f268ad836f2780c685087b7e362c58591f))
* **deps:** Update plugin-sdk for sqlite to v1.8.0 ([#4978](https://github.com/cloudquery/cloudquery/issues/4978)) ([e183336](https://github.com/cloudquery/cloudquery/commit/e183336d8b062d4a492cdc4d4a8dba634aaef04a))

## [1.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.7...plugins-destination-sqlite-v1.0.8) (2022-11-21)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.6.0 ([#4853](https://github.com/cloudquery/cloudquery/issues/4853)) ([460c6ff](https://github.com/cloudquery/cloudquery/commit/460c6ffddab8aba843d150cafb82825c46597b41))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.6...plugins-destination-sqlite-v1.0.7) (2022-11-16)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.5.3 ([#4652](https://github.com/cloudquery/cloudquery/issues/4652)) ([28faf58](https://github.com/cloudquery/cloudquery/commit/28faf5819e612e30b09e6c390561e835cce3eff0))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.5...plugins-destination-sqlite-v1.0.6) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.5.2 ([#4567](https://github.com/cloudquery/cloudquery/issues/4567)) ([d4b3b8c](https://github.com/cloudquery/cloudquery/commit/d4b3b8cfd9beafa906d88bf97cc246e1a7124deb))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.4...plugins-destination-sqlite-v1.0.5) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.5.1 ([#4507](https://github.com/cloudquery/cloudquery/issues/4507)) ([de43ab2](https://github.com/cloudquery/cloudquery/commit/de43ab295bcb427cd0a6e1cf0fedd39cebf1ae88))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.3...plugins-destination-sqlite-v1.0.4) (2022-11-11)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.5.0 ([#4397](https://github.com/cloudquery/cloudquery/issues/4397)) ([01a34fe](https://github.com/cloudquery/cloudquery/commit/01a34fe547d3380bae7e18110d824af60506af7e))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.2...plugins-destination-sqlite-v1.0.3) (2022-11-11)


### Bug Fixes

* Quote tables and columns for SQLite ([#4342](https://github.com/cloudquery/cloudquery/issues/4342)) ([12bb673](https://github.com/cloudquery/cloudquery/commit/12bb673e32dfeae1d3cc4f150d30f350c5798b1d))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.1...plugins-destination-sqlite-v1.0.2) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.4.1 ([#4301](https://github.com/cloudquery/cloudquery/issues/4301)) ([5b36929](https://github.com/cloudquery/cloudquery/commit/5b36929f18f56fc58e0bbc4d41b5fc08d4f7b0e5))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-sqlite-v1.0.0...plugins-destination-sqlite-v1.0.1) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for sqlite to v1.4.0 ([#4278](https://github.com/cloudquery/cloudquery/issues/4278)) ([ddf46dd](https://github.com/cloudquery/cloudquery/commit/ddf46dd9c2418c3f5e07ec0f808c1ab5960684d4))

## 1.0.0 (2022-11-10)


### Features

* Add SQLite destination plugin ([#3471](https://github.com/cloudquery/cloudquery/issues/3471)) ([8306f72](https://github.com/cloudquery/cloudquery/commit/8306f7235539262199bfb52bd0d4b44301363b3b))


### Bug Fixes

* **deps:** Update plugin-sdk for csv to v1.1.0 ([#3918](https://github.com/cloudquery/cloudquery/issues/3918)) ([f1acd68](https://github.com/cloudquery/cloudquery/commit/f1acd688fcd90011cc9be1be2285e3fe9369e341))
