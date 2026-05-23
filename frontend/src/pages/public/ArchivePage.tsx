import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Card, Spin, Empty } from 'antd';
import { articleApi } from '../../api/article';
import type { ArticleListItem } from '../../types/article';
import { formatDate } from '../../utils/format';

export default function ArchivePage() {
  const [articles, setArticles] = useState<ArticleListItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    articleApi.listPublic({ page: 1, page_size: 1000 }).then((res) => {
      setArticles(res.data.data?.list || []);
    }).finally(() => setLoading(false));
  }, []);

  const grouped = articles.reduce<Record<string, ArticleListItem[]>>((acc, article) => {
    const key = article.published_at ? article.published_at.substring(0, 7) : '未发布';
    if (!acc[key]) acc[key] = [];
    acc[key].push(article);
    return acc;
  }, {});

  const sortedKeys = Object.keys(grouped).sort((a, b) => b.localeCompare(a));

  return (
    <Card title="文章归档" style={{ borderRadius: 8 }}>
      <Spin spinning={loading}>
        {sortedKeys.length === 0 && !loading ? (
          <Empty description="暂无文章" />
        ) : (
          sortedKeys.map((key) => (
            <div key={key} style={{ marginBottom: 24 }}>
              <h3 style={{ borderBottom: '1px solid #f0f0f0', paddingBottom: 8, color: '#1677ff' }}>{key}</h3>
              {grouped[key].map((article) => (
                <div key={article.id} style={{ padding: '6px 0', display: 'flex', justifyContent: 'space-between' }}>
                  <Link to={`/articles/${article.slug}`}>{article.title}</Link>
                  <span style={{ color: '#999', fontSize: 13 }}>{formatDate(article.published_at)}</span>
                </div>
              ))}
            </div>
          ))
        )}
      </Spin>
    </Card>
  );
}
