FROM node:18-alpine3.14 as dev

ENV ROOT=/src/client
WORKDIR ${ROOT}

RUN apk update && \
  apk add --no-cache libc6-compat=1.2.2-r3 && \
  yarn install && \
  yarn cache clean
COPY . ${ROOT}

CMD ["yarn", "dev"]