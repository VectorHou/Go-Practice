package main
/*
#include <stdio.h>
#include <string.h>
void printInt(int a)
{
	printf("Convert Go_int to C_int: %d\n", a);
}
void printString(char* string )
{
	printf("Convert Go_string to C_string: \"%s\"\n", string);
}
*/
import "C"
import "fmt"

func main() {
fmt.Println("Go数据类型转换为C数据类型")
	var a int = 255
	var msg string = "Hello World"
	
	C.printInt(C.int(a))//C.int(var_int) return int(C) 
	C.printString(C.CString(msg))//C.CString(var_string) return char*(C)
	
	fmt.Println("程序结束")
}