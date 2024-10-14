import { convertConnectionStringToFields } from '../convertConnectionStringToFields';
import {
  generateConnectionStringKeyValue,
  generateConnectionStringURI,
} from '../generateConnectionString';

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

describe('generateConnectionStringURI', () => {
  test('returns a barebones connection string', async () => {
    const result = generateConnectionStringURI({ database: 'db', connectionParams: {} } as any);

    expect(result).toBe('postgres:///db');
  });

  test('returns a simple connection string', async () => {
    const result = generateConnectionStringURI(baseTestFormValues);

    expect(result).toBe('postgres://username:${password}@host:1234/db');
  });

  test('returns a simple connection string with no port', async () => {
    const result = generateConnectionStringURI({ ...baseTestFormValues, port: '' });

    expect(result).toBe('postgres://username:${password}@host/db');
  });

  test('returns a no-database connection string', async () => {
    const result = generateConnectionStringURI({ ...baseTestFormValues, database: '' });

    expect(result).toBe('postgres://username:${password}@host:1234/');
  });

  test('returns a no-address connection string', async () => {
    const result = generateConnectionStringURI({ ...baseTestFormValues, host: '' });

    expect(result).toBe('postgres://username:${password}@:1234/db');
  });

  test('returns an ssl connection string', async () => {
    const result = generateConnectionStringURI({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, ssl: true },
    });

    expect(result).toBe('postgres://username:${password}@host:1234/db?sslmode=require');
  });

  test('returns a search_path connection string', async () => {
    const result = generateConnectionStringURI({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, search_path: 'myschema' },
    });

    expect(result).toBe('postgres://username:${password}@host:1234/db?search_path=myschema');
  });

  test('returns a kitchen sink string', async () => {
    const result = generateConnectionStringURI({
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

describe('generateConnectionStringKeyValue', () => {
  test('returns a barebones connection string', async () => {
    const result = generateConnectionStringKeyValue(baseTestFormValues);

    expect(result).toBe(
      "user='username' password='${password}' host='host' dbname='db' port='1234'",
    );
  });

  test('return a connection string with slashes and backslashes escaped', async () => {
    const result = generateConnectionStringKeyValue({
      ...baseTestFormValues,
      host: "host'with'single'quotes",
      database: "db'with'single'quotes",
    });

    expect(result).toBe(
      "user='username' password='${password}' host='host\\'with\\'single\\'quotes' dbname='db\\'with\\'single\\'quotes' port='1234'",
    );
  });

  test('returns a connection string with no port', async () => {
    const result = generateConnectionStringKeyValue({ ...baseTestFormValues, port: '' });

    expect(result).toBe("user='username' password='${password}' host='host' dbname='db'");
  });

  test('returns a connection string with ssl mode', async () => {
    const result = generateConnectionStringKeyValue({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, ssl: true, sslmode: 'require' },
    });

    expect(result).toBe(
      "user='username' password='${password}' host='host' dbname='db' port='1234' sslmode='require'",
    );
  });

  test('returns a connection string with search path', async () => {
    const result = generateConnectionStringKeyValue({
      ...baseTestFormValues,
      connectionParams: { ...baseTestFormValues.connectionParams, search_path: 'myschema' },
    });

    expect(result).toBe(
      "user='username' password='${password}' host='host' dbname='db' port='1234' search_path='myschema'",
    );
  });
});
