FROM node:24-slim AS builder

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN pnpm install --frozen-lockfile

COPY . .

RUN pnpm run build

FROM node:24-slim AS final

WORKDIR /app

RUN corepack enable

COPY --from=builder ./app/dist ./dist

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN pnpm install --frozen-lockfile --prod

EXPOSE 7777

ENTRYPOINT ["node", "dist/main.js"]

CMD [ "serve", "--address", "[::]:7777", "--log-format", "json", "--log-level", "info" ]
