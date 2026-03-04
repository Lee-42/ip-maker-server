# Inspo (灵感) API 使用示例文档

**API Base URL**: 根据实际部署环境配置，例如 `http://localhost:8000/api`

---

## 1. 创建灵感

**请求方式**: `POST`
**路径**: `/inspo`
**Content-Type**: `application/json`

### 1.1 场景一：创建文本灵感

**请求参数 (Body)**:

```json
{
  "title": "关于夏日海滩的新点子",
  "type": "text",
  "content": "可以在沙滩上画一些带有霓虹色彩的图案..."
}
```

### 1.2 场景二：创建图片灵感

**请求参数 (Body)**:

```json
{
  "title": "夏日海滩参考图",
  "type": "image",
  "content": "https://example.com/images/summer_beach.jpg"
}
```

**响应成功 (200 OK)**:

```json
{
  "code": 0,
  "message": "",
  "data": {
    "id": 1001
  }
}
```

---

## 2. 更新灵感

**请求方式**: `PUT`
**路径**: `/inspo/{id}` (例如 `/inspo/1001`)
**Content-Type**: `application/json`

### 2.1 场景一：修改内容

**请求参数 (Body)**:

```json
{
  "id": 1001,
  "title": "关于夏日海滩更酷的点子",
  "type": "text",
  "content": "加上赛博朋克的霓虹色彩会有奇效！"
}
```

### 2.2 场景二：将灵感标记为“已生产”（修改状态）

**请求参数 (Body)**:

```json
{
  "id": 1001,
  "status": 1
}
```

_(注意：在更新接口中，参数是可选的，只传 `status` 就会只更新状态)_

**响应成功 (200 OK)**:

```json
{
  "code": 0,
  "message": "",
  "data": {}
}
```

---

## 3. 获取灵感详情

**请求方式**: `GET`
**路径**: `/inspo/{id}` (例如 `/inspo/1001`)

**请求参数 (Query)**:
_(无，ID从路径中获取)_

**响应成功 (200 OK)**:

```json
{
  "code": 0,
  "message": "",
  "data": {
    "id": 1001,
    "title": "关于夏日海滩更酷的点子",
    "type": "text",
    "content": "加上赛博朋克的霓虹色彩会有奇效！",
    "status": 1,
    "createdAt": "2026-03-04 15:00:00",
    "updatedAt": "2026-03-04 15:10:00"
  }
}
```

---

## 4. 获取灵感列表

**请求方式**: `GET`
**路径**: `/inspo`

### 4.1 场景一：分页获取所有灵感

**请求参数 (Query)**:
`/inspo?page=1&size=20`

### 4.2 场景二：获取所有“未生产”的灵感

**请求参数 (Query)**:
`/inspo?page=1&size=20&status=0`

**响应成功 (200 OK)**:

```json
{
  "code": 0,
  "message": "",
  "data": {
    "total": 50,
    "list": [
      {
        "id": 1002,
        "title": "冬季滑雪系列企划案",
        "type": "text",
        "content": "计划推出一系列滑雪道具...",
        "status": 0,
        "createdAt": "2026-03-04 14:20:00",
        "updatedAt": "2026-03-04 14:20:00"
      },
      {
        "id": 1003,
        "title": "滑雪板材质参考图",
        "type": "image",
        "content": "https://example.com/images/snowboard.png",
        "status": 0,
        "createdAt": "2026-03-04 14:15:00",
        "updatedAt": "2026-03-04 14:15:00"
      }
    ]
  }
}
```

---

## 5. 删除灵感

**请求方式**: `DELETE`
**路径**: `/inspo/{id}` (例如 `/inspo/1001`)

**请求参数 (Query)**:
_(无，ID从路径中获取)_

**响应成功 (200 OK)**:

```json
{
  "code": 0,
  "message": "",
  "data": {}
}
```
