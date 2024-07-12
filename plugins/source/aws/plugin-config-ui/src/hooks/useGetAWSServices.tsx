import { awsServiceLabelMap } from '../utils/constants';
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
    const { data: tablesData } = useGetPluginTables();

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
