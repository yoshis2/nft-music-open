import type { NextPage } from "next";
import Link from "next/link";
import Image from "next/image";

const About: NextPage = async () => {
  return (
    <main className="w-full bg-gray-50">
      {/* Hero Section - Updated */}
      <section className="w-full py-4 md:py-0">
        <div className="max-w-5xl mx-auto px-6">
          <div className="text-center mb-12 pt-12">
            <h1 className="text-3xl md:text-5xl font-bold mb-4 text-gray-900">スリーネクストについて</h1>
            <p className="text-lg md:text-xl text-gray-600">最先端技術で、新たな便利と未来を創出する</p>
          </div>
          <div className="relative w-full max-w-4xl mx-auto aspect-video rounded-lg overflow-hidden shadow-lg">
            <Image
              src="/img/about-threenext.png"
              alt="Threenext Eyecatch"
              fill
              sizes="(max-width: 896px) 100vw, 896px"
              style={{ objectFit: "cover" }}
              priority
            />
          </div>
        </div>
      </section>

      {/* Content Section */}
      <div className="max-w-4xl mx-auto px-6 py-16">
        <p className="mb-16 text-lg text-gray-600 leading-relaxed">
          スリーネクストは、代表のyoshi
          seki氏が運営するWeb3関連事業体です。最先端技術であるブロックチェーンや暗号資産領域での開発・分析事業を主軸としつつ、自身の経験を活かしたエンジニア向けの情報発信も行うなど、多角的な活動を展開しています。
        </p>

        {/* 会社概要 */}
        <section className="w-full mb-16">
          <h2 className="heading2 text-center">会社概要</h2>
          <div className="overflow-x-auto bg-white rounded-lg shadow-md">
            <table className="w-full">
              <tbody>
                <tr>
                  <th className="border-b border-gray-200 p-4 w-[30%] bg-gray-50 text-left font-semibold">
                    会社名/屋号
                  </th>
                  <td className="border-b border-gray-200 p-4">スリーネクスト (Threenext)</td>
                </tr>
                <tr>
                  <th className="border-b border-gray-200 p-4 bg-gray-50 text-left font-semibold">代表者</th>
                  <td className="border-b border-gray-200 p-4">yoshi seki 氏</td>
                </tr>
                <tr>
                  <th className="border-b border-gray-200 p-4 bg-gray-50 text-left font-semibold">所在地</th>
                  <td className="border-b border-gray-200 p-4">東京都港区</td>
                </tr>
                <tr>
                  <th className="border-b border-gray-200 p-4 bg-gray-50 text-left font-semibold">連絡先</th>
                  <td className="border-b border-gray-200 p-4">seki@threenext.com</td>
                </tr>
                <tr>
                  <th className="p-4 bg-gray-50 text-left font-semibold">公式サイト</th>
                  <td className="p-4">
                    <Link href="https://www.threenext.com" target="_blank" className="text-blue-600 hover:underline">
                      スリーネクスト フリーランスエンジニアのための情報サイト
                    </Link>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>

        {/* 事業内容 */}
        <section className="w-full mb-16">
          <h2 className="heading2 text-center">事業内容詳細</h2>
          <div className="space-y-10">
            <div className="bg-white p-8 rounded-lg shadow-md">
              <h3 className="text-2xl font-semibold mb-4 text-gray-800">1. Web3・暗号資産関連事業</h3>
              <p className="mb-4 text-gray-700 leading-relaxed">
                中核事業として、Web3領域における開発および暗号資産の分析を手掛けています。
              </p>
              <div className="ml-4 space-y-4 border-l-4 border-blue-500 pl-6">
                <div>
                  <p className="font-semibold text-gray-800">Web3開発:</p>
                  <p className="text-gray-700 leading-relaxed">
                    ブロックチェーン技術を活用した分散型アプリケーション（DApps）や、スマートコントラクトなどの設計・開発を行っていると考えられます。具体的な開発実績については公開されていませんが、先進技術を用いて新たな価値を創出するという企業理念に基づき、様々なプロジェクトに取り組んでいることが推測されます。
                  </p>
                </div>
                <div>
                  <p className="font-semibold text-gray-800">暗号資産の分析・開発:</p>
                  <p className="text-gray-700 leading-relaxed">
                    市場データやブロックチェーン上のデータを分析し、価値評価や将来予測などを行っています。また、独自の暗号資産（トークン）の開発に関わる事業も展開している可能性があります。
                  </p>
                </div>
              </div>
            </div>

            <div className="bg-white p-8 rounded-lg shadow-md">
              <h3 className="text-2xl font-semibold mb-4 text-gray-800">2. NFTミュージック事業</h3>
              <p className="mb-4 text-gray-700 leading-relaxed">
                ご提供情報にある「NFTミュージック」は、音楽とNFT（非代替性トークン）を組み合わせた新しいデジタルコンテンツ領域の事業です。アーティストの楽曲や関連アートワークの権利をNFT化し、ブロックチェーン上で売買・管理するプラットフォームの開発や、関連サービスの提供を目指しているものと考えられます。
              </p>
              <p className="text-gray-700 leading-relaxed">
                現時点では、これらのWeb3関連事業に関する具体的なサービス内容や実績は公開されていません。今後の正式な情報公開が待たれます。
              </p>
            </div>

            <div className="bg-white p-8 rounded-lg shadow-md">
              <h3 className="text-2xl font-semibold mb-4 text-gray-800">3. メディア事業</h3>
              <p className="mb-4 text-gray-700 leading-relaxed">
                代表自身の経歴を活かし、フリーランスのITエンジニアを対象としたウェブサイトを運営しています。
              </p>
              <div className="ml-4 space-y-4 border-l-4 border-blue-500 pl-6">
                <div>
                  <p className="font-semibold text-gray-800">
                    運用サイト①：フリーランスエンジニアのための情報サイト「スリーネクスト」
                  </p>
                  <p className="text-gray-700 leading-relaxed">
                    代表のseki氏は、2013年からフリーランスエンジニアとして活動してきた経験を持ちます。その実体験に基づき、独立準備から案件獲得の戦略、単価交渉術、税金・社会保険の知識まで、フリーランスとして成功するために不可欠な情報を網羅的に発信しています。技術情報だけでなく、事業運営やお金にまつわる実践的なノウハウを提供することで、エンジニアのキャリアをサポートしています。
                  </p>
                </div>
                <div>
                  <p className="font-semibold text-gray-800">運用サイト②（情報提供より）：仮想通貨の学校</p>
                  <p className="text-gray-700 leading-relaxed">
                    暗号資産に関する教育的なコンテンツを提供するサイトを運営されているとのことですが、現時点では「スリーネクスト」との関連性を公に確認できる情報はありません。暗号資産分析事業の一環として、投資家や学習者向けの情報提供を行っている可能性があります。
                  </p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section className="w-full bg-white p-8 rounded-lg shadow-md">
          <h2 className="heading2 text-center">企業理念・ビジョン</h2>
          <ul className="list-disc list-inside space-y-2 mb-4 text-gray-700 leading-relaxed">
            <li>
              <span className="font-semibold text-gray-800">新たな便利を創出する:</span>{" "}
              常に最新技術のアンテナを張り巡らせ、アイデアをサービスとして具現化し、新たな価値を社会に提供する。
            </li>
            <li>
              <span className="font-semibold text-gray-800">みんなの元気を創出する:</span>{" "}
              やりがいのある仕事を通じて人々の成功を導き、夢をかなえるサポートを行う。
            </li>
            <li>
              <span className="font-semibold text-gray-800">新たな未来を実現する:</span>{" "}
              新しく住みやすい社会の実現を目指し続ける。
            </li>
          </ul>
          <p className="text-gray-700 leading-relaxed">
            これらの理念は、Web3という未来を形作る技術領域での事業展開と、フリーランスエンジニアという個人の働き方を支援するメディア事業の両方に通底する思想となっています。
          </p>
        </section>
      </div>
    </main>
  );
};
export default About;
