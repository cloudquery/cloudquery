import { useEffect, useState } from 'react';
import { pluginUiMessageHandler } from '../messageHandler';
import { FormValues } from '../../types';

export function useFormInit() {
  const [initialized, setInitialized] = useState(false);
  const [initialValues, setInitialValues] = useState<FormValues | undefined>();

  useEffect(() => {
    return pluginUiMessageHandler.subscribeToMessage('init', ({ initialValues }) => {
      const values = initialValues as FormValues;
      setInitialValues(values);
      setInitialized(true);
    });
  }, []);

  return { initialized, initialValues };
}
