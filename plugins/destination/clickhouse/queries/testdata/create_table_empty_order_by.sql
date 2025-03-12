CREATE TABLE IF NOT EXISTS `table_name` (
  `extra_col` Nullable(Float64),
  `extra_inet_col` Nullable(String),
  `extra_inet_arr_col` Array(Nullable(String))
) ENGINE = MergeTree() ORDER BY tuple() SETTINGS allow_nullable_key=1