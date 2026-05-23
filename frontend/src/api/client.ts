import axios from 'axios';
import { message } from 'antd';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  timeout: 10000,
  headers: { 'Content-Type': 'application/json' },
});

apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

apiClient.interceptors.response.use(
  (response) => {
    const data = response.data;
    if (data.code !== 0) {
      message.error(data.message || '请求失败');
      return Promise.reject(new Error(data.message));
    }
    return response;
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/admin/login';
    }
    message.error(error.message || '网络错误');
    return Promise.reject(error);
  }
);

export default apiClient;
