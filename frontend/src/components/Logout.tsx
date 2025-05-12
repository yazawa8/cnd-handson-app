import React, { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { logout } from '../features/sessions/slice';
import { useNavigate } from 'react-router-dom';

const Logout: React.FC = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  useEffect(() => {
    dispatch(logout());
    navigate('/login', { replace: true });
  }, [dispatch, navigate]);

  return null;
}

export default Logout;