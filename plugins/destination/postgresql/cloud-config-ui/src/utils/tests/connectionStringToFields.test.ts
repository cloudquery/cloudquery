import { convertConnectionStringToFields } from '../convertConnectionStringToFields';

describe('connectionStringToFields (URI)', () => {
  test('empty connection string', () => {
    const result = convertConnectionStringToFields();

    expect(result).toEqual({
      connectionParams: {},
    });
  });

  test('basic empty connection string', () => {
    const result = convertConnectionStringToFields('postgres://');

    expect(result).toMatchObject({});
  });

  test('connection string with host', () => {
    const result = convertConnectionStringToFields('postgres://myhost');

    expect(result).toMatchObject({ host: 'myhost', port: undefined });
  });

  test('connection string with host and port', () => {
    const result = convertConnectionStringToFields('postgres://myhost:1234');

    expect(result).toMatchObject({ host: 'myhost', port: 1234 });
  });

  test('connection string with host, port, and database', () => {
    const result = convertConnectionStringToFields('postgres://myhost:1234/db');

    expect(result).toMatchObject({ host: 'myhost', port: 1234, database: 'db' });
  });

  test('connection string with host, port, database, username, and password', () => {
    const result = convertConnectionStringToFields('postgres://user:pass@myhost:1234/db');

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'db',
      user: 'user',
      password: 'pass',
    });
  });

  test('connection string with host, port, database, and username and environment variable password', () => {
    const result = convertConnectionStringToFields('postgres://user:${password}@myhost:1234/db');

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'db',
      user: 'user',
      password: '${password}',
    });
  });

  test('connection string with host, port, database, username, password, sslmode and search_path', () => {
    const result = convertConnectionStringToFields(
      'postgres://user:pass@myhost:1234/db?sslmode=require&search_path=myschema',
    );

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'db',
      user: 'user',
      password: 'pass',
      connectionParams: {
        ssl: true,
        sslmode: 'require',
        search_path: 'myschema',
      },
    });
  });
});

describe('connectionStringToFields (key-value)', () => {
  test('connection string with host, port, database, username, and password', () => {
    const result = convertConnectionStringToFields(
      "dbtype='postgresql' user='user' password='pass' host='myhost' port='1234' dbname='db' sslmode='require'",
    );

    console.log(result);

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'db',
      user: 'user',
      password: 'pass',
      connectionParams: {
        ssl: true,
        sslmode: 'require',
      },
    });
  });

  test('connection string with empty port', () => {
    const result = convertConnectionStringToFields(
      "dbtype='postgresql' user='user' password='pass' host='myhost' port='' dbname='db' sslmode='require'",
    );

    console.log(result);

    expect(result).toMatchObject({
      host: 'myhost',
      port: undefined,
      database: 'db',
      user: 'user',
      password: 'pass',
      connectionParams: {
        ssl: true,
        sslmode: 'require',
      },
    });
  });
});
