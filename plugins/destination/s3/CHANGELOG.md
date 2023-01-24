# Changelog

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
