import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';
import { ServiceTypes } from '../components/todoShare/serviceList';

export function prepareSubmitValues(
  values: FormValues,
  allServices: ServiceTypes,
): PluginUiMessagePayload['validation_passed']['values'] {
  return {
    // @ts-ignore TODO
    name: values.name,
    envs: values.envs,
    tables: getTablesFromServices(values.services, allServices),
  };
}

function getTablesFromServices(
  selectedServiceNames: string[],
  allServices: ServiceTypes,
): string[] {
  return selectedServiceNames.map((service: string) => allServices[service].tables).flat();
}
