package main
/*
#include <stdio.h>
#include <string.h>

typedef char bool;

char cchar = 'M';
float cfloat = 3.1415926;
double cdouble = 3.1415927;
long clong = 1234567;
char* cstring = "c type string";
char cchararray[64] = "c type char array";

void printInt(int a)
{
	printf("Convert Go_int to C_int: %d\n", a);
}
void printString(char* string )
{
	printf("Convert Go_string to C_string: \"%s\"\n", string);
}
void printChar(char ch)
{
	printf("Convert Go_byte to C_char: %c\n", ch);
}
void printFloat(float flt)
{
	printf("Convert Go_float32 to C_float: %f\n", flt);
}
void printBool(bool bl)
{
	printf("Convert Go_bool to C_bool: %d\n", bl);
}
///////////////////////////////////////////////////
int add(int a, int b)
{
	return a + b;
}

*/
import "C"
import "fmt"

func main() {
fmt.Println("Go数据类型转换为C数据类型\n")
	var c1 byte = 'a'
	var b float32 = 12.2
	ok := true
	var c_type_bool byte = 0
	var a int = 255
	var msg string = "Hello World"
	
	C.printChar(C.char(c1))//C.char(var byte) return char(C)
	C.printFloat(C.float(b))//C.float(var float32) return float(C)
	if ok{
		c_type_bool = 1
	}else{
		c_type_bool = 0
	}	
	C.printBool(C.bool(c_type_bool))//C have no bool type, so use C.bool(var bool) return C type(function argument)
	C.printInt(C.int(a))//C.int(var_int) return int(C) 
	C.printString(C.CString(msg))//C.CString(var_string) return char*(C)

	fmt.Println("\nC数据类型转换为Go数据类型\n")
	var gochar byte
	var gofloat float32
	var godouble float64
	var golong int64
	var gostring string
	var gochararray []byte//slice 声明时，指定了长度的是array，未指定长度的是slice，
						  //array可以使用[i]赋值，slice使用append(var_name, value)增加元素

	a = int(C.add(255, 255))
	fmt.Println("C function add return:", a)
	gochar = byte(C.cchar)//byte(C_char) return var_byte
	fmt.Println("C type char:", gochar)
	gofloat = float32(C.cfloat)//float32(C_float) return var_float32
	fmt.Println("C type float:", gofloat)
	godouble = float64(C.cdouble)//float64(C_doulbe) return var_float64
	fmt.Println("C type double:", godouble)
	golong = int64(C.clong)//int64(C_long) return var_int64
	fmt.Println("C type long:", golong)
	gostring = C.GoString(C.cstring)//C.GoString(C_char*) return var_string
	fmt.Println("C type string:", gostring)
	for i := range C.cchararray{
		if C.cchararray[i] != 0{
			gochararray = append(gochararray, byte(C.cchararray[i]))
		}		
	}
	fmt.Println("C type char array:", string(gochararray))

	fmt.Println("\n程序结束!")
}