import { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, InputNumber, Space, Popconfirm, message } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { categoryApi } from '../../api/category';
import type { Category, CreateCategoryRequest } from '../../types/category';

export default function CategoryManagePage() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [editing, setEditing] = useState<Category | null>(null);
  const [form] = Form.useForm();

  const fetch = () => {
    setLoading(true);
    categoryApi.list().then((res) => setCategories(res.data.data || [])).finally(() => setLoading(false));
  };

  useEffect(() => { fetch(); }, []);

  const handleCreate = () => {
    setEditing(null);
    form.resetFields();
    setModalOpen(true);
  };

  const handleEdit = (record: Category) => {
    setEditing(record);
    form.setFieldsValue(record);
    setModalOpen(true);
  };

  const handleSubmit = async (values: CreateCategoryRequest) => {
    try {
      if (editing) {
        await categoryApi.update(editing.id, values);
        message.success('更新成功');
      } else {
        await categoryApi.create(values);
        message.success('创建成功');
      }
      setModalOpen(false);
      fetch();
    } catch {
      message.error('操作失败');
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await categoryApi.delete(id);
      message.success('删除成功');
      fetch();
    } catch {
      message.error('删除失败');
    }
  };

  const columns = [
    { title: '名称', dataIndex: 'name', key: 'name' },
    { title: 'Slug', dataIndex: 'slug', key: 'slug' },
    { title: '排序', dataIndex: 'sort_order', key: 'sort_order', width: 80 },
    {
      title: '操作', key: 'action', width: 160,
      render: (_: unknown, record: Category) => (
        <Space>
          <Button type="link" icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
          <Popconfirm title="确定删除?" onConfirm={() => handleDelete(record.id)}>
            <Button type="link" danger icon={<DeleteOutlined />}>删除</Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>新建分类</Button>
      </div>
      <Table dataSource={categories} columns={columns} rowKey="id" loading={loading} pagination={false} />

      <Modal title={editing ? '编辑分类' : '新建分类'} open={modalOpen} onCancel={() => setModalOpen(false)} onOk={() => form.submit()} destroyOnClose>
        <Form form={form} layout="vertical" onFinish={handleSubmit}>
          <Form.Item name="name" label="名称" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item name="slug" label="Slug" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item name="sort_order" label="排序">
            <InputNumber style={{ width: '100%' }} />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}
