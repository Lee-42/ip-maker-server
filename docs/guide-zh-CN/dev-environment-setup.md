# 开发环境配置指南

## 问题背景

在多台电脑上进行开发时，传统方式需要在每台电脑上重新安装和配置相同的开发环境（Go、MySQL、Redis 等），非常耗时且容易出错。本文档介绍几种现代化的解决方案，让你能够在任何电脑上快速启动开发环境。

## 环境需求

### 必需组件

| 组件            | 版本要求    | 用途           | 安装状态 |
| :-------------- | :---------- | :------------- | :------- |
| **Go**          | 1.18+       | 后端开发语言   | 需要安装 |
| **GoFrame CLI** | 2.0+        | 框架命令行工具 | 需要安装 |
| **Node.js**     | 18+         | 前端开发环境   | 需要安装 |
| **pnpm**        | 8+          | 前端包管理器   | 需要安装 |
| **MySQL**       | 5.7+ / 8.0+ | 数据库         | 需要安装 |
| **Redis**       | 5.0+        | 缓存和队列     | 建议安装 |

### 数据库配置要求

根据 `server/manifest/config/config.yaml` 配置：

```yaml
# MySQL 配置
database:
  default:
    link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo"

# Redis 配置
redis:
  default:
    address: "127.0.0.1:6379"
    db: "2"
    pass: ""
```

**配置说明**：

- MySQL 地址：`127.0.0.1:3306`
- MySQL 用户名：`hotgo`
- MySQL 密码：`hg123456.`
- MySQL 数据库：`hotgo`
- Redis 地址：`127.0.0.1:6379`

## 解决方案对比

| 方案               | 配置时间（首次） | 配置时间（换电脑） | 数据同步 | 适用场景     | 推荐度     |
| :----------------- | :--------------- | :----------------- | :------- | :----------- | :--------- |
| **传统安装**       | 30-60 分钟       | 30-60 分钟         | 手动     | 不推荐       | ⭐         |
| **Docker Compose** | 5 分钟           | 10 秒              | 自动     | 个人开发     | ⭐⭐⭐⭐   |
| **Dev Container**  | 3 分钟           | 5 秒               | 自动     | 团队协作     | ⭐⭐⭐⭐⭐ |
| **云端开发**       | 0 秒             | 0 秒               | 无需同步 | 随时随地开发 | ⭐⭐⭐⭐⭐ |

---

## 方案 1：Docker Compose（推荐）

### 核心思路

将所有依赖服务（MySQL、Redis）容器化，通过 Docker Compose 统一管理。只需在新电脑上安装 Docker，就能一条命令启动完整的开发环境。

### 准备工作

1. **安装 Docker Desktop**

