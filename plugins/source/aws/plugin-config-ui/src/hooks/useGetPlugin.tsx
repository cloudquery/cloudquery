import {
  UseQueryOptions,
  UseQueryResult,
  QueryKey,
  useQuery,
  QueryFunction,
} from '@tanstack/react-query';
import { useContext } from 'react';
import { AuthContext } from '../contexts/authContext';

interface BasicError {
  message: string;
  status: number;
}

const getPlugin = (authToken: string, signal?: AbortSignal) => {
  const headers = new Headers();
  headers.append('Authorization', `Bearer ${authToken}`);
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  return fetch(`https://api.cloudquery.io/plugins/cloudquery/source/aws`, { headers, signal }).then(
    (res) => res.json(),
  );
};

export const useGetPlugin = <
  TData = Awaited<ReturnType<typeof getPlugin>>,
  TError = BasicError,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof getPlugin>>, TError, TData>;
}): UseQueryResult<TData, TError> => {
  const authToken = useContext(AuthContext);
  const queryKey = [`/plugins/cloudquery/source/aws`];

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getPlugin>>> = ({ signal }) =>
    getPlugin(authToken.value, signal);

  const query = useQuery({
    queryFn,
    queryKey,
    ...options,
  } as UseQueryOptions<Awaited<ReturnType<typeof getPlugin>>, TError, TData> & {
    queryKey: QueryKey;
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey };

  return query;
};
