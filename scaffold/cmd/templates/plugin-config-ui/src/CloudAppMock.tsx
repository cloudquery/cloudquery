import {
  MessageHandler,
  FormMessageType,
  FormMessagePayload,
  PluginUiMessageType,
  PluginUiMessagePayload,
  formMessageTypes,
  pluginUiMessageTypes,
} from '@cloudquery/plugin-config-ui-connector';
import { Button, Stack } from '@mui/material';
import { ReactNode, useEffect, useState } from 'react';

interface Props {
  children: ReactNode;
}

const formMessageHandler = new MessageHandler<
  FormMessageType,
  FormMessagePayload,
  PluginUiMessageType,
  PluginUiMessagePayload
>(formMessageTypes, pluginUiMessageTypes, window);

export function CloudAppMock({ children }: Props) {
  const [values, setValues] = useState<string>('');
  const [errors, setErrors] = useState<string>('');

  useEffect(() => {
    formMessageHandler.sendMessage('init', {
      initialValues: {
        migrateMode: undefined,
        tables: undefined,
        writeMode: undefined,
        spec: {},
        envs: [
          {
            name: 'username',
            value: '',
          },
          {
            name: 'password',
            value: '',
          },
        ],
      },
    });
  }, []);

  const handleSubmit = async () => {
    formMessageHandler.sendMessage('validate', undefined);
    let unsubscribeValidationPassed: (() => void) | undefined;
    let unsubscribeValidationFailed: (() => void) | undefined;

    try {
      const values = await new Promise((resolve, reject) => {
        unsubscribeValidationPassed = formMessageHandler.subscribeToMessageOnce(
          'validation_passed',
          ({ values }) => {
            resolve(values);
          },
        );
        unsubscribeValidationFailed = formMessageHandler.subscribeToMessageOnce(
          'validation_failed',
          ({ errors }) => reject(errors),
        );
      }).finally(() => {
        unsubscribeValidationPassed?.();
        unsubscribeValidationFailed?.();
      });

      setErrors('');
      setValues(JSON.stringify(values, null, 2));
    } catch (error) {
      unsubscribeValidationPassed?.();
      unsubscribeValidationFailed?.();

      setValues('');
      setErrors(JSON.stringify(error, null, 2));
    }
  };

  return (
    <>
      {children}
      <Stack direction="row" justifyContent="flex-end" spacing={2} padding={2}>
        <Button onClick={handleSubmit} variant="contained">
          Submit
        </Button>
      </Stack>
      <Stack padding={2}>
        <div>Values:</div>
        <pre style={{ wordBreak: 'break-all', whiteSpace: 'break-spaces' }}>{values || '-'}</pre>
        <div>Errors:</div>
        <pre style={{ wordBreak: 'break-all', whiteSpace: 'break-spaces' }}>{errors || '-'}</pre>
      </Stack>
    </>
  );
}
