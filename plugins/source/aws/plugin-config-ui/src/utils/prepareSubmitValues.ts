import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';
import { AWSServices } from '../hooks/useGetAWSServices';

export function prepareSubmitValues(
  values: FormValues,
  awsServices: AWSServices,
): PluginUiMessagePayload['validation_passed']['values'] {
  return {
    migrateMode: values.migrateMode, // TODO:SUBMIT, needed?
    writeMode: values.writeMode, // TODO:SUBMIT, needed?
    envs: [], //TODO:SUBMIT, needed?
    // connector_id: values.connector_id, // TODO:SUBMIT correct location?
    tables: getTablesFromServices(values.services, awsServices),
    spec: {
      regions: values.spec.regions,
    },
  };
}

function getTablesFromServices(services: string[], awsServices: AWSServices): string[] {
  return services.map((service: string) => awsServices[service].tables).flat();
}
