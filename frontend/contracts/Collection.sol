// SPDX-License-Identifier: MIT
// Solidityのバージョンを最新版に更新
pragma solidity ^0.8.24;

// OpenZeppelinの最新版（v5.x）をインポート
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title Collection
 * @dev ERC721の作成、修正、削除の機能を持つNFTコレクションコントラクト
 * @dev OpenZeppelin v5.x と Solidity v0.8.26+ に対応
 */
contract Collection is ERC721, Ownable {
  // トークンIDを管理するための変数。0から始まります。
  uint256 private _nextTokenId;

  // ベースURI（例: "https://api.example.com/nfts/"）
  string private _baseTokenURI;

  /**
   * @dev コンストラクタ
   * @param name_ NFTコレクションの名前
   * @param symbol_ NFTコレクションのシンボル
   * @param initialOwner コントラクトの初期オーナーアドレス
   */
  constructor(
    string memory name_,
    string memory symbol_,
    address initialOwner
  ) ERC721(name_, symbol_) Ownable(initialOwner) {}

  /**
   * @notice ベースURIを設定する関数（オーナーのみ実行可能）
   * @param baseURI_ 新しいベースURI
   */
  function setBaseURI(string memory baseURI_) public onlyOwner {
    _baseTokenURI = baseURI_;
  }

  /**
   * @dev 全てのトークンURIの前に付与されるベースURIを返す
   */
  function _baseURI() internal view override returns (string memory) {
    return _baseTokenURI;
  }

  // =================================================================
  // 1. 作成 (Creation)
  // =================================================================
  event Minted(address indexed to, uint256 indexed tokenId);
  /**
   * @notice 新しいNFTをミント（作成）する関数
   * @dev オーナーのみが実行可能
   */
  function safeMint(address to) public onlyOwner {
    // 現在のIDを新しいトークンIDとして使用
    uint256 tokenId = _nextTokenId;
    // 次のIDのためにカウンターを1増やす
    _nextTokenId++;
    _safeMint(to, tokenId);

    emit Minted(to, tokenId);
  }

  // =================================================================
  // 2. 修正 (Modification)
  // =================================================================

  /**
   * @notice NFTのメタデータ（画像や情報）は、通常、ベースURIとトークンIDを
   * 組み合わせて生成されるJSONファイルで管理されます。
   * メタデータ自体を変更したい場合は、サーバー上のJSONファイルを更新します。
   * もし、全てのトークンのURI構造の基点となる部分を変更したい場合は、
   * 上記の`setBaseURI`関数を使用します。
   */

  // =================================================================
  // 3. 削除 (Deletion)
  // =================================================================

  /**
   * @notice NFTをバーン（削除）する関数
   * @dev トークンの所有者、または承認されたアドレスのみが実行可能
   * @param tokenId バーンするNFTのID
   */
  function burn(uint256 tokenId) public virtual {
    // 呼び出し元がトークンの所有者または承認されたアドレスであることを確認します。
    // このチェックは、トークンが存在しない場合にもエラーを発生させます。
    _checkAuthorized(ownerOf(tokenId), msg.sender, tokenId);
    _burn(tokenId);
  }
}
