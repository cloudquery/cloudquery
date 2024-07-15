import { awsServiceLabelMap, serviceNameResolutions } from '../utils/constants';
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
    const { data } = useGetPluginTables();

    const tablesData = data?.items ?? [];

    let awsServices = {} as AWSServices;

    for (const table of tablesData) {
      let serviceName = table.name.split('_')[1];

      if (serviceNameResolutions[serviceName]) {
        serviceName = serviceNameResolutions[serviceName];
      }

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
    // TODO:ERROR, toast?
  }
  return {} as AWSServices;
};
