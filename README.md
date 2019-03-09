# go-wex-skeleton

golang web应用骨架

主要用于：

- api API接口应用
- cmd CLI命令行应用
- web web应用

## 项目结构

```text
api/ API接口应用 handlers
app/ 公共目录(公共方法，应用初始化，公共组件等)
cmd/ CLI命令行应用 commands
 |_ bin/    命令行应用入口文件(main)
config/   应用配置目录（基础配置加各个环境配置）
controller
 |_ api 
 |_ web
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

> 参考 https://github.com/golang-standards/project-layout

## 使用的包

- http路由: [gookit/rux](https://github.com/gookit/rux) 
- 配置读取管理：
  - ini配置：[gookit/ini](https://github.com/gookit/ini)
  - 多种格式配置：[gookit/config](https://github.com/gookit/config)
- 日志记录：[go.uber.org/zap](https://github.com/uber-go/zap)
- 日志记录：[sirupsen/logrus](https://github.com/sirupsen/logrus)
  - 日志分割：[rifflock/lfshook](https://github.com/rifflock/lfshook)
  - 日志分割：[lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
- mysql等：[go-xorm/xorm](https://github.com/go-xorm/xorm)
- mongodb: [github.com/globalsign/mgo](https://github.com/globalsign/mgo)
- cache: [gookit/cache](https://github.com/gookit/cache) 
- redis: [gomodule/redigo](https://github.com/gomodule/redigo/redis)
- language: [gookit/i18n](https://github.com/gookit/i18n)
- view渲染: [gookit/view](https://github.com/gookit/view)
- 命令行应用: [gookit/gcli](https://github.com/gookit/gcli)
- 表单数据验证:
  - [gookit/validate](https://github.com/gookit/validate) 
  - [go-playground/validator](https://github.com/go-playground/validator)
- 高性能的json序列化库: [json-iterator/go](https://github.com/json-iterator/go) 
- ~eureka client: [PDOK/go-eureka-client](https://github.com/PDOK/go-eureka-client)~ 未使用

### 辅助库

- `dep` 使用dep来安装管理依赖库
- swagger 文档生成：
  - go-swagger 文档复杂，功能更强大
  - [swaggo/swag](https://github.com/swaggo/swag) 文档和使用比较简单，仅生成文档足够用了
- 测试辅助库，方便快速断言 [stretchr/testify](https://github.com/stretchr/testify)
- 调试工具：[davecgh/go-spew](https://github.com/davecgh/go-spew) 深度打印golang变量数据

## 额外组件

- swagger UI: swagger文档渲染
- Dockerfile: 可用于生产环境的docker镜像构建脚本，基于alpine，构建完成一个项目镜像估计大小为：30 M 左右
- makefile: 已经内置一些快速使用的make命令，帮助快速生成文档，构建镜像

## 开始使用

- 首先，将骨架clone到 GOPATH 的 src下，重命名 `go-wex-skeleton` 目录为你的项目名
- 进入到项目，将项目中 `github.com/inhere/go-wex-skeleton` 替换为你的项目名(针对go文件)
- 再搜索将所有的`go-wex-skeleton`替换为你的项目名（主要是Dockerfile,makefile里）
- 运行 `dep ensure` 安装依赖库到vendor
- 运行项目：`go run main.go`

## swagger文档生成

安装：

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

> 使用请查看 `swaggo/swag` 的文档和示例

生成到指定目录下：

```bash
swag init -s static
# 同时会生成这个文件，不需要可以删除掉
rm docs/docs.go
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

## License

**MIT**
