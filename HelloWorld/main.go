package main
/*
#include <stdio.h>
int add(int a, int b)
{
	return (a + b);
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Hello World")
	var a int = 15
	b := 16
	var c C.int = C.int(b)
	fmt.Println(C.add(c,21))
	fmt.Println( "a + b = ", (a+b) )
}
