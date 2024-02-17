import type { UseFetchOptions } from "#app";
import { useApi } from "./useApi";

export function useDeleteApi<T>(url: string, options: UseFetchOptions<T> = {}) {
    return useApi<T>(url, { method: "DELETE", ...options });
  }