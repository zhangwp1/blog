import { createBrowserRouter } from 'react-router-dom';
import AuthGuard from './AuthGuard';
import PublicLayout from '../components/layout/PublicLayout';
import AdminLayout from '../components/layout/AdminLayout';
import HomePage from '../pages/public/HomePage';
import ArticleDetailPage from '../pages/public/ArticleDetailPage';
import ArchivePage from '../pages/public/ArchivePage';
import LoginPage from '../pages/admin/LoginPage';
import DashboardPage from '../pages/admin/DashboardPage';
import ArticleManagePage from '../pages/admin/ArticleManagePage';
import ArticleEditPage from '../pages/admin/ArticleEditPage';
import CategoryManagePage from '../pages/admin/CategoryManagePage';
import TagManagePage from '../pages/admin/TagManagePage';
import CommentManagePage from '../pages/admin/CommentManagePage';

const router = createBrowserRouter([
  {
    path: '/',
    element: <PublicLayout />,
    children: [
      { index: true, element: <HomePage /> },
      { path: 'articles/:slug', element: <ArticleDetailPage /> },
      { path: 'archive', element: <ArchivePage /> },
    ],
  },
  {
    path: '/admin/login',
    element: <LoginPage />,
  },
  {
    path: '/admin',
    element: (
      <AuthGuard>
        <AdminLayout />
      </AuthGuard>
    ),
    children: [
      { index: true, element: <DashboardPage /> },
      { path: 'articles', element: <ArticleManagePage /> },
      { path: 'articles/new', element: <ArticleEditPage /> },
      { path: 'articles/:id/edit', element: <ArticleEditPage /> },
      { path: 'categories', element: <CategoryManagePage /> },
      { path: 'tags', element: <TagManagePage /> },
      { path: 'comments', element: <CommentManagePage /> },
    ],
  },
]);

export default router;
