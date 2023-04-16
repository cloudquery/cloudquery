module.exports = () => {
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
    const sources = fs.readdirSync("plugins/source", {withFileTypes: true}).filter(dirent => dirent.isDirectory()).map(dirent => `plugins/source/${dirent.name}`)
    const destinations = fs.readdirSync("plugins/destination", {withFileTypes: true}).filter(dirent => dirent.isDirectory()).map(dirent => `plugins/destination/${dirent.name}`)
    const allComponents = ["cli", "scaffold", ...sources, ...destinations]
    console.log(`All components: ${allComponents.join(", ")}`)
    let actions = allComponents.filter(action => matchesFile(action))
    if (actions.length === 0) {
        console.log("No actions to wait for")
        return
    }

    // Most modules should have a 'validate-release' job
    for (const action of actions) {
        actions = [...actions, 'validate-release']
    }

    // We test the CLI on multiple OSes, so we need to wait for all of them
    if (actions.includes("cli")) {
        actions = actions.filter(action => action !== "cli")
        actions = ["cli (ubuntu-latest)", "cli (windows-latest)", "cli (macos-latest)", ...actions]
    }

    // Enforce policy tests for AWS,GCP and K8s plugins
    const pluginsWithPolicyTests = ["plugins/source/aws", "plugins/source/azure", "plugins/source/gcp", "plugins/source/k8s"]
    for (const plugin of pluginsWithPolicyTests) {
        if (actions.includes(plugin)) {
            actions = [...actions, 'test-policies']
        }
    }

    console.log(`Required actions: [${actions.join(", ")}]`)
    return actions
}