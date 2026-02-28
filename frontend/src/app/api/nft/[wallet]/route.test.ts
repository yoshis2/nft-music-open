import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET } from "./route";
import { NextRequest } from "next/server";

describe("api/nft/[wallet]", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("GET: ウォレット所有のNFTを取得できること", async () => {
    const params = { wallet: "0x123" };
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => [{ id: "1" }],
    } as Response);

    const request = new NextRequest("http://localhost/api/nft/0x123");
    const response = await GET(request, { params });
    expect(response.status).toBe(200);
  });
});
