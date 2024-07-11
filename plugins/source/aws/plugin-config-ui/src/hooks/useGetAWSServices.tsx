import { useGetPlugin } from './useGetPlugin';
import { useGetPluginTables } from './useGetPluginTables';

export type AWSService = {
  name: string;
  logo: string;
  tables: string[];
};

export type AWSServices = Record<string, AWSService>;

// TODO: finish this!!
export const useGetAWSServices = (): AWSServices => {
  try {
    const {
      data: pluginData,
      // error: pluginDataError,
      // isLoading: pluginDataLoading,
    } = useGetPlugin();
    console.log({ pluginData });
    // TODO: get version
    const {
      data: tablesData,
      // error: pluginDataError,
      // isLoading: pluginDataLoading,
    } = useGetPluginTables('v27.7.0'); // TODO

    //TODO:
    console.log({ tablesData });

    const tablesDataTODO = [
      {
        description:
          'https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html',
        is_incremental: false,
        is_paid: true,
        name: 'aws_backup_region_settings',
        relations: [],
        title: 'Backup Region Settings',
      },
      {
        description: 'https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ReportPlan.html',
        is_incremental: false,
        is_paid: true,
        name: 'aws_backup_report_plans',
        relations: [],
        title: 'Backup Report Plans',
      },
      {
        description:
          'https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html',
        is_incremental: false,
        is_paid: true,
        name: 'aws_cloudwatchlogs_log_groups',
        relations: [
          'aws_cloudwatchlogs_log_group_subscription_filters',
          'aws_cloudwatchlogs_log_group_data_protection_policies',
        ],
        title: 'Cloudwatchlogs Log Groups',
      },
    ];

    let awsServices = {} as AWSServices;

    for (const table of tablesDataTODO) {
      const serviceName = table.name.split('_')[1];

      if (awsServices[serviceName]) {
        awsServices[serviceName].tables.push(table.name);
      } else {
        awsServices[serviceName] = {
          name: serviceName,
          logo: `/icons/${serviceName}.svg`,
          tables: [table.name],
        };
      }
    }

    return awsServices;
  } catch (e) {
    console.log(e);
  }
  return {} as AWSServices;
  // TODO: note about possibly optimizing to use a single fetch call*
  // 1 fetch plugin
  // 2 fetch tables by latest_version
};
