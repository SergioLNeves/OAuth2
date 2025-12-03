import { exchangeCodeForTokens } from "@/api/oauth";
import { useAuth } from "@/hooks/use-auth";
import { AlertCircleIcon, LoaderCircleIcon } from "lucide-react";
import type { TokenResponse, User } from "@/types/oauth";
import { useMutation } from "@tanstack/react-query";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { z } from "zod";
import { useEffect } from "react";
import { ParseJWT } from "@/lib/utils";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";

export const Route = createFileRoute("/callback/")({
  validateSearch: (search) => searchSchema.parse(search),
  component: RouteCallback,
});

const searchSchema = z.object({
  code: z.string().optional(),
  state: z.string().optional(),
  error: z.string().optional(),
  error_description: z.string().optional(),
});

function RouteCallback() {
  const searchParams = Route.useSearch();
  const navigate = useNavigate();
  const { login } = useAuth();

  const { mutate, isPending, isError, error } = useMutation({
    mutationFn: exchangeCodeForTokens,
    retry: 0,
    onSuccess: (data: TokenResponse) => {
      const userData = ParseJWT<User>(data.id_token);
      login(data, userData);

      localStorage.removeItem("code_verifier");
      localStorage.removeItem("oauth_state");

      navigate({ to: "/dashboard" });
    },
  });

  useEffect(() => {
    const { code, state, error: apiError, error_description } = searchParams;

    if (apiError) {
      console.error("OAuth error: ", error_description);
      navigate({
        to: "/login",
        search: { error: apiError },
      });
      return;
    }

    const savedState = localStorage.getItem("oauth_state");
    if (state !== savedState) {
      navigate({
        to: "/login",
        search: { error: "invalid_state" },
      });
      return;
    }

    const codeVerifier = localStorage.getItem("code_verifier");
    if (!code || !codeVerifier) {
      navigate({
        to: "/login",
        search: { error: "missing_params" },
      });
      return;
    }

    mutate({ code, codeVerifier });
  }, [searchParams, mutate, navigate]);

  if (isPending) {
    return (
      <div className="flex min-h-screen w-full items-center justify-center bg-background p-4">
        <div className="flex flex-col items-center space-y-4">
          <LoaderCircleIcon className="h-10 w-10 animate-spin text-primary" />
          <p className="text-muted-foreground text-sm animate-pulse">
            Validando credenciais...
          </p>
        </div>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="flex min-h-screen w-full items-center justify-center bg-background p-4">
        <Card className="w-full max-w-md border-destructive/50">
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-destructive">
              <AlertCircleIcon className="h-5 w-5" />
              Falha na Autenticação
            </CardTitle>
            <CardDescription>
              Não foi possível concluir o login.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <Alert variant="destructive">
              <AlertTitle>Erro</AlertTitle>
              <AlertDescription>
                {error instanceof Error ? error.message : "Erro desconhecido"}
              </AlertDescription>
            </Alert>

            <Button
              variant="outline"
              className="w-full"
              onClick={() => navigate({ to: "/login" })}
            >
              Voltar para o Login
            </Button>
          </CardContent>
        </Card>
      </div>
    );
  }

  return null;
}
