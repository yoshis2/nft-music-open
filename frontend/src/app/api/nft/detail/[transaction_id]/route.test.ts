import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET } from "./route";
import { NextRequest } from "next/server";

describe("api/nft/detail/[transaction_id]", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("GET: NFTの詳細を取得できること", async () => {
    // コード内の変数名 transacctionId に合わせる
    const params = { transacctionId: "tx123" };
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "1" }),
    } as Response);

    const request = new NextRequest("http://localhost/api/nft/detail/tx123");
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const response = await GET(request, { params } as any);
    expect(response.status).toBe(200);
  });
});
