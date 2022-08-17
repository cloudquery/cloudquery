DROP VIEW IF EXISTS aws_resources;

DO $$
DECLARE
    tbl TEXT;
    strSQL TEXT = '';
BEGIN
-- iterate over every table in our information_schema that has an `arn` column available
FOR tbl IN
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME = 'account_id'
    INTERSECT
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME = 'arn'
LOOP 
    -- UNION each table query to create one view
 	IF NOT (strSQL = ''::TEXT) THEN
		strSQL = strSQL || ' UNION ALL ';
	END IF;
	-- create an SQL query to select from table and transform it into our resources view schema
	strSQL = strSQL || FORMAT('
        select  cq_id,  cq_meta, %L as cq_table, account_id, %s as region, arn, %s as tags,
        COALESCE(%s, (cq_meta->>''last_updated'')::timestamp) as fetch_date
        FROM %s',
        tbl,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='region' AND table_name=tbl) THEN 'region' ELSE E'\'unavailable\'' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='tags' AND table_name=tbl) THEN 'tags' ELSE '''{}''::jsonb' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='fetch_date' AND table_name=tbl) THEN 'fetch_date' ELSE 'NULL::timestamp' END,
        tbl);
END LOOP;

IF strSQL = ''::TEXT THEN
    RAISE EXCEPTION 'No tables found with ARN and ACCOUNT_ID columns. Run a fetch first and try again.';
ELSE
	EXECUTE FORMAT('CREATE VIEW aws_resources AS (%s)', strSQL);
END IF;

END $$;
