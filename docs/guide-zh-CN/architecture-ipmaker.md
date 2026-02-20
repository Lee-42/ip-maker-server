# IP Maker 服务解耦与架构开发规范

此文档记录了 `ip` 相关功能从 Addon（独立插件）全面解耦，并接入项目标准化（`internal` 及 `api`）生命周期的演进和后续开发规范。

## 1. 架构改动与优化

我们通过移除独立插件结构，将其整合进标准的系统中，并完成后台 `Admin` 权限及前台 `API` 服务的物理隔离。相比之前：

- **数据结构解耦统合 (`appin`)**：全面废弃专用的 `ipin` 目录。将核心入参、出参统一整合进公共规范的 `internal/model/input/appin`，使得 C 端模块对外标准接口高度一致，利于后期拓展。
- **安全的权限分流 (`ApiAuth`)**：原有直接使用 `admin/admin` 作为后台控制器的设计，容易引发基于 Casbin 管理权限与普通用户越权的混淆。前台 API 现已被提取到 `internal/controller/api/ipmaker`，并接入标准 `ApiAuth` 控制，安全隔离性大幅增强。
- **最佳实践的路由注册 (`g.Meta`)**：接口完全接入框架的 `g.Meta` 标签自动嗅探挂载机制，舍弃了原始硬编码写死的路径绑定。
- **强内聚的服务暴露体系 (`Service`)**：规范定义 `internal/service/ipmaker/ip.go` 服务协议以及通过 `RegisterIp` 将逻辑端 `logic` 注入到接口中，确保控制层仅感知接口、不感知底层详细结构。

---

## 2. 目录功能与职责说明

开发新接口时，代码需按照职责被存放在以下分层体系中。

### 表现分发层（Api & Controller）

- **`api/api/ipmaker/v1/ip/ip.go`**:
  存放入参和出参的数据结构体。这是 HTTP 层收发数据的**最外层皮**，并在此通过 `g.Meta` 声明对应的请求方法（GET/POST）、路径和 Swagger 文档注解。
- **`internal/controller/api/ipmaker/ip/ip.go`**:
  负责处理与接收上述 `api` 定义进来的 HTTP 请求数据。提取上下文或 Session 中的鉴权信息，并将请求抛给底层服务处理，然后将响应按照 `v1` 结果直接包装返回。

### 业务组织与逻辑层（Service & Logic）

- **`internal/service/ipmaker/ip.go`**:
  接口生命周期。里面定义了对外暴漏能力的 `IIp` 协议及服务访问注册器 `Ip()` 。任何外界（如 `Controller` ）调用能力，均必须引向 此 `service/ipmaker`。
- **`internal/logic/ipmaker/ip/ip.go`**:
  实现 `IIp` 接口定义，所有具体的判断、计算、缓存更新都在这完成，并使用 `init()` 自动向 `service` 注册暴露自身能力。

### 数据流转层（Model）

- **`internal/model/entity/ip.go`**:
  数据库实体的精确映射。
- **`internal/model/input/appin/ip.go`**:
  定义内部从 `Controller` 传递向 `Logic` 的复合业务结构入参模型，不再同暴露给前台客户端的模型结构体混合复用，以保证服务层稳定。

---

## 3. 追加新接口标准流程 (SOP)

如果你需要在此模块上增加一个新的业务流（例如 `ExportIpList` / 导出 IP），请遵循以下开发工序：

### 第一步：在 `api/api/ipmaker/v1/ip/ip.go` 申明请求协议

明确你的路由路径及参数。

```go
// IpExportReq is the request for exporting ips.
type IpExportReq struct {
	g.Meta `path:"/ip/export" method:"get" tags:"IP前台管理" summary:"导出IP列表"`
	Type   string `json:"type" v:"required#缺少导出类型"`
}

type IpExportRes struct {
	DownloadUrl string `json:"downloadUrl"`
}
```

### 第二步：在 `internal/model/input/appin/ip.go` 中建构传输体

定义服务端系统接收到的通用信息结构。

```go
type IpExportInput struct {
    Type     string
    Operator int64 // 内部追加当前操作人
}
```

### 第三步：在 `internal/service/ipmaker/ip.go` 新增能力

```go
type IIp interface {
	...
	Export(ctx context.Context, in *appin.IpExportInput) (downloadUrl string, err error)
}
```

### 第四步：在 `internal/logic/ipmaker/ip/ip.go` 完成实现

实现上文的借口逻辑并对接至数据库（或缓存）系统。

```go
func (s *SIp) Export(ctx context.Context, in *appin.IpExportInput) (downloadUrl string, err error) {
    // 处理文件导出逻辑
    return "http://xxxxx/download.csv", nil
}
```

### 第五步：打通 Controller 转发 (`internal/controller/api/ipmaker/ip/ip.go`)

挂载方法接入 `v1` 数据和底层 Service 交互。

```go
func (c *cIp) Export(ctx context.Context, req *v1.IpExportReq) (res *v1.IpExportRes, err error) {
    url, err := ipmaker.Ip().Export(ctx, &appin.IpExportInput{
         Type: req.Type,
         // 补充鉴权用户
         Operator: contexts.GetUserId(ctx),
    })

    if err == nil {
       res = &v1.IpExportRes{DownloadUrl: url}
    }
    return
}
```

_（无需手动配置 router 路由，完成这五步，运行服务，`g.Meta` 将自动挂载至 `/api/ip/export`）_
