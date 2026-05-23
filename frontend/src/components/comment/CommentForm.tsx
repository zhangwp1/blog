import { useState } from 'react';
import { Form, Input, Button, message } from 'antd';

interface Props {
  onSubmit: (data: { author_name: string; author_email: string; content: string }) => Promise<void>;
  placeholder?: string;
}

export default function CommentForm({ onSubmit, placeholder = '写下你的评论...' }: Props) {
  const [loading, setLoading] = useState(false);
  const [form] = Form.useForm();

  const handleSubmit = async (values: { author_name: string; author_email: string; content: string }) => {
    setLoading(true);
    try {
      await onSubmit(values);
      message.success('评论提交成功，等待审核');
      form.resetFields();
    } catch {
      message.error('评论提交失败');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Form form={form} layout="vertical" onFinish={handleSubmit} style={{ marginTop: 16 }}>
      <Form.Item name="content" rules={[{ required: true, message: '请输入评论内容' }]}>
        <Input.TextArea rows={3} placeholder={placeholder} />
      </Form.Item>
      <div style={{ display: 'flex', gap: 12 }}>
        <Form.Item name="author_name" rules={[{ required: true, message: '请输入昵称' }]} style={{ flex: 1 }}>
          <Input placeholder="昵称" />
        </Form.Item>
        <Form.Item name="author_email" rules={[{ type: 'email', message: '邮箱格式不正确' }]} style={{ flex: 1 }}>
          <Input placeholder="邮箱（选填）" />
        </Form.Item>
      </div>
      <Button type="primary" htmlType="submit" loading={loading}>
        提交评论
      </Button>
    </Form>
  );
}
