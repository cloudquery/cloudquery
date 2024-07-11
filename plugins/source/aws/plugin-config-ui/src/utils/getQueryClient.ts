import { QueryClient } from '@tanstack/react-query';

const retryCodes = new Set([408, 413, 429, 500, 502, 503, 504, 521, 522, 524]);
export function getQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: {
        cacheTime: 0,
        refetchOnWindowFocus: false,
        retry(failureCount, error: any) {
          return retryCodes.has(error?.response?.status) && failureCount < 3;
        },
      },
    },
  });
}
