package utils

import (
	bee "github.com/beego/beego/v2/server/web"
)

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func until(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func InitTemplate() {
	bee.AddFuncMap("add", add)
	bee.AddFuncMap("sub", sub)
	bee.AddFuncMap("until", until)
}
