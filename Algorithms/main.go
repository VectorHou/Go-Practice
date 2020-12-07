package main

import(
	"fmt"
	"math/rand"
	"time"
	"Algorithms/Sort"
)

func main(){
	fmt.Println("Let's go to the world of algorithm")
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}	
	sort.BubbleSort(a)	
}	