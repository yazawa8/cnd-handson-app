import type React from "react";
import { useSelector } from "react-redux";
import { Navigate, Outlet } from "react-router-dom";
import type { RootState } from "../store";
import Header from "./Header";

const PrivateLayout: React.FC = () => {
  const isLoggedIn = useSelector(
    (state: RootState) => state.session.isLoggedIn,
  );

  if (!isLoggedIn) return <Navigate to="/login" replace />;

  return (
    <>
      <Header />
      <Outlet />
    </>
  );
};

export default PrivateLayout;
