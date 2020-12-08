package main

import (
	"fmt"
	"unsafe"
	"math"
	"Algorithms/Sort"
) 

type h =  int // means typedef in C++
type hk int // hk means a new type, but hk has the feature of int

func main() {
	fmt.Println("Hello World")
	aa := []int{1,2,3,4,5,6,7}
	sort.QuickSort(aa)
	var a h = 10
	var b int = 26
	var c float64
	var d string
	var e hk
	var f rune// rune is an alias for int32 and is equivalent to int32 in all ways
	c, d = sqr_sum(a,b)
	fmt.Println("sqr_sum(", a, ",", b, ")", "result:", c, "error:", d)
	a = -20
	b = 10
	c, d = sqr_sum(a,b)
	fmt.Println("sqr_sum(", a, ",", b, ")", "result:", c, "error:", d)

	strmsg := "中华人民共和国"
	fmt.Println("strmsg length:", len(strmsg))
	var slice []byte = []byte(strmsg)
	fmt.Println(slice, len(slice))
	var data [10]byte
	data[0] = 'T'
	data[1] = 'E'
	strmsg = string(data[:])
	fmt.Println(strmsg)

	//package unsafe
	fmt.Println( unsafe.Sizeof(f) )//unsafe.Sizeof return the number of bytes for var
	fmt.Printf( "var a h type is: %T\n", a )
	fmt.Printf( "var e hk type is: %T\n", e )
	fmt.Println( len(strmsg))

}

func sqr_sum(a int, b int) (ret float64, err string){// func: function keyword; (a int, b int): function argument; (int, int): return type, note golang allow multiple return value
	var sum float64 = float64(a + b)
	if sum > 0{
		return math.Sqrt( sum ), "sucess"
	}else{
		return 0, "failed"
	}
}