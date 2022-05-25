# 依存パッケージのインストール
FROM node:18-alpine3.14 AS deps

ENV ROOT=/src/client
WORKDIR ${ROOT}

COPY package.json yarn.lock ./
RUN apk update && \
  apk add --no-cache libc6-compat=1.2.2-r3 && \
  yarn install --frozen-lockfile --ignore-optional && \
  yarn cache clean

# ビルド
FROM node:18-alpine3.14 AS builder

ENV ROOT=/src/client
WORKDIR ${ROOT}
COPY --from=deps /app/node_modules ./node_modules
COPY . ${ROOT}

RUN yarn build

# 実行
FROM node:18-alpine3.14 AS runner
ENV ROOT=/src/client
WORKDIR ${ROOT}

ENV NODE_ENV production

COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

RUN addgroup --system --gid 1001 nodejs && \
  adduser --system --uid 1001 nextjs
USER nextjs

CMD ["node", "server.js"]