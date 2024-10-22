import test, { expect } from '@playwright/test';

// TODO: rewrite as a test aginst production - live API Key needed.
test('Submit the form', async ({ page }) => {
  await page.goto('/');
  await page.getByLabel('Hosts', { exact: true }).click();
  await page.getByLabel('Hosts', { exact: true }).fill('localhost:9000');
  await page.getByLabel('Database *', { exact: true }).click();
  await page.getByLabel('Database *', { exact: true }).fill('my_db');
  await page.getByLabel('Username *', { exact: true }).click();
  await page.getByLabel('Username *', { exact: true }).fill('admin');
  await page.getByLabel('Password *', { exact: true }).click();
  await page.getByLabel('Password *', { exact: true }).fill('abc123');

  await page.getByRole('button', { name: 'Test connection' }).click();

  await expect(page.getByText('Testing the destination connection')).toBeVisible();
});
