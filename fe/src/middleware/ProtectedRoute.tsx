import { Navigate } from "react-router";
import { useAuth } from "../context/AuthContext.tsx";
import { ReactNode } from "react";

export default function ProtectedRoute({
                                           children,
                                       }: {
    children: ReactNode;
}) {
    const { user } = useAuth();

    if (!user) return <Navigate to="/signin" replace />;

    return children;
}
