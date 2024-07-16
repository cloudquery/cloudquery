import { serviceLabelMap, serviceNameResolutions } from '../utils/constants';
import tableData from './TODO-tabledata.json';
import { ServiceTypes } from '../components/serviceList';

export const useGetServices = (): ServiceTypes => {
  try {
    // const { data } = useGetPluginTables();

    const tablesData = tableData;

    let services = {} as ServiceTypes;

    for (const table of tablesData) {
      let serviceName = table.name.split('_')[1];

      if (serviceNameResolutions[serviceName]) {
        serviceName = serviceNameResolutions[serviceName];
      }

      if (services[serviceName]) {
        services[serviceName].tables.push(table.name);
      } else {
        services[serviceName] = {
          name: serviceName,
          label: serviceLabelMap[serviceName] ?? serviceName,
          logo: `/icons/${serviceName}.svg`,
          tables: [table.name],
        };
      }
    }

    return services;
  } catch (e) {
    console.log(e);
    // TODO:ERROR, toast?
  }
  return {} as ServiceTypes;
};
