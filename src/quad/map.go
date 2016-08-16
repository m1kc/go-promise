package quad

import (
	"aminus"
)

// Map returns values of all promises as a map. If one of the promises
// is rejected, it doesn't wait for the others.
func Map(arg map[string]aminus.Promise) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{})
	for key, p := range arg {
		val, err := p.Wait()
		if err != nil {
			return ret, err
		}
		ret[key] = val
	}
	return
}
