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

## Plugin tables

In case your plugin is a source plugin and you want to use the list of tables in your plugin UI, then make sure to build the plugin first to get a generated list of tables. Inside the plugin root directory run `go build`. This will generate a plugin file and then you should navigate back to the `cloud-config-ui` directory inside your plugin and run `npm start` or `npm run build`, they both will generate a `__tables.json` file in the `src/data` folder. After that you can import the list of tables inside the application like this:

```ts
import pluginTables from 'data/__tables.json';
```

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3001](http://localhost:3001) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

### Tips

- When using Material UI select, always set `MenuProps` with `autoFocus` value `false` and `disableAutoFocus` - `true` to avoid content jump when select is becomes focused.