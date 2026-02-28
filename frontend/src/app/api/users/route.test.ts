import { describe, it, expect, vi, beforeEach } from "vitest";
import { GET, POST, PUT } from "./route";
import { NextRequest } from "next/server";

describe("api/users", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("GET: ユーザー情報を取得できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "1", wallet: "0x123" }),
    } as Response);

    const request = new NextRequest("http://localhost/api/users?wallet=0x123");
    const response = await GET(request);
    expect(response.status).toBe(200);
  });

  it("POST: ユーザーを作成できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "1" }),
    } as Response);

    const request = new NextRequest("http://localhost/api/users", {
      method: "POST",
      body: JSON.stringify({ wallet: "0x123" }),
    });

    const response = await POST(request);
    expect(response.status).toBe(200);
  });

  it("PUT: ユーザーを更新できること", async () => {
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ id: "1" }),
    } as Response);

    const request = new NextRequest("http://localhost/api/users", {
      method: "PUT",
      body: JSON.stringify({ id: "1", name: "User" }),
    });

    const response = await PUT(request);
    expect(response.status).toBe(200);
  });
});
