import { convertConnectionStringToFields } from '../connectionStringToFields';

describe('connectionStringToFields (URI)', () => {
  test('empty connection string', () => {
    const result = convertConnectionStringToFields();

    expect(result).toEqual({});
  });

  test('basic empty connection string', () => {
    const result = convertConnectionStringToFields('postgres://');

    expect(result).toMatchObject({});
  });

  test('connection string with host', () => {
    const result = convertConnectionStringToFields('postgres://myhost');

    expect(result).toMatchObject({ host: 'myhost' });
  });

  test('connection string with host and port', () => {
    const result = convertConnectionStringToFields('postgres://myhost:1234');

    expect(result).toMatchObject({ host: 'myhost', port: 1234 });
  });

  test('connection string with host, port, and database', () => {
    const result = convertConnectionStringToFields('postgres://myhost:1234/mydb');

    expect(result).toMatchObject({ host: 'myhost', port: 1234, database: 'mydb' });
  });

  test('connection string with host, port, database, username, and password', () => {
    const result = convertConnectionStringToFields('postgres://user:pass@myhost:1234/mydb');

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'mydb',
      username: 'user',
      password: 'pass',
    });
  });

  test('connection string with host, port, database, and username and environment variable password', () => {
    const result = convertConnectionStringToFields('postgres://user:${password}@myhost:1234/mydb');

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'mydb',
      username: 'user',
      password: '${password}',
    });
  });

  test('connection string with host, port, database, username, password, and sslmode', () => {
    const result = convertConnectionStringToFields('postgres://user:pass@myhost:1234/mydb?sslmode=require');

    expect(result).toMatchObject({
      host: 'myhost',
      port: 1234,
      database: 'mydb',
      username: 'user',
      password: 'pass',
      queryParams: {
        sslmode: 'require',
      },
    });
  });
});
