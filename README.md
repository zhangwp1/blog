# 个人博客

前后端分离的个人博客网站，采用 React + Go 技术栈，支持 Markdown 文章编辑、分类标签管理、评论审核等功能。

## 技术栈

| 层级 | 技术 |
|------|------|
| 前端框架 | React 19 + TypeScript |
| 构建工具 | Vite 8 |
| UI 框架 | Ant Design 6（后台）+ Tailwind CSS 4（前台） |
| 路由 | React Router v7 |
| 状态管理 | Zustand |
| Markdown | @uiw/react-md-editor（编辑）+ react-markdown（渲染） |
| 后端框架 | Go + Gin |
| ORM | GORM v2 |
| 数据库 | SQLite（纯 Go 驱动，零配置） |
| 认证 | JWT Bearer Token |

## 功能

- **文章管理**：Markdown 编辑器，支持草稿/发布、置顶、分类、多标签
- **前台展示**：文章列表（分页、筛选、搜索）、文章详情（代码高亮）、文章归档
- **分类 & 标签**：独立管理，文章按分类/标签筛选
- **评论系统**：树形评论（支持回复）、审核机制（待审/通过/拒绝）、管理员回复
- **后台管理**：仪表盘统计、文章/分类/标签/评论 CRUD、JWT 登录认证

## 快速开始

### 环境要求

- Go 1.26+
- Node.js 18+

### 1. 启动后端

```bash
cd backend

# 安装依赖（国内需设置代理）
GOPROXY=https://goproxy.cn,direct go mod tidy

# 启动服务（默认端口 8080）
go run ./cmd/server/
```

首次运行会自动创建 SQLite 数据库（`backend/data/blog.db`）并生成默认管理员账号。

### 2. 启动前端

```bash
cd frontend

# 安装依赖
npm install

# 启动开发服务器（默认端口 5173）
npm run dev
```

### 3. 访问

| 地址 | 说明 |
|------|------|
| http://localhost:5173 | 前台首页 |
| http://localhost:5173/admin/login | 后台登录 |

**默认管理员**：`admin` / `admin123`

## 项目结构

```
zhuomin/
├── frontend/                 # React 前端
│   └── src/
│       ├── api/              # Axios 接口封装
│       ├── components/       # 通用组件（布局/Markdown/评论）
│       ├── pages/            # public/ 前台页面  admin/ 后台页面
│       ├── router/           # 路由配置 + 认证守卫
│       ├── store/            # Zustand 状态管理
│       └── types/            # TypeScript 类型定义
├── backend/                  # Go 后端
│   ├── cmd/server/main.go    # 入口文件
│   ├── config/               # 配置（Viper + YAML）
│   └── internal/
│       ├── handler/          # 控制器层（参数绑定 → 响应）
│       ├── service/          # 业务逻辑层
│       ├── repository/       # 数据访问层（GORM）
│       ├── model/            # 数据库模型
│       ├── dto/              # 请求/响应 DTO
│       ├── middleware/        # 中间件（JWT/CORS/日志/恢复）
│       ├── router/           # 路由注册
│       └── utils/            # 工具函数（JWT/统一响应）
├── ARCHITECTURE.md           # 架构设计文档
├── DEPENDENCIES.md           # 依赖清单
└── CLAUDE.md                 # Claude Code 项目指南
```

## API 接口

基础路径 `/api/v1`，统一响应格式 `{ code, message, data }`。

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /auth/login | 管理员登录 |
| GET | /articles | 文章列表（筛选/分页/搜索） |
| GET | /articles/:slug | 文章详情 |
| GET | /categories | 分类列表 |
| GET | /tags | 标签列表 |
| GET | /articles/:slug/comments | 文章评论（树形） |
| POST | /articles/:slug/comments | 提交评论 |

### 管理接口（需 JWT 认证）

所有管理接口位于 `/api/v1/admin/` 路径下，请求头需附带 `Authorization: Bearer <token>`。

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /admin/dashboard | 仪表盘统计 |
| GET/POST/PUT/DELETE | /admin/articles | 文章 CRUD |
| POST/PUT/DELETE | /admin/categories | 分类管理 |
| POST/PUT/DELETE | /admin/tags | 标签管理 |
| GET/PATCH/DELETE | /admin/comments | 评论审核 |
| POST | /admin/comments/:id/reply | 管理员回复 |

## 配置

后端配置文件 `backend/config/config.yaml`：

```yaml
server:
  port: 8080          # 服务端口
  mode: debug         # debug / release

database:
  file: "data/blog.db"  # SQLite 数据库文件路径

jwt:
  secret: "your-secret-key"  # JWT 签名密钥
  expire_hours: 24            # Token 过期时间（小时）
```

## 相关文档

- [ARCHITECTURE.md](./ARCHITECTURE.md) — 架构设计文档（ER 图、数据流、分层详解）
- [DEPENDENCIES.md](./DEPENDENCIES.md) — 完整依赖清单
- [CLAUDE.md](./CLAUDE.md) — Claude Code 开发指南
