import { describe, it, expect, vi, beforeEach } from "vitest";
import { POST } from "./route";
import { NextRequest } from "next/server";

describe("api/ipfs/file", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("POST: ファイルをIPFSにアップロードできること", async () => {
    const mockRes = { hash: "Qm..." };
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockRes,
    } as Response);

    const formData = new FormData();
    formData.append("file", new Blob(["test"], { type: "text/plain" }), "test.txt");

    const request = new NextRequest("http://localhost/api/ipfs/file", {
      method: "POST",
    });
    // request.formData() を直接モックする
    request.formData = vi.fn().mockResolvedValue(formData);

    const response = await POST(request);
    const data = await response.json();

    expect(response.status).toBe(200);
    expect(data).toEqual(mockRes);
  });
});
