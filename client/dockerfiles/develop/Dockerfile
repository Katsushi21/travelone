FROM node:18-alpine3.14 as dev

ENV ROOT=/traveling_alone
WORKDIR ${ROOT}
COPY . .

RUN apk update && \
  yarn install && \
  yarn cache clean

CMD ["yarn", "dev"]