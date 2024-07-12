// export const listPluginVersionTables = (
//   teamName: string,
//   pluginKind: 'source' | 'destination',
//   pluginName: string,
//   versionName: string,
//   params?: ListPluginVersionTablesParams,
//   signal?: AbortSignal,
// ) => {
//   return customFetch<ListPluginVersionTables200>({
//     method: 'GET',
//     params,
//     signal,
//     url: `/plugins/${teamName}/${pluginKind}/${pluginName}/versions/${versionName}/tables`,
//   });
// };

import {
  UseQueryOptions,
  UseQueryResult,
  QueryFunction,
  useQuery,
  QueryKey,
} from '@tanstack/react-query';
import { useContext } from 'react';
import { AuthContext } from '../contexts/authContext';

interface BasicError {
  message: string;
  status: number;
}

const getPluginTables = (authToken: string, version?: string, signal?: AbortSignal) => {
  const headers = new Headers();
  headers.append('Authorization', `Bearer ${authToken}`);
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  return fetch(
    `https://api.cloudquery.io/plugins/cloudquery/source/aws/versions/${version}/tables`,
    { headers, signal },
  ).then((res) => res.json());
};

export const useGetPluginTables = <
  TData = Awaited<ReturnType<typeof getPluginTables>>,
  TError = BasicError,
>(
  version?: string,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getPluginTables>>, TError, TData>;
  },
): UseQueryResult<TData, TError> => {
  const authToken = useContext(AuthContext);
  const queryKey = [`/plugins/cloudquery/source/aws/versions/${version}/tables`];

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getPluginTables>>> = ({ signal }) =>
    getPluginTables(authToken.value, version, signal);

  const query = useQuery({
    queryFn,
    queryKey,
    ...options,
    enabled: !!version,
  } as UseQueryOptions<Awaited<ReturnType<typeof getPluginTables>>, TError, TData> & {
    queryKey: QueryKey;
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey };

  return query;
};
