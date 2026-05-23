import { useEffect, useState } from 'react';
import { Card, Row, Col, Statistic } from 'antd';
import { FileTextOutlined, CommentOutlined, CheckCircleOutlined, ClockCircleOutlined } from '@ant-design/icons';
import { articleApi } from '../../api/article';
import { commentApi } from '../../api/comment';

export default function DashboardPage() {
  const [stats, setStats] = useState<Record<string, number>>({});
  const [commentStats, setCommentStats] = useState<Record<string, number>>({});

  useEffect(() => {
    articleApi.dashboard().then((res) => setStats(res.data.data || {}));
    commentApi.listAdmin({ page: 1, page_size: 1 }).then((res) => {
      const total = res.data.data?.total || 0;
      setCommentStats({ total_comments: total });
    });
    commentApi.listAdmin({ page: 1, page_size: 1, is_approved: 0 }).then((res) => {
      setCommentStats((prev) => ({ ...prev, pending_comments: res.data.data?.total || 0 }));
    });
  }, []);

  return (
    <div>
      <h2 style={{ marginBottom: 24 }}>仪表盘</h2>
      <Row gutter={[16, 16]}>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic title="文章总数" value={stats.total_articles || 0} prefix={<FileTextOutlined />} />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic title="已发布" value={stats.published_articles || 0} prefix={<CheckCircleOutlined />} valueStyle={{ color: '#3f8600' }} />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic title="草稿" value={stats.draft_articles || 0} prefix={<ClockCircleOutlined />} valueStyle={{ color: '#faad14' }} />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic title="待审核评论" value={commentStats.pending_comments || 0} prefix={<CommentOutlined />} valueStyle={{ color: '#ff4d4f' }} />
          </Card>
        </Col>
      </Row>
    </div>
  );
}
