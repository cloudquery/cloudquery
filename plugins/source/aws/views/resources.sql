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
        select _cq_id, _cq_source_name, _cq_sync_time, %L as _cq_table, account_id, %s as region, arn, %s as tags
        FROM %s',
        tbl,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='region' AND table_name=tbl) THEN 'region' ELSE E'\'unavailable\'' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='tags' AND table_name=tbl) THEN 'tags' ELSE '''{}''::jsonb' END,
        tbl);
END LOOP;

IF strSQL = ''::TEXT THEN
    RAISE EXCEPTION 'No tables found with ARN and ACCOUNT_ID columns. Run a sync first and try again.';
ELSE
	EXECUTE FORMAT('CREATE VIEW aws_resources AS (%s)', strSQL);
END IF;

END $$;
