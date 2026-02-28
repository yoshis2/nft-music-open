"use client";

import type { NextPage } from "next";
import Image from "next/image";
import { useEffect, useState } from "react";

import { ethers } from "ethers";

import { type RenderableMarketItem } from "@/types/types";

const NftListing: NextPage = () => {
  const [itemList, setItemList] = useState<RenderableMarketItem[]>([]);
  const [error, setError] = useState<Error | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      if (!window.ethereum) {
        setError(
          new Error("ウォレットが接続されていません。MetaMaskなどのウォレットをインストールして接続してください。"),
        );
        setIsLoading(false);
        return;
      }

      try {
        const provider = new ethers.BrowserProvider(window.ethereum);
        const signer = await provider.getSigner();
        const wallet = await signer.getAddress();

        if (!wallet) {
          setError(new Error("ウォレットアドレスが取得できませんでした。"));
          setIsLoading(false);
          return;
        }

        const response = await fetch(`/api/nft/${wallet}`, {
          method: "GET",
        });

        if (!response.ok) {
          throw new Error(`APIエラー: ${response.status} ${response.statusText}`);
        }

        const items = await response.json();
        setItemList(items || []);
      } catch (err) {
        console.error("データの取得に失敗しました:", err);
        setError(err instanceof Error ? err : new Error(String(err)));
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  const renderList = itemList.map((item) => {
    let NFTImage = "https://ipfs.io" + item.token_url;
    if (item.chain_id == 1337) {
      NFTImage = "http://ipfs:8080" + item.image_url;
    }
    return (
      <div
        key={item.id}
        className="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 card-item flex flex-col"
      >
        <div className="p-4">
          <h3 className="text-xl font-bold text-gray-900 truncate dark:text-white">{item.name}</h3>
        </div>
        <div className="relative w-full h-56">
          <Image src={NFTImage} alt={item.name} fill sizes="(max-width: 640px) 100vw, 50vw" className="object-cover" />
        </div>
        <div className="p-4 flex-grow flex flex-col">
          <p className="text-sm text-gray-600 mt-1 truncate dark:text-gray-400">{item.description}</p>
          <div className="mt-4">
            <p className="text-xs text-gray-500">Price</p>
            <p className="text-lg font-semibold text-gray-800 dark:text-gray-200">{item.price} ETH</p>
          </div>
          <div className="mt-2">
            <p className="text-xs text-gray-500">Insentive</p>
            <p className="text-lg font-semibold text-gray-800 dark:text-gray-200">{item.insentive} %</p>
          </div>
        </div>
      </div>
    );
  });

  return (
    <main className="main-container">
      <h1 className="heading1">NFT一覧</h1>
      {isLoading ? (
        <p>読み込み中...</p>
      ) : error ? (
        <p className="error-text">{error.message}</p>
      ) : itemList.length > 0 ? (
        <div className="w-full grid grid-cols-2 sm:grid-cols-6 lg:grid-cols-6 2xl:grid-cols-3 gap-1">{renderList}</div>
      ) : (
        <p>表示するNFTがありません。</p>
      )}
    </main>
  );
};

export default NftListing;
