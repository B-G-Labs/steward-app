import { useCookie, useFetch, type UseFetchOptions, type FetchResult } from "#app";
import { defu } from "defu";

export function useApi<T>(url: string, options: UseFetchOptions<T> = {}) {
  const authToken = useCookie("token");
  
  const defaults: UseFetchOptions<T> = {
    baseURL: "http://127.0.0.1:3000/api",
    // set user token if connected
    headers: authToken.value
      ? { Authorization: `Bearer ${authToken.value}` }
      : {},
  };

  // for nice deep defaults, please use unjs/defu
  const params = defu(options, defaults);

  return useFetch(url, params);
}





