import { describe, it, expect, vi, beforeEach } from "vitest";
import { renderHook, act, waitFor } from "@testing-library/react";
import { useSearch } from "./useSearch";

describe("useSearch hook", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    global.fetch = vi.fn();
  });

  it("初期化時にジャンルを取得すること", async () => {
    const mockGenres = [{ id: "1", name: "Rock" }];
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockGenres,
    } as Response);

    const { result } = renderHook(() => useSearch());

    await waitFor(() => {
      expect(result.current.genres).toEqual(mockGenres);
    });
  });

  it("search関数を呼ぶと結果が取得されること", async () => {
    // genres fetch
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => [],
    } as Response);

    // search fetch
    const mockResults = [{ id: "1", name: "NFT" }];
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockResults,
    } as Response);

    const { result } = renderHook(() => useSearch());

    await act(async () => {
      await result.current.search();
    });

    expect(result.current.results).toEqual(mockResults);
    expect(result.current.hasSearched).toBe(true);
  });
});
