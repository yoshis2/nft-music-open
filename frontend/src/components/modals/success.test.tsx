import { describe, it, expect, vi } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import SuccessModal from "./success";
import React from "react";

describe("SuccessModal Component", () => {
  const defaultProps = {
    open: true,
    kind: "NFTミント",
    wallet: "0x1234567890abcdef",
    onCancel: vi.fn(),
    onHandler: vi.fn(),
  };

  it("オープン状態の時に正しく表示されること", () => {
    render(<SuccessModal {...defaultProps} />);
    expect(screen.getByText("NFTミント 確認画面")).toBeInTheDocument();
    expect(screen.getByText("0x1234567890abcdef")).toBeInTheDocument();
  });

  it("クローズ状態の時に表示されないこと", () => {
    render(<SuccessModal {...defaultProps} open={false} />);
    expect(screen.queryByText("NFTミント 確認画面")).not.toBeInTheDocument();
  });

  it("ミントボタンをクリックすると onHandler が呼ばれること", () => {
    render(<SuccessModal {...defaultProps} />);
    const mintButton = screen.getByRole("button", { name: "ミント" });
    fireEvent.click(mintButton);
    expect(defaultProps.onHandler).toHaveBeenCalled();
  });

  it("キャンセルボタンをクリックすると onCancel が呼ばれること", () => {
    render(<SuccessModal {...defaultProps} />);
    const cancelButton = screen.getByRole("button", { name: "キャンセル" });
    fireEvent.click(cancelButton);
    expect(defaultProps.onCancel).toHaveBeenCalled();
  });
});
