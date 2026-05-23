# CLAUDE.md

本文件为 Claude Code（claude.ai/code）在此仓库中工作时提供指导。

## 常用命令

### 后端（Go）

```bash
# 安装依赖（国内需使用 Go 代理）
cd backend && GOPROXY=https://goproxy.cn,direct go mod tidy

# 启动开发服务器（监听 :8080 端口）
cd backend && go run ./cmd/server/

# 编译二进制文件
cd backend && go build -o server ./cmd/server/
```

### 前端（React + TypeScript）

```bash
cd frontend && npm install
npm run dev        # 开发服务器，端口 :5173，/api 代理到 localhost:8080
npm run build      # 类型检查 + 生产构建
npm run lint       # ESLint 检查
```

### 默认管理员账号

```
admin / admin123
```

## 项目架构

```
frontend/                          backend/
├── src/                           ├── cmd/server/main.go    # 入口：初始化数据库、自动建表、创建默认管理员
│   ├── api/client.ts              ├── config/               # Viper 配置加载（需要 mapstructure 标签）
│   ├── api/*.ts      # 各模块接口调用                    ├── internal/
│   ├── store/        # Zustand 状态管理（authSlice）        │   ├── handler/        # Gin 控制器（参数绑定 → 调用 service → 返回响应）
│   ├── router/       # React Router v7 + 路由守卫           │   ├── service/        # 业务逻辑层
│   ├── components/   # 通用组件（layout/markdown/comment）   │   ├── repository/     # 数据访问层（GORM）
│   ├── pages/        # public/（前台）admin/（后台）         │   ├── model/          # 数据库模型
│   ├── types/        # TypeScript 类型定义                   │   ├── dto/            # 请求/响应 DTO
│   └── hooks/        # 自定义 hooks                          │   ├── middleware/      # JWT 认证、CORS、日志、异常恢复
│                                                            │   ├── router/         # 路由注册 + 依赖注入
│                                                            │   └── utils/          # JWT 工具、统一响应格式
│                                                            └── config/config.yaml
```

**分层调用链**：`handler → service → repository → model（GORM）` —— 每层仅依赖下一层，不允许跨层调用。所有依赖在 `handler/init.go` 中手动注入。

**数据库**：使用 `github.com/glebarez/sqlite` 驱动（纯 Go 实现，无需 CGO）。数据库文件位于 `backend/data/blog.db`，首次运行时通过 GORM AutoMigrate 自动创建。

**API 路由**（`backend/internal/router/router.go`）：
- 公开路由 `/api/v1/*` —— 无需认证（文章、分类、标签、评论、登录）
- 需认证路由 `/api/v1/auth/profile`
- 管理端路由 `/api/v1/admin/*` —— 全部需要 JWT 认证，通过 `Authorization: Bearer <token>` 传递

**统一响应格式**：成功时返回 `{ code: 0, message: "ok", data }`，失败时 code 为非零。分页数据在 data 内使用 `{ list, total, page, page_size }` 格式。

## 重要注意事项

### Viper 配置必须加 `mapstructure` 标签

`config/config.go` 中的每个结构体字段都必须有 `mapstructure` 标签，与 YAML 中的 snake_case 键名对应，否则 Viper 会静默地将字段设为零值。

```go
type JWTConfig struct {
    Secret      string `mapstructure:"secret"`
    ExpireHours int    `mapstructure:"expire_hours"`
}
```

### SQLite 驱动的选择

必须使用 `github.com/glebarez/sqlite`（纯 Go 实现），不要使用 `gorm.io/driver/sqlite`。后者底层依赖 `mattn/go-sqlite3`，需要 CGO 和 C 编译器。

```go
import "github.com/glebarez/sqlite"
db, _ := gorm.Open(sqlite.Open("data/blog.db"), &gorm.Config{})
```

### 文章更新必须同时 Save 和 Replace 标签

更新文章及其多对多标签关联时，必须同时调用 `db.Save()`（保存文章字段）和 `Association("Tags").Replace()`（替换标签关联）。仅 Save 不会持久化标签变更。

### 前端 API 基础地址

`frontend/src/api/client.ts` 中的 Axios 实例直接使用 `http://localhost:8080/api/v1` 作为 baseURL（而非走 Vite 代理）。Vite 代理仅作为开发时的备用方案。

### Windows 上重启后端端口被占用

Windows 上 Go 进程退出后可能仍占用 8080 端口，用以下命令强制释放：
```powershell
Get-NetTCPConnection -LocalPort 8080 | ForEach-Object { Stop-Process -Id $_.OwningProcess -Force }
```

### JWT 认证流程

1. 前端 POST `/auth/login` 提交用户名和密码
2. 后端返回 `{ token, user }`
3. 前端将 token 存入 localStorage，Zustand authSlice 持久化用户信息
4. Axios 请求拦截器自动为每个请求附加 `Authorization: Bearer <token>` 头
5. 收到 401 响应时，拦截器清除本地存储并重定向到 `/admin/login`
