package observer

type Subject[T any] struct {
	observers []Observer[T]
	value     T
}

func NewSubject[T any](value T) *Subject[T] {
	return &Subject[T]{
		value: value,
	}
}

func (s *Subject[T]) RegisterObserver(observer Observer[T]) {
	s.observers = append(s.observers, observer)
}

func (s *Subject[T]) Set(value T) {
	s.value = value
	for _, observer := range s.observers {
		observer.Notify(s)
	}
}
