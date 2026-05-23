# 前后端分离个人博客 — 架构设计文档

## 一、总体架构

```
┌─────────────────────────────────────────────────────────────┐
│                        浏览器客户端                          │
│                  http://localhost:5173                       │
└────────────┬────────────────────────────┬───────────────────┘
             │                            │
    前台页面（公开）                 后台管理页面（需登录）
    首页 / 文章详情 / 归档          仪表盘 / 文章管理 / 分类标签 / 评论审核
             │                            │
             └────────────┬───────────────┘
                          │  HTTP REST API (JSON)
                          │  Authorization: Bearer <token>
                          ▼
┌─────────────────────────────────────────────────────────────┐
│              Gin Web Server (:8080)                          │
│                                                             │
│  Middleware Chain: Recovery → Logger → CORS → [Auth]        │
│                                                             │
│  /api/v1/*         → 公开路由（无需认证）                     │
│  /api/v1/admin/*   → 管理路由（JWT 认证）                    │
└────────────────────────────┬────────────────────────────────┘
                             │
                    handler → service → repository
                             │
                             ▼
┌─────────────────────────────────────────────────────────────┐
│                   SQLite (data/blog.db)                      │
│                                                             │
│  users │ categories │ tags │ articles │ article_tags │ comments │
└─────────────────────────────────────────────────────────────┘
```

**核心原则**：
- 前后端完全分离，通过 HTTP JSON API 通信
- 后端严格分层：handler（控制器）→ service（业务逻辑）→ repository（数据访问）
- 公开接口与认证接口路由隔离
- 统一响应格式 `{ code, message, data }`

---

## 二、项目目录结构

