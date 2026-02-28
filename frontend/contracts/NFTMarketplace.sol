// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract NFTMarketplace is ERC721URIStorage, Ownable, Pausable {
  uint256 private _tokenIds;
  uint256 private _itemsSold;

  uint256 public listingPrice;
  uint256 public royaltyFeeBps; // クリエイターへのロイヤリティ手数料率 (BPS: 1/100 of 1%)

  mapping(uint256 => MarketItem) private idToMarketItem;

  struct MarketItem {
    uint256 tokenId;
    address payable seller;
    address owner; // MarketItemの所有者は支払いを受ける必要がないため、payableではありません
    uint256 price;
    bool sold;
    address payable creator; // NFTの最初の作成者
  }

  event MarketItemCreated(
    uint256 indexed tokenId,
    address seller,
    address owner,
    uint256 price,
    bool sold,
    address creator
  );

  constructor(
    string memory _name,
    string memory _symbol,
    uint256 _initialListingPrice,
    uint256 _initialRoyaltyFeeBps
  ) ERC721(_name, _symbol) Ownable(msg.sender) {
    listingPrice = _initialListingPrice;
    royaltyFeeBps = _initialRoyaltyFeeBps;
  }

  function pause() public onlyOwner {
    _pause();
  }

  function unpause() public onlyOwner {
    _unpause();
  }

  function updateRoyaltyFee(uint256 _royaltyFeeBps) public onlyOwner {
    royaltyFeeBps = _royaltyFeeBps;
  }

  function getListingPrice() public view returns (uint256) {
    return listingPrice;
  }

  function updateListingPrice(uint _listingPrice) public onlyOwner {
    listingPrice = _listingPrice;
  }

  // ミント
  function createToken(string memory tokenURL) public payable whenNotPaused returns (uint) {
    _tokenIds++;
    require(msg.value == listingPrice, "Please submit the asking price and listing fee");

    uint256 newTokenId = _tokenIds;

    _mint(msg.sender, newTokenId);
    _setTokenURI(newTokenId, tokenURL);

    idToMarketItem[newTokenId] = MarketItem(
      newTokenId,
      payable(msg.sender),
      address(this), // 所有者はマーケットプレイスのコントラクトです
      msg.value,
      false,
      payable(msg.sender) // このトークンを最初に作成したクリエイター
    );

    _transfer(
      msg.sender,
      address(this), // 送信者のアドレスからこのコントラクトへ
      newTokenId
    );

    emit MarketItemCreated(
      newTokenId,
      msg.sender,
      address(this), // 所有者はマーケットプレイスのコントラクトです
      msg.value,
      false,
      msg.sender
    );

    return newTokenId;
  }

  function resellToken(uint256 tokenId, uint256 price) public payable whenNotPaused {
    require(idToMarketItem[tokenId].owner == msg.sender, unicode"アイテムの所有者のみがこの操作を実行できます");
    require(msg.value == listingPrice, unicode"価格は出品価格と等しくなければなりません");

    idToMarketItem[tokenId].price = price;
    idToMarketItem[tokenId].seller = payable(msg.sender);
    idToMarketItem[tokenId].owner = address(this); // アイテムが販売中のため、所有者はマーケットプレイスのコントラクトになります
    idToMarketItem[tokenId].sold = false;

    _transfer(msg.sender, address(this), tokenId);
  }

  event MarketItemSold(uint256 indexed tokenId, address indexed seller, address indexed buyer, uint256 price);

  error InsufficientFunds(uint256 tokenId, uint256 sent);

  function createMarketSale(uint256 tokenId) public payable whenNotPaused {
    uint price = idToMarketItem[tokenId].price;
    address seller = idToMarketItem[tokenId].seller;
    address payable creator = idToMarketItem[tokenId].creator;
    // ✅ ここでイベントを発生させてログを記録する
    emit MarketItemSold(
      tokenId, // 売れたアイテムのID
      seller, // 元の販売者
      msg.sender, // 新しい所有者（購入者）
      price // 価格
    );

    // requireの代わりにif文を使う
    if (msg.value != price) {
      // もし支払額が価格と等しくなければ...
      revert InsufficientFunds(price, msg.value);
    }

    // 手数料とロイヤリティを計算
    uint256 royaltyAmount = (price * royaltyFeeBps) / 10000;
    uint256 marketplaceFee = listingPrice; // 固定の出品手数料
    uint256 sellerProceeds = price - royaltyAmount - marketplaceFee;

    // 各所への送金処理（より安全な .call を使用）
    (bool sentMarketplace, ) = owner().call{ value: marketplaceFee }("");
    require(sentMarketplace, "Failed to send funds to marketplace owner");

    (bool sentCreator, ) = creator.call{ value: royaltyAmount }("");
    require(sentCreator, "Failed to send funds to creator");

    (bool sentSeller, ) = payable(seller).call{ value: sellerProceeds }("");
    require(sentSeller, "Failed to send funds to seller");

    // 販売を反映するためにマーケットアイテムを更新します
    idToMarketItem[tokenId].owner = msg.sender; // 購入者が新しい所有者です
    idToMarketItem[tokenId].sold = true;
    idToMarketItem[tokenId].seller = payable(address(0)); // アイテムが売れたため、販売者はクリアされます

    _itemsSold++;

    _transfer(address(this), msg.sender, tokenId);
  }

  function fetchMarketItems() public view returns (MarketItem[] memory) {
    uint totalItemCount = _tokenIds;
    uint itemCount = 0;
    uint currentIndex = 0;

    for (uint i = 0; i < totalItemCount; i++) {
      // 販売中のアイテム（つまり、このコントラクトが所有しているアイテム）を確認します
      if (idToMarketItem[i + 1].owner == address(this)) {
        itemCount += 1;
      }
    }

    MarketItem[] memory items = new MarketItem[](itemCount);
    for (uint i = 0; i < totalItemCount; i++) {
      // 販売中のアイテムを再度確認し、配列に追加します
      if (idToMarketItem[i + 1].owner == address(this)) {
        uint currentId = i + 1;
        MarketItem storage currentItem = idToMarketItem[currentId];
        items[currentIndex] = currentItem;
        currentIndex += 1;
      }
    }

    return items;
  }

  function fetchMyNFTs() public view returns (MarketItem[] memory) {
    uint totalItemCount = _tokenIds;
    uint itemCount = 0;
    uint currentIndex = 0;

    for (uint i = 0; i < totalItemCount; i++) {
      if (idToMarketItem[i + 1].owner == msg.sender) {
        itemCount += 1;
      }
    }

    MarketItem[] memory items = new MarketItem[](itemCount);

    for (uint i = 0; i < totalItemCount; i++) {
      if (idToMarketItem[i + 1].owner == msg.sender) {
        uint currentId = i + 1;

        MarketItem storage currentItem = idToMarketItem[currentId];

        items[currentIndex] = currentItem;

        currentIndex += 1;
      }
    }

    return items;
  }

  function fetchItemsListed() public view returns (MarketItem[] memory) {
    uint totalItemCount = _tokenIds;
    uint itemCount = 0;
    uint currentIndex = 0;

    for (uint i = 0; i < totalItemCount; i++) {
      if (idToMarketItem[i + 1].seller == msg.sender) {
        itemCount += 1;
      }
    }

    MarketItem[] memory items = new MarketItem[](itemCount);

    for (uint i = 0; i < totalItemCount; i++) {
      if (idToMarketItem[i + 1].seller == msg.sender) {
        uint currentId = i + 1;

        MarketItem storage currentItem = idToMarketItem[currentId];

        items[currentIndex] = currentItem;

        currentIndex += 1;
      }
    }

    return items;
  }

  /**
   * @dev これまでに作成された全てのマーケットアイテムを取得します（販売中、売却済みなど状態を問わない）。
   * @return MarketItem[] 全てのマーケットアイテムの配列
   */
  function fetchAllMarketItems() public view returns (MarketItem[] memory) {
    uint256 totalItemCount = _tokenIds;
    MarketItem[] memory items = new MarketItem[](totalItemCount);

    // ループを回して全てのアイテム情報を配列に格納します
    for (uint256 i = 0; i < totalItemCount; i++) {
      uint256 currentId = i + 1;
      items[i] = idToMarketItem[currentId];
    }

    return items;
  }
}
