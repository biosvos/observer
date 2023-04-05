package observer

type Observer[T any] interface {
	Notify(subject *Subject[T])
}
