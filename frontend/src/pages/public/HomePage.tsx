import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Card, List, Tag, Pagination, Spin, Empty } from 'antd';
import { CalendarOutlined, EyeOutlined, FolderOutlined } from '@ant-design/icons';
import { articleApi } from '../../api/article';
import { categoryApi } from '../../api/category';
import { tagApi } from '../../api/tag';
import type { ArticleListItem } from '../../types/article';
import type { Category } from '../../types/category';
import type { Tag as TagType } from '../../types/tag';
import { formatDate } from '../../utils/format';

export default function HomePage() {
  const [articles, setArticles] = useState<ArticleListItem[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [categories, setCategories] = useState<Category[]>([]);
  const [tags, setTags] = useState<TagType[]>([]);
  const pageSize = 10;

  useEffect(() => {
    categoryApi.list().then((res) => setCategories(res.data.data || []));
    tagApi.list().then((res) => setTags(res.data.data || []));
  }, []);

  useEffect(() => {
    setLoading(true);
    articleApi.listPublic({ page, page_size: pageSize }).then((res) => {
      setArticles(res.data.data?.list || []);
      setTotal(res.data.data?.total || 0);
    }).finally(() => setLoading(false));
  }, [page]);

  return (
    <div>
      <Spin spinning={loading}>
        {articles.length === 0 && !loading ? (
          <Empty description="暂无文章" />
        ) : (
          <List
            dataSource={articles}
            renderItem={(item) => (
              <Card
                hoverable
                style={{ marginBottom: 16, borderRadius: 8 }}
                bodyStyle={{ padding: '20px 24px' }}
                key={item.id}
              >
                {item.pinned && <Tag color="red" style={{ marginBottom: 8 }}>置顶</Tag>}
                <h2 style={{ margin: 0, marginBottom: 8 }}>
                  <Link to={`/articles/${item.slug}`} style={{ color: '#222' }}>{item.title}</Link>
                </h2>
                <p style={{ color: '#666', marginBottom: 12, lineHeight: 1.6 }}>{item.summary}</p>
                <div style={{ display: 'flex', alignItems: 'center', gap: 16, color: '#999', fontSize: 13, flexWrap: 'wrap' }}>
                  <span><CalendarOutlined style={{ marginRight: 4 }} />{formatDate(item.published_at)}</span>
                  <span><EyeOutlined style={{ marginRight: 4 }} />{item.view_count}</span>
                  {item.category && (
                    <Link to={`/?category=${item.category.slug}`}>
                      <span><FolderOutlined style={{ marginRight: 4 }} />{item.category.name}</span>
                    </Link>
                  )}
                  {item.tags?.map((tag) => (
                    <Link to={`/?tag=${tag.slug}`} key={tag.id}>
                      <Tag>{tag.name}</Tag>
                    </Link>
                  ))}
                </div>
              </Card>
            )}
          />
        )}
      </Spin>
      {total > pageSize && (
        <div style={{ textAlign: 'center', marginTop: 24 }}>
          <Pagination current={page} total={total} pageSize={pageSize} onChange={setPage} />
        </div>
      )}

      {/* Sidebar content - rendered inside the sidebar area by PublicLayout */}
      <div className="sidebar-widgets" style={{ position: 'sticky', top: 80 }}>
        <Card title="分类" size="small" style={{ marginBottom: 16 }}>
          {categories.map((cat) => (
            <Link to={`/?category=${cat.slug}`} key={cat.id} style={{ display: 'block', padding: '4px 0', color: '#333' }}>
              {cat.name}
            </Link>
          ))}
          {categories.length === 0 && <span style={{ color: '#999' }}>暂无分类</span>}
        </Card>
        <Card title="标签" size="small">
          {tags.map((tag) => (
            <Link to={`/?tag=${tag.slug}`} key={tag.id}>
              <Tag style={{ marginBottom: 8 }}>{tag.name}</Tag>
            </Link>
          ))}
          {tags.length === 0 && <span style={{ color: '#999' }}>暂无标签</span>}
        </Card>
      </div>
    </div>
  );
}
