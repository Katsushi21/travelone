/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    domains: ['travel-client.com', 'api.lorem.space'],
  },
  experimental: {
    appDir: true,
  },
};

module.exports = nextConfig;
