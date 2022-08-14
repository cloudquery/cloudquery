


## Automation:
- All commit messages will be validated by GitHub Action
    - (This action should be required for merge)
- Release Notes will be triggered manually during 


## Commit Messages


| Title      | Message | Action |
| ----------- | ----------- |----------- |
| `feat: <Message>`      |  `<String>`       | minor release|
| `enhancement: <Message>`      |  `<String>`       | patch release|
| `docs: <Message>`      |  `<String>`       | patch release|
| `ci: <Message>`      |  `<String>`       | patch release|
| `chore: <Message>`      |  `<String>`       | patch release|
| `refactor: <Message>`      |  `<String>`       | patch release|
| `test: <Message>`      |  `<String>`       | patch release|
| `fix: <Message>`      |  `<String>`      | patch release|

    

## Running locally:

You can run the changelog + release locally if you want to see what it will do:

1. install dependencies:

    `npm install`
2. Run tool:

    `GITHUB_TOKEN=<PERSONAL_ACCESS_TOKEN> npx semantic-release -d`



            
            
            
            
            
            