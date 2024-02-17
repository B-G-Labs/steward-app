import type { UseFetchOptions } from "#app";
import { useApi } from "./useApi";

export function usePutApi<T>(url: string, options: UseFetchOptions<T> = {}) {
    return useApi<T>(url, { method: "PUT", ...options });
  }