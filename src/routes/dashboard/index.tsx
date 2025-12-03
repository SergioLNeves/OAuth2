import { useAuth } from "@/hooks/use-auth";
import { createFileRoute, useNavigate } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/")({
  component: RouteDashboard,
});

function RouteDashboard() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  function handleLogout() {
    logout();
    navigate({ to: "/" });
  }

  return (
    <div className="flex flex-col items-center justify-center p-6">
      <h1>Dashboard</h1>
      <p>Welcome, {user?.name || user?.email || "User"}!</p>
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
}
