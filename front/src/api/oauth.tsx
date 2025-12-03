import { OAUTH_CONFIG } from "@/config/oauth";

type exchangeCodeForTokensType = {
  code: string;
  codeVerifier: string;
};

export async function exchangeCodeForTokens({
  code,
  codeVerifier,
}: exchangeCodeForTokensType) {
  const response = await fetch(OAUTH_CONFIG.tokenUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      grant_type: "authorization_code",
      code,
      code_verifier: codeVerifier,
      client_id: OAUTH_CONFIG.clientId,
      redirect_uri: OAUTH_CONFIG.redirectUri,
    }),
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error_description || "Failed to exchange code");
  }

  return response.json();
}
