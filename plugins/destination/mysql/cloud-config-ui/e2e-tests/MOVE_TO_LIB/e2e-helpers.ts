import { Page, Locator, Frame, expect } from '@playwright/test';

export const getMainTestUser = () => {
  if (!process.env.CQ_CI_PLAYWRIGHT_TEST_USER_EMAIL) {
    throw new Error('CQ_CI_PLAYWRIGHT_TEST_USER_EMAIL is not set');
  }
  if (!process.env.CQ_CI_PLAYWRIGHT_TEST_USER_PASSWORD) {
    throw new Error('CQ_CI_PLAYWRIGHT_TEST_USER_PASSWORD is not set');
  }

  return {
    email: process.env.CQ_CI_PLAYWRIGHT_TEST_USER_EMAIL,
    password: process.env.CQ_CI_PLAYWRIGHT_TEST_USER_PASSWORD,
  };
};

export function getRootUrl() {
  return process.env.CQ_CI_PLAYWRIGHT_PREVIEW_LINK
    ? process.env.CQ_CI_PLAYWRIGHT_PREVIEW_LINK.replace('cloudquery-test', 'cq-bot-team')
    : 'https://cloud.cloudquery.io';
}

export function getPluginUrl(url: string) {
  if (process.env.CQ_CI_PLAYWRIGHT_PREVIEW_LINK) {
    return `${url}?${process.env.CQ_CI_PLAYWRIGHT_PREVIEW_LINK.split('?')[1]}`;
  } else {
    return url;
  }
}

export function getRandomTestUser(browserType: string) {
  return {
    email: `e2e-test-${process.env.GITHUB_JOB_ID}-${browserType}@cloudquery.io`,
    newPassword: 'A87654321a',
    password: '12345678Aa',
  };
}

export async function assertUrlPathname(page: Page, value: string) {
  await page.waitForTimeout(500);
  await page.waitForURL((url) => url.pathname === value);
}

export async function assertUrlParam(page: Page, key: string, value: string | null) {
  await page.waitForTimeout(500);
  await page.waitForFunction(
    ({ key, value }) => new URL(document.location.href).searchParams.get(key) === value,
    { key, value },
  );
}

export async function fillInput(page: Page | Frame, selector: string | Locator, value: string) {
  const locator = typeof selector === 'string' ? page.locator(selector) : selector;
  expect(locator).toBeVisible();
  await locator.focus();
  await locator.clear();
  if (value) {
    await ('keyboard' in page ? page.keyboard.type(value) : locator.fill(value));
  }
  await expect(typeof selector === 'string' ? page.locator(selector) : selector).toHaveValue(value);
}

async function waitForLoadState(page: Page) {
  await page.waitForLoadState('domcontentloaded', { timeout: 15_000 });
  await page.waitForLoadState('load', { timeout: 15_000 });
}

export async function goTo(page: Page, url: string) {
  await page.goto(url);
  await page.waitForTimeout(1000);
  await waitForLoadState(page);
}

export async function click(_: Page | Frame, element: Locator) {
  await expect(element).toBeVisible();
  await element.focus();
  await element.click();
}
