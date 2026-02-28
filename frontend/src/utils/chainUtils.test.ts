import { describe, it, expect } from "vitest";
import { getChainInfo } from "./chainUtils";

describe("chainUtils", () => {
  describe("getChainInfo", () => {
    it("Ethereum (1) の情報を返すこと", () => {
      const info = getChainInfo(1);
      expect(info.name).toBe("Ethereum");
      expect(info.explorerUrl).toBe("https://etherscan.io");
    });

    it("Polygon (137) の情報を返すこと", () => {
      const info = getChainInfo(137);
      expect(info.name).toBe("Polygon");
      expect(info.explorerUrl).toBe("https://polygonscan.com");
    });

    it("Sepolia (11155111) の情報を返すこと", () => {
      const info = getChainInfo(11155111);
      expect(info.name).toBe("Sepolia");
      expect(info.explorerUrl).toBe("https://sepolia.etherscan.io");
    });

    it("Local Network (1337) の情報を返すこと", () => {
      const info = getChainInfo(1337);
      expect(info.name).toBe("Local Network");
      expect(info.explorerUrl).toBe("");
    });

    it("未知のチェーンIDに対して適切に返すこと", () => {
      const info = getChainInfo(999);
      expect(info.name).toBe("Unknown Chain (999)");
      expect(info.explorerUrl).toBe("");
    });
  });
});
