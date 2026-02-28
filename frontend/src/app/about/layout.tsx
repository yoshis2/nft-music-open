import type { Metadata } from "next";

import "../globals.css";
import Header from "@/components/layouts/header";
import Footer from "@/components/layouts/footer";

export const metadata: Metadata = {
  title: "スリーネクストNFTミュージックについて",
  description: "スリーネクストNFTミュージックの事業内容などの詳細情報を紹介します。",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <>
      <Header />
      {children}
      <Footer />
    </>
  );
}
