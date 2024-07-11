import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  return {
    spec: {
      token: values.token,
    },
    envs: [],

    // required for source plugin
    // tables: [],

    // required for destination plugin
    // migrateMode: 'forced',
    // writeMode: 'append',
  };
}
