import { BigNumberish } from "ethers";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

export type NFTItem = {
  price: string;
  name: string;
  description: string;
  image: string;
};

export type NFTForm = Omit<NFTItem, "image">;

export type MarketItem = {
  tokenId: BigNumberish;
  price: BigNumberish;
  seller: string;
  owner: string;
  sold?: boolean;
  description: string;
  name: string;
  tokenURI: string;
};

export type RenderableMarketItem = {
  id: string;
  user_id: string;
  chain_id: number;
  tokenId: number;
  nonce: number;
  name: string;
  description: string;
  file_type: string;
  image_url: string;
  audio_url: string;
  video_url: string;
  contract_address: string;
  token_url: string;
  genre_id: string;
  genre_name: string;
  creator_address: string;
  created_at: string;
  updated_at: string;
  from: string;
  to: string;
  price: string;
  sale?: boolean;
  insentive: number;
};

export type RenderableMarketItem2 = Omit<MarketItem, "tokenId" | "price"> & {
  tokenId: number;
  price: string;
  image: string; // todo maybe in smart contract also?
};

type Waitable<T> = T & {
  wait: () => Promise<void>;
};

export type MarketContract = {
  updateListingPrice: (listingPrice: BigNumberish) => Promise<void>;
  getListingPrice: () => Promise<number>;
  createToken: (url: string, price: BigNumberish, message: { value: string }) => Promise<Waitable<void>>;
  createMarketItem: (tokenId: BigNumberish, price: BigNumberish) => Promise<void>;
  resellToken: (tokenId: BigNumberish, price: BigNumberish, message: { value: string }) => Promise<Waitable<void>>;
  createMarketSale: (tokenId: BigNumberish, message: { value: BigNumberish }) => Promise<Waitable<void>>;
  fetchMarketItems: () => Promise<MarketItem[]>;
  fetchMyNFTs: () => Promise<MarketItem[]>;
  fetchItemsListed: () => Promise<MarketItem[]>;
  tokenURI: (tokenId: BigNumberish) => Promise<string>;
};

export type UploadToIPFS = (file: File, toThirdWebStorage: boolean) => Promise<string | undefined>;

export type CreateNFT = (formInput: NFTForm, fileUrl: string, router: AppRouterInstance) => Promise<void>;

export type FetchNFTs = () => Promise<RenderableMarketItem[]>;

export type FetchMyNFTsOrListedNFTs = (type: "fetchItemsListed" | "fetchMyNFTs") => Promise<RenderableMarketItem[]>;

export type BuyNFT = (nft: RenderableMarketItem) => Promise<void>;

export type CreateSale = (url: string, formInputPrice: string, isReselling?: boolean, id?: string) => Promise<void>;

export type Context = {
  currentAccount: string;
  nftCurrency: string;
  connectWallet: () => Promise<void>;
  uploadToIPFS: UploadToIPFS;
  createNFT: CreateNFT;
  fetchNFTs: FetchNFTs;
  fetchMyNFTsOrListedNFTs: FetchMyNFTsOrListedNFTs;
  buyNFT: BuyNFT;
  createSale: CreateSale;
  isLoadingNFT: boolean;
};

export type Genre = {
  id: string;
  name: string;
};
