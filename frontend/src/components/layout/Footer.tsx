export default function Footer() {
  return (
    <footer style={{ textAlign: 'center', padding: '24px 0', color: '#999', fontSize: 13, background: '#fff', borderTop: '1px solid #f0f0f0' }}>
      <div>Powered by React + Go · © {new Date().getFullYear()} 我的博客</div>
    </footer>
  );
}
