# Getting Started with custom Plugin UI

### `npm install`

This will install all the dependencies required for the Plugin UI to run.

Make sure to copy `.env.example.json` to `.env.json` inside `src` folder before running locally:
* `authToken` and `teamName` is required if you need to use CloudQuery API. You can get the token by navigating to https://cloud.cloudquery.io and inspecting any fetch request: you can extract the token from the `Authorization` header.
* `initialValues` is required if you want to see how your plugin behaves with initial values

## Plugin tables

In case your plugin is a source plugin and you want to use the list of tables in your plugin UI, then make sure to build the plugin first to get a generated list of tables. Inside the plugin root directory run `go build`. This will generate a plugin file. Then navigate back to the `cloud-config-ui` directory inside your plugin and run `npm start` or `npm run build`, they both will generate a `__tables.json` file and move it to the `src/data` folder. After that you can import the list of tables inside the application like this:

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

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the `data/__tables.json` file.

Copies the `.env.example.json` to `.env.json`.

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

### Tips

* When using Material UI select, always set `MenuProps` with `autoFocus` value `false` and `disableAutoFocus` - `true` to avoid content jump when select is becomes focused.