import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns

from .invoices import Invoices

locations_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Location"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Locations(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_locations",
            title="Square Locations",
            columns=locations_columns,
            relations=[
                Invoices(),
            ],
        )

    @property
    def resolver(self):
        child_resolvers: list[TableResolver] = []
        for rel in self.relations:
            child_resolvers.append(rel.resolver)

        return LocationsResolver(self, child_resolvers)


class LocationsResolver(TableResolver):
    def __init__(self, table: Table, child_resolvers: list[TableResolver]) -> None:
        super().__init__(table=table, child_resolvers=child_resolvers)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        response = client.client.locations.list()
        for location in response.locations or []:
            yield location.model_dump(mode="json")
