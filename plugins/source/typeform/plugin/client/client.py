from dataclasses import dataclass, field

from plugin.typeform.client import TypeformClient

DEFAULT_CONCURRENCY = 100
DEFAULT_QUEUE_SIZE = 10000


@dataclass
class Spec:
    access_token: str
    base_url: str = field(default="https://api.typeform.com")
    concurrency: int = field(default=DEFAULT_CONCURRENCY)
    queue_size: int = field(default=DEFAULT_QUEUE_SIZE)

    def validate(self):
        if self.access_token is None:
            raise Exception("access_token must be provided")

    @staticmethod
    def json_schema():
        return '''{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://github.com/cloudquery/cloudquery/plugins/source/typeform/spec",
  "$ref": "#/$defs/Spec",
  "$defs": {
    "Spec": {
      "properties": {
        "access_token": {
          "type": "string",
          "minLength": 1,
          "description": "Your [personal access token](https://www.typeform.com/developers/get-started/personal-access-token/) from the Typeform Dashboard."
        },
        "base_url": {
          "type": "string",
          "default": "https://api.typeform.com",
          "description": "The base URL to fetch data from. Use `https://api.eu.typeform.com` if your account is stored in the EU."
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
        "access_token"
      ]
    }
  }
}
'''

class Client:
    def __init__(self, spec: Spec) -> None:
        self._spec = spec
        self._client = TypeformClient(spec.access_token, spec.base_url)

    def id(self):
        return "typeform"

    @property
    def client(self) -> TypeformClient:
        return self._client
