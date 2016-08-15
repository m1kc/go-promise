package promise

type Promise interface {
	Wait() (interface{}, error)

	Fulfill(interface{})
	Reject(error)

	// TODO: get state
}
