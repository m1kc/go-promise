package quad

import (
	"aminus"
)

func All(arg ...aminus.Promise) (ret []interface{}, err error) {
	ret = make([]interface{}, 0, len(arg))
	for _, p := range arg {
		result, err := p.Wait()
		if err != nil {
			return ret, err
		}
		ret = append(ret, result)
	}
	return
}

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
