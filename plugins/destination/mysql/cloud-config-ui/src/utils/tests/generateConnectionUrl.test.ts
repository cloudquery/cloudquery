import { FormValues } from '../formSchema';
import { generateConnectionUrl } from '../generateConnectionUrl';

const baseTestFormValues = {
  username: 'username',
  password: 'password',
  host: 'host',
  port: 'port',
  database: 'database',
  connectionParams: {
    tls: false,
    tlsMode: 'preferred',
    schemaName: '',
    parseTime: false,
    charset: '',
    loc: '',
  },
} as unknown as FormValues;

describe('generateConnectionUrl', () => {
  test('returns a simple connection string', async () => {
    const result = generateConnectionUrl(baseTestFormValues);

    expect(result).toBe('username:${password}@host:port/database');
  });

  test('returns a no-db connection string', async () => {
    const result = generateConnectionUrl({ ...baseTestFormValues, database: undefined });

    expect(result).toBe('username:${password}@host:port/');
  });

  test('returns a no-address connection string', async () => {
    const result = generateConnectionUrl({ ...baseTestFormValues, host: undefined });

    expect(result).toBe('username:${password}@/database');
  });

  test('returns a tls connection string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, tls: true },
    });

    expect(result).toBe('username:${password}@host:port/database?tlsMode=preferred');
  });

  test('returns a tcp connection string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      tcp: true,
      connectionParams: { ...baseTestFormValues.connectionParams },
    });

    expect(result).toBe('username:${password}@tcp(host:port)/database');
  });

  test('returns a parseTime, charset, loc connection string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      connectionParams: {
        ...baseTestFormValues.connectionParams,
        parseTime: true,
        charset: 'utf8',
        loc: 'UTC',
      },
    });

    expect(result).toBe(
      'username:${password}@host:port/database?parseTime=True&charset=utf8&loc=UTC',
    );
  });

  test('returns a kitchen sink string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      tcp: true,
      connectionParams: {
        ...baseTestFormValues.connectionParams,
        tls: true,
        tlsMode: 'preferred',
        parseTime: true,
        charset: 'utf8',
        loc: 'UTC',
      },
    });

    expect(result).toBe(
      'username:${password}@tcp(host:port)/database?tlsMode=preferred&parseTime=True&charset=utf8&loc=UTC',
    );
  });
});
