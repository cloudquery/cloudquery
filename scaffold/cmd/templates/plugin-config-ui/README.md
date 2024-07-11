# Getting Started with custom Plugin UI

### `npm install`

This will install all the dependencies required for the Plugin UI to run.

## Plugin tables

In case your plugin is a source plugin and you want to use the list of tables in your plugin UI, then make sure to build the plugin first to get generated list of tables. Inside the plugin root directory run `go build`. This will generate a plugin file and then you should run `./{generated-plugin-file-name} doc --format=json docs`. The generated file with table should be located inside `docs/__tables.json`. Then navigate back to the `plugin-config-ui` directory inside your plugin and run `npm install`. After that you can import the list of tables inside the application like this:

```ts
import pluginTables from '@cloudquery-plugin/tables';
```

If you change your plugin configuration that should update the list of plugins, you only need to regenerate `tables.json`, the frontend app will automatically detect changes there.

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

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!
