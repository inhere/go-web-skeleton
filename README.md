# Go Web Skeleton

A complete Golang web application skeleton.

> **[中文说明](README.zh-CN.md)**

Contains:

- Can be used for API interface applications, CLI command line applications, WEB applications
- Log library `gookit/slog` configuration used
- swagger API documentation configuration generation
- Multi-language support, view rendering, request data validation
- Configure read management, load according to environment, multi-file support
- Contains redis, mysql, mongo for initialization and simple use
- Use `go mod` to install the management dependencies

## Project Structure

> Github Project https://github.com/inhere/go-web-skeleton

```text
api/ API interface application handlers
  |- controller/
  |- middleware/
  |_ routes.go
app/ Common directory (public methods, application initialization, public components, etc.)
cmd/ CLI command line application commands
  |_ cliapp/ command line application entry file (main)
config/ Application configuration directory (basic configuration plus various environment configurations)
model/  Data and logic code directory
  |- form/  Request form structure data definition, form validation configuration
  |- logic/ Logic processing
  |- mongo/ MongoDB data collection model definition
  |- mysql/ MySQL data form model definition
  |_ rds/   Redis data model definition
resource/   Non-code resources used by some projects (language files, view template files, etc.)
runtime/    Temporary file directory (file cache, log files, etc.)
static/     Static resource directory (js, css, etc.)
main.go     Web application entry file
Dockerfile  Dockerfile
Makefile    Has written some common shortcut commands to help package, build docker, generate documentation, run tests, etc.
...
```

> render by `tree -d 2 ./`

## Go Packages

> use ✅ mark current used go package

### Http service

Provide HTTP service and routing

- [gookit/rux](https://github.com/gookit/rux) ✅

**Other**:

### CLI application

Command line application

- [gookit/gcli](https://github.com/gookit/gcli) ✅

### Configuration

Configuration management

- Multiple format configurations: [gookit/config](https://github.com/gookit/config) ✅
  - INI configuration: [gookit/ini](https://github.com/gookit/ini)

### Logging

  - [gookit/slog](https://github.com/gookit/slog) ✅
  - [sirupsen/logrus](https://github.com/sirupsen/logrus)
    - Log splitting: [rifflock/lfshook](https://github.com/rifflock/lfshook)
    - Log splitting: [lestrrat-go/file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)
  - [go.uber.org/zap](https://github.com/uber-go/zap)

### Database ORM

Mysql:

- [go-xorm/xorm](https://github.com/go-xorm/xorm)
- [jinzhu/gorm](https://github.com/jinzhu/gorm) ✅

Mongodb:

- [github.com/globalsign/mgo](https://github.com/globalsign/mgo)

### Cache

- Cache: [gookit/cache](https://github.com/gookit/cache) ✅

Redis:

- [go-redis/redis](https://github.com/go-redis/redis) ✅
- [gomodule/redigo](https://github.com/gomodule/redigo/redis)

### Data validation

Request data validation

- [gookit/validate](https://github.com/gookit/validate) ✅
- [go-playground/validator](https://github.com/go-playground/validator)

### Data serialization

High-performance  serialization library

JSON:

- [json-iterator/go](https://github.com/json-iterator/go)

### Other packages

- I18n language: [gookit/i18n](https://github.com/gookit/i18n) ✅
- View rendering: [gookit/easytpl](https://github.com/gookit/easytpl) ✅

Config register center:

- eureka client: [PDOK/go-eureka-client](https://github.com/PDOK/go-eureka-client) Not used
- nacos client:

## Extra Compoents

### Auxiliary Library

- swagger document generation:
  - [go-swagger](https://github.com/go-swagger/go-swagger) The documentation is complex but more powerful
  - [swaggo/swag](https://github.com/swaggo/swag) Documents and usage are relatively simple, only generating documents is enough
- Test the auxiliary library for quick assertion [stretchr/testify](https://github.com/stretchr/testify)
- Debugging tool: [davecgh/go-spew](https://github.com/davecgh/go-spew) Deep printing golang variable data

### Additional components

- swagger UI: swagger document rendering
- `Dockerfile`: docker image build script for production environment, based on alpine, build a project image with an estimated size of around 30 M
- `makefile`: Some quick-on make commands have been built to help quickly generate documentation and build images.

## Start

- First, clone the skeleton repository to your local directory
- Rename the `go-web-skeleton` directory to your project name.
- Go to the project and replace `github.com/inhere/go-web-skeleton` with your project name (for go file)
- Search again and replace all `go-web-skeleton` with your project name (mainly Dockerfile, makefile)
- Run `go mod tidy` to install dependent libraries
- Run the project: `go run main.go`

### Init project

```shell
go run ./cmd/appinit
```

### Swagger Docs Generation

installation:

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

> Please check the documentation and examples of `swaggo/swag`

Generated to the specified directory:

```bash
swag init -o static
# This file will be generated at the same time. It can be deleted if it is not needed.
rm static/docs.go
```

Notice:

> `swaggo/swag` is the parsing field description information from the comment of the field

```go
type SomeModel struct {
	// the name description
	Name   string `json:"name" example:"tom"`
}	
```

## Help

- Run the test

```bash
Go test
// output coverage
Go test -cover
```

- Formatting project

```bash
gofmt -s -l ./
go fmt ./...
```

- Run GoLint check

> Note: You need to install `GoLint` first.

```bash
golint ./...
```

## Refer

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## License

**MIT**
