package main

// struct 自动生成假数据

import (
	"fmt"
	"github.com/bxcodec/faker"
)

type T1 struct {
	A1 string
	A2 T2
}

type T2 struct {
	B1 int
}

func main() {
	var foo T1
	err := faker.FakeData(&foo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(foo)
	fmt.Printf("%+v", foo)
}

/**
Output:
{VZJYrVcczMndhBskifpLECVew {6107950388915260082}}
{A1:VZJYrVcczMndhBskifpLECVew A2:{B1:6107950388915260082}}[Finished in 0.5s]
*/
