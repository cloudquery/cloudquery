import { act } from 'react';
import { renderHook } from '@testing-library/react';
import { useFormSubmit } from '../useFormSubmit';
import {
  MessageHandler,
  FormMessagePayload,
  FormMessageType,
  PluginUiMessagePayload,
  PluginUiMessageType,
  formMessageTypes,
  pluginUiMessageTypes,
} from '@cloudquery/plugin-config-ui-connector';

const formMessageHandler = new MessageHandler<
  FormMessageType,
  FormMessagePayload,
  PluginUiMessageType,
  PluginUiMessagePayload
>(formMessageTypes, pluginUiMessageTypes, window.parent);

describe('useFormSubmit', () => {
  let originalPostMessage: typeof window.postMessage;

  afterEach(() => {
    window.parent.postMessage = originalPostMessage;
    jest.restoreAllMocks();
  });

  test('validation succeeded', async () => {
    const onValidate = jest.fn(() =>
      Promise.resolve({ values: { email: 'john@doe.com', name: 'John Doe' } }),
    );
    renderHook(() => useFormSubmit(onValidate as any));

    await act(async () => {
      formMessageHandler.sendMessage('validate', undefined);

      originalPostMessage = window.parent.postMessage;
      window.parent.postMessage = jest.fn();
      jest.spyOn(window, 'addEventListener');
      await new Promise((resolve) => setTimeout(resolve, 100));
    });

    expect(onValidate).toBeCalledTimes(1);
    expect(window.parent.postMessage).toBeCalledWith(
      {
        type: 'validation_passed',
        payload: { values: { key: 'value' } },
        id: expect.any(String),
      },
      '*',
    );
  });

  test('validation failed', async () => {
    const onValidate = jest.fn(() => Promise.resolve({ errors: { key: 'value' } }));
    renderHook(() => useFormSubmit(onValidate as any));

    await act(async () => {
      formMessageHandler.sendMessage('validate', undefined);

      originalPostMessage = window.parent.postMessage;
      window.parent.postMessage = jest.fn();
      jest.spyOn(window, 'addEventListener');
      await new Promise((resolve) => setTimeout(resolve, 100));
    });

    expect(onValidate).toBeCalledTimes(1);
    expect(window.parent.postMessage).toBeCalledWith(
      {
        type: 'validation_failed',
        payload: { errors: { key: 'value' } },
        id: expect.any(String),
      },
      '*',
    );
  });
});
