package promise

import ()

func New() *P {
	return &P{
		valueChan: make(chan interface{}, 1),
		errChan:   make(chan error, 1),
	}
}

func All(arg ...*P) (ret []interface{}, err error) {
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
