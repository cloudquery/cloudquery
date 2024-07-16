import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';
import { serviceNameResolutions } from './constants';

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  return {
    ...(initialValues as any),
    services: convertTablesToServices(initialValues?.tables),
    all_tables: [],
  };
}

function convertTablesToServices(tables?: string[]): string[] {
  return Array.from(
    new Set(
      tables?.map((table) => {
        const serviceNames = table.split('_')[1];
        return serviceNameResolutions[serviceNames] ?? serviceNames;
      }),
    ),
  );
}
