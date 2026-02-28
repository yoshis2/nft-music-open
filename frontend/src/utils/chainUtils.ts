/**
 * チェーンIDに基づいてネットワーク名とエクスプローラーURLを返す
 */
export const getChainInfo = (chainId: number) => {
  switch (chainId) {
    case 1:
      return { name: "Ethereum", explorerUrl: "https://etherscan.io" };
    case 137:
      return { name: "Polygon", explorerUrl: "https://polygonscan.com" };
    case 11155111:
      return { name: "Sepolia", explorerUrl: "https://sepolia.etherscan.io" };
    case 1337:
      return { name: "Local Network", explorerUrl: "" }; // ローカルではリンクなし
    default:
      return { name: `Unknown Chain (${chainId})`, explorerUrl: "" };
  }
};
