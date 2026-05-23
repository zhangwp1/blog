import { useEffect, useState } from 'react';
import { Table, Button, Space, Tag, Popconfirm, message, Modal, Input, Select } from 'antd';
import { CheckOutlined, CloseOutlined, DeleteOutlined, CommentOutlined } from '@ant-design/icons';
import { commentApi } from '../../api/comment';
import type { Comment } from '../../types/comment';
import { formatDateTime } from '../../utils/format';

export default function CommentManagePage() {
  const [comments, setComments] = useState<Comment[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [statusFilter, setStatusFilter] = useState<number | undefined>(undefined);
  const [replyModal, setReplyModal] = useState<{ open: boolean; commentId: number | null }>({ open: false, commentId: null });
  const [replyContent, setReplyContent] = useState('');
  const pageSize = 10;

  const fetch = () => {
    setLoading(true);
    commentApi.listAdmin({ page, page_size: pageSize, is_approved: statusFilter }).then((res) => {
      setComments(res.data.data?.list || []);
      setTotal(res.data.data?.total || 0);
    }).finally(() => setLoading(false));
  };

  useEffect(() => { fetch(); }, [page, statusFilter]);

  const handleApprove = async (id: number) => {
    try { await commentApi.approve(id); message.success('已通过'); fetch(); } catch { message.error('操作失败'); }
  };

  const handleReject = async (id: number) => {
    try { await commentApi.reject(id); message.success('已拒绝'); fetch(); } catch { message.error('操作失败'); }
  };

  const handleDelete = async (id: number) => {
    try { await commentApi.delete(id); message.success('已删除'); fetch(); } catch { message.error('删除失败'); }
  };

  const handleReply = async () => {
    if (!replyModal.commentId || !replyContent) return;
    try {
      await commentApi.reply(replyModal.commentId, replyContent);
      message.success('回复成功');
      setReplyModal({ open: false, commentId: null });
      setReplyContent('');
      fetch();
    } catch {
      message.error('回复失败');
    }
  };

  const statusMap: Record<number, { color: string; text: string }> = {
    0: { color: 'orange', text: '待审核' },
    1: { color: 'green', text: '已通过' },
    2: { color: 'red', text: '已拒绝' },
  };

  const columns = [
    { title: '作者', dataIndex: 'author_name', key: 'author_name', width: 100, render: (v: string, r: Comment) => <span>{v}{r.is_admin && <Tag color="blue" style={{ marginLeft: 4 }}>博主</Tag>}</span> },
    { title: '内容', dataIndex: 'content', key: 'content', ellipsis: true, width: 300 },
    { title: '状态', dataIndex: 'is_approved', key: 'is_approved', width: 80, render: (v: number) => <Tag color={statusMap[v]?.color}>{statusMap[v]?.text}</Tag> },
    { title: '时间', dataIndex: 'created_at', key: 'created_at', width: 160, render: (v: string) => formatDateTime(v) },
    {
      title: '操作', key: 'action', width: 240,
      render: (_: unknown, record: Comment) => (
        <Space>
          {record.is_approved === 0 && (
            <>
              <Button type="link" size="small" icon={<CheckOutlined />} onClick={() => handleApprove(record.id)}>通过</Button>
              <Button type="link" size="small" danger icon={<CloseOutlined />} onClick={() => handleReject(record.id)}>拒绝</Button>
            </>
          )}
          <Button type="link" size="small" icon={<CommentOutlined />} onClick={() => setReplyModal({ open: true, commentId: record.id })}>回复</Button>
          <Popconfirm title="确定删除?" onConfirm={() => handleDelete(record.id)}>
            <Button type="link" size="small" danger icon={<DeleteOutlined />}>删除</Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Select
          placeholder="审核状态" allowClear style={{ width: 150 }}
          value={statusFilter} onChange={(v) => { setStatusFilter(v); setPage(1); }}
          options={[{ value: 0, label: '待审核' }, { value: 1, label: '已通过' }, { value: 2, label: '已拒绝' }]}
        />
      </div>
      <Table dataSource={comments} columns={columns} rowKey="id" loading={loading}
        pagination={{ current: page, total, pageSize, onChange: setPage }} />

      <Modal title="回复评论" open={replyModal.open} onCancel={() => setReplyModal({ open: false, commentId: null })}
        onOk={handleReply} okText="回复">
        <Input.TextArea rows={4} value={replyContent} onChange={(e) => setReplyContent(e.target.value)} placeholder="输入回复内容..." />
      </Modal>
    </div>
  );
}
