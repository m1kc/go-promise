package promise

import ()

type P struct {
	settled bool

	value interface{}
	err   error

	subs []chan interface{}
}

// New returns a new promise in pending state.
func New() *P {
	return &P{
		subs: make([]chan interface{}, 0, 1),
	}
}

// When implements the corresponding method required by aminus.Promise iface.
// Its behaviour is defined by the A- spec.
func (p *P) When() chan interface{} {
	ch := make(chan interface{})
	p.subs = append(p.subs, ch)
	return ch
}

// Wait implements the corresponding method required by aminus.Promise iface.
// Its behaviour is defined by the A- spec.
func (p *P) Wait() (ret interface{}, err error) {
	// If settled, return immediately
	if p.settled {
		return p.value, p.err
	}

	// If not, wait for fulfill/reject
	<-p.When()
	return p.value, p.err
}

func (p *P) settle() {
	p.settled = true
	for _, ch := range p.subs {
		ch <- nil
	}
}

// Fulfill implements the corresponding method required by aminus.Promise iface.
// Its behaviour is defined by the A- spec.
func (p *P) Fulfill(arg interface{}) {
	if p.settled {
		return
	}
	p.value = arg
	p.settle()
}

// Reject implements the corresponding method required by aminus.Promise iface.
// Its behaviour is defined by the A- spec.
func (p *P) Reject(arg error) {
	if p.settled {
		return
	}
	p.err = arg
	p.settle()
}
