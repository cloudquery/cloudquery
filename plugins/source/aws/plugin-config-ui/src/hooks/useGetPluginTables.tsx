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

export const useGetPluginTables = () => {
  // TODO: note about possibly optimizing to use a single fetch call*
};
