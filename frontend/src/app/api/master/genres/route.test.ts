import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET, POST } from "./route";
import { NextRequest } from "next/server";

describe("api/master/genres", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  describe("GET", () => {
    it("ジャンル一覧を取得できること", async () => {
      const mockData = [{ id: "1", name: "Rock" }];
      vi.mocked(global.fetch).mockResolvedValueOnce({
        ok: true,
        json: async () => mockData,
      } as Response);

      const response = await GET();
      const data = await response.json();

      expect(response.status).toBe(200);
      expect(data).toEqual(mockData);
      expect(global.fetch).toHaveBeenCalledWith("http://backend:1323/api/v1/genres", expect.any(Object));
    });

    it("エラー時に500を返すこと", async () => {
      vi.mocked(global.fetch).mockRejectedValueOnce(new Error("Fetch error"));

      const response = await GET();
      expect(response.status).toBe(500);
    });
  });

  describe("POST", () => {
    it("ジャンルを作成できること", async () => {
      const mockReq = { name: "Jazz" };
      const mockRes = { id: "2", ...mockReq };
      vi.mocked(global.fetch).mockResolvedValueOnce({
        ok: true,
        json: async () => mockRes,
      } as Response);

      const request = new NextRequest("http://localhost/api/master/genres", {
        method: "POST",
        body: JSON.stringify(mockReq),
      });

      const response = await POST(request);
      const data = await response.json();

      expect(response.status).toBe(200);
      expect(data).toEqual(mockRes);
    });
  });
});
