const fs = require('fs');

const COMPARE_FILES_CAP = 300;

const eventCommits = (context) => {
  if (context.eventName === 'pull_request') {
    return { base: context.payload.pull_request.base.sha, head: context.payload.pull_request.head.sha };
  }
  if (context.eventName === 'merge_group') {
    return { base: context.payload.merge_group.base_sha, head: context.payload.merge_group.head_sha };
  }
  if (context.eventName === 'push' && context.ref.startsWith('refs/heads/') && !/^0+$/.test(context.payload.before)) {
    return { base: context.payload.before, head: context.payload.after };
  }
  return undefined;
};

const listChangedFiles = async ({ github, context, withPreviousNames = true }) => {
  const commits = eventCommits(context);
  if (!commits) return undefined;

  const { data } = await github.rest.repos.compareCommits({ ...context.repo, ...commits });
  if (data.files.length >= COMPARE_FILES_CAP) return undefined;
  return [...new Set(data.files.flatMap((f) => [f.filename, withPreviousNames && f.previous_filename].filter(Boolean)))];
};

const triggerPaths = (workflowFile) => {
  const lines = fs.readFileSync(workflowFile, 'utf8').split('\n');
  const start = lines.findIndex((line) => /^\s+paths:\s*$/.test(line));
  if (start === -1) throw new Error(`no paths filter in ${workflowFile}`);

  const patterns = [];
  for (const line of lines.slice(start + 1)) {
    const match = line.match(/^\s+-\s+["']?([^"'\s]+)["']?\s*$/);
    if (!match) break;
    if (match[1].startsWith('!')) throw new Error(`negated path ${match[1]} in ${workflowFile} is not supported`);
    patterns.push(match[1]);
  }
  if (patterns.length === 0) throw new Error(`empty paths filter in ${workflowFile}`);
  return patterns;
};

const workflowFileOf = (workflowRef) => workflowRef.replace(/^[^/]+\/[^/]+\//, '').replace(/@.*$/, '');

module.exports = { eventCommits, listChangedFiles, triggerPaths, workflowFileOf };