```
zhuomin/
├── frontend/                          # React + TypeScript 前端
│   ├── public/                        # 静态资源
│   ├── src/
│   │   ├── api/                       # Axios 实例 + 各模块 API 封装
│   │   │   ├── client.ts             # Axios 实例（baseURL、拦截器）
│   │   │   ├── auth.ts
│   │   │   ├── article.ts
│   │   │   ├── category.ts
│   │   │   ├── tag.ts
│   │   │   └── comment.ts
│   │   ├── components/
│   │   │   ├── layout/
│   │   │   │   ├── PublicLayout.tsx  # 前台布局（Header + 内容 + Footer）
│   │   │   │   ├── AdminLayout.tsx   # 后台布局（侧边栏 + 内容）
│   │   │   │   ├── Header.tsx
│   │   │   │   └── Footer.tsx
│   │   │   ├── markdown/
│   │   │   │   └── MdRenderer.tsx    # Markdown 渲染组件
│   │   │   └── comment/
│   │   │       ├── CommentList.tsx   # 评论列表（递归树形）
│   │   │       └── CommentForm.tsx   # 评论提交表单
│   │   ├── hooks/                    # 自定义 React Hooks
│   │   ├── pages/
│   │   │   ├── public/              # 前台页面
│   │   │   │   ├── HomePage.tsx     # 首页（文章列表 + 侧边栏）
│   │   │   │   ├── ArticleDetailPage.tsx  # 文章详情 + 评论
│   │   │   │   └── ArchivePage.tsx  # 文章归档
│   │   │   └── admin/               # 后台管理页面
│   │   │       ├── LoginPage.tsx
│   │   │       ├── DashboardPage.tsx
│   │   │       ├── ArticleManagePage.tsx
│   │   │       ├── ArticleEditPage.tsx    # Markdown 编辑器
│   │   │       ├── CategoryManagePage.tsx
│   │   │       ├── TagManagePage.tsx
│   │   │       └── CommentManagePage.tsx
│   │   ├── router/
│   │   │   ├── index.tsx            # 路由配置表
│   │   │   └── AuthGuard.tsx        # 路由守卫
│   │   ├── store/
│   │   │   └── authSlice.ts         # Zustand 认证状态
│   │   ├── types/                   # TypeScript 类型定义
│   │   ├── App.tsx
│   │   └── main.tsx                 # 应用入口
│   ├── vite.config.ts               # Vite 配置（含 API 代理）
│   ├── tsconfig.json
│   └── package.json
│
├── backend/                           # Go + Gin 后端
│   ├── cmd/server/
│   │   └── main.go                   # 入口：初始化数据库、自动建表、启动服务
│   ├── config/
│   │   ├── config.go                 # 配置结构体（需 mapstructure 标签）
│   │   └── config.yaml               # YAML 配置文件
│   ├── internal/
│   │   ├── handler/                  # 控制器层
│   │   │   ├── init.go              # 依赖注入（手动组装所有依赖）
│   │   │   ├── auth.go
│   │   │   ├── article.go
│   │   │   ├── category.go
│   │   │   ├── tag.go
│   │   │   └── comment.go
│   │   ├── service/                  # 业务逻辑层
│   │   │   ├── auth.go
│   │   │   ├── article.go
│   │   │   ├── category.go
│   │   │   ├── tag.go
│   │   │   └── comment.go           # 含评论树构建算法
│   │   ├── repository/               # 数据访问层（GORM）
│   │   │   ├── user.go
│   │   │   ├── article.go           # 含复杂筛选查询（标签/分类/归档/关键词）
│   │   │   ├── category.go
│   │   │   ├── tag.go
│   │   │   └── comment.go
│   │   ├── model/                    # 数据库模型
│   │   │   ├── user.go
│   │   │   ├── category.go
│   │   │   ├── tag.go
│   │   │   ├── article.go           # 含 GORM 关联定义（BelongsTo、Many2Many）
│   │   │   └── comment.go           # 含 parent_id 自引用
│   │   ├── dto/                      # 数据传输对象
│   │   │   ├── auth.go
│   │   │   ├── article.go
│   │   │   ├── category.go
│   │   │   ├── tag.go
│   │   │   └── comment.go
│   │   ├── middleware/               # 中间件
│   │   │   ├── auth.go              # JWT 认证中间件
│   │   │   ├── cors.go              # CORS 跨域中间件
│   │   │   ├── logger.go            # 请求日志
│   │   │   └── recovery.go          # Panic 恢复
│   │   ├── router/
│   │   │   └── router.go            # 路由注册（公开/认证/管理三组路由）
│   │   └── utils/
│   │       ├── jwt.go               # JWT 生成与解析
│   │       └── response.go          # 统一响应格式（Success / PageSuccess / Error）
│   └── migrations/
│       └── 001_init.sql             # MySQL 参考 DDL（实际使用 GORM AutoMigrate）
│
├── docker-compose.yml                # MySQL 参考配置（未使用）
├── CLAUDE.md                        # Claude Code 项目指南
├── DEPENDENCIES.md                  # 依赖清单
└── ARCHITECTURE.md                  # 本文件
```

---

## 三、数据库设计

### 3.1 ER 图（实体关系）

```
┌──────────┐         ┌──────────────┐         ┌──────────┐
│  users   │ 1     N │   articles   │ N     1 │ categories│
│──────────│◄────────│──────────────│────────►│──────────│
│ id (PK)  │         │ id (PK)      │         │ id (PK)  │
│ username │         │ title        │         │ name     │
│ password │         │ slug (UQ)    │         │ slug     │
│ nickname │         │ content      │         │ sort_order│
│ avatar   │         │ summary      │         └──────────┘
└──────────┘         │ cover_image  │
                     │ is_published │         ┌──────────┐
                     │ pinned       │ N     M │   tags   │
                     │ view_count   │◄───────►│──────────│
                     │ category_id  │         │ id (PK)  │
                     │ author_id    │         │ name     │
                     │ published_at │         │ slug     │
                     └──────┬───────┘         └──────────┘
                            │ 1                    ▲
                            │                      │
                            │ N            ┌───────┴──────┐
                     ┌──────▼───────┐     │ article_tags  │
                     │   comments   │     │───────────────│
                     │──────────────│     │ article_id(PK)│
                     │ id (PK)      │     │ tag_id (PK)   │
                     │ article_id   │     └───────────────┘
                     │ parent_id ◄──┤ (自引用：回复链)
                     │ author_name  │
                     │ content      │
                     │ is_approved  │ (0=待审 1=通过 2=拒绝)
                     │ is_admin     │
                     └──────────────┘
```

### 3.2 表结构详解

