<h1 align="center">behappy</h1>

<div align="center">

![Build](https://github.com/Bpazy/behappy/workflows/Build/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Bpazy_behappy&metric=alert_status)](https://sonarcloud.io/dashboard?id=Bpazy_behappy)

Don't worry, be happy. 风力掀天浪打头，只须一笑不须愁。

本项目的目标将用户的比赛信息推送到 QQ 群中，供群友点评。

</div>

## 效果预览
![2](https://user-images.githubusercontent.com/9838749/123549535-1a3ce300-d79c-11eb-9996-12bf03ce6882.png)


## 使用
```shell
$ ./behappy --help
风力掀天浪打头，只须一笑不须愁

Usage:
  behappy [command]

Available Commands:
  help        Help about any command
  run         运行 behappy 主程序
  version     查看 behappy 版本号

Flags:
      --debug   Debug Mode
  -h, --help    help for behappy

Use "behappy [command] --help" for more information about a command.
```
第一次启动会自动初始化 schema, tables

配置文件 `~/.behappy.yaml`:
```yaml
mirai:
  botqq: 10001
steamapi:
  key: AXAFJQPDJV2312
datasource:
  url: USERNAME:PASSWORD@tcp(example-mysql.com:3306)/behappy?charset=utf8mb4&parseTime=True&loc=Local
```

## 开发
### 项目原理？
基于 opendota 扫描战绩，基于 mirai 提供 QQ 交互能力

### 如何编译？
> 请注意，编译具有以下依赖项:
> 1. Golang
> 2. Make

简单的执行 make 命令即可，编译的文件在 `bin` 目录下。
如果只需要编译到 windows 平台，可以执行命令: `make windows-amd64`。更详细的用法，你可以直接查看根目录中 `Makefile` 文件。


## 注意事项
如果想要使用“最佳劳模”功能，需要在系统中安装 simkai.ttf 字体。
