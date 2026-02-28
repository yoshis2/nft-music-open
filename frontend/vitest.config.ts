import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import path from "path";

export default defineConfig(({ mode }) => {
  const isDev = mode === "development";

  return {
    build: {
      chunkSizeWarningLimit: 1000,
      // 開発モード（npm run dev）の時だけ true、本番（npm run build）では false になる
      sourcemap: isDev,
    },
    css: {
      // 開発時のCSSソースマップも有効化しておくと便利です
      devSourcemap: isDev,
      preprocessorOptions: {
        scss: {
          quietDeps: true,
        },
      },
    },
    plugins: [react()],
    test: {
      environment: "jsdom",
      globals: true,
      setupFiles: ["./src/test/setup.ts"],
      alias: {
        "@": path.resolve(__dirname, "./src"),
        "@Contract": path.resolve(__dirname, "./artifacts/contracts"),
      },
      exclude: ["**/node_modules/**", "**/e2e/**"],
      coverage: {
        provider: "v8",
        reporter: ["text", "json", "html"],
        exclude: ["node_modules/", "src/test/setup.ts"],
      },
    },
  };
});
