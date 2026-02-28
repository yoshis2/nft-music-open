"use client"; // モバイルメニューの開閉状態を管理するために必要

import React, { useState } from "react";
import Link from "next/link";
import { ethers } from "ethers";

const Header = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };
  // --- シンプルなウォレット接続のための状態管理 ---
  const [isConnected, setIsConnected] = useState(false);
  const [walletAddress, setWalletAddress] = useState("");

  // --- シンプルなウォレット接続/切断処理 ---
  const handleWalletConnect = async () => {
    // window.ethereum が存在するか確認 (MetaMaskなどのウォレット拡張機能)
    if (typeof window.ethereum === "undefined") {
      alert("ウォレット拡張機能が見つかりません。MetaMaskなどをインストールしてください。");
      return;
    }

    try {
      if (!isConnected) {
        // ethers.js v6 を使用したウォレット接続
        const provider = new ethers.BrowserProvider(window.ethereum);
        // アカウントへのアクセスを要求 (ユーザーにウォレット接続のプロンプトを表示)
        await provider.send("eth_requestAccounts", []);
        const signer = await provider.getSigner();
        const _walletAddress = await signer.getAddress();

        setIsConnected(true);
        setWalletAddress(_walletAddress);
        console.log("Wallet connected:", _walletAddress);
      } else {
        // 「切断」はアプリケーション側の状態をリセットすることを意味します
        setIsConnected(false);
        setWalletAddress("");
        console.log("Wallet disconnected by user action in app.");
      }
    } catch (error) {
      console.error("ウォレット接続または切断中にエラーが発生しました:", error);
      alert(`エラーが発生しました: ${error instanceof Error ? error.message : String(error)}`);
      // エラー発生時も状態をリセット
      setIsConnected(false);
      setWalletAddress("");
    }
  };

  return (
    <header className="flex h-20 w-full shrink-0 items-center px-4 md:px-6 bg-slate-50">
      {/* Mobile Menu Button */}
      <button
        className="lg:hidden p-2 rounded-md focus:outline-none focus:ring-2 focus:ring-inset focus:ring-gray-500"
        onClick={toggleMobileMenu}
        aria-expanded={isMobileMenuOpen}
        aria-controls="mobile-menu"
      >
        <MenuIcon className="h-6 w-6" />
        <span className="sr-only">Toggle navigation menu</span>
      </button>

      <Link href="/" className="flex items-center gap-2 font-semibold lg:hidden" prefetch={false}>
        <MountainIcon className="h-6 w-6" />
        NFTミュージック
        <span className="sr-only">Acme Inc</span>
      </Link>
      {/* Mobile Menu (Drawer) */}
      {isMobileMenuOpen && (
        <div
          id="mobile-menu"
          className="absolute top-20 left-0 w-full bg-white dark:bg-gray-950 shadow-lg lg:hidden z-50 p-4"
        >
          <div className="flex items-center gap-2 font-semibold mb-2"></div>
          <p className="text-sm text-gray-600 dark:text-gray-400 mb-4 px-6">
            スリーネクストが運営する NFT Musicの販売サイト
          </p>
          <nav>
            <div className="grid px-6">
              <Link href="/search" className="moblink-primary" prefetch={false}>
                <SearchIcon className="h-5 w-5" />
                検索
              </Link>
              <Link href="/about" className="moblink-primary" prefetch={false}>
                <UserIcon className="h-5 w-5" />
                プロフィール
              </Link>
              <Link href="/contact" className="moblink-primary" prefetch={false}>
                <MailIcon className="h-5 w-5" />
                お問合せ
              </Link>
              <Link href="/creators" className="moblink" prefetch={false}>
                <UserIcon className="h-5 w-5" />
                製作者画面
              </Link>
              <Link href="/master/businesses/list" className="moblink" prefetch={false}>
                <UserIcon className="h-5 w-5" />
                管理画面
              </Link>
              {/* --- ウォレット接続ボタンを他のメニューアイテムと同じグループに移動 --- */}
              <button onClick={handleWalletConnect} type="button" className="moblink-wallet">
                <WalletIcon className="h-5 w-5 mr-1" />
                {isConnected
                  ? `接続済み: ${walletAddress.substring(0, 6)}...${walletAddress.substring(walletAddress.length - 4)}`
                  : "ウォレット接続"}
              </button>
            </div>
            <div className="flex h-5 items-center justify-between py-6"></div>{" "}
          </nav>
        </div>
      )}

      <Link href="/" className="mr-6 hidden lg:flex items-center gap-2 font-semibold" prefetch={false}>
        <MountainIcon className="h-6 w-6" />
        スリーネクストNFTミュージック
        <span className="sr-only">Acme Inc</span>
      </Link>
      {/* Desktop Menu */}
      <nav className="hidden lg:flex h-full space-x-4 ml-auto items-center">
        <div className="inline-flex h-full" role="group">
          <Link href="/search" className="group weblink-primary" prefetch={false}>
            <SearchIcon className="h-5 w-5 mr-1" />
            検索
          </Link>
          <Link href="/about" className="group weblink-primary" prefetch={false}>
            <UserIcon className="h-5 w-5 mr-1" />
            プロフィール
          </Link>
          <Link href="/contact" className="group weblink-primary" prefetch={false}>
            <MailIcon className="h-5 w-5 mr-1" />
            お問合せ
          </Link>
          <Link href="/creators" className="group weblink" prefetch={false}>
            <UserIcon className="h-5 w-5 mr-1" />
            製作者画面
          </Link>
          <Link href="/master/businesses/list" className="group weblink" prefetch={false}>
            <UserIcon className="h-5 w-5 mr-1" />
            管理画面
          </Link>
          {/* --- ここからシンプルなウォレット接続ボタン --- */}
          <button onClick={handleWalletConnect} type="button" className="group  weblink-wallet">
            <WalletIcon className="h-5 w-5 mr-1" />
            {isConnected
              ? `接続済み: ${walletAddress.substring(0, 6)}...${walletAddress.substring(walletAddress.length - 4)}`
              : "ウォレット接続"}
          </button>
          {/* --- ここまでシンプルなウォレット接続ボタン --- */}
        </div>
      </nav>
      <div className="ml-auto lg:hidden">{/* モバイル表示時の右側のスペース確保用、またはここに別の要素を配置 */}</div>
    </header>
  );
};

