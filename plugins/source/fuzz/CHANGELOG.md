# Changelog

<!-- Release notes generated using configuration in .github/release.yml at main -->

## What's Changed
### Breaking Changes 🛠
* fix!: Fix Azure credential chain by @hermanschaaf in https://github.com/cloudquery/cloudquery/pull/1283
### Added 🎉
* feat: Implement EC2 Key Pairs (#1403) by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1325
* feat: Add aws elasticache resources by @shimonp21 in https://github.com/cloudquery/cloudquery/pull/1327
* feat: Add fsx filesystems by @irmatov in https://github.com/cloudquery/cloudquery/pull/1277
* feat: Add website, docs and blog to our main repo by @yevgenypats in https://github.com/cloudquery/cloudquery/pull/1159
* feat: Add AWS SSM Parameters resource by @spangenberg in https://github.com/cloudquery/cloudquery/pull/1222
* feat: Add IAM policy tags by @sunil494 in https://github.com/cloudquery/cloudquery/pull/1433
* feat: Add tags for Glue Databases by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1326
* feat: Add fsx snapshots by @irmatov in https://github.com/cloudquery/cloudquery/pull/1278
* feat: Extend sns subscription data by @irmatov in https://github.com/cloudquery/cloudquery/pull/1424
* feat: Add transfer servers by @irmatov in https://github.com/cloudquery/cloudquery/pull/1284
* feat: Website sitemap by @dj-stormtrooper in https://github.com/cloudquery/cloudquery/pull/1471
* feat: Add queries examples section to the landing page by @anton-kuptsov in https://github.com/cloudquery/cloudquery/pull/1472
* feat: Add fsx data repo tasks by @irmatov in https://github.com/cloudquery/cloudquery/pull/1279
* feat: Add fsx data repo associations by @irmatov in https://github.com/cloudquery/cloudquery/pull/1280
* feat: Add fsx storage virtual machines by @irmatov in https://github.com/cloudquery/cloudquery/pull/1296
* feat: Add cq-gen hcl and a new field to sqs queues by @irmatov in https://github.com/cloudquery/cloudquery/pull/1453
* feat: Added azure cdn profiles by @amanenk in https://github.com/cloudquery/cloudquery/pull/1460
* feat: Add fsx volumes by @irmatov in https://github.com/cloudquery/cloudquery/pull/1322
* feat: Add cq-gen config for apigateway by @irmatov in https://github.com/cloudquery/cloudquery/pull/1541
* feat(website): Plugins layout by @dj-stormtrooper in https://github.com/cloudquery/cloudquery/pull/1548
* feat(website): Remove Tempus logo until permission by @yevgenypats in https://github.com/cloudquery/cloudquery/pull/1550
* feat: Added throttling for digitalocean API calls by @amanenk in https://github.com/cloudquery/cloudquery/pull/1546
### Fixed 🪳
* fix(cli): Don't prepend cq-provider to monorepo binary name by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1264
* fix: Update endpoints by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1273
* fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1272
* fix(deps): Update module github.com/cloudquery/cq-gen to v0.0.9 by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1286
* fix(deps): Update to Go 1.18 by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1292
* fix(deps): Update github-actions by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1293
* fix(cli): Keep docker old entrypoint name by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1297
* fix(cli): Update binary name by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1304
* fix(deps): Update github-actions (major) by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1294
* fix(ec2): Add ARN to key pair by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1339
* fix: Update endpoints by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1432
* fix: Validate github config by @roneli in https://github.com/cloudquery/cloudquery/pull/1438
* fix: Blog titles and og:image from config hook by @dj-stormtrooper in https://github.com/cloudquery/cloudquery/pull/1450
* fix(docs): Add the missing path 'cli' for compiled bin by @hardy4yooz in https://github.com/cloudquery/cloudquery/pull/1457
* fix: K8s Mock Testing by @bbernays in https://github.com/cloudquery/cloudquery/pull/1434
* fix: Website fixes by @dj-stormtrooper in https://github.com/cloudquery/cloudquery/pull/1461
* fix(deps): Update module github.com/cloudquery/cq-gen to v0.0.10 by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1474
* fix: Website links by @dj-stormtrooper in https://github.com/cloudquery/cloudquery/pull/1477
* fix(deps): Update module github.com/cloudquery/cq-gen to v0.0.11 by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1491
* fix: Fix output path of build-cli make command by @hermanschaaf in https://github.com/cloudquery/cloudquery/pull/1464
* fix(deps): Update module github.com/cloudquery/cq-gen to v0.0.12 by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1503
* fix(website): Fix position of corner in top image by @hermanschaaf in https://github.com/cloudquery/cloudquery/pull/1508
* fix: Update endpoints by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1514
* fix(cli): Fix cli version checking to work with monorepo by @hermanschaaf in https://github.com/cloudquery/cloudquery/pull/1510
* fix: ECS Tags by @bbernays in https://github.com/cloudquery/cloudquery/pull/1515
* fix: Update endpoints by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1539
* fix: Fix broken links on website by @anton-kuptsov in https://github.com/cloudquery/cloudquery/pull/1547
* fix: Update endpoints by @cq-bot in https://github.com/cloudquery/cloudquery/pull/1563
* fix: Remove deprecated firebase by @yevgenypats in https://github.com/cloudquery/cloudquery/pull/1568
* fix: Update Okta docs by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1572
* fix: Update Terraform docs by @erezrokah in https://github.com/cloudquery/cloudquery/pull/1573

## New Contributors
* @sunil494 made their first contribution in https://github.com/cloudquery/cloudquery/pull/1433
* @hardy4yooz made their first contribution in https://github.com/cloudquery/cloudquery/pull/1457

**Full Changelog**: https://github.com/cloudquery/cloudquery/compare/plugins/source/fuzz/v0.0.15...plugins/source/fuzz/v0.0.16

## [0.0.15](https://github.com/cloudquery/cloudquery/compare/plugins/source/fuzz-v0.0.14...plugins/source/fuzz/v0.0.15) (2022-08-14)


### Features

* **plugin-source-docs:** Add docs ([#1257](https://github.com/cloudquery/cloudquery/issues/1257)) ([06a8671](https://github.com/cloudquery/cloudquery/commit/06a86719f5bb8699c4926175a731412e43737d51))

## [0.0.14](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.13...v0.0.14) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#49](https://github.com/cloudquery/cq-provider-fuzz/issues/49)) ([4dd0a06](https://github.com/cloudquery/cq-provider-fuzz/commit/4dd0a06b35e8f76f244dd8a0b115bc21fb6a138a))

## [0.0.13](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.12...v0.0.13) (2022-08-07)


### Bug Fixes

* **deps:** Update tubone24/update_release digest to 2146f15 ([#44](https://github.com/cloudquery/cq-provider-fuzz/issues/44)) ([2629965](https://github.com/cloudquery/cq-provider-fuzz/commit/2629965e2740ddb39398f212731c4e5272ab9984))

## [0.0.12](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.11...v0.0.12) (2022-07-27)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.4 ([#35](https://github.com/cloudquery/cq-provider-fuzz/issues/35)) ([83801ad](https://github.com/cloudquery/cq-provider-fuzz/commit/83801add10e1213c7b75db0997fda47db9061250))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.5 ([#38](https://github.com/cloudquery/cq-provider-fuzz/issues/38)) ([a959f34](https://github.com/cloudquery/cq-provider-fuzz/commit/a959f34152711c746502d96623b7c1221d35fc43))
* **deps:** Update tubone24/update_release digest to ef10529 ([#31](https://github.com/cloudquery/cq-provider-fuzz/issues/31)) ([f39994a](https://github.com/cloudquery/cq-provider-fuzz/commit/f39994a60ac6290e65cdc2b5805f25cd36724467))

## [0.0.11](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.10...v0.0.11) (2022-07-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.3 ([#33](https://github.com/cloudquery/cq-provider-fuzz/issues/33)) ([1ddf3dc](https://github.com/cloudquery/cq-provider-fuzz/commit/1ddf3dc4b61cecb2187ffcdde95c1983c23256d1))

## [0.0.10](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.9...v0.0.10) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.2 ([#28](https://github.com/cloudquery/cq-provider-fuzz/issues/28)) ([ccb46a1](https://github.com/cloudquery/cq-provider-fuzz/commit/ccb46a14146e55e3bd5bd61e37f15f18bade6e78))

## [0.0.9](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.8...v0.0.9) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.1 ([#24](https://github.com/cloudquery/cq-provider-fuzz/issues/24)) ([b0e4554](https://github.com/cloudquery/cq-provider-fuzz/commit/b0e4554e52bd64a964183b27e3de0fae65ad65c1))

## [0.0.8](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.7...v0.0.8) (2022-06-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.0 ([#22](https://github.com/cloudquery/cq-provider-fuzz/issues/22)) ([8c3d812](https://github.com/cloudquery/cq-provider-fuzz/commit/8c3d81211302e32e568d9f958466a4243fbd4ceb))

## [0.0.7](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.6...v0.0.7) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.5 ([#18](https://github.com/cloudquery/cq-provider-fuzz/issues/18)) ([81455c5](https://github.com/cloudquery/cq-provider-fuzz/commit/81455c52911d09facf99a02436d31459dd56d5fe))

## [0.0.6](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.5...v0.0.6) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.4 ([#16](https://github.com/cloudquery/cq-provider-fuzz/issues/16)) ([9915e1c](https://github.com/cloudquery/cq-provider-fuzz/commit/9915e1c435d86d3a2dd632c562125120005102a3))

## [0.0.5](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.4...v0.0.5) (2022-06-26)


### Bug Fixes

* Embed correct version ([2d0757c](https://github.com/cloudquery/cq-provider-fuzz/commit/2d0757c36f527286ce2cc9f986fdcd5f039665a8))

## [0.0.4](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.3...v0.0.4) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.3 ([#10](https://github.com/cloudquery/cq-provider-fuzz/issues/10)) ([fcdbd54](https://github.com/cloudquery/cq-provider-fuzz/commit/fcdbd545b76015f322119a81c41a1fdc952a2499))
* Embed correct version ([7ecbad0](https://github.com/cloudquery/cq-provider-fuzz/commit/7ecbad0f37376e9361712d001f0a2d95f2e7c814))

## [0.0.3](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.2...v0.0.3) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.2 ([#7](https://github.com/cloudquery/cq-provider-fuzz/issues/7)) ([f9d3e92](https://github.com/cloudquery/cq-provider-fuzz/commit/f9d3e92df35131c0ca8229c60f1d7387d30210e9))

## [0.0.2](https://github.com/cloudquery/cq-provider-fuzz/compare/v0.0.1...v0.0.2) (2022-06-23)


### Bug Fixes

* Embed correct version ([#5](https://github.com/cloudquery/cq-provider-fuzz/issues/5)) ([0db0a81](https://github.com/cloudquery/cq-provider-fuzz/commit/0db0a8125a5172f85d6547b796080794890e17a7))

## 0.0.1 (2022-06-23)


### Features

* Add readme ([c0f7e3b](https://github.com/cloudquery/cq-provider-fuzz/commit/c0f7e3bad2952e2e598f438277c9fb64a52e4b1f))
* Initial commit ([044fcb4](https://github.com/cloudquery/cq-provider-fuzz/commit/044fcb41721bb3af8fefc5009233c61282b53e6a))


### Bug Fixes

* Change defaults ([b762dca](https://github.com/cloudquery/cq-provider-fuzz/commit/b762dca938477e8cc819ced838267920c5602045))
* Improve logging ([a919a7b](https://github.com/cloudquery/cq-provider-fuzz/commit/a919a7b3ecc06710139239f3a29823ec70624049))
* Lint error ([015a913](https://github.com/cloudquery/cq-provider-fuzz/commit/015a913296d27fd364ad506292e1aa383b1a44cd))
* Readme ([a831ac3](https://github.com/cloudquery/cq-provider-fuzz/commit/a831ac325d318054cca7913eaded3ea15816715f))
* Readme ([fdcf924](https://github.com/cloudquery/cq-provider-fuzz/commit/fdcf9242179543031b115da656a20ea2a7901da0))
* Rename env variables ([ea7ee9b](https://github.com/cloudquery/cq-provider-fuzz/commit/ea7ee9b0aaf081d731c0ca8c4b58269dde725e26))
* Update defaults ([0d661e2](https://github.com/cloudquery/cq-provider-fuzz/commit/0d661e2ae9e1058c02226834db955edfd0055eac))


### Miscellaneous Chores

* Release 0.0.1 ([4b72cb6](https://github.com/cloudquery/cq-provider-fuzz/commit/4b72cb6fbf05efafe922b3d4a840d2750edc03db))
