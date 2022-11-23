# CloudQuery Website & Blog

## Running locally

We recommend using `nvm` to manage npm versions. See [nvm installation instructions](https://github.com/nvm-sh/nvm#installing-and-updating).

With `nvm` installed, run the following command from within the `website` directory:

```shell
nvm use
```

(and potentially `nvm install`) to set the correct version of `npm`.

Next, install dependencies via `yarn`:

```shell
yarn install
```

And now run the server in development mode:

```shell
yarn dev
```

This should start a server running on `localhost:3000`.