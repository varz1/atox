1. N: turn into anything from string or panic
2. F: turn anything into string

simple use: 
```go
half := atox.MustN[int64]("1")

one := atox.F(1)
```
