import { convertConnectionStringToFields } from '../convertConnectionStringToFields';

describe('connectionStringToFields', () => {
  test('returns fields from a simple connection string', async () => {
    const result = convertConnectionStringToFields(
      'user:password@localhost:3306/dbname?timeout=30s\u0026readTimeout=1s\u0026writeTimeout=1s',
    );

    expect(result).toMatchObject({
      username: 'user',
      password: 'password',
      host: 'localhost',
      port: '3306',
      database: 'dbname',
      tcp: false,
      connectionParams: {
        timeout: 30,
        readTimeout: 1,
        writeTimeout: 1,
      },
    });
  });

  test('returns fields from a kitchen sink connection string', async () => {
    const result = convertConnectionStringToFields(
      'user@gmail.com:${password}@tcp(host:port)/database?tlsMode=preferred\u0026parseTime=True\u0026charset=utf8\u0026loc=UTC\u0026timeout=30s\u0026readTimeout=60s\u0026writeTimeout=90s\u0026allowNativePasswords=true',
    );

    expect(result).toMatchObject({
      username: 'user@gmail.com',
      password: '${password}',
      host: 'host',
      port: 'port',
      database: 'database',
      tcp: true,
      connectionParams: {
        tls: true,
        tlsMode: 'preferred',
        parseTime: true,
        charset: 'utf8',
        loc: 'UTC',
        timeout: 30,
        readTimeout: 60,
        writeTimeout: 90,
        allowNativePasswords: true,
      },
    });
  });
});
