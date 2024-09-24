// request.ts file
/* eslint-disable */
// @ts-nocheck

import axios from "axios";
import type { RawAxiosRequestHeaders } from "axios";

import type { ApiRequestOptions } from "./ApiRequestOptions";
import { CancelablePromise } from "./CancelablePromise";
import type { OpenAPIConfig } from "./OpenAPI";
import { getAccessToken } from "../../../lib/tokenProvider"; // トークンプロバイダーからアクセストークン取得関数をインポート

// Optional: Get and link the cancelation token, so the request can be aborted.
const source = axios.CancelToken.source();

const axiosInstance = axios.create({
  // Your custom Axios instance config
  baseURL: "http://localhost:9090/api/v1/",
  headers: {
    "Content-Type": "application/json",
  } satisfies RawAxiosRequestHeaders,
});

// Add a request interceptor
axiosInstance.interceptors.request.use(
  async function (config) {
    // アクセストークンを取得
    const token = await getAccessToken();
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }

    // URL 内のパラメータをエンコードして置換
    if (config.url && config.params) {
      Object.entries<any>(config.params).forEach(([key, value]) => {
        const stringToSearch = `{${key}}`;
        if (config.url.search(stringToSearch) !== -1) {
          config.url = config.url.replace(
            `{${key}}`,
            encodeURIComponent(String(value))
          );
          delete config.params[key];
        }
      });
    }

    return config;
  },
  function (error) {
    // リクエストエラーを処理
    return Promise.reject(error);
  }
);

// Add a response interceptor
axiosInstance.interceptors.response.use(
  function (response) {
    // 2xx ステータスコードのレスポンスを処理
    return response;
  },
  function (error) {
    // 2xx 以外のステータスコードのレスポンスを処理
    return Promise.reject(error);
  }
);

/**
 * カスタムリクエスト関数
 * @param config OpenAPI 設定
 * @param options API リクエストオプション
 * @returns CancelablePromise
 */
export const request = <T>(
  config: OpenAPIConfig,
  options: ApiRequestOptions
): CancelablePromise<T> => {
  return new CancelablePromise((resolve, reject, onCancel) => {
    onCancel(() => source.cancel("The user aborted a request."));

    let formattedHeaders = options.headers as RawAxiosRequestHeaders;
    if (options.mediaType) {
      formattedHeaders = {
        ...options.headers,
        "Content-Type": options.mediaType,
      } satisfies RawAxiosRequestHeaders;
    }

    return axiosInstance
      .request({
        url: options.url,
        data: options.body,
        method: options.method,
        params: options.path,
        headers: formattedHeaders,
        cancelToken: source.token,
      })
      .then((res) => {
        resolve(res.data);
      })
      .catch((error) => {
        reject(error);
      });
  });
};
