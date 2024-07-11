import json
from unittest import mock

from cloudquery.sdk.message import (
    SyncInsertMessage,
    SyncMigrateTableMessage,
)
from cloudquery.sdk.plugin import plugin
from cloudquery.sdk.schema import Table

from plugin import TypeformPlugin


@mock.patch("plugin.client.client.TypeformClient")
def test_forms(mock_typeform_client):
    client = mock_typeform_client()
    client.list_forms.return_value = [
        {
            "id": "form1",
            "type": "quiz",
            "title": "Form 1",
            "last_updated_at": "2023-07-31T18:13:54.780837Z",
            "created_at": "2023-07-31T11:11:53.459949Z",
            "settings": {"is_public": True, "is_trial": False},
            "self": {"href": "https://api.typeform.com/forms/XYZ"},
            "theme": {"href": "https://api.typeform.com/themes/XYZ"},
            "_links": {
                "display": "https://cloudquery.typeform.com/to/XYZ",
                "responses": "https://api.typeform.com/forms/XYZ/responses",
            },
            "tags": ["tag1", "tag2"],
        },
        {
            "id": "form2",
            "title": "Form 2",
        },
    ]
    client.list_form_responses.return_value = [
        {
            "answers": [],
            "submitted_at": "2017-09-14T22:38:22Z",
        },
    ]

    p = TypeformPlugin()
    p.init(
        json.dumps(
            {
                "access_token": "test",
            }
        )
    )
    msgs = list(
        p.sync(
            options=plugin.SyncOptions(
                tables=["typeform_forms"],
                skip_tables=[],
                skip_dependent_tables=False,
            )
        )
    )
    client.list_forms.assert_called()
    client.list_form_responses.assert_called()

    migrations = get_migration_messages(msgs)
    assert len(migrations) == 2
    assert Table.from_arrow_schema(migrations[0].table).name == "typeform_forms"
    assert (
        Table.from_arrow_schema(migrations[1].table).name == "typeform_form_responses"
    )

    form_msgs = get_insert_messages(table="typeform_forms", msgs=msgs)
    assert len(form_msgs) == 2
    assert form_msgs[0].record["id"][0].as_py() == "form1"
    assert form_msgs[0].record["title"][0].as_py() == "Form 1"
    assert form_msgs[1].record["id"][0].as_py() == "form2"
    assert form_msgs[1].record["title"][0].as_py() == "Form 2"

    response_msgs = get_insert_messages(table="typeform_form_responses", msgs=msgs)
    assert len(response_msgs) == 2

    assert len(msgs) == 6


def get_migration_messages(msgs):
    migrations = []
    for msg in msgs:
        if type(msg) == SyncMigrateTableMessage:
            migrations.append(msg)
    return migrations


def get_insert_messages(table, msgs):
    inserts = []
    for msg in msgs:
        if (
            type(msg) == SyncInsertMessage
            and Table.from_arrow_schema(msg.record.schema).name == table
        ):
            inserts.append(msg)
    return inserts
