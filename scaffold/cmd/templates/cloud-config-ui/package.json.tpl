{
  "name": "cloud-config-ui",
  "version": "0.0.1",
  "private": true,
  "homepage": "./",
  "dependencies": {
    "@cloudquery/cloud-ui": "^0.1.13",
    "@cloudquery/plugin-config-ui-connector": "^0.2.12",
    "@cloudquery/plugin-config-ui-lib": "^0.0.56",
    "@emotion/react": "^11.11.4",
    "@emotion/styled": "^11.11.5",
    "@mui/icons-material": "^5.15.20",
    "@mui/lab": "^5.0.0-alpha.170",
    "@mui/material": "^5.15.20",
    "@mui/system": "^5.15.20",
    "@mui/x-date-pickers": "^7.6.2",
    "@mui/x-tree-view": "^7.6.2",
    "humanize-string": "^3.0.0",
    "react": "^18.3.1",
    "react-dom": "^18.3.1",
    "react-hook-form": "^7.52.2",
    "react-scripts": "^5.0.1",
    "typescript": "^4.9.5",
    "yup": "^1.4.0"
  },
  "devDependencies": {
    "@babel/plugin-proposal-private-property-in-object": "^7.21.11",
    "@playwright/test": "^1.46.0",
    "@types/node": "^22.2.0",
    "@types/react": "^18.3.3",
    "@types/react-dom": "^18.3.0",
    "eslint": "^8.57.0",
    "eslint-config-prettier": "^9.1.0",
    "eslint-import-resolver-typescript": "^3.6.1",
    "eslint-plugin-custom-rules": "file:./eslint-plugin-custom-rules",
    "eslint-plugin-import": "^2.29.1",
    "eslint-plugin-jsx-a11y": "^6.9.0",
    "eslint-plugin-prettier": "^5.2.1",
    "eslint-plugin-react": "^7.35.0",
    "eslint-plugin-react-hooks": "^4.6.2",
    "eslint-plugin-sort-destructure-keys": "^2.0.0",
    "eslint-plugin-unicorn": "^55.0.0",
    "http-server": "^14.1.1",
    "prettier": "^3.3.3",
    "yaml": "^2.5.0"
  },
  "scripts": {
    "start": "node scripts/initialize.js && PORT=3001 react-scripts start",
    "build": "node scripts/initialize.js && react-scripts build",
    "lint": "eslint src --ext .ts,.tsx --max-warnings 0",
    "lint:fix": "eslint src --ext .ts,.tsx --max-warnings 0 --fix",
    "test:e2e": "playwright test"
  },
  "eslintConfig": {
    "extends": [
      "react-app"
    ]
  },
  "browserslist": {
    "production": [
      ">0.2%",
      "not dead",
      "not op_mini all"
    ],
    "development": [
      "last 1 chrome version",
      "last 1 firefox version",
      "last 1 safari version"
    ]
  }
}
