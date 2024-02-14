import { useCookie } from "#app";
import { useApi } from "./useApi";

export interface UserData {
  name: string;
  password: string;
}

export function useAuth() {
  // TODO: implement isAuthenticated method
  // TODO: add token expiry date
  // TODO: redirect user to log-in when not authenticated  
  const token = useCookie("token");

  async function register(userData: UserData) {
    const { status, error } = await useApi("/auth/register", {
      method: "POST",
      body: userData,
    });

    if (error) throw new Error(error.value?.message);

    if (status) return true;

    return false;
  }

  async function logIn(userData: UserData) {
    const { data, error } = await useApi<any>("/auth/login", {
      method: "POST",
      body: userData,
    });

    token.value = data.value.data;
  }

  return {
    register,
    logIn,
  };
}
