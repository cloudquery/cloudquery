import test, { expect } from '@playwright/test';

test('Submit the form', async ({ page }) => {
  await page.goto('/');
  await page.getByRole('button', { name: 'Test connection' }).click();

  await expect(page.getByText('Testing the source connection')).toBeVisible();
});
