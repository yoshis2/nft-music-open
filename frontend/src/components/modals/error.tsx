import type { NextPage } from "next";

export type ModalProps = {
  open: boolean;
  kind: string;
  wallet: string;
  onCancel: () => void;
  createAccount: () => void;
};

/**
 * A modal component for displaying an error when a user is not registered.
 * It prompts the user to create an account.
 *
 * @param {ModalProps} props The component props.
 * @param {boolean} props.open Controls if the modal is open.
 * @param {string} props.kind The kind of action that failed (e.g., "NFTデータアップロード").
 * @param {string} props.wallet The user's wallet address.
 * @param {() => void} props.onCancel Function to call when the cancel button is clicked.
 * @param {() => void} props.createAccount Function to call to navigate to the account creation page.
 * @returns {JSX.Element | null} The modal component or null if not open.
 */
const Modal: NextPage<ModalProps> = (props) => {
  if (!props.open) {
    return null;
  }

  return (
    <>
      <div className="bg-white top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-120 h-auto p-6 flex flex-col items-start absolute z-20 rounded-lg shadow-xl">
        <h1 className="text-xl font-bold mb-4">アカウントの登録が必要です</h1>
        <div className="text-base mb-6 w-full">
          <p className="mb-4">
            「{props.kind}」を行うには、このウォレットアドレスに紐づくクリエイターアカウントの登録が必要です。
          </p>
          <div className="mb-4 p-3 bg-slate-100 rounded-md">
            <p className="text-sm font-semibold text-slate-600">ウォレットアドレス:</p>
            <p className="text-sm break-all">{props.wallet}</p>
          </div>
          <p>
            アカウントを作成しますか？
            <br />
            「アカウント作成」ボタンを押すと、登録ページに移動します。
          </p>
        </div>
        <div className="flex mt-auto w-full justify-end gap-4">
          <button className="bg-slate-500 hover:bg-slate-600 text-white px-6 py-2 rounded-md" onClick={props.onCancel}>
            キャンセル
          </button>
          <button
            className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-md"
            onClick={props.createAccount}
          >
            アカウント作成
          </button>
        </div>
      </div>
      <div className="fixed bg-black bg-opacity-50 w-full h-full z-10" onClick={props.onCancel}></div>
    </>
  );
};

export default Modal;
