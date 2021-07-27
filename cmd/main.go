package main

import (
	"fmt"
	"goset/sethelper"
	"time"
)

// type helper interface {
// 	Runner1(a int64, b string, c float64)
// 	Runner2(a int64, b string, c float64)
// }

type Test struct {
	A int
	B string
	T *Test
}

type MyTypeForCheck int64

func main() {
	// wg := new(sync.WaitGroup)
	// // wg.Add(2)
	// dispatch.Dispatch(wg, Runner1, 12, "qwerty", 12.12)
	// dispatch.Dispatch(wg, Runner2, 34, "asdfgh")
	// wg.Wait()

	// a := []MyTypeForCheck{1, 2, 3, 4, 5}
	a := []interface{}{33, "eyert", 56.78}
	p := 56.78

	// b := []int{1, 2, 4}
	// var p MyTypeForCheck = 6
	// A := &Test{}
	// B := &Test{
	// 	T: A,
	// }
	// A.T = B
	// t1 := &Test{
	// 	A: 1,
	// 	B: "qwe",
	// }
	// t2 := &Test{
	// 	A: 2,
	// 	B: "asd",
	// }
	// A = append(A, t1, t2)

	n, err := sethelper.GetIndexOF(p, a)
	if err != nil {
		fmt.Printf("[error] %v\n", err)
	} else {
		fmt.Printf("index = %v\n", n)
	}

	// if reflect.DeepEqual(a, a) {
	// fmt.Println(true)
	// }

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
