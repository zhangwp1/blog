# 项目依赖清单

## 后端（Go）

### 核心框架

| 包名 | 版本 | 说明 |
|------|------|------|
| `github.com/gin-gonic/gin` | v1.12.0 | HTTP Web 框架，路由、中间件、参数绑定 |
| `gorm.io/gorm` | v1.31.1 | ORM 框架，提供模型定义、查询构建、关联加载、事务支持 |

### 数据库驱动

| 包名 | 版本 | 说明 |
|------|------|------|
| `github.com/glebarez/sqlite` | v1.11.0 | SQLite 驱动（纯 Go 实现，无需 CGO），基于 modernc.org/sqlite |

### 认证与安全

| 包名 | 版本 | 说明 |
|------|------|------|
| `github.com/golang-jwt/jwt/v5` | v5.3.1 | JWT 令牌生成与解析 |
| `golang.org/x/crypto` | v0.52.0 | bcrypt 密码哈希 |

### 配置管理

| 包名 | 版本 | 说明 |
|------|------|------|
| `github.com/spf13/viper` | v1.21.0 | 配置文件加载，支持 YAML/JSON/ENV 等多种格式 |

### 中间件与工具

| 包名 | 版本 | 说明 |
|------|------|------|
| `github.com/gin-contrib/cors` | v1.7.7 | CORS 跨域中间件 |
| `github.com/gosimple/slug` | v1.15.0 | 文章标题自动生成 URL 友好的 slug |

### 间接依赖（核心）

| 包名 | 版本 | 说明 |
|------|------|------|
| `github.com/gin-contrib/sse` | v1.1.0 | Gin 的 SSE（Server-Sent Events）支持 |
| `github.com/glebarez/go-sqlite` | v1.21.2 | glebarez/sqlite 的底层 SQLite 绑定 |
| `github.com/go-playground/validator/v10` | v10.30.1 | Gin 的参数校验器 |
| `github.com/goccy/go-json` | v0.10.5 | 高性能 JSON 编解码 |
| `github.com/json-iterator/go` | v1.1.12 | 另一套高性能 JSON 实现 |
| `github.com/fsnotify/fsnotify` | v1.9.0 | Viper 的文件变更监听 |
| `github.com/go-viper/mapstructure/v2` | v2.4.0 | Viper 的结构体映射 |
| `github.com/spf13/afero` | v1.15.0 | Viper 的文件系统抽象层 |
| `github.com/spf13/cast` | v1.10.0 | Viper 的类型转换 |
| `github.com/jinzhu/inflection` | v1.0.0 | GORM 的单复数转换 |
| `github.com/jinzhu/now` | v1.1.5 | GORM 的时间处理 |
| `github.com/dustin/go-humanize` | v1.0.1 | SQLite 的可读大小格式化 |
| `github.com/google/uuid` | v1.3.0 | UUID 生成 |
| `github.com/remyoudompheng/bigfft` | v0.0.0 | SQLite 的大数运算 |
| `modernc.org/libc` | v1.22.5 | 纯 Go 的 libc 实现（SQLite 依赖） |
| `modernc.org/mathutil` | v1.5.0 | 纯 Go 的数学工具（SQLite 依赖） |
| `modernc.org/memory` | v1.5.0 | 纯 Go 的内存管理（SQLite 依赖） |
| `modernc.org/sqlite` | v1.23.1 | 纯 Go 的 SQLite 实现 |
| `golang.org/x/net` | v0.54.0 | Go 网络库扩展 |
| `golang.org/x/sys` | v0.45.0 | Go 系统调用扩展 |
| `golang.org/x/text` | v0.37.0 | Go 文本处理扩展 |

---

## 前端（React + TypeScript）

### 核心框架

| 包名 | 版本 | 说明 |
|------|------|------|
| `react` | ^19.2.6 | React 核心库 |
| `react-dom` | ^19.2.6 | React DOM 渲染器 |

### 构建工具

| 包名 | 版本 | 说明 |
|------|------|------|
| `vite` | ^8.0.14 | 前端构建工具，开发服务器 + 生产打包 |
| `@vitejs/plugin-react` | ^6.0.2 | Vite 的 React 插件（支持 React Fast Refresh） |
| `typescript` | ~6.0.3 | TypeScript 编译器 |

### UI 框架与样式

| 包名 | 版本 | 说明 |
|------|------|------|
| `antd` | ^6.4.3 | Ant Design 组件库，用于后台管理界面 |
| `tailwindcss` | ^4.3.0 | Tailwind CSS 原子化样式框架，用于前台页面定制 |
| `@tailwindcss/vite` | ^4.3.0 | Tailwind CSS 的 Vite 插件 |

### 路由

| 包名 | 版本 | 说明 |
|------|------|------|
| `react-router-dom` | ^7.15.1 | React 路由库，支持嵌套路由、动态参数 |

### 状态管理

| 包名 | 版本 | 说明 |
|------|------|------|
| `zustand` | ^5.0.13 | 轻量级状态管理库，用于认证状态持久化 |

### HTTP 客户端

| 包名 | 版本 | 说明 |
|------|------|------|
| `axios` | ^1.16.1 | HTTP 请求库，拦截器自动注入 JWT token |

### Markdown 编辑器与渲染

| 包名 | 版本 | 说明 |
|------|------|------|
| `@uiw/react-md-editor` | ^4.1.1 | Markdown 编辑器组件，用于文章编辑 |
| `react-markdown` | ^10.1.0 | Markdown 渲染为 React 组件，用于文章展示 |
| `remark-gfm` | ^4.0.1 | remark 的 GFM（GitHub Flavored Markdown）插件，支持表格/任务列表等 |
| `rehype-highlight` | ^7.0.2 | rehype 的代码高亮插件 |
| `highlight.js` | ^11.11.1 | 代码语法高亮库 |

### 日期处理

| 包名 | 版本 | 说明 |
|------|------|------|
| `dayjs` | ^1.11.20 | 轻量级日期处理库，用于时间格式化 |

### 代码质量

| 包名 | 版本 | 说明 |
|------|------|------|
| `eslint` | ^10.4.0 | JavaScript/TypeScript 代码检查工具 |
| `@eslint/js` | ^10.0.1 | ESLint 的 JavaScript 规则集 |
| `typescript-eslint` | ^8.59.4 | ESLint 的 TypeScript 支持 |
| `eslint-plugin-react-hooks` | ^7.1.1 | React Hooks 规则的 ESLint 插件 |
| `eslint-plugin-react-refresh` | ^0.5.2 | React Fast Refresh 的 ESLint 插件 |
| `globals` | ^17.6.0 | 全局变量定义（ESLint 配置使用） |

### 类型定义

| 包名 | 版本 | 说明 |
|------|------|------|
| `@types/react` | ^19.2.15 | React 的 TypeScript 类型定义 |
| `@types/react-dom` | ^19.2.3 | ReactDOM 的 TypeScript 类型定义 |
| `@types/node` | ^24.12.4 | Node.js 的 TypeScript 类型定义 |

---

## 开发环境

| 工具 | 版本要求 | 说明 |
|------|----------|------|
| Go | 1.26.3+ | 后端运行时 |
| Node.js | 18+ | 前端运行时 |
| npm | 9+ | 前端包管理器 |
