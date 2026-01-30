# 数据库恢复操作指南

本文档记录了当 MySQL 服务出现异常（如被攻击、数据丢失）时，如何重新启动容器并恢复表结构的操作流程。

## 场景描述

- **环境**：远程 Docker 环境 (e.g., `meme-dev`)
- **问题**：MySQL 容器丢失或数据卷被破坏/污染。
- **目标**：重启 MySQL 服务，并确保数据库表结构已正确生成。

## 恢复步骤

### 1. 检查当前状态

首先确认 MySQL 容器和数据卷是否确实存在或损坏。

```bash
# 检查容器
docker ps -a | grep mysql

# 检查数据卷
docker volume ls | grep mysql
```

如果容器或数据卷已被恶意删除（例如被攻击后），则可以直接进行下一步。如果存在但已损坏，建议先手动删除：

```bash
# 停止并删除容器
docker stop meme-mysql
docker rm meme-mysql

# 删除数据卷 (警告：将清除所有数据)
docker volume rm meme-server_mysql_data
```

### 2. 启动服务

使用 Docker Compose 启动 MySQL 服务。服务启动时，Docker 会自动检测到由于数据卷为空（或新建），从而执行挂载目录下的初始化脚本（通常位于 `storage/init/`）。

```bash
docker compose -f docker-compose.dev.yml up -d mysql
```

> **注意**：`docker-compose.dev.yml` 中必须配置了初始化脚本的挂载，例如：
>
> ```yaml
> volumes:
>   - ./storage/init:/docker-entrypoint-initdb.d
> ```

### 3. 验证恢复结果

#### 查看初始化日志

检查容器日志，确认 `hotgo.sql` (或其他初始化脚本) 已被成功执行：

```bash
docker logs meme-mysql
```

应看到类似如下的日志输出：

```log
[Note] [Entrypoint]: /usr/local/bin/docker-entrypoint.sh: running /docker-entrypoint-initdb.d/hotgo.sql
...
[Note] [Entrypoint]: MySQL init process done. Ready for start up.
```

#### 检查数据库表

连接数据库，查询表列表以确结构已生成：

```bash
docker exec -it meme-mysql mysql -u leo -pjmqs18666 -D meme -e "SHOW TABLES;"
```

如果输出大量表名（如 `hg_admin_member`, `hg_sys_config` 等），则说明恢复成功。

## 常见问题

- **Q: 为什么重启没有生成表？**
  - A: 如果旧的数据卷 `meme-server_mysql_data` 仍然存在且包含数据，MySQL 会跳过初始化脚本的执行。必须先删除旧数据卷才能重新初始化。

- **Q: 生产环境如何处理？**
  - A: 生产环境**严禁**直接删除数据卷。应优先尝试从备份恢复。本文档流程主要适用于开发/测试环境，或数据已完全丢失且无备份的紧急重建场景。
