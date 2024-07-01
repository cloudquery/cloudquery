import { useEffect, useState } from 'react';
import { pluginUiMessageHandler } from '../messageHandler';
import { FormValues } from '../formSchema';
import { prepareInitialValues } from '../prepareInitialValues';

export function useFormInit() {
  const [initialized, setInitialized] = useState(false);
  const [initialValues, setInitialValues] = useState<FormValues | undefined>();

  useEffect(() => {
    return pluginUiMessageHandler.subscribeToMessage('init', ({ initialValues }) => {
      if (initialValues?.spec?.connection_string) {
        setInitialValues(prepareInitialValues(initialValues));
      }

      setInitialized(true);
    });
  }, []);

  return { initialized, initialValues };
}
