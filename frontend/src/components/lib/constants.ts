// MarketAddress
// 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0

// これはコントラクトのコンパイルバージョンです（コントラクトを作成するたびに、あなたはそれを更新する必要があります）。

import market from "@Contract/NFTMarketplace.sol/NFTMarketplace.json";
import sample from "@Contract/EthEcho.sol/EthEcho.json";
import pure from "@Contract/PureView.sol/PureView.json";

export const MarketAddressABI = market.abi;
export const SampleAddressABI = sample.abi;
export const PureAddressABI = pure.abi;

// Local
export const LocalMarketAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
export const LocalEthEchoAddress = "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9";
export const LocalPureAddress = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707";

// Sepolia
export const SepoliaMarketAddress = "0xEdBf654a8F9849eFba41e0FF09873601DC93df3F";

// Amoy
export const AmoyMarketAddress = "0x700b1FB54c1814e1F4Bf343a2fb5B30350bfFB3e";

// API URL
export const FLARE_API_URL = "https://api.flare.network";
export const ETHERSCAN_MAIN_URL = "https://api.etherscan.io/api";
export const ETHERSCAN_SEPOLIA_URL = "https://api-sepolia.etherscan.io/api";

// IPFSローカル
// export const IPFS_HOST="http://127.0.0.1" // ローカル起動の場合
export const LOCAL_HOST = "http://127.0.0.1";
export const IPFS_HOST = "http://ipfs"; // Docker起動の場合
export const IPFS_API_PORT = ":5001";
export const IPFS_GATEWAY_PORT = ":8080";

export const SEPOLIA = 11155111;
