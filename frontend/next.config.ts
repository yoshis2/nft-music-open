import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  productionBrowserSourceMaps: true,
  webpack: (config, { dev, isServer }) => {
    if (dev && !isServer) {
      config.devtool = "source-map";
    }
    return config;
  },
  turbopack: {},
  images: {
    remotePatterns: [
      {
        protocol: "http",
        hostname: "127.0.0.1",
        port: "8080",
        pathname: "/ipfs/**",
      },
      {
        protocol: "http",
        hostname: "ipfs",
        port: "8080",
        pathname: "/ipfs/**",
      },
      {
        protocol: "https",
        hostname: "ipfs.io",
        pathname: "/ipfs/**",
      },
    ],
  },
};

export default nextConfig;
