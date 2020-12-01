package syncPk

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch = make(chan string, 10)// channel communication between routines

func MultiDownload(){// public function
	for i :=0; i < 3; i++{
		wg.Add(1)
		go singleDownload("www.vector.cn/" + string(i +'0') + ".zip", i)
	}
	for i :=0; i < 3; i++{
		msg := <- ch
		fmt.Println("msg is:", msg)
	}
	wg.Wait()	
	fmt.Println("Done!")
}

func singleDownload(url string, dur int){// private function
	ch <- string("Msg is: " + string(dur + '0'))
	fmt.Println("Start to download", url)
	time.Sleep((time.Second) * time.Duration(dur))
	fmt.Println("Finished to download", url)
	wg.Done()
}