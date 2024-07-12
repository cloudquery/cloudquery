import {
  UseQueryOptions,
  UseQueryResult,
  QueryFunction,
  useQuery,
  QueryKey,
} from '@tanstack/react-query';
import { useContext } from 'react';
import { AuthContext } from '../contexts/authContext';
import { getPlugin } from './useGetPlugin';

interface BasicError {
  message: string;
  status: number;
}

const getPluginTables = async (authToken: string, signal?: AbortSignal) => {
  const headers = new Headers();
  headers.append('Authorization', `Bearer ${authToken}`);
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  const plugin = await getPlugin(authToken, signal);

  // TODO: BE note about possibly optimizing to use a single fetch call*
  return fetch(
    `https://api.cloudquery.io/plugins/cloudquery/source/aws/versions/${plugin.latest_version}/tables?page=1&per_page=1000`,
    { headers, signal },
  ).then((res) => res.json());
};

export const useGetPluginTables = <
  TData = Awaited<ReturnType<typeof getPluginTables>>,
  TError = BasicError,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof getPluginTables>>, TError, TData>;
}): UseQueryResult<TData, TError> => {
  const authToken = useContext(AuthContext);
  const queryKey = [`/plugins/cloudquery/source/aws/versions/latest/tables`];

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getPluginTables>>> = ({ signal }) =>
    getPluginTables(authToken.value, signal);

  const query = useQuery({
    queryFn,
    queryKey,
    ...options,
  } as UseQueryOptions<Awaited<ReturnType<typeof getPluginTables>>, TError, TData> & {
    queryKey: QueryKey;
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey };

  return query;
};
