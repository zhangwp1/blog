import { Link, useLocation } from 'react-router-dom';
import { Menu } from 'antd';

const items = [
  { key: '/', label: <Link to="/">首页</Link> },
  { key: '/archive', label: <Link to="/archive">归档</Link> },
];

export default function Header() {
  const location = useLocation();
  const selectedKey = location.pathname === '/' ? '/' : '/archive';

  return (
    <header style={{ background: '#fff', boxShadow: '0 2px 8px rgba(0,0,0,0.06)', position: 'sticky', top: 0, zIndex: 100 }}>
      <div style={{ maxWidth: 1100, margin: '0 auto', display: 'flex', alignItems: 'center', justifyContent: 'space-between', padding: '0 16px', height: 56 }}>
        <Link to="/" style={{ fontSize: 20, fontWeight: 700, color: '#1677ff', textDecoration: 'none' }}>
          我的博客
        </Link>
        <Menu mode="horizontal" selectedKeys={[selectedKey]} items={items} style={{ border: 'none', flex: 1, justifyContent: 'flex-end' }} />
      </div>
    </header>
  );
}
