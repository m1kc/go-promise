package quad

import (
	"aminus"
)

// All waits for all promises to be settled and returns their values as array
// or an error, if any of them was rejected.
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
