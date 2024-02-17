
import type { UseFetchOptions } from "#app";
import { useApi } from "./useApi";

export function usePostApi<T>(url: string, options: UseFetchOptions<T> = {}) {
    return useApi<T>(url, { ...options, method: "POST"});
  }