export default Header;

function MailIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect width="20" height="16" x="2" y="4" rx="2" />
      <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
    </svg>
  );
}

function MenuIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <line x1="4" x2="20" y1="12" y2="12" />
      <line x1="4" x2="20" y1="6" y2="6" />
      <line x1="4" x2="20" y1="18" y2="18" />
    </svg>
  );
}

function WalletIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M21 12V7H5a2 2 0 0 1 0-4h14v4" />
      <path d="M3 5v14a2 2 0 0 0 2 2h16v-5" />
      <path d="M18 12a2 2 0 0 0 0 4h4v-4Z" />
    </svg>
  );
}

function MountainIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="m8 3 4 8 5-5 5 15H2L8 3z" />
    </svg>
  );
}

function UserIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
      <circle cx="12" cy="7" r="4" />
    </svg>
  );
}

function SearchIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props} // propsを展開して、呼び出し元からのclassNameなどを適用
      xmlns="http://www.w3.org/2000/svg"
      // width="32"   // classNameでサイズが指定されることを期待する場合、
      // height="32"  // これらのハードコードされた属性は削除またはコメントアウトします。
      viewBox="0 0 24 24"
    >
      <g fill="none" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2">
        <path d="m13.5 8.5l-5 5m0-5l5 5" />
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21l-4.3-4.3" />
      </g>
    </svg>
  );
}
