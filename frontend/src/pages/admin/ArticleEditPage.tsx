import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Form, Input, Select, Switch, Button, Card, Space, message } from 'antd';
import MDEditor from '@uiw/react-md-editor';
import { articleApi } from '../../api/article';
import { categoryApi } from '../../api/category';
import { tagApi } from '../../api/tag';
import type { Category } from '../../types/category';
import type { Tag } from '../../types/tag';
import type { CreateArticleRequest } from '../../types/article';

export default function ArticleEditPage() {
  const { id } = useParams<{ id: string }>();
  const isEdit = !!id;
  const navigate = useNavigate();
  const [form] = Form.useForm();
  const [content, setContent] = useState('');
  const [loading, setLoading] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [categories, setCategories] = useState<Category[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);

  useEffect(() => {
    categoryApi.list().then((res) => setCategories(res.data.data || []));
    tagApi.list().then((res) => setTags(res.data.data || []));
  }, []);

  useEffect(() => {
    if (isEdit && id) {
      setLoading(true);
      articleApi.getById(Number(id)).then((res) => {
        const article = res.data.data;
        form.setFieldsValue({
          title: article.title,
          summary: article.summary,
          cover_image: article.cover_image,
          is_published: article.is_published,
          pinned: article.pinned,
          category_id: article.category_id,
          tag_ids: article.tags?.map((t) => t.id) || [],
        });
        setContent(article.content || '');
      }).finally(() => setLoading(false));
    }
  }, [id, isEdit, form]);

  const handleSubmit = async (values: Record<string, unknown>) => {
    const data: CreateArticleRequest = {
      title: values.title as string,
      content,
      summary: values.summary as string,
      cover_image: values.cover_image as string,
      is_published: values.is_published as boolean,
      pinned: values.pinned as boolean,
      category_id: values.category_id as number,
      tag_ids: values.tag_ids as number[],
    };

    setSubmitting(true);
    try {
      if (isEdit) {
        await articleApi.update(Number(id), data);
      } else {
        await articleApi.create(data);
      }
      message.success(isEdit ? '更新成功' : '创建成功');
      navigate('/admin/articles');
    } catch {
      message.error('操作失败');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div>
      <Card title={isEdit ? '编辑文章' : '新建文章'} loading={loading} extra={<Button onClick={() => navigate('/admin/articles')}>返回列表</Button>}>
        <Form form={form} layout="vertical" onFinish={handleSubmit}
          initialValues={{ is_published: false, pinned: false, category_id: undefined, tag_ids: [] }}>
          <Form.Item name="title" label="标题" rules={[{ required: true, message: '请输入标题' }]}>
            <Input placeholder="文章标题" />
          </Form.Item>

          <Form.Item label="内容" required>
            <div data-color-mode="light">
              <MDEditor value={content} onChange={(v) => setContent(v || '')} height={500} />
            </div>
          </Form.Item>

          <Form.Item name="summary" label="摘要">
            <Input.TextArea rows={2} placeholder="文章摘要（选填）" />
          </Form.Item>

          <Form.Item name="cover_image" label="封面图片URL">
            <Input placeholder="封面图片链接（选填）" />
          </Form.Item>

          <Space size="large" style={{ display: 'flex' }}>
            <Form.Item name="category_id" label="分类" style={{ minWidth: 200 }}>
              <Select placeholder="选择分类" allowClear options={categories.map((c) => ({ value: c.id, label: c.name }))} />
            </Form.Item>
            <Form.Item name="tag_ids" label="标签" style={{ minWidth: 300 }}>
              <Select mode="multiple" placeholder="选择标签" options={tags.map((t) => ({ value: t.id, label: t.name }))} />
            </Form.Item>
          </Space>

          <Space size="large">
            <Form.Item name="is_published" label="发布状态" valuePropName="checked">
              <Switch />
            </Form.Item>
            <Form.Item name="pinned" label="置顶" valuePropName="checked">
              <Switch />
            </Form.Item>
          </Space>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={submitting} style={{ marginRight: 12 }}>
              {isEdit ? '更新' : '创建'}
            </Button>
            <Button onClick={() => navigate('/admin/articles')}>取消</Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
}
