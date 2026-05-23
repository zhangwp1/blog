import { useEffect, type ReactNode } from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { useAuthStore } from '../store/authSlice';
import { Spin } from 'antd';

export default function AuthGuard({ children }: { children: ReactNode }) {
  const { token, user, fetchProfile } = useAuthStore();
  const location = useLocation();

  useEffect(() => {
    if (token && !user) {
      fetchProfile();
    }
  }, [token, user, fetchProfile]);

  if (!token) {
    return <Navigate to="/admin/login" state={{ from: location }} replace />;
  }

  if (!user) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Spin size="large" />
      </div>
    );
  }

  return <>{children}</>;
}
