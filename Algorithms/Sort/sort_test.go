package sort

import(
	"testing"
	"math/rand"
	"time"
)

func TestBubbleSort(t *testing.T){
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}
	res, err := BubbleSort(a)
	if err != nil{
		t.Errorf("Execution function BubbleSort failed: %v", err)
	}else{
		item := res[0]
		for n, num := range(res){
			if item <= num{
				item = num
			} else {
				t.Errorf("Sort failed [%d].value(%d) > [%d].value(%d)\n", n-1, res[n-1], n, res[n])
				break;
			}
		}
	}	
}

func TestInsertionSort(t *testing.T){
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}
	res, err := InsertionSort(a)
	if err != nil{
		t.Errorf("Execution function InsertionSort failed: %v", err)
	}else{
		item := res[0]
		for n, num := range(res){
			if item <= num{
				item = num
			} else {
				t.Errorf("Sort failed [%d].value(%d) > [%d].value(%d)\n", n-1, res[n-1], n, res[n])
				break;
			}
		}
	}	
}

func TestSelectionSort(t *testing.T){
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}
	res, err := SelectionSort(a)
	if err != nil{
		t.Errorf("Execution function SelectionSort failed: %v", err)
	}else{
		item := res[0]
		for n, num := range(res){
			if item <= num{
				item = num
			} else {
				t.Errorf("Sort failed [%d].value(%d) > [%d].value(%d)\n", n-1, res[n-1], n, res[n])
				break;
			}
		}
	}	
}

func TestHeapSort(t *testing.T){
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}
	res, err := HeapSort(a)
	if err != nil{
		t.Errorf("Execution function HeapSort failed: %v", err)
	}else{
		item := res[0]
		for n, num := range(res){
			if item <= num{
				item = num
			} else {
				t.Errorf("Sort failed [%d].value(%d) > [%d].value(%d)\n", n-1, res[n-1], n, res[n])
				break;
			}
		}
	}	
}

func TestMergeSort(t *testing.T){
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}
	res, err := MergeSort(a)
	if err != nil{
		t.Errorf("Execution function MergeSort failed: %v", err)
	}else{
		item := res[0]
		for n, num := range(res){
			if item <= num{
				item = num
			} else {
				t.Errorf("Sort failed [%d].value(%d) > [%d].value(%d)\n", n-1, res[n-1], n, res[n])
				break;
			}
		}
	}	
}

func TestQuickSort(t *testing.T){
	var a []int
	rand.Seed(time.Now().Unix())
	num := rand.Intn(45)
	for i :=0; i < num; i++{
		a = append(a, rand.Intn(1000))
	}
	res, err := QuickSort(a)
	if err != nil{
		t.Errorf("Execution function QuickSort failed: %v", err)
	}else{
		item := res[0]
		for n, num := range(res){
			if item <= num{
				item = num
			} else {
				t.Errorf("Sort failed [%d].value(%d) > [%d].value(%d)\n", n-1, res[n-1], n, res[n])
				break;
			}
		}
	}	
}