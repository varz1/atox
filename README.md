Generic wrapper around strconv.Parse* forked from [josharian/atox](https://github.com/josharian/atox)

1. N: turn into anything from string ( or panic ) 转换字符串内容为指定类型
2. F: turn anything into string 将任何内容转换为字符串

simple use: 
```go
one := atox.MustN[int64]("1")

two := atox.F(1)

fmt.Println(reflect.TypeOf(one))
fmt.Println(two)
```
