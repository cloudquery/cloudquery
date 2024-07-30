# Getting Started with custom Plugin UI

### `npm install`

This will install all the dependencies required for the Plugin UI to run.

Make sure to copy `.env.example.json` to `.env.json` inside `src` folder before running locally:

- `authToken` and `teamName` is required if you need to use CloudQuery API. You can get the token by navigating to https://cloud.cloudquery.io and inspecting any fetch request: you can extract the token from the `Authorization` header.
- `initialValues` is required if you want to see how your plugin behaves with initial values

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
