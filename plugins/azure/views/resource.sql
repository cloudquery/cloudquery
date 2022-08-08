DROP VIEW IF EXISTS azure_resources;
DO $$
DECLARE
    tbl text;
    strSQL text = '';
BEGIN
    -- iterate over every table in our information_schema that has an `subscription_id` column available
    FOR tbl IN SELECT table_name FROM information_schema.columns WHERE table_name LIKE 'azure_%s' AND COLUMN_NAME = 'subscription_id'
        INTERSECT
    -- iterate over every table in our information_schema that has an `id` column available
    SELECT table_name FROM information_schema.columns WHERE table_name LIKE 'azure_%s' AND COLUMN_NAME = 'id'
LOOP
    -- UNION each table query to create one view
    IF NOT (strSQL = ''::text) THEN
        strSQL = strSQL || ' UNION ALL ';
    END IF;
    -- create an SQL query to select from table and transform it into our resources view schema
    -- we use the double reverse here because split_part with negative indexes is not available in PostgreSQL < 14; https://pgpedia.info/postgresql-versions/postgresql-14.html#system_function_changes
    strSQL = strSQL || format('
        SELECT cq_id, cq_meta, %L as cq_table, subscription_id, reverse(split_part(reverse(id), ''/''::text, 1)) as id,
        %s as name, %s as kind, %s as location,
        COALESCE(%s, (cq_meta->>''last_updated'')::timestamp) as fetch_date
        FROM %s', tbl,
            CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='name' AND table_name=tbl) THEN 'name' ELSE 'NULL' END,
            CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='kind' AND table_name=tbl) THEN 'kind' ELSE 'NULL' END,
            CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='location' AND table_name=tbl) THEN 'location' ELSE E'\'unavailable\'' END,
            CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='fetch_date' AND table_name=tbl) THEN 'fetch_date' ELSE 'NULL::timestamp' END,
    tbl);

END LOOP;
    EXECUTE format('CREATE VIEW azure_resources AS (%s)', strSQL);
END $$;