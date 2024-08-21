import sqlite3
from typing import Generator
from plugin.sqlite.migrate import MigrateSQLClient


class SQLiteColumn:
    def __init__(
        self,
        *,
        name: str,
        type: str,
        description: str = "",
        primary_key: bool = False,
        not_null: bool = False,
        incremental_key: bool = False,
        unique: bool = False,
    ) -> None:
        self.name = name
        self.type = type
        self.description = description
        self.primary_key = primary_key
        self.not_null = not_null
        self.incremental_key = incremental_key
        self.unique = unique

    def to_create_sql(self):
        sql = f"{self.name} {self.type}"
        if self.primary_key:
            sql += " PRIMARY KEY"
        if self.not_null:
            sql += " NOT NULL"
        return sql


def _identifier(name):
    return f'"{name}"'


class SQLClient:
    def __init__(self, connection_string: str):
        self.connection_string = connection_string
        self.conn = sqlite3.connect(
            connection_string, check_same_thread=False, isolation_level=None
        )
        self.migrate_client = MigrateSQLClient(self.conn)

    def close(self):
        self.conn.close()

    def create_table(
        self, table_name: str, cols: list[SQLiteColumn], migrate_force: bool
    ):
        self.conn.execute(
            f"CREATE TABLE IF NOT EXISTS {table_name} ({', '.join(col.to_create_sql() for col in cols)})"
        )

    def insert(
        self,
        table_name: str,
        col_names: list[str],
        values: list[tuple],
        primary_keys: list[str] = None,
    ):
        placeholders = ", ".join(f"?{i+1}" for i in range(len(col_names)))
        columns_list = ", ".join(_identifier(col) for col in col_names)

        if primary_keys:
            sql_string = f"INSERT OR REPLACE INTO {_identifier(table_name)} ({columns_list}) VALUES ({placeholders})"
        else:
            sql_string = f"INSERT INTO {_identifier(table_name)} ({columns_list}) VALUES ({placeholders})"

        for v in values:
            try:
                self.conn.execute(sql_string, v)
            except sqlite3.Error as e:
                raise RuntimeError(f"Failed to execute '{sql_string}': {e}")

    def read(
        self, *, table_name: str, col_names: list[str]
    ) -> Generator[tuple, None, None]:
        cols = ", ".join(col_names)
        cursor = self.conn.cursor()
        rows = []
        try:
            cursor.execute("SELECT {} FROM {}".format(cols, table_name))
            while True:
                row = cursor.fetchone()
                if row is None:
                    break
                rows.append(row)
        finally:
            cursor.close()
        for row in rows:
            yield row

    def delete_stale(
        self,
        *,
        table_name: str,
        source_name: str,
        sync_time: str,
        cq_sync_time_column: str,
        cq_source_name_column: str,
    ):
        cursor = self.conn.cursor()
        sql = f"""
        DELETE FROM "{table_name}"
        WHERE "{cq_source_name_column}" = ?
        AND datetime("{cq_sync_time_column}") < datetime(?)
        """
        cursor.execute(sql, (source_name, sync_time))
        cursor.close()
