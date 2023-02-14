SELECT `_cq_id`, `_cq_parent_id`, `_cq_source_name`, `_cq_sync_time`, `extra_col`
FROM `table_name`
WHERE `_cq_source_name` = @sourceName
ORDER BY `_cq_sync_time`