import {
  UseMutationOptions,
  UseMutationResult,
  MutationFunction,
  useMutation,
} from '@tanstack/react-query';
import { useContext } from 'react';
import { AuthContext } from '../contexts/authContext';

export interface ConnectorAuthRequestAWS {
  /** List of AWS account IDs to authenticate */
  account_ids?: string[];
  /** Kind of the plugin */
  plugin_kind: string;
  /** Name of the plugin */
  plugin_name: string;
  /** Team that owns the plugin we are authenticating the connector for */
  plugin_team: string;
}

interface BasicError {
  message: string;
  status: number;
}

export const createConnector = async (authToken: string, team: string, connectorName: string) => {
  const headers = new Headers();
  headers.append('Authorization', `Bearer ${authToken}`);
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  const name = connectorName || 'AWS Connector';

  return fetch(`https://api.cloudquery.io/teams/${team}/connectors`, {
    headers,
    method: 'POST',
    body: JSON.stringify({
      type: 'aws',
      name,
    }),
  }).then((res) => res.json());
};

const authenticateConnectorAWS = async (authToken: string, team: string, name: string) => {
  const headers = new Headers();
  headers.append('Authorization', `Bearer ${authToken}`);
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  const newConnector = await createConnector(authToken, team, name);

  return fetch(
    `https://api.cloudquery.io/teams/${team}/connectors/${newConnector.id}/authenticate/aws`,
    {
      headers,
      method: 'POST',
      body: JSON.stringify({
        plugin_kind: 'source',
        plugin_name: 'aws',
        plugin_team: 'cloudquery',
      }),
    },
  ).then((res) => ({ connector_id: newConnector.id, ...res.json() }));
};

export const useAuthenticateConnectorAWS = <TError = BasicError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof authenticateConnectorAWS>>,
    TError,
    {},
    TContext
  >;
}): UseMutationResult<
  Awaited<ReturnType<typeof authenticateConnectorAWS>>,
  TError,
  { name: string },
  TContext
> => {
  const { token, team } = useContext(AuthContext);

  const { mutation: providedMutationOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof authenticateConnectorAWS>>,
    { name: string }
  > = ({ name }) => {
    return authenticateConnectorAWS(token, team, name);
  };

  const mutationOptions = { mutationFn, ...providedMutationOptions };

  return useMutation(mutationOptions);
};
