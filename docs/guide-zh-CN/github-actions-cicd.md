# GitHub Actions CI/CD 自动化部署

本项目使用 GitHub Actions 实现全自动化的 Docker 镜像构建和发布流程。

## 工作流文件位置

[`.github/workflows/docker-build.yml`](file://../../.github/workflows/docker-build.yml)

## 工作流概览

**名称**：`Build and Push Docker Images`（构建和推送 Docker 镜像）

**作用**：自动检测代码变更，智能构建并发布 Docker 镜像到 GitHub Container Registry (GHCR)。

## 触发条件

工作流会在以下情况自动运行：

| 触发条件             | 说明                           | 示例                                |
| :------------------- | :----------------------------- | :---------------------------------- |
| **推送到 main 分支** | 提交代码到 main 分支时自动构建 | `git push origin main`              |
| **创建版本标签**     | 创建 `v*` 格式的标签时         | `git tag v1.0.0 && git push --tags` |
| **手动触发**         | 在 GitHub Actions 页面手动运行 | Actions → 选择工作流 → Run workflow |

## 环境变量

工作流定义了以下镜像仓库地址：

```yaml
REGISTRY: ghcr.io
WEB_IMAGE: ghcr.io/polpo-space/wownow-server-web
SERVER_IMAGE: ghcr.io/polpo-space/wownow-server
```

## 工作流程（3 个 Jobs）

### Job 1: `changes` - 智能变更检测

**作用**：检测哪些目录的代码发生了变化，避免不必要的构建。

```yaml
filters:
  web:
    - "web/**"
  server:
    - "server/**"
```

**优势**：

- ✅ 只改了前端代码？不会重新构建后端镜像
- ✅ 只改了后端代码？不会重新构建前端镜像
- ✅ 节省构建时间和 CI 资源

### Job 2: `build-web` - 构建前端镜像

**运行条件**：仅当 `web/` 目录有代码变更时执行

**构建步骤**：

1. **准备环境**
   - 检出代码
   - 设置 Node.js 20
   - 安装 pnpm 8
   - 缓存 pnpm 依赖（加速后续构建）

2. **Docker 构建**
   - 设置 Docker Buildx（支持多平台构建）
   - 登录 GitHub Container Registry
   - 提取镜像元数据（自动生成标签）
   - 构建并推送镜像到 GHCR

**构建配置**：

```yaml
context: ./web
dockerfile: ./web/Dockerfile
platform: linux/amd64
cache: GitHub Actions Cache
```

### Job 3: `build-server` - 构建后端镜像

**运行条件**：仅当 `server/` 目录有代码变更时执行

**构建步骤**：

1. **准备环境**
   - 检出代码
   - 设置 Docker Buildx
   - 登录容器仓库

2. **Docker 构建**
   - 提取镜像元数据
   - **多层缓存策略**：
     - GitHub Actions 缓存（`type=gha`）
     - 远程镜像缓存（`buildcache` 标签）
   - 传入构建参数（Git commit hash）
   - 构建并推送镜像

**构建配置**：

```yaml
context: ./server
dockerfile: ./server/Dockerfile
platform: linux/amd64
build-args:
  - GIT_HASH=${{ github.sha }}
  - BUILDKIT_INLINE_CACHE=1
```

## 镜像标签策略

GitHub Actions 会根据触发条件自动生成镜像标签：

### 推送到 main 分支

```bash
ghcr.io/polpo-space/wownow-server-web:latest
ghcr.io/polpo-space/wownow-server-web:main

ghcr.io/polpo-space/wownow-server:latest
ghcr.io/polpo-space/wownow-server:main
```

### 创建版本标签（如 `v1.2.3`）

```bash
# Web 镜像
ghcr.io/polpo-space/wownow-server-web:1.2.3
ghcr.io/polpo-space/wownow-server-web:1.2
ghcr.io/polpo-space/wownow-server-web:1
ghcr.io/polpo-space/wownow-server-web:latest

# Server 镜像
ghcr.io/polpo-space/wownow-server:1.2.3
ghcr.io/polpo-space/wownow-server:1.2
ghcr.io/polpo-space/wownow-server:1
ghcr.io/polpo-space/wownow-server:latest
```

**标签规则**：

- `latest`：main 分支的最新构建
- `main`：main 分支标识
- `1.2.3`：完整语义化版本
- `1.2`：主版本 + 次版本
- `1`：主版本

## 缓存策略

为了加速构建，工作流使用了多层缓存：

### 前端缓存

- **pnpm 依赖缓存**：缓存 `node_modules`，避免重复下载
- **Docker 层缓存**：GitHub Actions 缓存（`type=gha`）

### 后端缓存

- **GitHub Actions 缓存**：缓存 Docker 构建层
- **远程镜像缓存**：使用 `buildcache` 标签存储缓存
- **内联缓存**：`BUILDKIT_INLINE_CACHE=1` 启用内联缓存

**效果**：首次构建可能需要 5-10 分钟，后续构建通常只需 1-2 分钟！

## 使用方法

### 自动构建和发布

1. **开发代码**

   ```bash
   # 修改前端代码
   cd web
   # ... 编辑代码 ...
   ```

2. **提交并推送**

   ```bash
   git add .
   git commit -m "feat: 添加新功能"
   git push origin main
   ```

3. **GitHub Actions 自动运行**
   - 自动检测到 `web/` 目录变更
   - 只构建前端镜像（跳过后端）
   - 推送到 GHCR，打上 `latest` 和 `main` 标签

### 发布版本

1. **创建版本标签**

   ```bash
   git tag v1.2.3
   git push --tags
   ```

2. **自动构建多个标签**
   - `1.2.3`（完整版本）
   - `1.2`（次版本）
   - `1`（主版本）
   - `latest`（最新版）

### 手动触发构建

1. 访问 GitHub 仓库
2. 点击 **Actions** 标签
3. 选择 **Build and Push Docker Images**
4. 点击 **Run workflow**
5. 选择分支并运行

## 查看构建状态

### GitHub Actions 页面

访问：`https://github.com/polpo-space/wownow-server/actions`

可以查看：

- ✅ 构建成功/失败状态
- 📊 构建日志
- ⏱️ 构建时间
- 📦 生成的镜像标签

### 镜像仓库

访问：`https://github.com/orgs/polpo-space/packages`

可以查看：

- 📦 所有已发布的镜像版本
- 🏷️ 镜像标签列表
- 📁 镜像大小
- 📅 发布时间

## 在服务器上使用镜像

### 拉取最新镜像

```bash
# 需要先登录 GHCR
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin

# 拉取 latest 标签
docker pull ghcr.io/polpo-space/wownow-server:latest
docker pull ghcr.io/polpo-space/wownow-server-web:latest
```

### 使用特定版本

```bash
# 拉取指定版本
docker pull ghcr.io/polpo-space/wownow-server:1.2.3
docker pull ghcr.io/polpo-space/wownow-server-web:1.2.3
```

### 在 docker-compose.yml 中使用

```yaml
version: "3.8"

services:
  server:
    image: ghcr.io/polpo-space/wownow-server:latest
    # ...

  web:
    image: ghcr.io/polpo-space/wownow-server-web:latest
    # ...
```

## 关键优势

✅ **自动化**：提交代码即自动构建，无需手动操作  
✅ **智能检测**：只构建有变更的部分，节省时间  
✅ **多版本管理**：自动生成语义化版本标签  
✅ **缓存加速**：多层缓存策略，大幅提升构建速度  
✅ **安全发布**：自动推送到 GitHub Container Registry  
✅ **可追溯性**：每个镜像包含 Git commit hash，方便回溯  
✅ **零停机部署**：结合 Portainer 可实现蓝绿部署

## 故障排查

### 构建失败

**查看日志**：

1. 进入 GitHub Actions 页面
2. 点击失败的工作流运行
3. 查看具体步骤的错误日志

**常见问题**：

- **Docker 构建失败**：检查 Dockerfile 语法
- **依赖安装失败**：检查 `package.json` 或 `requirements.txt`
- **权限错误**：确保 `GITHUB_TOKEN` 有 `packages: write` 权限

### 镜像推送失败

**原因**：

- GHCR 认证失败
- 仓库权限不足

**解决方法**：

1. 确保仓库设置中启用了 "Actions" 权限
2. 检查工作流的 `permissions` 配置

### 缓存失效

**现象**：每次构建都很慢，缓存似乎没生效

**解决方法**：

```bash
# 在 GitHub Actions 页面手动清除缓存
Settings → Actions → Caches → Delete all caches
```

## 扩展阅读

- [GitHub Actions 官方文档](https://docs.github.com/en/actions)
- [GitHub Container Registry 使用指南](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
- [Docker Buildx 文档](https://docs.docker.com/buildx/working-with-buildx/)
- [语义化版本规范](https://semver.org/lang/zh-CN/)
