import { default as Ajv } from 'ajv';
import camelcaseKeys from 'camelcase-keys';

const spec = {
  type: 'object',
  properties: {
    concurrency: { type: 'integer' },
    // eslint-disable-next-line @typescript-eslint/naming-convention
    access_token: { type: 'string' },
    // eslint-disable-next-line @typescript-eslint/naming-convention
    endpoint_url: { type: 'string' },
  },
  required: ['access_token'],
};

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
