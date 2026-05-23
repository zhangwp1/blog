import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Card, Spin, Tag, Empty, message } from 'antd';
import { CalendarOutlined, EyeOutlined, FolderOutlined } from '@ant-design/icons';
import { articleApi } from '../../api/article';
import { commentApi } from '../../api/comment';
import type { Article } from '../../types/article';
import type { Comment } from '../../types/comment';
import { formatDate } from '../../utils/format';
import MdRenderer from '../../components/markdown/MdRenderer';
import CommentList from '../../components/comment/CommentList';
import CommentForm from '../../components/comment/CommentForm';

export default function ArticleDetailPage() {
  const { slug } = useParams<{ slug: string }>();
  const [article, setArticle] = useState<Article | null>(null);
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchArticle = async () => {
    if (!slug) return;
    setLoading(true);
    try {
      const res = await articleApi.getBySlug(slug);
      setArticle(res.data.data);
    } catch {
      message.error('文章加载失败');
    } finally {
      setLoading(false);
    }
  };

  const fetchComments = async () => {
    if (!slug) return;
    try {
      const res = await commentApi.listByArticle(slug);
      setComments(res.data.data || []);
    } catch {
      // silently fail
    }
  };

  useEffect(() => {
    fetchArticle();
    fetchComments();
  }, [slug]);

  if (loading) return <Spin style={{ display: 'block', margin: '100px auto' }} size="large" />;
  if (!article) return <Empty description="文章不存在" />;

  const handleCommentSubmit = async (data: { author_name: string; author_email: string; content: string }) => {
    await commentApi.create(article.slug, {
      article_id: article.id,
      author_name: data.author_name,
      author_email: data.author_email,
      content: data.content,
    });
    await fetchComments();
  };

  return (
    <div>
      <Card style={{ borderRadius: 8, marginBottom: 24 }}>
        <h1 style={{ marginBottom: 12, fontSize: 28 }}>{article.title}</h1>
        <div style={{ display: 'flex', gap: 16, color: '#999', fontSize: 13, marginBottom: 20, flexWrap: 'wrap' }}>
          <span><CalendarOutlined style={{ marginRight: 4 }} />{formatDate(article.published_at)}</span>
          <span><EyeOutlined style={{ marginRight: 4 }} />{article.view_count}</span>
          {article.category && (
            <span><FolderOutlined style={{ marginRight: 4 }} />{article.category.name}</span>
          )}
        </div>
        {article.tags && article.tags.length > 0 && (
          <div style={{ marginBottom: 20 }}>
            {article.tags.map((tag) => <Tag key={tag.id}>{tag.name}</Tag>)}
          </div>
        )}
        {article.cover_image && (
          <img src={article.cover_image} alt={article.title} style={{ width: '100%', borderRadius: 8, marginBottom: 20, maxHeight: 400, objectFit: 'cover' }} />
        )}
        <MdRenderer content={article.content || ''} />
      </Card>

      <Card title={`评论 (${comments.length})`} style={{ borderRadius: 8 }}>
        <CommentList comments={comments} onReply={() => {}} />
        <CommentForm onSubmit={handleCommentSubmit} />
      </Card>
    </div>
  );
}
