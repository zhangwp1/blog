import { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, Space, Popconfirm, message } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { tagApi } from '../../api/tag';
import type { Tag, CreateTagRequest } from '../../types/tag';

export default function TagManagePage() {
  const [tags, setTags] = useState<Tag[]>([]);
  const [loading, setLoading] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [editing, setEditing] = useState<Tag | null>(null);
  const [form] = Form.useForm();

  const fetch = () => {
    setLoading(true);
    tagApi.list().then((res) => setTags(res.data.data || [])).finally(() => setLoading(false));
  };

  useEffect(() => { fetch(); }, []);

  const handleCreate = () => {
    setEditing(null);
    form.resetFields();
    setModalOpen(true);
  };

  const handleEdit = (record: Tag) => {
    setEditing(record);
    form.setFieldsValue(record);
    setModalOpen(true);
  };

  const handleSubmit = async (values: CreateTagRequest) => {
    try {
      if (editing) {
        await tagApi.update(editing.id, values);
        message.success('更新成功');
      } else {
        await tagApi.create(values);
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
      await tagApi.delete(id);
      message.success('删除成功');
      fetch();
    } catch {
      message.error('删除失败');
    }
  };

  const columns = [
    { title: '名称', dataIndex: 'name', key: 'name' },
    { title: 'Slug', dataIndex: 'slug', key: 'slug' },
    {
      title: '操作', key: 'action', width: 160,
      render: (_: unknown, record: Tag) => (
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
        <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>新建标签</Button>
      </div>
      <Table dataSource={tags} columns={columns} rowKey="id" loading={loading} pagination={false} />

      <Modal title={editing ? '编辑标签' : '新建标签'} open={modalOpen} onCancel={() => setModalOpen(false)} onOk={() => form.submit()} destroyOnClose>
        <Form form={form} layout="vertical" onFinish={handleSubmit}>
          <Form.Item name="name" label="名称" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item name="slug" label="Slug" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}