#### users（管理员用户）

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | INTEGER | PK, AUTO | 主键 |
| username | VARCHAR(64) | UNIQUE, NOT NULL | 用户名 |
| password | VARCHAR(255) | NOT NULL | bcrypt 加密存储，JSON 序列化时隐藏 `json:"-"` |
| nickname | VARCHAR(64) | | 显示昵称 |
| avatar | VARCHAR(255) | | 头像 URL |
| created_at | DATETIME | | 创建时间 |
| updated_at | DATETIME | | 更新时间 |

**设计要点**：单用户博客，通过 seedAdmin() 自动创建默认管理员 admin/admin123，password 字段使用 `json:"-"` 标签确保不会泄露到 API 响应中。

#### categories（文章分类）

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | INTEGER | PK, AUTO | 主键 |
| name | VARCHAR(64) | NOT NULL | 分类名称 |
| slug | VARCHAR(64) | UNIQUE, NOT NULL | URL 友好标识 |
| sort_order | INTEGER | DEFAULT 0 | 排序权重 |
| created_at | DATETIME | | 创建时间 |
| updated_at | DATETIME | | 更新时间 |

#### tags（文章标签）

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | INTEGER | PK, AUTO | 主键 |
| name | VARCHAR(64) | NOT NULL | 标签名称 |
| slug | VARCHAR(64) | UNIQUE, NOT NULL | URL 友好标识 |
| created_at | DATETIME | | 创建时间 |
| updated_at | DATETIME | | 更新时间 |

#### articles（文章）— 核心表

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | INTEGER | PK, AUTO | 主键 |
| title | VARCHAR(255) | NOT NULL | 文章标题 |
| slug | VARCHAR(255) | UNIQUE, NOT NULL | URL 友好标识（由标题自动生成） |
| content | MEDIUMTEXT | NOT NULL | Markdown 格式正文 |
| summary | VARCHAR(512) | | 文章摘要 |
| cover_image | VARCHAR(255) | | 封面图 URL |
| is_published | BOOLEAN | DEFAULT false | 是否发布 |
| pinned | BOOLEAN | DEFAULT false | 是否置顶 |
| view_count | INTEGER | DEFAULT 0 | 阅读量 |
| category_id | INTEGER | DEFAULT 0 | 外键 → categories.id |
| author_id | INTEGER | NOT NULL | 外键 → users.id |
| published_at | DATETIME | | 首次发布时间（用于排序和归档） |
| created_at | DATETIME | | 创建时间 |
| updated_at | DATETIME | | 更新时间 |

**GORM 关联定义**：
```go
Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
Author   User     `gorm:"foreignKey:AuthorID"  json:"author,omitempty"`
Tags     []Tag    `gorm:"many2many:article_tags" json:"tags,omitempty"`
```

**查询策略**：使用 `Preload("Category").Preload("Tags").Preload("Author")` 一次性加载关联数据，避免 N+1 查询问题。

#### article_tags（文章-标签关联表）

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| article_id | INTEGER | PK | 联合主键 |
| tag_id | INTEGER | PK | 联合主键 |

GORM 自动管理此表，无需手动操作。

#### comments（评论）

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | INTEGER | PK, AUTO | 主键 |
| article_id | INTEGER | NOT NULL, INDEX | 外键 → articles.id |
| parent_id | INTEGER | DEFAULT 0 | 父评论 ID（0 = 顶级评论，自引用实现回复链） |
| author_name | VARCHAR(64) | NOT NULL | 评论者昵称 |
| author_email | VARCHAR(128) | | 评论者邮箱 |
| author_website | VARCHAR(255) | | 评论者网站 |
| content | TEXT | NOT NULL | 评论内容 |
| is_approved | TINYINT | DEFAULT 0 | 审核状态：0=待审，1=通过，2=拒绝 |
| is_admin | BOOLEAN | DEFAULT false | 是否管理员回复 |
| ip | VARCHAR(45) | | 评论者 IP（`json:"-"` 不输出到 API） |
| user_agent | VARCHAR(255) | | 浏览器 UA（`json:"-"` 不输出到 API） |
| created_at | DATETIME | | 评论时间 |

