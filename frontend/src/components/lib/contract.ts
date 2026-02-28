import { ethers, type Signer } from "ethers";

import { type MarketContract, MarketItem, RenderableMarketItem2 } from "@/types/types";

import {
  LocalMarketAddress,
  SepoliaMarketAddress,
  AmoyMarketAddress,
  MarketAddressABI,
} from "@/components/lib/constants";

export const fetchContract = async (
  chainId: number | undefined,
  signerOrProvider?: Signer | ethers.BrowserProvider,
): Promise<MarketContract> => {
  const contractAddress = ContractAddress(chainId);
  return new ethers.Contract(contractAddress, MarketAddressABI, signerOrProvider) as unknown as MarketContract;
};

export const formatBalance = (rawBalance: string) => {
  const balance = (parseInt(rawBalance) / 1000000000000000000).toFixed(2);
  return balance;
};

export const formatChainAsNum = (chainIdHex: string) => {
  const chainIdNum = parseInt(chainIdHex);
  return chainIdNum;
};

export const formatAddress = (addr: string | undefined) => {
  return `${addr?.substring(0, 8)}...`;
};

export const getRenderableData = async (
  data: MarketItem[],
  chainId: number,
  contract: MarketContract,
): Promise<RenderableMarketItem2[]> => {
  const items: (RenderableMarketItem2 | undefined)[] = await Promise.all(
    data.map(async ({ tokenId, seller, owner, price: unformattedPrice }) => {
      let tokenURI = await contract.tokenURI(tokenId);
      if (chainId == 1337) {
        tokenURI = tokenURI.replace(
          "https://ipfs.io",
          "http://ipfs:8080", // TODO: URLの再修正が必要
        );
      }

      let metaData;
      try {
        const res = await fetch("/api/blockchain/rendering", {
          method: "POST",
          body: JSON.stringify(tokenURI),
        });

        if (res.status == 500) {
          return undefined;
        }

        const response = await res.json();
        metaData = JSON.parse(response);

        if (!metaData) return undefined;
      } catch (e) {
        console.error("Block Chainの処理中にエラーが発生しました");
        console.error(e);
        return undefined;
      }
      const { name, description, image, insentive } = metaData;
      const price = ethers.formatUnits(unformattedPrice, "ether");

      return {
        tokenId: Number(tokenId),
        price,
        seller,
        owner,
        sold: false,
        description,
        name,
        tokenURI,
        image,
        insentive,
      } as RenderableMarketItem2;
    }),
  );
  return items.filter((item): item is RenderableMarketItem2 => item !== undefined);
};

export const ContractAddress = (chainId: number | undefined) => {
  switch (String(chainId)) {
    case "1337":
      return LocalMarketAddress;
    case "80002":
      return AmoyMarketAddress;
    case "11155111":
      console.log("セポリア");
      return SepoliaMarketAddress;
    default:
      console.log("無アドレス");
      return LocalMarketAddress;
  }
};
