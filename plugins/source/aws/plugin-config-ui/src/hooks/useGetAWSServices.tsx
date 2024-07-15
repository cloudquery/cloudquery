import { awsServiceLabelMap } from '../utils/constants';
import { useGetPluginTables } from './useGetPluginTables';

export type AWSService = {
  name: string;
  label: string;
  logo: string;
  tables: string[];
};

export type AWSServices = Record<string, AWSService>;

const serviceNameResolutions: Record<string, string> = {
  acmpca: 'acm',
  apigatewayv2: 'apigateway',
  cloudwatchlogs: 'cloudwatch',
  elbv1: 'elb',
  elbv2: 'elb',
  route53recoverycontrolconfig: 'route53',
  route53resolver: 'route53',
  ssmincidents: 'ssm',
  wafregional: 'waf',
  wafv2: 'waf',
};

export const useGetAWSServices = (): AWSServices => {
  try {
    const { data: tablesData } = useGetPluginTables();

    const tablesDataTODO = tablesData?.items ?? [];

    let awsServices = {} as AWSServices;

    for (const table of tablesDataTODO) {
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
    // TODO: toast? through message broker or install on the plugin?
  }
  return {} as AWSServices;
};
