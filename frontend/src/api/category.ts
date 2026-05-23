import apiClient from './client';
import type { Category, CreateCategoryRequest, UpdateCategoryRequest } from '../types/category';

export const categoryApi = {
  list: () =>
    apiClient.get<{ code: number; data: Category[] }>('/categories'),

  create: (data: CreateCategoryRequest) =>
    apiClient.post<{ code: number; data: Category }>('/admin/categories', data),

  update: (id: number, data: UpdateCategoryRequest) =>
    apiClient.put<{ code: number; data: Category }>(`/admin/categories/${id}`, data),

  delete: (id: number) =>
    apiClient.delete(`/admin/categories/${id}`),
};
