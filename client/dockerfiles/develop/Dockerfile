FROM node:18-alpine3.14 as dev

ENV ROOT=/src
WORKDIR ${ROOT}
COPY ./src ./
COPY ./package.json ./

# hadolint ignore=DL3018
RUN apk update && \
  yarn install && \
  yarn cache clean

CMD ["yarn", "dev"]