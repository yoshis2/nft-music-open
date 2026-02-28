import { useState, useEffect } from "react";
import { RenderableMarketItem, Genre } from "@/types/types";

export const useSearch = () => {
  const [query, setQuery] = useState("");
  const [genre, setGenre] = useState("");
  const [minPrice, setMinPrice] = useState("");
  const [maxPrice, setMaxPrice] = useState("");
  const [sort, setSort] = useState("newest");
  const [genres, setGenres] = useState<Genre[]>([]);

  const [results, setResults] = useState<RenderableMarketItem[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [hasSearched, setHasSearched] = useState(false);

  useEffect(() => {
    const fetchGenres = async () => {
      try {
        const res = await fetch(`/api/master/genres`, {
          method: "GET",
        });
        if (!res.ok) {
          throw new Error("ジャンルの取得に失敗しました。");
        }
        const data: Genre[] = await res.json();
        setGenres(data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchGenres();
  }, []);

  const search = async () => {
    setIsLoading(true);
    setError(null);
    setHasSearched(true);

    try {
      const params = new URLSearchParams();
      if (query.trim()) params.append("q", query);
      if (genre) params.append("genre", genre);
      if (minPrice) params.append("min_price", minPrice);
      if (maxPrice) params.append("max_price", maxPrice);
      if (sort) params.append("sort", sort);

      const res = await fetch(`/api/nft/search?params=${params.toString()}`, {
        cache: "no-store",
      });
      if (!res.ok) {
        throw new Error("検索結果の取得に失敗しました。");
      }
      const data: RenderableMarketItem[] = await res.json();
      setResults(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "不明なエラーが発生しました。");
      setResults([]);
    } finally {
      setIsLoading(false);
    }
  };

  return {
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
  };
};
