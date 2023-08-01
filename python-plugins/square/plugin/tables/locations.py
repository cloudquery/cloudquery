import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from cloudquery.sdk.types import JSONType
from plugin.client import Client
from square.api.locations_api import LocationsApi


class Locations(Table):
    def __init__(self) -> None:
        super().__init__(
            "locations",
            [
                Column("id", pa.string(), primary_key=True),
                Column("name", pa.string()),
                Column("address", JSONType()),
                Column("timezone", pa.string()),
                Column("capabilities", JSONType()),
                Column("status", pa.string()),
                Column("created_at", pa.timestamp("ms")),
                Column("merchant_id", pa.string()),
                Column("country", pa.string()),
                Column("language_code", pa.string()),
                Column("currency", pa.string()),
            ],
        )

    @property
    def resolver(self):
        return LocationsResolver(self)


class LocationsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        locations: LocationsApi = client.client.locations
        response = locations.list_locations()
        if response.is_error():
            raise Exception(response.errors)

        for location in response.locations:
            yield location
