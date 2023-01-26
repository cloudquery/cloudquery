# Changelog

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.0.0...plugins-destination-mssql-v2.1.0) (2023-01-24)


### Features

* **mssql:** Use `LEFT JOIN` for overwrite ([#7086](https://github.com/cloudquery/cloudquery/issues/7086)) ([97c7981](https://github.com/cloudquery/cloudquery/commit/97c79816850b62989670916fd93f4f52528ce2b3))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v1.0.0...plugins-destination-mssql-v2.0.0) (2023-01-23)


### ⚠ BREAKING CHANGES

* **mssql:** Change column types from `varchar` to `nvarchar` to properly store Unicode characters

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))
* **mssql:** Change column types from `varchar` to `nvarchar` to properly store Unicode characters ([748c7df](https://github.com/cloudquery/cloudquery/commit/748c7df498dfd7ff342c3e380af1ee75f5d1034b))
* **mssql:** Fix JSON transformer to store unescaped data ([748c7df](https://github.com/cloudquery/cloudquery/commit/748c7df498dfd7ff342c3e380af1ee75f5d1034b))

## 1.0.0 (2023-01-19)


### Features

* **mssql:** Microsoft SQL Server destination ([#6417](https://github.com/cloudquery/cloudquery/issues/6417)) ([c71bc88](https://github.com/cloudquery/cloudquery/commit/c71bc887280c9efb05afb98321a611bc33e387b0))
