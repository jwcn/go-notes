package main

import (
	"fmt"
)

// 第二条和第三条总结下来，就是 defer 后如果跟函数就是读取返回值，如果直接跟语句表达式就是读取当前行的值

func main() {
	// 1. 加了 defer 后倒序打印，因为 defer 执行顺序是先进后出
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // Output: 4 3 2 1 0
	}
	defer fmt.Println("===============")

	// 2. defer 声明时，参数就已经实时解析，不会在 return 的时候再解析取值
	j := 0
	defer fmt.Println(j) // Output: 0
	j++

	foo()
}

// 3. defer 可以读取有名的函数返回值
func foo() (i int) {
	i = 3
	defer func() {
		i++
		// 1) defer 的函数解析了并没有执行
		// 2) 函数读取了返回值 1
		// 3) 变相=修改了函数返回值，达到了另一个效果
		fmt.Println(i) // Output: 2
	}()
	// 这个才输出 3 对应第二条规则
	defer fmt.Println(i) // Output: 3
	return 1
}

// FileOutput:
// 3
// 2
// 0
// ===============
// 4
// 3
// 2
// 1
// 0
