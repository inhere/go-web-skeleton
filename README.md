# go-web-skeleton

使用golang gin框架的应用骨架

## 包含应用

- api API接口应用
- cli CLI命令行应用
- web web应用

## 使用的包

- 路由器: [gookit/sux](https://github.com/gookit/sux) 
- ini配置：[gookit/ini](https://github.com/gookit/ini)
- 日志记录：[go.uber.org/zap](https://github.com/uber-go/zap)
- 日志记录：[sirupsen/logrus](https://github.com/sirupsen/logrus)
  - 日志分割：[rifflock/lfshook](https://github.com/rifflock/lfshook)
  - 日志分割：[lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
- mysql等：[go-xorm/xorm](https://github.com/go-xorm/xorm)
- mongodb: [github.com/globalsign/mgo](https://github.com/globalsign/mgo)
- cache, redis: [garyburd/redigo](https://github.com/garyburd/redigo/redis)
- language: [gookit/i18n](https://github.com/gookit/i18n)
- 表单数据验证: [go-playground/validator](https://github.com/go-playground/validator)
- 高性能的json序列化库: [json-iterator/go](https://github.com/json-iterator/go) 
- ~eureka client: [PDOK/go-eureka-client](https://github.com/PDOK/go-eureka-client)~ 未使用

### 辅助库

- `dep` 使用dep来安装管理依赖库
- swagger 文档生成：
  - go-swagger 
  - [swaggo/swag](https://github.com/swaggo/swag)
- 测试辅助库，方便快速断言 [stretchr/testify](https://github.com/stretchr/testify)
- 调试工具：[davecgh/go-spew](https://github.com/davecgh/go-spew) 深度打印golang变量数据

## 额外组件

- swagger UI: swagger文档渲染
- Dockerfile: 可用于生产环境的docker镜像构建脚本，基于alpine，构建完成一个项目镜像估计大小为：30 M 左右
- makefile: 已经内置一些快速使用的make命令，帮助快速生成文档，构建镜像

## 开始使用

- 首先，将骨架clone到 GOPATH 的 src下，重命名 `go-web-skeleton` 目录为你的项目名
- 进入到项目，将项目中 `github.com/inhere/go-web-skeleton` 替换为你的项目名(针对go文件)
- 再搜索将所有的`go-web-skeleton`替换为你的项目名（主要是Dockerfile,makefile里）
- 运行 `dep ensure` 安装依赖库到vendor
- 运行项目：`go run main.go`

## 使用帮助

- 运行测试

```bash
go test
// 输出覆盖率
go test -cover
```

- 格式化项目

```bash
go fmt ./...
```

- 运行GoLint检查

> 需先安装 GoLint

```bash
golint ./...
```

## License

**MIT**
