package main

import (
	"aminus"
	"bird"
	"quad"

	"fmt"
	"time"
)

func longOperation(n uint64) (ret aminus.Promise) {
	ret = promise.New()
	go func() {
		for i := uint64(0); i < n; i++ {
			<-time.After(time.Second)
		}
		ret.Fulfill(n*10 + n)
	}()
	return
}

func main() {
	go func() {
		for i := 1; true; i++ {
			<-time.After(1 * time.Second)
			fmt.Printf("... %d sec\n", i)
		}
	}()

	p1 := longOperation(3)
	p2 := longOperation(2)

	values, err := quad.All(p1, p2)
	if err != nil {
		return
	}

	v1 := values[0].(uint64)
	v2 := values[1].(uint64)

	fmt.Printf("Value 1 is %v\n", v1)
	fmt.Printf("Value 2 is %v\n", v2)
	fmt.Printf("Sum is %v\n", v1+v2)

	p3 := longOperation(3)
	p4 := longOperation(2)
	first, err := quad.Race(p3, p4)
	if err != nil {
		return
	}

	fmt.Printf("First of 2 is %v\n", first)

	p5 := longOperation(2)
	go func() {
		val, err := p5.Wait()
		if err != nil {
			return
		}
		fmt.Printf("Routine 1: %v\n", val)
	}()
	go func() {
		val, err := p5.Wait()
		if err != nil {
			return
		}
		fmt.Printf("Routine 2: %v\n", val)
	}()
	v5, err := p5.Wait()
	if err != nil {
		return
	}
	fmt.Printf("Main thread: %v\n", v5)

	m := make(map[string]aminus.Promise)
	m["first"] = longOperation(1)
	m["second"] = longOperation(2)
	m["third"] = longOperation(3)
	m["forth"] = longOperation(4)

	m2, err := quad.Map(m)
	if err != nil {
		return
	}
	for key, value := range m2 {
		fmt.Printf("%v -> %v\n", key, value)
	}
}
