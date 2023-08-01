import json
from unittest import mock

from cloudquery.sdk.plugin import plugin
from cloudquery.sdk.schema import Table

from plugin import TypeformPlugin


@mock.patch("plugin.client.client.TypeformClient")
def test_forms(mock_typeform_client):
    client = mock_typeform_client()
    client.list_forms.return_value = [
        {
            'id': 'form1', 'type': 'quiz', 'title': 'Form 1',
            'last_updated_at': '2023-07-31T18:13:54.780837Z',
            'created_at': '2023-07-31T11:11:53.459949Z',
            'settings': {'is_public': True, 'is_trial': False},
            'self': {'href': 'https://api.typeform.com/forms/XYZ'},
            'theme': {'href': 'https://api.typeform.com/themes/XYZ'},
            '_links': {
                'display': 'https://cloudquery.typeform.com/to/XYZ',
                'responses': 'https://api.typeform.com/forms/XYZ/responses',
            },
        },
        {
            "id": "form2",
            "title": "Form 2",
        },
    ]
    client.list_form_responses.return_value = [
        {
            "response_id": "response_id",
        },
    ]

    p = TypeformPlugin()
    p.init(json.dumps({
        "access_token": "test",
    }))
    msgs = list(p.sync(options=plugin.SyncOptions(
        tables=["typeform_forms"],
    )))
    client.list_forms.assert_called()
    client.list_form_responses.assert_called()

    print(msgs)
    assert len(msgs) == 6  # 2 migrations + 2 forms + 1 response each
    assert Table.from_arrow_schema(msgs[0].table).name == "typeform_forms"
    assert Table.from_arrow_schema(msgs[1].table).name == "typeform_form_responses"

    assert msgs[2].record["id"][0].as_py() == "form1"
    assert msgs[2].record["title"][0].as_py() == "Form 1"
    assert msgs[3].record["id"][0].as_py() == "form2"
    assert msgs[3].record["title"][0].as_py() == "Form 2"
