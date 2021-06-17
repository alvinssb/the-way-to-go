package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of int is: ", size)
}