**树形结构**：`Children []*Comment` 字段标记 `gorm:"-"`（非数据库字段），通过 `buildCommentTree()` 算法在内存中构建。

---

## 四、后端架构

### 4.1 分层架构

```
┌─────────────────────────────────────────────────────────┐
│                     handler 层                           │
│  职责：参数绑定与校验、调用 service、返回统一 JSON 响应     │
│  依赖：service 接口                                      │
│  禁止：直接访问 repository 或 GORM                        │
├─────────────────────────────────────────────────────────┤
│                     service 层                           │
│  职责：业务逻辑、数据转换（model ↔ dto）、权限判断          │
│  依赖：repository 接口                                    │
│  禁止：直接操作 HTTP 上下文（gin.Context）                  │
├─────────────────────────────────────────────────────────┤
│                    repository 层                          │
│  职责：数据库 CRUD、复杂查询构建、分页处理                  │
│  依赖：*gorm.DB                                          │
│  禁止：包含业务逻辑                                       │
├─────────────────────────────────────────────────────────┤
│                    model 层                               │
│  职责：GORM 模型定义、关联关系声明                          │
│  依赖：无                                                 │
└─────────────────────────────────────────────────────────┘
```

### 4.2 依赖注入

所有依赖在 `handler/init.go` 中手动组装，不使用依赖注入框架：

```go
func InitHandlers(db *gorm.DB) *Handlers {
    // 1. 创建 Repository 实例
    userRepo    := repository.NewUserRepository(db)
    articleRepo := repository.NewArticleRepository(db)
    categoryRepo := repository.NewCategoryRepository(db)
    tagRepo     := repository.NewTagRepository(db)
    commentRepo := repository.NewCommentRepository(db)

    // 2. 创建 Service 实例（注入 Repository）
    authService    := service.NewAuthService(userRepo)
    articleService := service.NewArticleService(articleRepo, categoryRepo, tagRepo)
    categoryService := service.NewCategoryService(categoryRepo)
    tagService     := service.NewTagService(tagRepo)
    commentService := service.NewCommentService(commentRepo)

    // 3. 创建 Handler 实例（注入 Service）
    return &Handlers{
        Auth:     NewAuthHandler(authService),
        Article:  NewArticleHandler(articleService),
        Category: NewCategoryHandler(categoryService),
        Tag:      NewTagHandler(tagService),
        Comment:  NewCommentHandler(commentService, articleRepo),
    }
}
```

### 4.3 中间件链

每个请求依次经过以下中间件：

```
请求 → Recovery（panic 恢复）
     → Logger（请求日志）
     → CORS（跨域处理，AllowAllOrigins）
     → [Auth]（JWT 令牌校验，仅需认证路由）
     → 路由处理函数
```

**CORS 配置**：由于 `gin-contrib/cors` v1.7+ 不允许空的 AllowOrigins，使用 `AllowAllOrigins: true` 并配合 `AllowCredentials: false`。

**JWT 认证**：
1. 从 `Authorization` 头提取 `Bearer <token>`
2. 使用 `golang-jwt/jwt/v5` 解析 token
3. 将 `user_id` 和 `username` 注入 `gin.Context`
4. token 有效期通过配置文件的 `jwt.expire_hours` 控制（默认 24 小时）

### 4.4 统一响应格式

```go
// 成功
{ "code": 0, "message": "ok", "data": {...} }

// 分页
{ "code": 0, "message": "ok", "data": { "list": [...], "total": 100, "page": 1, "page_size": 10 } }

// 错误（HTTP 状态码仍为 200，通过 code 区分）
{ "code": 400, "message": "参数错误", "data": null }

// 认证错误（HTTP 状态码 401）
{ "code": 401, "message": "token无效或已过期", "data": null }
```

### 4.5 关键业务逻辑

#### 文章查询（FindList）

支持多重筛选条件组合，按优先级叠加：

1. **发布状态过滤**：`is_published = ?`（公开列表固定为 true）
2. **分类过滤**：`category_id = ?`（通过 category slug 反查 ID）
3. **标签过滤**：JOIN article_tags + tags 表 → `tags.slug = ?`
4. **关键词搜索**：`title LIKE %keyword% OR content LIKE %keyword%`
5. **归档过滤**：`YEAR(published_at) = ? AND MONTH(published_at) = ?`
6. **置顶过滤**：`pinned = ?`
7. **排序**：`pinned DESC, published_at DESC`（置顶优先，然后按发布时间倒序）
8. **分页**：Offset + Limit，默认每页 10 条

