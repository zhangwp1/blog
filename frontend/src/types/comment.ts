export interface Comment {
  id: number;
  article_id: number;
  parent_id: number;
  author_name: string;
  author_email?: string;
  author_website?: string;
  content: string;
  is_approved: number;
  is_admin: boolean;
  created_at: string;
  children?: Comment[];
}

export interface CreateCommentRequest {
  article_id: number;
  parent_id?: number;
  author_name: string;
  author_email?: string;
  author_website?: string;
  content: string;
}

export interface CommentQuery {
  page?: number;
  page_size?: number;
  is_approved?: number;
}
