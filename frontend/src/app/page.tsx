import type { NextPage } from "next";
import Link from "next/link";

import Header from "@/components/layouts/header";
import Footer from "@/components/layouts/footer";
import { NftCard } from "@/components/layouts/nft-card";
import { RenderableMarketItem } from "@/types/types";

// APIからNFTデータを取得する関数
async function getFeaturedNfts(): Promise<RenderableMarketItem[]> {
  try {
    // 例: 新しいものから4件取得
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/nfts?limit=4`, {
      cache: "no-store", // ページがリクエストされるたびに最新の情報を取得
    });
    if (!res.ok) {
      // エラーレスポンスの場合は空の配列を返す
      console.error("Failed to fetch NFTs. Status:", res.status);
      return [];
    }
    return res.json();
  } catch (error) {
    // ネットワークエラーなどが発生した場合
    console.error("An error occurred while fetching NFTs:", error);
    return [];
  }
}

/**
 * トップページ
 *
 * Hero SectionとFeatured NFTs Sectionを含みます。
 *
 * Hero Sectionには、NFT Music Marketplaceの概要とマーケットを見るボタンを配置しています。
 *
 * Featured NFTs Sectionには、APIから取得したNFT情報をCardコンポーネントに渡し、Gridレイアウトで配置しています。
 *
 * @returns トップページのJSX
 */
const Home: NextPage = async () => {
  const featuredNfts = await getFeaturedNfts();

  return (
    <>
      <Header />
      <main className="w-full flex flex-col items-center">
        {/* Hero Section */}
        <section className="w-full text-center bg-gray-100 py-20 px-4">
          <h1 className="heading1">音楽の新しい価値を見つけよう</h1>
          <p className="text-lg text-gray-600 mb-8 max-w-2xl mx-auto">
            NFT Music Marketplaceは、お気に入りの音楽を所有し、アーティストを直接サポートできる場所です。
          </p>
          <Link
            href="/market"
            className="bg-blue-600 text-white font-bold py-3 px-8 rounded-lg hover:bg-blue-700 transition-colors"
          >
            マーケットを見る
          </Link>
        </section>

        {/* Featured NFTs Section */}
        <section className="py-16 px-4 w-full max-w-7xl">
          <h2 className="text-3xl font-bold text-center mb-10">注目のNFT</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
            {featuredNfts && featuredNfts.length > 0 ? (
              featuredNfts.map((nft) => <NftCard key={nft.id} nft={nft} />)
            ) : (
              <p className="col-span-full text-center text-gray-500">現在、注目のNFTはありません。</p>
            )}
          </div>
        </section>
      </main>
      <Footer />
    </>
  );
};
export default Home;
