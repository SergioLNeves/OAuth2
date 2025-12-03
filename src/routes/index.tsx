import { LoginForm } from "@/components/templates/login-form";
import { OAUTH_CONFIG } from "@/config/oauth";
import {
  generateCodeChallenge,
  generateCodeVerifier,
  generateState,
} from "@/lib/pkce";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: RouteLogin,
});

function RouteLogin() {
  async function handleLogin() {
    try {
      const codeVerifier = generateCodeVerifier();
      const state = generateState();
      const codeChallenge = await generateCodeChallenge(codeVerifier);

      localStorage.setItem("code_verifier", codeVerifier);
      localStorage.setItem("oauth_state", state);

      const params = new URLSearchParams({
        client_id: OAUTH_CONFIG.clientId,
        redirect_uri: OAUTH_CONFIG.redirectUri,
        response_type: OAUTH_CONFIG.responseType,
        scope: OAUTH_CONFIG.scope,
        state: state,
        code_challenge: codeChallenge,
        code_challenge_method: "S256",
      });

      const authUrl = `${OAUTH_CONFIG.authUrl}?${params.toString()}`;
      window.location.href = authUrl;
    } catch (error) {
      console.error("Init login failed, err: ", error);
    }
  }

  return (
    <div className="bg-background flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm onClick={handleLogin} />
      </div>
    </div>
  );
}
