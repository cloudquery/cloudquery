import { awsServiceLabelMap } from '../utils/constants';
import { useGetPlugin } from './useGetPlugin';
import { useGetPluginTables } from './useGetPluginTables';

export type AWSService = {
  name: string;
  label: string;
  logo: string;
  tables: string[];
};

export type AWSServices = Record<string, AWSService>;

export const useGetAWSServices = (): AWSServices => {
  try {
    // TODO: BE note about possibly optimizing to use a single fetch call*
    const { data: pluginData } = useGetPlugin();

    const { data: tablesData } = useGetPluginTables(pluginData.latest_version);

    const tablesDataTODO = tablesData?.items ?? [];

    let awsServices = {} as AWSServices;

    for (const table of tablesDataTODO) {
      const serviceName = table.name.split('_')[1];

      if (awsServices[serviceName]) {
        awsServices[serviceName].tables.push(table.name);
      } else {
        awsServices[serviceName] = {
          name: serviceName,
          label: awsServiceLabelMap[serviceName] ?? serviceName,
          logo: `/icons/${serviceName}.svg`,
          tables: [table.name],
        };
      }
    }

    return awsServices;
  } catch (e) {
    console.log(e);
    // TODO: toast? through message broker or install on the plugin?
  }
  return {} as AWSServices;
};
