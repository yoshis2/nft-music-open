import { describe, it, expect, vi } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import Header from "@/components/layouts/header";
import React from "react";

// next/link をモック化
vi.mock("next/link", () => {
  return {
    default: ({ children, href }: { children: React.ReactNode; href: string }) => {
      return <a href={href}>{children}</a>;
    },
  };
});

describe("Header Component", () => {
  it("レンダリングされること", () => {
    render(<Header />);
    // デスクトップメニューのテキストを確認
    expect(screen.getByText("スリーネクストNFTミュージック")).toBeInTheDocument();
  });

  it("モバイルメニューボタンをクリックするとメニューが開閉すること", () => {
    render(<Header />);
    const menuButton = screen.getByRole("button", { name: /toggle navigation menu/i });

    // 初期状態ではメニューは存在しない
    expect(screen.queryByText("スリーネクストが運営する NFT Musicの販売サイト")).not.toBeInTheDocument();

    // クリックして開く
    fireEvent.click(menuButton);
    expect(screen.getByText("スリーネクストが運営する NFT Musicの販売サイト")).toBeInTheDocument();

    // 再度クリックして閉じる
    fireEvent.click(menuButton);
    expect(screen.queryByText("スリーネクストが運営する NFT Musicの販売サイト")).not.toBeInTheDocument();
  });

  it("ウォレット接続ボタンが表示されていること", () => {
    render(<Header />);
    // Desktop と Mobile 両方にあるため getAllByRole を使用
    const connectButtons = screen.getAllByRole("button", { name: /ウォレット接続/i });
    expect(connectButtons.length).toBeGreaterThan(0);
  });
});
