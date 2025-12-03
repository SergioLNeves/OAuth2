import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

export function getContext() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 1000 * 60 * 5, // 5 minutes
        refetchOnWindowFocus: false,
        retry: false,
      },
    },
  });

  return {
    queryClient,
  };
}

interface TanStackQueryProviderProps {
  children: React.ReactNode;
  context: ReturnType<typeof getContext>;
}

export function Provider({ children, context }: TanStackQueryProviderProps) {
  return (
    <QueryClientProvider client={context.queryClient}>
      {children}
    </QueryClientProvider>
  );
}
