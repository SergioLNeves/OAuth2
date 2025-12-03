export const OAUTH_CONFIG = {
  clientId: import.meta.env.VITE_OAUTH_CLIENT_ID,
  authUrl: import.meta.env.VITE_OAUTH_AUTH_URL,
  tokenUrl: import.meta.env.VITE_OAUTH_TOKEN_URL,
  redirectUri: import.meta.env.VITE_OAUTH_REDIRECT_URI,
  scope: "openid profile email",
  responseType: "code",
};
