const isProd = process.env.NODE_ENV === 'development'
module.exports = {
  reactStrictMode: true,
  env: {
    isProd: isProd,
  }
}
