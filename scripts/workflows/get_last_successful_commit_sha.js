module.exports = async ({github, context}) => {
  const allCommits = context.github.pull_request.commits;
  // loop from second-to-last commit to first
  for (const commit of allCommits.reverse().slice(1)) {
    const checkRuns = await github.paginate(github.rest.checks.listForRef, {
      owner: 'cloudquery',
      repo: context.repo.repo,
      ref: commit.sha,
      status: 'completed',
      per_page: 100
    })
    for (const checkRun of checkRuns) {
      if (checkRun.conclusion === 'success') {
        return checkRun.head_sha
      }
    }
  }
  // by default return base sha
  return context.github.pull_request.base.sha
}