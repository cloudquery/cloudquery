import 'dotenv/config';

import { defineConfig, devices } from '@playwright/test';

/**
 * See https://playwright.dev/docs/test-configuration.
 */
export default defineConfig({
  forbidOnly: true,
  fullyParallel: false,
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },
  ],
  reporter: 'html',
  retries: 0,
  testDir: './e2e-tests',
  timeout: 2 * 60 * 1000,
  use: {
    baseURL: 'http://localhost:3000',
    headless: process.env.CI ? true : false,
    trace: process.env.CI ? 'on-first-retry' : 'retain-on-failure',
    video: {
      mode: process.env.CI ? 'on-first-retry' : 'retain-on-failure',
      size: { height: 480, width: 640 },
    },
  },
  workers: 1,
  webServer: {
    command: 'http-server ./build -p 3000 --silent',
    reuseExistingServer: !process.env.CI,
    stderr: 'pipe',
    stdout: 'pipe',
    url: 'http://localhost:3000',
  },
});
