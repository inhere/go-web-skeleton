# Dir refer

## `/assets`

一些资源 (images, logos, etc).

## 插件 `/addons|plugin`

## `/build`

打包和持续集成脚本

## 文档 `/docs`    

## `/githooks`

Git hooks.

## `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

## `/pkg`     

引入的外部包

## 工具 `/tools|utils`    

该项目的支持工具。请注意，这些工具可以从 `/pkg` 和 `/internal` 目录导入代码。

## 脚本目录 `/scripts`

脚本执行各种构建，安装，分析等操作。

这些脚本使根级别Makefile保持简洁。

## `/web`

```text
/web
    /app
    /api
    /static
    /template
```

特定于Web应用程序的组件：静态Web资产，服务器端模板和SPA

## `/website`

如果您不使用Github页面，这是放置项目的网站数据的地方
