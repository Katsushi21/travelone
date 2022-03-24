FROM node:16-alpine3.12 AS build

RUN mkdir /app
WORKDIR /app

COPY . .

RUN yarn install && yarn cache clean

CMD [ "yarn", "dev" ]
