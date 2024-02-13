from dataclasses import dataclass, field
import square
import square.client

DEFAULT_CONCURRENCY = 100
DEFAULT_QUEUE_SIZE = 10000


@dataclass
class Spec:
    access_token: str
    environment: str
    concurrency: int = field(default=DEFAULT_CONCURRENCY)
    queue_size: int = field(default=DEFAULT_QUEUE_SIZE)

    def validate(self):
        if self.environment not in ["sandbox", "production"]:
            raise Exception("environment must be one of 'sandbox' or 'production'")
        if self.access_token is None:
            raise Exception("access_token must be provided")

    @staticmethod
    def json_schema():
        return """{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://github.com/cloudquery/cloudquery/plugins/source/square/spec",
  "$ref": "#/$defs/Spec",
  "$defs": {
    "Spec": {
      "properties": {
        "access_token": {
          "type": "string",
          "minLength": 1,
          "description": "Your access token from Square."
        },
        "environment": {
          "type": "string",
          "enum": [
            "sandbox",
            "production"
          ],
          "description": "The environment to use."
        },
        "concurrency": {
          "type": "integer",
          "minimum": 1,
          "default": 100,
          "description": "Maximum number of requests to perform concurrently."
        },
        "queue_size": {
          "type": "integer",
          "minimum": 1,
          "default": 10000,
          "description": "Maximum number of items to have in the queue before waiting for an unfinished request to complete."
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "access_token",
        "environment"
      ]
    }
  }
}
"""


class Client:
    def __init__(self, spec: Spec) -> None:
        self._spec = spec
        self._client = square.client.Client(
            access_token=self._spec.access_token, environment=self._spec.environment
        )

    def id(self):
        return "square"

    @property
    def client(self) -> square.client.Client:
        return self._client
