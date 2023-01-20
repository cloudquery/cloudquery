# Changelog

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-gitlab-v2.0.0...plugins-source-gitlab-v2.1.0) (2023-01-17)


### Features

* **gitlab-saas:** Pass minimal access level on GitLab SaaS ([#6852](https://github.com/cloudquery/cloudquery/issues/6852)) ([1316fda](https://github.com/cloudquery/cloudquery/commit/1316fda156b5695cf282e60891161084c750243d))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-gitlab-v1.1.0...plugins-source-gitlab-v2.0.0) (2023-01-10)


### ⚠ BREAKING CHANGES

* **deps:** `gitlab_users` columns `last_sign_in_ip, current_sign_in_ip` type changed from `IntArray` to `Inet`

### Features

* **gitlab:** Migrate codegen to transformation ([#6429](https://github.com/cloudquery/cloudquery/issues/6429)) ([34d9f1f](https://github.com/cloudquery/cloudquery/commit/34d9f1fa6fb3d6eb0aeaa0a86ad2bfa5b8b08994))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([b821650](https://github.com/cloudquery/cloudquery/commit/b821650a9d4e97aa7eb8fa0fe210c835cea4146b))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-gitlab-v1.0.1...plugins-source-gitlab-v1.1.0) (2023-01-03)


### Features

* **gitlab:** Add Branches ([#6135](https://github.com/cloudquery/cloudquery/issues/6135)) ([31185dd](https://github.com/cloudquery/cloudquery/commit/31185ddcafbe50918dfd96ce289c030908afee26))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.0 ([#6256](https://github.com/cloudquery/cloudquery/issues/6256)) ([b19f6cd](https://github.com/cloudquery/cloudquery/commit/b19f6cd8e2c39994aeb19d78e78e927d6c3cf580))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-gitlab-v1.0.0...plugins-source-gitlab-v1.0.1) (2022-12-27)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.13.1 ([#5897](https://github.com/cloudquery/cloudquery/issues/5897)) ([ad15915](https://github.com/cloudquery/cloudquery/commit/ad15915f2951a75729859f6f1377ed789f8ba115))

## 1.0.0 (2022-12-23)


### Features

* New Gitlab (Self Hosted) Plugin ([#5222](https://github.com/cloudquery/cloudquery/issues/5222)) ([f45cb0a](https://github.com/cloudquery/cloudquery/commit/f45cb0abb59cae30787e531395dd3be2a227f07c))
