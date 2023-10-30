package main

import "fmt"

func main() {
	var a map[int]string

	a = make(map[int]string, 10)

	a[20095452] = "zhangsan"
	a[20095387] = "lisi"
	a[20097291] = "wangwu"
	value, flag := a[20095452]

	fmt.Println(a)
	fmt.Println(value, flag)

	delete(a, 20095387)

	fmt.Println(a)

	for k, v := range a {
		fmt.Println(k, v)
	}
}
