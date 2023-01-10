DROP VIEW IF EXISTS azure_resources;
DO $$
DECLARE
    tbl TEXT;
    strSQL TEXT = '';
BEGIN
-- iterate over every table in our information_schema that has an `subscription_id` column available
FOR tbl IN
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'azure_%s' AND COLUMN_NAME = 'subscription_id'
    INTERSECT
    -- iterate over every table in our information_schema that has an `id` column available
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'azure_%s' AND COLUMN_NAME = 'id'
LOOP
    -- UNION each table query to create one view
    IF NOT (strSQL = ''::TEXT) THEN
        strSQL = strSQL || ' UNION ALL ';
    END IF;
    -- create an SQL query to select from table and transform it into our resources view schema
    -- we use the double reverse here because split_part with negative indexes is not available in PostgreSQL < 14; https://pgpedia.info/postgresql-versions/postgresql-14.html#system_function_changes
    strSQL = strSQL || format('
        SELECT _cq_id, _cq_source_name, _cq_sync_time, %L as _cq_table, subscription_id, reverse(split_part(reverse(id), ''/''::TEXT, 1)) as id,
        %s as name, %s as kind, %s as location, id AS full_id
        FROM %s',
        tbl,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='name' AND table_name=tbl) THEN 'name' ELSE 'NULL' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='kind' AND table_name=tbl) THEN 'kind' ELSE 'NULL' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='location' AND table_name=tbl) THEN 'location' ELSE E'\'unavailable\'' END,
        tbl);

END LOOP;

IF strSQL = ''::TEXT THEN
    RAISE EXCEPTION 'No tables found with ID and SUBSCRIPTION_ID columns. Run a sync first and try again.';
ELSE
    EXECUTE FORMAT('CREATE VIEW azure_resources AS (%s)', strSQL);
END IF;

END $$;
