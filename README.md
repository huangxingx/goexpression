### golang 实现的表达式计算

## 功能清单

1. [x] 四则运算 +-*/ 实现；
2. [x] 逻辑运算 and or && || 实现；
3. [x] 参数传递；
4. [x] 逻辑运算符 > >= < <= ! 实现；
5. [ ] 自定义方法注入；

## TODO

1. 优雅的错误返回；
2. 自定义方法注入；

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
➜  go-utils (main) ✗ tree express              
express
├── README.md
├── express.go
├── express_test.go
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
└── parse_test.go
```