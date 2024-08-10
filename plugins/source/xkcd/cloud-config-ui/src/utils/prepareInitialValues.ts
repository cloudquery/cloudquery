import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  return {
    name: initialValues?.name || '',
    envs: initialValues?.envs || [],
    spec: {},
    tables: Object.fromEntries((initialValues?.tables || []).map((table) => [table, true])),
  };
}
