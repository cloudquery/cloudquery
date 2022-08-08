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
	 strSQL = strSQL || format('select  cq_id,  cq_meta, %L as cq_table, project_id, %s as region, id, %s as name, %s as description,
		  					    COALESCE(%s, (cq_meta->>''last_updated'')::timestamp) as fetch_date
							   FROM %s', tbl,
							   CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='region' AND table_name=tbl) THEN 'region' ELSE E'\'unavailable\'' END,
                               CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='name' AND table_name=tbl) THEN 'name' ELSE 'NULL' END,
                               CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='description' AND table_name=tbl) THEN 'description' ELSE 'NULL' END,
							   CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='fetch_date' AND table_name=tbl) THEN 'fetch_date' ELSE 'NULL::timestamp' END,
							   tbl);

END LOOP;
execute format('CREATE VIEW gcp_resources AS (%s)', strSQL);

end $$;