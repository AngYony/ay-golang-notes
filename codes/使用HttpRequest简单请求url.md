# 使用HttpRequest简单请求url

这里使用的是一个第三方包。

```go
import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)
func main() {
	req := HttpRequest.NewRequest()
	res, err := req.Get("http://baidu.com")
	if err != nil {
		return
	}
	body, _ := res.Body()
	fmt.Println(string(body))
}
```

