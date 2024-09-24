type ConnectionFields = {
  protocol: 'postgres' | 'postgresql';
  username: string | null;
  password: string | null;
  host: string | null;
  port: number;
  database: string | null;
  queryParams: Record<string, string>;
};

export const convertConnectionStringToFields = (connectionString?: string): ConnectionFields => {
  if (!connectionString) {
    return {} as ConnectionFields;
  }

  // Check if the string starts with 'postgres://' or 'postgresql://'
  if (
    !connectionString.startsWith('postgres://') &&
    !connectionString.startsWith('postgresql://')
  ) {
    throw new Error('Invalid connection string: must start with "postgres://" or "postgresql://"');
  }

  // Remove the protocol part ('postgres://' or 'postgresql://')
  const protocolEndIndex = connectionString.indexOf('://') + 3;
  const withoutProtocol = connectionString.slice(protocolEndIndex);

  // Split the connection string into the main part and optional query parameters
  const [mainPart, queryString] = withoutProtocol.split('?');

  // Split the main part into user info, host, and database
  const [userInfoHost, database] = mainPart.split('/');

  // Initialize components
  let username = null,
    password = null,
    host = null,
    port = '5432'; // default port is 5432

  // Check if there is user info (username and password)
  if (userInfoHost.includes('@')) {
    const [userInfo, hostPort] = userInfoHost.split('@');

    // Check if password is included
    if (userInfo.includes(':')) {
      [username, password] = userInfo.split(':');
    } else {
      username = userInfo;
    }

    // Check if port is included
    if (hostPort.includes(':')) {
      [host, port] = hostPort.split(':');
    } else {
      host = hostPort;
    }
  } else {
    // No user info, just host and port
    if (userInfoHost.includes(':')) {
      [host, port] = userInfoHost.split(':');
    } else {
      host = userInfoHost;
    }
  }

  // Parse query parameters if present
  const queryParams = {};
  if (queryString) {
    const pairs = queryString.split('&');
    for (const pair of pairs) {
      const [key, value] = pair.split('=');
      queryParams[key] = decodeURIComponent(value || '');
    }
  }

  // Return parsed components
  return {
    protocol: connectionString.startsWith('postgresql://') ? 'postgresql' : 'postgres',
    username,
    password,
    host,
    port: Number.parseInt(port, 10),
    database,
    queryParams,
  };
};
