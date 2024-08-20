import sqlite3
from typing import List, Any, Tuple, Optional
from cloudquery.sdk.schema import Table, Column, TableColumnChange, TableColumnChangeType, get_table_changes, flatten_tables
import pyarrow as pa
from cloudquery.sdk.message import WriteMigrateTableMessage
from plugin.sqlite.type_conversions import arrow_type_to_sqlite_str, arrow_type_to_sqlite, sqlite_type_to_arrow_type

SQL_TABLE_INFO = "PRAGMA table_info('{}');"

class ColumnInfo:
    def __init__(self, index: int, name: str, typ: str, not_null: bool, default_value: Any, pk: int):
        self.index = index
        self.name = name
        self.typ = typ.lower()
        self.not_null = not_null
        self.default_value = default_value
        self.pk = pk

class TableInfo:
    def __init__(self):
        self.columns: List[ColumnInfo] = []

class MigrateSQLClient:
    def __init__(self, db: sqlite3.Connection):
        self.db = db

    def identifier(self, string: str) -> str:
        return f'"{string}"'

    def sqlite_tables(self, tables: List[Table]) -> List[Table]:
        schema_tables = []
        for table in tables:
            columns = []
            info = self.get_table_info(table.name)
            if info is None:
                continue
            for col in info.columns:
                columns.append(Column(
                    name=col.name,
                    type=sqlite_type_to_arrow_type(col.typ),
                    primary_key=col.pk != 0,
                    not_null=col.not_null
                ))
            schema_tables.append(Table(name=table.name, columns=columns))
        return schema_tables

    def normalize_tables(self, tables: List[Table]) -> List[Table]:
        flattened = flatten_tables(tables)
        normalized = [self.normalize_table(table) for table in flattened]
        return normalized

    def normalize_table(self, table: Table) -> Table:
        columns = [self.normalize_field(col.to_arrow_field()) for col in table.columns]
        columns = [Column.from_arrow_field(col) for col in columns]
        return Table(name=table.name, columns=columns)

    def normalize_field(self, field: pa.Field) -> pa.Field:
        return pa.field(
            name=field.name,
            type=arrow_type_to_sqlite(field.type),
            nullable=field.nullable,
            metadata=field.metadata
        )

    def non_auto_migratable_tables(self, tables: List[Table], sqlite_tables: List[Table]) -> Tuple[List[str], List[List[TableColumnChange]]]:
        result = []
        table_changes = []
        for t in tables:
            sqlite_table = self.get_table_by_name(sqlite_tables, t.name)
            if sqlite_table is None:
                continue
            changes = get_table_changes(sqlite_table, t)
            if not self.can_auto_migrate(changes):
                result.append(t.name)
                table_changes.append(changes)
        return result, table_changes

    def auto_migrate_table(self, table: Table, changes: List[TableColumnChange]) -> Optional[Exception]:
        for change in changes:
            if change.type == TableColumnChangeType.ADD:
                err = self.add_column(table.name, change.current.name, arrow_type_to_sqlite_str(change.current.type))
                if err:
                    return err
        return None

    def can_auto_migrate(self, changes: List[TableColumnChange]) -> bool:
        for change in changes:
            if change.type == TableColumnChangeType.ADD:
                if change.current.primary_key or change.current.not_null:
                    return False
            elif change.type == TableColumnChangeType.REMOVE:
                if change.previous.primary_key or change.previous.not_null:
                    return False
            elif change.type != TableColumnChangeType.REMOVE_UNIQUE_CONSTRAINT:
                return False
        return True

    def migrate_tables(self, msgs: List[WriteMigrateTableMessage]):
        for msg in msgs:
            force = msg.migrate_force
            tables = [msg.table]
            normalized_tables = self.normalize_tables(tables)
            sqlite_tables = self.sqlite_tables(normalized_tables)

            if not force:
                non_auto_migratable_tables, changes = self.non_auto_migratable_tables(normalized_tables, sqlite_tables)
                if non_auto_migratable_tables:
                    return Exception(f"Tables {', '.join(non_auto_migratable_tables)} with changes {changes} require migration. Migrate manually or consider using 'migrate_mode: forced'")

            for table in normalized_tables:
                if len(table.columns) == 0:
                    continue

                sqlite_table = self.get_table_by_name(sqlite_tables, table.name)
                if sqlite_table is None:
                    err = self.create_table_if_not_exist(table)
                    if err:
                        raise err
                else:
                    changes = get_table_changes(table, sqlite_table)
                    if self.can_auto_migrate(changes):
                        err = self.auto_migrate_table(table, changes)
                        if err:
                            raise err
                    else:
                        err = self.recreate_table(table)
                        if err:
                            raise err
        return None

    def recreate_table(self, table: Table) -> Optional[Exception]:
        sql = f"DROP TABLE IF EXISTS {self.identifier(table.name)}"
        try:
            self.db.execute(sql)
        except sqlite3.Error as err:
            return Exception(f"Failed to drop table {table.name}: {err}")
        return self.create_table_if_not_exist(table)

    def add_column(self, table_name: str, column_name: str, column_type: str) -> Optional[Exception]:
        sql = f"ALTER TABLE {self.identifier(table_name)} ADD COLUMN {self.identifier(column_name)} {self.identifier(column_type)}"
        try:
            self.db.execute(sql)
        except sqlite3.Error as err:
            return Exception(f"Failed to add column {column_name} on table {table_name}: {err}")
        return None

    def create_table_if_not_exist(self, table: Table) -> Optional[Exception]:
        sb = []

        sb.append(f"CREATE TABLE IF NOT EXISTS {self.identifier(table.name)} (")
        primary_keys = []
        for i, col in enumerate(table.columns):
            sql_type = arrow_type_to_sqlite_str(col.type)
            if not sql_type:
                continue
            field_def = f"{self.identifier(col.name)} {sql_type}"
            if col.not_null:
                field_def += " NOT NULL"
            sb.append(field_def)
            if col.primary_key:
                primary_keys.append(self.identifier(col.name))
            if i < len(table.columns) - 1:
                sb.append(", ")

        if primary_keys:
            sb.append(f", CONSTRAINT {self.identifier(table.name + '_cqpk')} PRIMARY KEY ({', '.join(primary_keys)})")
        sb.append(")")

        try:
            self.db.execute("".join(sb))
        except sqlite3.Error as err:
            return Exception(f"Failed to create table with '{''.join(sb)}': {err}")
        return None

    def get_table_info(self, table_name: str) -> Optional[TableInfo]:
        info = TableInfo()
        cursor = self.db.cursor()
        try:
            cursor.execute(SQL_TABLE_INFO.format(table_name))
            for row in cursor.fetchall():
                col_info = ColumnInfo(*row)
                info.columns.append(col_info)
        except sqlite3.Error as err:
            raise err
        finally:
            cursor.close()

        if len(info.columns) == 0:
            return None

        return info

    def get_table_by_name(self, tables: List[Table], name: str) -> Optional[Table]:
        for table in tables:
            if table.name == name:
                return table
        return None

