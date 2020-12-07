package main

import (
	"advance/structPk"// when import local package, input the module_name/dir_name
	"advance/syncPk"
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