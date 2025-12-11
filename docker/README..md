目录说明，该目录下都是容器运行时的文件，不应该被修改，目录结构如下，目录文件适配/docker-compose.yml：
```
docker/
├── app/
│   ├── config/
│   │   └── config.toml  容器启动时自动拷贝自 config/config.toml
│   └── logs/  容器运行时的日志目录，映射到主机的 ./docker/app/logs 目录
├── redis/
│   ├── conf/
│   │   └── redis.conf  容器启动时Redis配置信息 config/redis.conf
│   └── data/
└── README.md  容器运行时的说明文档

```
