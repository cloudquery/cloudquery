import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareInitialValues(
  initialValues: Exclude<FormMessagePayload['init']['initialValues'], undefined>,
): FormValues {
  return {
    token: initialValues.spec?.token || '',
  };
}
