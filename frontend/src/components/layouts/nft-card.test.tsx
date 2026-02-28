import { describe, it, expect, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { NftCard } from "./nft-card";
import { RenderableMarketItem } from "@/types/types";
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

const mockNft: RenderableMarketItem = {
  id: "1",
  user_id: "user1",
  chain_id: 1,
  tokenId: 1,
  nonce: 0,
  name: "Test NFT",
  description: "Test Description",
  file_type: "image",
  image_url: "/ipfs/image.png",
  audio_url: "",
  video_url: "",
  contract_address: "0x123",
  token_url: "http://example.com",
  genre_id: "1",
  genre_name: "Rock",
  creator_address: "0x456",
  created_at: "2024-01-01",
  updated_at: "2024-01-01",
  from: "0x789",
  to: "0xabc",
  price: "0.1",
  insentive: 0,
};

describe("NftCard Component", () => {
  it("NFTの情報が正しく表示されること", () => {
    render(<NftCard nft={mockNft} />);

    expect(screen.getByText("Test NFT")).toBeInTheDocument();
    expect(screen.getByText("Test Description")).toBeInTheDocument();
    expect(screen.getByText("0.1 ETH")).toBeInTheDocument();
  });

  it("画像URLが正しく設定されること", () => {
    render(<NftCard nft={mockNft} />);
    const image = screen.getByAltText("Test NFT");
    expect(image).toHaveAttribute("src", "https://ipfs.io/ipfs/image.png");
  });

  it("ローカルネットワークの場合にローカルIPFSゲートウェイを使用すること", () => {
    const localNft = { ...mockNft, chain_id: 1337 };
    render(<NftCard nft={localNft} />);
    const image = screen.getByAltText("Test NFT");
    expect(image).toHaveAttribute("src", "http://127.0.0.1:8080/ipfs/image.png");
  });

  it("音声ファイルの場合にaudioタグが表示されること", () => {
    const audioNft = { ...mockNft, file_type: "audio", audio_url: "/ipfs/audio.mp3" };
    render(<NftCard nft={audioNft} />);
    // audioタグが存在することを確認
    const audioElement = document.querySelector("audio");
    expect(audioElement).toBeInTheDocument();
    expect(audioElement).toHaveAttribute("src", "https://ipfs.io/ipfs/audio.mp3");
  });

  it("動画ファイルの場合にvideoタグが表示されること", () => {
    const videoNft = { ...mockNft, file_type: "video", video_url: "/ipfs/video.mp4" };
    render(<NftCard nft={videoNft} />);
    // videoタグが存在することを確認
    const videoElement = document.querySelector("video");
    expect(videoElement).toBeInTheDocument();
    expect(videoElement).toHaveAttribute("src", "https://ipfs.io/ipfs/video.mp4");
  });
});
