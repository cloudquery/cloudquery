import {
  UseMutationOptions,
  UseMutationResult,
  MutationFunction,
  useMutation,
} from '@tanstack/react-query';
import { useContext } from 'react';
import { AuthContext } from '../contexts/authContext';

interface ConnectorAuthFinishRequestAWS {
  /** External ID in the role definition. Optional. If not provided the previously suggested external ID will be used. Empty string will remove the external ID. */
  external_id?: string;
  /** ARN of role created by the user */
  role_arn: string;
}

interface BasicError {
  message: string;
  status: number;
}

/**
 * Complete authentication for the given AWS connector
 */
export const authenticateConnectorFinishAWS = (
  authToken: string,
  connectorId: string,
  connectorAuthFinishRequestAWS: ConnectorAuthFinishRequestAWS,
) => {
  const headers = new Headers();
  headers.append('Authorization', `Bearer ${authToken}`);
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  return fetch(
    `https://api.cloudquery.io/teams/cloudquery/connectors/${connectorId}/authenticate/aws`,
    {
      headers,
      method: 'PATCH',
      body: JSON.stringify(connectorAuthFinishRequestAWS),
    },
  );
};

export const useAuthenticateConnectorFinishAWS = <
  TError = BasicError,
  TContext = unknown,
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof authenticateConnectorFinishAWS>>,
    TError,
    { connectorId: string; data: ConnectorAuthFinishRequestAWS },
    TContext
  >;
}): UseMutationResult<
  Awaited<ReturnType<typeof authenticateConnectorFinishAWS>>,
  TError,
  { connectorId: string; data: ConnectorAuthFinishRequestAWS },
  TContext
> => {
  const authToken = useContext(AuthContext);

  const { mutation: providedMutationOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof authenticateConnectorFinishAWS>>,
    { connectorId: string; data: ConnectorAuthFinishRequestAWS }
  > = (props) => {
    const { connectorId, data } = props ?? {};

    return authenticateConnectorFinishAWS(authToken.value, connectorId, data);
  };

  const mutationOptions = { mutationFn, ...providedMutationOptions };

  return useMutation(mutationOptions);
};
