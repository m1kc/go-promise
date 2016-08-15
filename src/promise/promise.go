package promise

import ()

type P struct {
	settled bool

	value interface{}
	err   error

	valueChan chan interface{}
	errChan   chan error
}

func New() *P {
	return &P{
		valueChan: make(chan interface{}, 1),
		errChan:   make(chan error, 1),
	}
}

func (p *P) Wait() (ret interface{}, err error) {
	// If settled, return immediately
	if p.settled {
		return p.value, p.err
	}

	// If not, wait for fulfill/reject

	select {
	case p.value = <-p.valueChan:
		// do nothing
	case p.err = <-p.errChan:
		// do nothing
	}

	p.settled = true
	return p.value, p.err
}

func (p *P) Fulfill(arg interface{}) {
	p.valueChan <- arg
}

func (p *P) Reject(arg error) {
	p.errChan <- arg
}
