import test, { expect } from '@playwright/test';

test('Submit the form', async ({ page }) => {
  await page.goto('/');
  await page.getByLabel('Start time', { exact: true }).click();
  await page.getByLabel('Start time', { exact: true }).fill('07/07/2024 10:00 PM');
  await page.getByLabel('Item concurrency', { exact: true }).click();
  await page.getByLabel('Item concurrency', { exact: true }).fill('50');
  await page.getByRole('button', { name: 'Test connection' }).click();

  await expect(page.getByText('Testing the source connection')).toBeVisible();
});
