import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareInitialValues(
  initialValues: FormMessagePayload['init']['initialValues'],
): FormValues {
  return {
    ...(initialValues as any),
    arn: getArnFromAccounts(initialValues?.spec?.accounts),
    services: convertTablesToServices(initialValues?.tables),
    all_tables: [],
    spec: initialValues?.spec ?? {
      regions: [],
      accounts: [{ role_arn: '' }],
    },
  };
}

function convertTablesToServices(tables?: string[]): string[] {
  return Array.from(new Set(tables?.map((table) => table.split('_')[1])));
}

function getArnFromAccounts(accounts: { role_arn: string }[]) {
  return accounts?.[0]?.role_arn ?? '';
}
