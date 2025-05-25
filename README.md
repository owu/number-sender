# number sender

Number-sender is a tool that categorizes auto-incrementing numbers based on predefined rules, designed for scenarios where unique numbers need to be assigned to entities (e.g., allocating user IDs during user registration, assigning room IDs to live streamers, or distributing group IDs to chat groups). In practical business scenarios, developer can obtain numbers by specifying types (e.g., assigning regular numbers to standard users, assigning premium numbers for paying users).

---

English | [中文](./README_cn.md) | [日本語](./README_jp.md)

## build
```
go mod vendor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o owu-number-sender-go main.go

```

## run
```
./owu-number-sender-go --config=config/config-dev.toml

```


## api
#### Check service status
- request:
```
curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/ping'
```

- response:
```
{"data":{"time":1747067149074},"error":0,"msg":"success"}
```
---

#### Query the count of various plans in the cache
- request:
```
curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/len'

```
- response:
```
{"data":{"starter":3472008,"standard":2237018,"premium":82317,"ultimate":5657},"error":0,"msg":"success"}
```

#### Retrieve a numeric value of a specified plan
- request:

```
# The value of {:plan} can be starter, standard, premium, or ultimate.

curl \
-H "Token:f98f84f7939a56f6ec8c42ef088139f5" \
-H "Milli:1746460799000" \
'http://localhost:8080/api/pop/{:plan}'
```
- response:
```
{"data":{"starter":0,"standard":0,"premium":0,"ultimate":10101},"error":0,"msg":"success"}
```

#### API Authentication Protocol
```
Header: Milli , timestamp in milliseconds
Header: Token , md5({Milli},{Encrypt})

---

Token generation example:
{Milli}, current timestamp in milliseconds, which is 174646079900 
{Encrypt}, is configured in the file config/config-*.toml, with the value of api.encrypt being 9b0b7484bc65e241804ce8eeb014f247. 

Then Token = MD5(1746460799000, 9b0b7484bc65e241804ce8eeb014f247) = f98f84f7939a56f6ec8c42ef088139f5
```

