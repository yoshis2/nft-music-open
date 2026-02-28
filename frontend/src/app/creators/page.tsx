"use client";

import type { NextPage } from "next";
import Link from "next/link";
import { useEffect, useState, useCallback } from "react";
import { ethers } from "ethers";

// APIから返されるユーザー情報の型
interface UserProfile {
  id: number;
  name: string;
  email: string;
  // 必要に応じて他のプロパティを追加
}

const CreatorsPage: NextPage = () => {
  const [user, setUser] = useState<UserProfile | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchUser = useCallback(async () => {
    if (!window.ethereum) {
      setError("ウォレットプロバイダーが見つかりません。ページをリフレッシュして再度お試しください。");
      setIsLoading(false);
      return;
    }
    setIsLoading(true);
    setError(null);
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const walletAddress = await signer.getAddress();

      // バックエンドにユーザー情報を問い合わせ (APIエンドポイントは要確認)
      const response = await fetch(`/api/users?wallet=${walletAddress}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        const userData = await response.json();
        if (userData && userData.id) {
          setUser(userData);
          console.log("ユーザー情報:", userData);
        } else {
          setUser(null);
        }
      } else if (response.status === 404) {
        setUser(null);
      } else {
        console.error(`APIエラー: ステータス ${response.status}`);
        const errorText = await response.text();
        console.error("レスポンスボディ:", errorText);
        setError("ユーザー情報の取得に失敗しました。");
      }
    } catch (err: unknown) {
      if (err instanceof Error) {
        console.error(err);
        setError(err.message || "ユーザー情報の取得中にエラーが発生しました。");
      } else {
        console.error("Unknown error:", err);
        setError("ユーザー情報の取得中に不明なエラーが発生しました。");
      }
    } finally {
      setIsLoading(false);
    }
  }, []);

  const handleConnectWallet = useCallback(async () => {
    if (!window.ethereum) {
      setError("ウォレットがインストールされていません。");
      return;
    }
    try {
      setIsLoading(true);
      setError(null);
      const provider = new ethers.BrowserProvider(window.ethereum);
      await provider.send("eth_requestAccounts", []);
      await fetchUser();
    } catch (err: unknown) {
      if (
        typeof err === "object" &&
        err !== null &&
        "code" in err &&
        (err as { code: unknown }).code === "ACTION_REJECTED"
      ) {
        setError("ウォレットへの接続が拒否されました。");
      } else {
        console.error("ウォレット接続中にエラーが発生しました:", err);
        setError("ウォレット接続中にエラーが発生しました。");
      }
      setIsLoading(false);
    }
  }, [fetchUser]);

  useEffect(() => {
    const checkConnection = async () => {
      if (!window.ethereum) {
        setError("ウォレットがインストールされていません。");
        setIsLoading(false);
        return;
      }

      const provider = new ethers.BrowserProvider(window.ethereum);
      const accounts = await provider.listAccounts();
      if (accounts.length > 0) {
        fetchUser();
      } else {
        setError("ウォレットに接続されていません。");
        setIsLoading(false);
      }
    };
    checkConnection();
  }, [fetchUser]);

  const renderContent = () => {
    if (isLoading) {
      return <p>読み込み中...</p>;
    }

    if (error) {
      return (
        <>
          <div
            className="p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
            role="alert"
          >
            <span className="font-medium">エラー:</span> {error}
          </div>
          {error === "ウォレットがインストールされていません。" && (
            <div className="mt-6 p-4 bg-blue-50 border border-blue-200 rounded-lg text-center">
              <p className="text-blue-800 mb-3">Web3機能を利用するには、MetaMaskウォレットが必要です。</p>
              <a
                href="https://metamask.io/download/"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-block px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-orange-500 hover:bg-orange-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500"
              >
                MetaMaskをインストールする
              </a>
            </div>
          )}
          {(error === "ウォレットに接続されていません。" || error === "ウォレットへの接続が拒否されました。") && (
            <div className="mt-6 p-4 bg-blue-50 border border-blue-200 rounded-lg text-center">
              <p className="text-blue-800 mb-3">Web3機能を利用するには、ウォレットを接続してください。</p>
              <button
                onClick={handleConnectWallet}
                className="inline-block px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                ウォレットを接続
              </button>
            </div>
          )}
        </>
      );
    }

    if (user) {
      // 登録済みユーザー向けのダッシュボード
      return (
        <div>
          <h2 className="text-2xl font-bold mb-4">ようこそ、{user.name}さん</h2>
          <p className="mb-6">クリエイターダッシュボードから、活動を始めましょう。</p>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <Link
              href="/creators/nft/mint"
              className="p-6 bg-gray-50 hover:bg-gray-100 border border-gray-200 rounded-lg shadow-sm block"
            >
              <h3 className="text-xl font-semibold">NFTを作成する</h3>
              <p className="text-gray-600">新しい音楽NFTを作成して出品します。</p>
            </Link>
            <Link
              href={`/creators/users/update`}
              className="p-6 bg-gray-50 hover:bg-gray-100 border border-gray-200 rounded-lg shadow-sm block"
            >
              <h3 className="text-xl font-semibold">ユーザー編集</h3>
              <p className="text-gray-600">あなたのユーザー登録情報を更新します。</p>
            </Link>
          </div>
        </div>
      );
    }

    // 未登録ユーザー向けの案内
    return (
      <div>
        <h2 className="text-2xl font-bold mb-4">クリエイター登録へようこそ</h2>
        <p className="mb-6">
          まだクリエイターとして登録されていません。
          <br />
          以下のボタンからアカウントを作成し、あなたの音楽を世界に届けましょう。
        </p>
        <Link href="/creators/users/create" className="submit-button inline-block">
          アカウントを作成する
        </Link>
      </div>
    );
  };

  return (
    <main className="main-container max-w-4xl mx-auto">
      <h1 className="heading1">クリエイターページ</h1>
      <section className="p-6 mb-6 bg-white border border-gray-200 rounded-lg shadow w-full">{renderContent()}</section>
    </main>
  );
};
export default CreatorsPage;
