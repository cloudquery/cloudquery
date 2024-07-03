from typing import Any, Generator

import pyarrow as pa
from cloudquery.sdk.scheduler import TableResolver
from cloudquery.sdk.schema import Column
from cloudquery.sdk.schema import Table
from cloudquery.sdk.schema.resource import Resource
from cloudquery.sdk.types import JSONType
from cloudquery.sdk.stateclient.stateclient import StateClient

from plugin.client import Client


class FormResponses(Table):
    def __init__(self) -> None:
        super().__init__(
            name="typeform_form_responses",
            title="Typeform Form Responses",
            columns=[
                Column("form_id", pa.string(), primary_key=True),
                Column("response_id", pa.string(), primary_key=True),
                Column("landing_id", pa.string()),
                Column("landed_at", pa.timestamp(unit="s")),
                Column("submitted_at", pa.timestamp(unit="s")),
                Column("token", pa.string()),
                Column("metadata", JSONType()),
                Column("answers", JSONType()),
                Column("hidden", JSONType()),
                Column("calculated", JSONType()),
                Column("variables", JSONType()),
                Column("tags", JSONType()),
            ],
        )

    @property
    def resolver(self):
        return FormResponsesResolver(table=self)


class FormResponsesResolver(TableResolver):
    def __init__(self, table) -> None:
        super().__init__(table=table)

    def resolve(
        self, client: Client, parent_resource: Resource
    ) -> Generator[Any, None, None]:
        print("In FormResponsesResolver.resolve, I found this stateclient: ", self.state_client)
        for form_response in client.client.list_form_responses(
            form_id=parent_resource.item["id"]
        ):
            form_response["form_id"] = parent_resource.item["id"]
            yield form_response
