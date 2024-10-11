with aws_tables as (
    select table_name
    from information_schema.columns
    where table_name like 'aws_%s'
        and column_name = 'account_id'
    intersect
    select table_name
    from information_schema.columns
    where table_name like 'aws_%s'
        and column_name = 'arn'
),

select_statements as (
    select concat('select _cq_id, _cq_source_name, _cq_sync_time, ',
        '''',
        table_name,
        '''',
        ' as _cq_table, account_id, ',
        case
            when
                exists (
                    select 1
                    from
                        information_schema.columns
                    where
                        information_schema.columns.column_name = 'region'
                        and information_schema.columns.table_name = aws_tables.table_name
                ) then 'region'
            else concat('''', 'unavailable', '''')
        end,
        ' as region, arn, ',
        case
            when
                exists (
                    select 1
                    from
                        information_schema.columns
                    where
                        information_schema.columns.column_name = 'tags'
                        and information_schema.columns.table_name = aws_tables.table_name
                ) then 'tags'
            else concat('''', '{}', '''')
        end,
        ' as tags from ',
        table_name
    ) as select_statement
    from aws_tables
)

select concat('create or replace view aws_resources as (',
    array_join(array_agg(select_statement), ' union all '),
    ');') as create_view_statement
from select_statements;
