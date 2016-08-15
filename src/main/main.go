package main

import (
	"promise"

	"fmt"
	"time"
)

func longOperation(n uint64) (ret *promise.P) {
	ret = promise.New()
	go func() {
		<-time.After(2 * time.Second)
		ret.Fulfill(n*10 + n)
	}()
	return
}

func main() {
	go func() {
		i := 0
		for {
			i++
			<-time.After(1 * time.Second)
			fmt.Printf("... %d sec\n", i)
		}
	}()

	p1 := longOperation(1)
	p2 := longOperation(3)

	v2, err := p2.Wait()
	if err != nil {
		return
	}
	vv2 := v2.(uint64)

	fmt.Printf("Value 2 is %v\n", vv2)

	v1, err := p1.Wait()
	if err != nil {
		return
	}
	vv1 := v1.(uint64)

	fmt.Printf("Value 1 is %v\n", vv1)

	fmt.Printf("Sum is %v\n", vv1+vv2)
}
