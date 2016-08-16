package quad

import (
	"aminus"
)

// Race returns value or reason for the first promise to be settled,
// ignoring and not waiting for the others.
func Race(arg ...aminus.Promise) (ret interface{}, err error) {
	ch := make(chan aminus.Promise, len(arg))
	for _, p := range arg {
		go func(p aminus.Promise) {
			<-p.When()
			ch <- p
		}(p)
	}
	first := <-ch
	return first.Wait()
}
