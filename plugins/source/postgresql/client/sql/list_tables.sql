-- Extract the list of tables with columns, types and constraints. Columns with multiple constraints will be
-- listed multiple times.
--
-- +----------------+-------------+-------------+------------+----------------+-----------+-----------+---------------------+
-- | ordinal_position | table_name | column_name | data_type | is_primary_key | not_null  | is_unique | constraint_name     |
-- +----------------+-------------+-------------+------------+----------------+-----------+-----------+---------------------+
-- |              1 | users       | id          | bigint     | YES            | true      | true      | cq_users_pk         |
-- |              1 | users       | id          | bigint     | NO             | true      | true      | extra_constraint    |
-- |              2 | users       | name        | text       | NO             | false     | false     |                     |
-- |              3 | users       | email       | text       | NO             | true      | false     | cq_users_pk         |
-- |              1 | posts       | id          | bigint     | YES            | true      | true      | cq_posts_pk         |
-- |              2 | posts       | title       | text       | NO             | false     | false     |                     |
WITH base_query AS (SELECT
    columns.ordinal_position AS ordinal_position,
    pg_class.relname AS table_name,
    pg_attribute.attname AS column_name,
    CASE
        -- This is required per the differences in pg_catalog.format_type implementations
        -- between PostgreSQL & CockroachDB.
        -- namely, numeric(20,0)[] is returned as numeric[] unless we use the typelem format + []
        WHEN pg_type.typcategory = 'A' AND pg_type.typelem != 0
            THEN pg_catalog.format_type(pg_type.typelem, pg_attribute.atttypmod) || '[]'
        ELSE pg_catalog.format_type(pg_attribute.atttypid, pg_attribute.atttypmod)
        END AS data_type,
    CASE
        WHEN conkey IS NOT NULL AND contype = 'p' AND array_position(conkey, pg_attribute.attnum) > 0 THEN true
        ELSE false
        END AS is_primary_key,
    CASE
        WHEN pg_attribute.attnotnull THEN true
        ELSE false
        END AS not_null,
    CASE
        WHEN
                conkey IS NOT NULL
                AND (contype = 'p' OR contype = 'u')
                AND array_length(conkey, 1) = 1  -- we don't handle composite unique keys
                AND array_position(conkey, pg_attribute.attnum) > 0
            THEN true
        ELSE false
        END AS is_unique,
    COALESCE(pg_constraint.conname, '') AS constraint_name
FROM
    pg_catalog.pg_attribute
        INNER JOIN
    pg_catalog.pg_type ON pg_type.oid = pg_attribute.atttypid
        INNER JOIN
    pg_catalog.pg_class ON pg_class.oid = pg_attribute.attrelid
        INNER JOIN
    pg_catalog.pg_namespace ON pg_namespace.oid = pg_class.relnamespace
        LEFT JOIN
    pg_catalog.pg_constraint ON pg_constraint.conrelid = pg_attribute.attrelid
        AND conkey IS NOT NULL AND array_position(conkey, pg_attribute.attnum) > 0
        AND (contype = 'p' OR contype = 'u')
        INNER JOIN
    information_schema.columns ON columns.table_name = pg_class.relname AND columns.column_name = pg_attribute.attname AND columns.table_schema = pg_catalog.pg_namespace.nspname
WHERE
        pg_attribute.attnum > 0
  AND NOT pg_attribute.attisdropped
  AND pg_catalog.pg_namespace.nspname = '%s')

-- This is the outer query that aggregates the results of the base query, removing duplicate rows due to multiple constraints.
-- The constraint_name selects the primary key constraint or inserts an empty string if a column is not due to a primary key.
SELECT
    bq.ordinal_position as ordinal_position,
    bq.table_name as table_name,
    bq.column_name as column_name,
    bq.data_type as data_type,
    bool_or(bq.is_primary_key) as is_primary_key,
    bool_or(bq.not_null) as not_null,
    bool_or(bq.is_unique) as is_unique,
    coalesce((
        SELECT
            constraint_name
        FROM base_query
        WHERE
            ordinal_position = bq.ordinal_position
            AND table_name = bq.table_name
            AND column_name = bq.column_name
            AND data_type = bq.data_type
            AND is_primary_key = TRUE
        ),'') AS constraint_name
FROM
    base_query AS bq
GROUP BY
    ordinal_position,
    table_name,
    column_name,
    data_type
ORDER BY
    table_name, ordinal_position;
