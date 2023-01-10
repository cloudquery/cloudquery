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
        strSQL = strSQL || format('SELECT _cq_id, _cq_source_name, _cq_sync_time, %L AS _cq_table, context, uid FROM %s', tbl, tbl);
    END LOOP;
    IF strSQL = ''::TEXT THEN
        RAISE EXCEPTION 'No tables found with UID and CONTEXT columns. Run a sync first and try again.';
    ELSE
        EXECUTE FORMAT('CREATE VIEW k8s_resources AS (%s)', strSQL);
    END IF;
END $$;
