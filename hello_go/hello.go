package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello misaki!")
	fmt.Println(sortFunc())
}

// 排序尝试
func sortFunc() string {
	a := "misaki"
	s := []byte(a)                                            // 转为Slice
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 用sorted的包，后面是比较函数
	sortedStr := string(s)
	return sortedStr
}
