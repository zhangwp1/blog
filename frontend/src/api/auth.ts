import apiClient from './client';
import type { LoginRequest, LoginResponse, User } from '../types/user';

export const authApi = {
  login: (data: LoginRequest) =>
    apiClient.post<{ code: number; data: LoginResponse }>('/auth/login', data),

  profile: () =>
    apiClient.get<{ code: number; data: User }>('/auth/profile'),
};
