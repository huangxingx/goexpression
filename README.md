## golang 实现的表达式计算

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/huangxingx/goexpression/Go)
![Codecov](https://img.shields.io/codecov/c/github/huangxingx/goexpression)
![GitHub](https://img.shields.io/github/license/huangxingx/goexpression)

### 功能清单

1. [x] 四则运算 +-*/ 实现；
2. [x] 逻辑运算 and or && || 实现；
3. [x] 参数传递；
4. [x] 逻辑运算符 > >= < <= ! 实现；
5. [ ] 自定义方法注入；

### TODO

1. 优雅的错误返回；
2. 自定义方法注入；

### 使用

```bash
go get -u github.com/huangxingx/goexpression
```

`cat example.go`

```go
package main

import (
	"fmt"

	"github.com/huangxingx/goexpression"
)

func main() {
	expressStr := "1+2+3"
	exp := goexpression.NewExpress(expressStr)
	result := exp.Execute(nil)
	fmt.Printf("%s = %.2f", expressStr, result.(float64))
}
```

执行 `go run example.go`

```bash
1+2+3 = 6.00
Process finished with the exit code 0
```

**文件树**

```shell
➜  goexpression (main) ✗ tree                              
.
├── LICENSE
├── README.md
├── example
│   └── example.go
├── expression.go
├── expression_test.go
├── go.mod
├── operate
│   ├── add_operate.go
│   ├── and_operate.go
│   ├── div_operate.go
│   ├── logic_operate.go
│   ├── multi_operate.go
│   ├── operate.go
│   ├── or_operate.go
│   └── sub_operate.go
├── parse.go
├── parse_test.go
└── stack.go
```

`go test .`<p>

```shell
➜  goexpression (main) ✗ go test .       
ok      github.com/huangxingx/goexpression      0.167s

```

`go test -bench .`

```shell
➜  goexpression (main) ✗ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/huangxingx/goexpression
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAdd-12          1364682               864.0 ns/op
BenchmarkDiv-12           975276              1144 ns/op
PASS
ok      github.com/huangxingx/goexpression      3.555s

```