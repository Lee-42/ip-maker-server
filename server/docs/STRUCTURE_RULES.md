# 项目结构与开发规范 (Meme Server)

为了区分原有的 `hotgo` 基础服务与本项目（Meme Server）的特定业务逻辑，我们在各个层级引入了 `meme` 命名空间。所有新增的业务代码应严格遵循以下结构。

## 1. 目录结构概览

所有业务相关的代码必须放置在 `meme` 目录下，禁止直接在 `hotgo` 原生目录下混淆业务代码。

| 层级 | 路径 | 说明 |
| :--- | :--- | :--- |
| **API (Admin)** | `server/api/admin/meme/{module}/` | 后台管理接口定义 |
| **API (Client)** | `server/api/api/meme/{version}/{module}/` | 客户端接口定义 (需带版本号，如 v1) |
| **Controller** | `server/internal/controller/api/meme/{module}/` | 接口控制器实现 |
| **Service** | `server/internal/service/meme/` | 服务接口定义 |
| **Logic** | `server/internal/logic/meme/{module}/` | 业务逻辑实现 |

## 2. 详细规范

### 2.1 API 定义 (`server/api`)

*   **后台管理 (Admin)**:
    *   路径: `server/api/admin/meme/{模块名}/{功能名}.go`
    *   包名: `package {模块名}` 或 `package {功能名}`
    *   示例: `server/api/admin/meme/user/user.go`

*   **客户端 (Client)**:
    *   路径: `server/api/api/meme/{版本号}/{模块名}/{功能名}.go`
    *   包名: `package {模块名}`
    *   示例: `server/api/api/meme/v1/user/login.go`

### 2.2 控制器 (`server/internal/controller`)

*   路径: `server/internal/controller/api/meme/{模块名}/`
*   包名: `package {模块名}`
*   引用: 引用 `api` 定义时，使用完整的包路径。
*   示例: `server/internal/controller/api/meme/user/user.go`

### 2.3 服务接口 (`server/internal/service`)

*   路径: `server/internal/service/meme/`
*   文件: `{模块名}.go`
*   包名: `package meme`
*   说明: 在此目录下定义接口，并提供 `Register{服务名}` 和 `{服务名}()` 方法。
*   示例: `server/internal/service/meme/user.go`

### 2.4 业务逻辑 (`server/internal/logic`)

*   路径: `server/internal/logic/meme/{模块名}/`
*   文件: `{功能名}.go`
*   包名: `package {模块名}`
*   说明: 实现 `service` 中定义的接口，并在 `init()` 方法中调用 `meme.Register{服务名}` 进行注册。
*   注册: 需要在 `server/internal/logic/logic.go` 中添加匿名导入: `_ "hotgo/internal/logic/meme/{模块名}"`

### 2.5 路由注册 (`server/internal/router`)

*   **API 路由**: `server/internal/router/api.go`
*   **Admin 路由**: `server/internal/router/admin.go`
*   **说明**: 必须在路由文件中导入控制器的包，并将其绑定到路由组中，否则接口无法访问且 Swagger 文档不会生成。
*   示例:
    ```go
    import "hotgo/internal/controller/api/meme/comment"

    // ...
    group.Bind(
        comment.Comment, // 评论模块
    )
    ```

## 3. 示例：添加新模块 "Comment"

1.  **API**: 定义 `server/api/api/meme/v1/comment/create.go`
2.  **Service**: 定义 `server/internal/service/meme/comment.go` (interface `IComment`)
3.  **Logic**: 实现 `server/internal/logic/meme/comment/comment.go` (struct `sComment`)
4.  **Logic 注册**: 修改 `server/internal/logic/logic.go` 引入 `hotgo/internal/logic/meme/comment`
5.  **Controller**: 实现 `server/internal/controller/api/meme/comment/comment.go`
6.  **路由注册**: 修改 `server/internal/router/api.go` 注册 `comment.Comment` 控制器
