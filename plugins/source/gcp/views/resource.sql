DROP VIEW IF EXISTS gcp_resources;
do $$
declare
  tbl text;
  strSQL text = '';
begin
 -- iterate over every table in our information_schema that has an `arn` column available
FOR tbl IN SELECT table_name from information_schema.columns where table_name like 'gcp_%s' and COLUMN_NAME = 'project_id'
           intersect
select table_name from information_schema.columns where table_name like 'gcp_%s' and COLUMN_NAME = 'id'
 LOOP 
     -- UNION each table query to create one view
 	 IF NOT (strSQL = ''::text) THEN
	      strSQL = strSQL || ' UNION ALL ';
	 END IF;
	 -- create an SQL query to select from table and transform it into our resources view schema
	 strSQL = strSQL || format('select _cq_id, _cq_source_name, _cq_sync_time, %L as _cq_table, project_id, %s as region, id::text as id, %s as name, %s as description
							   FROM %s', tbl,
							   CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='region' AND table_name=tbl) THEN 'region' ELSE E'\'unavailable\'' END,
                               CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='name' AND table_name=tbl) THEN 'name' ELSE 'NULL' END,
                               CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='description' AND table_name=tbl) THEN 'description' ELSE 'NULL' END,
							   tbl);

END LOOP;
execute format('CREATE VIEW gcp_resources AS (%s)', strSQL);

end $$;