#### 评论树构建（buildCommentTree）

```
输入：扁平评论列表（按时间排序）
输出：树形结构 [{...children: [{...children: []}]}]

算法：
1. 遍历所有评论，创建 nodeMap[id] = &CommentResponse
2. 再次遍历，根据 parent_id 挂载：
   - parent_id == 0 → 放入 roots[]
   - parent_id != 0 → 找到 nodeMap[parent_id]，追加到其 Children
```

#### 文章编辑器更新流程

```
前端 ArticleEditPage
  → articleApi.update(id, data)
    → PUT /api/v1/admin/articles/:id
      → handler.Update（参数绑定）
        → service.Update（slug 重新生成、发布时间处理）
          → repository.Update（db.Save + Association.Replace）
```

**关键注意点**：repository.Update 必须先 `db.Save(article)` 保存文章字段，再 `Association("Tags").Replace(article.Tags)` 替换多对多关联。仅调用 Replace 不会保存文章自身字段。

---

## 五、前端架构

### 5.1 路由设计

```
┌──────────────────────────────────────────────────────┐
│                    React Router v7                     │
├──────────────────────────────────────────────────────┤
│                                                      │
│  /                          → HomePage               │
│  /articles/:slug            → ArticleDetailPage      │
│  /archive                   → ArchivePage            │
│                           （使用 PublicLayout 布局）    │
│                                                      │
│  /admin/login               → LoginPage              │
│                           （独立页面，无布局）           │
│                                                      │
│  /admin                     → DashboardPage          │
│  /admin/articles            → ArticleManagePage      │
│  /admin/articles/new        → ArticleEditPage        │
│  /admin/articles/:id/edit   → ArticleEditPage        │
│  /admin/categories          → CategoryManagePage     │
│  /admin/tags                → TagManagePage          │
│  /admin/comments            → CommentManagePage      │
│                           （使用 AdminLayout + AuthGuard）│
└──────────────────────────────────────────────────────┘
```

### 5.2 路由守卫（AuthGuard）

```
AuthGuard 组件
  ├── 无 token → 重定向到 /admin/login
  ├── 有 token 但无用户信息 → 调用 fetchProfile() 获取用户信息，显示 Loading
  └── 有 token 且有用户信息 → 渲染子组件（AdminLayout）
```

### 5.3 状态管理（Zustand AuthSlice）

```
useAuthStore
  ├── token: string | null        // 从 localStorage 恢复
  ├── user: User | null           // 从 localStorage 恢复
  ├── setToken(token)             // 存储到 localStorage + 更新状态
  ├── setUser(user)               // 存储到 localStorage + 更新状态
  ├── fetchProfile()              // 调用 /auth/profile → setUser
  ├── logout()                    // 清除 localStorage + 重置状态
  └── isAuthenticated()           // 检查 token 是否存在
```

**设计决策**：选择 Zustand 而非 Redux，因为博客状态复杂度低，仅需管理认证状态，Zustand 无需 Provider、Reducer、Action Creator 等样板代码。

### 5.4 Axios 拦截器

```
请求拦截器：
  config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`

响应拦截器：
  ├── 业务错误（code !== 0）→ message.error() + Promise.reject
  └── HTTP 401 → 清除 localStorage + 跳转 /admin/login
