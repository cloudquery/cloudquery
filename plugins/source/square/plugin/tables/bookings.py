import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns
from square.core.api_error import ApiError

bookings_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Booking"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Bookings(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_bookings",
            title="Square Bookings",
            columns=bookings_columns,
        )

    @property
    def resolver(self):
        return BookingsResolver(self)


class BookingsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        try:
            for booking in client.client.bookings.list():
                yield booking.model_dump(mode="json")
        except ApiError as e:
            for error in e.errors:
                if (
                    error.category == "AUTHENTICATION_ERROR"
                    and error.code == "UNAUTHORIZED"
                    and error.detail == "Merchant not onboarded to Appointments"
                ):
                    # TODO log
                    return
            raise
