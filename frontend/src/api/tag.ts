import apiClient from './client';
import type { Tag, CreateTagRequest, UpdateTagRequest } from '../types/tag';

export const tagApi = {
  list: () =>
    apiClient.get<{ code: number; data: Tag[] }>('/tags'),

  create: (data: CreateTagRequest) =>
    apiClient.post<{ code: number; data: Tag }>('/admin/tags', data),

  update: (id: number, data: UpdateTagRequest) =>
    apiClient.put<{ code: number; data: Tag }>(`/admin/tags/${id}`, data),

  delete: (id: number) =>
    apiClient.delete(`/admin/tags/${id}`),
};
