"use client";

import Link from "next/link";
import React, { useState } from "react";
import { ethers } from "ethers";

const AdminHeader = () => {
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

      {/* Mobile Menu (Drawer) */}
      {isMobileMenuOpen && (
        <div
          id="mobile-menu"
          className="absolute top-20 left-0 w-full bg-white dark:bg-gray-950 shadow-lg lg:hidden z-50 p-4"
        >
          <div className="flex items-center gap-2 font-semibold mb-2">
            <Link href="/" className="flex items-center gap-2 font-semibold" prefetch={false}>
              <MountainIcon className="h-6 w-6" />
              NFTミュージック (管理者)
              <span className="sr-only">Acme Inc</span>
            </Link>
          </div>
          <p className="text-sm text-gray-600 dark:text-gray-400 mb-4">
            スリーネクストが運営する NFT Musicの販売サイト
          </p>
          <nav>
            <div className="grid px-6">
              <Link href="/master/businesses/list" className="moblink-primary" prefetch={false}>
                <PackageIcon className="h-5 w-5" />
                職種一覧
              </Link>
              <Link href="/master/genres/list" className="moblink-primary" prefetch={false}>
                <PackageIcon className="h-5 w-5" />
                ジャンル一覧
              </Link>
              <Link href="/" className="moblink" prefetch={false}>
                <UserIcon className="h-5 w-5" />
                一般ページ
              </Link>
              <Link href="/creators" className="moblink" prefetch={false}>
                <UserIcon className="h-5 w-5" />
                クリエイター
              </Link>
              {/* --- ウォレット接続ボタンを他のメニューアイテムと同じグループに移動 --- */}
              <button onClick={handleWalletConnect} type="button" className="moblink-wallet">
                <WalletIcon className="h-5 w-5" />
                {isConnected
                  ? `接続済み: ${walletAddress.substring(0, 6)}...${walletAddress.substring(walletAddress.length - 4)}`
                  : "ウォレット接続"}
              </button>
            </div>
          </nav>
        </div>
      )}

      <Link href="/" className="mr-6 hidden lg:flex items-center gap-2 font-semibold" prefetch={false}>
        <MountainIcon className="h-6 w-6" />
        スリーネクストNFTミュージック (管理者)
        <span className="sr-only">Acme Inc</span>
      </Link>
      {/* Desktop Menu */}
      <nav className="hidden lg:flex h-full space-x-4 ml-auto items-center">
        <div className="inline-flex h-full" role="group">
          <Link href="/master/businesses/list" className="group weblink-primary" prefetch={false}>
            <PackageIcon className="h-5 w-5 mr-1" />
            職種一覧
          </Link>
          <Link href="/master/genres/list" className="group weblink-primary" prefetch={false}>
            <PackageIcon className="h-5 w-5 mr-1" />
            ジャンル一覧
          </Link>
          <Link href="/" className="group weblink" prefetch={false}>
            <UserIcon className="h-5 w-5 mr-1" />
            一般ページ
          </Link>
          <Link href="/creators" className="group weblink" prefetch={false}>
            <UserIcon className="h-5 w-5 mr-1" />
            クリエイター
          </Link>
          {/* --- ここからシンプルなウォレット接続ボタン --- */}
          <button onClick={handleWalletConnect} type="button" className="group weblink-wallet">
            <WalletIcon className="h-5 w-5 mr-1" />
            {isConnected
              ? `接続済み: ${walletAddress.substring(0, 6)}...${walletAddress.substring(walletAddress.length - 4)}`
              : "ウォレット接続"}
          </button>
          {/* --- ここまでシンプルなウォレット接続ボタン --- */}
        </div>
      </nav>
      <div className="ml-auto lg:hidden">{/* モバイル表示時の右側のスペース確保用 */}</div>
    </header>
  );
};

export default AdminHeader;

// Icon components (copied from creators-header.tsx or header.tsx for consistency)
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

function PackageIcon(props: React.SVGProps<SVGSVGElement>) {
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
      <path d="m7.5 4.27 9 5.15" />
      <path d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z" />
      <path d="m3.3 7 8.7 5 8.7-5" />
      <path d="M12 22V12" />
    </svg>
  );
}
