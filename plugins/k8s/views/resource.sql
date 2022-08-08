DROP VIEW IF EXISTS k8s_resources;
DO $$
DECLARE
  tbl TEXT;
  strSQL TEXT = '';
BEGIN
    -- iterate over every table in our information_schema that has `uid` and `context` columns available
    FOR tbl IN
        SELECT table_name
        FROM information_schema.columns
        WHERE table_name LIKE 'k8s_%s' AND COLUMN_NAME = 'context'
        INTERSECT
        SELECT table_name
        FROM information_schema.columns
        WHERE table_name LIKE 'k8s_%s' AND COLUMN_NAME = 'uid'
    LOOP
        -- UNION each table query to create one view
 	    IF NOT (strSQL = ''::TEXT) THEN
	        strSQL = strSQL || ' UNION ALL ';
	    END IF;
	    -- create an SQL query to select from table and transform it into our resources view schema
	    strSQL = strSQL || format('SELECT cq_id, cq_meta, %L AS cq_table, context, uid,
		  					       COALESCE(%s, (cq_meta->>''last_updated'')::timestamp) AS fetch_date
							       FROM %s',
                                   tbl,
							       CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='fetch_date' AND table_name=tbl) THEN 'fetch_date' ELSE 'NULL::timestamp' END,
							       tbl);
    END LOOP;
    EXECUTE FORMAT('CREATE VIEW k8s_resources AS (%s)', strSQL);
END $$;
