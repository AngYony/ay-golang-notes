# Go-时间与日期

时间与日期相关的操作，需要引入time包。

时间类型是time.Time。



示例一，获取当前时间：

```go
// 获取当前时间
now := time.Now()
fmt.Printf("type=%T value=%v \n", now, now)
// 获取年月日时分秒
fmt.Printf("年：%v \n", now.Year())
fmt.Printf("月(英文）：%v \n", now.Month())
fmt.Printf("月(中文）：%v \n", int(now.Month()))
fmt.Printf("日：%v \n", now.Day())
fmt.Printf("时：%v \n", now.Hour())
fmt.Printf("分：%v \n", now.Minute())
fmt.Printf("秒：%v \n", now.Second())
```

输出：

```
type=time.Time value=2021-12-13 17:33:03.1074455 +0800 CST m=+0.001720801
年：2021                                                   
月(英文）：December                                        
月(中文）：12                                              
日：13                                                     
时：17                                                     
分：33                                                     
秒：3   
```



示例二，格式化时间：

```go
// 格式化日期和时间
// 方式一：Printf()或者Sprintf()
fmt.Printf("当前格式化后的时间：%02d年%02d月%02d日 %02d:%02d:%02d",
	now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
fmt.Println()
// 方式二：Format，注意：这个时间对应的数字是固定写法不可调整，不能随便改为其他年份其他时间点，必须如下时间点
fmt.Println(now.Format("2006/01/02 15:04:05"))
fmt.Println(now.Format("2006年01月02日 15时04分05秒"))
// 获取年份，只能指定为2006
fmt.Println("年份：", now.Format("2006"))
// 获取月份，只能指定为01
fmt.Println("月份：", now.Format("01"))
// 获取小时，只能指定为15
fmt.Println("小时：", now.Format("15"))
```

输出：

```
当前格式化后的时间：2021年12月13日 17:33:03                
2021/12/13 17:33:03                                        
2021年12月13日 17时33分03秒                                
年份： 2021                                                
月份： 12                                                  
小时： 17                  
```

示例三，获取时间戳：

```go
// unix时间戳
fmt.Printf("unix时间戳：%v\tunixnano时间戳：%v", now.Unix(), now.UnixNano())
```

输出：

```
unix时间戳：1639387983  unixnano时间戳：1639387983107445500
```





## 时间与日期常用函数

| 函数          | 说明                                                      |
| ------------- | --------------------------------------------------------- |
| time.Sleep(t) | 休眠指定时长                                              |
| time.Unix     | Unix时间戳，从1970的UTC时间到时间点所经过的时间（单位秒） |
| time.UnixNano | Unix时间戳，单位纳秒                                      |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |
|               |                                                           |

