import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET, PUT, DELETE } from "./route";
import { NextRequest } from "next/server";

describe("api/master/genres/[id]", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  const params = { id: "1" };

  it("GET: 特定のジャンルを取得できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "1", name: "Rock" }),
    } as Response);

    const request = new NextRequest("http://localhost/api/master/genres/1");
    const response = await GET(request, { params: Promise.resolve(params) });
    expect(response.status).toBe(200);
  });

  it("PUT: ジャンルを更新できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "1", name: "Updated Rock" }),
    } as Response);

    const request = new NextRequest("http://localhost/api/master/genres/1", {
      method: "PUT",
      body: JSON.stringify({ name: "Updated Rock" }),
    });
    const response = await PUT(request, { params: Promise.resolve(params) });
    expect(response.status).toBe(200);
  });

  it("DELETE: ジャンルを削除できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ success: true }),
    } as Response);

    const request = new NextRequest("http://localhost/api/master/genres/1", {
      method: "DELETE",
    });
    const response = await DELETE(request, { params: Promise.resolve(params) });
    expect(response.status).toBe(200);
  });
});
