package main

import (
	"Advance/structpk"// when import local package, the name of the package must be lower-case
	"Advance/syncpk"
	"fmt"
)
func main(){
	fmt.Println("程序开始")
	for{
		fmt.Printf("1: download task\n2: struct task\n3: eixt program\n")
		fmt.Println("Please input your cmd num:")
		var num int = 0
		fmt.Scanf("%d", &num)		
		fmt.Println()
		switch num {
		case 1:
			syncPk.MultiDownload()
		case 2:
			structPk.Structfunc()
		default:
		}
		if num != 1 && num != 2 && num != 0{
			break
		}
		fmt.Println()
	}
	fmt.Println("程序结束")
}