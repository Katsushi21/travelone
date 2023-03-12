FROM node:18-alpine3.14 AS deps
WORKDIR /src/client

RUN apk update && \
  apk add --no-cache libc6-compat=1.2.2-r3

COPY package.json .
RUN yarn install --frozen-lockfile --ignore-optional && \
  yarn cache clean

FROM deps AS builder
COPY --from=deps /src/client/node_modules ./node_modules
COPY . .

RUN yarn build

FROM builder AS runner
ENV NODE_ENV production

COPY --from=builder /src/client/next.config.js .
COPY --from=builder /src/client/public ./public
COPY --from=builder /src/client/package.json ./package.json
COPY --from=builder --chown=nextjs:nextjs /src/client/.next/standalone .
COPY --from=builder --chown=nextjs:nextjs /src/client/.next/static ./.next/static

RUN addgroup --system --gid 1001 nextjs && \
  adduser --system --uid 1001 nextjs
USER nextjs:nextjs

CMD ["node", "server.js"]