package dispatch

import (
	"log"
	"reflect"
	"sync"
)

// Dispatch asynchronous function start with arguments
func Dispatch(wg *sync.WaitGroup, fu interface{}, args ...interface{}) {
	fuVal := reflect.ValueOf(fu)

	if fuVal.Kind() == reflect.Func {
		refArgs := make([]reflect.Value, fuVal.Type().NumIn())

		if fuVal.Type().NumIn() != len(args) {
			log.Println("[error] number args not valid")
			return
		}

		for i, arg := range args {
			if arg != nil {
				argVal := reflect.ValueOf(arg)

				if argVal.Type().ConvertibleTo(fuVal.Type().In(i)) {
					refArgs[i] = argVal.Convert(fuVal.Type().In(i))
				} else {
					log.Printf("[error] arg: %v cannot convert to: %v\n", argVal.Type(), fuVal.Type().In(i))
					return
				}

			} else {
				log.Println("arg is nil")
				return
			}
		}
		wg.Add(1)
		go func() {
			fuVal.Call(refArgs)
			wg.Done()
		}()

	} else {
		log.Println("[error] arg not a function")
		return
	}
}
