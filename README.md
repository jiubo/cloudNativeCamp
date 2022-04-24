本项目采用monorepo结构,多个微服务都在同一个代码库里面, 惯例是每个项目入口在`./cmd/<project>`.

### 运行
以运行HTTP SERVER为例

    go run cmd/httpServer/main.go