let getToken: () => Promise<string> = async () => "";

/**
 * アクセストークンを取得する関数を設定します。
 * @param fn トークンを取得する非同期関数
 */
export const setGetToken = (fn: () => Promise<string>) => {
  getToken = fn;
}

/**
 * アクセストークンを取得する関数を呼び出します。
 * @returns アクセストークン
 */
export const getAccessToken = () => getToken();
