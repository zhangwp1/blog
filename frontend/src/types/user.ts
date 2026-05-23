export interface User {
  id: number;
  username: string;
  nickname: string;
  avatar: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  expires_in: number;
}
