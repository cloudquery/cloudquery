# Getting Started with custom Plugin UI

### `npm install`

This will install all the dependencies required for the Plugin UI to run.

Make sure to copy `.env.example` to `.env` inside root folder before running locally:

- `REACT_APP_PLUGIN_TEAM` is the team name of the plugin
- `REACT_APP_PLUGIN_KIND` is the kind of the plugin (source or destination)
- `REACT_APP_PLUGIN_NAME` is the name of the plugin
- `REACT_APP_PLUGIN_VERSION` is the version of the plugin (e.g. `v1.0.0`)

In addition, you can copy `.env.example.json` to `.env.json` inside `src` folder before running locally:

- `authToken` and `teamName` is required if you need to use CloudQuery API outside of CloudQuery Cloud App (https://cloud.cloudquery.io). You can get the token by navigating to https://cloud.cloudquery.io and inspecting any fetch request: you can extract the token from the `Cookie` header with value that start with`__session=`.
- `initialValues` is required if you want to see how your plugin behaves with initial values

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3001](http://localhost:3001) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `npm run build`

Copies `.env.example.json` to `.env.json` inside `src` folder.

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

### `npm run dev:test:e2e`

This command enables the ability to locally write and run e2e tests.

This mimics how e2e tests will run in the CI environment. The tests are conducted against the production cloud app with the preview URL from a built PR.

The `.env` file should be copied from `.env.example` and populated with the needed environment variables for the e2e tests to run properly. Specifically, the following variables are required for the e2e tests to run:

- CQ_CI_PLAYWRIGHT_TEST_USER_EMAIL: this can be any valid CloudQuery username
- CQ_CI_PLAYWRIGHT_TEST_USER_PASSWORD: this can be any valid CloudQuery password
- CQ_CI_PLAYWRIGHT_PREVIEW_LINK: this will be generated and commented on a successfully built PR. It will be in a format like: `https://cloud.cloudquery.io/teams/cloudquery-test/destinations/create?plugin-cloud-ui=cloudquery|destination|mysql|https://plugin-destination-mysql-cloud-ui-19223.vercel.app`
