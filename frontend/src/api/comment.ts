import apiClient from './client';
import type { Comment, CreateCommentRequest, CommentQuery } from '../types/comment';

export const commentApi = {
  listByArticle: (slug: string) =>
    apiClient.get<{ code: number; data: Comment[] }>(`/articles/${slug}/comments`),

  create: (slug: string, data: CreateCommentRequest) =>
    apiClient.post<{ code: number; data: Comment }>(`/articles/${slug}/comments`, data),

  listAdmin: (params: CommentQuery) =>
    apiClient.get<{ code: number; data: { list: Comment[]; total: number; page: number; page_size: number } }>('/admin/comments', { params }),

  approve: (id: number) =>
    apiClient.patch(`/admin/comments/${id}/approve`),

  reject: (id: number) =>
    apiClient.patch(`/admin/comments/${id}/reject`),

  delete: (id: number) =>
    apiClient.delete(`/admin/comments/${id}`),

  reply: (id: number, content: string) =>
    apiClient.post<{ code: number; data: Comment }>(`/admin/comments/${id}/reply`, { content }),
};
