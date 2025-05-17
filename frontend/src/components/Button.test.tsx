import { fireEvent, render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import Button from "./Button";

describe("Button", () => {
  it("正しくレンダリングされること", () => {
    render(<Button>テストボタン</Button>);
    expect(screen.getByText("テストボタン")).toBeInTheDocument();
  });

  it("クリックイベントが発火すること", () => {
    const handleClick = vi.fn();
    render(<Button onClick={handleClick}>クリック</Button>);

    fireEvent.click(screen.getByText("クリック"));
    expect(handleClick).toHaveBeenCalledTimes(1);
  });

  it("無効化されている場合、クリックイベントが発火しないこと", () => {
    const handleClick = vi.fn();
    render(
      <Button onClick={handleClick} disabled>
        無効ボタン
      </Button>,
    );

    const button = screen.getByText("無効ボタン");
    expect(button).toBeDisabled();

    fireEvent.click(button);
    expect(handleClick).not.toHaveBeenCalled();
  });

  it("適切なバリアントスタイルが適用されること", () => {
    const { rerender } = render(<Button variant="primary">プライマリ</Button>);
    const primaryButton = screen.getByText("プライマリ");
    expect(primaryButton).toHaveStyle({
      backgroundColor: "#007bff",
    });

    rerender(<Button variant="secondary">セカンダリ</Button>);
    const secondaryButton = screen.getByText("セカンダリ");
    expect(secondaryButton).toHaveStyle({
      backgroundColor: "#6c757d",
    });

    rerender(<Button variant="danger">危険</Button>);
    const dangerButton = screen.getByText("危険");
    expect(dangerButton).toHaveStyle({
      backgroundColor: "#dc3545",
    });
  });

  it("デフォルトでボタンタイプであること", () => {
    render(<Button>デフォルト</Button>);
    const button = screen.getByText("デフォルト");
    expect(button).toHaveAttribute("type", "button");
  });

  it("タイプが指定されていれば、そのタイプが適用されること", () => {
    render(<Button type="submit">送信</Button>);
    const button = screen.getByText("送信");
    expect(button).toHaveAttribute("type", "submit");
  });
});
