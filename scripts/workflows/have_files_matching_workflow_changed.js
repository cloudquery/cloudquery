module.exports = async ({github, context}) => {
  const files = process.env.FILES.split(' ')
  console.log(files)
  if (files.length >= 300) {
    // This is a GitHub limitation https://github.com/cloudquery/cloudquery/issues/2688
    throw new Error('Too many files changed. Skipping check. Please split your PR into multiple ones.')
  }

  const matchesWorkflow = (file, action) => {
    if (!file.startsWith('.github/workflows/')) {
      return false
    }
    try {
      const contents = fs.readFileSync(file, 'utf8');
      return contents.includes(`name: "${action}"`)
    } catch {
      return false
    }
  }
  const matchesFile = (action) => {
    return files.some(file => file.startsWith(`${action}/`) || matchesWorkflow(file, action))
  }

  if (matchesFile(context.github.action)) {
    return true
  }
  return false
}