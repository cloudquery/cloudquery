const assert = require('node:assert/strict');
const fs = require('node:fs');
const os = require('node:os');
const path = require('node:path');
const { describe, it } = require('node:test');
const { eventCommits, listChangedFiles, triggerPaths, workflowFileOf } = require('./changed_files.js');

const writeWorkflow = (contents) => {
  const file = path.join(fs.mkdtempSync(path.join(os.tmpdir(), 'changed-files-')), 'workflow.yml');
  fs.writeFileSync(file, contents);
  return file;
};

const githubWithFiles = (files) => ({
  rest: {
    repos: {
      compareCommits: async ({ owner, repo, base, head }) => {
        assert.deepEqual({ owner, repo, base, head }, { owner: 'cloudquery', repo: 'cloudquery', base: 'base', head: 'head' });
        return { data: { files } };
      },
    },
  },
});

const repo = { owner: 'cloudquery', repo: 'cloudquery' };

describe('eventCommits', () => {
  it('reads the pull request commits', () => {
    const context = { eventName: 'pull_request', payload: { pull_request: { base: { sha: 'b' }, head: { sha: 'h' } } } };
    assert.deepEqual(eventCommits(context), { base: 'b', head: 'h' });
  });

  it('reads the merge group commits', () => {
    const context = { eventName: 'merge_group', payload: { merge_group: { base_sha: 'b', head_sha: 'h' } } };
    assert.deepEqual(eventCommits(context), { base: 'b', head: 'h' });
  });

  it('reads the push commits', () => {
    const context = { eventName: 'push', ref: 'refs/heads/main', payload: { before: 'b', after: 'h' } };
    assert.deepEqual(eventCommits(context), { base: 'b', head: 'h' });
  });

  it('returns undefined for a push that creates a branch', () => {
    const context = { eventName: 'push', ref: 'refs/heads/main', payload: { before: '0000000', after: 'h' } };
    assert.equal(eventCommits(context), undefined);
  });

  it('returns undefined for a tag push', () => {
    const context = { eventName: 'push', ref: 'refs/tags/v1.0.0', payload: { before: 'b', after: 'h' } };
    assert.equal(eventCommits(context), undefined);
  });

  it('returns undefined for other events', () => {
    assert.equal(eventCommits({ eventName: 'workflow_dispatch' }), undefined);
  });
});

describe('listChangedFiles', () => {
  const context = { eventName: 'merge_group', repo, payload: { merge_group: { base_sha: 'base', head_sha: 'head' } } };

  it('lists the changed files and the names they moved from', async () => {
    const github = githubWithFiles([
      { filename: 'cli/main.go' },
      { filename: 'docs/a.md', previous_filename: 'plugins/source/xkcd/a.md' },
    ]);
    assert.deepEqual(await listChangedFiles({ github, context }), ['cli/main.go', 'docs/a.md', 'plugins/source/xkcd/a.md']);
  });

  it('leaves out the names files moved from when asked', async () => {
    const github = githubWithFiles([{ filename: 'docs/a.md', previous_filename: 'plugins/source/xkcd/a.md' }]);
    assert.deepEqual(await listChangedFiles({ github, context, withPreviousNames: false }), ['docs/a.md']);
  });

  it('returns undefined when the comparison hits the file cap', async () => {
    const github = githubWithFiles(Array.from({ length: 300 }, (_, i) => ({ filename: `f${i}` })));
    assert.equal(await listChangedFiles({ github, context }), undefined);
  });

  it('returns undefined without commits to compare', async () => {
    assert.equal(await listChangedFiles({ github: githubWithFiles([]), context: { eventName: 'workflow_dispatch' } }), undefined);
  });
});

describe('triggerPaths', () => {
  it('reads the first paths filter', () => {
    const file = writeWorkflow(`on:
  pull_request:
    paths:
      - "plugins/source/xkcd/**"
      - '.github/workflows/source_xkcd.yml'
      - .github/cloudquery-cli-version
  push:
    paths:
      - "other/**"
`);
    assert.deepEqual(triggerPaths(file), ['plugins/source/xkcd/**', '.github/workflows/source_xkcd.yml', '.github/cloudquery-cli-version']);
  });

  it('fails without a paths filter', () => {
    assert.throws(() => triggerPaths(writeWorkflow('on:\n  pull_request:\n')), /no paths filter/);
  });

  it('fails on a negated path', () => {
    assert.throws(() => triggerPaths(writeWorkflow('on:\n  pull_request:\n    paths:\n      - "!a/**"\n')), /negated path/);
  });

  it('fails on an empty paths filter', () => {
    assert.throws(() => triggerPaths(writeWorkflow('on:\n  pull_request:\n    paths:\n  push:\n')), /empty paths filter/);
  });
});

describe('workflowFileOf', () => {
  it('reads the workflow file from a workflow ref', () => {
    assert.equal(
      workflowFileOf('cloudquery/cloudquery/.github/workflows/source_xkcd.yml@refs/heads/gh-readonly-queue/main/pr-1-abc'),
      '.github/workflows/source_xkcd.yml',
    );
  });
});
