import pyarrow as pa
import pytest
from cloudquery.sdk.message import WriteMigrateTableMessage
from cloudquery.sdk.schema import Column, Table

from plugin.sqlite.client import SQLClient, SQLiteColumn, _identifier

HOSTILE_TABLE = "quoted\" (x); CREATE TABLE injected (y); --'"
HOSTILE_COLUMN = "id\"; --'"
SOURCE_COLUMN = "_cq_source_name"
SYNC_TIME_COLUMN = "_cq_sync_time"


@pytest.fixture
def client():
    sql_client = SQLClient(":memory:")
    yield sql_client
    sql_client.close()


def table_names(client: SQLClient) -> list[str]:
    rows = client.conn.execute("SELECT name FROM sqlite_master WHERE type = 'table'")
    return [row[0] for row in rows]


def hostile_table(extra_columns: list[Column] = ()) -> Table:
    return Table(
        name=HOSTILE_TABLE,
        columns=[
            Column(name=HOSTILE_COLUMN, type=pa.string()),
            Column(name=SOURCE_COLUMN, type=pa.string()),
            Column(name=SYNC_TIME_COLUMN, type=pa.string()),
            *extra_columns,
        ],
    )


def migrate(client: SQLClient, table: Table):
    error = client.migrate_client.migrate_tables(
        [WriteMigrateTableMessage(table=table, migrate_force=False)]
    )
    assert error is None


@pytest.mark.parametrize(
    "name, expected",
    [
        ("plain", '"plain"'),
        ('with"quote', '"with""quote"'),
        ('a"; drop; --', '"a""; drop; --"'),
        ('""', '""""""'),
        ("", '""'),
    ],
)
def test_identifier_doubles_embedded_quotes(client, name, expected):
    assert _identifier(name) == expected
    assert client.migrate_client.identifier(name) == expected


def test_migrate_write_and_delete_stale_with_hostile_identifiers(client):
    columns = [HOSTILE_COLUMN, SOURCE_COLUMN, SYNC_TIME_COLUMN]
    migrate(client, hostile_table())
    assert table_names(client) == [HOSTILE_TABLE]

    client.insert(
        HOSTILE_TABLE,
        columns,
        [
            ("old", "source-test", "2020-01-01T00:00:00"),
            ("new", "source-test", "2030-01-01T00:00:00"),
        ],
    )

    migrate(client, hostile_table([Column(name='added"; --', type=pa.string())]))
    info = client.migrate_client.get_table_info(HOSTILE_TABLE)
    assert [col.name for col in info.columns] == [*columns, 'added"; --']
    assert len(list(client.read(table_name=HOSTILE_TABLE, col_names=columns))) == 2

    client.delete_stale(
        table_name=HOSTILE_TABLE,
        source_name="source-test",
        sync_time="2025-01-01T00:00:00",
        cq_sync_time_column=SYNC_TIME_COLUMN,
        cq_source_name_column=SOURCE_COLUMN,
    )

    rows = list(client.read(table_name=HOSTILE_TABLE, col_names=columns))
    assert rows == [("new", "source-test", "2030-01-01T00:00:00")]
    assert table_names(client) == [HOSTILE_TABLE]


def test_create_table_with_hostile_identifiers(client):
    client.create_table(
        HOSTILE_TABLE, [SQLiteColumn(name=HOSTILE_COLUMN, type="text")], False
    )

    assert table_names(client) == [HOSTILE_TABLE]
    info = client.migrate_client.get_table_info(HOSTILE_TABLE)
    assert [col.name for col in info.columns] == [HOSTILE_COLUMN]


def test_get_table_info_returns_none_for_missing_table(client):
    assert client.migrate_client.get_table_info(HOSTILE_TABLE) is None
