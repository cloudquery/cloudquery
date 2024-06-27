import { useEffect } from 'react';
import { pluginUiMessageHandler } from '../messageHandler';
import { FormValues } from '../../types';

interface Success {
  values: FormValues;
  errors?: never;
}

interface Failure {
  values?: never;
  errors: Record<string, any>;
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
