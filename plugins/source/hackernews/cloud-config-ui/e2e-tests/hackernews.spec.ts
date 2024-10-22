import { Frame, test } from '@playwright/test';

import {
  createPlugin,
  deletePlugin,
  editPlugin,
  fillInput,
  getPersistentName,
  login,
} from '@cloudquery/plugin-config-ui-lib/e2e-utils';

test.describe.configure({ mode: 'serial' });

test.describe('HackerNews Source', () => {
  const parameters = {
    pluginNewName: getPersistentName(),
    kind: 'source' as 'source',
    pluginName: 'hackernews',
    pluginLabel: 'HackerNews',
    pluginUrl: '',
  };

  test.beforeEach('login', async ({ page }) => {
    await login(page);
  });

  test('create plugin', async ({ page }) => {
    parameters.pluginUrl = await createPlugin({
      ...parameters,
      page,
      fillFieldsSteps: async (iframeElement: Frame) => {
        await fillInput(
          iframeElement,
          iframeElement.getByLabel('Start time'),
          '07/07/2024 10:00 PM',
        );

        await fillInput(iframeElement, iframeElement.getByLabel('Item concurrency'), '50');
      },
    });
  });

  test('edit plugin', async ({ page }) => {
    await editPlugin({
      ...parameters,
      page,
      fillFieldsSteps: async (iframeElement: Frame) => {
        await fillInput(iframeElement, iframeElement.getByLabel('Item concurrency'), '150');
      },
    });
  });

  test('delete plugin', async ({ page }) => {
    await deletePlugin({
      ...parameters,
      page,
    });
  });
});
