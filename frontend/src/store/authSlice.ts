import { create } from 'zustand';
import type { User } from '../types/user';
import { authApi } from '../api/auth';

interface AuthState {
  token: string | null;
  user: User | null;
  setToken: (token: string) => void;
  setUser: (user: User) => void;
  fetchProfile: () => Promise<void>;
  logout: () => void;
  isAuthenticated: () => boolean;
}

export const useAuthStore = create<AuthState>((set, get) => ({
  token: localStorage.getItem('token'),
  user: JSON.parse(localStorage.getItem('user') || 'null'),

  setToken: (token: string) => {
    localStorage.setItem('token', token);
    set({ token });
  },

  setUser: (user: User) => {
    localStorage.setItem('user', JSON.stringify(user));
    set({ user });
  },

  fetchProfile: async () => {
    try {
      const res = await authApi.profile();
      get().setUser(res.data.data);
    } catch {
      get().logout();
    }
  },

  logout: () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    set({ token: null, user: null });
  },

  isAuthenticated: () => !!get().token,
}));
