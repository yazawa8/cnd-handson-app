import type React from "react";
import { Outlet, Navigate } from "react-router-dom";
import { useSelector } from "react-redux";
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
