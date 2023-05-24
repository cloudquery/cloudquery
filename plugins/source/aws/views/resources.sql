DROP VIEW IF EXISTS aws_resources;

DO $$
DECLARE
    tbl TEXT;
    strSQL TEXT = '';
BEGIN
-- iterate over every table in our information_schema that has an `arn` column available
FOR tbl IN
    SELECT DISTINCT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME IN ('account_id', 'request_account_id')
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
	strSQL = strSQL || FORMAT(E'
        SELECT _cq_id,
            _cq_source_name,
            _cq_sync_time,
            %L AS _cq_table,
            COALESCE(%s, SPLIT_PART(arn, \':\', 5)) AS account_id,
            COALESCE(%s, %s, SPLIT_PART(arn, \':\', 5)) AS request_account_id,
            %s AS region,
            SPLIT_PART(arn, \':\', 2) AS PARTITION,
            SPLIT_PART(arn, \':\', 3) AS service,
            CASE
            WHEN SPLIT_PART(SPLIT_PART(ARN, \':\', 6), \'/\', 2) = \'\' AND SPLIT_PART(arn, \':\', 7) = \'\' THEN NULL
                ELSE SPLIT_PART(SPLIT_PART(arn, \':\', 6), \'/\', 1)
            END AS TYPE,
            arn, %s AS tags
        FROM %s',
        tbl,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='account_id' AND table_name=tbl) THEN 'account_id' ELSE 'NULL' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='request_account_id' AND table_name=tbl) THEN 'request_account_id' ELSE 'NULL' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='account_id' AND table_name=tbl) THEN 'account_id' ELSE 'NULL' END,
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
