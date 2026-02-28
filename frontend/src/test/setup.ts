import "@testing-library/jest-dom";
import "vitest-canvas-mock";
import { cleanup } from "@testing-library/react";
import { afterEach } from "vitest";

// 各テストの後に DOM をクリーンアップ
afterEach(() => {
  cleanup();
});
