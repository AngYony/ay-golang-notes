# Go - GRPC

关于如何使用RPC进行方法调用，可参见源码：\src\imooc.go_dev_engineer_part2\rpc_start





RPC：Remote Procedure Call，远程过程调用，简单的理解就是一个节点请求另一个节点提供的服务。







## 开发环境配置

### 安装 Protocol Buffers（protoc）

GitHub地址：https://github.com/protocolbuffers/protobuf

下载链接：https://github.com/protocolbuffers/protobuf/releases/tag/v3.13.0

### 配置环境变量

需要配置的环境变量有：

GOPATH：C:\Users\AngYony\go

在Path变量后面追加GOPATH下的bin目录和Protobuf的bin目录：

D:\MyProgramFiles\protoc\protoc-3.13.0-win64\bin

C:\Users\AngYony\go\bin

在cmd中输入如下命令验证是否安装成功：

```powershell
protoc
```

### 安装go语言插件

进入到go.mod所在的目录，执行下述命令：

mac os：

```
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go
```

windows：

```
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get github.com/golang/protobuf/protoc-gen-go
```



## 创建proto文件

### .proto文件介绍

```protobuf
syntax = "proto3";
package coolcar;  //proto文件的package
option go_package="coolcar/proto/gen/go;trippb"; //最终生成的包路径

message Location{
    double latitude=1;
    double longitude=2;
}
//定义枚举类型
enum TripStatus{
    TS_NoT_SPECIFIED=0;
    NOT_STARTED=1;
    IN_PROGRESS=2;
}

message Trip{
    string start=1;
    string end=2;
    int64 duration_sec=3;
    int64 fee_cent=4;
    Location start_pos=5; //复合数据类型
    Location end_pos=6;
    repeated Location path_locations=7;  //集合数据类型
    TripStatus status=8;

}

```

### 使用protoc命令生成go文件

参考链接：https://github.com/golang/protobuf

参考链接：https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go

```powershell
protoc --go_out=paths=source_relative:gen/go  .\trip.proto
```

`:`后面表示的的是路径。

示例二：

```
protoc -I . helloworld.proto --go_out=plugins=grpc:.
```

命令用法说明：

`-I`：表示路径在什么位置之下，上面命令描述的是helloworld.proto文件在当前目录（.）下。

`plugins=grpc`表示运行插件protoc-gen-go。

`:`后面表示的的是路径，这里是生成到当前路径下（.）。





### 编写go文件

```go
package main

import (
	trippb "coolcar/proto/gen/go" 
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "abc",
		End:         "def",
		DurationSec: 3600,
		FeeCent:     10000,
		StartPos: &trippb.Location{
			Latitude:  35,
			Longitude: 100,
		},
		EndPos: &trippb.Location{
			Latitude:  40,
			Longitude: 130,
		},
		PathLocations: []*trippb.Location{
			{
				Latitude:  50,
				Longitude: 100,
			},
			{
				Latitude:  66,
				Longitude: 77,
			},
		},

		Status: trippb.TripStatus_IN_PROGRESS,
	}
	fmt.Println(&trip)
	//将类型编码为二进制流
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X \n", b)

	var trip2 trippb.Trip
	//将二进制编码进行解码
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}

	fmt.Println(&trip2)

	b, err = json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
```



## RPC 模式

- 简单模式（Simple RPC）：这种模式最为传统，即客户端发起一次请求，服务端响应一个数据。
- 服务端数据流模式（Server-side streaming RPC）：这种模式是客户端发起一次请求，服务返回一段连续的数据流。例如，客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端。
- 客户端数据流模式（Client-side streaming RPC）：与服务端数据流模式相反，这次是客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。例如，物联网终端向服务器报送数据。
- 双向数据流模式（Bidirectional streaming RPC）：客户端和服务端都可以向对方发送数据流，这个时候双方的数据可以同时互相发送，也就是可以实现实时交互。例如，聊天机器人。



























