# go web demo

---

BB 两句
```note
error 处理真恶心啊
go 脚本滴神!
web 效率碾压了 java
反人类语法入门
一周入门这玩意 web 么得问题
```

## 使用到的技术

- gin
- gorm
- http
- mysql

## 使用方法
编译

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

启动

```shell
sudo ./main
```

debug

```shell
go mod download
go mod verify
go mod tidy
```

```shell
go run main.go
```

