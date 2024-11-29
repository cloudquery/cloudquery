FROM node:22-slim as builder

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

RUN npm run build

FROM node:22-slim AS final

WORKDIR /app

COPY --from=builder ./app/dist ./dist

COPY package*.json ./

RUN npm ci --omit=dev

EXPOSE 7777

ENTRYPOINT ["node", "dist/main.js"]

CMD [ "serve", "--address", "[::]:7777", "--log-format", "json", "--log-level", "info" ]