import { useEffect } from 'react';
import { pluginUiMessageHandler } from '../messageHandler';
import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';

interface Success {
  values: PluginUiMessagePayload['validation_passed']['values'];
  errors?: never;
}

interface Failure {
  values?: never;
  errors: PluginUiMessagePayload['validation_failed']['errors'];
}

export function useFormSubmit(onValidate: () => Promise<Success | Failure>) {
  useEffect(() => {
    const handleValidate = async () => {
      const { errors, values } = await onValidate();

      if (errors) {
        pluginUiMessageHandler.sendMessage('validation_failed', {
          errors,
        });
      } else {
        pluginUiMessageHandler.sendMessage('validation_passed', {
          values,
        });
      }
    };

    return pluginUiMessageHandler.subscribeToMessage('validate', handleValidate);
  }, [onValidate]);
}
