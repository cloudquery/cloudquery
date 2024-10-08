import { generateConnectionStringURI } from '../generateConnectionString';

const _secretKeys = ['username', 'password', 'database', 'connection_string'];

const baseTestFormValues = {
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
  _secretKeys,
} as any;

describe('generateConnectionStringURI', () => {
  test('returns a simple connection string with single host', async () => {
    const result = generateConnectionStringURI({
      username: 'username',
      password: 'password',
      hosts: ['host1:9000'],
      database: 'database',
      connectionParams: {},
      _secretKeys,
    } as any);

    expect(result.connection_string).toBe(
      'clickhouse://${username}:${password}@host1:9000/${database}',
    );
  });

  test('returns a simple connection string with multiple hosts', async () => {
    const result = generateConnectionStringURI({
      username: 'username',
      password: 'password',
      hosts: ['host1:9000', 'host2:9000'],
      database: 'database',
      connectionParams: {},
      _secretKeys,
    } as any);

    expect(result.connection_string).toBe(
      'clickhouse://${username}:${password}@host1:9000,host2:9000/${database}',
    );
  });

  test('returns a kitchen sink string', async () => {
    const result = generateConnectionStringURI(baseTestFormValues);

    expect(result.connection_string).toBe(
      'clickhouse://${username}:${password}@host1:9000,host2:9000/${database}?dial_timeout=200ms&block_buffer_size=3&debug=true&connection_open_strategy=round_robin&compress=br&compress_level=10&read_timeout=400s',
    );
  });
});
