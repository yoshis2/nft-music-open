"use client";

import type { NextPage } from "next";
import { FormEvent } from "react";

import Footer from "@/components/layouts/footer";
import Header from "@/components/layouts/header";
import { NftCard } from "@/components/layouts/nft-card";
import { useSearch } from "@/hooks/useSearch";

const SearchPage: NextPage = () => {
  const {
    query,
    setQuery,
    genre,
    setGenre,
    minPrice,
    setMinPrice,
    maxPrice,
    setMaxPrice,
    sort,
    setSort,
    genres,
    results,
    isLoading,
    error,
    hasSearched,
    search,
  } = useSearch();

  const handleSearch = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await search();
  };

  return (
    <>
      <Header />
      <main className="main-container">
        <h1 className="heading1">NFTを検索</h1>
        <form onSubmit={handleSearch} className="max-w-4xl mx-auto mb-12 bg-white p-8 rounded-lg shadow-md">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="md:col-span-2">
              <label htmlFor="search-query" className="block text-sm font-medium text-gray-700 mb-1">
                キーワード
              </label>
              <input
                id="search-query"
                type="text"
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                placeholder="キーワードで検索..."
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label htmlFor="genre-select" className="block text-sm font-medium text-gray-700 mb-1">
                ジャンル
              </label>
              <select
                id="genre-select"
                value={genre}
                onChange={(e) => setGenre(e.target.value)}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">すべてのジャンル</option>
                {genres.map((g) => (
                  <option key={g.id} value={g.id}>
                    {g.name}
                  </option>
                ))}
              </select>
            </div>
            <div>
              <label htmlFor="sort-select" className="block text-sm font-medium text-gray-700 mb-1">
                並び順
              </label>
              <select
                id="sort-select"
                value={sort}
                onChange={(e) => setSort(e.target.value)}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="newest">新着順</option>
                <option value="price_asc">価格の安い順</option>
                <option value="price_desc">価格の高い順</option>
              </select>
            </div>
            <div>
              <label htmlFor="min-price" className="block text-sm font-medium text-gray-700 mb-1">
                価格（最小）
              </label>
              <input
                id="min-price"
                type="number"
                value={minPrice}
                onChange={(e) => setMinPrice(e.target.value)}
                placeholder="例: 100"
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label htmlFor="max-price" className="block text-sm font-medium text-gray-700 mb-1">
                価格（最大）
              </label>
              <input
                id="max-price"
                type="number"
                value={maxPrice}
                onChange={(e) => setMaxPrice(e.target.value)}
                placeholder="例: 1000"
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
          </div>
          <div className="mt-6">
            <button
              type="submit"
              className="w-full bg-blue-600 text-white font-bold px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors disabled:bg-gray-400"
              disabled={isLoading}
            >
              {isLoading ? "検索中..." : "検索"}
            </button>
          </div>
        </form>

        {isLoading && <p className="text-center">検索しています...</p>}
        {error && <p className="text-center text-red-500">{error}</p>}

        {hasSearched && !isLoading && !error && (
          <div>
            {results.length > 0 ? (
              <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                {results.map((nft) => (
                  <NftCard key={nft.id} nft={nft} />
                ))}
              </div>
            ) : (
              <p className="text-center text-gray-500">検索結果が見つかりませんでした。</p>
            )}
          </div>
        )}
      </main>
      <Footer />
    </>
  );
};

export default SearchPage;
