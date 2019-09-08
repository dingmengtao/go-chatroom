# go语言实现的简易聊天室系统

原生纯手写框架，练手项目，从单文件逐渐融入分层思想，构建基于MVC架构思想的聊天室小工具，适合新手学习。

## 要求
- 安装redis，需要本地或服务器安装redis扩展

- 使用redigo扩展，项目内使用redigo扩展包需redis数据库操作
[redigo GitHub地址传送门](https://github.com/gomodule/redigo)
```
    go get github.com/gomodule/redigo/redis
```
- 如果使用go get无法完成下载，则可直接下载后放到 github.com/gomodule/ 该路径下

## 安装

进入GOPATH目录，从github上clone本项目即可。

## 使用

在GOPATH目录下分别执行如下代码生成可执行文件(默认项目放在GOPATH目录下)：

生成server端可执行文件：
```
go build -o server.exe .\chatroom\server\main
```

生活曾客户端可执行文件
```
go build -o client.exe .\chatroom\client\main
```

分别启动server.exe和client.exe即可使用该聊天室。