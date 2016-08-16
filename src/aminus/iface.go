package aminus

type Promise interface {
	// When returns a channel which emits a single nil value when the promise
	// is settled.
	When() chan interface{}
	// Wait is a blocking call which returns when the promise is settled,
	// returning its value and reason.
	Wait() (interface{}, error)

	// Fulfill fulfills the promise with given value.
	Fulfill(interface{})
	// Reject rejects the promise with given reason.
	Reject(error)

	// TODO: get state
}
