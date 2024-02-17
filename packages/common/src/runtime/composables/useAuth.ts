import { useCookie } from "#app";
import { usePostApi } from "./usePostApi";

export interface UserData {
  name: string;
  password: string;
}

export function useAuth() {
  // TODO: implement isAuthenticated method
  // TODO: add token expiry date
  // TODO: redirect user to log-in when not authenticated
  const token = useCookie("token");
  const tokenExpiryDate = useCookie<Number>("tokenExpiryDate");

  async function register(userData: UserData) {
    const { status, error, data } = await usePostApi("/auth/register", {
      body: userData,
    });

    if (error) throw new Error(error.value?.message);

    if (status) return true;

    return data.value;
  }

  async function logIn(userData: UserData) {
    const { data, error } = await usePostApi<any>("/auth/login", {
      body: userData,
    });

    token.value = data.value.data.token;
    tokenExpiryDate.value = data.value.data.expiry_data;
  }

  function isAuthenticated() {
    if (!tokenExpiryDate.value || token.value) return false;

    const expiryDate = new Date(+tokenExpiryDate.value * 1000);
    const tokenIsValid = token.value && expiryDate > new Date();

    return tokenIsValid;
  }

  return {
    register,
    logIn,
    isAuthenticated,
  };
}
