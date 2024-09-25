export const convertConnectionStringToFields = (connectionString?: string) => {
  const connectionParams: Record<string, any> = {};

  if (!connectionString) {
    return {
      connectionParams,
    };
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
  let user = null,
    password = null,
    host = null,
    port = '5432'; // default port is 5432

  // Check if there is user info (user and password)
  if (userInfoHost.includes('@')) {
    const [userInfo, hostPort] = userInfoHost.split('@');

    // Check if password is included
    if (userInfo.includes(':')) {
      [user, password] = userInfo.split(':');
    } else {
      user = userInfo;
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
  if (queryString) {
    const pairs = queryString.split('&');
    for (const pair of pairs) {
      const [key, value] = pair.split('=');
      if (key === 'sslmode') {
        connectionParams.ssl = true;
        connectionParams.sslmode = decodeURIComponent(value || '');
      } else {
        connectionParams[key] = decodeURIComponent(value || '');
      }
    }
  }

  // Return parsed components
  return {
    protocol: connectionString.startsWith('postgresql://') ? 'postgresql' : 'postgres',
    user,
    password,
    host,
    port: Number.parseInt(port, 10),
    database,
    connectionParams,
  };
};
