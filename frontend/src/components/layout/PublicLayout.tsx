import { Outlet } from 'react-router-dom';
import Header from './Header';
import Footer from './Footer';

export default function PublicLayout() {
  return (
    <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column', background: '#f5f5f5' }}>
      <Header />
      <main style={{ flex: 1, maxWidth: 1100, width: '100%', margin: '24px auto', padding: '0 16px' }}>
        <div style={{ display: 'flex', gap: 24 }}>
          <div style={{ flex: 1 }}>
            <Outlet />
          </div>
          <aside style={{ width: 280, flexShrink: 0 }}>
            {/* Sidebar widgets will go here */}
          </aside>
        </div>
      </main>
      <Footer />
    </div>
  );
}
