import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import SearchPage from "./page";
import React from "react";

// next/link をモック化
vi.mock("next/link", () => {
  return {
    default: ({ children, href }: { children: React.ReactNode; href: string }) => {
      return <a href={href}>{children}</a>;
    },
  };
});

// next/image をモック化
vi.mock("next/image", () => {
  return {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    default: (props: any) => {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { fill, ...rest } = props;
      return <img {...rest} alt={props.alt} />;
    },
  };
});

describe("SearchPage", () => {
  beforeEach(() => {
    vi.resetAllMocks();
    // fetch のグローバルモック
    global.fetch = vi.fn();
  });

  it("ジャンルリストが初期表示時に取得されること", async () => {
    const mockGenres = [{ id: "1", name: "Rock" }];
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockGenres,
    } as Response);

    render(<SearchPage />);

    await waitFor(() => {
      expect(screen.getByText("Rock")).toBeInTheDocument();
    });
  });

  it("キーワードを入力して検索すると結果が表示されること", async () => {
    // 1回目の fetch (genres)
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => [],
    } as Response);

    // 2回目の fetch (search)
    const mockResults = [
      {
        id: "1",
        name: "Searched NFT",
        description: "Desc",
        price: "100",
        image_url: "/img.png",
        chain_id: 1,
      },
    ];
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => mockResults,
    } as Response);

    render(<SearchPage />);

    const input = screen.getByPlaceholderText("キーワードで検索...");
    fireEvent.change(input, { target: { value: "Searched" } });

    const searchButton = screen.getByRole("button", { name: "検索" });
    fireEvent.click(searchButton);

    await waitFor(() => {
      expect(screen.getByText("Searched NFT")).toBeInTheDocument();
    });
  });

  it("検索結果が空の場合にメッセージが表示されること", async () => {
    // 1回目の fetch (genres)
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => [],
    } as Response);

    // 2回目の fetch (search)
    vi.mocked(global.fetch).mockResolvedValueOnce({
      ok: true,
      json: async () => [],
    } as Response);

    render(<SearchPage />);

    const searchButton = screen.getByRole("button", { name: "検索" });
    fireEvent.click(searchButton);

    await waitFor(() => {
      expect(screen.getByText("検索結果が見つかりませんでした。")).toBeInTheDocument();
    });
  });
});
