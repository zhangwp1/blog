export interface Tag {
  id: number;
  name: string;
  slug: string;
  created_at: string;
  updated_at: string;
}

export interface CreateTagRequest {
  name: string;
  slug: string;
}

export interface UpdateTagRequest {
  name?: string;
  slug?: string;
}
