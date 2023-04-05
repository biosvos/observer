package observer

type Subject struct {
	observers []*Observer
	value     int
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) RegisterObserver(observer *Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Set(value int) {
	s.value = value
	for _, observer := range s.observers {
		observer.Notify(s)
	}
}
