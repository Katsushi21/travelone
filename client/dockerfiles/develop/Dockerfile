FROM node:18-alpine3.14 AS dev
WORKDIR /src/client

RUN apk update && \
  apk add --no-cache libc6-compat=1.2.2-r3

COPY package.json .
RUN yarn install && \
  yarn cache clean

COPY . .

CMD ["yarn", "dev"]