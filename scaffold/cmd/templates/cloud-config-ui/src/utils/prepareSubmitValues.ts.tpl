import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  return {
    name: values.name,
    spec: {
      token: values.token,
    },
    envs: [],

    // required for source plugin
    // tables: [],
    // skipTables: [],

    // required for destination plugin
    // migrateMode: 'forced',
    // writeMode: 'append',
  };
}
