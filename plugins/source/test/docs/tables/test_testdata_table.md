# Table: test_testdata_table

This table shows data for Test Testdata Table.

Testdata table

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|test_cq_source_name|`utf8`|
|test_cq_sync_time|`timestamp[us, tz=UTC]`|
|test_cq_id|`uuid`|
|test_cq_parent_id|`uuid`|
|binary|`binary`|
|boolean|`bool`|
|date32|`date32`|
|date64|`date64`|
|daytimeinterval|`day_time_interval`|
|duration_ms|`duration[ms]`|
|duration_ns|`duration[ns]`|
|duration_s|`duration[s]`|
|duration_us|`duration[us]`|
|float32|`float32`|
|float64|`float64`|
|inet|`inet`|
|int16|`int16`|
|int32|`int32`|
|int64|`int64`|
|int8|`int8`|
|largebinary|`large_binary`|
|largestring|`large_utf8`|
|mac|`mac`|
|monthdaynanointerval|`month_day_nano_interval`|
|monthinterval|`month_interval`|
|string|`utf8`|
|time32ms|`time32[ms]`|
|time32s|`time32[s]`|
|time64ns|`time64[ns]`|
|time64us|`time64[us]`|
|timestamp_ms|`timestamp[ms, tz=UTC]`|
|timestamp_ns|`timestamp[ns, tz=UTC]`|
|timestamp_s|`timestamp[s, tz=UTC]`|
|timestamp_us|`timestamp[us, tz=UTC]`|
|uint16|`uint16`|
|uint32|`uint32`|
|uint64|`uint64`|
|uint8|`uint8`|
|uuid|`uuid`|
|json|`json`|
|json_array|`json`|
|boolean_list|`list<item: bool, nullable>`|
|date32_list|`list<item: date32, nullable>`|
|date64_list|`list<item: date64, nullable>`|
|daytimeinterval_list|`list<item: day_time_interval, nullable>`|
|duration_ms_list|`list<item: duration[ms], nullable>`|
|duration_ns_list|`list<item: duration[ns], nullable>`|
|duration_s_list|`list<item: duration[s], nullable>`|
|duration_us_list|`list<item: duration[us], nullable>`|
|float32_list|`list<item: float32, nullable>`|
|float64_list|`list<item: float64, nullable>`|
|inet_list|`list<item: inet, nullable>`|
|int16_list|`list<item: int16, nullable>`|
|int32_list|`list<item: int32, nullable>`|
|int64_list|`list<item: int64, nullable>`|
|int8_list|`list<item: int8, nullable>`|
|largestring_list|`list<item: large_utf8, nullable>`|
|mac_list|`list<item: mac, nullable>`|
|monthdaynanointerval_list|`list<item: month_day_nano_interval, nullable>`|
|monthinterval_list|`list<item: month_interval, nullable>`|
|string_list|`list<item: utf8, nullable>`|
|time32ms_list|`list<item: time32[ms], nullable>`|
|time32s_list|`list<item: time32[s], nullable>`|
|time64ns_list|`list<item: time64[ns], nullable>`|
|time64us_list|`list<item: time64[us], nullable>`|
|timestamp_ms_list|`list<item: timestamp[ms, tz=UTC], nullable>`|
|timestamp_ns_list|`list<item: timestamp[ns, tz=UTC], nullable>`|
|timestamp_s_list|`list<item: timestamp[s, tz=UTC], nullable>`|
|timestamp_us_list|`list<item: timestamp[us, tz=UTC], nullable>`|
|uint16_list|`list<item: uint16, nullable>`|
|uint32_list|`list<item: uint32, nullable>`|
|uint64_list|`list<item: uint64, nullable>`|
|uint8_list|`list<item: uint8, nullable>`|
|uuid_list|`list<item: uuid, nullable>`|
|struct|`struct<binary: binary, boolean: bool, date32: date32, date64: date64, daytimeinterval: day_time_interval, duration_ms: duration[ms], duration_ns: duration[ns], duration_s: duration[s], duration_us: duration[us], float32: float32, float64: float64, inet: inet, int16: int16, int32: int32, int64: int64, int8: int8, largebinary: large_binary, largestring: large_utf8, mac: mac, monthdaynanointerval: month_day_nano_interval, monthinterval: month_interval, string: utf8, time32ms: time32[ms], time32s: time32[s], time64ns: time64[ns], time64us: time64[us], timestamp_ms: timestamp[ms, tz=UTC], timestamp_ns: timestamp[ns, tz=UTC], timestamp_s: timestamp[s, tz=UTC], timestamp_us: timestamp[us, tz=UTC], uint16: uint16, uint32: uint32, uint64: uint64, uint8: uint8, uuid: uuid, json: json, json_array: json>`|
|nested_struct|`struct<inner: struct<binary: binary, boolean: bool, date32: date32, date64: date64, daytimeinterval: day_time_interval, duration_ms: duration[ms], duration_ns: duration[ns], duration_s: duration[s], duration_us: duration[us], float32: float32, float64: float64, inet: inet, int16: int16, int32: int32, int64: int64, int8: int8, largebinary: large_binary, largestring: large_utf8, mac: mac, monthdaynanointerval: month_day_nano_interval, monthinterval: month_interval, string: utf8, time32ms: time32[ms], time32s: time32[s], time64ns: time64[ns], time64us: time64[us], timestamp_ms: timestamp[ms, tz=UTC], timestamp_ns: timestamp[ns, tz=UTC], timestamp_s: timestamp[s, tz=UTC], timestamp_us: timestamp[us, tz=UTC], uint16: uint16, uint32: uint32, uint64: uint64, uint8: uint8, uuid: uuid, json: json, json_array: json>>`|