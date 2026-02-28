import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET } from "./route";
import { NextRequest } from "next/server";

describe("api/nft/search", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("GET: NFTを検索できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => [{ id: "1" }],
    } as Response);

    const request = new NextRequest("http://localhost/api/nft/search?params=q=test");
    const response = await GET(request);
    expect(response.status).toBe(200);
  });
});
