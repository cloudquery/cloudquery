DROP VIEW IF EXISTS gcp_resources;
DO $$
DECLARE
    tbl TEXT;
    strSQL TEXT = '';
BEGIN
-- iterate over every table in our information_schema that has an `arn` column available
FOR tbl IN
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'gcp_%s' AND column_name = 'project_id'
    INTERSECT
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'gcp_%s' AND COLUMN_NAME = 'id'
LOOP 
    -- UNION each table query to create one view
    IF NOT (strSQL = ''::TEXT) THEN
        strSQL = strSQL || ' UNION ALL ';
    END IF;
    -- create an SQL query to select from table and transform it into our resources view schema
    strSQL = strSQL || FORMAT('
        SELECT cq_id, cq_meta, %L AS cq_table, project_id, %s AS region, id, %s AS name, %s AS description,
        COALESCE(%s, (cq_meta->>''last_updated'')::timestamp) AS fetch_date
        FROM %s',
        tbl,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='region' AND table_name=tbl) THEN 'region' ELSE E'\'unavailable\'' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='name' AND table_name=tbl) THEN 'name' ELSE 'NULL' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='description' AND table_name=tbl) THEN 'description' ELSE 'NULL' END,
        CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns WHERE column_name='fetch_date' AND table_name=tbl) THEN 'fetch_date' ELSE 'NULL::timestamp' END,
        tbl);

END LOOP;

IF strSQL = ''::TEXT THEN
    RAISE EXCEPTION 'No tables found with ID and PROJECT_ID columns. Run a fetch first and try again.';
ELSE
    EXECUTE FORMAT('CREATE VIEW gcp_resources AS (%s)', strSQL);
END IF;

END $$;
