export const convertConnectionStringToFields = (connectionString?: string) => {
  try {
    const connectionParams: Record<string, any> = {};

    if (!connectionString) {
      return {
        connectionParams,
      };
    }

    // Return parsed components
    return parseConnectionFieldsFromURI(connectionString);
  } catch {
    return {
      connectionParams: {},
    };
  }
};

/* eslint-disable unicorn/no-abusive-eslint-disable */
/* eslint-disable */

/**
 * Parses connection fields from URI i.e. "clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms"
 *
 * @param connectionString
 *
 * @returns
 */
function parseConnectionFieldsFromURI(connectionString: string) {
  const connectionParams: Record<string, any> = {};

  const withoutProtocol = connectionString.replace('clickhouse://', '');

  // Split the connection string into the main part and optional query parameters
  const [mainPart, queryString] = withoutProtocol.split('?');

  // Split the main part into [user info, hosts] and database
  const splitDBAt = mainPart.lastIndexOf('/');

  const userInfoHosts = mainPart.slice(0, splitDBAt);
  const database = mainPart.slice(splitDBAt + 1);

  console.log({ userInfoHosts, database });

  // Initialize components
  let username = null,
    password = null,
    hosts = null;

  const [userInfo, hostsString] = userInfoHosts.split('@');

  // Check if password is included
  if (userInfo.includes(':')) {
    [username, password] = userInfo.split(':');
  } else {
    username = userInfo;
  }

  hosts = hostsString.split(',');

  // Parse query parameters if present
  if (queryString) {
    const pairs = queryString.split('&');
    for (const pair of pairs) {
      const [key, value] = pair.split('=');
      connectionParams[key] = decodeURIComponent(value || '');

      switch (key) {
        // handle boolean values
        case 'debug': {
          connectionParams[key] = value === 'true';

          break;
        }

        // handle postfixes
        case 'dial_timeout': {
          connectionParams[key] = value.replace('ms', '');

          break;
        }
        case 'read_timeout': {
          connectionParams[key] = value.replace('s', '');

          break;
        }
        // No default
      }
    }
  }

  // Return parsed components
  return {
    protocol: 'clickhouse',
    username,
    password,
    hosts,
    database,
    connectionParams,
  };
}
