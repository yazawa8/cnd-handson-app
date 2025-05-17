import { describe, expect, it } from "vitest";
import { add, capitalize, isEmpty } from "./index";

describe("isEmpty", () => {
  it("nullの場合はtrueを返す", () => {
    expect(isEmpty(null)).toBe(true);
  });

  it("undefinedの場合はtrueを返す", () => {
    expect(isEmpty(undefined)).toBe(true);
  });

  it("空文字列の場合はtrueを返す", () => {
    expect(isEmpty("")).toBe(true);
  });

  it("空白のみの文字列の場合はtrueを返す", () => {
    expect(isEmpty("   ")).toBe(true);
  });

  it("文字列が含まれる場合はfalseを返す", () => {
    expect(isEmpty("hello")).toBe(false);
  });
});

describe("add", () => {
  it("2つの正の数を足し算できる", () => {
    expect(add(1, 2)).toBe(3);
  });

  it("負の数を処理できる", () => {
    expect(add(-1, -2)).toBe(-3);
    expect(add(-1, 2)).toBe(1);
  });

  it("小数を処理できる", () => {
    expect(add(1.5, 2.5)).toBe(4);
  });
});

describe("capitalize", () => {
  it("文字列の先頭を大文字に変換する", () => {
    expect(capitalize("hello")).toBe("Hello");
  });

  it("すでに大文字で始まる文字列はそのままにする", () => {
    expect(capitalize("Hello")).toBe("Hello");
  });

  it("空文字列は空文字列を返す", () => {
    expect(capitalize("")).toBe("");
  });

  it("1文字の文字列を正しく処理する", () => {
    expect(capitalize("a")).toBe("A");
  });
});
