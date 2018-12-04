#SHELL = /bin/sh
.DEFAULT_GOAL := help
# 每行命令之前必须有一个tab键。如果想用其他键，可以用内置变量.RECIPEPREFIX 声明
# mac 下这条声明 没起作用 !!
.RECIPEPREFIX = >
.PHONY: all usage help clean

# 需要注意的是，每行命令在一个单独的shell中执行。这些Shell之间没有继承关系。
# - 解决办法是将两行命令写在一行，中间用分号分隔。
# - 或者在换行符前加反斜杠转义 \

##There some make command for the project
##

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | sed -e 's/: / /'

##Available Commands:

  all:        ## Run all the commands for the entire publishing process
all: route apidoc phar pbimg

  clean:      ## Clean all created artifacts
clean:
	git clean --exclude=.idea/ -fdx

  apidoc:     ## Generate swagger UI document json
apidoc:
	swag init -s static
	rm docs/docs.go

  pack:       ## Build and package the application
pack:
	# collect git info to current env config file.
	go build -o ./go-webx

  pbprod:     ## Build prod docker image and push to your hub
pbprod:
	go build ./cli/cliapp.go && ./cliapp git
	docker build -f Dockerfile -t your.dockerhub.com/go-webx --build-arg app_env=prod .
	docker push your.dockerhub.com/go-webx

  pbtest:     ## Build test docker image and push to your hub
pbtest:
	go build ./cli/cliapp.go && ./cliapp git
	docker build -f Dockerfile -t your.dockerhub.com/go-webx:test --build-arg app_env=test .
	docker push your.dockerhub.com/go-webx:test

  pbaudit:    ## Build audit docker image and push to your hub
pbaudit:
	go build ./cli/cliapp.go && ./cliapp git
	docker build -f Dockerfile -t your.dockerhub.com/go-webx:audit --build-arg app_env=audit .
	docker push your.dockerhub.com/go-webx:audit

  devimg:     ## Build dev docker image
devimg:
	docker build -f Dockerfile --build-arg app_env=dev -t go-webx:dev .

##
##Helper Commands:

  test:   ## Run all the tests
test: fmt lt

  echo:   ## echo test
echo:
	echo hello

  fmt:    ## Run the go fmt
fmt:
	go fmt ./...

  lt:     ## Run the golint tool
lt:
	go lint ./...

  tc:     ## Run the unit tests with code coverage
tc:
	go test -cover ./...
