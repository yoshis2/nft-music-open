import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET, POST } from "./route";
import { NextRequest } from "next/server";

describe("api/nft", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("GET: NFT一覧を取得できること", async () => {
    const mockData = [{ id: "1", name: "NFT" }];
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockData,
    } as Response);

    const response = await GET();
    const data = await response.json();

    expect(response.status).toBe(200);
    expect(data).toEqual(mockData);
  });

  it("POST: NFTを作成できること", async () => {
    const mockReq = { name: "New NFT" };
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "2", ...mockReq }),
    } as Response);

    const request = new NextRequest("http://localhost/api/nft", {
      method: "POST",
      body: JSON.stringify(mockReq),
    });

    const response = await POST(request);
    expect(response.status).toBe(200);
  });
});