```

### 5.5 组件层次（前台文章详情页示例）

```
PublicLayout
├── Header（导航栏）
├── ArticleDetailPage
│   ├── 文章标题 / 元信息（分类、标签、发布日期、阅读量）
│   ├── MdRenderer（react-markdown 渲染 Markdown 正文）
│   ├── 上一篇 / 下一篇导航
│   └── CommentList（递归组件）
│       ├── CommentForm（评论提交表单）
│       └── CommentItem[]
│           ├── 评论内容
│           ├── 回复按钮
│           └── CommentItem[]（子回复，递归渲染）
└── Footer
```

---

## 六、API 接口完整清单

### 6.1 公开接口

| 方法 | 路径 | 功能 | 请求参数 |
|------|------|------|----------|
| GET | /api/v1/health | 健康检查 | 无 |
| POST | /api/v1/auth/login | 管理员登录 | `{ username, password }` |
| GET | /api/v1/auth/profile | 当前用户信息 | Header: Bearer token |
| GET | /api/v1/articles | 文章列表 | `?keyword&category&tag&year&month&page&page_size` |
| GET | /api/v1/articles/:slug | 文章详情 | 路径参数 slug |
| GET | /api/v1/categories | 分类列表 | 无 |
| GET | /api/v1/tags | 标签列表 | 无 |
| GET | /api/v1/articles/:slug/comments | 文章评论 | 路径参数 slug |
| POST | /api/v1/articles/:slug/comments | 提交评论 | `{ article_id, parent_id, author_name, author_email, author_website, content }` |

### 6.2 管理端接口（全部需 JWT 认证）

| 方法 | 路径 | 功能 |
|------|------|------|
| GET | /api/v1/admin/dashboard | 仪表盘统计 |
| GET | /api/v1/admin/articles | 文章列表（含未发布） |
| GET | /api/v1/admin/articles/:id | 文章详情（按 ID） |
| POST | /api/v1/admin/articles | 创建文章 |
| PUT | /api/v1/admin/articles/:id | 更新文章 |
| DELETE | /api/v1/admin/articles/:id | 删除文章 |
| POST | /api/v1/admin/categories | 创建分类 |
| PUT | /api/v1/admin/categories/:id | 更新分类 |
| DELETE | /api/v1/admin/categories/:id | 删除分类 |
| POST | /api/v1/admin/tags | 创建标签 |
| PUT | /api/v1/admin/tags/:id | 更新标签 |
| DELETE | /api/v1/admin/tags/:id | 删除标签 |
| GET | /api/v1/admin/comments | 评论列表（审核管理） |
| PATCH | /api/v1/admin/comments/:id/approve | 通过评论 |
| PATCH | /api/v1/admin/comments/:id/reject | 拒绝评论 |
| DELETE | /api/v1/admin/comments/:id | 删除评论 |
| POST | /api/v1/admin/comments/:id/reply | 管理员回复 |

---

## 七、数据流图

### 7.1 用户浏览文章

```
浏览器                          Gin Server                    SQLite
  │                               │                             │
  │  GET /api/v1/articles         │                             │
  │  ?page=1&page_size=10         │                             │
  ├──────────────────────────────►│                             │
  │                               │  FindList(filter)           │
  │                               ├────────────────────────────►│
  │                               │◄────────────────────────────┤
  │                               │  articles[] + total         │
  │◄──────────────────────────────┤                             │
  │  { list: [...], total: 15 }  │                             │
  │                               │                             │
  │  GET /api/v1/articles/my-post │                             │
  ├──────────────────────────────►│                             │
  │                               │  FindBySlug("my-post")      │
  │                               ├────────────────────────────►│
  │                               │◄────────────────────────────┤
  │                               │  article (Preload 关联)      │
  │                               │  IncrementViewCount(1)      │
  │◄──────────────────────────────┤                             │
  │  { title, content(MD),       │                             │
  │    category, tags, author }  │                             │
  │                               │                             │
  │  react-markdown 渲染为 HTML   │                             │
```

### 7.2 管理员创建文章

```
浏览器 (Markdown编辑器)          Gin Server                    SQLite
  │                               │                             │
  │  POST /api/v1/admin/articles  │                             │
  │  { title, content("# 标题"), │                             │
  │    category_id, tag_ids }     │                             │
  ├──────────────────────────────►│                             │
  │                               │  1. slug.Make("文章标题")    │
  │                               │  2. find tags by ids        │
  │                               │  3. set PublishedAt         │
  │                               │  4. Create(article)         │
  │                               ├────────────────────────────►│
  │                               │◄────────────────────────────┤
  │                               │  article (with ID)          │
  │                               │  5. FindByID → Preload      │
  │◄──────────────────────────────┤                             │
  │  { id, slug, category, tags }│                             │
