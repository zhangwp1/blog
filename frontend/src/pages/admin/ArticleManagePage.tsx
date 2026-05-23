import { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Table, Button, Space, Tag, Popconfirm, message, Input } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons';
import { articleApi } from '../../api/article';
import type { ArticleListItem } from '../../types/article';
import { formatDateTime } from '../../utils/format';

export default function ArticleManagePage() {
  const [articles, setArticles] = useState<ArticleListItem[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [keyword, setKeyword] = useState('');
  const navigate = useNavigate();
  const pageSize = 10;

  const fetchArticles = () => {
    setLoading(true);
    articleApi.listAdmin({ page, page_size: pageSize, keyword }).then((res) => {
      setArticles(res.data.data?.list || []);
      setTotal(res.data.data?.total || 0);
    }).finally(() => setLoading(false));
  };

  useEffect(() => { fetchArticles(); }, [page]);

  const handleDelete = async (id: number) => {
    try {
      await articleApi.delete(id);
      message.success('删除成功');
      fetchArticles();
    } catch {
      message.error('删除失败');
    }
  };

  const columns = [
    { title: '标题', dataIndex: 'title', key: 'title', render: (text: string, record: ArticleListItem) => <Link to={`/articles/${record.slug}`}>{text}</Link> },
    { title: '状态', dataIndex: 'is_published', key: 'is_published', width: 80, render: (v: boolean) => v ? <Tag color="green">已发布</Tag> : <Tag color="orange">草稿</Tag> },
    { title: '分类', dataIndex: ['category', 'name'], key: 'category', width: 100, render: (v: string) => v || '-' },
    { title: '浏览量', dataIndex: 'view_count', key: 'view_count', width: 80 },
    { title: '发布时间', dataIndex: 'published_at', key: 'published_at', width: 180, render: (v: string) => formatDateTime(v) },
    {
      title: '操作', key: 'action', width: 160,
      render: (_: unknown, record: ArticleListItem) => (
        <Space>
          <Button type="link" icon={<EditOutlined />} onClick={() => navigate(`/admin/articles/${record.id}/edit`)}>编辑</Button>
          <Popconfirm title="确定删除?" onConfirm={() => handleDelete(record.id)}>
            <Button type="link" danger icon={<DeleteOutlined />}>删除</Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 16 }}>
        <Input.Search placeholder="搜索文章..." value={keyword} onChange={(e) => setKeyword(e.target.value)} onSearch={fetchArticles} style={{ width: 300 }} prefix={<SearchOutlined />} />
        <Button type="primary" icon={<PlusOutlined />} onClick={() => navigate('/admin/articles/new')}>新建文章</Button>
      </div>
      <Table dataSource={articles} columns={columns} rowKey="id" loading={loading}
        pagination={{ current: page, total, pageSize, onChange: setPage }} />
    </div>
  );
}
