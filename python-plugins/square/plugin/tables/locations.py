import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from cloudquery.sdk.types import JSONType
from plugin.client import Client
from square.api.locations_api import LocationsApi
from square.http.api_response import ApiResponse
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns

columns = oapi_definition_to_columns(
    OAPILoader.get_definition('Location'),
    override_columns=[Column(name='id', type=pa.string(), primary_key=True)])


class Locations(Table):
    def __init__(self) -> None:
        super().__init__(
            "square_locations",
            columns=columns,
        )

    @property
    def resolver(self):
        return LocationsResolver(self)


class LocationsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        locations: LocationsApi = client.client.locations
        response: ApiResponse = locations.list_locations()
        if response.is_error():
            raise Exception(response)
        for location in response.body.get('locations', []):
            yield location
