import { describe, it, expect, vi, beforeEach } from "vitest";
import { POST } from "./route";
import { NextRequest } from "next/server";

describe("api/ipfs/meta", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("POST: メタデータをIPFSにアップロードできること", async () => {
    const mockReq = { name: "NFT Name", description: "Desc" };
    const mockRes = { hash: "QmMeta..." };
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockRes,
    } as Response);

    const request = new NextRequest("http://localhost/api/ipfs/meta", {
      method: "POST",
      body: JSON.stringify(mockReq),
    });

    const response = await POST(request);
    const data = await response.json();

    expect(response.status).toBe(200);
    expect(data).toEqual(mockRes);
  });
});
