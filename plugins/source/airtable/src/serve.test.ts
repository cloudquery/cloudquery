import test from 'ava';
import { pathExists } from 'path-exists';
import { temporaryDirectoryTask } from 'tempy';

import { createAirtableServeCommand } from './serve.js';

const serve = createAirtableServeCommand().exitProcess(false);

test('should build docker', async (t) => {
  delete process.env.CQ_TELEMETRY_LEVEL;
  await temporaryDirectoryTask(async (outputDirectory) => {
    await serve.parse(['package', '-m', 'test', 'v1.0.0', '.', '--dist-dir', outputDirectory, '--log-level', 'debug']);
    t.true(await pathExists(`${outputDirectory}/tables.json`));
    t.true(await pathExists(`${outputDirectory}/package.json`));
    t.true(await pathExists(`${outputDirectory}/plugin-airtable-v1.0.0-linux-amd64.tar`));
    t.true(await pathExists(`${outputDirectory}/plugin-airtable-v1.0.0-linux-arm64.tar`));
    t.true(await pathExists(`${outputDirectory}/docs/overview.md`));
  });
});
