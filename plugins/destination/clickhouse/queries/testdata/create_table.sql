CREATE TABLE `table_name` (
  `_cq_id` UUID,
  `_cq_parent_id` Nullable(UUID),
  `_cq_source_name` Nullable(String),
  `_cq_sync_time` Nullable(DateTime64(9)),
  `extra_col` Float64,
  `extra_inet_col` Nullable(String),
  `extra_inet_arr_col` Array(Nullable(String))
) ENGINE = MergeTree ORDER BY (`_cq_id`, `extra_col`)