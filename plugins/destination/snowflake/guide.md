Create Stage:

JSON

```
create or replace file format myjsonformat
  type = 'JSON'
  strip_outer_array = true;

create or replace stage my_json_stage
  file_format = myjsonformat;

put file:///tmp/snowflakejson/* @my_json_stage auto_compress=true;

list @my_json_stage;

copy into "gcp_compute_instances"
  from @my_json_stage1
  pattern='.*gcp_compute_instances\.json.*'
  file_format = (format_name = myjsonformat)
  on_error = 'skip_file';
```