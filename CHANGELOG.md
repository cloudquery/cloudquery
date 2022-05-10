# Changelog

All notable changes to CloudQuery will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
## [0.23.0](https://github.com/cloudquery/cloudquery/compare/v0.23.0...v0.23.0) (2022-05-10)


### Features

* Add '--disable-fetch-check' flag ([#567](https://github.com/cloudquery/cloudquery/issues/567)) ([e064f59](https://github.com/cloudquery/cloudquery/commit/e064f591a51607d79d8ccddcb45c29ef8540d1ac))
* Add console colour disable ([#559](https://github.com/cloudquery/cloudquery/issues/559)) ([903d369](https://github.com/cloudquery/cloudquery/commit/903d36963d92433ad4ea0dafac78cb8b8a11cbc9))
* Add file() func to config parser ([#522](https://github.com/cloudquery/cloudquery/issues/522)) ([9f05cd4](https://github.com/cloudquery/cloudquery/commit/9f05cd4c92ed057fd9071183bffcb92982931d72))
* Add max_goroutines to providers ([#477](https://github.com/cloudquery/cloudquery/issues/477)) ([6c627bd](https://github.com/cloudquery/cloudquery/commit/6c627bdefe5183bff81337e0873e12e80f2a716b))
* Added checking of finished fetches before running policy ([#444](https://github.com/cloudquery/cloudquery/issues/444)) ([327398f](https://github.com/cloudquery/cloudquery/commit/327398f0fadbc0a72c6205c15181ec69e180b863))
* Added execution time to policy output file ([#424](https://github.com/cloudquery/cloudquery/issues/424)) ([420330c](https://github.com/cloudquery/cloudquery/commit/420330c932f770dafa7b8b8cf0c1e310147ecc70))
* Added store fetch summary routine ([#356](https://github.com/cloudquery/cloudquery/issues/356)) ([d6e24db](https://github.com/cloudquery/cloudquery/commit/d6e24db44307b4531a6142ccb1f3dd5b06a38958))
* Adjust log messages when .cq dir is absent ([#411](https://github.com/cloudquery/cloudquery/issues/411)) ([b965db4](https://github.com/cloudquery/cloudquery/commit/b965db40d380514ed5e61c1bc103ad746c849fa4))
* Change to rudder ([#650](https://github.com/cloudquery/cloudquery/issues/650)) ([8f3f4c1](https://github.com/cloudquery/cloudquery/commit/8f3f4c14be4b7f95b7c673b1de6d4c2153556f93))
* Check for slashes in policy ([#561](https://github.com/cloudquery/cloudquery/issues/561)) ([3ad0b5b](https://github.com/cloudquery/cloudquery/commit/3ad0b5b2d302e87dab414c3171409e2a5bf3cd45))
* Configurable data folder ([#456](https://github.com/cloudquery/cloudquery/issues/456)) ([038bb26](https://github.com/cloudquery/cloudquery/commit/038bb26d22c24730e51e33c8d7d5aba1e27513b5))
* Configure to return diagnostics, upgrade SDK ([#594](https://github.com/cloudquery/cloudquery/issues/594)) ([879438d](https://github.com/cloudquery/cloudquery/commit/879438db44244bf220555263d1c06670750d6517))
* Drift: Add [@getbool](https://github.com/getbool) modifier ([#515](https://github.com/cloudquery/cloudquery/issues/515)) ([38cfa36](https://github.com/cloudquery/cloudquery/commit/38cfa36ff698a0b7e206b8ba06a990b034c6c7f8))
* Drift: Improvements ([#563](https://github.com/cloudquery/cloudquery/issues/563)) ([a7b9455](https://github.com/cloudquery/cloudquery/commit/a7b94550eca9ed6be7bf0035e06974f573794f37))
* Drift: List unimplemented resources in debug mode ([#558](https://github.com/cloudquery/cloudquery/issues/558)) ([0d2a9b9](https://github.com/cloudquery/cloudquery/commit/0d2a9b921cb54b31274a2ae7c4029567115a996b))
* Enable remote fetch ([#591](https://github.com/cloudquery/cloudquery/issues/591)) ([1d5bfb7](https://github.com/cloudquery/cloudquery/commit/1d5bfb7979ec3df48eb691793798bfa753e1e6cf))
* Expose max_parallel_resource_fetch_limit ([#413](https://github.com/cloudquery/cloudquery/issues/413)) ([b447809](https://github.com/cloudquery/cloudquery/commit/b44780936f4034cb02aea68868ae34d3d5779c3f))
* Fetch only specified providers/resources ([#468](https://github.com/cloudquery/cloudquery/issues/468)) ([2b55e76](https://github.com/cloudquery/cloudquery/commit/2b55e7647fa488f033a2ffdd936924a4d92a967b))
* Ignore ignored diags unless -v is specified ([#545](https://github.com/cloudquery/cloudquery/issues/545)) ([b3ba500](https://github.com/cloudquery/cloudquery/commit/b3ba5002d61c5dfaae9bfd582c086f33e3772c70))
* Increase ulimit in unix environment ([#416](https://github.com/cloudquery/cloudquery/issues/416)) ([7440844](https://github.com/cloudquery/cloudquery/commit/7440844c2c7e9dd5100923af5635859319598067))
* Limit scope of views ([#556](https://github.com/cloudquery/cloudquery/issues/556)) ([83d82b2](https://github.com/cloudquery/cloudquery/commit/83d82b20cad027b7f37874f18957e1fe5a11f326))
* Linting action ([#75](https://github.com/cloudquery/cloudquery/issues/75)) ([d0835e9](https://github.com/cloudquery/cloudquery/commit/d0835e916929e54e992b6a6db90e20699f1734c8))
* Migrations v2 (core) ([#406](https://github.com/cloudquery/cloudquery/issues/406)) ([c4e1b5e](https://github.com/cloudquery/cloudquery/commit/c4e1b5e7021bad6bc92c07bce4b2d975ffe205e1))
* Modules v2 support ([#489](https://github.com/cloudquery/cloudquery/issues/489)) ([b4c4d2c](https://github.com/cloudquery/cloudquery/commit/b4c4d2c6ffe74e5bb721dd77327a54d8ff202e5a))
* Policy Snapshot Testing ([#449](https://github.com/cloudquery/cloudquery/issues/449)) ([c6c24ed](https://github.com/cloudquery/cloudquery/commit/c6c24ed47da5914e822214e0992e08c6dddc21fe))
* Purge stale data ([#565](https://github.com/cloudquery/cloudquery/issues/565)) ([d06bad3](https://github.com/cloudquery/cloudquery/commit/d06bad30f59f0090fbcfd12582f12e47a2451f81))
* Remote config ([#568](https://github.com/cloudquery/cloudquery/issues/568)) ([227f398](https://github.com/cloudquery/cloudquery/commit/227f39821ef0888dada17ba8f3dd041c5bcaac1d))
* Remove filtering from executor ([#429](https://github.com/cloudquery/cloudquery/issues/429)) ([0317672](https://github.com/cloudquery/cloudquery/commit/0317672b3e3138144c0bb0733110427f31a3c953))
* Report on multiple aliased providers ([#437](https://github.com/cloudquery/cloudquery/issues/437)) ([d987a34](https://github.com/cloudquery/cloudquery/commit/d987a34de397c3eb39c0ba6efb9348f8edca4d97))
* Simplify fetch warning/error counts (unless -v is on) ([#554](https://github.com/cloudquery/cloudquery/issues/554)) ([3d738f2](https://github.com/cloudquery/cloudquery/commit/3d738f2a4f48a60fb7bfd4e1276447cd8f57a079))
* Small improvements ([#459](https://github.com/cloudquery/cloudquery/issues/459)) ([fabd906](https://github.com/cloudquery/cloudquery/commit/fabd9063a15ad1eb97c7fa20fd1b88e17bc28500))
* Store cq_fetch_id in meta ([#523](https://github.com/cloudquery/cloudquery/issues/523)) ([083b5fb](https://github.com/cloudquery/cloudquery/commit/083b5fbb23e309bc36a72faef17e582a9a888317))
* Support build-schema for all providers in config ([#414](https://github.com/cloudquery/cloudquery/issues/414)) ([ad1ae13](https://github.com/cloudquery/cloudquery/commit/ad1ae13c30c489a260771be7e6ac451b486fb5ac))
* Support selector on describe ([#425](https://github.com/cloudquery/cloudquery/issues/425)) ([fc37402](https://github.com/cloudquery/cloudquery/commit/fc374025ed9abfa5cbf6f32baa74f02fe0f97d09))
* Support timeouts in fetch ([#537](https://github.com/cloudquery/cloudquery/issues/537)) ([2442461](https://github.com/cloudquery/cloudquery/commit/244246194f67d1ca9183b26902975543d60172fb))
* Track db installations ([#652](https://github.com/cloudquery/cloudquery/issues/652)) ([e38acb7](https://github.com/cloudquery/cloudquery/commit/e38acb7d70297f764b1683dffe8389d908636369))
* Update drift for AWS IOT resources ([#434](https://github.com/cloudquery/cloudquery/issues/434)) ([c6709f3](https://github.com/cloudquery/cloudquery/commit/c6709f3d4ea096ccfbf18080c991fece16cf6774))
* Update number of args purge command ([#589](https://github.com/cloudquery/cloudquery/issues/589)) ([e5522fa](https://github.com/cloudquery/cloudquery/commit/e5522fac050d53c3d1e357ea09df10529a8969ca))
* Update SDK and show resource id on diagnostics ([#480](https://github.com/cloudquery/cloudquery/issues/480)) ([9daead7](https://github.com/cloudquery/cloudquery/commit/9daead7be47db12ad86e2d1c1178d71920c3f7e8))
* Update SDK v0.7.0 ([#435](https://github.com/cloudquery/cloudquery/issues/435)) ([b232ed2](https://github.com/cloudquery/cloudquery/commit/b232ed233a042e3ee054180373bb364514b10aa3))
* Use firestore for releases ([#483](https://github.com/cloudquery/cloudquery/issues/483)) ([63f8d34](https://github.com/cloudquery/cloudquery/commit/63f8d34ddf9915b0d1433988e219856f295b3169))
* Use provider name/version as exception type when reporting diagnostics to Sentry ([#475](https://github.com/cloudquery/cloudquery/issues/475)) ([196c818](https://github.com/cloudquery/cloudquery/commit/196c8188cb5bd4ab503921aabb7fdcf67fc0aac7))
* Use redacted errors for sentry/telemetry ([#461](https://github.com/cloudquery/cloudquery/issues/461)) ([fbf5fb9](https://github.com/cloudquery/cloudquery/commit/fbf5fb9ed7f81ad919035558ea60550c9d9b7876))


### Bug Fixes

* Add ProviderVersion to the 'cloudquery.fetches' meta table ([#555](https://github.com/cloudquery/cloudquery/issues/555)) ([422905d](https://github.com/cloudquery/cloudquery/commit/422905d5a2a17a1e2baa795be5ffa35876235ca7))
* Add safeguard for drop provider ([#580](https://github.com/cloudquery/cloudquery/issues/580)) ([4c2b3e7](https://github.com/cloudquery/cloudquery/commit/4c2b3e7dca1b283824299d7c63299ba4167b2026))
* Added error when connection configuration is not set ([#476](https://github.com/cloudquery/cloudquery/issues/476)) ([2674f8e](https://github.com/cloudquery/cloudquery/commit/2674f8ea4b4cfbd4426a485b13dc65c668de2d0f))
* Adjust policy describe message ([#409](https://github.com/cloudquery/cloudquery/issues/409)) ([19dc6f7](https://github.com/cloudquery/cloudquery/commit/19dc6f7685517002a099264e08ef8f060787cf01))
* Better drift tfstate error ([#430](https://github.com/cloudquery/cloudquery/issues/430)) ([001a06d](https://github.com/cloudquery/cloudquery/commit/001a06df14f243efcb0e11e91a6ac89f237b828c))
* Bug where policy_run always fails ([#667](https://github.com/cloudquery/cloudquery/issues/667)) ([402266e](https://github.com/cloudquery/cloudquery/commit/402266ec8995bcd36d58093a2072efa795d89a1b))
* DB version validation shouldn't block on vanilla PG ([#447](https://github.com/cloudquery/cloudquery/issues/447)) ([bf58d82](https://github.com/cloudquery/cloudquery/commit/bf58d820b60d53d154560248df0402410f1e523f))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.2 ([#637](https://github.com/cloudquery/cloudquery/issues/637)) ([55a60a9](https://github.com/cloudquery/cloudquery/commit/55a60a95328e4b5db00a5689ce5da5aed46dcbe5))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.3 ([#658](https://github.com/cloudquery/cloudquery/issues/658)) ([351cce5](https://github.com/cloudquery/cloudquery/commit/351cce50ecfa03d1be6dfd1d3fd7e268368a8aeb))
* Diag: Don't report errors we don't care about ([#521](https://github.com/cloudquery/cloudquery/issues/521)) ([8079070](https://github.com/cloudquery/cloudquery/commit/8079070eb15ac1bb324e4ab4af2bb87e45caa0ba))
* Don't connect on cloudquery init ([#436](https://github.com/cloudquery/cloudquery/issues/436)) ([d053299](https://github.com/cloudquery/cloudquery/commit/d053299e17f7adea17c5498fa20e04edf78aa993))
* Don't report unmanaged provider diags to sentry ([#492](https://github.com/cloudquery/cloudquery/issues/492)) ([36e13e4](https://github.com/cloudquery/cloudquery/commit/36e13e470a427dd8905a2919b49aeb42a48b9443))
* Don't return an error if encountering an empty subpolicy ([#486](https://github.com/cloudquery/cloudquery/issues/486)) ([6a126a6](https://github.com/cloudquery/cloudquery/commit/6a126a673e69afe34e352c03c2e5beb1578687d1))
* Don't show telemetry notice when it's not enabled ([#418](https://github.com/cloudquery/cloudquery/issues/418)) ([a379bc2](https://github.com/cloudquery/cloudquery/commit/a379bc2faec59396c06b09713f6bd62e343cbe12))
* Downgrade provider ([#531](https://github.com/cloudquery/cloudquery/issues/531)) ([5e7c84a](https://github.com/cloudquery/cloudquery/commit/5e7c84a61e73e476ab4fd4a9386fd589ff1ef333))
* Drift: Remove version constraints, always include all parents ([#575](https://github.com/cloudquery/cloudquery/issues/575)) ([b1f8263](https://github.com/cloudquery/cloudquery/commit/b1f826313cafdde2327820a03c17f152346138ec))
* Empty source in provider ([#465](https://github.com/cloudquery/cloudquery/issues/465)) ([0144698](https://github.com/cloudquery/cloudquery/commit/01446988f7d4b25cf4cc32497784de45113eddc2))
* Enable logging config through HCL ([#604](https://github.com/cloudquery/cloudquery/issues/604)) ([51bd06c](https://github.com/cloudquery/cloudquery/commit/51bd06c83f2a371b9e18969faf6edb1967b62e62))
* Encode json ([#641](https://github.com/cloudquery/cloudquery/issues/641)) ([1c04e45](https://github.com/cloudquery/cloudquery/commit/1c04e4515a9d865b92475200c9959253735ca9cd))
* Exit with 1 on policy error ([#412](https://github.com/cloudquery/cloudquery/issues/412)) ([a9d7847](https://github.com/cloudquery/cloudquery/commit/a9d78473598143f21427ac54165e87880bc74b4c))
* Fetch summary fixes ([#417](https://github.com/cloudquery/cloudquery/issues/417)) ([ff24dd9](https://github.com/cloudquery/cloudquery/commit/ff24dd9548551440c53a12f370f5418a557f7a1c))
* Fetch summary SQL state 54000 ([#487](https://github.com/cloudquery/cloudquery/issues/487)) ([9e6dbd9](https://github.com/cloudquery/cloudquery/commit/9e6dbd9182cad6bb48f5e0f7846e156b79d8621c))
* Fixed local policy run ([#535](https://github.com/cloudquery/cloudquery/issues/535)) ([3330b0b](https://github.com/cloudquery/cloudquery/commit/3330b0b809b8c90fda10d505a89fb8dc78fc0f67))
* Foreign Keys in Snapshots ([#525](https://github.com/cloudquery/cloudquery/issues/525)) ([a4e289c](https://github.com/cloudquery/cloudquery/commit/a4e289c95ed3ba7a2ed7c9b57063c63950dd3617))
* Getter implementation ([#450](https://github.com/cloudquery/cloudquery/issues/450)) ([5fcd775](https://github.com/cloudquery/cloudquery/commit/5fcd7756515489e66a21b485dceee7f8eea90925))
* Go Dependencies for `go-fsimpl` ([#595](https://github.com/cloudquery/cloudquery/issues/595)) ([e06fd0f](https://github.com/cloudquery/cloudquery/commit/e06fd0f1ebdf1ffe6a19e6eeee3aec9b332bce37))
* Handle if Path doesn't exist ([#518](https://github.com/cloudquery/cloudquery/issues/518)) ([d88e455](https://github.com/cloudquery/cloudquery/commit/d88e455f7beb9433a9add161240d414e998a0728))
* Handle panic level diags ([#579](https://github.com/cloudquery/cloudquery/issues/579)) ([b624b22](https://github.com/cloudquery/cloudquery/commit/b624b22084183995d707fa924ff466af1a6b10ca))
* Ignore warning/access diagnostics ([#590](https://github.com/cloudquery/cloudquery/issues/590)) ([fb5a144](https://github.com/cloudquery/cloudquery/commit/fb5a1447e0851d96af0fb7d522585bf78566ff3f))
* Include all responses in log, log invocation params ([#574](https://github.com/cloudquery/cloudquery/issues/574)) ([4661582](https://github.com/cloudquery/cloudquery/commit/466158252db8e39dd2f9074595a18d69e1f83350))
* Init command cleanup ([#502](https://github.com/cloudquery/cloudquery/issues/502)) ([a9ebc84](https://github.com/cloudquery/cloudquery/commit/a9ebc84f68f52cd765818c04ae99a23550cbe027))
* Keep binary reference consistent ([#516](https://github.com/cloudquery/cloudquery/issues/516)) ([9848cc0](https://github.com/cloudquery/cloudquery/commit/9848cc0d45fb26a54e14881adb9b560b1cade2f3))
* Make Date a static value ([#534](https://github.com/cloudquery/cloudquery/issues/534)) ([2cd1d3c](https://github.com/cloudquery/cloudquery/commit/2cd1d3c04a495f5946946236440ff7fafd6d4480))
* Max parallel resource limit ([#471](https://github.com/cloudquery/cloudquery/issues/471)) ([f090e87](https://github.com/cloudquery/cloudquery/commit/f090e8724bde5ca2cc337c1e4def32b0141ab51a))
* Missing views in history mode ([#508](https://github.com/cloudquery/cloudquery/issues/508)) ([da12335](https://github.com/cloudquery/cloudquery/commit/da123354c7a2b8ad6b9e1a708c78bc3a8b5edab6))
* Nil connection ([#442](https://github.com/cloudquery/cloudquery/issues/442)) ([d1fbaf1](https://github.com/cloudquery/cloudquery/commit/d1fbaf197aad5657a58c0daab9222ad6b977b10d))
* Panic On Policy Describe ([#438](https://github.com/cloudquery/cloudquery/issues/438)) ([27bbf11](https://github.com/cloudquery/cloudquery/commit/27bbf116b9f9282298459fffb6021f5fafb3dadb))
* panic on sync failure ([#676](https://github.com/cloudquery/cloudquery/issues/676)) ([27d574f](https://github.com/cloudquery/cloudquery/commit/27d574f6262417071c615675ec22b586317c50aa))
* Panic when canceled ([#443](https://github.com/cloudquery/cloudquery/issues/443)) ([117904a](https://github.com/cloudquery/cloudquery/commit/117904aeb7949e54d8753a381b446ecfdd5329f0))
* Policy Describe command  ([#466](https://github.com/cloudquery/cloudquery/issues/466)) ([fbb033f](https://github.com/cloudquery/cloudquery/commit/fbb033fff100af6cafb721796985f96c3e22d85d))
* Policy describe should run without config.hcl requirement ([#546](https://github.com/cloudquery/cloudquery/issues/546)) ([9bad3eb](https://github.com/cloudquery/cloudquery/commit/9bad3eb7fe2726d71afb9332b4b9299af75c3058))
* **policy:** Add missing GitHub getter ([#613](https://github.com/cloudquery/cloudquery/issues/613)) ([e3fc361](https://github.com/cloudquery/cloudquery/commit/e3fc361c12139c58de14e42ab7ba89f2a967508a))
* **policy:** Use firebase instead of GitHub API to get latest version ([#618](https://github.com/cloudquery/cloudquery/issues/618)) ([455ed23](https://github.com/cloudquery/cloudquery/commit/455ed23ca3f0d075028385359a47436b8b05ead9))
* Provider version check ([#582](https://github.com/cloudquery/cloudquery/issues/582)) ([acf79c3](https://github.com/cloudquery/cloudquery/commit/acf79c3c12e14df4e26f45aa0ce8906c4f690bcc))
* query --output should write file as a list of json objects ([5d15860](https://github.com/cloudquery/cloudquery/commit/5d15860c7d0d130066867e1b2ae7a76259f9810b))
* Refer to cloudquery as cloudquery ([#440](https://github.com/cloudquery/cloudquery/issues/440)) ([9d8b1a9](https://github.com/cloudquery/cloudquery/commit/9d8b1a981d3cf91622d47b9034b03a74ca569060))
* Relax tfstate version check ([#460](https://github.com/cloudquery/cloudquery/issues/460)) ([5c33bfc](https://github.com/cloudquery/cloudquery/commit/5c33bfc7630d570b86e7dbeb346bb619d4e05b8a))
* Remove enable_partial_fetch and support for migrationless providers ([#495](https://github.com/cloudquery/cloudquery/issues/495)) ([4457eb2](https://github.com/cloudquery/cloudquery/commit/4457eb25e6132981ac63acc1bb62afdcb48af8c9))
* Replace non-supported satori/go.uuid with google/uuid ([#576](https://github.com/cloudquery/cloudquery/issues/576)) ([d9170a5](https://github.com/cloudquery/cloudquery/commit/d9170a5140a2a8b0bc151fc15502607be103e070))
* Revert remote config ([#583](https://github.com/cloudquery/cloudquery/issues/583)) ([a411761](https://github.com/cloudquery/cloudquery/commit/a411761ad75d6c29e5048b1fc56feb075842b49d))
* skip lint for the weekend. ([f486d64](https://github.com/cloudquery/cloudquery/commit/f486d64c2df4ba10573765393a24e959cc245d72))
* Snapshot Persistance ([#500](https://github.com/cloudquery/cloudquery/issues/500)) ([d394e5c](https://github.com/cloudquery/cloudquery/commit/d394e5cc6b21c03e8e1351732be8bc1e32ac745a))
* Store fetch summary adjustements ([#426](https://github.com/cloudquery/cloudquery/issues/426)) ([49bd25a](https://github.com/cloudquery/cloudquery/commit/49bd25a954794b2404b52e891f35684d71c91966))
* Store fetch summary adjustements ([#426](https://github.com/cloudquery/cloudquery/issues/426)) ([d5b781e](https://github.com/cloudquery/cloudquery/commit/d5b781e8110bb85e06ed0a6e28d7c63969735739))
* Sync support optional provider args ([#642](https://github.com/cloudquery/cloudquery/issues/642)) ([5eac023](https://github.com/cloudquery/cloudquery/commit/5eac02321222f6a50b95308274cc631402ab213a))
* Typo in error message ([#441](https://github.com/cloudquery/cloudquery/issues/441)) ([166949f](https://github.com/cloudquery/cloudquery/commit/166949f7d5d40be6219d16c5944664906fb671ca))
* Update SDK, report squashed diagnostics ([#467](https://github.com/cloudquery/cloudquery/issues/467)) ([1f62bfb](https://github.com/cloudquery/cloudquery/commit/1f62bfb179553377c0815e7c025c5568a20ffca6))
* Updated the policy download command message ([#422](https://github.com/cloudquery/cloudquery/issues/422)) ([bac7369](https://github.com/cloudquery/cloudquery/commit/bac7369a629160fdefef7458e9ccb6799b3f84f5))
* Upgrade schema in history mode ([#494](https://github.com/cloudquery/cloudquery/issues/494)) ([d1e49ad](https://github.com/cloudquery/cloudquery/commit/d1e49adff490d1aca3e6d4f466bb3ddcb3e61f70))
* Upgrade/downgrade providers cleanup ([#498](https://github.com/cloudquery/cloudquery/issues/498)) ([bafcb15](https://github.com/cloudquery/cloudquery/commit/bafcb15d0c93ad191511fc937625869bce3d48c5))
* Use provider source ([#458](https://github.com/cloudquery/cloudquery/issues/458)) ([880edc5](https://github.com/cloudquery/cloudquery/commit/880edc5179d4da2cfc007f71cfb295bc8cd6ee2d))
* Use URI DSNs ([#482](https://github.com/cloudquery/cloudquery/issues/482)) ([dd8d0f3](https://github.com/cloudquery/cloudquery/commit/dd8d0f32063f16835b0cd9d072d129abdc06a13d))
* Validate DB connection with an explicit timeout, rather than the… ([#510](https://github.com/cloudquery/cloudquery/issues/510)) ([f5b0d04](https://github.com/cloudquery/cloudquery/commit/f5b0d049eda149f41b7b2f2ddddcf4b37c484d6a))
* Validate db version before proceeding ([#653](https://github.com/cloudquery/cloudquery/issues/653)) ([5af7f61](https://github.com/cloudquery/cloudquery/commit/5af7f615c580e94d319e2ad99b470ead9afd18f2))
* Validate provider semantic version ([#445](https://github.com/cloudquery/cloudquery/issues/445)) ([18256fa](https://github.com/cloudquery/cloudquery/commit/18256fa1b935012daf2049d3c7aea11bbec34c8b))
* Validate Timescale Version ([#540](https://github.com/cloudquery/cloudquery/issues/540)) ([1fad3b8](https://github.com/cloudquery/cloudquery/commit/1fad3b805a55bcb2be4d03ebe8474397a4b9b666))
* Warn on providers with no resources requested ([#528](https://github.com/cloudquery/cloudquery/issues/528)) ([5b67635](https://github.com/cloudquery/cloudquery/commit/5b67635a4abaf09ee06a22d592689a3661686077)), closes [#514](https://github.com/cloudquery/cloudquery/issues/514)


### Miscellaneous Chores

* Release 0.23.0 ([#674](https://github.com/cloudquery/cloudquery/issues/674)) ([d4a2502](https://github.com/cloudquery/cloudquery/commit/d4a250288832b28104ae7e5497fbe6dc9a8f1231))

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

## [Unreleased] 

### :rocket: Added
* Added core migrations implementation
* Added fetch summary saving to `fetches` table


<!-- 
## Unreleased
### Added
### Changed
### Fixed
### Breaking Changes
-->

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
