package main

import (
	"fmt"
	"goset/dispatch"
	"sync"
	"time"
)

// type helper interface {
// 	Runner1(a int64, b string, c float64)
// 	Runner2(a int64, b string, c float64)
// }

func main() {
	wg := new(sync.WaitGroup)
	// wg.Add(2)
	dispatch.Dispatch(wg, Runner1, 12, "qwerty", 12.12)
	dispatch.Dispatch(wg, Runner2, 34, "asdfgh")
	wg.Wait()

	// n, err := sethelper.GetIndexOF("wer", []string{"qwer", "wer", "tyu"})
	// if err != nil {
	// 	fmt.Printf("[error] %v\n", err)
	// }
	// fmt.Printf("[index] %v\n", n)

}

func Runner1(a int64, b string, c float64) {
	i := 0
	for {
		if i == 3 {
			break
		}
		i++
		fmt.Printf("run Runner1 [%v] %v, %v, %v\n", i, a, b, c)
		time.Sleep(time.Duration(time.Second))
	}
}

func Runner2(a int64, b string) {
	i := 0
	for {
		if i == 3 {
			break
		}
		i++
		fmt.Printf("run Runner2 [%v] %v, %v\n", i, a, b)
		time.Sleep(time.Duration(time.Second * 2))
	}
}