访问 [Docker 官网](https://www.docker.com/products/docker-desktop/) 下载并安装。

```bash
# 验证安装
docker --version
docker-compose --version
```

### 配置步骤

#### 1. 创建开发环境配置

在项目根目录创建 `docker-compose.dev.yml`：

```yaml
version: "3.8"

services:
  # MySQL 数据库
  mysql:
    image: mysql:8.0
    container_name: hotgo-mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: root123456
      MYSQL_DATABASE: hotgo
      MYSQL_USER: hotgo
      MYSQL_PASSWORD: hg123456.
      TZ: Asia/Shanghai
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
      - ./storage/init:/docker-entrypoint-initdb.d # 自动执行初始化 SQL
    command:
      - --default-authentication-plugin=mysql_native_password
      - --character-set-server=utf8mb4
      - --collation-server=utf8mb4_unicode_ci
    healthcheck:
      test:
        [
          "CMD",
          "mysqladmin",
          "ping",
          "-h",
          "localhost",
          "-uhotgo",
          "-phg123456.",
        ]
      interval: 10s
      timeout: 5s
      retries: 5

  # Redis 缓存
  redis:
    image: redis:7-alpine
    container_name: hotgo-redis
    restart: unless-stopped
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --appendonly yes
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5

volumes:
  mysql_data:
    driver: local
  redis_data:
    driver: local
```

#### 2. 添加到版本控制

```bash
git add docker-compose.dev.yml
git commit -m "feat: 添加 Docker Compose 开发环境配置"
git push
```

### 使用流程

#### 第一次使用（任何电脑）

```bash
# 1. 克隆项目
git clone https://github.com/your-username/meme-server.git
cd meme-server

# 2. 启动开发环境
docker-compose -f docker-compose.dev.yml up -d

# 3. 查看服务状态
docker-compose -f docker-compose.dev.yml ps

# 4. 查看日志（可选）
docker-compose -f docker-compose.dev.yml logs -f mysql

# 5. 等待 MySQL 初始化完成（约 10-30 秒）
# 6. 启动前端
cd web
npm install  # 或 pnpm install
npm run dev

# 7. 启动后端
cd ../server
go mod download
make http
```

#### 换电脑后

```bash
# 1. 拉取最新代码
git pull

# 2. 启动开发环境（10 秒）
docker-compose -f docker-compose.dev.yml up -d

# 3. 启动项目
cd web && npm run dev
cd server && make http
```

#### 停止环境

```bash
# 停止但保留数据
docker-compose -f docker-compose.dev.yml stop

# 停止并删除容器（保留数据卷）
docker-compose -f docker-compose.dev.yml down

# 完全清理（包括数据）
docker-compose -f docker-compose.dev.yml down -v
```

### 常用命令

```bash
# 查看容器状态
docker-compose -f docker-compose.dev.yml ps

# 查看日志
docker-compose -f docker-compose.dev.yml logs -f

# 重启服务
docker-compose -f docker-compose.dev.yml restart

# 进入 MySQL 容器
docker exec -it hotgo-mysql mysql -uhotgo -phg123456. hotgo

# 进入 Redis 容器
docker exec -it hotgo-redis redis-cli

# 备份数据库
docker exec hotgo-mysql mysqldump -uhotgo -phg123456. hotgo > backup.sql

# 恢复数据库
docker exec -i hotgo-mysql mysql -uhotgo -phg123456. hotgo < backup.sql
```

### 优势

✅ **快速启动**：换电脑后 10 秒恢复环境  
✅ **环境一致**：所有开发者使用相同的环境配置  
✅ **无污染**：不影响本地系统环境  
✅ **易于清理**：一条命令删除所有数据  
✅ **自动初始化**：数据库自动导入初始数据

---

## 方案 2：VS Code Dev Container

### 核心思路

将整个开发环境（包括 Go、Node.js、MySQL、Redis、VS Code 插件）打包成容器，实现"一键启动完整开发环境"。

### 准备工作

1. **安装 Docker Desktop**
2. **安装 VS Code**
3. **安装 VS Code 扩展**：Remote - Containers

### 配置步骤

#### 1. 创建 Dev Container 配置

创建 `.devcontainer/devcontainer.json`：

```json
{
  "name": "HotGo Full Stack Development",
  "dockerComposeFile": "docker-compose.yml",
  "service": "workspace",
  "workspaceFolder": "/workspace",

  "forwardPorts": [8000, 3000, 3306, 6379],
  "portsAttributes": {
    "8000": { "label": "Backend API", "onAutoForward": "notify" },
    "3000": { "label": "Frontend Web", "onAutoForward": "notify" },
    "3306": { "label": "MySQL", "onAutoForward": "silent" },
    "6379": { "label": "Redis", "onAutoForward": "silent" }
  },

  "customizations": {
    "vscode": {
      "extensions": [
        "golang.go",
        "Vue.volar",
        "dbaeumer.vscode-eslint",
        "esbenp.prettier-vscode",
        "ms-azuretools.vscode-docker"
      ],
      "settings": {
        "go.toolsManagement.autoUpdate": true,
        "editor.formatOnSave": true
      }
    }
  },

  "postCreateCommand": "cd server && go mod download && cd ../web && pnpm install",

  "remoteUser": "vscode"
}
```

#### 2. 创建 Docker Compose 配置

创建 `.devcontainer/docker-compose.yml`：

```yaml
version: "3.8"

services:
  workspace:
    build:
      context: .
      dockerfile: Dockerfile
    volumes:
      - ../..:/workspaces:cached
    command: sleep infinity
    network_mode: service:mysql
    environment:
      - TZ=Asia/Shanghai

  mysql:
    image: mysql:8.0
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: root123456
      MYSQL_DATABASE: hotgo
      MYSQL_USER: hotgo
      MYSQL_PASSWORD: hg123456.
      TZ: Asia/Shanghai
    volumes:
      - mysql-data:/var/lib/mysql
    command: --default-authentication-plugin=mysql_native_password

  redis:
    image: redis:7-alpine
    restart: unless-stopped
    network_mode: service:mysql
    volumes:
      - redis-data:/data

volumes:
  mysql-data:
  redis-data:
```

#### 3. 创建 Dockerfile

创建 `.devcontainer/Dockerfile`：

```dockerfile
FROM mcr.microsoft.com/devcontainers/go:1.24

# 安装 Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# 安装 pnpm
RUN npm install -g pnpm

# 安装 GoFrame CLI
RUN go install github.com/gogf/gf/cmd/gf/v2@latest

# 安装常用工具
RUN apt-get update && apt-get install -y \
    mysql-client \
    redis-tools \
    && rm -rf /var/lib/apt/lists/*
```

### 使用流程

#### 第一次使用

1. **克隆项目**

   ```bash
   git clone https://github.com/your-username/meme-server.git
   ```

2. **在 VS Code 中打开项目**

   ```bash
   cd meme-server
   code .
   ```

3. **在容器中重新打开**
   - VS Code 会自动检测到 `.devcontainer` 配置
   - 点击提示框中的 "Reopen in Container"
   - 或按 `F1` → 输入 "Dev Containers: Reopen in Container"

4. **等待环境构建**（首次约 3-5 分钟）

5. **自动完成**：
   - Go 依赖下载
   - Node.js 依赖安装
   - MySQL、Redis 启动
   - VS Code 扩展安装

6. **开始开发**

#### 换电脑后

1. `git pull`
2. VS Code 打开项目
3. "Reopen in Container"
4. 等待 5 秒，环境就绪

### 优势

✅ **零配置**：打开项目即可开发  
✅ **环境标准化**：包含所有工具和扩展  
✅ **与生产一致**：开发环境和生产环境完全一致  
✅ **团队协作**：所有成员使用相同的配置  
✅ **秒级切换**：换电脑后 5 秒恢复

---

## 方案 3：云端开发环境

### 选项 A：GitHub Codespaces

#### 配置步骤

在项目中创建 `.devcontainer/devcontainer.json`（与方案 2 相同）。

#### 使用流程

1. **在 GitHub 仓库页面**
2. **点击 Code → Codespaces → Create codespace on main**
3. **等待 2-3 分钟**，自动启动完整开发环境
4. **在浏览器中开发**，或在本地 VS Code 中连接

#### 优势

✅ **任何设备**：包括 iPad、低配电脑  
✅ **零安装**：只需要浏览器  
✅ **自动备份**：代码云端存储  
✅ **免费额度**：60 小时/月（个人账号）

#### 定价

- 免费：60 小时/月（2 核）
- Pro：90 小时/月（2 核）
- 按需付费：$0.18/小时（2 核）

### 选项 B：自建云开发环境

#### 在云服务器上部署

```bash
# 1. SSH 连接到云服务器
ssh user@your-server.com

# 2. 安装 code-server（基于浏览器的 VS Code）
curl -fsSL https://code-server.dev/install.sh | sh

# 3. 启动 code-server
export PASSWORD="your-password"
code-server --bind-addr 0.0.0.0:8081

# 4. 克隆项目
git clone https://github.com/your-username/meme-server.git
cd meme-server

# 5. 启动开发环境
docker-compose -f docker-compose.dev.yml up -d
```

#### 本地通过浏览器访问

访问：`http://your-server.com:8081`

#### 或使用 VS Code Remote SSH

1. 安装 VS Code 扩展：Remote - SSH
2. 连接到服务器
3. 打开项目文件夹
4. 开始开发

#### 优势

✅ **数据安全**：代码和数据在自己的服务器上  
✅ **高性能**：可以选择高配服务器  
✅ **随时访问**：任何地方都能继续开发  
✅ **成本可控**：轻量应用服务器月费 30-100 元

---

## 推荐方案

### 个人开发者

**推荐：Docker Compose 方案**

理由：

- 配置简单，5 分钟搞定
- 换电脑后 10 秒恢复
- 完全免费
- 数据在本地，安全可控

### 团队协作

**推荐：VS Code Dev Container**

理由：

- 环境完全标准化
- 新成员零配置入职
- 自动安装扩展和工具
- 与 CI/CD 环境一致

### 移动办公

**推荐：GitHub Codespaces 或自建云环境**

理由：

- 任何设备都能开发
- 不受本地性能限制
- 数据自动备份
- 随时随地继续工作

---

## 数据迁移和备份

### 导出开发数据

```bash
# 导出 MySQL 数据
docker exec hotgo-mysql mysqldump -uhotgo -phg123456. hotgo > backup-$(date +%Y%m%d).sql

# 备份到 Git（建议）
mkdir -p storage/backup
mv backup-*.sql storage/backup/
git add storage/backup/
git commit -m "backup: 数据库备份"
git push
```

### 导入到新环境

```bash
# 方式 1：通过 docker exec
docker exec -i hotgo-mysql mysql -uhotgo -phg123456. hotgo < backup.sql

# 方式 2：放入初始化目录
mkdir -p storage/init
cp backup.sql storage/init/
docker-compose -f docker-compose.dev.yml up -d
```

---

## 故障排查

### Docker 相关

**问题：端口被占用**

```bash
# 查看端口占用
lsof -i :3306
lsof -i :6379

# 修改端口映射
# 编辑 docker-compose.dev.yml，将 3306:3306 改为 3307:3306
```

**问题：MySQL 连接失败**

```bash
# 查看 MySQL 日志
docker-compose -f docker-compose.dev.yml logs mysql

# 等待 MySQL 完全启动
docker-compose -f docker-compose.dev.yml ps
# 等待 health 状态变为 healthy
```

**问题：数据丢失**

```bash
# 检查数据卷
docker volume ls | grep mysql

# 数据卷位置
docker volume inspect meme-server_mysql_data
```

### Go 项目相关

**问题：找不到配置文件**

```bash
# 复制配置文件
cd server
cp manifest/config/config.example.yaml manifest/config/config.yaml
```

**问题：数据库连接被拒绝**

检查 `server/manifest/config/config.yaml`：

```yaml
database:
  default:
    link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo"
```

确保：

- 地址为 `127.0.0.1:3306`（Docker 已映射到本地）
- 用户名、密码、数据库名与 `docker-compose.dev.yml` 一致

---

## 扩展阅读

- [Docker Compose 官方文档](https://docs.docker.com/compose/)
- [VS Code Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers)
- [GitHub Codespaces](https://github.com/features/codespaces)
- [GoFrame 框架文档](https://goframe.org/)
