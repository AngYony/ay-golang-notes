# Encoding格式检测其他编码格式转换为UTF8

### GBK转换为UTF8

引入`golang.org/x/text`包

```go
import "golang.org/x/text/transform"
...
utf8Reader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
all, err := ioutil.ReadAll(utf8Reader)
```



### 其他编码格式转换为UTF8的通用代码

需要安装以下两个包：

```
go get golang.org/x/text
go get golang.org/x/net/html
```

导入包：

```go
import (
	...
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)
```

核心代码：

```go
// 检测Reader中的编码格式并返回
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func main() {
	response, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error:status code", response.StatusCode)
		return
	}

	bodyReader := bufio.NewReader(response.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)

}
```

