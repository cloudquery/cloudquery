module.exports = async ({github, context}) => {
    const actions = JSON.parse(process.env.ACTIONS)
    if (actions.length === 0) {
        console.log("No actions to wait for")
        return
    }
    console.log(`Required actions: [${actions.join(", ")}]`)

    let now = new Date().getTime()
    const deadline = now + 60 * 1000 * 50
    pendingActions = [...actions]
    console.log(`Waiting for ${pendingActions.join(", ")}`)

    while (now <= deadline) {
        const checkRuns = await github.paginate(github.rest.checks.listForRef, {
            owner: 'cloudquery',
            repo: 'cloudquery',
            ref: context.payload.pull_request.head.sha,
            status: 'completed',
            per_page: 100
        })
        const runsWithPossibleDuplicates = checkRuns.map(({id, name, conclusion}) => ({id, name, conclusion}))
        const runs = runsWithPossibleDuplicates.filter((run, index, self) => self.findIndex(({id}) => id === run.id) === index)
        console.log(`Got the following check runs: ${JSON.stringify(runs)}`)
        const matchingRuns = runs.filter(({name}) => actions.includes(name))
        const failedRuns = matchingRuns.filter(({conclusion}) => conclusion !== 'success')
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
        await new Promise(r => setTimeout(r, 5000));
        now = new Date().getTime()
    }
    throw new Error(`Timed out waiting for ${pendingActions.join(', ')}`)
}