```

### 7.3 JWT 认证流程

```
┌──────────┐     ┌──────────┐     ┌──────────┐     ┌──────────┐
│  LoginPage│     │  Server  │     │  Axios   │     │  Admin   │
│           │     │          │     │Intercept.│     │  Routes  │
└─────┬─────┘     └─────┬────┘     └────┬─────┘     └────┬─────┘
      │                 │               │                 │
      │ POST /auth/login│               │                 │
      │ {admin,admin123}│               │                 │
      ├────────────────►│               │                 │
      │                 │ bcrypt.Compare│                 │
      │                 │ GenerateToken │                 │
      │◄────────────────┤               │                 │
      │ {token, user}   │               │                 │
      │                 │               │                 │
      │ localStorage.setItem('token')   │                 │
      │ zustand.setToken()              │                 │
      │ zustand.setUser()              │                 │
      │                 │               │                 │
      │ navigate('/admin')             │                 │
      │                 │               │ GET /admin/dash │
      │                 │               ├────────────────►│
      │                 │               │ Bearer <token>  │
      │                 │               │                 │──► Auth MW
      │                 │               │                 │    ParseToken
      │                 │               │                 │    set user_id
      │                 │               │                 │──► Handler
      │                 │               │◄────────────────┤
      │                 │               │ {stats}         │
```

---

## 八、关键设计决策

| 决策 | 选择 | 原因 |
|------|------|------|
| 数据库 | SQLite（非 MySQL） | 零配置、单文件、无需独立服务；个人博客数据量小，SQLite 完全满足 |
| SQLite 驱动 | glebarez/sqlite | 纯 Go 实现，无需 CGO，跨平台编译简单 |
| 配置管理 | Viper + YAML | Go 生态最流行的配置方案，支持多种格式和环境变量覆盖 |
| 密码加密 | bcrypt | 自适应哈希算法，抗暴力破解 |
| API 路径 | `/api/v1/` 前缀 | 预留版本升级空间 |
| 文章 URL | slug 而非 ID | SEO 友好，URL 可读 |
| 评论结构 | parent_id 自引用 | 支持无限层级回复，查询简单 |
| 状态管理 | Zustand | 轻量、无 Provider 包裹、API 简洁 |
| 依赖注入 | 手动组装 | 依赖关系简单，无需引入 Wire/Dig 等框架 |
| 中间件顺序 | Recovery → Logger → CORS → Auth | Recovery 在最外层确保 panic 被捕获；Auth 在最内层按需保护路由 |

---

## 九、部署架构（开发环境）

```
┌──────────────────────────────────────────────┐
│              开发机 (Windows 11)               │
│                                              │
│  ┌──────────────────┐  ┌──────────────────┐  │
│  │ Vite Dev Server  │  │  Gin Server      │  │
│  │ :5173            │──│  :8080           │  │
│  │                  │  │                  │  │
│  │ HMR 热更新       │  │ 直接读写          │  │
│  └──────────────────┘  └───────┬──────────┘  │
│                                │              │
│                        ┌───────▼──────────┐  │
│                        │  SQLite          │  │
│                        │  data/blog.db    │  │
│                        └──────────────────┘  │
└──────────────────────────────────────────────┘
```

- 前端通过 Axios 直接请求 `http://localhost:8080`（CORS 方式）
- Vite proxy `/api` → `http://localhost:8080` 作为备用方案
- SQLite 数据库文件自动创建于 `backend/data/blog.db`
- 无需 Docker 或外部数据库服务

---

## 十、安全性设计

| 层面 | 措施 |
|------|------|
| 认证 | JWT Bearer Token，24 小时过期 |
| 密码存储 | bcrypt 哈希，不可逆 |
| API 防护 | 管理端全部接口需 JWT 中间件校验 |
| 敏感字段 | password（`json:"-"`）、ip、user_agent 不对外输出 |
| 评论审核 | 所有公开评论默认待审（is_approved=0），防止垃圾内容 |
| CORS | AllowAllOrigins（开发阶段），生产环境应限制具体域名 |
| 前端路由守卫 | AuthGuard 组件拦截未登录访问，无法绕过 |
| 401 处理 | Axios 响应拦截器自动清除 token 并重定向登录页 |
