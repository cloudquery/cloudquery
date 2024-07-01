import { act } from 'react';
import { renderHook } from '@testing-library/react';
import { useFormInit } from '../useFormInit';
import {
  MessageHandler,
  FormMessagePayload,
  FormMessageType,
  PluginUiMessagePayload,
  PluginUiMessageType,
  formMessageTypes,
  pluginUiMessageTypes,
} from '@cloudquery/plugin-config-ui-connector';

test('useFormInit', async () => {
  const { rerender, result } = renderHook(useFormInit);
  expect(result.current).toEqual({ initialValues: undefined, muiThemeOptions: undefined });

  await act(async () => {
    const formMessageHandler = new MessageHandler<
      FormMessageType,
      FormMessagePayload,
      PluginUiMessageType,
      PluginUiMessagePayload
    >(formMessageTypes, pluginUiMessageTypes, window.parent);
    formMessageHandler.sendMessage('init', { initialValues: undefined });
    await new Promise((resolve) => setTimeout(resolve, 0));
  });

  rerender();
  expect(result.current).toEqual({ initialValues: undefined });
});
