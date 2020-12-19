
# API设计

## API Define
API
## API Errors
### 1.API响应处理
#### 错误习惯
大部分开发习惯在处理返回给页面的错误时，都是定义一个结构体，包装错误信息，然后HTTP响应码基本上都是`200`，如下：
```go
import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	CodeUserNotExit = 10001
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func TestResErr(t *testing.T) {
	res := Response{
		Code: CodeUserNotExit, // 用户不存在业务状态码
		Msg:  "用户不存在",
		Data: nil,
	}

	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, res) // 正常http响应
	})
	t.Logf("res:%v\n", res)

	r.Run()

}
```
这样做很不好：
>  上层运维监控系统，调用你提供的接口时，永远都是200。为了将正确的服务指标报给普罗米修斯等平台，还要解析响应内容，看Code字段和msg字段内容，额外增加了工作量，运维人员无法做统一的解析，也不好扩展。因为不同部门，不同系统定义的`Response`并非完全一致。

#### 正确处理：使用一小组标准错误配合大量资源

[HTTP错误码整理](https://blog.csdn.net/elevenqiao/article/details/6763040)

grpc 标准错误：

![](../image/01_grpc%20标准错误码.png)

尽可能把通用错误映射成标准错误码中的一类，比如 `找不到资源`类可以统一使用`404`类。

但是又衍生处一个问题：**如何区分用户找不到、邮箱不存在、订单不存在这些错误？**

 设计成如下：
```

```
### 2.API错误传播处理
如果您的 API 服务依赖于其他服务，则不应盲目地将这些服务的错误传播到您的客户端。在翻译错误时，我们建议执行以下操作：

隐藏实现详细信息和机密信息。

调整负责该错误的一方。例如，从另一个服务接收 INVALID_ARGUMENT 错误的服务器应该将 INTERNAL 传播给它自己的调用者。

#### 错误习惯

**直接透传错误**。比如某个BFF层调用UserService的服务，直接将错误return了。这样做会导致，写API文档时，讲不清API具体会返回什么错误码，因为都是直接透传err。如果调用多个service，不同service的错误码可能会重叠，无法让上游明白到底是什么错误。

当然有解决办法，就是使用全局错误码，比如：
- 账号组使用 `1XXXX`
- 订单组使用 `2XXXX`
- ....

对于以前的单体巨石架构，这样做是可以的，因为在一个项目里。

但是面对分布式系统，全局错误码是松散的、易被破坏契约的。系统多了，服务多了，定义的错误码就会增多，不能保证其他部门就非得使用你定义好的错误码，并且容易引用错误。


#### 正确处理
基于我们上述讨论的，在每个服务传播错误的时候，做一次翻译，这样保证每个服务 + 错误枚举，应该是唯一的，而且在 proto 定义中是可以写出来文档的。

在翻译错误时，有以下两个原则：
- 隐藏实现的详细信息和隐密信息。
- 调整负责该错误的一方。例如，从另一个服务接收 `INVALID_ARGUMENT` 错误的服务器应该将 `INTERNAL` 传播给上游

处理方式可以大致分两种：
1. 如果并不在乎下游的错误码，那么返回`UNKNOW`（500）错误码。
2. 如果在乎下游的错误码，那么需要翻译为自己自定义的错误码，然后返回

## API Design

推荐阅读 [Google API 设计指南](https://cloud.google.com/apis/design)

### Update 接口

### 建议update行为，对外只提供一个接口，入参为字段较全的结构体。
不推荐针对单个条件单独定义一个接口，防止接口定义过多。

但是会有一个问题：部分更新和全量更新的判断不同，针对部分字段的更新在调用时会增加很多判断。

**解决：** 使用`FieldMask`，是grpc提供的一个字段

`FieldMask` 部分更新的方案:
- 客户端可以执行需要更新的字段信息:
```go
paths: "author"
paths: "submessage.submessage.field"
```

示例： TODO
```go
service LibraryService {
rpc UpdateBook(UpdateBookRequest)returns(UpdateBookReply);
}
message UpdateBookRequest { 
    Book book = 1; 
    google.protobuf.FieldMask mask = 2;
}

message UpdateBookReply { Book book = 1; }

message Book {
string name = 1;
string author=2;
string title = 3;
}
```


**注意：** 空 `FieldMask` 默认应用到 “所有字段”
### 读写的请求和响应结构体推荐分开定义

分别定义`XXXRequest`和`XXXReply`

```go
service LibraryService {
rpc UpdateBook(UpdateBookRequest)returns(UpdateBookReply);
}
message UpdateBookRequest { Book book = 1; }

message UpdateBookReply { Book book = 1; }

message Book {
string name = 1;
string author=2;
string title = 3;
}
```