// Copyright 2023 xfy150150@gmail.com. All rights reserved.
// @Time        : 2023/11/3 10:54
// @Author      : Createitv
// @FileName    : main.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

package main

import (
	"fmt"
	"github.com/createitv/knifes/containers/lists/arraylist"
)

func main() {
	c := arraylist.New("array", "array", "demo", "arraylist")
	fmt.Println(c.String())
}
