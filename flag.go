package main

import (
	"flag"
	"fmt"
)

var (
	foo = flag.String("foo", "value", "desc")	// 返回的是内存地址
)

func init() {
	flag.Parse()
}

func main() {

	fmt.Println(*foo)
}


/*
➜ default go run test.go
value
➜ default go run test.go -foo aaa
aaa
*/
