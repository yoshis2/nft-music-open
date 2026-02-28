import Image from "next/image";
import Link from "next/link";
import { RenderableMarketItem } from "@/types/types";

type NftCardProps = {
  nft: RenderableMarketItem;
};

export const NftCard = ({ nft }: NftCardProps) => {
  const isLocal = nft.chain_id === 1337;
  const ipfsGateway = isLocal ? "http://127.0.0.1:8080" : "https://ipfs.io";

  const imageUrl = nft.image_url ? `${ipfsGateway}${nft.image_url}` : "/placeholder-image.png";
  const audioUrl = nft.audio_url ? `${ipfsGateway}${nft.audio_url}` : "";
  const videoUrl = nft.video_url ? `${ipfsGateway}${nft.video_url}` : "";

  const priceInEth = nft.price; // 価格はAPIから取得したものをそのまま使う

  return (
    <Link href={`/nfts/detail/${nft.id}`} className="block group">
      <div className="bg-white rounded-lg shadow-md overflow-hidden transition-shadow duration-300 hover:shadow-xl">
        <div className="relative w-full h-56 bg-gray-200">
          {nft.file_type === "video" && videoUrl ? (
            <video
              loop
              muted
              autoPlay
              playsInline
              src={videoUrl}
              className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
              poster={imageUrl}
            ></video>
          ) : (
            <>
              <Image
                src={imageUrl}
                alt={nft.name}
                fill
                sizes="(max-width: 640px) 100vw, (max-width: 1024px) 50vw, 25vw"
                className="object-cover transition-transform duration-300 group-hover:scale-105"
              />
              {nft.file_type === "audio" && audioUrl && (
                <div className="absolute bottom-0 left-0 right-0 p-1 bg-black bg-opacity-25 backdrop-blur-sm">
                  <audio controls src={audioUrl} className="w-full h-8"></audio>
                </div>
              )}
            </>
          )}
        </div>
        <div className="p-4">
          <h3 className="text-lg font-bold text-gray-900 truncate">{nft.name}</h3>
          <p className="text-sm text-gray-600 mt-1 truncate">{nft.description}</p>
          <div className="mt-4">
            <p className="text-xs text-gray-500">Price</p>
            <p className="text-lg font-semibold text-gray-800">{priceInEth} ETH</p>
          </div>
        </div>
      </div>
    </Link>
  );
};
