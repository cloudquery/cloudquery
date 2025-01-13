import { default as Ajv } from 'ajv';
import camelcaseKeys from 'camelcase-keys';

const spec = {
  type: 'object',
  properties: {
    concurrency: { type: 'integer', minimum: 1 },
    // eslint-disable-next-line @typescript-eslint/naming-convention
    access_token: { type: 'string', minLength: 1 },
    // eslint-disable-next-line @typescript-eslint/naming-convention
    endpoint_url: { type: 'string' },
  },
  required: ['access_token'],
};

export const JSON_SCHEMA = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://github.com/cloudquery/cloudquery/plugins/source/airtable/spec",
  "$ref": "#/$defs/Spec",
  "$defs": {
    "Spec": {
      "properties": {
        "access_token": {
          "type": "string",
          "minLength": 1,
          "description": "Your Airtable API [personal access token](https://airtable.com/developers/web/guides/personal-access-tokens)."
        },
        "endpoint_url": {
          "type": "string",
          "default": "https://api.airtable.com",
          "description": "The endpoint URL to fetch data from."
        },
        "concurrency": {
          "type": "integer",
          "minimum": 1,
          "default": 10000,
          "description": "Best effort maximum number of tables to sync concurrently."
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "access_token"
      ]
    }
  }
}`;

type JSONSpec = {
  concurrency: number;
  // eslint-disable-next-line @typescript-eslint/naming-convention
  access_token: string;
  // eslint-disable-next-line @typescript-eslint/naming-convention
  endpoint_url: string;
};

const ajv = new Ajv.default();
const validate = ajv.compile(spec);

export type Spec = {
  concurrency: number;
  apiKey: string;
  endpointUrl: string;
};

export const parseSpec = (spec: string): Spec => {
  const parsed = JSON.parse(spec) as Partial<JSONSpec>;
  const valid = validate(parsed);
  if (!valid) {
    throw new Error(`Invalid spec: ${JSON.stringify(validate.errors)}`);
  }
  const { concurrency = 10_000, accessToken = '', endpointUrl = 'https://api.airtable.com' } = camelcaseKeys(parsed);
  return { concurrency, apiKey: accessToken, endpointUrl };
};
