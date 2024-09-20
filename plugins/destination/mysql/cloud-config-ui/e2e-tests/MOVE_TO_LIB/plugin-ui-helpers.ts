import cryptoRandomString from 'crypto-random-string';
import { expect, Frame, Page, test } from '@playwright/test';
import { click, fillInput, getMainTestUser } from './e2e-helpers';

export const getPersistentName = () => `name-${cryptoRandomString(10)}`;

type PluginControlOpts = {
  page: Page;
  kind: 'source' | 'destination';
  pluginName: string;
  pluginLabel: string;
  pluginNewName: string;
  fillFieldsSteps?: (iframeElement: Frame) => Promise<void>;
};

export const login = async (page: Page) => {
  await page.goto('https://cloud.cloudquery.io/auth/login');

  const { email, password } = getMainTestUser();

  await page.getByLabel('Email Address').click();
  await page.getByLabel('Email Address').fill(email);

  await page.locator(String.raw`button[type="submit"]`).click();

  await page.getByLabel('Password').click();
  await page.getByLabel('Password').fill(password);

  await page.locator(String.raw`button[type="submit"]`).click();

  await expect(page.getByRole('heading', { name: 'Overview' })).toBeVisible({ timeout: 5000 });

  await page.goto(process.env.CQ_CI_PLAYWRIGHT_PREVIEW_LINK ?? 'https://cloud.cloudquery.io');
};

export const createPlugin = async ({
  page,
  kind,
  pluginName,
  pluginNewName,
  pluginLabel,
  fillFieldsSteps,
}: PluginControlOpts) => {
  test.setTimeout(300_000);
  await page.goto(`https://cloud.cloudquery.io/teams/cq-bot-team/${kind}s/create`);

  await expect(page.getByText(`Create a ${kind}`)).toBeVisible();

  await fillInput(page, 'input[name="search"]', pluginName);
  await click(page, page.getByRole('button', { name: pluginLabel }));

  await expect(page.getByText(pluginLabel)).toBeTruthy();
  await expect(page.locator('iframe[name="Plugin UI"]')).toBeVisible({ timeout: 30_000 });
  const iframeElement = page.frame({ name: 'Plugin UI' });

  if (!iframeElement) {
    throw new Error('iframe not found');
  }

  await expect(iframeElement.getByText('Previewing')).toBeVisible();
  await fillInput(iframeElement, '[name="displayName"]', pluginNewName);
  await fillFieldsSteps?.(iframeElement);

  await click(iframeElement, iframeElement.getByRole('button', { name: 'Test connection' }));
  await expect(iframeElement.locator('div:has-text("Connection test")')).toBeTruthy();
  await expect(iframeElement.locator('button:has-text("Cancel test")')).toBeTruthy();
  await expect(page.getByText(`Edit ${kind}`)).toBeVisible({
    timeout: 30_000,
  });
};

export const editPlugin = async ({
  page,
  kind,
  pluginNewName,
  fillFieldsSteps,
}: PluginControlOpts) => {
  test.setTimeout(300_000);
  await page.goto(`https://cloud.cloudquery.io/teams/cq-bot-team/${kind}s`);

  await fillInput(page, 'input[type="text"]', pluginNewName);
  await click(page, page.getByText(pluginNewName));

  await expect(page.getByText(pluginNewName)).toBeTruthy();
  await page.getByRole('tab', { name: `Edit ${kind}` }).click();
  await expect(page.locator('iframe[name="Plugin UI"]')).toBeVisible({ timeout: 30_000 });
  const iframeElement = page.frame({ name: 'Plugin UI' });

  if (!iframeElement) {
    throw new Error('iframe not found');
  }

  await expect(iframeElement.getByText('Previewing')).toBeVisible();
  await expect(
    iframeElement.getByRole('textbox', {
      name: `${kind === 'destination' ? 'Destination' : 'Source'} name`,
    }),
  ).toHaveValue(pluginNewName);

  await fillFieldsSteps?.(iframeElement);

  await click(iframeElement, iframeElement.getByRole('button', { name: 'Test connection' }));
  await expect(iframeElement.locator('div:has-text("Connection test")')).toBeTruthy();
  await expect(iframeElement.locator('button:has-text("Cancel test")')).toBeTruthy();
  await expect(page.getByText(`Edit ${kind}`)).toBeVisible({
    timeout: 30_000,
  });
};

export const deletePlugin = async ({ page, kind, pluginNewName }: PluginControlOpts) => {
  test.setTimeout(300_000);
  await page.goto(`https://cloud.cloudquery.io/teams/cq-bot-team/${kind}s`);

  await fillInput(page, 'input[type="text"]', pluginNewName);
  await click(page, page.getByText(pluginNewName));

  await expect(page.getByText(pluginNewName)).toBeTruthy();
  await page.getByRole('tab', { name: `Edit ${kind}` }).click();
  await expect(page.locator('iframe[name="Plugin UI"]')).toBeVisible({ timeout: 30_000 });
  const iframeElement = page.frame({ name: 'Plugin UI' });

  if (!iframeElement) {
    throw new Error('iframe not found');
  }

  await expect(iframeElement.getByText('Previewing')).toBeVisible();
  await expect(
    iframeElement.getByRole('textbox', {
      name: `${kind === 'destination' ? 'Destination' : 'Source'} name`,
    }),
  ).toHaveValue(pluginNewName);

  await iframeElement.getByRole('button', { name: `Delete this ${kind}` }).click();
  await click(page, page.getByText(`Delete ${kind}`));

  await expect(page.getByText(pluginNewName)).toHaveCount(0);
};
