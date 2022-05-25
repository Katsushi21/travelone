FROM node:18-alpine3.14 AS dev
WORKDIR /src/client

COPY package.json .
RUN apk update && \
  apk add --no-cache libc6-compat=1.2.2-r3 && \
  yarn install && \
  yarn cache clean
COPY . .

CMD ["yarn", "dev"]