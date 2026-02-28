// frontend/eslint.config.mjs
import { dirname } from "path";
import { fileURLToPath } from "url";
import js from "@eslint/js";
import tseslint from "typescript-eslint";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default tseslint.config(
  {
    // 無視するディレクトリ
    ignores: [".next/**", "node_modules/**", "coverage/**"],
  },
  // 1. 基本的なJavaScriptの推奨ルール
  js.configs.recommended,

  // 2. TypeScriptの推奨ルール（パーサー設定も自動で含まれます）
  ...tseslint.configs.recommended,

  // 3. プロジェクト固有の設定
  {
    files: ["**/*.ts", "**/*.tsx"],
    languageOptions: {
      parserOptions: {
        project: "./tsconfig.json",
        tsconfigRootDir: __dirname,
      },
    },
    rules: {
      // 開発中は警告レベルに留める（エラーで止まらないようにする）
      "@typescript-eslint/no-explicit-any": "warn",
      "@typescript-eslint/no-unused-vars": "warn",

      // Next.js (React 17+) では不要なルールをOFF
      "react/react-in-jsx-scope": "off",
    },
  },
);
