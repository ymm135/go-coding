## 第5章 Go语言的依赖管理
### 5-1 依赖管理
- 依赖管理经历三个阶段: GOPATH, GOVENDOR, go mod

### 5-2 GOPATH 和 GOVENDOR
**GOPATH**
- 默认在~/go路径, 所有的import都放在gopath下

```
go env -w GOPATH=
export GOPATH=

//关闭gomoudle, 要不然gopath不生效
export GO111MODULE=off

//需要在gopath目录下建立一个文件夹src

//下载依赖库 usage: go get [-d] [-t] [-u] [-v] [build flags] [packages], 下载包的其他依赖
go get -u go.uber.org/zap

//存放的目录就是在 $GOPATH/src/go.uber.org/zap下
//如果两个项目依赖不同的zap,需要在项目目录下新建vendor目录,把对应zap版本放在对应的vendor目录下

//寻找目录结构: vendor tree => goroot => gopath 
// 第三方依赖管理工具: glide、 dep、 go dep 
```

### 5-3 go mod的使用
- go.mod文件管理
- go命令统一管理, 用户不必关闭目录结构
- 初始化: go mod init 
- 增加依赖: go get ,更新依赖go get [@v...], go mod tidy 去除无关依赖
- 项目迁移到go mod: go moid init , go build ./...

```
module gomodtest
go 1.13

require (
    go.uber.org/zap v1.12.0
)
```

```
go get -u go.uber.org/zap@v1.12.0
//依赖库存放目录: $GOPATH/pkg/mod/go.uber.org/zap@v1.12.0
//会把其他依赖也放在go.mod文件中, 这样就有依赖版本管理功能了

go get -u go.uber.org/zap@v1.11.0

//如果需要升级, 不加任何版本, 就会更新到最新
go get -u go.uber.org/zap

//增加依赖
go get 
在对应的文件中import 

//旧项目迁移到go mod
把项目的依赖库都去掉

//触发编译当前目录及所有子目录, 遇到import就会触发go mod
go build ./...
```

### 5-4 目录的整理

> 之前是所有的go文件放在一个文件夹下,每个go文件都有main方法, go build 编译包下所有的go文件时,就会报错,那就需要把每个有main方法的go文件放在单独的文件夹中. 

