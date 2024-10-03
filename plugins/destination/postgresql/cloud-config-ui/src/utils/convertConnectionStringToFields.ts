export const convertConnectionStringToFields = (connectionString?: string) => {
  try {
    const connectionParams: Record<string, any> = {};

    if (!connectionString) {
      return {
        connectionParams,
      };
    }

    if (
      connectionString.startsWith('postgres://') ||
      connectionString.startsWith('postgresql://')
    ) {
      return parseConnectionFieldsFromURI(connectionString);
    }

    // Return parsed components
    return parseConnectionFieldsFromKeyValue(connectionString);
  } catch {
    return {
      connectionParams: {},
    };
  }
};

/**
 * Parses connection fields from URI i.e. "postgres://user:pass@myhost:1234/db?sslmode=require"
 *
 * @param connectionString
 *
 * @returns
 */
function parseConnectionFieldsFromURI(connectionString: string) {
  const connectionParams: Record<string, any> = {};

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
    port = null;

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
    port: port ? Number.parseInt(port, 10) : undefined,
    database,
    connectionParams,
  };
}

/**
 * Parses connection fields from key value pairs i.e. "dbtype='postgresql' user='user' password='pass' host='myhost' port='1234' dbname='db' sslmode='require'"
 *
 *
 * @param connectionString - connection string
 *
 * @returns {Record<string, any>} connection fields
 */
function parseConnectionFieldsFromKeyValue(connectionString: string): Record<string, any> {
  const connectionFields: Record<string, any> = {};
  const connectionParams: Record<string, any> = {};

  // Split the connection string into key-value pairs
  const pairs = connectionString.split(' ');

  // Parse key-value pairs
  for (const pair of pairs) {
    const [key, value] = pair.split('=');

    const cleanedValue = value.replace(/'/g, '');

    switch (key) {
      case 'sslmode': {
        connectionParams.ssl = true;
        connectionParams.sslmode = cleanedValue;

        break;
      }
      case 'search_path': {
        connectionParams.search_path = cleanedValue;

        break;
      }
      case 'dbname': {
        connectionFields.database = cleanedValue;

        break;
      }
      case 'port': {
        connectionFields.port = cleanedValue ? Number.parseInt(cleanedValue, 10) : undefined;

        break;
      }
      default: {
        connectionFields[key] = cleanedValue;
      }
    }
  }

  return {
    ...connectionFields,
    connectionParams,
  };
}
