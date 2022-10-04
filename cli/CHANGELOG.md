# Changelog

All notable changes to CloudQuery will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0-pre.1](https://github.com/cloudquery/cloudquery/compare/cli-v1.2.1-pre.1...cli-v1.3.0-pre.1) (2022-09-30)


### Features

* **cli:** Use SourceClient and DestinationClient directly ([#2165](https://github.com/cloudquery/cloudquery/issues/2165)) ([e594e61](https://github.com/cloudquery/cloudquery/commit/e594e615de5217ca695592a23a2a18e9fee9cfe7))


### Bug Fixes

* ProgressBar output fix ([#2163](https://github.com/cloudquery/cloudquery/issues/2163)) ([85fbd58](https://github.com/cloudquery/cloudquery/commit/85fbd58ceb755f870b097731e33f36b70d9d97db))

## [1.2.1-pre.1](https://github.com/cloudquery/cloudquery/compare/cli-v1.2.0-pre.1...cli-v1.2.1-pre.1) (2022-09-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* Version fetching code and test ([#2097](https://github.com/cloudquery/cloudquery/issues/2097)) ([f174f30](https://github.com/cloudquery/cloudquery/commit/f174f30822d390e41114828fbb54c8505ec02cd5))

## [1.2.0-pre.1](https://github.com/cloudquery/cloudquery/compare/cli-v1.1.0-pre.1...cli-v1.2.0-pre.1) (2022-09-27)


### Features

* Remove spinner ([#2089](https://github.com/cloudquery/cloudquery/issues/2089)) ([5915181](https://github.com/cloudquery/cloudquery/commit/5915181287d1ddb7a01b06c8ad8732050f1bd4e0))
* Update SDK to v0.10.1 ([#2100](https://github.com/cloudquery/cloudquery/issues/2100)) ([d968032](https://github.com/cloudquery/cloudquery/commit/d9680327b600d0c52c1d18dacc09d2d825564b20))

## [1.1.0-pre.1](https://github.com/cloudquery/cloudquery/compare/cli-v1.1.0-pre.0...cli-v1.1.0-pre.1) (2022-09-26)

### ⚠ BREAKING CHANGES

- **cli:** Remove gen command (#2022)
- CloudQuery V2. (#1463)

### Features

- **cli:** Disable sentry in development ([#1939](https://github.com/cloudquery/cloudquery/issues/1939)) ([e2c417e](https://github.com/cloudquery/cloudquery/commit/e2c417e42e018f1e6e6bbc2b6fe9ace0c990b30c))
- **cli:** Remove gen command ([#2022](https://github.com/cloudquery/cloudquery/issues/2022)) ([83a32dd](https://github.com/cloudquery/cloudquery/commit/83a32ddb6d2973c2235975cf437862fed40371a0))
- CloudQuery V2. ([#1463](https://github.com/cloudquery/cloudquery/issues/1463)) ([d1799f3](https://github.com/cloudquery/cloudquery/commit/d1799f347d1387dcc2b9a4f05aa2f48999ed1090))
- Create directory structure when generating configuration ([#1845](https://github.com/cloudquery/cloudquery/issues/1845)) ([4f9c8de](https://github.com/cloudquery/cloudquery/commit/4f9c8de59ac2e9d6c88e6dd9c4ae6cf457359e61))
- **gcp:** Remove Classify and IgnoreError ([#1757](https://github.com/cloudquery/cloudquery/issues/1757)) ([3d34ca5](https://github.com/cloudquery/cloudquery/commit/3d34ca526941b1579f5c6f4360c2d364bfe96cc2))
- Generate `postgresql` as default destination in `gen` command ([#1863](https://github.com/cloudquery/cloudquery/issues/1863)) ([aad6218](https://github.com/cloudquery/cloudquery/commit/aad6218b92684620850f5cffb7d7ebdaa9997392))
- Generate auto-filled config ([#1764](https://github.com/cloudquery/cloudquery/issues/1764)) ([2255404](https://github.com/cloudquery/cloudquery/commit/2255404012afa97d38b64c11d2f66405dfa84c6f))
- Move to standalone postgresql plugin ([#2074](https://github.com/cloudquery/cloudquery/issues/2074)) ([a0de6d3](https://github.com/cloudquery/cloudquery/commit/a0de6d3dfc0f43aad9b465c469b92a96121db0a1))
- **sync:** Default input directory to '.' ([#1869](https://github.com/cloudquery/cloudquery/issues/1869)) ([005c915](https://github.com/cloudquery/cloudquery/commit/005c9152f7dee2f081d2020612603e1764ea6a5e))
- Use jsonb for json columns ([#1870](https://github.com/cloudquery/cloudquery/issues/1870)) ([78e37fc](https://github.com/cloudquery/cloudquery/commit/78e37fcdd8f960739031c67771ec9ee751bc12c4))
- Use new tag format to download plugins ([#1985](https://github.com/cloudquery/cloudquery/issues/1985)) ([583f54e](https://github.com/cloudquery/cloudquery/commit/583f54ea3ecb8f96b6f76498ddca29770c691831))
- Use spinner instead of progress bar ([#1829](https://github.com/cloudquery/cloudquery/issues/1829)) ([af9129e](https://github.com/cloudquery/cloudquery/commit/af9129ef7f5dea2f87dc0a7e3f5668eb0b261dbe))

### Bug Fixes

- Can't gen with specific version ([#1915](https://github.com/cloudquery/cloudquery/issues/1915)) ([b7ae169](https://github.com/cloudquery/cloudquery/commit/b7ae169974894850f80949e3060c57aad4bb69f1))
- Checksum validation ([#1637](https://github.com/cloudquery/cloudquery/issues/1637)) ([a899cce](https://github.com/cloudquery/cloudquery/commit/a899ccef5f22eb90bb195e902fd766d5c921c916))
- CLI dev version ([#1864](https://github.com/cloudquery/cloudquery/issues/1864)) ([9c18606](https://github.com/cloudquery/cloudquery/commit/9c186062cdfe3fd77292deca5c45aa4d5ede320b))
- **cli:** Fix cli version checking to work with monorepo ([#1510](https://github.com/cloudquery/cloudquery/issues/1510)) ([e0ddfcf](https://github.com/cloudquery/cloudquery/commit/e0ddfcf4f12f991db00c0706bbfc430dc0dd673d))
- **cli:** Remove debug output printed to stdout ([#1929](https://github.com/cloudquery/cloudquery/issues/1929)) ([eec9627](https://github.com/cloudquery/cloudquery/commit/eec962745edf5c1dab8c4bd8aef11ef0f0c22b56))
- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to dc4ac0a ([#1650](https://github.com/cloudquery/cloudquery/issues/1650)) ([9222b5a](https://github.com/cloudquery/cloudquery/commit/9222b5aaadcf3a81fcf5d6d723755f48398a0797))
- **deps:** Update github.com/ProtonMail/go-crypto digest to 4b6e5c5 ([#1649](https://github.com/cloudquery/cloudquery/issues/1649)) ([c71d859](https://github.com/cloudquery/cloudquery/commit/c71d85915f9a80a589d5942728ac15233b2dff49))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.1.2 ([#1750](https://github.com/cloudquery/cloudquery/issues/1750)) ([fbe1b78](https://github.com/cloudquery/cloudquery/commit/fbe1b7835b0677a3d1c79bc10d95b991e2eb5129))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.4 ([#1761](https://github.com/cloudquery/cloudquery/issues/1761)) ([7a83a65](https://github.com/cloudquery/cloudquery/commit/7a83a65446119b5339f5f3a3759f7f160a3716b4))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.5 ([#1769](https://github.com/cloudquery/cloudquery/issues/1769)) ([c9c8c05](https://github.com/cloudquery/cloudquery/commit/c9c8c05b97ae349d7a21deef1e524d8a180c512d))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.6 ([#1770](https://github.com/cloudquery/cloudquery/issues/1770)) ([5bc205e](https://github.com/cloudquery/cloudquery/commit/5bc205ec7f4a7b2fa8a34793e8d746a43cbb03ed))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.7 ([#1783](https://github.com/cloudquery/cloudquery/issues/1783)) ([c291499](https://github.com/cloudquery/cloudquery/commit/c2914999f8607b6d313c346cc037829d21f84cfb))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.8 ([#1784](https://github.com/cloudquery/cloudquery/issues/1784)) ([b64e2d1](https://github.com/cloudquery/cloudquery/commit/b64e2d18abfaeb6056b6a3aed56f8fdc07e7d535))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.9 ([#1785](https://github.com/cloudquery/cloudquery/issues/1785)) ([c6e8cb0](https://github.com/cloudquery/cloudquery/commit/c6e8cb03c5851f96fef09e11d8c8d34ac74a70fb))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.4.0 ([#1786](https://github.com/cloudquery/cloudquery/issues/1786)) ([cba274b](https://github.com/cloudquery/cloudquery/commit/cba274b3dda610e06129b843ed0c1376f83f09bb))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.4.1 ([#1787](https://github.com/cloudquery/cloudquery/issues/1787)) ([bad385c](https://github.com/cloudquery/cloudquery/commit/bad385c39bbed55a74894591f7cfdd092bcded55))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.4.2 ([#1789](https://github.com/cloudquery/cloudquery/issues/1789)) ([79a46a2](https://github.com/cloudquery/cloudquery/commit/79a46a2b4719d66df40db128f841e9c3640a9128))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.5.0 ([#1792](https://github.com/cloudquery/cloudquery/issues/1792)) ([0b4834e](https://github.com/cloudquery/cloudquery/commit/0b4834e38bb66d65bd6d44c84847d3538e264d2c))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.5.2 ([#1793](https://github.com/cloudquery/cloudquery/issues/1793)) ([36fd6a1](https://github.com/cloudquery/cloudquery/commit/36fd6a18acd5c405bba8c1c52df294e056533ee0))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.0 ([#1817](https://github.com/cloudquery/cloudquery/issues/1817)) ([bd68a9c](https://github.com/cloudquery/cloudquery/commit/bd68a9c8b691f7af4c956259a06eb18ac50b374a))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.1 ([#1820](https://github.com/cloudquery/cloudquery/issues/1820)) ([2613e23](https://github.com/cloudquery/cloudquery/commit/2613e2374ea451e7bf031bda8ea26e895e65528c))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.2 ([#1838](https://github.com/cloudquery/cloudquery/issues/1838)) ([5b16c59](https://github.com/cloudquery/cloudquery/commit/5b16c59dd415cf0a775dbc38cd62c99b97f04ea5))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.3 ([#1858](https://github.com/cloudquery/cloudquery/issues/1858)) ([9e3ace7](https://github.com/cloudquery/cloudquery/commit/9e3ace775da2d600968ef4275e9e0013d4dfd825))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.4 ([#1862](https://github.com/cloudquery/cloudquery/issues/1862)) ([5d141cf](https://github.com/cloudquery/cloudquery/commit/5d141cf6006e26cf240ddf295dda53c16f7386a4))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.1 ([#1865](https://github.com/cloudquery/cloudquery/issues/1865)) ([474bb70](https://github.com/cloudquery/cloudquery/commit/474bb7081b6e9b6ffc5ac949ed3a664f92083c82))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.12 ([#1916](https://github.com/cloudquery/cloudquery/issues/1916)) ([27d8153](https://github.com/cloudquery/cloudquery/commit/27d81534baaa1312a6bd87294d298dd8b5348a79))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.2 ([#1872](https://github.com/cloudquery/cloudquery/issues/1872)) ([49ed26d](https://github.com/cloudquery/cloudquery/commit/49ed26d231c91ac1b5b00cc55d3d0a8a5a6306f7))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.3 ([#1886](https://github.com/cloudquery/cloudquery/issues/1886)) ([7435d59](https://github.com/cloudquery/cloudquery/commit/7435d593e51ca829d3a328eebc9517e9cb2a4ef0))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.4 ([#1889](https://github.com/cloudquery/cloudquery/issues/1889)) ([63a5362](https://github.com/cloudquery/cloudquery/commit/63a5362995aa680b291f2411d01e776e884896d4))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.9 ([#1891](https://github.com/cloudquery/cloudquery/issues/1891)) ([3469f20](https://github.com/cloudquery/cloudquery/commit/3469f20e76e9dcbf48b9c6e3e7c0c2224c5b8ad3))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))
- **deps:** Update module github.com/hashicorp/go-plugin to v1.4.5 ([#1665](https://github.com/cloudquery/cloudquery/issues/1665)) ([6107aef](https://github.com/cloudquery/cloudquery/commit/6107aef96677e3b44011e8865e47b64d517ae4de))
- **deps:** Update module github.com/mattn/go-isatty to v0.0.16 ([#1666](https://github.com/cloudquery/cloudquery/issues/1666)) ([44f0998](https://github.com/cloudquery/cloudquery/commit/44f099875101c42697b4fe9bf4d4fedcc07a9d72))
- Don't print skip download message to console ([#2008](https://github.com/cloudquery/cloudquery/issues/2008)) ([a947d44](https://github.com/cloudquery/cloudquery/commit/a947d44172a657f3d730b31d3baa9eddce703601))
- **pg:** Fix PKs recreation when nothing changed ([#1938](https://github.com/cloudquery/cloudquery/issues/1938)) ([5db7baa](https://github.com/cloudquery/cloudquery/commit/5db7baa929186e71debfd0cdc5ec4477f42098a3))
- Quote kind string in generated yml files ([#1824](https://github.com/cloudquery/cloudquery/issues/1824)) ([6c184ce](https://github.com/cloudquery/cloudquery/commit/6c184ce2dc927a6a87edc4e0fcc0e18fcb4cc092))
- **release:** Embed correct modules version ([#1849](https://github.com/cloudquery/cloudquery/issues/1849)) ([2095a3b](https://github.com/cloudquery/cloudquery/commit/2095a3be90c7e3986fc4704e6d613dd5a667199c))
- Remove deprecated firebase ([#1568](https://github.com/cloudquery/cloudquery/issues/1568)) ([a879709](https://github.com/cloudquery/cloudquery/commit/a8797092df6c228ae519dbb5af4fb55b4ce0cb52))
- Resolve plugin version from GitHub registry instead of gRPC call ([#1856](https://github.com/cloudquery/cloudquery/issues/1856)) ([14d2ca6](https://github.com/cloudquery/cloudquery/commit/14d2ca6894bdf51597082fbadd1aad7401316e59))
- Use correct binary path on Windows ([#1894](https://github.com/cloudquery/cloudquery/issues/1894)) ([1ee41e7](https://github.com/cloudquery/cloudquery/commit/1ee41e740143663ffe67c48dc5db193c94b3eafb))
- Use postgres defaults in destination config ([#1846](https://github.com/cloudquery/cloudquery/issues/1846)) ([24bbbc1](https://github.com/cloudquery/cloudquery/commit/24bbbc1d8d2bc1d65c54c9adad73a3311e95cb88))
- Use uppercase downloading during progress ([#2006](https://github.com/cloudquery/cloudquery/issues/2006)) ([e6a7a44](https://github.com/cloudquery/cloudquery/commit/e6a7a4465b20a696fa4b773c9beb5481993213aa))

## [1.1.0-pre.0](https://github.com/cloudquery/cloudquery/compare/cli-v1.0.0-pre.0...cli-v1.1.0-pre.0) (2022-09-26)

### Features

- Move to standalone postgresql plugin ([#2074](https://github.com/cloudquery/cloudquery/issues/2074)) ([a0de6d3](https://github.com/cloudquery/cloudquery/commit/a0de6d3dfc0f43aad9b465c469b92a96121db0a1))

## [1.0.0-pre.0](https://github.com/cloudquery/cloudquery/compare/cli-v0.33.4-pre.0...cli-v1.0.0-pre.0) (2022-09-25)

### ⚠ BREAKING CHANGES

- **cli:** Remove gen command (#2022)

### Features

- **cli:** Remove gen command ([#2022](https://github.com/cloudquery/cloudquery/issues/2022)) ([83a32dd](https://github.com/cloudquery/cloudquery/commit/83a32ddb6d2973c2235975cf437862fed40371a0))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))
- Don't print skip download message to console ([#2008](https://github.com/cloudquery/cloudquery/issues/2008)) ([a947d44](https://github.com/cloudquery/cloudquery/commit/a947d44172a657f3d730b31d3baa9eddce703601))
- Use uppercase downloading during progress ([#2006](https://github.com/cloudquery/cloudquery/issues/2006)) ([e6a7a44](https://github.com/cloudquery/cloudquery/commit/e6a7a4465b20a696fa4b773c9beb5481993213aa))

## [0.33.4-pre.0](https://github.com/cloudquery/cloudquery/compare/cli/v0.33.3-pre.0...cli-v0.33.4-pre.0) (2022-09-22)

### Features

- Use new tag format to download plugins ([#1985](https://github.com/cloudquery/cloudquery/issues/1985)) ([583f54e](https://github.com/cloudquery/cloudquery/commit/583f54ea3ecb8f96b6f76498ddca29770c691831))

## [0.33.3-pre.0](https://github.com/cloudquery/cloudquery/compare/cli/v0.33.2-pre.0...cli/v0.33.3-pre.0) (2022-09-22)

### Features

- **cli:** Disable sentry in development ([#1939](https://github.com/cloudquery/cloudquery/issues/1939)) ([e2c417e](https://github.com/cloudquery/cloudquery/commit/e2c417e42e018f1e6e6bbc2b6fe9ace0c990b30c))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.12 ([#1916](https://github.com/cloudquery/cloudquery/issues/1916)) ([27d8153](https://github.com/cloudquery/cloudquery/commit/27d81534baaa1312a6bd87294d298dd8b5348a79))
- **pg:** Fix PKs recreation when nothing changed ([#1938](https://github.com/cloudquery/cloudquery/issues/1938)) ([5db7baa](https://github.com/cloudquery/cloudquery/commit/5db7baa929186e71debfd0cdc5ec4477f42098a3))

## [0.33.2-pre.0](https://github.com/cloudquery/cloudquery/compare/cli/v0.33.1-pre.0...cli/v0.33.2-pre.0) (2022-09-21)

### Bug Fixes

- **cli:** Remove debug output printed to stdout ([#1929](https://github.com/cloudquery/cloudquery/issues/1929)) ([eec9627](https://github.com/cloudquery/cloudquery/commit/eec962745edf5c1dab8c4bd8aef11ef0f0c22b56))

## [0.33.1-pre.0](https://github.com/cloudquery/cloudquery/compare/cli/v0.33.0-pre.0...cli/v0.33.1-pre.0) (2022-09-21)

### Bug Fixes

- Can't gen with specific version ([#1915](https://github.com/cloudquery/cloudquery/issues/1915)) ([b7ae169](https://github.com/cloudquery/cloudquery/commit/b7ae169974894850f80949e3060c57aad4bb69f1))

## [0.33.0-pre.0](https://github.com/cloudquery/cloudquery/compare/cli-v0.32.12-pre.0...cli/v0.33.0-pre.0) (2022-09-21)

### ⚠ BREAKING CHANGES

- CloudQuery V2. (#1463)

### Features

- CloudQuery V2. ([#1463](https://github.com/cloudquery/cloudquery/issues/1463)) ([d1799f3](https://github.com/cloudquery/cloudquery/commit/d1799f347d1387dcc2b9a4f05aa2f48999ed1090))
- Create directory structure when generating configuration ([#1845](https://github.com/cloudquery/cloudquery/issues/1845)) ([4f9c8de](https://github.com/cloudquery/cloudquery/commit/4f9c8de59ac2e9d6c88e6dd9c4ae6cf457359e61))
- **gcp:** Remove Classify and IgnoreError ([#1757](https://github.com/cloudquery/cloudquery/issues/1757)) ([3d34ca5](https://github.com/cloudquery/cloudquery/commit/3d34ca526941b1579f5c6f4360c2d364bfe96cc2))
- Generate `postgresql` as default destination in `gen` command ([#1863](https://github.com/cloudquery/cloudquery/issues/1863)) ([aad6218](https://github.com/cloudquery/cloudquery/commit/aad6218b92684620850f5cffb7d7ebdaa9997392))
- Generate auto-filled config ([#1764](https://github.com/cloudquery/cloudquery/issues/1764)) ([2255404](https://github.com/cloudquery/cloudquery/commit/2255404012afa97d38b64c11d2f66405dfa84c6f))
- **sync:** Default input directory to '.' ([#1869](https://github.com/cloudquery/cloudquery/issues/1869)) ([005c915](https://github.com/cloudquery/cloudquery/commit/005c9152f7dee2f081d2020612603e1764ea6a5e))
- Use jsonb for json columns ([#1870](https://github.com/cloudquery/cloudquery/issues/1870)) ([78e37fc](https://github.com/cloudquery/cloudquery/commit/78e37fcdd8f960739031c67771ec9ee751bc12c4))
- Use spinner instead of progress bar ([#1829](https://github.com/cloudquery/cloudquery/issues/1829)) ([af9129e](https://github.com/cloudquery/cloudquery/commit/af9129ef7f5dea2f87dc0a7e3f5668eb0b261dbe))

### Bug Fixes

- Checksum validation ([#1637](https://github.com/cloudquery/cloudquery/issues/1637)) ([a899cce](https://github.com/cloudquery/cloudquery/commit/a899ccef5f22eb90bb195e902fd766d5c921c916))
- CLI dev version ([#1864](https://github.com/cloudquery/cloudquery/issues/1864)) ([9c18606](https://github.com/cloudquery/cloudquery/commit/9c186062cdfe3fd77292deca5c45aa4d5ede320b))
- **cli-docs:** Fix typos in docs/index ([#1256](https://github.com/cloudquery/cloudquery/issues/1256)) ([424da6d](https://github.com/cloudquery/cloudquery/commit/424da6da710f8294398f49e5bb7f837036f9c81a))
- **cli:** Don't pre-pend cq-provider to monorepo binary name ([#1264](https://github.com/cloudquery/cloudquery/issues/1264)) ([2b1e082](https://github.com/cloudquery/cloudquery/commit/2b1e0820ff731c5d8a6dca3036572aba674e34b5))
- **cli:** Fix cli version checking to work with monorepo ([#1510](https://github.com/cloudquery/cloudquery/issues/1510)) ([e0ddfcf](https://github.com/cloudquery/cloudquery/commit/e0ddfcf4f12f991db00c0706bbfc430dc0dd673d))
- **cli:** Keep old entrypoint name ([#1297](https://github.com/cloudquery/cloudquery/issues/1297)) ([bab4f39](https://github.com/cloudquery/cloudquery/commit/bab4f3972bc8853cdb3bf74cde8da01a399e182b))
- **cli:** Update binary name ([#1304](https://github.com/cloudquery/cloudquery/issues/1304)) ([432c404](https://github.com/cloudquery/cloudquery/commit/432c40444fafc78f9d7a5c882c6203e3433c8627))
- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to dc4ac0a ([#1650](https://github.com/cloudquery/cloudquery/issues/1650)) ([9222b5a](https://github.com/cloudquery/cloudquery/commit/9222b5aaadcf3a81fcf5d6d723755f48398a0797))
- **deps:** Update github.com/ProtonMail/go-crypto digest to 4b6e5c5 ([#1649](https://github.com/cloudquery/cloudquery/issues/1649)) ([c71d859](https://github.com/cloudquery/cloudquery/commit/c71d85915f9a80a589d5942728ac15233b2dff49))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#1272](https://github.com/cloudquery/cloudquery/issues/1272)) ([8546173](https://github.com/cloudquery/cloudquery/commit/85461731a03c9d2e5f84267e5eb7012226389a24))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.1.2 ([#1750](https://github.com/cloudquery/cloudquery/issues/1750)) ([fbe1b78](https://github.com/cloudquery/cloudquery/commit/fbe1b7835b0677a3d1c79bc10d95b991e2eb5129))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.4 ([#1761](https://github.com/cloudquery/cloudquery/issues/1761)) ([7a83a65](https://github.com/cloudquery/cloudquery/commit/7a83a65446119b5339f5f3a3759f7f160a3716b4))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.5 ([#1769](https://github.com/cloudquery/cloudquery/issues/1769)) ([c9c8c05](https://github.com/cloudquery/cloudquery/commit/c9c8c05b97ae349d7a21deef1e524d8a180c512d))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.6 ([#1770](https://github.com/cloudquery/cloudquery/issues/1770)) ([5bc205e](https://github.com/cloudquery/cloudquery/commit/5bc205ec7f4a7b2fa8a34793e8d746a43cbb03ed))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.7 ([#1783](https://github.com/cloudquery/cloudquery/issues/1783)) ([c291499](https://github.com/cloudquery/cloudquery/commit/c2914999f8607b6d313c346cc037829d21f84cfb))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.8 ([#1784](https://github.com/cloudquery/cloudquery/issues/1784)) ([b64e2d1](https://github.com/cloudquery/cloudquery/commit/b64e2d18abfaeb6056b6a3aed56f8fdc07e7d535))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.2.9 ([#1785](https://github.com/cloudquery/cloudquery/issues/1785)) ([c6e8cb0](https://github.com/cloudquery/cloudquery/commit/c6e8cb03c5851f96fef09e11d8c8d34ac74a70fb))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.4.0 ([#1786](https://github.com/cloudquery/cloudquery/issues/1786)) ([cba274b](https://github.com/cloudquery/cloudquery/commit/cba274b3dda610e06129b843ed0c1376f83f09bb))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.4.1 ([#1787](https://github.com/cloudquery/cloudquery/issues/1787)) ([bad385c](https://github.com/cloudquery/cloudquery/commit/bad385c39bbed55a74894591f7cfdd092bcded55))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.4.2 ([#1789](https://github.com/cloudquery/cloudquery/issues/1789)) ([79a46a2](https://github.com/cloudquery/cloudquery/commit/79a46a2b4719d66df40db128f841e9c3640a9128))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.5.0 ([#1792](https://github.com/cloudquery/cloudquery/issues/1792)) ([0b4834e](https://github.com/cloudquery/cloudquery/commit/0b4834e38bb66d65bd6d44c84847d3538e264d2c))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.5.2 ([#1793](https://github.com/cloudquery/cloudquery/issues/1793)) ([36fd6a1](https://github.com/cloudquery/cloudquery/commit/36fd6a18acd5c405bba8c1c52df294e056533ee0))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.0 ([#1817](https://github.com/cloudquery/cloudquery/issues/1817)) ([bd68a9c](https://github.com/cloudquery/cloudquery/commit/bd68a9c8b691f7af4c956259a06eb18ac50b374a))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.1 ([#1820](https://github.com/cloudquery/cloudquery/issues/1820)) ([2613e23](https://github.com/cloudquery/cloudquery/commit/2613e2374ea451e7bf031bda8ea26e895e65528c))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.2 ([#1838](https://github.com/cloudquery/cloudquery/issues/1838)) ([5b16c59](https://github.com/cloudquery/cloudquery/commit/5b16c59dd415cf0a775dbc38cd62c99b97f04ea5))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.3 ([#1858](https://github.com/cloudquery/cloudquery/issues/1858)) ([9e3ace7](https://github.com/cloudquery/cloudquery/commit/9e3ace775da2d600968ef4275e9e0013d4dfd825))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.4 ([#1862](https://github.com/cloudquery/cloudquery/issues/1862)) ([5d141cf](https://github.com/cloudquery/cloudquery/commit/5d141cf6006e26cf240ddf295dda53c16f7386a4))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.1 ([#1865](https://github.com/cloudquery/cloudquery/issues/1865)) ([474bb70](https://github.com/cloudquery/cloudquery/commit/474bb7081b6e9b6ffc5ac949ed3a664f92083c82))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.2 ([#1872](https://github.com/cloudquery/cloudquery/issues/1872)) ([49ed26d](https://github.com/cloudquery/cloudquery/commit/49ed26d231c91ac1b5b00cc55d3d0a8a5a6306f7))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.3 ([#1886](https://github.com/cloudquery/cloudquery/issues/1886)) ([7435d59](https://github.com/cloudquery/cloudquery/commit/7435d593e51ca829d3a328eebc9517e9cb2a4ef0))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.4 ([#1889](https://github.com/cloudquery/cloudquery/issues/1889)) ([63a5362](https://github.com/cloudquery/cloudquery/commit/63a5362995aa680b291f2411d01e776e884896d4))
- **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.9 ([#1891](https://github.com/cloudquery/cloudquery/issues/1891)) ([3469f20](https://github.com/cloudquery/cloudquery/commit/3469f20e76e9dcbf48b9c6e3e7c0c2224c5b8ad3))
- **deps:** Update module github.com/hashicorp/go-plugin to v1.4.5 ([#1665](https://github.com/cloudquery/cloudquery/issues/1665)) ([6107aef](https://github.com/cloudquery/cloudquery/commit/6107aef96677e3b44011e8865e47b64d517ae4de))
- **deps:** Update module github.com/mattn/go-isatty to v0.0.16 ([#1666](https://github.com/cloudquery/cloudquery/issues/1666)) ([44f0998](https://github.com/cloudquery/cloudquery/commit/44f099875101c42697b4fe9bf4d4fedcc07a9d72))
- Quote kind string in generated yml files ([#1824](https://github.com/cloudquery/cloudquery/issues/1824)) ([6c184ce](https://github.com/cloudquery/cloudquery/commit/6c184ce2dc927a6a87edc4e0fcc0e18fcb4cc092))
- **release:** Embed correct modules version ([#1849](https://github.com/cloudquery/cloudquery/issues/1849)) ([2095a3b](https://github.com/cloudquery/cloudquery/commit/2095a3be90c7e3986fc4704e6d613dd5a667199c))
- Remove deprecated firebase ([#1568](https://github.com/cloudquery/cloudquery/issues/1568)) ([a879709](https://github.com/cloudquery/cloudquery/commit/a8797092df6c228ae519dbb5af4fb55b4ce0cb52))
- Resolve plugin version from GitHub registry instead of gRPC call ([#1856](https://github.com/cloudquery/cloudquery/issues/1856)) ([14d2ca6](https://github.com/cloudquery/cloudquery/commit/14d2ca6894bdf51597082fbadd1aad7401316e59))
- Use correct binary path on Windows ([#1894](https://github.com/cloudquery/cloudquery/issues/1894)) ([1ee41e7](https://github.com/cloudquery/cloudquery/commit/1ee41e740143663ffe67c48dc5db193c94b3eafb))
- Use postgres defaults in destination config ([#1846](https://github.com/cloudquery/cloudquery/issues/1846)) ([24bbbc1](https://github.com/cloudquery/cloudquery/commit/24bbbc1d8d2bc1d65c54c9adad73a3311e95cb88))

## [0.32.12](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.11...cli/v0.32.12) (2022-09-13)

### Bug Fixes

- **cli:** Windows verification check ([#1804](https://github.com/cloudquery/cloudquery/issues/1804)) ([aed04ed](https://github.com/cloudquery/cloudquery/commit/aed04edc30c1af41c68405469f01f5cf7ae02ea5))

## [0.32.11](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.10...cli/v0.32.11) (2022-09-01)

### Bug Fixes

- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to dc4ac0a ([#1650](https://github.com/cloudquery/cloudquery/issues/1650)) ([9222b5a](https://github.com/cloudquery/cloudquery/commit/9222b5aaadcf3a81fcf5d6d723755f48398a0797))
- **deps:** Update github.com/ProtonMail/go-crypto digest to 4b6e5c5 ([#1649](https://github.com/cloudquery/cloudquery/issues/1649)) ([c71d859](https://github.com/cloudquery/cloudquery/commit/c71d85915f9a80a589d5942728ac15233b2dff49))
- **deps:** Update module github.com/hashicorp/go-plugin to v1.4.5 ([#1665](https://github.com/cloudquery/cloudquery/issues/1665)) ([6107aef](https://github.com/cloudquery/cloudquery/commit/6107aef96677e3b44011e8865e47b64d517ae4de))
- **deps:** Update module github.com/mattn/go-isatty to v0.0.16 ([#1666](https://github.com/cloudquery/cloudquery/issues/1666)) ([44f0998](https://github.com/cloudquery/cloudquery/commit/44f099875101c42697b4fe9bf4d4fedcc07a9d72))

## [0.32.10](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.9...cli/v0.32.10) (2022-08-31)

### Bug Fixes

- Checksum validation ([#1637](https://github.com/cloudquery/cloudquery/issues/1637)) ([a899cce](https://github.com/cloudquery/cloudquery/commit/a899ccef5f22eb90bb195e902fd766d5c921c916))

## [0.32.9](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.8...cli/v0.32.9) (2022-08-27)

### Bug Fixes

- Remove deprecated firebase ([#1568](https://github.com/cloudquery/cloudquery/issues/1568)) ([a879709](https://github.com/cloudquery/cloudquery/commit/a8797092df6c228ae519dbb5af4fb55b4ce0cb52))

## [0.32.8](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.7...cli/v0.32.8) (2022-08-21)

### Bug Fixes

- **cli:** Fix cli version checking to work with monorepo ([#1510](https://github.com/cloudquery/cloudquery/issues/1510)) ([e0ddfcf](https://github.com/cloudquery/cloudquery/commit/e0ddfcf4f12f991db00c0706bbfc430dc0dd673d))

## [0.32.7](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.6...cli/v0.32.7) (2022-08-15)

### Bug Fixes

- **cli:** Update binary name ([#1304](https://github.com/cloudquery/cloudquery/issues/1304)) ([432c404](https://github.com/cloudquery/cloudquery/commit/432c40444fafc78f9d7a5c882c6203e3433c8627))

## [0.32.6](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.5...cli/v0.32.6) (2022-08-15)

### Bug Fixes

- **cli:** Keep old entrypoint name ([#1297](https://github.com/cloudquery/cloudquery/issues/1297)) ([bab4f39](https://github.com/cloudquery/cloudquery/commit/bab4f3972bc8853cdb3bf74cde8da01a399e182b))

## [0.32.5](https://github.com/cloudquery/cloudquery/compare/cli/v0.32.4...cli/v0.32.5) (2022-08-15)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#1272](https://github.com/cloudquery/cloudquery/issues/1272)) ([8546173](https://github.com/cloudquery/cloudquery/commit/85461731a03c9d2e5f84267e5eb7012226389a24))

## [0.32.4](https://github.com/cloudquery/cloudquery/compare/cli-v0.32.3...cli/v0.32.4) (2022-08-14)

### Bug Fixes

- **cli-docs:** Fix typos in docs/index ([#1256](https://github.com/cloudquery/cloudquery/issues/1256)) ([424da6d](https://github.com/cloudquery/cloudquery/commit/424da6da710f8294398f49e5bb7f837036f9c81a))
- **cli:** Don't pre-pend cq-provider to monorepo binary name ([#1264](https://github.com/cloudquery/cloudquery/issues/1264)) ([2b1e082](https://github.com/cloudquery/cloudquery/commit/2b1e0820ff731c5d8a6dca3036572aba674e34b5))

## [0.32.3](https://github.com/cloudquery/cloudquery/compare/v0.32.2...v0.32.3) (2022-08-14)

### Features

- Update CLI to work with monorepo release conventions ([#1224](https://github.com/cloudquery/cloudquery/issues/1224)) ([9cf254d](https://github.com/cloudquery/cloudquery/commit/9cf254d37131abfb110415a83b61b323f116d0e5))

## [0.32.2](https://github.com/cloudquery/cloudquery/compare/v0.32.1...v0.32.2) (2022-08-07)

### Bug Fixes

- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to 5cfe3b4 ([#1140](https://github.com/cloudquery/cloudquery/issues/1140)) ([a991dfc](https://github.com/cloudquery/cloudquery/commit/a991dfc0ee0ceb6f590fbd7e5ad30f9acea6d13a))
- **deps:** Update github.com/ProtonMail/go-crypto digest to d6ffb76 ([#1151](https://github.com/cloudquery/cloudquery/issues/1151)) ([2c8b814](https://github.com/cloudquery/cloudquery/commit/2c8b814cd8e7796955d867f63816871a4c72914e))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.5 ([#1138](https://github.com/cloudquery/cloudquery/issues/1138)) ([c9ed3ae](https://github.com/cloudquery/cloudquery/commit/c9ed3ae4927d39a7bb3479f54d6e495f62b01f11))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.6 ([#1149](https://github.com/cloudquery/cloudquery/issues/1149)) ([c06ab38](https://github.com/cloudquery/cloudquery/commit/c06ab38a145d0b18dc4debd7d999d35fe05f7c07))
- **deps:** Update module github.com/hashicorp/go-hclog to v1.2.2 ([#1152](https://github.com/cloudquery/cloudquery/issues/1152)) ([3801568](https://github.com/cloudquery/cloudquery/commit/380156800c66c8bc93b77460bf1c1ae46d4118c1))
- **deps:** Update module go to 1.18 ([#1153](https://github.com/cloudquery/cloudquery/issues/1153)) ([ccebe70](https://github.com/cloudquery/cloudquery/commit/ccebe70b5a1f82a1bfa8bff3a962369e08d0fc76))
- **deps:** Update tubone24/update_release digest to 2146f15 ([#1142](https://github.com/cloudquery/cloudquery/issues/1142)) ([7a75bbc](https://github.com/cloudquery/cloudquery/commit/7a75bbc9a7bde3917d754b4d573538a219168cc0))

## [0.32.1](https://github.com/cloudquery/cloudquery/compare/v0.32.0...v0.32.1) (2022-07-27)

### Bug Fixes

- **deps:** Update golang.org/x/term digest to a9ba230 ([#1141](https://github.com/cloudquery/cloudquery/issues/1141)) ([99b63fd](https://github.com/cloudquery/cloudquery/commit/99b63fd65777ef0b3d4c48ef3ac9413e40f87384))
- **deps:** Update module github.com/spf13/afero to v1.9.2 ([#1143](https://github.com/cloudquery/cloudquery/issues/1143)) ([a214d0c](https://github.com/cloudquery/cloudquery/commit/a214d0cb1ff1c7197b1b45646678447d324e2b2e))

## [0.32.0](https://github.com/cloudquery/cloudquery/compare/v0.31.9...v0.32.0) (2022-07-25)

### ⚠ BREAKING CHANGES

- Remove policy command and use pure SQL/psql (#1133)

### Features

- Remove policy command and use pure SQL/psql ([#1133](https://github.com/cloudquery/cloudquery/issues/1133)) ([f1e75ee](https://github.com/cloudquery/cloudquery/commit/f1e75ee40e361d7e69090574ca39eae5c1cc11e9))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.2 ([#1136](https://github.com/cloudquery/cloudquery/issues/1136)) ([8913946](https://github.com/cloudquery/cloudquery/commit/89139466e2345357e5b88e45c0601d11d6186be8))

## [0.31.9](https://github.com/cloudquery/cloudquery/compare/v0.31.8...v0.31.9) (2022-07-20)

### Features

- Add dsn_file connection option ([#1131](https://github.com/cloudquery/cloudquery/issues/1131)) ([114c4ed](https://github.com/cloudquery/cloudquery/commit/114c4edb7341702367cf25c011959a9a166dde70))

### Bug Fixes

- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to 152ecd2 ([#1115](https://github.com/cloudquery/cloudquery/issues/1115)) ([830ae91](https://github.com/cloudquery/cloudquery/commit/830ae91bb8ebc2f94b9c724a5c0c19b3e2435e6c))
- **deps:** Update github.com/ProtonMail/go-crypto digest to e85cedf ([#1123](https://github.com/cloudquery/cloudquery/issues/1123)) ([71421fc](https://github.com/cloudquery/cloudquery/commit/71421fcb17ef009ca538917b95074818ac7f4ab9))
- **deps:** Update module github.com/google/go-github/v35 to v45 ([#1075](https://github.com/cloudquery/cloudquery/issues/1075)) ([94d0f01](https://github.com/cloudquery/cloudquery/commit/94d0f01f8f9a3b831ccdc200063c4d74df7ef810))
- **deps:** Update module github.com/spf13/afero to v1.9.0 ([#1124](https://github.com/cloudquery/cloudquery/issues/1124)) ([fac9954](https://github.com/cloudquery/cloudquery/commit/fac9954defc4de54f7c166a072ab050e3aace0d0))
- **deps:** Update module github.com/spf13/viper to v1.12.0 ([#1021](https://github.com/cloudquery/cloudquery/issues/1021)) ([2fa73aa](https://github.com/cloudquery/cloudquery/commit/2fa73aa95ce35bb91a88d23ee4942eef293f9fc7))
- **deps:** Update module google.golang.org/grpc to v1.48.0 ([#1125](https://github.com/cloudquery/cloudquery/issues/1125)) ([c2e483b](https://github.com/cloudquery/cloudquery/commit/c2e483b279d38f6ddf241da908131b26b5917598))
- **deps:** Update tubone24/update_release digest to 87bc28c ([#1116](https://github.com/cloudquery/cloudquery/issues/1116)) ([0748135](https://github.com/cloudquery/cloudquery/commit/074813584eb263412a3fd9b4784409e90ac2bc7d))

## [0.31.8](https://github.com/cloudquery/cloudquery/compare/v0.31.7...v0.31.8) (2022-07-14)

### Features

- Skip Provider Update Check ([#1121](https://github.com/cloudquery/cloudquery/issues/1121)) ([1e3358d](https://github.com/cloudquery/cloudquery/commit/1e3358d30837998550f8ca1eed3b968b5b48aa8b))

### Bug Fixes

- Add GH actions to install src detection ([#1117](https://github.com/cloudquery/cloudquery/issues/1117)) ([398fe43](https://github.com/cloudquery/cloudquery/commit/398fe435cd03e0141e32d20f6a55247f4fc0ebd8))

## [0.31.7](https://github.com/cloudquery/cloudquery/compare/v0.31.6...v0.31.7) (2022-07-08)

### Features

- Add Binary and Installation Source tracking ([#1113](https://github.com/cloudquery/cloudquery/issues/1113)) ([006c54f](https://github.com/cloudquery/cloudquery/commit/006c54f2e21a6ead33fc51d9b08615bbcd02e573))
- Send telemetry about failed COPY FROMs ([#1050](https://github.com/cloudquery/cloudquery/issues/1050)) ([1da0478](https://github.com/cloudquery/cloudquery/commit/1da0478aa05ab409a74d2170bced181aff78e74e))

### Bug Fixes

- Add schema option to connection config ([#1112](https://github.com/cloudquery/cloudquery/issues/1112)) ([4e597e4](https://github.com/cloudquery/cloudquery/commit/4e597e4f4d298c8aaeeea47aefded375cc6467d7))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.5 ([#1114](https://github.com/cloudquery/cloudquery/issues/1114)) ([8e6cd35](https://github.com/cloudquery/cloudquery/commit/8e6cd35c5dd05fd901e75eedb6c8c0ebd5e7008d))

## [0.31.6](https://github.com/cloudquery/cloudquery/compare/v0.31.5...v0.31.6) (2022-07-06)

### Features

- Validate OOB config keys ([#1106](https://github.com/cloudquery/cloudquery/issues/1106)) ([0509354](https://github.com/cloudquery/cloudquery/commit/050935452308f00fc279558202324e4b5641ec93))

### Bug Fixes

- **deps:** Update module github.com/hashicorp/go-getter to v2 ([#1076](https://github.com/cloudquery/cloudquery/issues/1076)) ([31c67fa](https://github.com/cloudquery/cloudquery/commit/31c67faaf7b51619f26dfdd65bf3aebe2ec6e68f))
- Env Var Replacement ([#1108](https://github.com/cloudquery/cloudquery/issues/1108)) ([ead598f](https://github.com/cloudquery/cloudquery/commit/ead598f5634785f19f022e477c25022b9a01a9ca))
- Ulimit bug ([#1061](https://github.com/cloudquery/cloudquery/issues/1061)) ([38f1ce5](https://github.com/cloudquery/cloudquery/commit/38f1ce5d38723694613644fc07b403439ff1d5d0))

## [0.31.5](https://github.com/cloudquery/cloudquery/compare/v0.31.4...v0.31.5) (2022-07-04)

### Bug Fixes

- **cmd-version:** Embded commit and date ([#1099](https://github.com/cloudquery/cloudquery/issues/1099)) ([c6a6c23](https://github.com/cloudquery/cloudquery/commit/c6a6c23b841cefaa2e6f13528be7eab793347f43))

## [0.31.4](https://github.com/cloudquery/cloudquery/compare/v0.31.3...v0.31.4) (2022-07-04)

### Bug Fixes

- **cmd-doc:** Generate docs for completion command ([#1096](https://github.com/cloudquery/cloudquery/issues/1096)) ([ff3f940](https://github.com/cloudquery/cloudquery/commit/ff3f94099dd076b77e780f0f3040685bef79905d))
- **config:** Align provider version validation ([#915](https://github.com/cloudquery/cloudquery/issues/915)) ([d5de510](https://github.com/cloudquery/cloudquery/commit/d5de510181adc82c7b054f45d984ec234d6d45c3))

## [0.31.3](https://github.com/cloudquery/cloudquery/compare/v0.31.2...v0.31.3) (2022-07-04)

### Bug Fixes

- Classify MkDir fails as diag.USER ([#1091](https://github.com/cloudquery/cloudquery/issues/1091)) ([c52149e](https://github.com/cloudquery/cloudquery/commit/c52149ebcf2c4a77bab90a91ca580ca45d9fedb6))
- **cmd-providers:** Don't print help for sync command twice ([#1094](https://github.com/cloudquery/cloudquery/issues/1094)) ([fd2813f](https://github.com/cloudquery/cloudquery/commit/fd2813fe1dfe149bf82dccdb4d3974213eedda8e))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.4 ([#1093](https://github.com/cloudquery/cloudquery/issues/1093)) ([2ca124c](https://github.com/cloudquery/cloudquery/commit/2ca124ca180bbb3b9a2ae2deb3cad98f359b2c43))

## [0.31.2](https://github.com/cloudquery/cloudquery/compare/v0.31.1...v0.31.2) (2022-07-04)

### Bug Fixes

- **cmd-doc:** Don't add autogenerated tag ([#1087](https://github.com/cloudquery/cloudquery/issues/1087)) ([5d47f54](https://github.com/cloudquery/cloudquery/commit/5d47f548b194b4d8dfa0703fd3e5e65c79baf670))

## [0.31.1](https://github.com/cloudquery/cloudquery/compare/v0.31.0...v0.31.1) (2022-07-04)

### Bug Fixes

- **cmd:** Doc generation ([#1086](https://github.com/cloudquery/cloudquery/issues/1086)) ([4837053](https://github.com/cloudquery/cloudquery/commit/4837053c1ff41cb33403e8ac3b710b7dd29c7f1c))
- **deps:** Update module github.com/vbauerster/mpb/v6 to v7 ([#1077](https://github.com/cloudquery/cloudquery/issues/1077)) ([0e0d387](https://github.com/cloudquery/cloudquery/commit/0e0d3875f5e92d44d9e87efd351a1fa8b24138c5))

## [0.31.0](https://github.com/cloudquery/cloudquery/compare/v0.30.5...v0.31.0) (2022-07-04)

### ⚠ BREAKING CHANGES

- Remove support for HCL configuration ([#1044](https://github.com/cloudquery/cloudquery/issues/1044)) ([c6a9872](https://github.com/cloudquery/cloudquery/commit/c6a9872d5421c71fe84a23e014b4a315d10aa0ce)). For information on the new YAML configuration format [visit our docs](https://docs.cloudquery.io/docs/configuration/overview).

### Bug Fixes

- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to 3402ea6 ([#1069](https://github.com/cloudquery/cloudquery/issues/1069)) ([67ea5cb](https://github.com/cloudquery/cloudquery/commit/67ea5cb63aedb442081ed013168e11310266427d))
- **deps:** Update github.com/johannesboyne/gofakes3 digest to c3ac35d ([#1070](https://github.com/cloudquery/cloudquery/issues/1070)) ([156f757](https://github.com/cloudquery/cloudquery/commit/156f7573fcb89fad7dd90f237da84bb2d4d2697b))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.3 ([#1081](https://github.com/cloudquery/cloudquery/issues/1081)) ([e4522ae](https://github.com/cloudquery/cloudquery/commit/e4522ae328929d63bb139f936552d9f7f65b6d57))
- **deps:** Update module github.com/hashicorp/go-version to v1.6.0 ([#1071](https://github.com/cloudquery/cloudquery/issues/1071)) ([bb22b4a](https://github.com/cloudquery/cloudquery/commit/bb22b4a50265f554b1723c11f6be7dea4af120d0))
- **deps:** Update module github.com/stretchr/testify to v1.8.0 ([#1072](https://github.com/cloudquery/cloudquery/issues/1072)) ([23068a8](https://github.com/cloudquery/cloudquery/commit/23068a82f070b7388429e4528fefb5ccf1b48e1d))
- **deps:** Update module github.com/zclconf/go-cty to v1.10.0 ([#1073](https://github.com/cloudquery/cloudquery/issues/1073)) ([6350fa3](https://github.com/cloudquery/cloudquery/commit/6350fa3aaae7f8ef0ba10c2496d96bdafd2eb40d))

### Miscellaneous Chores

- Remove support for HCL configuration ([#1044](https://github.com/cloudquery/cloudquery/issues/1044)) ([c6a9872](https://github.com/cloudquery/cloudquery/commit/c6a9872d5421c71fe84a23e014b4a315d10aa0ce)). For information on the new YAML configuration format [visit our docs](https://docs.cloudquery.io/docs/configuration/overview).

## [0.30.5](https://github.com/cloudquery/cloudquery/compare/v0.30.4...v0.30.5) (2022-07-03)

### Features

- Remove custom completion command ([#1068](https://github.com/cloudquery/cloudquery/issues/1068)) ([c4ab6c5](https://github.com/cloudquery/cloudquery/commit/c4ab6c5d0ee36ccf77529d7120fe878af6c9b3db))

### Bug Fixes

- **deps:** Update module github.com/spf13/cobra to v1.5.0 ([#1019](https://github.com/cloudquery/cloudquery/issues/1019)) ([0facde6](https://github.com/cloudquery/cloudquery/commit/0facde6ac1e12bfd5355a777991d396718e16598))

## [0.30.4](https://github.com/cloudquery/cloudquery/compare/v0.30.3...v0.30.4) (2022-07-03)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.2 ([#1064](https://github.com/cloudquery/cloudquery/issues/1064)) ([2429095](https://github.com/cloudquery/cloudquery/commit/2429095c390a68a8953aad80f024e605a35755a3))

## [0.30.3](https://github.com/cloudquery/cloudquery/compare/v0.30.2...v0.30.3) (2022-07-03)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.1 ([#1056](https://github.com/cloudquery/cloudquery/issues/1056)) ([f216d90](https://github.com/cloudquery/cloudquery/commit/f216d9012bfe45436c1821455149af1a4b27d618))
- Update 'policy run' help message ([#1037](https://github.com/cloudquery/cloudquery/issues/1037)) ([2235e82](https://github.com/cloudquery/cloudquery/commit/2235e826c956ea1873d18f3d4155f6e08b5edfc2))

## [0.30.2](https://github.com/cloudquery/cloudquery/compare/v0.30.1...v0.30.2) (2022-06-30)

### Bug Fixes

- Don't print duplicate error during init ([#984](https://github.com/cloudquery/cloudquery/issues/984)) ([79d1b54](https://github.com/cloudquery/cloudquery/commit/79d1b540e26b04da0581b24be2aea268e51ee719))

## [0.30.1](https://github.com/cloudquery/cloudquery/compare/v0.30.0...v0.30.1) (2022-06-30)

### Features

- Track config format ([#1040](https://github.com/cloudquery/cloudquery/issues/1040)) ([3a3ad17](https://github.com/cloudquery/cloudquery/commit/3a3ad17dcdd6a5a7d5ac1df3128617140e07b6e8))

### Bug Fixes

- Doc command to generate id meta without space ([#1046](https://github.com/cloudquery/cloudquery/issues/1046)) ([be12ebf](https://github.com/cloudquery/cloudquery/commit/be12ebf97b5bee54c985b8bece00d62757253e2f))
- **docs:** Add code of conduct contact details & commit message guidelines to CONTRIBUTING.md ([#1043](https://github.com/cloudquery/cloudquery/issues/1043)) ([b304bbc](https://github.com/cloudquery/cloudquery/commit/b304bbcaba0024705a4bdb703a7c7a937527d4a0))
- Support env var substitution in YAML config ([#1041](https://github.com/cloudquery/cloudquery/issues/1041)) ([3bc8014](https://github.com/cloudquery/cloudquery/commit/3bc8014c34655d8de7a8d7210be167e7e14f772b))

## [0.30.0](https://github.com/cloudquery/cloudquery/compare/v0.29.0...v0.30.0) (2022-06-29)

### ⚠ BREAKING CHANGES

- Remove 'policy download' feature (#1033)

### Features

- Generate metadata for CLI docs ([#1034](https://github.com/cloudquery/cloudquery/issues/1034)) ([f7ae08a](https://github.com/cloudquery/cloudquery/commit/f7ae08a4803adc69a11e26eddd8bfe5db5648f30))
- Remove 'policy download' feature ([#1033](https://github.com/cloudquery/cloudquery/issues/1033)) ([e376aed](https://github.com/cloudquery/cloudquery/commit/e376aed83c625b92be9d5f44e5bd0b902b50ca9f))

## [0.29.0](https://github.com/cloudquery/cloudquery/compare/v0.28.3...v0.29.0) (2022-06-27)

### ⚠ BREAKING CHANGES

- Remove support for cq init HCL (#993)
- Rename default config to be cloudquery.yml (#1030)

### Bug Fixes

- Configure Provider Yml ([#1001](https://github.com/cloudquery/cloudquery/issues/1001)) ([23b6695](https://github.com/cloudquery/cloudquery/commit/23b66956148885bb5e9b532a80c0136568868eda))
- **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.5 ([#1027](https://github.com/cloudquery/cloudquery/issues/1027)) ([57db8db](https://github.com/cloudquery/cloudquery/commit/57db8db759e24b6563f8a4828b5e093320f897d7))
- Don't decrease ulimit. Change ulimit logs from 'debug' to 'info' ([#1028](https://github.com/cloudquery/cloudquery/issues/1028)) ([079cef1](https://github.com/cloudquery/cloudquery/commit/079cef159718face4d4abb7d40e7353d5c32c267))

### Miscellaneous Chores

- Remove support for cq init HCL ([#993](https://github.com/cloudquery/cloudquery/issues/993)) ([0504e8](https://github.com/cloudquery/cloudquery/commit/0504e849482fe8d76ef52ee9432186509ec254e8))
- Rename default config to be cloudquery.yml ([#1030](https://github.com/cloudquery/cloudquery/issues/1030)) ([c15a9d](https://github.com/cloudquery/cloudquery/commit/c15a9d72ee5b56b6e75951b9c0985a1f3ac04d2e))

## [0.28.3](https://github.com/cloudquery/cloudquery/compare/v0.28.2...v0.28.3) (2022-06-27)

### Features

- More verbose summary ([#996](https://github.com/cloudquery/cloudquery/pull/996)) ([d5d6ac6](https://github.com/cloudquery/cloudquery/commit/d5d6ac613a9abd11c99d4309a151af079704c4bb))

### Bug Fixes

- **deps:** fix(deps): Update module github.com/georgysavva/scany to v0.3.0 ([#1013](https://github.com/cloudquery/cloudquery/issues/1013)) ([a44a3a7](https://github.com/cloudquery/cloudquery/commit/a44a3a7ff99d16fd3792cb15e1f1f403de717fbd))

## [0.28.2](https://github.com/cloudquery/cloudquery/compare/v0.28.1...v0.28.2) (2022-06-27)

### Bug Fixes

- **deps:** fix(deps): Update github.com/hairyhenderson/go-fsimpl digest to d4f0b5a ([#1011](https://github.com/cloudquery/cloudquery/issues/1011)) ([2db5c30](https://github.com/cloudquery/cloudquery/commit/2db5c30c626a273dcace1e43e524cab9bf6f1200))
- **deps:** fix(deps): Update github.com/ProtonMail/go-crypto digest to 5afb4c2 ([#1010](https://github.com/cloudquery/cloudquery/issues/1010)) ([19ba992](https://github.com/cloudquery/cloudquery/commit/19ba9929e49567d97c7829ae835bb461158cc0ce))
- **deps:** fix(deps): Update module github.com/google/go-github/v35 to v35.3.0 ([#1014](https://github.com/cloudquery/cloudquery/issues/1014)) ([1c35b45](https://github.com/cloudquery/cloudquery/commit/1c35b454419c2d5e41545ec85b6c3073a9d40261))
- **deps:** fix(deps): Update module github.com/hashicorp/hcl/v2 to v2.13.0 ([#1015](https://github.com/cloudquery/cloudquery/issues/1015)) ([29e5ed5](https://github.com/cloudquery/cloudquery/commit/29e5ed5bd7fe4f02dfb88a9af4c22be5d961a7c9))
- **deps:** fix(deps): Update module github.com/rs/zerolog to v1.27.0 ([#1016](https://github.com/cloudquery/cloudquery/issues/1016)) ([3643d3d](https://github.com/cloudquery/cloudquery/commit/3643d3dda0439605680e4f39209fdc3e6e122083))
- **deps:** fix(deps): Update module github.com/spf13/afero to v1.8.2 ([#1017](https://github.com/cloudquery/cloudquery/issues/1017)) ([44e8e9b](https://github.com/cloudquery/cloudquery/commit/44e8e9bc27e4420f48069bb1994be218860f1419))
- **deps:** fix(deps): Update module github.com/spf13/cast to v1.5.0 ([#1018](https://github.com/cloudquery/cloudquery/issues/1018)) ([5fb9cc6](https://github.com/cloudquery/cloudquery/commit/5fb9cc6d8e6e716fbe28f25ce3a634a825204f5d))
- **deps:** fix(deps): Update module github.com/stretchr/testify to v1.7.5 ([#1012](https://github.com/cloudquery/cloudquery/issues/1012)) ([a96c37b](https://github.com/cloudquery/cloudquery/commit/a96c37b44542862183b006caedaac71fdd18f777))

## [0.28.1](https://github.com/cloudquery/cloudquery/compare/v0.28.0...v0.28.1) (2022-06-26)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.3 ([#1008](https://github.com/cloudquery/cloudquery/issues/1008)) ([2eb7ecf](https://github.com/cloudquery/cloudquery/commit/2eb7ecf22c625ac645d539d6e767b3a2c783db17))
- Policy download test ([#1009](https://github.com/cloudquery/cloudquery/issues/1009)) ([a3312e9](https://github.com/cloudquery/cloudquery/commit/a3312e9af6cc114557adf273e303beb19d7f2cad))
- **providers-sync:** Improve error message ([#1006](https://github.com/cloudquery/cloudquery/issues/1006)) ([a38b443](https://github.com/cloudquery/cloudquery/commit/a38b4430d515dcff7a5d55029afa7bf4186b26a7))

## [0.28.0](https://github.com/cloudquery/cloudquery/compare/v0.27.3...v0.28.0) (2022-06-26)

### ⚠ BREAKING CHANGES

- Merge upgrade/downgrade provider command to sync command (#973)

### Features

- Merge upgrade/downgrade provider command to sync command ([#973](https://github.com/cloudquery/cloudquery/issues/973)) ([d76255a](https://github.com/cloudquery/cloudquery/commit/d76255a4c441deb9539c7616fbcbbb428330b3a9))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.2 ([#999](https://github.com/cloudquery/cloudquery/issues/999)) ([a735082](https://github.com/cloudquery/cloudquery/commit/a735082c2b42606d53dcaa00222a2fd2bd611d68))
- Handle Diags ([#997](https://github.com/cloudquery/cloudquery/issues/997)) ([ca41b84](https://github.com/cloudquery/cloudquery/commit/ca41b84812680f7efcf8bcee46a135219603abad))
- Show identifiers values correctly ([#992](https://github.com/cloudquery/cloudquery/issues/992)) ([1487893](https://github.com/cloudquery/cloudquery/commit/148789372cc662da8367427bfa9181d672b3554e))
- Work around the panic in createOutputTable ([#995](https://github.com/cloudquery/cloudquery/issues/995)) ([3e3c957](https://github.com/cloudquery/cloudquery/commit/3e3c95760d449c010b2afa310294bb3ecbaf09f1))

## [0.27.3](https://github.com/cloudquery/cloudquery/compare/v0.27.2...v0.27.3) (2022-06-23)

### Features

- Add doc command to generation doc markdown ([#989](https://github.com/cloudquery/cloudquery/issues/989)) ([3572560](https://github.com/cloudquery/cloudquery/commit/357256056b5784f386c516f5179aae857c6da921))

## [0.27.2](https://github.com/cloudquery/cloudquery/compare/v0.27.1...v0.27.2) (2022-06-23)

### Bug Fixes

- Disallow running of db-using policy operations without config ([#985](https://github.com/cloudquery/cloudquery/issues/985)) ([e4334c2](https://github.com/cloudquery/cloudquery/commit/e4334c2d08fbb43884d9e3d662d58a5b532694f0))
- More panics in policy commands ([#988](https://github.com/cloudquery/cloudquery/issues/988)) ([0fd0867](https://github.com/cloudquery/cloudquery/commit/0fd08677f60dac4ef91b65c5db5ce34a61d35de2))

## [0.27.1](https://github.com/cloudquery/cloudquery/compare/v0.27.0...v0.27.1) (2022-06-22)

### Bug Fixes

- Don't print duplicate errors ([#955](https://github.com/cloudquery/cloudquery/issues/955)) ([18c6b6c](https://github.com/cloudquery/cloudquery/commit/18c6b6c0533533916b444f262747e53dbcec4246))
- Remove unused console print ([#959](https://github.com/cloudquery/cloudquery/issues/959)) ([f1896f1](https://github.com/cloudquery/cloudquery/commit/f1896f160ede732bfae8bc15a1eeb9689eb15a6b))

## [0.27.0](https://github.com/cloudquery/cloudquery/compare/v0.26.4...v0.27.0) (2022-06-22)

### ⚠ BREAKING CHANGES

- Remove drift (#887)
- Remove `plugin_directory` and `policy_directory` from `cloudquery` configuration block. Please use the `--data-dir` CLI flag instead or remove if using default values (#887)

### Features

- Remove drift ([#887](https://github.com/cloudquery/cloudquery/issues/887)) ([3d387bd](https://github.com/cloudquery/cloudquery/commit/3d387bda0ed8afcdb0b32b5ec1ae2d0e9c279e5e))
- Add YAML configuration support ([#887](https://github.com/cloudquery/cloudquery/issues/887)) ([3d387bd](https://github.com/cloudquery/cloudquery/commit/3d387bda0ed8afcdb0b32b5ec1ae2d0e9c279e5e))
- Remove `plugin_directory` and `policy_directory` from `cloudquery` configuration block. Please use the `--data-dir` CLI flag instead or remove if using default values (#887)

### Bug Fixes

- Default alias to name ([#966](https://github.com/cloudquery/cloudquery/issues/966)) ([706447c](https://github.com/cloudquery/cloudquery/commit/706447ce55bedc4f30260ac49100a7715e300ba3))
- Support modules tag in config ([#965](https://github.com/cloudquery/cloudquery/issues/965)) ([379344f](https://github.com/cloudquery/cloudquery/commit/379344f818768fde05c3fd61e976283678ed09e3))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.1 ([#972](https://github.com/cloudquery/cloudquery/issues/972)) ([1f871e9](https://github.com/cloudquery/cloudquery/commit/1f871e9191db48bddc7ee7e69ba5ef01a2b284d9))
- Request correct config format (YAML) from provider ([#968](https://github.com/cloudquery/cloudquery/issues/968)) ([999b68d](https://github.com/cloudquery/cloudquery/commit/999b68da85cb71dd3307761cd4051a4357197afe))

## [0.27.0-rc3](https://github.com/cloudquery/cloudquery/compare/v0.27.0-rc2...v0.27.0-rc3) (2022-06-22)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.1 ([#972](https://github.com/cloudquery/cloudquery/issues/972)) ([1f871e9](https://github.com/cloudquery/cloudquery/commit/1f871e9191db48bddc7ee7e69ba5ef01a2b284d9))
- Request correct config format (YAML) from provider ([#968](https://github.com/cloudquery/cloudquery/issues/968)) ([999b68d](https://github.com/cloudquery/cloudquery/commit/999b68da85cb71dd3307761cd4051a4357197afe))

### Miscellaneous Chores

- Release 0.27.0-rc3 ([#970](https://github.com/cloudquery/cloudquery/issues/970)) ([a0f432a](https://github.com/cloudquery/cloudquery/commit/a0f432aa29281c8c5d3d4388ef216b6100bca034))

## [0.27.0-rc2](https://github.com/cloudquery/cloudquery/compare/v0.27.0-rc1...v0.27.0-rc2) (2022-06-21)

### Bug Fixes

- Default alias to name ([#966](https://github.com/cloudquery/cloudquery/issues/966)) ([706447c](https://github.com/cloudquery/cloudquery/commit/706447ce55bedc4f30260ac49100a7715e300ba3))
- Support modules tag in config ([#965](https://github.com/cloudquery/cloudquery/issues/965)) ([379344f](https://github.com/cloudquery/cloudquery/commit/379344f818768fde05c3fd61e976283678ed09e3))

### Miscellaneous Chores

- Release 0.27.0-rc2 ([#967](https://github.com/cloudquery/cloudquery/issues/967)) ([cd37a11](https://github.com/cloudquery/cloudquery/commit/cd37a119581da8f6dbe138562c6897e3e5db5b6a))

## [0.27.0-rc1](https://github.com/cloudquery/cloudquery/compare/v0.26.4...v0.27.0-rc1) (2022-06-21)

### ⚠ BREAKING CHANGES

- Remove drift (#887)

### Features

- Remove drift ([#887](https://github.com/cloudquery/cloudquery/issues/887)) ([3d387bd](https://github.com/cloudquery/cloudquery/commit/3d387bda0ed8afcdb0b32b5ec1ae2d0e9c279e5e))

### Miscellaneous Chores

- Release 0.27.0-rc1 ([#962](https://github.com/cloudquery/cloudquery/issues/962)) ([3a2ec6d](https://github.com/cloudquery/cloudquery/commit/3a2ec6d256d488fe06c47a0602978213859555ed))

## [0.26.4](https://github.com/cloudquery/cloudquery/compare/v0.26.3...v0.26.4) (2022-06-21)

### Bug Fixes

- Silence usage on command errors ([#956](https://github.com/cloudquery/cloudquery/issues/956)) ([474473d](https://github.com/cloudquery/cloudquery/commit/474473d29947fae4bef151ae03c275bb0095f5d1))

## [0.26.3](https://github.com/cloudquery/cloudquery/compare/v0.26.2...v0.26.3) (2022-06-20)

### Bug Fixes

- **deps:** Update github.com/hairyhenderson/go-fsimpl digest to 3a8e791 ([#923](https://github.com/cloudquery/cloudquery/issues/923)) ([974933f](https://github.com/cloudquery/cloudquery/commit/974933f6f3b33499237d50cde0796bc9b208ad96))
- **deps:** Update module github.com/aws/aws-sdk-go to v1.44.37 ([#939](https://github.com/cloudquery/cloudquery/issues/939)) ([cc0670d](https://github.com/cloudquery/cloudquery/commit/cc0670dce30e6242983788cb5d6b335a43178e24))
- Revert "fix(deps): Update github.com/hairyhenderson/go-fsimpl digest to 3a8e791 ([#923](https://github.com/cloudquery/cloudquery/issues/923))" ([#950](https://github.com/cloudquery/cloudquery/issues/950)) ([75f870d](https://github.com/cloudquery/cloudquery/commit/75f870d971977b760b6a0224de1beb7d90fc6012))

## [0.26.2](https://github.com/cloudquery/cloudquery/compare/v0.26.1...v0.26.2) (2022-06-20)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.4 ([#947](https://github.com/cloudquery/cloudquery/issues/947)) ([2823123](https://github.com/cloudquery/cloudquery/commit/2823123db5ce8e3e12b33c2e04362043faaae330))

## [0.26.1](https://github.com/cloudquery/cloudquery/compare/v0.26.0...v0.26.1) (2022-06-20)

### Features

- Remove handle_command wrapper function ([#918](https://github.com/cloudquery/cloudquery/issues/918)) ([e75a5f9](https://github.com/cloudquery/cloudquery/commit/e75a5f936d2d35f541c6d5c457dc7b527d78e4fb))

### Bug Fixes

- **deps:** Update github.com/jackc/pgerrcode digest to 469b46a ([#924](https://github.com/cloudquery/cloudquery/issues/924)) ([da0c201](https://github.com/cloudquery/cloudquery/commit/da0c20174f19f938371431e3d6f52cd098779336))
- **deps:** Update github.com/johannesboyne/gofakes3 digest to 83a58ec ([#925](https://github.com/cloudquery/cloudquery/issues/925)) ([01c8c92](https://github.com/cloudquery/cloudquery/commit/01c8c9233ecd1159f6767dc4f6f1808ec507e194))
- Progress bar diag count ([#945](https://github.com/cloudquery/cloudquery/issues/945)) ([e3549b3](https://github.com/cloudquery/cloudquery/commit/e3549b392fc43e79d392d5b9e160c9d0dc968c1d)), closes [#883](https://github.com/cloudquery/cloudquery/issues/883)

## [0.26.0](https://github.com/cloudquery/cloudquery/compare/v0.25.7...v0.26.0) (2022-06-20)

### ⚠ BREAKING CHANGES

- Remove fail-on-error (#916)

### Features

- Remove fail-on-error ([#916](https://github.com/cloudquery/cloudquery/issues/916)) ([00fd817](https://github.com/cloudquery/cloudquery/commit/00fd817ef92f81f065c5d8f3c07a431632d8efd8))

### Bug Fixes

- **deps:** Update github.com/ProtonMail/go-crypto digest to 88bb529 ([#922](https://github.com/cloudquery/cloudquery/issues/922)) ([20ce36e](https://github.com/cloudquery/cloudquery/commit/20ce36e5a2d928259aad7728e20c7fbc94da3b8a))

## [0.25.7](https://github.com/cloudquery/cloudquery/compare/v0.25.6...v0.25.7) (2022-06-20)

### Bug Fixes

- **deps:** Update golang.org/x/term digest to 065cf7b ([#926](https://github.com/cloudquery/cloudquery/issues/926)) ([81df0d0](https://github.com/cloudquery/cloudquery/commit/81df0d0621d622d519d7466a843b1a6717b8c066))
- **deps:** Update module github.com/doug-martin/goqu/v9 to v9.18.0 ([#940](https://github.com/cloudquery/cloudquery/issues/940)) ([90e9574](https://github.com/cloudquery/cloudquery/commit/90e95742630db559d24f1f80e9106c424609b7f8))
- **deps:** Update module github.com/golang-migrate/migrate/v4 to v4.15.2 ([#928](https://github.com/cloudquery/cloudquery/issues/928)) ([9dde364](https://github.com/cloudquery/cloudquery/commit/9dde364291bbff3d2557addaaa566fdef4d76a67))
- **deps:** Update module github.com/google/go-cmp to v0.5.8 ([#929](https://github.com/cloudquery/cloudquery/issues/929)) ([ed8d1ea](https://github.com/cloudquery/cloudquery/commit/ed8d1ea7476aecde5c42c5578c59c07d607554e6))
- **deps:** Update module github.com/hashicorp/go-getter to v1.6.2 ([#930](https://github.com/cloudquery/cloudquery/issues/930)) ([a04bd1e](https://github.com/cloudquery/cloudquery/commit/a04bd1e1a174a39e6afdb06cd463e106b78b8c96))
- **deps:** Update module github.com/hashicorp/go-plugin to v1.4.4 ([#931](https://github.com/cloudquery/cloudquery/issues/931)) ([b6a0c09](https://github.com/cloudquery/cloudquery/commit/b6a0c09ea23199a3a210c849478e15c88fff2073))
- **deps:** Update module github.com/jackc/pgconn to v1.12.1 ([#933](https://github.com/cloudquery/cloudquery/issues/933)) ([7939aa9](https://github.com/cloudquery/cloudquery/commit/7939aa9269deefe556cdf87c2d322ea9cb85950f))
- **deps:** Update module github.com/jackc/pgx/v4 to v4.16.1 ([#934](https://github.com/cloudquery/cloudquery/issues/934)) ([d76ed32](https://github.com/cloudquery/cloudquery/commit/d76ed32a51a03febba7b5b48a9f508e7870b96f6))
- **deps:** Update module github.com/lib/pq to v1.10.6 ([#935](https://github.com/cloudquery/cloudquery/issues/935)) ([66d4d51](https://github.com/cloudquery/cloudquery/commit/66d4d5149e9b98fe1f844b31f7217a542ed576ff))
- **deps:** Update module github.com/stretchr/testify to v1.7.2 ([#936](https://github.com/cloudquery/cloudquery/issues/936)) ([8ece5bb](https://github.com/cloudquery/cloudquery/commit/8ece5bb91769a91646d73060a925865a623ff26a))
- **deps:** Update module github.com/vbauerster/mpb/v6 to v6.0.4 ([#938](https://github.com/cloudquery/cloudquery/issues/938)) ([5481d89](https://github.com/cloudquery/cloudquery/commit/5481d89bf74ef5ecb699608de616d12fade30493))

## [0.25.6](https://github.com/cloudquery/cloudquery/compare/v0.25.5...v0.25.6) (2022-06-19)

### Bug Fixes

- **console:** Don't print [@latest](https://github.com/latest) when fetching ([#919](https://github.com/cloudquery/cloudquery/issues/919)) ([3cf9789](https://github.com/cloudquery/cloudquery/commit/3cf97899ef6ede096542964d8c7081a62e97efe7))

## [0.25.5](https://github.com/cloudquery/cloudquery/compare/v0.25.4...v0.25.5) (2022-06-19)

### Bug Fixes

- **command-options:** Hide internal debug flags ([#906](https://github.com/cloudquery/cloudquery/issues/906)) ([67be21e](https://github.com/cloudquery/cloudquery/commit/67be21e8bf51cd0061d42a2300ea8a9bd8e199db))

## [0.25.4](https://github.com/cloudquery/cloudquery/compare/v0.25.3...v0.25.4) (2022-06-15)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.2 ([#902](https://github.com/cloudquery/cloudquery/issues/902)) ([573c4f2](https://github.com/cloudquery/cloudquery/commit/573c4f2dfdc9b351e7c8a0a0d599b83aa5427d87))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.3 ([#905](https://github.com/cloudquery/cloudquery/issues/905)) ([1d1a321](https://github.com/cloudquery/cloudquery/commit/1d1a3210d9e24dc25bc86c1add7a2380538ccafa))

## [0.25.3](https://github.com/cloudquery/cloudquery/compare/v0.25.2...v0.25.3) (2022-06-15)

### Bug Fixes

- Classify "failed to open file" as USER diags ([#891](https://github.com/cloudquery/cloudquery/issues/891)) ([9e8c489](https://github.com/cloudquery/cloudquery/commit/9e8c48942a01e1836ccdec1727d5a179b9e5c018))
- Remove plugin-dir and data-dir CLI flags ([#899](https://github.com/cloudquery/cloudquery/issues/899)) ([50afe25](https://github.com/cloudquery/cloudquery/commit/50afe253e7a1afd389cdd5d60d6a121c07a76dee))

## [0.25.2](https://github.com/cloudquery/cloudquery/compare/v0.25.1...v0.25.2) (2022-06-14)

### Bug Fixes

- Add missing space when printing provider update message ([#897](https://github.com/cloudquery/cloudquery/issues/897)) ([9694baf](https://github.com/cloudquery/cloudquery/commit/9694baf2dcaba9a67967c964295f8ea2012cf689))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.1 ([#892](https://github.com/cloudquery/cloudquery/issues/892)) ([eaecdfe](https://github.com/cloudquery/cloudquery/commit/eaecdfedcfee9d41a99b7ab78e7985621700e805))
- Index out of range in `policy snapshot` ([#894](https://github.com/cloudquery/cloudquery/issues/894)) ([5d0a46d](https://github.com/cloudquery/cloudquery/commit/5d0a46dac1fd48033cc9cdaec2a98328261c1acb))

## [0.25.1](https://github.com/cloudquery/cloudquery/compare/v0.25.0...v0.25.1) (2022-06-13)

### Features

- Add clickable links support ([#882](https://github.com/cloudquery/cloudquery/issues/882)) ([366d585](https://github.com/cloudquery/cloudquery/commit/366d585d34947c1d3cbcb7864bad281798d86dbe))

### Bug Fixes

- **deps:** Bump github.com/hashicorp/go-getter from 1.5.11 to 1.6.1 ([#881](https://github.com/cloudquery/cloudquery/issues/881)) ([1dddb95](https://github.com/cloudquery/cloudquery/commit/1dddb952ec51d6fd17fc5f2c8190c8b23cc76a93))
- Improve invalid provider version error ([#879](https://github.com/cloudquery/cloudquery/issues/879)) ([6a03444](https://github.com/cloudquery/cloudquery/commit/6a0344435d9a7611100e62e1c03fbb01b377bcdf))

## [0.25.0](https://github.com/cloudquery/cloudquery/compare/v0.24.11...v0.25.0) (2022-06-09)

### ⚠ BREAKING CHANGES

- Disable policies in config file, Disallow running more than one policy (#841)

### Features

- Sentry: Report provider version as release ([#874](https://github.com/cloudquery/cloudquery/issues/874)) ([349d4cf](https://github.com/cloudquery/cloudquery/commit/349d4cfc16a7f93cc23346fff9df11dc664a3705))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.0 ([#875](https://github.com/cloudquery/cloudquery/issues/875)) ([890ac09](https://github.com/cloudquery/cloudquery/commit/890ac09931db55397fcd4c49a1ee8578465cef16))
- Github detector should not use firebase ([#871](https://github.com/cloudquery/cloudquery/issues/871)) ([2175524](https://github.com/cloudquery/cloudquery/commit/2175524107f41b74d19fc45a33b3f0504f839220))
- Purge command failure ([#869](https://github.com/cloudquery/cloudquery/issues/869)) ([9010f18](https://github.com/cloudquery/cloudquery/commit/9010f1808f59118f6d8803cf70b5dcb3b670d982))
- Redact errors again ([#867](https://github.com/cloudquery/cloudquery/issues/867)) ([9e9f81a](https://github.com/cloudquery/cloudquery/commit/9e9f81a6843f8811607b904ba28324fc0afcd7e2))
- Small fix in deprecation warning for "policy in config" ([#873](https://github.com/cloudquery/cloudquery/issues/873)) ([11d27ab](https://github.com/cloudquery/cloudquery/commit/11d27aba7449af5d6d5720f9a6f6e7a5349e25fc))

### breaking

- Disable policies in config file, Disallow running more than one policy ([#841](https://github.com/cloudquery/cloudquery/issues/841)) ([f97de5d](https://github.com/cloudquery/cloudquery/commit/f97de5df4b45a16a7ad5952cebcc3fe750e66bdf))

## [0.24.11](https://github.com/cloudquery/cloudquery/compare/v0.24.10...v0.24.11) (2022-06-07)

### Features

- Improve policy analytics ([#862](https://github.com/cloudquery/cloudquery/issues/862)) ([7789f82](https://github.com/cloudquery/cloudquery/commit/7789f82563e644e1e7b87f8c2d2f1cc6c4ec47b1))

### Bug Fixes

- **console:** Output to console when console log is enabled in non termial envs ([#853](https://github.com/cloudquery/cloudquery/issues/853)) ([83731e5](https://github.com/cloudquery/cloudquery/commit/83731e5e56c57839d7abb47ba150eca0b8626083))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.10 ([#865](https://github.com/cloudquery/cloudquery/issues/865)) ([13c0b7e](https://github.com/cloudquery/cloudquery/commit/13c0b7e1eb04eead871032dc2cc7c3c7a6712061))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.11 ([#866](https://github.com/cloudquery/cloudquery/issues/866)) ([ed50a3c](https://github.com/cloudquery/cloudquery/commit/ed50a3ca222150fae4ff401329a91ea7683844fa))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.9 ([#863](https://github.com/cloudquery/cloudquery/issues/863)) ([eb998ad](https://github.com/cloudquery/cloudquery/commit/eb998ada94e9f56134d8c57b49d246b9e6b54536))
- Remove deprecated 'policy_directory' and 'plugin_directory' from hcl file ([#855](https://github.com/cloudquery/cloudquery/issues/855)) ([cd3eb90](https://github.com/cloudquery/cloudquery/commit/cd3eb90806f2a73ffba20f9b5f8752fcc429db07))

## [0.24.10](https://github.com/cloudquery/cloudquery/compare/v0.24.9...v0.24.10) (2022-06-07)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.8 ([#860](https://github.com/cloudquery/cloudquery/issues/860)) ([9ec43f5](https://github.com/cloudquery/cloudquery/commit/9ec43f566371f8e081f1016aaa4bab8746b6171d))

## [0.24.9](https://github.com/cloudquery/cloudquery/compare/v0.24.8...v0.24.9) (2022-06-06)

### Bug Fixes

- Improve missing git binary message ([#846](https://github.com/cloudquery/cloudquery/issues/846)) ([28abc99](https://github.com/cloudquery/cloudquery/commit/28abc9917c7b392afb7dc5edfd844c3013f27588))

## [0.24.8](https://github.com/cloudquery/cloudquery/compare/v0.24.7...v0.24.8) (2022-06-02)

### Features

- Change minimum postgres version to 10 ([#838](https://github.com/cloudquery/cloudquery/issues/838)) ([33c3c2c](https://github.com/cloudquery/cloudquery/commit/33c3c2c4d0df9dde2f29725b98dc8ca9a7bf1505))

### Bug Fixes

- Disable Telemetry with special CQ team cookie ([#832](https://github.com/cloudquery/cloudquery/issues/832)) ([550b17b](https://github.com/cloudquery/cloudquery/commit/550b17be108adb2f0393507db4a0ee1c46a0bdac))

## [0.24.7](https://github.com/cloudquery/cloudquery/compare/v0.24.6...v0.24.7) (2022-06-02)

### Bug Fixes

- Policy Output Table ([#830](https://github.com/cloudquery/cloudquery/issues/830)) ([f5f3dc9](https://github.com/cloudquery/cloudquery/commit/f5f3dc912789630b268306f32b42e0e601a9422a))

### [0.24.6](https://github.com/cloudquery/cloudquery/compare/v0.24.5...v0.24.6) (2022-06-01)

### Features

- Add goroutine count to sentry ([#825](https://github.com/cloudquery/cloudquery/issues/825)) ([d4d56c2](https://github.com/cloudquery/cloudquery/commit/d4d56c2011e3f5665feab629ce9ac2d3e46990f9))

### Bug Fixes

- Bad condition db version analytics ([#829](https://github.com/cloudquery/cloudquery/issues/829)) ([6cd8d6d](https://github.com/cloudquery/cloudquery/commit/6cd8d6daeb92a8fc3930d166ebcd3906b0bc7db5))
- Changed level of bad config diagnostics ([#822](https://github.com/cloudquery/cloudquery/issues/822)) ([67b0bd3](https://github.com/cloudquery/cloudquery/commit/67b0bd3ef32401443c20c74b76968a14ed6ac62a))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.7 ([#826](https://github.com/cloudquery/cloudquery/issues/826)) ([f075726](https://github.com/cloudquery/cloudquery/commit/f075726a99122509b10f21ffa9b7384787f28edb))
- Don't fail on PG_TRGM missing ([#821](https://github.com/cloudquery/cloudquery/issues/821)) ([a0672b5](https://github.com/cloudquery/cloudquery/commit/a0672b5411a3858e9c69e590a387b6bd96730fdd)), closes [#816](https://github.com/cloudquery/cloudquery/issues/816)
- Incorrect diagnostics summary count ([#823](https://github.com/cloudquery/cloudquery/issues/823)) ([58760b4](https://github.com/cloudquery/cloudquery/commit/58760b451527c14cb61877c3d75c0aa172b99782))

### [0.24.5](https://github.com/cloudquery/cloudquery/compare/v0.24.4...v0.24.5) (2022-06-01)

### Features

- Add ulimit info to sentry ([#819](https://github.com/cloudquery/cloudquery/issues/819)) ([540878b](https://github.com/cloudquery/cloudquery/commit/540878b5cbc12f07fadc530099fb008f84142569))

### [0.24.4](https://github.com/cloudquery/cloudquery/compare/v0.24.3...v0.24.4) (2022-06-01)

### Features

- Update Error message for Failed Policy Execution ([#814](https://github.com/cloudquery/cloudquery/issues/814)) ([8b9a7b9](https://github.com/cloudquery/cloudquery/commit/8b9a7b9725c66dd6099e44dd83f908905896c459))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.6 ([#818](https://github.com/cloudquery/cloudquery/issues/818)) ([ae48d5f](https://github.com/cloudquery/cloudquery/commit/ae48d5ffdf1c5aa5950f2326fa333bacc7bdb65d))

### [0.24.3](https://github.com/cloudquery/cloudquery/compare/v0.24.2...v0.24.3) (2022-05-31)

### Features

- Add database info metrics ([#810](https://github.com/cloudquery/cloudquery/issues/810)) ([75e0f02](https://github.com/cloudquery/cloudquery/commit/75e0f0278dc345d14ae89963926284788ed50522))

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.4 ([#808](https://github.com/cloudquery/cloudquery/issues/808)) ([b91abf2](https://github.com/cloudquery/cloudquery/commit/b91abf2b2378e856eb9403cee25e9a3c4886dfde))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.5 ([#812](https://github.com/cloudquery/cloudquery/issues/812)) ([0d23f6c](https://github.com/cloudquery/cloudquery/commit/0d23f6ced6181177943548a23e7fa7fc446a07b5))
- Panic During Policy Storage ([#803](https://github.com/cloudquery/cloudquery/issues/803)) ([570d9e2](https://github.com/cloudquery/cloudquery/commit/570d9e24faa7e92adaf26e824ad32126082bf482))

### [0.24.2](https://github.com/cloudquery/cloudquery/compare/v0.24.1...v0.24.2) (2022-05-27)

### Bug Fixes

- Add Git Binary To Docker ([f30c2fe](https://github.com/cloudquery/cloudquery/commit/f30c2fe8f5a262d3f1f044aff9e51ef508f59ee8))
- Skip reattached provider update checks ([#801](https://github.com/cloudquery/cloudquery/issues/801)) ([714b446](https://github.com/cloudquery/cloudquery/commit/714b44678f8afcb2a27b967b89b86f59a5ecb391))

### [0.24.1](https://github.com/cloudquery/cloudquery/compare/v0.24.0...v0.24.1) (2022-05-26)

### Features

- Store Policy Output ([#709](https://github.com/cloudquery/cloudquery/issues/709)) ([bda4a50](https://github.com/cloudquery/cloudquery/commit/bda4a50d7ddf446f92ba72850e8d1e778620be6a))

### Bug Fixes

- Console log log level ([#786](https://github.com/cloudquery/cloudquery/issues/786)) ([8a7b76e](https://github.com/cloudquery/cloudquery/commit/8a7b76ecdeaceb62dd7b5cd376aaf8760ea5f9e3))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.2 ([#784](https://github.com/cloudquery/cloudquery/issues/784)) ([f904794](https://github.com/cloudquery/cloudquery/commit/f9047942373cd0650412b7eb6e9a4a1694f31596))
- Error classifier improvements ([#788](https://github.com/cloudquery/cloudquery/issues/788)) ([f32a701](https://github.com/cloudquery/cloudquery/commit/f32a701a8893cdb2f0b5f28c304a98d341ef4f82))
- Panic in loadPolicyFromSource ([#787](https://github.com/cloudquery/cloudquery/issues/787)) ([5953d09](https://github.com/cloudquery/cloudquery/commit/5953d091110c30f6d8760aadc1784fe9a88eff95))
- Panic in OsFs.downloadFile ([#789](https://github.com/cloudquery/cloudquery/issues/789)) ([f2c76e8](https://github.com/cloudquery/cloudquery/commit/f2c76e8d9e95d77b479f0d12bac4bfdfebb9e15a))
- Remove always nil return value ([#778](https://github.com/cloudquery/cloudquery/issues/778)) ([dca8745](https://github.com/cloudquery/cloudquery/commit/dca8745c24fbff70c7b3ba89eafb87b90259daeb))
- Test flakiness ([#790](https://github.com/cloudquery/cloudquery/issues/790)) ([1497c51](https://github.com/cloudquery/cloudquery/commit/1497c511da1f31e7b0cf8eca2a9b159580bbe8e0))

## [0.24.0](https://github.com/cloudquery/cloudquery/compare/v0.23.4...v0.24.0) (2022-05-24)

### ⚠ BREAKING CHANGES

- Remove provider migrations (#731)

### Features

- Classify some policy download errors as USER ([#742](https://github.com/cloudquery/cloudquery/issues/742)) ([8224e60](https://github.com/cloudquery/cloudquery/commit/8224e60d3a76d3b3f181d3b32b9153a63b04816a))
- Remove provider migrations ([#731](https://github.com/cloudquery/cloudquery/issues/731)) ([bb93967](https://github.com/cloudquery/cloudquery/commit/bb93967803a28dfc1fccc0fb45af207d76208b2d))

### Bug Fixes

- Added root policy config inheritance to selected subpolicy ([#702](https://github.com/cloudquery/cloudquery/issues/702)) ([0d8b3c7](https://github.com/cloudquery/cloudquery/commit/0d8b3c77741177e6417c59e269856639dce0155e))
- Check policy version on it's core version ([#773](https://github.com/cloudquery/cloudquery/issues/773)) ([c7c9ad0](https://github.com/cloudquery/cloudquery/commit/c7c9ad0384e2b90c6507c52ab87da65db2bc836f))
- Classify "no policies in config" as USER error ([#743](https://github.com/cloudquery/cloudquery/issues/743)) ([4cbc03e](https://github.com/cloudquery/cloudquery/commit/4cbc03e22f5a0bbfa33812b407e65704727a88fd))
- **deps:** Update SDK ([#758](https://github.com/cloudquery/cloudquery/issues/758)) ([c9b4094](https://github.com/cloudquery/cloudquery/commit/c9b4094cfe2ebbae9b6cbbc8df4568c7460208e5))
- Diag import clean up ([#744](https://github.com/cloudquery/cloudquery/issues/744)) ([437c956](https://github.com/cloudquery/cloudquery/commit/437c956ee941c3ed2c9859a0af6a0b88401b58d8))
- Don't attempt to download provider in re-attach mode ([#748](https://github.com/cloudquery/cloudquery/issues/748)) ([59973b8](https://github.com/cloudquery/cloudquery/commit/59973b84826599915f7b76fc8d8b16626dd26c74))
- FetchId column regression ([#745](https://github.com/cloudquery/cloudquery/issues/745)) ([585d395](https://github.com/cloudquery/cloudquery/commit/585d39589ef6c27ae2aab5d224fc00a2387d7628))
- Handle DeadlineExceeded errors ([#741](https://github.com/cloudquery/cloudquery/issues/741)) ([0167ce4](https://github.com/cloudquery/cloudquery/commit/0167ce4158d4795fc3a4b0f6661c19ae197c20c9))
- Handle Outputting Policies With Selectors ([a3ecfc9](https://github.com/cloudquery/cloudquery/commit/a3ecfc9166170e1bb77011befd11a5fbe1c86007))
- Policy executor ([#769](https://github.com/cloudquery/cloudquery/issues/769)) ([d5b6aef](https://github.com/cloudquery/cloudquery/commit/d5b6aef25f1cccaaf30618c53d5d7204f83d74aa))
- Policy output file name ([#770](https://github.com/cloudquery/cloudquery/issues/770)) ([1a87c25](https://github.com/cloudquery/cloudquery/commit/1a87c259ddf9d0d1694976f551503918ca1557bd))
- Space trimming in telemetry file ([#734](https://github.com/cloudquery/cloudquery/issues/734)) ([16c4cfc](https://github.com/cloudquery/cloudquery/commit/16c4cfce7e15f4474af3ab5d7e0cdb3698d2d08e))
- Upgrade protocol version to V5 ([#774](https://github.com/cloudquery/cloudquery/issues/774)) ([69b405d](https://github.com/cloudquery/cloudquery/commit/69b405d686031dec3443ebc018047f42dd259d0e))
- Use consistent descriptions for flags ([#753](https://github.com/cloudquery/cloudquery/issues/753)) ([cedeb3d](https://github.com/cloudquery/cloudquery/commit/cedeb3d0b1d733b352a03d45753914a6e1d11ee4))

### [0.23.4](https://github.com/cloudquery/cloudquery/compare/v0.23.3...v0.23.4) (2022-05-17)

### Bug Fixes

- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.4 ([#725](https://github.com/cloudquery/cloudquery/issues/725)) ([69afca7](https://github.com/cloudquery/cloudquery/commit/69afca7dd34200ef0fda2341293e3fb46ee75faa))
- Don't generate telemetry-random-id in current dir ([#729](https://github.com/cloudquery/cloudquery/issues/729)) ([5eb493b](https://github.com/cloudquery/cloudquery/commit/5eb493b7215dc488d515106beff1a863a384b002))
- Panic on nil fetch response ([#728](https://github.com/cloudquery/cloudquery/issues/728)) ([8118554](https://github.com/cloudquery/cloudquery/commit/811855475622955dcfb323298292bde958f4372d))
- Panic on nil fetch result ([#730](https://github.com/cloudquery/cloudquery/issues/730)) ([7f224d0](https://github.com/cloudquery/cloudquery/commit/7f224d0371ed0014948fb6c572adef20bdb16094))
- squash redact errors ([#727](https://github.com/cloudquery/cloudquery/issues/727)) ([bccf7b7](https://github.com/cloudquery/cloudquery/commit/bccf7b71094eef4552e9227e4290aeec9a47896f))

### [0.23.3](https://github.com/cloudquery/cloudquery/compare/v0.23.2...v0.23.3) (2022-05-17)

### Features

- Add global log id ([#714](https://github.com/cloudquery/cloudquery/issues/714)) ([cece150](https://github.com/cloudquery/cloudquery/commit/cece150a78c83365a36cb3c295de8218ae959995))
- Policy Output ([#664](https://github.com/cloudquery/cloudquery/issues/664)) ([31f7e19](https://github.com/cloudquery/cloudquery/commit/31f7e19463da541b5ec13e18f4faf6d91dcfe6b0))
- Resource list enhancements ([#706](https://github.com/cloudquery/cloudquery/issues/706)) ([1952a27](https://github.com/cloudquery/cloudquery/commit/1952a27f212e109bac7bc74761cf193478aa1289))
- Use database id as unique id ([#705](https://github.com/cloudquery/cloudquery/issues/705)) ([dc00381](https://github.com/cloudquery/cloudquery/commit/dc0038158924b48ac41cbe57f7140084f2059ec3))

### Bug Fixes

- Add missing descriptions ([#700](https://github.com/cloudquery/cloudquery/issues/700)) ([c3c288c](https://github.com/cloudquery/cloudquery/commit/c3c288c62ff134109b2f35ec1a73b6cdd63c2d72))
- Classify not found policies and improve errors ([#697](https://github.com/cloudquery/cloudquery/issues/697)) ([413a2cf](https://github.com/cloudquery/cloudquery/commit/413a2cfe757f6a29ebc2fdb2db07b99b1fa9c4a1))
- Classify policy parse errors as User ([#716](https://github.com/cloudquery/cloudquery/issues/716)) ([f5947bf](https://github.com/cloudquery/cloudquery/commit/f5947bf443631454d41c2764c45bb32e5cfc2058))
- Classify subdir not found error ([#701](https://github.com/cloudquery/cloudquery/issues/701)) ([1a30732](https://github.com/cloudquery/cloudquery/commit/1a307321dab4c75c6697b20f8756d7282689a5cf))
- Completion issue ([#703](https://github.com/cloudquery/cloudquery/issues/703)) ([21c7bfe](https://github.com/cloudquery/cloudquery/commit/21c7bfeeb7afee4f1da7b8492e7be3a4c92b2bca))
- Handle empty policy directory ([#699](https://github.com/cloudquery/cloudquery/issues/699)) ([6acd308](https://github.com/cloudquery/cloudquery/commit/6acd3087cb3a81d990c77f351b969445e12d2bfd))
- Remove empty keys from init config ([#696](https://github.com/cloudquery/cloudquery/issues/696)) ([0e8dda1](https://github.com/cloudquery/cloudquery/commit/0e8dda1aecf5ac8ca785f1f9d4912b412b040ae8))
- Remove lambda support ([#710](https://github.com/cloudquery/cloudquery/issues/710)) ([5254f34](https://github.com/cloudquery/cloudquery/commit/5254f34f30f96b27d82e627a2be6c302bcb174af))
- Remove unused lambda dependency ([#717](https://github.com/cloudquery/cloudquery/issues/717)) ([7c78974](https://github.com/cloudquery/cloudquery/commit/7c78974668ad4144c7d9ded285cb4290fb0b01e6))
- Set ID For all Versions ([#724](https://github.com/cloudquery/cloudquery/issues/724)) ([ac46d2a](https://github.com/cloudquery/cloudquery/commit/ac46d2ad77bc8987e693028211a034bfe70cb06f))

### [0.23.2](https://github.com/cloudquery/cloudquery/compare/v0.23.1...v0.23.2) (2022-05-11)

### Bug Fixes

- **deps:** Bump github.com/hashicorp/go-getter from 1.5.10 to 1.5.11 ([#691](https://github.com/cloudquery/cloudquery/issues/691)) ([2ef215e](https://github.com/cloudquery/cloudquery/commit/2ef215e70af2de6243e2fd424c6785a920a8bfb2))

### [0.23.1](https://github.com/cloudquery/cloudquery/compare/v0.23.0...v0.23.1) (2022-05-11)

### Features

- DSN credentials ([#670](https://github.com/cloudquery/cloudquery/issues/670)) ([35e27d0](https://github.com/cloudquery/cloudquery/commit/35e27d03bb4d1102c93b04b981ed435720171386))

### Bug Fixes

- Handle nil policy run response ([#688](https://github.com/cloudquery/cloudquery/issues/688)) ([bd3e3bd](https://github.com/cloudquery/cloudquery/commit/bd3e3bd36e7a531f0fdb56378c658a9822b1166e))
- Run detectors in order ([#690](https://github.com/cloudquery/cloudquery/issues/690)) ([a39b2b6](https://github.com/cloudquery/cloudquery/commit/a39b2b6c878d41bcd78e81e84daf1ee95f05d125))

## [0.23.0](https://github.com/cloudquery/cloudquery/compare/v0.22.10...v0.23.0) (2022-05-10)

### Features

- Change to rudder ([#650](https://github.com/cloudquery/cloudquery/issues/650)) ([8f3f4c1](https://github.com/cloudquery/cloudquery/commit/8f3f4c14be4b7f95b7c673b1de6d4c2153556f93))
- Track db installations ([#652](https://github.com/cloudquery/cloudquery/issues/652)) ([e38acb7](https://github.com/cloudquery/cloudquery/commit/e38acb7d70297f764b1683dffe8389d908636369))

### Bug Fixes

- Bug where policy_run always fails ([#667](https://github.com/cloudquery/cloudquery/issues/667)) ([402266e](https://github.com/cloudquery/cloudquery/commit/402266ec8995bcd36d58093a2072efa795d89a1b))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.2 ([#637](https://github.com/cloudquery/cloudquery/issues/637)) ([55a60a9](https://github.com/cloudquery/cloudquery/commit/55a60a95328e4b5db00a5689ce5da5aed46dcbe5))
- **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.3 ([#658](https://github.com/cloudquery/cloudquery/issues/658)) ([351cce5](https://github.com/cloudquery/cloudquery/commit/351cce50ecfa03d1be6dfd1d3fd7e268368a8aeb))
- Enable logging config through HCL ([#604](https://github.com/cloudquery/cloudquery/issues/604)) ([51bd06c](https://github.com/cloudquery/cloudquery/commit/51bd06c83f2a371b9e18969faf6edb1967b62e62))
- Encode json ([#641](https://github.com/cloudquery/cloudquery/issues/641)) ([1c04e45](https://github.com/cloudquery/cloudquery/commit/1c04e4515a9d865b92475200c9959253735ca9cd))
- panic on sync failure ([#676](https://github.com/cloudquery/cloudquery/issues/676)) ([27d574f](https://github.com/cloudquery/cloudquery/commit/27d574f6262417071c615675ec22b586317c50aa))
- **policy:** Add missing GitHub getter ([#613](https://github.com/cloudquery/cloudquery/issues/613)) ([e3fc361](https://github.com/cloudquery/cloudquery/commit/e3fc361c12139c58de14e42ab7ba89f2a967508a))
- **policy:** Use firebase instead of GitHub API to get latest version ([#618](https://github.com/cloudquery/cloudquery/issues/618)) ([455ed23](https://github.com/cloudquery/cloudquery/commit/455ed23ca3f0d075028385359a47436b8b05ead9))
- Sync support optional provider args ([#642](https://github.com/cloudquery/cloudquery/issues/642)) ([5eac023](https://github.com/cloudquery/cloudquery/commit/5eac02321222f6a50b95308274cc631402ab213a))
- Validate db version before proceeding ([#653](https://github.com/cloudquery/cloudquery/issues/653)) ([5af7f61](https://github.com/cloudquery/cloudquery/commit/5af7f615c580e94d319e2ad99b470ead9afd18f2))

### Miscellaneous Chores

- Release 0.23.0 ([#674](https://github.com/cloudquery/cloudquery/issues/674)) ([d4a2502](https://github.com/cloudquery/cloudquery/commit/d4a250288832b28104ae7e5497fbe6dc9a8f1231))

## [v0.19.0] - 2022-01-10

### Breaking Changes

- Policy command updated and spec changed [#369](https://github.com/cloudquery/cloudquery/pull/369)

### Fixed

- Fixed empty policy bug [#399](https://github.com/cloudquery/cloudquery/pull/399).
- Fixed lambda json conversion [#397](https://github.com/cloudquery/cloudquery/pull/397).
- Removed confusing error message [#391](https://github.com/cloudquery/cloudquery/pull/391).
- Respected absolute file path in policies [#395](https://github.com/cloudquery/cloudquery/pull/395).
- Fixed isLevel for logger [#385](https://github.com/cloudquery/cloudquery/pull/385).
- Fixed pathing for hub to use real source path [#394](https://github.com/cloudquery/cloudquery/pull/394).
- CreateDatabase: check for err in correct place [#389](https://github.com/cloudquery/cloudquery/pull/389).
- Prevented reporting of errors to sentry twice [#386](https://github.com/cloudquery/cloudquery/pull/386).

### :gear: Changed

- Removed stack traces from sentry [#387](https://github.com/cloudquery/cloudquery/pull/87).
- Sentry send stack trace only on panic [#390](https://github.com/cloudquery/cloudquery/pull/390).

## [v0.18.0]- 2022-01-03

### 🚀 Added

- On cancel show error [#371](https://github.com/cloudquery/cloudquery/pull/371)

### 💥 Breaking Changes

- Upgrade to sdk [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/releases/tag/v0.6.1)

### :gear: Changed

- remove the need for json hcl2json convertor [#375](https://github.com/cloudquery/cloudquery/pull/375)
- removed gen config [#370](https://github.com/cloudquery/cloudquery/pull/370)

### :spider: Fixed

- Table upgrades with history mode enabled [#381](https://github.com/cloudquery/cloudquery/pull/381).

## [v0.17.4]- 2021-12-23

### 💥 Breaking Changes

- Removed old terraform deployment from core package, new deployment located [here](https://github.com/cloudquery/terraform-aws-cloudquery) [#357](https://github.com/cloudquery/cloudquery/pull/357).

### :rocket: Added

- Drift: Use correct ID for EMR clusters [#360](https://github.com/cloudquery/cloudquery/pull/360).
- Policy: added more logging to policy execution [#341](https://github.com/cloudquery/cloudquery/pull/341).
- Added hash of config to telemetry [#358](https://github.com/cloudquery/cloudquery/pull/359).

### :spider: Fixed

- Fixed Sentry issues [#347](https://github.com/cloudquery/cloudquery/pull/347).

### :gear: Changed

- Changed how we classify errors for sentry reducing errors sent, so only critical errors are report [#350](https://github.com/cloudquery/cloudquery/pull/350).
- Disable sentry module reporting [#351](https://github.com/cloudquery/cloudquery/pull/351).
- Made `source` attribute optional in CloudQuery config [#352](https://github.com/cloudquery/cloudquery/pull/352).
- Improved misleading help messaeg in cloudquery init [#359](https://github.com/cloudquery/cloudquery/pull/359).

## [v0.17.3]- 2021-12-16

### :spider: Fixed

- Report panics to Sentry [#347](https://github.com/cloudquery/cloudquery/pull/347).

## [v0.17.2] - 2021-12-16

### :spider: Fixed

- Panic on `cloudquery fetch`

## [v0.17.1] - 2021-12-15

### :rocket: Added

- Added [#210](https://github.com/cloudquery/cloudquery/issues/210) contribution [guide](https://github.com/cloudquery/cloudquery/blob/main/CONTRIBUTING.md) [#331](https://github.com/cloudquery/cloudquery/pull/331).
- Added new provider update available notification [#336](https://github.com/cloudquery/cloudquery/pull/336) fixes [#299](https://github.com/cloudquery/cloudquery/issues/299).
- Added notification if an update to CQ core is available [#338](https://github.com/cloudquery/cloudquery/pull/338).
- Added sentry for crash error reporting to improve stability [#342](https://github.com/cloudquery/cloudquery/pull/342).

### :gear: Changed

- Telemetry: collect hash of MAC + Hostname [#339](https://github.com/cloudquery/cloudquery/pull/339).

### :spider: Fixed

- Provider download routine added before to policy run command [#335](https://github.com/cloudquery/cloudquery/pull/335) fixes [#316](https://github.com/cloudquery/cloudquery/issues/316).
- Fixed [#303](https://github.com/cloudquery/cloudquery/issues/303) UUID output in policies [#332](https://github.com/cloudquery/cloudquery/pull/332).
- Fixed Telemetry error counting, changed `debug-telemetry` flag to only set open-telelmetry client to debug mode [#340](https://github.com/cloudquery/cloudquery/pull/340)

## [v0.17.0] - 2021-12-06

### 💥 Breaking Changes

- `policy run` flag `--subpath` has been removed to execute sub policy pass it as second argument i.e `policy run <policy_name> <subpath>`

### :rocket: Added

- Added `policy describe <policy_name>` subcommand, allowing to see all policies and sub-policies available and execution paths
- Added support for CloudQuery History **Alpha** for more info see [docs](https://docs.cloudquery.io/cli/history/overview)
- Exposed diagnostic counts on fetch for telemetry [#319](https://github.com/cloudquery/cloudquery/pull/319)

### :spider: Fixed

- Fixed resource fetch summary total fetched resources wouldn't sum correctly [#326](https://github.com/cloudquery/cloudquery/pull/326)
- Provider fetch failure cancels out other provider fetches [#325](https://github.com/cloudquery/cloudquery/pull/325)

### :gear: Changed

- Upgraded to SDK Version [v0.5.3](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).
- Read persisted telemetry ID if exists [#313](https://github.com/cloudquery/cloudquery/pull/313)
- Cleanup init command [#320](https://github.com/cloudquery/cloudquery/pull/320)
- Improve logging for policy execution errors [#323](https://github.com/cloudquery/cloudquery/pull/323)
- Updated drift aws configuration for new version [#329](https://github.com/cloudquery/cloudquery/pull/329)

## [v0.16.2] - 2021-11-29

### :rocket: Added

- Added support for telemetry, to gain better insight on usage to improve features and tool performance. For additional info see [docs](https://docs.cloudquery.io/docs/cli/telemetry) [#280](https://github.com/cloudquery/cloudquery/pull/280).
- Added support for executing policy in policy [#302](https://github.com/cloudquery/cloudquery/issues/302)

### :spider: Fixed

- Fixed Policy Not Found unclear message [#306](https://github.com/cloudquery/cloudquery/issues/306)
- Fixed Logging Statements Output [#305](https://github.com/cloudquery/cloudquery/issues/305)

## [v0.16.1] - 2021-11-22

### :spider: Fixed

- Fix fetch failure on providers that don't support upgrade [#295](https://github.com/cloudquery/cloudquery/pull/295)

## [v0.16.0] - 2021-11-19

### :rocket: Added

- Added support for [Terraform Drift detection](https://www.cloudquery.io/blog/announcing-cloudquery-terraform-drift-detection).
- Allow regex patterns for drift configuration (both local files and s3 bucket + keys [#281](https://github.com/cloudquery/cloudquery/issues/281)
- Run provider upgrades before fetch [#283](https://github.com/cloudquery/cloudquery/pull/283)
- Support running policies from configuration [#269](https://github.com/cloudquery/cloudquery/pull/269)
- Added a changelog :rocket:

### :spider: Fixed

- Fixed Confusing Error when config.hcl doesn't exist [#277](https://github.com/cloudquery/cloudquery/issues/277)

## [0.15.11] - 2021-11-18

Base version at which changelog was introduced.
