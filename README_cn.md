# number sender

Number-sender 是一款基于预设规则对自增数字进行分类管理的工具，专为需要为实体分配唯一数字的场景设计（例如：在用户注册时分配用户ID，为主播分配直播间ID，或为聊天群组分配群号等）。在实际业务场景中，开发人员可以通过指定类型获取号码（例如：为普通用户分配常规号码，为付费用户分配专属靓号）。

---

[English](./README.md) | 中文 | [日本語](./README_jp.md)

## 编译
```
go mod vendor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o owu-number-sender-go main.go

```

## 运行
```
./owu-number-sender-go --config=config/config-dev.toml

```


## api接口
#### 服务状态检查
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

#### 查询缓存中各套餐数量
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

#### 获取指定套餐的号码
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

#### API鉴权协议
```
请求头: Milli , 时间戳（毫秒格式）
请求头: Token , 由 md5({Milli},{Encrypt}) 生成

---

Token生成示例：
{Milli} 为当前毫秒时间戳，假设为 1746460799000 
{Encrypt} 配置于 config/config-*.toml 文件，假设 api.encrypt 值为 9b0b7484bc65e241804ce8eeb014f247 

则 Token = MD5(1746460799000,9b0b7484bc65e241804ce8eeb014f247) = f98f84f7939a56f6ec8c42ef088139f5
```

