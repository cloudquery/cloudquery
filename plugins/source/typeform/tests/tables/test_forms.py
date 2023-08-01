import json
from unittest import mock

from cloudquery.sdk.plugin import plugin

from plugin import TypeformPlugin


@mock.patch("plugin.client.client.TypeformClient")
def test_forms(mock_typeform_client):
    mock_typeform_client().list_forms.return_value = [
        {
            "id": "form1",
            "title": "Form 1",
        },
        {
            "id": "form2",
            "title": "Form 2",
        },
    ]

    p = TypeformPlugin()
    p.init(json.dumps({
        "access_token": "test",
    }))
    forms = list(p.sync(options=plugin.SyncOptions(
        tables=["typeform_forms"],
    )))
    mock_typeform_client().list_forms.assert_called()
    print(forms)
    assert len(forms) == 2
