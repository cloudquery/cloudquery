# Changelog

## [2.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.1.2...plugins-destination-postgresql-v2.1.3) (2023-02-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))
* **migrate:** Handle timestamps not set by CloudQuery ([#8416](https://github.com/cloudquery/cloudquery/issues/8416)) ([7831a28](https://github.com/cloudquery/cloudquery/commit/7831a280a188f4b0e44b2ba35f2b7075d525c6be))
* **migrate:** Use unique column option instead of hard coding cq_id ([#8370](https://github.com/cloudquery/cloudquery/issues/8370)) ([e692a06](https://github.com/cloudquery/cloudquery/commit/e692a063fa5d8d81230d904bb303cffb3fbe496f))

## [2.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.1.1...plugins-destination-postgresql-v2.1.2) (2023-02-21)


### Bug Fixes

* **migrate:** Make `_cq_id` unique ([#8326](https://github.com/cloudquery/cloudquery/issues/8326)) ([d2cf7ee](https://github.com/cloudquery/cloudquery/commit/d2cf7ee9c5dbdac52b465070b6f051d18f944b5c))

## [2.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.1.0...plugins-destination-postgresql-v2.1.1) (2023-02-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))
* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.10...plugins-destination-postgresql-v2.1.0) (2023-02-16)


### Features

* **pg:** Faster migrations ([#7819](https://github.com/cloudquery/cloudquery/issues/7819)) ([8f51733](https://github.com/cloudquery/cloudquery/commit/8f517337a48c25aa7471b7d1e9381a3188dc3c3b))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* Postgresql timestamps ([#7840](https://github.com/cloudquery/cloudquery/issues/7840)) ([e2c8b61](https://github.com/cloudquery/cloudquery/commit/e2c8b613696447602d052d0686b643a38694573c))

## [2.0.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.9...plugins-destination-postgresql-v2.0.10) (2023-02-08)


### Bug Fixes

* **postgresql:** Sanitize JSONs ([#7801](https://github.com/cloudquery/cloudquery/issues/7801)) ([8705c9a](https://github.com/cloudquery/cloudquery/commit/8705c9a658210294ede67e56d83612de4f14b3b1))

## [2.0.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.8...plugins-destination-postgresql-v2.0.9) (2023-02-07)


### Bug Fixes

* **deps:** Update github.com/jackc/pgx-zerolog digest to 7c83b3e ([#7530](https://github.com/cloudquery/cloudquery/issues/7530)) ([4d2617e](https://github.com/cloudquery/cloudquery/commit/4d2617ec722c28370fa2af77595be427c0350822))
* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))
* Sort PKs before reporting stale PKs error ([#7620](https://github.com/cloudquery/cloudquery/issues/7620)) ([2d9e196](https://github.com/cloudquery/cloudquery/commit/2d9e19686c1c6aedb90ef45e8058795131165e8e))

## [2.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.7...plugins-destination-postgresql-v2.0.8) (2023-01-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [2.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.6...plugins-destination-postgresql-v2.0.7) (2023-01-24)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))

## [2.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.5...plugins-destination-postgresql-v2.0.6) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))
* Fix number of writes reported by postgresql ([#6757](https://github.com/cloudquery/cloudquery/issues/6757)) ([018edf4](https://github.com/cloudquery/cloudquery/commit/018edf44322f0ccb25fbf97ad4b8fcf0973a944e))

## [2.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.4...plugins-destination-postgresql-v2.0.5) (2023-01-10)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([#6553](https://github.com/cloudquery/cloudquery/issues/6553)) ([392b848](https://github.com/cloudquery/cloudquery/commit/392b848b3124f9cf28f6234fdb9a43d671069879))
* Error if after the migration there are `not null` columns that are not part of the new schema ([#6519](https://github.com/cloudquery/cloudquery/issues/6519)) ([566da52](https://github.com/cloudquery/cloudquery/commit/566da52d730f8ce0b88277f71d88fe0d861a336e))

## [2.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.3...plugins-destination-postgresql-v2.0.4) (2023-01-09)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **pg:** Dont overwrite _cq_id, _cq_parent_id on conflict ([#6521](https://github.com/cloudquery/cloudquery/issues/6521)) ([d1398a5](https://github.com/cloudquery/cloudquery/commit/d1398a5a081142db57291704c7f1df8479668b9c))

## [2.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.2...plugins-destination-postgresql-v2.0.3) (2023-01-07)


### Bug Fixes

* **destinations:** Handle nulls in JSONs ([#6466](https://github.com/cloudquery/cloudquery/issues/6466)) ([f434f00](https://github.com/cloudquery/cloudquery/commit/f434f00285ed8fc5edbacf03194fb983d4d98f86))

## [2.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.1...plugins-destination-postgresql-v2.0.2) (2023-01-06)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update plugin-sdk to v1.21.0 for destinations ([#6419](https://github.com/cloudquery/cloudquery/issues/6419)) ([f3b989f](https://github.com/cloudquery/cloudquery/commit/f3b989f7cbe335481dc01ad2a56cf7eff48e01d5))
* Error if after the migration there are `not null` columns that are not part of the new schema ([#6282](https://github.com/cloudquery/cloudquery/issues/6282)) ([c5a4bf5](https://github.com/cloudquery/cloudquery/commit/c5a4bf596cc145295ae74add917c1d29c05cb493))
* **pg:** Return more detailed pg errors ([#6421](https://github.com/cloudquery/cloudquery/issues/6421)) ([acb3e21](https://github.com/cloudquery/cloudquery/commit/acb3e216b73405763e1e3980f15d43e7166c3f5b))
* **postgresql:** Revert [#6282](https://github.com/cloudquery/cloudquery/issues/6282) ([#6434](https://github.com/cloudquery/cloudquery/issues/6434)) ([2cecacd](https://github.com/cloudquery/cloudquery/commit/2cecacde39edb9dd16498d816415b4a1fb8efb4e))
* **tests:** Postgresql: Read only columns defined in the schema ([#6371](https://github.com/cloudquery/cloudquery/issues/6371)) ([491941a](https://github.com/cloudquery/cloudquery/commit/491941a0c165ee18ce96d71c2e6aecf3d7830d50))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v2.0.0...plugins-destination-postgresql-v2.0.1) (2023-01-03)


### Bug Fixes

* **deps:** Update github.com/jackc/pgservicefile digest to 091c0ba ([#6192](https://github.com/cloudquery/cloudquery/issues/6192)) ([d73d03c](https://github.com/cloudquery/cloudquery/commit/d73d03c10ba1ca0414e7f6f8c82baac0e41ee4f7))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.10.2...plugins-destination-postgresql-v2.0.0) (2022-12-29)


### ⚠ BREAKING CHANGES

* **postgres-spec:** Move `batch_size` from the plugin spec to the top level spec ([#6091](https://github.com/cloudquery/cloudquery/issues/6091))

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))
* **postgres-spec:** Move `batch_size` from the plugin spec to the top level spec ([#6091](https://github.com/cloudquery/cloudquery/issues/6091)) ([c504423](https://github.com/cloudquery/cloudquery/commit/c50442397e3e0ded68940f0d3121d00eae22d912))

## [1.10.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.10.1...plugins-destination-postgresql-v1.10.2) (2022-12-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))

## [1.10.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.10.0...plugins-destination-postgresql-v1.10.1) (2022-12-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))

## [1.10.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.9.0...plugins-destination-postgresql-v1.10.0) (2022-12-23)


### Features

* **destinations:** Migrate to managed batching SDK ([#5805](https://github.com/cloudquery/cloudquery/issues/5805)) ([2f130c1](https://github.com/cloudquery/cloudquery/commit/2f130c12c6e83ccd8a2d036ab5c47b55e2fb5280))

## [1.9.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.8.0...plugins-destination-postgresql-v1.9.0) (2022-12-20)


### Features

* **postgresql:** Update pgx to v5 ([#5757](https://github.com/cloudquery/cloudquery/issues/5757)) ([ce2aaf5](https://github.com/cloudquery/cloudquery/commit/ce2aaf535865d7a44173231732a7bee87ce2ab8c))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.2 ([#5583](https://github.com/cloudquery/cloudquery/issues/5583)) ([d721c4e](https://github.com/cloudquery/cloudquery/commit/d721c4e06b8a97b5373215aca0e4ed64942ac489))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.3 ([#5639](https://github.com/cloudquery/cloudquery/issues/5639)) ([6452d0e](https://github.com/cloudquery/cloudquery/commit/6452d0ed5a44abad9d7530af6e79cde6504d0c4c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.4 ([#5649](https://github.com/cloudquery/cloudquery/issues/5649)) ([b4aa889](https://github.com/cloudquery/cloudquery/commit/b4aa889e396db3b0887d1684e4bc07da6050af43))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.5 ([#5661](https://github.com/cloudquery/cloudquery/issues/5661)) ([b354b8a](https://github.com/cloudquery/cloudquery/commit/b354b8a3683fa2bc918c1002afac487427d65a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.6 ([#5790](https://github.com/cloudquery/cloudquery/issues/5790)) ([8e2663c](https://github.com/cloudquery/cloudquery/commit/8e2663c17c3347afd5e53f665462adc3e709c96c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.7 ([#5797](https://github.com/cloudquery/cloudquery/issues/5797)) ([15da529](https://github.com/cloudquery/cloudquery/commit/15da5294786fa2656228ca5bbc48ef1fc44e486b))

## [1.8.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.17...plugins-destination-postgresql-v1.8.0) (2022-12-13)


### Features

* **postgres:** Log table name with pg errors ([#5552](https://github.com/cloudquery/cloudquery/issues/5552)) ([ee90823](https://github.com/cloudquery/cloudquery/commit/ee908236d2d24cd31b51c59f8d16e62e2fb64c61))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.1 ([#5458](https://github.com/cloudquery/cloudquery/issues/5458)) ([58b7432](https://github.com/cloudquery/cloudquery/commit/58b74321cd253c9a843c8c103f324abb93952195))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.2 ([#5497](https://github.com/cloudquery/cloudquery/issues/5497)) ([c1876cf](https://github.com/cloudquery/cloudquery/commit/c1876cf793b43d825a25fb3c9ba4996e4b09964f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.0 ([#5539](https://github.com/cloudquery/cloudquery/issues/5539)) ([fb71293](https://github.com/cloudquery/cloudquery/commit/fb71293d5cfe1b2ef32ba83d604ac3c48e662bce))

## [1.7.17](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.16...plugins-destination-postgresql-v1.7.17) (2022-12-06)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.0 ([#5416](https://github.com/cloudquery/cloudquery/issues/5416)) ([2e7ca35](https://github.com/cloudquery/cloudquery/commit/2e7ca35922fdb14fd717f582aaaa9693dae2ef4c))

## [1.7.16](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.15...plugins-destination-postgresql-v1.7.16) (2022-12-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.10.0 ([#5153](https://github.com/cloudquery/cloudquery/issues/5153)) ([ea1f77e](https://github.com/cloudquery/cloudquery/commit/ea1f77e910f430287600e74cedd7d3f4ae79eb18))
* Handling of NULL bytes in Postgresql Text fields ([#5249](https://github.com/cloudquery/cloudquery/issues/5249)) ([936c311](https://github.com/cloudquery/cloudquery/commit/936c311e6cd5cc76e2c10b2f991e85de6e1fadb4))

## [1.7.15](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.14...plugins-destination-postgresql-v1.7.15) (2022-11-29)


### Bug Fixes

* DeleteStale from relational tables as well ([#5143](https://github.com/cloudquery/cloudquery/issues/5143)) ([23aa159](https://github.com/cloudquery/cloudquery/commit/23aa1597a2db275df85c15a52bcd2986f19d9621))
* **deps:** Update plugin-sdk for postgresql to v1.8.1 ([#5043](https://github.com/cloudquery/cloudquery/issues/5043)) ([abd4d12](https://github.com/cloudquery/cloudquery/commit/abd4d1241a1b9353ecd43d9dabb9847943f392e3))
* **deps:** Update plugin-sdk for postgresql to v1.9.0 ([#5085](https://github.com/cloudquery/cloudquery/issues/5085)) ([be36fb0](https://github.com/cloudquery/cloudquery/commit/be36fb049c02ff973c31ee61fc7055a5d5a61549))

## [1.7.14](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.13...plugins-destination-postgresql-v1.7.14) (2022-11-23)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.8.0 ([#4977](https://github.com/cloudquery/cloudquery/issues/4977)) ([54583e9](https://github.com/cloudquery/cloudquery/commit/54583e9a5d751f969e00d5158cd1cc0ff422ba59))

## [1.7.13](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.12...plugins-destination-postgresql-v1.7.13) (2022-11-22)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.7.0 ([#4914](https://github.com/cloudquery/cloudquery/issues/4914)) ([c1383fe](https://github.com/cloudquery/cloudquery/commit/c1383fe9dfeece75f1b2d44ec6c31c94dc9601ad))

## [1.7.12](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.11...plugins-destination-postgresql-v1.7.12) (2022-11-21)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.6.0 ([#4852](https://github.com/cloudquery/cloudquery/issues/4852)) ([6c18d6a](https://github.com/cloudquery/cloudquery/commit/6c18d6a5b6984f2376619e9ebeb143585879d688))

## [1.7.11](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.10...plugins-destination-postgresql-v1.7.11) (2022-11-15)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.5.3 ([#4651](https://github.com/cloudquery/cloudquery/issues/4651)) ([6f8ea18](https://github.com/cloudquery/cloudquery/commit/6f8ea18365385831ad4da0414560c1cb8263255f))

## [1.7.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.9...plugins-destination-postgresql-v1.7.10) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.5.2 ([#4566](https://github.com/cloudquery/cloudquery/issues/4566)) ([e05fe42](https://github.com/cloudquery/cloudquery/commit/e05fe42a5f9f4058a973766e1056e57591dde2ce))

## [1.7.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.8...plugins-destination-postgresql-v1.7.9) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.5.1 ([#4511](https://github.com/cloudquery/cloudquery/issues/4511)) ([cb20fd2](https://github.com/cloudquery/cloudquery/commit/cb20fd222bac1a9ea6a6a30a3b86be7989898998))
* Revert "fix(deps): Update plugin-sdk for postgresql to v1.5.0 ([#4396](https://github.com/cloudquery/cloudquery/issues/4396)) ([#4493](https://github.com/cloudquery/cloudquery/issues/4493)) ([0d1eb14](https://github.com/cloudquery/cloudquery/commit/0d1eb14a13504efb37b4d10c539d1f334a220e85))

## [1.7.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.7...plugins-destination-postgresql-v1.7.8) (2022-11-11)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.5.0 ([#4396](https://github.com/cloudquery/cloudquery/issues/4396)) ([f4f73a4](https://github.com/cloudquery/cloudquery/commit/f4f73a4959af7d437038057c057045c3dee83c13))

## [1.7.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.6...plugins-destination-postgresql-v1.7.7) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.4.1 ([#4300](https://github.com/cloudquery/cloudquery/issues/4300)) ([c97f8fb](https://github.com/cloudquery/cloudquery/commit/c97f8fbabe1bc0c8b0c02fa7906085c7933c268e))

## [1.7.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.5...plugins-destination-postgresql-v1.7.6) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.4.0 ([#4237](https://github.com/cloudquery/cloudquery/issues/4237)) ([bdf5747](https://github.com/cloudquery/cloudquery/commit/bdf5747f6026b8b6b5d8de0b2bfd9ec38ca810e9))

## [1.7.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.4...plugins-destination-postgresql-v1.7.5) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.3.2 ([#4204](https://github.com/cloudquery/cloudquery/issues/4204)) ([9a7cf8c](https://github.com/cloudquery/cloudquery/commit/9a7cf8c1fb8bd7f8b300fe5d4c25fe444cbfacfa))

## [1.7.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.3...plugins-destination-postgresql-v1.7.4) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.3.1 ([#4154](https://github.com/cloudquery/cloudquery/issues/4154)) ([422672c](https://github.com/cloudquery/cloudquery/commit/422672c8c9618f7bef4196db08f36051bca6fcfc))

## [1.7.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.2...plugins-destination-postgresql-v1.7.3) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.2.0 ([#4047](https://github.com/cloudquery/cloudquery/issues/4047)) ([9e79b17](https://github.com/cloudquery/cloudquery/commit/9e79b176f155469296bc92fce7760d4260971dbf))
* **deps:** Update plugin-sdk for postgresql to v1.3.0 ([#4078](https://github.com/cloudquery/cloudquery/issues/4078)) ([7383e78](https://github.com/cloudquery/cloudquery/commit/7383e78724fbb92bc0b35102fc985c7f0d3e6e29))

## [1.7.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.1...plugins-destination-postgresql-v1.7.2) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.1.2 ([#4016](https://github.com/cloudquery/cloudquery/issues/4016)) ([2762eab](https://github.com/cloudquery/cloudquery/commit/2762eab9fbe2b545723919603b3d4bba59ec1366))

## [1.7.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.7.0...plugins-destination-postgresql-v1.7.1) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for csv to v1.1.0 ([#3918](https://github.com/cloudquery/cloudquery/issues/3918)) ([f1acd68](https://github.com/cloudquery/cloudquery/commit/f1acd688fcd90011cc9be1be2285e3fe9369e341))
* **deps:** Update plugin-sdk for postgresql to v1.1.1 ([#4001](https://github.com/cloudquery/cloudquery/issues/4001)) ([3a11b2b](https://github.com/cloudquery/cloudquery/commit/3a11b2b3e276e7c11ade65c5e1250dccaf324e0d))

## [1.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.7...plugins-destination-postgresql-v1.7.0) (2022-11-08)


### Features

* Update destination to more managed SDK ([#3842](https://github.com/cloudquery/cloudquery/issues/3842)) ([5e4fd5d](https://github.com/cloudquery/cloudquery/commit/5e4fd5dd4fe3e778e5e98719d16f9e5db85d37bc))

## [1.6.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.6...plugins-destination-postgresql-v1.6.7) (2022-11-08)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1.0.3 ([#3856](https://github.com/cloudquery/cloudquery/issues/3856)) ([df55de2](https://github.com/cloudquery/cloudquery/commit/df55de2185e383fcd47f9c0d276757a465e29963))
* **deps:** Update plugin-sdk for postgresql to v1.0.4 ([#3885](https://github.com/cloudquery/cloudquery/issues/3885)) ([7bcfe9f](https://github.com/cloudquery/cloudquery/commit/7bcfe9fd66c683b254b26665f788126636686dd6))

## [1.6.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.5...plugins-destination-postgresql-v1.6.6) (2022-11-07)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v1 ([#3796](https://github.com/cloudquery/cloudquery/issues/3796)) ([edc90f2](https://github.com/cloudquery/cloudquery/commit/edc90f24474009cd26668adaf14f554f275e9550))

## [1.6.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.4...plugins-destination-postgresql-v1.6.5) (2022-11-07)


### Bug Fixes

* Update SDK to v0.13.23 ([#3727](https://github.com/cloudquery/cloudquery/issues/3727)) ([f9769e5](https://github.com/cloudquery/cloudquery/commit/f9769e5240ea563d63f551407e802096d8ffd032))

## [1.6.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.3...plugins-destination-postgresql-v1.6.4) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.22 ([#3697](https://github.com/cloudquery/cloudquery/issues/3697)) ([541ea61](https://github.com/cloudquery/cloudquery/commit/541ea6124b9872741de67a3b6e63299808a3fafb))

## [1.6.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.2...plugins-destination-postgresql-v1.6.3) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.21 ([#3644](https://github.com/cloudquery/cloudquery/issues/3644)) ([480c150](https://github.com/cloudquery/cloudquery/commit/480c1503fb21f899b9d3117d0589b2fe334bdb8f))

## [1.6.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.1...plugins-destination-postgresql-v1.6.2) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.20 ([#3588](https://github.com/cloudquery/cloudquery/issues/3588)) ([7b106f9](https://github.com/cloudquery/cloudquery/commit/7b106f9849b5e4aafdbfc5b7930c79caf8203c92))

## [1.6.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.6.0...plugins-destination-postgresql-v1.6.1) (2022-11-03)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.19 ([#3525](https://github.com/cloudquery/cloudquery/issues/3525)) ([c0c705c](https://github.com/cloudquery/cloudquery/commit/c0c705c05e4f659f9f9f8ab3eb1ea11dc7719a04))

## [1.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.5.0...plugins-destination-postgresql-v1.6.0) (2022-11-02)


### Features

* Add unique constraint on _cq_id column ([#3449](https://github.com/cloudquery/cloudquery/issues/3449)) ([30fd8a3](https://github.com/cloudquery/cloudquery/commit/30fd8a370c73fff1effcb8474e49dac5cabf4653))

## [1.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.4.0...plugins-destination-postgresql-v1.5.0) (2022-11-01)


### Features

* Migrate cli, plugins and destinations to new type system ([#3323](https://github.com/cloudquery/cloudquery/issues/3323)) ([f265a94](https://github.com/cloudquery/cloudquery/commit/f265a94448ad55c968b26ba8a19681bc81086c11))


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.18 ([#3417](https://github.com/cloudquery/cloudquery/issues/3417)) ([5c1cfcd](https://github.com/cloudquery/cloudquery/commit/5c1cfcdd2851076571b8350d7301eae9a9ed4004))

## [1.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.11...plugins-destination-postgresql-v1.4.0) (2022-10-31)


### Features

* Update all plugins to SDK with metrics and DFS scheduler ([#3286](https://github.com/cloudquery/cloudquery/issues/3286)) ([a35b8e8](https://github.com/cloudquery/cloudquery/commit/a35b8e89d625287a9b9406ff18cfac78ffdb1241))

## [1.3.11](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.10...plugins-destination-postgresql-v1.3.11) (2022-10-27)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.14 ([#3248](https://github.com/cloudquery/cloudquery/issues/3248)) ([90e2fba](https://github.com/cloudquery/cloudquery/commit/90e2fba5052da7324e1ef32c275d872385052a3f))

## [1.3.10](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.9...plugins-destination-postgresql-v1.3.10) (2022-10-27)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.13 ([#3223](https://github.com/cloudquery/cloudquery/issues/3223)) ([1e90c14](https://github.com/cloudquery/cloudquery/commit/1e90c14fc43752ec67d3faf46c4f7b70ed9fa902))

## [1.3.9](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.8...plugins-destination-postgresql-v1.3.9) (2022-10-20)


### Bug Fixes

* Support Migrations in Custom Schemas ([#3144](https://github.com/cloudquery/cloudquery/issues/3144)) ([92ecd66](https://github.com/cloudquery/cloudquery/commit/92ecd6691b33ea4407146ea4282aee7ec7ccfb20))

## [1.3.8](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.7...plugins-destination-postgresql-v1.3.8) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.12 ([#3110](https://github.com/cloudquery/cloudquery/issues/3110)) ([9ec603a](https://github.com/cloudquery/cloudquery/commit/9ec603a7ecd0c55a043ea4c621c31fce6fe55625))

## [1.3.7](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.6...plugins-destination-postgresql-v1.3.7) (2022-10-19)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.11 ([#3044](https://github.com/cloudquery/cloudquery/issues/3044)) ([4281cbe](https://github.com/cloudquery/cloudquery/commit/4281cbe5ab17f5ec64f8794ede6a3dc4b6341589))

## [1.3.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.5...plugins-destination-postgresql-v1.3.6) (2022-10-19)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.10 ([#3025](https://github.com/cloudquery/cloudquery/issues/3025)) ([51b89c9](https://github.com/cloudquery/cloudquery/commit/51b89c9a3c527c933d5da30e1b584a19bdf9d00d))

## [1.3.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.4...plugins-destination-postgresql-v1.3.5) (2022-10-18)


### Bug Fixes

* **postgresql:** Structure plugin so version is embedded by Go Releaser ([#2947](https://github.com/cloudquery/cloudquery/issues/2947)) ([7535f31](https://github.com/cloudquery/cloudquery/commit/7535f319f38f80d4a0fb3a72504f38161f22ee9e))

## [1.3.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.3...plugins-destination-postgresql-v1.3.4) (2022-10-18)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.9 ([#2937](https://github.com/cloudquery/cloudquery/issues/2937)) ([89fd353](https://github.com/cloudquery/cloudquery/commit/89fd3532e73274185c66ed588de9d7e5cb7c4d69))

## [1.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.2...plugins-destination-postgresql-v1.3.3) (2022-10-16)


### Bug Fixes

* Improve PG Write Speeds ([#2887](https://github.com/cloudquery/cloudquery/issues/2887)) ([80cfed7](https://github.com/cloudquery/cloudquery/commit/80cfed7a7b0d005bd16a682b48cc00b380c41df0))

## [1.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.1...plugins-destination-postgresql-v1.3.2) (2022-10-14)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.8 ([#2871](https://github.com/cloudquery/cloudquery/issues/2871)) ([98cf7ea](https://github.com/cloudquery/cloudquery/commit/98cf7ea22eb8b0b33f8d4b177aa68e7ca110eaa4))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.3.0...plugins-destination-postgresql-v1.3.1) (2022-10-13)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.7 ([#2798](https://github.com/cloudquery/cloudquery/issues/2798)) ([3aca9fa](https://github.com/cloudquery/cloudquery/commit/3aca9faec9f139efbc99bf8ddf338490f19f790c))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.2.2...plugins-destination-postgresql-v1.3.0) (2022-10-13)


### Features

* **postgresql:** Add interval column type for `schema.TimeInterval` ([#2167](https://github.com/cloudquery/cloudquery/issues/2167)) ([f2e6a53](https://github.com/cloudquery/cloudquery/commit/f2e6a532c74bed1b20aae7d74b4a923eb533953c))

## [1.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.2.1...plugins-destination-postgresql-v1.2.2) (2022-10-12)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.6 ([#2731](https://github.com/cloudquery/cloudquery/issues/2731)) ([8916d33](https://github.com/cloudquery/cloudquery/commit/8916d3350a6c62ff0e6bd7ec4e9e8672878af54c))

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.2.0...plugins-destination-postgresql-v1.2.1) (2022-10-12)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.13.1 ([#2613](https://github.com/cloudquery/cloudquery/issues/2613)) ([63d3f22](https://github.com/cloudquery/cloudquery/commit/63d3f22dff516dcc4f315e3dec32e651fbccdc84))
* **deps:** Update plugin-sdk for postgresql to v0.13.5 ([#2665](https://github.com/cloudquery/cloudquery/issues/2665)) ([f274ba6](https://github.com/cloudquery/cloudquery/commit/f274ba6047dd7f568a3c8214e7d384b0858babcc))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.1.0...plugins-destination-postgresql-v1.2.0) (2022-10-11)


### Features

* **pg:** Use pgx batch to improve performance ([#2604](https://github.com/cloudquery/cloudquery/issues/2604)) ([ed16cef](https://github.com/cloudquery/cloudquery/commit/ed16cef82213be1d7a5282e3c598dc5a06a45adb))


### Bug Fixes

* Update postgresql plugin to work with new sdk (deleteStale feature) ([#2587](https://github.com/cloudquery/cloudquery/issues/2587)) ([f5f9257](https://github.com/cloudquery/cloudquery/commit/f5f9257a4eb1c153401d792a65aa65698a0fb321))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.0.1...plugins-destination-postgresql-v1.1.0) (2022-10-09)


### Features

* **pg:** Support CockroachDB ([#2531](https://github.com/cloudquery/cloudquery/issues/2531)) ([0b4c2c1](https://github.com/cloudquery/cloudquery/commit/0b4c2c1589b49708d57a18d4c8cf3fae5fe05d2c))


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.12.10 ([#2556](https://github.com/cloudquery/cloudquery/issues/2556)) ([571346e](https://github.com/cloudquery/cloudquery/commit/571346e29311d4a280310773817cdce4fcf00d94))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v1.0.0...plugins-destination-postgresql-v1.0.1) (2022-10-09)


### Bug Fixes

* **deps:** Update plugin-sdk for postgresql to v0.12.3 ([#2366](https://github.com/cloudquery/cloudquery/issues/2366)) ([6e5c2f6](https://github.com/cloudquery/cloudquery/commit/6e5c2f6822f512ec7a2682cd905986ec4891a77b))
* **deps:** Update plugin-sdk for postgresql to v0.12.4 ([#2406](https://github.com/cloudquery/cloudquery/issues/2406)) ([04b15ad](https://github.com/cloudquery/cloudquery/commit/04b15ad887a665df81c1ff8cf7fccee3f746b680))
* **deps:** Update plugin-sdk for postgresql to v0.12.5 ([#2427](https://github.com/cloudquery/cloudquery/issues/2427)) ([09d78cb](https://github.com/cloudquery/cloudquery/commit/09d78cbe0d7d5fbdbd2d333745a23ae2f3ba592d))
* **deps:** Update plugin-sdk for postgresql to v0.12.6 ([#2442](https://github.com/cloudquery/cloudquery/issues/2442)) ([41a31f3](https://github.com/cloudquery/cloudquery/commit/41a31f33e8c993ab33b90e6eee0632bf34542517))
* **deps:** Update plugin-sdk for postgresql to v0.12.7 ([#2455](https://github.com/cloudquery/cloudquery/issues/2455)) ([1e8c0d8](https://github.com/cloudquery/cloudquery/commit/1e8c0d8a8d5159c8abf0dad35e2f21610c1bd9f2))
* **deps:** Update plugin-sdk for postgresql to v0.12.8 ([#2506](https://github.com/cloudquery/cloudquery/issues/2506)) ([9d0356b](https://github.com/cloudquery/cloudquery/commit/9d0356b5aebecae1d167611d1bc6e304b1c3616f))
* **deps:** Update plugin-sdk for postgresql to v0.12.9 ([#2519](https://github.com/cloudquery/cloudquery/issues/2519)) ([91f4116](https://github.com/cloudquery/cloudquery/commit/91f4116e6170f3442cbe6532e94486a40941277d))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v0.3.0...plugins-destination-postgresql-v1.0.0) (2022-10-04)


### ⚠ BREAKING CHANGES

* Migrate k8s plugin to v2 (#2035)

### Features

* Migrate k8s plugin to v2 ([#2035](https://github.com/cloudquery/cloudquery/issues/2035)) ([955b742](https://github.com/cloudquery/cloudquery/commit/955b742c5be5d1419b671b1723efaca8032f48b2))
* **pg:** Add sentry ([#2139](https://github.com/cloudquery/cloudquery/issues/2139)) ([8891808](https://github.com/cloudquery/cloudquery/commit/8891808222dfcafb35d0820bb426203ba4f4e674))
* **postgresql:** Add support for overwrite-delete-stale ([#2220](https://github.com/cloudquery/cloudquery/issues/2220)) ([efdd136](https://github.com/cloudquery/cloudquery/commit/efdd136bdcf872f7a6104f23429e7ebfb4a7c7c6))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.2 ([#2162](https://github.com/cloudquery/cloudquery/issues/2162)) ([5701aa5](https://github.com/cloudquery/cloudquery/commit/5701aa5b0a8d04e9e99e3efe6e27d5f7ff29b216))
* **deps:** Update plugin-sdk for postgresql to v0.12.2 ([#2311](https://github.com/cloudquery/cloudquery/issues/2311)) ([e4778ea](https://github.com/cloudquery/cloudquery/commit/e4778ea72d9919af7228c07b9165df5a18c0fe71))
* **postgresql:** Upsert in postgresql ([#2133](https://github.com/cloudquery/cloudquery/issues/2133)) ([565728b](https://github.com/cloudquery/cloudquery/commit/565728b79d183081c27642f6bda0be9bdfaf25b3))

## [0.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v0.2.1...plugins-destination-postgresql-v0.3.0) (2022-09-28)


### Features

* Move to standalone postgresql plugin ([#2074](https://github.com/cloudquery/cloudquery/issues/2074)) ([a0de6d3](https://github.com/cloudquery/cloudquery/commit/a0de6d3dfc0f43aad9b465c469b92a96121db0a1))
* Update log ([#2051](https://github.com/cloudquery/cloudquery/issues/2051)) ([6efdf1c](https://github.com/cloudquery/cloudquery/commit/6efdf1c44434f7de13ab25a1f48d47a877c26a79))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))

## [0.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v0.2.0...plugins-destination-postgresql-v0.2.1) (2022-09-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))

## [0.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v0.1.0...plugins-destination-postgresql-v0.2.0) (2022-09-26)


### Features

* Move to standalone postgresql plugin ([#2074](https://github.com/cloudquery/cloudquery/issues/2074)) ([a0de6d3](https://github.com/cloudquery/cloudquery/commit/a0de6d3dfc0f43aad9b465c469b92a96121db0a1))

## [0.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-postgresql-v0.0.1...plugins-destination-postgresql-v0.1.0) (2022-09-25)


### Features

* Update log ([#2051](https://github.com/cloudquery/cloudquery/issues/2051)) ([6efdf1c](https://github.com/cloudquery/cloudquery/commit/6efdf1c44434f7de13ab25a1f48d47a877c26a79))
