# number sender

Number-sender 是一款基于预设规则对自增数字进行分类管理的工具，专为需要为实体分配唯一数字的场景设计（例如：在用户注册时分配用户ID，为主播分配直播间ID，或为聊天群组分配群号等）。在实际业务场景中，开发人员可以通过指定类型获取号码（例如：为普通用户分配常规号码，为付费用户分配专属靓号）。

---

[English](./README.md) | 中文 | [日本語](./README_jp.md)

## 1. 编译
```
go mod vendor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o number-sender-go main.go

```

## 2. 运行
```
./number-sender-go --config=config/config-test.toml

配置文件中的 app.env 变量，线下环境为 test ，生成环境为 prod 。
在接口鉴权中间件 /internal/pkg/mware/auth.go 中进行了判定test不做鉴权，方便测试。
```


## 3. API接口
### 3.1 服务状态检查
- 请求示例:
```
curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/ping'
```

- 返回示例:
```
{"data":{"time":1747067149074},"error":0,"msg":"success"}
```
---

### 3.2 查询缓存中各套餐数量
- 请求示例:
```
curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/len'

```
- 返回示例:
```
{"data":{"starter":3472008,"standard":2237018,"premium":82317,"ultimate":5657},"error":0,"msg":"success"}
```

### 3.3 获取指定套餐的号码
- 请求示例:

```
# {:plan} 可替换为 starter, standard, premium, ultimate.

curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/pop/{:plan}'
```
- 返回示例:
```
{"data":{"starter":0,"standard":0,"premium":0,"ultimate":10101},"error":0,"msg":"success"}
```

### 3.4 API鉴权协议
```
请求头: Milli , 时间戳（毫秒格式）
请求头: Token , 由 md5({Milli},{Encrypt}) 生成

---

Token生成示例：
{Milli} 为当前毫秒时间戳，假设为 1746460799000 
{Encrypt} 配置于 config/config-*.toml 文件，假设 api.encrypt 值为 9b0b7484bc65e241804ce8eeb014f247 

则 Token = MD5(1746460799000,9b0b7484bc65e241804ce8eeb014f247) = f98f84f7939a56f6ec8c42ef088139f5
```


## 4. Docker 部署

### 4.0 项目准备

在进行Docker部署前，需要先克隆项目代码并切换到指定分支：

```bash
git clone https://github.com/owu/number-sender.git && cd number-sender && git checkout main && git pull && chmod 777 -R ./docker && chmod +x docker-tools.sh 
```

### 4.1 Docker 工具脚本

项目提供了一个集成化的 Docker 管理脚本 `docker-tools.sh`，用于简化 Docker 镜像构建、容器管理和镜像迁移等操作。

#### 功能特点
- 集成交付镜像构建、容器启停、日志查看和镜像迁移功能于一体
- 支持交互式选择和命令行参数两种使用方式
- 每个操作都显示实际执行的 Docker 命令，提高透明度
- 模块化设计，便于后期维护和扩展

#### 使用方法

##### 交互式模式
```bash
./docker-tools.sh
```

运行后会显示如下菜单，输入对应的字母选择操作：
```
========================================
Docker 工具脚本
========================================
请选择要执行的操作:
b) 构建Docker镜像 (docker build)
d) 停止Docker容器 (docker compose down)
l) 查看容器日志 (docker logs)
m) 镜像迁移 (docker save/load)
u) 启动Docker容器 (docker compose up -d)
直接回车退出脚本
```

##### 命令行参数模式
```bash
./docker-tools.sh [选项]
```

支持的选项：
- `b` : 构建Docker镜像
- `d` : 停止Docker容器
- `l` : 查看容器日志
- `m` : 执行镜像迁移功能（进入子菜单）
- `u` : 启动Docker容器
- `h` : 显示帮助信息

#### 各功能说明

1. **构建Docker镜像 (b)**
   - 检查镜像是否已存在，避免重复构建
   - 检查镜像是否正在被运行，确保构建安全
   - 自动更新 docker-compose.yml 中的镜像版本号

2. **启动Docker容器 (u)**
   - 使用 `docker compose up -d` 启动容器
   - 显示容器启动状态

3. **停止Docker容器 (d)**
   - 使用 `docker compose down` 停止容器
   - 显示当前运行的容器列表

4. **查看容器日志 (l)**
   - 使用 `docker logs number-sender` 查看容器日志

5. **镜像迁移 (m)**
   - 导出镜像到本地文件（使用 `docker save`）
   - 从本地文件导入镜像（使用 `docker load`）
   - 支持交互式选择导出或导入操作

#### 镜像迁移子菜单
当选择 `m` 选项后，会进入镜像迁移子菜单：
```
========================================
Docker镜像迁移功能
当前配置:
  镜像名称: number-sender
  镜像版本: 0.0.1
  完整镜像名: number-sender:0.0.1
  导出文件名: number-sender.0.0.1.tar
========================================
请选择要执行的操作:
1) 导出镜像 (docker save)
2) 导入镜像 (docker load)
直接回车退出镜像迁移功能
```

### 4.2 传统Docker部署方式（可选）

如果需要使用传统方式部署，也可以手动执行以下命令：

1. 构建镜像：
```bash
docker build -t number-sender:0.0.1 .
```

2. 启动容器：
```bash
docker compose up -d
```

3. 停止容器：
```bash
docker compose down
```

4. 查看日志：
```bash
docker logs number-sender
```

5. 镜像迁移：
```bash
# 导出镜像
docker save -o number-sender.0.0.1.tar number-sender:0.0.1

# 导入镜像
docker load -i number-sender.0.0.1.tar
```

