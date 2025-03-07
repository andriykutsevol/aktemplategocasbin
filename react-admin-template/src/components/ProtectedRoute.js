import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export const ProtectedRoute = ({ children }) => {
  const { user } = useAuth();
    if (!user) {
    //if (false) {
    // user is not authenticated
    return <Navigate to="/login" />;
  }
  return children;
};