export interface AuthTokens {
  access_token: string;
  refresh_token: string;
}

export interface User {
  name?: string;
  email?: string;
  [key: string]: unknown;
}

export interface TokenResponse {
  id_token: string;
  access_token: string;
  refresh_token: string;
  expires_in: number;
}
