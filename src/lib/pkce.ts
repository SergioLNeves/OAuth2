// Converte array para base64url
function base64UrlEncode(buffer: ArrayBuffer | Uint8Array) {
  const bytes = buffer instanceof Uint8Array ? buffer : new Uint8Array(buffer);
  const base64 = btoa(String.fromCharCode(...bytes));
  return base64.replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
}

// Gera SHA256 hash
async function sha256(plain: string) {
  const encoder = new TextEncoder();
  const data = encoder.encode(plain);
  const hash = await crypto.subtle.digest("SHA-256", data);
  return hash;
}

// Gera string aleatória para code_verifier
export function generateRandomString(length: number) {
  const array = new Uint8Array(length);
  crypto.getRandomValues(array);
  return base64UrlEncode(array);
}

// Gera code_challenge a partir do code_verifier
export async function generateCodeChallenge(codeVerifier: string) {
  const hashed = await sha256(codeVerifier);
  return base64UrlEncode(hashed);
}

export function generateCodeVerifier() {
  return generateRandomString(43);
}

export function generateState() {
  return generateRandomString(32);
}
