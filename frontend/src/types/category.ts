export interface Category {
  id: number;
  name: string;
  slug: string;
  sort_order: number;
  created_at: string;
  updated_at: string;
}

export interface CreateCategoryRequest {
  name: string;
  slug: string;
  sort_order?: number;
}

export interface UpdateCategoryRequest {
  name?: string;
  slug?: string;
  sort_order?: number;
}
