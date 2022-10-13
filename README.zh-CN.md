# Go Web Skeleton

一个完整的 Golang web 应用骨架。

> **[EN README](README.md)**

包含：

- 可用于API接口应用，CLI命令行应用，WEB应用
- 日志库logrus配置使用
- swagger API文档配置生成
- 多语言支持，视图渲染，请求数据验证
- 配置读取管理，根据环境加载，多文件支持
- 包含 redis, mysql, mongo 的初始化和简单使用
- 使用`go mod`来安装管理依赖库

## 项目结构

> Github地址 https://github.com/inhere/go-web-skeleton

```text
api/ API接口应用 handlers
 |- controller
 |_ middleware
app/ 公共目录(公共方法，应用初始化，公共组件等)
cmd/ CLI命令行应用 commands
 |_ cliapp/    命令行应用入口文件(main)
config/   应用配置目录(基础配置加各个环境配置)
model/  数据和逻辑代码目录
 |_ form/   请求表单结构数据定义，表单验证配置
 |_ logic/  逻辑处理
 |_ mongo/  MongoDB的数据集合模型定义
 |_ mysql/  MySQL的数据表单模型定义
 |_ rds/    Redis的数据模型定义
resource/ 一些项目使用到的非代码资源（语言文件，视图模板文件等）
runtime/      临时文件目录(文件缓存，日志文件等)
static/   静态资源目录（js,css等）
route.go  路由注册文件
Dockerfile Dockerfile
makefile  编写了一些通用的快捷命令，帮助打包，构建docker，生成文档，运行测试等等
... ...
```

## 使用的包

- http路由: [gookit/rux](https://github.com/gookit/rux) 
- 配置读取管理：
  - ini配置：[gookit/ini](https://github.com/gookit/ini)
  - 多种格式配置：[gookit/config](https://github.com/gookit/config)
- 日志记录：
  - [sirupsen/logrus](https://github.com/sirupsen/logrus)
    - 日志分割：[rifflock/lfshook](https://github.com/rifflock/lfshook)
    - 日志分割：[lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
  - [go.uber.org/zap](https://github.com/uber-go/zap)
- mysql等：
  - [go-xorm/xorm](https://github.com/go-xorm/xorm)
  - [jinzhu/gorm](https://github.com/jinzhu/gorm)
- mongodb: [github.com/globalsign/mgo](https://github.com/globalsign/mgo)
- cache: [gookit/cache](https://github.com/gookit/cache) 
- redis
  - [go-redis/redis](https://github.com/go-redis/redis)
  - [gomodule/redigo](https://github.com/gomodule/redigo/redis)
- language: [gookit/i18n](https://github.com/gookit/i18n)
- view渲染: [gookit/view](https://github.com/gookit/view)
- 命令行应用: [gookit/gcli](https://github.com/gookit/gcli)
- 表单数据验证:
  - [gookit/validate](https://github.com/gookit/validate) 
  - [go-playground/validator](https://github.com/go-playground/validator)
- 高性能的json序列化库: [json-iterator/go](https://github.com/json-iterator/go) 
- ~eureka client: [PDOK/go-eureka-client](https://github.com/PDOK/go-eureka-client)~ 未使用

### 辅助库

- swagger 文档生成：
  - [go-swagger](https://github.com/go-swagger/go-swagger) 文档复杂，但是功能更强大
  - [swaggo/swag](https://github.com/swaggo/swag) 文档和使用比较简单，仅生成文档足够用了
- 测试辅助库，方便快速断言 [stretchr/testify](https://github.com/stretchr/testify)
- 调试工具：[davecgh/go-spew](https://github.com/davecgh/go-spew) 深度打印golang变量数据

## 额外组件

- swagger UI: swagger文档渲染
- Dockerfile: 可用于生产环境的docker镜像构建脚本，基于alpine，构建完成一个项目镜像估计大小为：30 M 左右
- makefile: 已经内置一些快速使用的make命令，帮助快速生成文档，构建镜像

## 开始使用

- 首先，将骨架仓库clone到你的本机目录
- 重命名 `go-web-skeleton` 目录为你的项目名
- 进入到项目，将项目中 `github.com/inhere/go-web-skeleton` 替换为你的项目名(针对go文件)
- 再搜索将所有的`go-web-skeleton`替换为你的项目名（主要是Dockerfile,makefile里）
- 运行 `go mod tidy` 安装依赖库
- 运行项目：`go run main.go`

## swagger文档生成

安装：

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

> 使用请查看 `swaggo/swag` 的文档和示例

生成到指定目录下：

```bash
swag init -o static
# 同时会生成这个文件，不需要的可以删除掉
rm static/docs.go
```

注意：

> `swaggo/swag` 是从字段的注释解析字段描述信息的

```go
type SomeModel struct {
	// the name description
	Name   string `json:"name" example:"tom"`
}
```

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

## 参考

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## License

**MIT**
