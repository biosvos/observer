package observer

type Observer struct {
	value int
}

func (o *Observer) Notify(subject *Subject) {
	o.value = subject.value
}

func NewObserver() *Observer {
	return &Observer{}
}
