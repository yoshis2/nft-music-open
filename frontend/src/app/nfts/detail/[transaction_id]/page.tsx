import type { NextPage } from "next";
import Image from "next/image";
import Link from "next/link";

import Footer from "@/components/layouts/footer";
import Header from "@/components/layouts/header";
import { RenderableMarketItem } from "@/types/types";
import { getChainInfo } from "@/utils/chainUtils";

type Props = {
  params: {
    transaction_id: string;
  };
};

async function getNftDetail(transactionId: string): Promise<RenderableMarketItem | null> {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/nfts/detail/${transactionId}`, {
      cache: "no-store",
    });

    if (!res.ok) {
      console.error(`Failed to fetch NFT detail. Status: ${res.status}`);
      return null;
    }
    return res.json();
  } catch (error) {
    console.error("An error occurred while fetching NFT detail:", error);
    return null;
  }
}

const NftDetailPage: NextPage<Props> = async ({ params }) => {
  const nft = await getNftDetail(params.transaction_id);
  console.log("NFTの内容 : ");
  console.log(nft);

  if (!nft) {
    return (
      <>
        <Header />
        <main className="main-container">
          <h1 className="heading1">NFTが見つかりません</h1>
          <p className="text-center text-gray-500">指定されたNFTは存在しないか、取得できませんでした。</p>
          <Link href="/market" className="submit-button mt-8">
            マーケットに戻る
          </Link>
        </main>
        <Footer />
      </>
    );
  }

  // IPFS URLとブロックチェーン情報の決定
  let imageUrl = "https://ipfs.io" + nft.image_url;
  let mediaUrl = nft.video_url ? "https://ipfs.io" + nft.video_url : "";
  if (nft.chain_id == 1337) {
    imageUrl = "http://127.0.0.1:8080" + nft.image_url;
    if (nft.video_url) {
      mediaUrl = "http://127.0.0.1:8080" + nft.video_url;
    }
  }

  const chainInfo = getChainInfo(nft.chain_id);

  return (
    <>
      <Header />
      <main className="main-container">
        <div className="grid w-full md:grid-cols-5 gap-8">
          {/* メディアセクション (左側) */}
          <div className="md:col-span-3">
            {nft.file_type === "video" && mediaUrl ? (
              <video
                controls
                src={mediaUrl}
                className="w-full rounded-lg shadow-lg aspect-video bg-black"
                poster={imageUrl}
              ></video>
            ) : (
              <div className="relative w-full aspect-square rounded-lg overflow-hidden shadow-lg bg-gray-200">
                <Image
                  src={imageUrl}
                  alt={nft.name}
                  fill
                  sizes="(max-width: 768px) 100vw, 60vw"
                  className="object-cover"
                  priority
                />
                {nft.file_type === "audio" && mediaUrl && (
                  <div className="absolute bottom-0 left-0 right-0 p-4 bg-black bg-opacity-50 backdrop-blur-sm">
                    <audio controls src={mediaUrl} className="w-full"></audio>
                  </div>
                )}
              </div>
            )}
          </div>

          {/* 情報セクション (右側) */}
          <div className="md:col-span-2 flex flex-col">
            <div className="mb-4">
              <p className="text-sm text-gray-500">{nft.genre_name || "Uncategorized"}</p>
              <h1 className="text-4xl font-bold text-gray-900">{nft.name}</h1>
            </div>
            <div className="mb-4">
              <p className="font-semibold text-gray-700">Creator</p>
              <a
                href={
                  chainInfo.explorerUrl && nft.creator_address
                    ? `${chainInfo.explorerUrl}/address/${nft.creator_address}`
                    : "#"
                }
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-600 hover:underline truncate block"
              >
                {nft.creator_address || "N/A"}
              </a>
            </div>

            <div className="bg-gray-50 p-6 rounded-lg border border-gray-200 flex-grow flex flex-col">
              <div className="mb-4">
                <p className="text-sm text-gray-500">価格</p>
                <p className="text-3xl font-bold text-gray-800">{nft.price} ETH</p>
              </div>
              <div className="mb-4">
                <p className="text-sm text-gray-500">クリエイターインセンティブ</p>
                <p className="text-xl font-semibold text-gray-800">{nft.insentive} %</p>
              </div>
              <button className="w-full bg-blue-600 text-white font-bold py-3 px-8 rounded-lg hover:bg-blue-700 transition-colors mt-auto">
                購入する
              </button>
            </div>
          </div>
        </div>

        {/* 説明 & 詳細セクション */}
        <div className="mt-8 grid w-full md:grid-cols-5 gap-8">
          <div className="md:col-span-3">
            <div className="border-b border-gray-200 mb-4">
              <h2 className="text-2xl font-semibold pb-2">説明</h2>
            </div>
            <p className="text-gray-700 leading-relaxed whitespace-pre-wrap">{nft.description}</p>
          </div>
          <div className="md:col-span-2">
            <div className="border-b border-gray-200 mb-4">
              <h2 className="text-2xl font-semibold pb-2">詳細</h2>
            </div>
            <dl className="space-y-3 text-sm">
              <div className="flex justify-between">
                <dt className="text-gray-500">コントラクトアドレス</dt>
                <dd className="text-right">
                  {nft.contract_address ? (
                    <a
                      href={chainInfo.explorerUrl ? `${chainInfo.explorerUrl}/address/${nft.contract_address}` : "#"}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-blue-600 hover:underline truncate"
                    >
                      {nft.contract_address.substring(0, 6)}...
                      {nft.contract_address.substring(nft.contract_address.length - 4)}
                    </a>
                  ) : (
                    <span className="text-gray-900">N/A</span>
                  )}
                </dd>
              </div>
              <div className="flex justify-between">
                <dt className="text-gray-500">トークンID</dt>
                <dd className="text-gray-900 truncate">{nft.tokenId ?? "N/A"}</dd>
              </div>
              <div className="flex justify-between">
                <dt className="text-gray-500">トークン規格</dt>
                <dd className="text-gray-900">ERC-721</dd>
              </div>
              <div className="flex justify-between">
                <dt className="text-gray-500">ブロックチェーン</dt>
                <dd className="text-gray-900">{chainInfo.name}</dd>
              </div>
            </dl>
          </div>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default NftDetailPage;
