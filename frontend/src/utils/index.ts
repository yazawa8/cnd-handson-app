/**
 * 文字列が空かどうかをチェックする
 * @param value 検証する文字列
 * @returns 空文字列またはnull/undefinedの場合はtrue、それ以外はfalse
 */
export const isEmpty = (value: string | null | undefined): boolean => {
  return value === null || value === undefined || value.trim() === "";
};

/**
 * 2つの数値を足し算する
 * @param a 1つ目の数値
 * @param b 2つ目の数値
 * @returns 足し算の結果
 */
export const add = (a: number, b: number): number => {
  return a + b;
};

/**
 * 文字列の先頭を大文字にする
 * @param str 変換する文字列
 * @returns 先頭が大文字の文字列
 */
export const capitalize = (str: string): string => {
  if (isEmpty(str)) return "";
  return str.charAt(0).toUpperCase() + str.slice(1);
};
