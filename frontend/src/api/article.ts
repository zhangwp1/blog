import apiClient from './client';
import type {
  Article,
  ArticleListItem,
  CreateArticleRequest,
  UpdateArticleRequest,
  ArticleQuery,
} from '../types/article';

export const articleApi = {
  listPublic: (params: ArticleQuery) =>
    apiClient.get<{ code: number; data: { list: ArticleListItem[]; total: number; page: number; page_size: number } }>('/articles', { params }),

  listAdmin: (params: ArticleQuery) =>
    apiClient.get<{ code: number; data: { list: ArticleListItem[]; total: number; page: number; page_size: number } }>('/admin/articles', { params }),

  getBySlug: (slug: string) =>
    apiClient.get<{ code: number; data: Article }>(`/articles/${slug}`),

  getById: (id: number) =>
    apiClient.get<{ code: number; data: Article }>(`/admin/articles/${id}`),

  create: (data: CreateArticleRequest) =>
    apiClient.post<{ code: number; data: Article }>('/admin/articles', data),

  update: (id: number, data: UpdateArticleRequest) =>
    apiClient.put<{ code: number; data: Article }>(`/admin/articles/${id}`, data),

  delete: (id: number) =>
    apiClient.delete(`/admin/articles/${id}`),

  dashboard: () =>
    apiClient.get<{ code: number; data: Record<string, number> }>('/admin/dashboard'),
};
