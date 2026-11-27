const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const { describe, it } = require('node:test');
const waitForRequiredWorkflows = require('./wait_for_required_workflows.js');

process.chdir(path.join(__dirname, '..', '..'));

const pullRequestEvent = {
  eventName: 'pull_request',
  payload: { pull_request: { base: { sha: 'base' }, head: { sha: 'head' } } },
};

const mergeGroupEvent = {
  eventName: 'merge_group',
  payload: { merge_group: { base_sha: 'base', head_sha: 'head' } },
};

const workflowFiles = fs.readdirSync('.github/workflows').filter((file) => file.endsWith('.yml'));
const readWorkflow = (file) => fs.readFileSync(path.join('.github/workflows', file), 'utf8');
const componentName = /^    name: ["'](cli|scaffold|plugins\/(?:source|destination|transformer)\/[^/"']+)["']$/m;
const componentWorkflows = workflowFiles.filter((file) => componentName.test(readWorkflow(file)));
const componentsWithWorkflow = componentWorkflows.map((file) => readWorkflow(file).match(componentName)[1]);
const allCheckNames = [
  ...componentsWithWorkflow.filter((name) => name !== 'cli'),
  'cli (ubicloud-standard-2)',
  'cli (windows-latest)',
  'cli (macos-latest)',
  'validate-release',
  'validate-fips',
];

const run = async ({ event, files, checkRuns = allCheckNames.map((name) => ({ name, conclusion: 'success' })) }) => {
  const checkRefs = [];
  const github = {
    rest: {
      repos: {
        compareCommits: async ({ base, head }) => {
          assert.deepEqual({ base, head }, { base: 'base', head: 'head' });
          return { data: { files: files.map((file) => (typeof file === 'string' ? { filename: file } : file)) } };
        },
      },
      checks: { listForRef: 'listForRef' },
    },
    paginate: async (_, { ref }) => {
      checkRefs.push(ref);
      return checkRuns.map((checkRun, id) => ({ id, ...checkRun }));
    },
  };
  const context = { ...event, repo: { owner: 'cloudquery', repo: 'cloudquery' } };
  const log = console.log;
  const lines = [];
  console.log = (...args) => lines.push(args.join(' '));
  try {
    await waitForRequiredWorkflows({ github, context });
  } finally {
    console.log = log;
  }
  const waitingFor = lines.find((line) => line.startsWith('Waiting for '));
  return {
    checkRefs,
    waitingFor: waitingFor ? waitingFor.replace('Waiting for ', '').split(', ').sort() : [],
  };
};

const manyFiles = Array.from({ length: 300 }, (_, i) => `docs/f${i}.md`);

describe('wait_for_required_workflows', () => {
  it('reads the merge group commits and its head check runs', async () => {
    const { checkRefs, waitingFor } = await run({ event: mergeGroupEvent, files: ['plugins/source/xkcd/main.go'] });
    assert.deepEqual(checkRefs, ['head']);
    assert.deepEqual(waitingFor, ['plugins/source/xkcd']);
  });

  it('reads the pull request head check runs', async () => {
    const { checkRefs, waitingFor } = await run({ event: pullRequestEvent, files: ['plugins/source/xkcd/main.go'] });
    assert.deepEqual(checkRefs, ['head']);
    assert.deepEqual(waitingFor, ['plugins/source/xkcd']);
  });

  it('waits for every CLI operating system and the release validation', async () => {
    const { waitingFor } = await run({ event: mergeGroupEvent, files: ['cli/main.go'] });
    assert.deepEqual(waitingFor, ['cli (macos-latest)', 'cli (ubicloud-standard-2)', 'cli (windows-latest)', 'validate-release']);
  });

  it('waits for the FIPS validation of a plugin with a FIPS version', async () => {
    const { waitingFor } = await run({ event: mergeGroupEvent, files: ['plugins/destination/postgresql/main.go'] });
    assert.deepEqual(waitingFor, ['plugins/destination/postgresql', 'validate-fips']);
  });

  for (const [workflow, component] of [['source_xkcd.yml', 'plugins/source/xkcd'], ['source_hackernews.yml', 'plugins/source/hackernews']]) {
    it(`waits for ${component} when ${workflow} changes`, async () => {
      const { waitingFor } = await run({ event: mergeGroupEvent, files: [`.github/workflows/${workflow}`] });
      assert.deepEqual(waitingFor, [component]);
    });
  }

  it('waits for the component a merge group moves a file out of', async () => {
    const { waitingFor } = await run({
      event: mergeGroupEvent,
      files: [{ filename: 'docs/a.md', previous_filename: 'plugins/source/xkcd/a.md' }],
    });
    assert.deepEqual(waitingFor, ['plugins/source/xkcd']);
  });

  it('ignores the component a pull request moves a file out of', async () => {
    const { waitingFor } = await run({
      event: pullRequestEvent,
      files: [{ filename: 'docs/a.md', previous_filename: 'plugins/source/xkcd/a.md' }],
    });
    assert.deepEqual(waitingFor, []);
  });

  it('waits for nothing when no component changes', async () => {
    const { checkRefs, waitingFor } = await run({ event: mergeGroupEvent, files: ['README.md'] });
    assert.deepEqual(checkRefs, []);
    assert.deepEqual(waitingFor, []);
  });

  it('waits for nothing when a plugin without a workflow changes', async () => {
    const { checkRefs } = await run({ event: mergeGroupEvent, files: ['plugins/source/aws/views/a.sql'] });
    assert.deepEqual(checkRefs, []);
  });

  it('waits for every component with a workflow when a merge group has too many files to list', async () => {
    const { waitingFor } = await run({ event: mergeGroupEvent, files: manyFiles });
    assert.deepEqual(waitingFor, [...allCheckNames].sort());
  });

  it('fails a pull request with too many files to list', async () => {
    await assert.rejects(run({ event: pullRequestEvent, files: manyFiles }), /Too many files changed/);
  });

  it('passes when more than one workflow reports the same check name', async () => {
    const checkRuns = [
      { name: 'plugins/destination/postgresql', conclusion: 'success' },
      { name: 'plugins/destination/kafka', conclusion: 'success' },
      { name: 'validate-fips', conclusion: 'success' },
      { name: 'validate-fips', conclusion: 'skipped' },
    ];
    const files = ['plugins/destination/postgresql/main.go', 'plugins/destination/kafka/main.go'];
    await run({ event: mergeGroupEvent, files, checkRuns });
  });

  it('fails when a required check fails', async () => {
    const checkRuns = [{ name: 'plugins/source/xkcd', conclusion: 'failure' }];
    await assert.rejects(
      run({ event: mergeGroupEvent, files: ['plugins/source/xkcd/main.go'], checkRuns }),
      /required workflows failed: plugins\/source\/xkcd/,
    );
  });
});

describe('component workflows', () => {
  it('find the component workflows', () => {
    assert.equal(componentWorkflows.length, 33);
  });

  for (const file of componentWorkflows) {
    it(`${file} runs in the merge queue and skips when not affected`, () => {
      const contents = readWorkflow(file);
      assert.match(contents, /^  merge_group:$/m);
      assert.match(contents, /^  changes:\n    if: github.event_name == 'merge_group'\n    uses: \.\/\.github\/workflows\/changed_files\.yml\n    with:\n      trigger_paths: true$/m);
    });
  }
});
