import { convertConnectionStringToFields } from '../convertConnectionStringToFields';

describe('connectionStringToFields', () => {
  test('returns fields from a simple connection string', async () => {
    const result = convertConnectionStringToFields(
      'clickhouse://username:password@host1:9000,host2:9000/database',
    );

    expect(result).toMatchObject({
      username: 'username',
      password: 'password',
      hosts: ['host1:9000', 'host2:9000'],
      database: 'database',
    });
  });

  test('returns fields from a simple connection string with params', async () => {
    const result = convertConnectionStringToFields(
      'clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&block_buffer_size=3',
    );

    expect(result).toMatchObject({
      username: 'username',
      password: 'password',
      hosts: ['host1:9000', 'host2:9000'],
      database: 'database',
      connectionParams: {
        dial_timeout: '200',
        block_buffer_size: '3',
      },
    });
  });

  test('returns fields from a kitchen sink connection string', async () => {
    const result = convertConnectionStringToFields(
      'clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&block_buffer_size=3&debug=true&connection_open_strategy=round_robin&compress=br&compress_level=10&read_timeout=400s',
    );

    expect(result).toMatchObject({
      username: 'username',
      password: 'password',
      hosts: ['host1:9000', 'host2:9000'],
      database: 'database',
      connectionParams: {
        dial_timeout: '200',
        block_buffer_size: '3',
        debug: true,
        connection_open_strategy: 'round_robin',
        compress: 'br',
        compress_level: '10',
        read_timeout: '400',
      },
    });
  });
});
