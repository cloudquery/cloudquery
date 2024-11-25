module.exports = async ({github, context}) => {
    const files = process.env.FILES.split(' ')
    console.log(files)
    if (files.length >= 300) {
        // This is a GitHub limitation https://github.com/cloudquery/cloudquery/issues/2688
        throw new Error('Too many files changed. Skipping check. Please split your PR into multiple ones.')
    }

    const fs = require("fs");
    let now = new Date().getTime()
    const deadline = now + 60 * 1000 * 50
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

    const getPluginsWithConfigUI = () => {
        const cloudConfigUIDir = "/cloud-config-ui"
        const configUIDirectories = fs.readdirSync("plugins", {recursive: true}).filter(directory => directory.endsWith(cloudConfigUIDir))
        const plugins = configUIDirectories.map(directory => `plugins/${directory.replace(cloudConfigUIDir, "")}`)
        console.log(`Found the following plugins with config UI: ${plugins.join(", ")}`)
        return plugins
    } 

    const sources = fs.readdirSync("plugins/source", {withFileTypes: true}).filter(dirent => dirent.isDirectory()).map(dirent => `plugins/source/${dirent.name}`)
    const destinations = fs.readdirSync("plugins/destination", {withFileTypes: true}).filter(dirent => dirent.isDirectory()).map(dirent => `plugins/destination/${dirent.name}`)
    const transformers = fs.readdirSync("plugins/transformer", {withFileTypes: true}).filter(dirent => dirent.isDirectory()).map(dirent => `plugins/transformer/${dirent.name}`)
    const allComponents = ["cli", "scaffold", ...sources, ...destinations, ...transformers]
    console.log(`All components: ${allComponents.join(", ")}`)
    let actions = allComponents.filter(action => matchesFile(action))
    if (actions.length === 0) {
        console.log("No actions to wait for")
        return
    }

    if (actions.includes("scaffold")) {
        actions = [...actions, 'validate-release']
    }

    // We test the CLI on multiple OSes, so we need to wait for all of them
    if (actions.includes("cli")) {
        actions = [...actions, 'validate-release']
        actions = actions.filter(action => action !== "cli")
        actions = ["cli (ubuntu-latest)", "cli (windows-latest)", "cli (macos-latest)", ...actions]
    }

    const pluginsWithConfigUI = getPluginsWithConfigUI()
    for (const action of actions) {
        if (pluginsWithConfigUI.includes(action)) {
            console.log(`Adding validate-config-ui to the list of required workflows for plugin ${action}`)
            actions = [...actions, 'validate-config-ui']
        }
    }

    const pluginsWithFipsVersion = ["plugins/source/test"]
    for (const action of actions) {
        if (pluginsWithFipsVersion.includes(action)) {
            console.log(`Adding validate-fips to the list of required workflows for plugin ${action}`)
            actions = [...actions, 'validate-fips']
        }
    }

    pendingActions = [...actions]
    console.log(`Waiting for ${pendingActions.join(", ")}`)
    while (now <= deadline) {
        const checkRuns = await github.paginate(github.rest.checks.listForRef, {
            owner: 'cloudquery',
            repo: context.repo.repo,
            ref: context.payload.pull_request.head.sha,
            status: 'completed',
            per_page: 100
        })
        const runsWithPossibleDuplicates = checkRuns.map(({id, name, conclusion}) => ({id, name, conclusion}))
        const runs = runsWithPossibleDuplicates.filter((run, index, self) => self.findIndex(({id}) => id === run.id) === index)
        console.log(`Got the following check runs: ${JSON.stringify(runs)}`)
        const matchingRuns = runs.filter(({name}) => actions.includes(name))
        const allowedConclusions = ['success', 'skipped']
        const failedRuns = matchingRuns.filter(({conclusion}) => !allowedConclusions.includes(conclusion))
        if (failedRuns.length > 0) {
            throw new Error(`The following required workflows failed: ${failedRuns.map(({name}) => name).join(", ")}`)
        }
        console.log(`Matching runs: ${matchingRuns.map(({name}) => name).join(", ")}`)
        console.log(`Actions: ${actions.join(", ")}`)
        if (matchingRuns.length === actions.length) {
            console.log("All required workflows have passed")
            return
        }
        pendingActions = actions.filter(action => !runs.some(({name}) => name === action))
        console.log(`Waiting for ${pendingActions.join(", ")}`)
        await new Promise(r => setTimeout(r, 60000));
        now = new Date().getTime()
    }
    throw new Error(`Timed out waiting for ${pendingActions.join(', ')}`)
}
