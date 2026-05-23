export interface Article {
  id: number;
  title: string;
  slug: string;
  content?: string;
  summary: string;
  cover_image: string;
  is_published: boolean;
  pinned: boolean;
  view_count: number;
  category_id: number;
  author_id: number;
  published_at: string | null;
  created_at: string;
  updated_at: string;
  category?: CategoryInfo;
  author?: AuthorInfo;
  tags?: TagInfo[];
}

export interface ArticleListItem {
  id: number;
  title: string;
  slug: string;
  summary: string;
  cover_image: string;
  is_published: boolean;
  pinned: boolean;
  view_count: number;
  category_id: number;
  published_at: string | null;
  created_at: string;
  category?: CategoryInfo;
  tags?: TagInfo[];
}

export interface CategoryInfo {
  id: number;
  name: string;
  slug: string;
}

export interface AuthorInfo {
  id: number;
  username: string;
  nickname: string;
  avatar: string;
}

export interface TagInfo {
  id: number;
  name: string;
  slug: string;
}

export interface CreateArticleRequest {
  title: string;
  content: string;
  summary?: string;
  cover_image?: string;
  is_published?: boolean;
  pinned?: boolean;
  category_id?: number;
  tag_ids?: number[];
}

export interface UpdateArticleRequest {
  title?: string;
  content?: string;
  summary?: string;
  cover_image?: string;
  is_published?: boolean;
  pinned?: boolean;
  category_id?: number;
  tag_ids?: number[];
}

export interface ArticleQuery {
  page?: number;
  page_size?: number;
  category?: string;
  tag?: string;
  keyword?: string;
  year?: number;
  month?: number;
}
