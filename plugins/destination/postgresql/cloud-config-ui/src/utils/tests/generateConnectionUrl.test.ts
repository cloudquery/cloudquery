import { convertConnectionStringToFields } from '../convertConnectionStringToFields';
import { generateConnectionUrl } from '../generateConnectionUrl';

const baseTestFormValues = {
  user: 'username',
  password: 'password',
  host: 'host',
  port: '1234',
  database: 'db',
  connectionParams: {
    ssl: false,
    sslmode: 'require',
  },
} as any;

describe('generateConnectionUrl', () => {
  test('returns a barebones connection string', async () => {
    const result = generateConnectionUrl({ database: 'db', connectionParams: {} } as any);

    expect(result).toBe('postgres:///db');
  });

  test('returns a simple connection string', async () => {
    const result = generateConnectionUrl(baseTestFormValues);

    expect(result).toBe('postgres://username:${password}@host:1234/db');
  });

  test('returns a simple connection string with no port', async () => {
    const result = generateConnectionUrl({ ...baseTestFormValues, port: '' });

    expect(result).toBe('postgres://username:${password}@host/db');
  });

  test('returns a no-database connection string', async () => {
    const result = generateConnectionUrl({ ...baseTestFormValues, database: '' });

    expect(result).toBe('postgres://username:${password}@host:1234/');
  });

  test('returns a no-address connection string', async () => {
    const result = generateConnectionUrl({ ...baseTestFormValues, host: '' });

    expect(result).toBe('postgres://username:${password}@:1234/db');
  });

  test('returns an ssl connection string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, ssl: true },
    });

    expect(result).toBe('postgres://username:${password}@host:1234/db?sslmode=require');
  });

  test('returns a search_path connection string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, search_path: 'myschema' },
    });

    expect(result).toBe('postgres://username:${password}@host:1234/db?search_path=myschema');
  });

  test('returns a kitchen sink string', async () => {
    const result = generateConnectionUrl({
      ...baseTestFormValues,
      ssl: true,
      connectionParams: {
        ...baseTestFormValues.connectionParams,
        ssl: true,
        sslmode: 'require',
      },
    });

    expect(result).toBe('postgres://username:${password}@host:1234/db?sslmode=require');
  });
});
